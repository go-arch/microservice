package util

import (

	"fmt"
)

//print error in future we will revert the client with bad request
func Fatal(err error){
	if err != nil {
		fmt.Println(err)
	}
}


