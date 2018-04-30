## SMIS Base Repo

### Dependency

> Golang

> RabbitMQ
## ubuntu docker installation script

    #!/bin/bash

    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
         sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
     sudo apt-get update
     apt-cache policy docker-ce
     sudo apt-get install -y docker-ce
     sudo systemctl status docker

        sudo chmod +x /usr/local/bin/docker-compose
     sudo curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
     sudo chmod +x /usr/local/bin/docker-compose


## Simple setup using Docker
    docker-compose build
    docker-compose up




### SetUp
 1. Git clone repo
 2. do
 `bash
 go get ./
 `
 3. To run it
 `bash
 go run main.go
 `
## functionality
 1. Http Api
 2. RPC for internal One way communication
 3. MQ for One way message passing

## Example
 client Example to call it

     var settings = {
       "async": true,
       "crossDomain": true,
       "url": "http://localhost:1234/user",
       "method": "POST",
       "headers": {
         "Content-Type": "application/json",
         "Cache-Control": "no-cache",
         "Postman-Token": "b669a328-b1b1-c263-1625-17cc0bc764b5"
       },
       "processData": false,
       "data": "{\n\t\"name\" : \"Rahul Kamboj3\",\n\t\"email\" : \"kamo3.rahul@gmail.com\",\n\t\"password\" : \"1234567788\",\n\t\"phoneNumber\" : \"+911234567890\"\n}"
     }

     $.ajax(settings).done(function (response) {
       console.log(response);
     });


## Example DB Table

        create database user_ms;
        use user_ms;

        create table users(`id` INT(10) NOT NULL AUTO_INCREMENT,
                `name` VARCHAR(64) NULL DEFAULT NULL,
                `email` VARCHAR(64) NULL DEFAULT NULL,
                `password` VARCHAR(64) NULL DEFAULT NULL,
                `phoneNumber` VARCHAR(64) NULL DEFAULT NULL,
                PRIMARY KEY (`id`))


## RPC Testing

    package main

    import (
        "net/rpc"
        "log"
        "fmt"
    )



    type User struct {
        Id string `json:"id"`
        Name string `json:"name"`
        Email string `json:"email"`
        Password string `json:"password"`
        PhoneNumber string `json:"phoneNumber"`
    }
    func main() {
        client, err := rpc.DialHTTP("tcp", "localhost" + ":1234")
        if err != nil {
            log.Fatal("dialing:", err)
        }// Synchronous call
        args := "kamo.rahul@gmail.com"
        var reply User
        err = client.Call("Usr.GetUser", args, &reply)
        if err != nil {
            log.Fatal("arith error:", err)
        }
        fmt.Println(reply)
    }


## UnderDevelopment
 > RPC client Support -- Done

 > MQ client  -- Done

## Directory structure explanation

## Format supported by Golang For Date

    ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"

### Sample Code for date conversion
    func GetSpanishDate(ts time.Time) string{
        location,err := time.LoadLocation("Chile/Continental");
        if err != nil{
            log.Fatal("Issue converting time")
        }
        return ts.In(location).Format("02/01/2006 15:04:05 MST")
    }

### sample code for date parsing
    format := "02/01/2006 15:04:05 MST"
    dateString := "02/12/2018 15:04:05 MST"
    time,err := time.Parse(format,dateString)

### Goplayground link for date handling

https://play.golang.org/p/9pXTkhFuwDA

### Goplay url
Made With Code by Rahul Kamboj