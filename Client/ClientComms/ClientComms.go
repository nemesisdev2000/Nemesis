package ClientComms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type userDetails struct {
	Username string
	Password string
}

type TcpListenerDetails struct {
	Type string
	Port string
	Host string
}

type ListenerDetails struct {
	Type string
	ID   string
}

func Login(username string, password string) bool {
	url := "http://192.168.181.136:8000/login"

	user := &userDetails{Username: username, Password: password}
	b, err := json.Marshal(user)

	if err != nil {
		return false
	}

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
		return false
	}

	defer response.Body.Close()

	fmt.Println("Response status : ", response.Status)
	fmt.Println("Response Headers : ", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body : ", string(body))

	stat := HandleSubmissions(response.Status)

	return stat
}

func HandleSubmissions(stat string) bool {
	if strings.Contains(stat, "200") {
		return true
	} else {
		return false
	}
}

func SendTcpListener(port string, host string) bool {
	url := "http://192.168.181.136:8000/listen"

	listener := &TcpListenerDetails{Type: "TCP", Port: port, Host: host}
	b, err := json.Marshal(listener)

	if err != nil {
		return false
	}

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
		return false
	}

	fmt.Println("Response status : ", response.Status)
	if strings.Contains(response.Status, "OK") {
		return true
	} else {
		return false
	}
	fmt.Println("Response Headers : ", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body : ", string(body))

	return false
}

func ShowListeners() []string {
	url := "http://192.168.181.136:8000/showListeners"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error ", err.Error())
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ", err.Error())
		return nil
	}

	fmt.Println("Response Body : ", strings.Split(string(body), "\""))
	return strings.Split(string(body), "\"")
}

func StopListener(listenerID string) {
	url := "http://192.168.181.136:8000/stopListener"

	listenerData := &ListenerDetails{Type: "TCP", ID: listenerID}
	b, err := json.Marshal(listenerData)

	if err != nil {
		return
	}

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
		return
	}

	fmt.Println("Response status : ", response.Status)
	if strings.Contains(response.Status, "OK") {
		return
	} else {
		return
	}

	fmt.Println("Response Headers : ", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body  : ", string(body))

	return
}
