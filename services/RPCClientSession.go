package services

import (
	"net/rpc"

	"students/constants"
	"students/util"
)


//execute remote procedure
func ExecRPC(functionName string, payload interface{}, reply *interface{}){
	client, err := rpc.DialHTTP("tcp", constants.HOST + ":" + constants.PORT)
	util.Fatal(err)
	err = client.Call(functionName, payload, &reply)

}
