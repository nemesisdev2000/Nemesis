package AuthenticationServer

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/nemesisdev2000/Nemesis/TeamServer/Jwt"
	"github.com/nemesisdev2000/Nemesis/TeamServer/UserDB"
)

var tokenList []string

func SignupHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Passwordhash missing"))
		return
	}

	//validate and then add the user

	check := UserDB.AddUserObject(r.Header["Username"][0], r.Header["Passwordhash"][0], 0)

	if !check {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username already exists"))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Client Added"))
}

func GetSignedToken() (string, error) {
	claimsMap := map[string]string{
		"aud": "dosxuz.gitlab.io",
		"iss": "gitlab.io",
		"exp": fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := "somepassword"
	header := "HS256"
	tokenString, err := Jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func ValidateUser(username string, passwordhash string) (bool, error) {
	user, exists := UserDB.GetUserObject(username)
	if !exists {
		return false, errors.New("Client does not exist")
	}
	passwordCheck := user.ValidatePasswordHash(passwordhash)
	if !passwordCheck {
		return false, nil
	}
	return true, nil
}

func SigninHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))
		return
	}

	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Passwordhash Missing"))
		return
	}

	valid, err := ValidateUser(r.Header["Username"][0], r.Header["Passwordhash"][0])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Client does not exist"))
		return
	}

	if !valid {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect password"))
		return
	}

	tokenString, err := GetSignedToken()
	tokenList = append(tokenList, tokenString)
	fmt.Println(tokenList)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
}

//this checks if a signed in user is sending the command or not
func ValidateRequest(token string) bool {
	flag := false
	for _, a := range tokenList {
		if a == token {
			flag = true
			break
		}
	}
	if flag == false {
		fmt.Println("Unauthorized request")
		return false
	} else {
		return true
	}
}
