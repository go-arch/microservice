package services

import (
	
	"github.com/streadway/amqp"
	"students/util"
	"encoding/json"
	"students/models"
	"log"
	"students/constants"
	"fmt"
	"time"
	"os"

)

//Queue and channel declaration
var Q amqp.Queue
var CH *amqp.Channel

//initlization of queue
func InitMQ(){
	go (func() {
		var env = os.Getenv("GO_ENV");
		amqpUrl := "amqp://guest:guest@rabbit:5672/"
		if env == "local" {
			amqpUrl = "amqp://guest:guest@localhost:5672/"
		}
		time.Sleep(10 * time.Second)
		conn, err := amqp.Dial(amqpUrl)
		util.Fatal(err)
		ch, err := conn.Channel()
		CH = ch
		util.Fatal(err)
		q, err := ch.QueueDeclare(
			"userq", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		util.Fatal(err)
		Q = q;
		QSubscribe()
	})()

}


//Publish data on Queue
func QPublish(eventType string, content interface{})  {
	publishContent := models.MQPublish{EventType:eventType,Content:content}
	body,err := json.Marshal(publishContent)
	util.Fatal(err)
	err = CH.Publish(
		"",     // exchange
		Q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:body,
		})
	util.Fatal(err)
}

//Data subscription
func QSubscribe()  {
	go func() {
		//time.Sleep(10 * time.Second)
	var qPublish models.MQPublish
	msgs, err := CH.Consume(
		Q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.Fatal(err)
	//forever := make(chan bool)


		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		for d := range msgs {

			err := json.Unmarshal(d.Body,&qPublish)
			util.Fatal(err);
			switch qPublish.EventType {
			case constants.QUserCreated:
				{
					//Action On User Created
				}
			default:
				{
					fmt.Println("Event Recieved But NO Match")
				}

			}

		}
	}()

	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever
}