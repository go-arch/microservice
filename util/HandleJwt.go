package util

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"students/constants"
)

var (
	verifyKey *rsa.PublicKey
)
// jwt initialization 
func InitiateJwt() {

	verifyKeyPath, err := filepath.Abs(constants.PUB_KEY_PATH)

	verifyByte, err := ioutil.ReadFile(verifyKeyPath)
	Fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyByte)
	Fatal(err)

}

//ValidateMiddleware HI Function

/**

 */
func ValidateMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token is Not valid")
	} else {
		if token.Valid {
			ctx := context.WithValue(r.Context(), "id", token.Claims)
			next(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized) //TODO code repeated
			fmt.Fprint(w, "Token is Not valid")
		}
	}
}
