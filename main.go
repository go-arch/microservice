/*
Main go file
Go code start from here
*/

package main

//Package importing
import (
	"net/http"
	"students/util"
	"students/services"
	"students/routes"
	"net/rpc"
)

import (
   "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)
var db *gorm.DB //define the gobal variable for database connection with datatype.
var err error   //defind the gobal variable for dispaly the error.
// define the table struct name of table name with feild list along with data type and field name.

//main function call
func main() {

	util.InitiateJwt() //jwt certificate reading from file
	services.GetDbSession(); // initiating database connection
	//services.Migration();
	//db, err = gorm.Open("mysql", "userms:randompass@tcp(userms.ctm6uenobr4l.us-west-2.rds.amazonaws.com:3306)/userms?charset=utf8")
	//   defer db.Close()
	//if err != nil {
	//  fmt.Println(err)
	//}
	//
	// After the dataconnection successfully we are migrate table with database.
	// AutoMigrate(structName) that we have already defined the person2 struct with fieldname and datatype.
	//db.AutoMigrate(&models.User{})
	//services.Migration()
	services.InitMQ() // initiating rabbitmq connection
	//services.QSubscribe() // subscribe to queue for listning  one way message from other services
	routes.RegisterHttp() // register routers
	routes.InitilizeRoleMap() // Initialize roles
	routes.RegisterRpc() // register RPC functions
	rpc.HandleHTTP() // binding rpc with http
	fmt.Println("Listning At 3001")

	http.ListenAndServe(":3001", nil) // start listning to http server

}


