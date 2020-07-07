package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"nats-streaming-webclient/pkg/natsclient"
	"net/http"
)



func Start(port string) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/send/", send)
	fmt.Println("Start listening port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

type SendRequest struct {
	Host    string `json:"host"`
	Cluster string `json:"cluster"`
	Queue   string `json:"queue"`
	Message string `json:"message"`
}

func send(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Body error"))
		return
	}
	sr := &SendRequest{}
	err = json.Unmarshal(body, sr)
	if err != nil {
		w.Write([]byte("Unmarshal error"))
		return
	}

	err = natsclient.FastPublish(sr.Host, sr.Cluster, sr.Queue, sr.Message)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/main.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}


