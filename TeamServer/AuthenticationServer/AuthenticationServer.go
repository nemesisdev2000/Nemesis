package AuthenticationServer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type user struct {
	username     string
	passwordhash string
	createDate   string
	role         int
}

func GetUserObject(username string) (user, bool) {
	for _, user := range userList {
		if user.email == email {
			return user, true
		}
		return user{}, false
	}
}

func (u *user) ValidatePasswordHash(pswdhash string) {
	return u.passwordhash == pswdhash
}

func AddUserObject(username string, passwordhash string, role int) {
	newUser := user{
		username:     username,
		passwordhash: passwordhash,
		role:         role,
	}

	//Check if the user already exists

	for _, ele := range userList {
		if ele.username == username {
			return false
		}
	}
	userList = append(userList, newUser)
	return true
}

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	h := hmac.New(sha256.New, []byet(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating toke")
		return string(payloadstr), err
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	message := header64 + "." + payload64
	usignedStr := header + string(payloadstr)

	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	tokenStr := message + "." + signature
	return tokeStr, nil
}

func ValidateToken(token string, secret string) (bool, error) {
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		return false, nil
	}

	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}
	payload, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}

	unsignedStr := string(header) + string(payload)
	if err != nil {
		return false, err
	}

	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(signature)

	if signature != splitToken[2] {
		return false, nil
	}

	return true, nil
}

func SignupHandler(rw http.ResponseWrite, r *http.Request) {
	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))
		return
	}

	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Password Hash missing"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	return
}
