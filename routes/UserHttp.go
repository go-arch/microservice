package routes

import (
	"students/models"
	"students/controllers"
	"net/http"
	"encoding/json"
	"students/util"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func HandleUserSet(w http.ResponseWriter, r *http.Request) {
	//setConnection();
	uc := controllers.NewUserController();
	var user models.User
	decoder := json.NewDecoder(r.Body);
	err := decoder.Decode(&user);
	util.Fatal(err)
	//db.Create(&user)
	created := uc.CreateUser(&user, w)
	//userResponse := models.UserCreateResponse{created,id}
	//SendJsonResponseSuccess(w,userResponse);
	SendJsonResponseSuccess(w, created);
}

func HandleUserGet(w http.ResponseWriter, r *http.Request) {
	//setConnection();
	uc := controllers.NewUserController();
	var user models.User
	user.Email = r.URL.Query().Get("email");
	_, err := uc.GetUser(&user);
	if (err != nil) {
		Send404(w)
		return;
	}
	SendJsonResponseSuccess(w, user);


}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	uc := controllers.NewUserController();
	var user models.GetAllUser
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		SendJsonResponseError(w);
		return
	}
	userList, err := uc.GetBulkUsers(user)
	if err != nil {
		SendJsonResponseError(w);
		return;
	}
	SendJsonResponseSuccess(w, userList)
	return;


}

func deleteStudnetHandler(w http.ResponseWriter, r *http.Request) {
	//setConnection();
	uc := controllers.NewUserController();
	log.Printf("calling delete Studnet Handler function")
	userId := r.URL.Query().Get("id")
	deletedUser := uc.DeleteUser(userId)
	SendJsonResponseSuccess(w, deletedUser);

}

func showStudentHandler(w http.ResponseWriter, r *http.Request) {
	//setConnection();
	uc := controllers.NewUserController();
	log.Printf("calling delete Studnet Handler function")
	userId := r.URL.Query().Get("id")
	user, err := uc.ShowUser(userId)
	if (err != nil) {
		SendJsonResponseError(w)
	}
	SendJsonResponseSuccess(w, user);
}

func updateStudentHandler(w http.ResponseWriter, r *http.Request) {
	uc := controllers.NewUserController();

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		SendJsonResponseError(w);
		return
	}
	err := uc.UpdateUser(&user)
	if (err != nil) {
		SendJsonResponseError(w)
	}
	SendJsonResponseSuccess(w, user);
}









