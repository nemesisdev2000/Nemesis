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

func SendPost(rw http.ResponseWriter, r *http.Request) {
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
}
