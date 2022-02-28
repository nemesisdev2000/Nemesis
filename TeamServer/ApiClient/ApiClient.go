package ApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmcvetta/napping"
	"github.com/nemesisdev2000/Nemesis/TeamServer/AuthenticationServer"
)

type Post struct {
	Type  string `json:"type"`
	Port  string `json:"port"`
	Token string `json:"token"`
}

type Config struct {
	Type  string `json:"type"`
	Port  string `json:"port"`
	Token string `json:"token"`
	Id    string `json:"id"`
}

func ApiStartListener(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	fmt.Println(post)
	if err != nil {
		fmt.Println(err)
	}

	t := AuthenticationServer.ValidateRequest(post.Token)
	if t == false {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Unauthorized Request"))
		return
	}
	url := "http://localhost:8000/listen"
	s := napping.Session{}
	h := &http.Header{}
	h.Set("Content-Type", "application/json")
	s.Header = h

	//var jsonStr = []byte(`{ "type":"TcpListener", "port":"1331"}`)
	jsonStr, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	}

	var data map[string]json.RawMessage
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status())
	fmt.Println(resp.RawText())
	rw.Write([]byte(resp.RawText()))
}

func ApiStopListener(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var config Config
	err := json.NewDecoder(r.Body).Decode(&config)
	fmt.Println("Config : ", config)
	if err != nil {
		fmt.Println(err)
	}

	t := AuthenticationServer.ValidateRequest(config.Token)
	if t == false {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Unauthorized Request"))
		return
	}

	url := "http://localhost:8000/stoplistener/" + config.Id
	s := napping.Session{}
	h := &http.Header{}
	h.Set("Content-Type", "text/html")
	s.Header = h

	jsonStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}

	var data map[string]json.RawMessage
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := s.Get(url, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status())
	fmt.Println(resp.RawText())
	rw.Write([]byte(resp.RawText()))
}
