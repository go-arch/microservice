/*
main file for routing to initiate rpc and http connection

*/

package routes

import (
	"net/http"
	"students/models"
	"encoding/json"
	"students/util"
	"net/rpc"
	"github.com/kamorahul/rbac"
)

var RoleMapper *rbac.RoleMapper

func InitilizeRoleMap() {
	RoleMapper = rbac.NewRoleMapper()
	RoleMapper.AddMethodMapping("/user/get", "GET", []string{"student"})
	RoleMapper.AddMethodMapping("/user/all", "GET", []string{"student"})
	RoleMapper.AddMethodMapping("/user/delete", "GET", []string{"student"})
	RoleMapper.AddMethodMapping("/user/show", "GET", []string{"student"})
	RoleMapper.AddMethodMapping("/user/update", "GET", []string{"student"})

}

//register http routing
func RegisterHttp() {
	http.HandleFunc("/user/set", HandleUserSet)
	http.HandleFunc("/user/get", withRole(HandleUserGet))
	http.HandleFunc("/user/all", GetStudents)
	http.HandleFunc("/user/delete", deleteStudnetHandler)
	http.HandleFunc("/user/show", showStudentHandler)
	http.HandleFunc("/user/update", updateStudentHandler)

}

//Role

func withRole(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.URL.Query().Get("role");
		valid := ValidateRole(r.URL.Path,r.Method,role)
		if (valid) {
			next.ServeHTTP(w, r)
		} else {
			SendJsonResponseUnauthorize(w);
		}
	}
}

func ValidateRole(path string,method string,roleR string) bool {
	if RoleMapper.RoleMethodValid(path,method, roleR) {
		return true
	}
	return false
}

//register rpc function for two way communications
func RegisterRpc() {
	u := new(Usr)
	err := rpc.Register(u);
	util.Fatal(err)
}



// its a utility function for making standard output
func SendJsonResponseSuccess(w http.ResponseWriter, data interface{}) {
	var response models.ResponseModel
	response.Status = http.StatusOK;
	response.Data = data
	responseData, err := json.Marshal(response)
	util.Fatal(err);
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);
	w.Write(responseData);
}
func SendJsonResponseError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusInternalServerError);
	response := models.ResponseModel{Status : http.StatusInternalServerError, Data:http.StatusText(http.StatusInternalServerError)}
	res, err := json.Marshal(&response)
	util.Fatal(err)
	w.Write(res);
}
func SendJsonResponseUnauthorize(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusUnauthorized);
	response := models.ResponseModel{Status : http.StatusUnauthorized, Data:http.StatusText(http.StatusUnauthorized)}
	res, err := json.Marshal(&response)
	util.Fatal(err)
	w.Write(res);
}

// utility for sending 404 in response
func Send404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound);
	//w.Write([]byte(http.StatusText(http.StatusNotFound)));
	response := models.ResponseModel{Status : http.StatusNotFound, Data:models.EmptyArray()}
	res, err := json.Marshal(&response)
	util.Fatal(err)
	w.Write(res);
}