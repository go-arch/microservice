package controllers

import (
	"students/models"
	"students/services"
	"students/util"
	"students/constants"
	"net/http"
	"fmt"
)

type UserController struct {

}

func NewUserController() *UserController {
	return &UserController{}
}


// get user
func (uc *UserController) GetUser(user *models.User) (*models.User,error){
	us := services.NewUserService()
	if user.Email != "" {
		err := us.SUserGet(user)

		return user,err
	}
	return &models.User{},nil

}
// create user controller
func (uc *UserController) CreateUser(user *models.User,w http.ResponseWriter) (created bool) {
	us := services.NewUserService()
	if user.Email != "" && user.Name != "" {
		fmt.Println(1)
		exists := us.SUserExists(user)
		if exists {
			return false
		}
		fmt.Println(2)

		err := us.SCreateUser(user);
		if err != nil {
			util.Fatal(err)
		} else {
			created = true
		}
	}
	services.QPublish(constants.QUserCreated,user)
	return created
}

func (uc *UserController) DeleteUser(id string ) bool{
	us := services.NewUserService()
	deleted := us.SDeleteUser(id)

	//services.QPublish(constants.QUserCreated,user)
	return deleted
}

func (uc *UserController) ShowUser(id string )  ( models.User,error){
	us := services.NewUserService()
	user,err := us.SShowUser(id)


	//services.QPublish(constants.QUserCreated,user)
	return user,err
}


func (uc *UserController) UpdateUser(user *models.User )  ( error){

	us := services.NewUserService()
	err := us.SUpdateUser(user)


	//services.QPublish(constants.QUserCreated,user)
	return err
}

func (uc *UserController) GetBulkUsers	(user models.GetAllUser )  ( []models.User,error){

	us := services.NewUserService()
	allUser,err := us.SGetBulkUser(user)


	//services.QPublish(constants.QUserCreated,user)
	return allUser,err
}
