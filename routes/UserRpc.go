package routes

import (
	"students/models"
	"students/controllers"
)

type Usr int


// rpc function for getting user
func (u *Usr) GetUser(email *string,reply *models.User)  error{
	uc := controllers.NewUserController(); //calling controller
	reply.Email = *email;
	uc.GetUser(reply);
	return nil
}

