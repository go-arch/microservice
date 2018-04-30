/*
user services
*/

package services

import (
	"students/models"
	"students/util"
	"fmt"
)

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

// get user from database
func (us *UserService) SUserGet(user *models.User) (err error) {

	fmt.Println("Reached")

	err = DB.Where("email = ?", user.Email).Find(&user).Error
	fmt.Println(err)
	fmt.Println(user)

	util.Fatal(err)

	return err
}

// check user exists in database
func (us *UserService) SUserExists(user *models.User) (exists bool) {

	//query := fmt.Sprintf(SqlQuery.UserExistsQuery, user.Name, user.Email)
	DB.Find(user, user.Email)

	//util.Fatal(err)
	//var count int
	//
	//for rs.Next() {
	//	rs.Scan(&count)
	//}
	//
	//if count == 1 {
	//	exists = true
	//} else {
	//	exists = false
	//}
	return
}

func (us *UserService) SDeleteUser(id string) (bool) {
	var user models.User
	err := DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return false
	} else {
		return true
	}

}

func (us *UserService) SShowUser(id string) (models.User, error) {
	var user models.User
	err := DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return models.User{}, err
	} else {
		return user, nil
	}
}

func (us *UserService) SUpdateUser(user *models.User) (error) {
	DB.Where("id = ?", user.ID).Find(&user, &user.ID)
	err := DB.Save(&user).Error
	fmt.Println(err)
	return err
}



// user create service
func (us *UserService) SCreateUser(user *models.User) error {
	err := DB.Create(&user).Error;
	return err
}

func (us *UserService) SGetBulkUser(q models.GetAllUser) (users []models.User,err  error) {
	allUser := []models.User{}
	err = DB.Find(&allUser).Offset(q.Skip).Limit(q.Limit).Error
	return allUser, err
}