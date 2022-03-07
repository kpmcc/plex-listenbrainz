package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Player struct {
	Local         bool   `json:local`
	PublicAddress string `json:publicAddress`
	Title         string `json:title`
	UUID          string `json:uuid`
}

type Server struct {
	Title string `json:title`
	UUID  string `json:uuid`
}

type Account struct {
	Id           int    `json:id`
	ThumbnailUrl string `json:thumb`
	Title        string `json:title`
}

type WebHook struct {
	EventType    string                 `json:event`
	User         bool                   `json:user`
	Owner        bool                   `json:owner`
	EventAccount Account                `json:Account`
	Server       Server                 `json:Server`
	Player       Player                 `json:Player`
	Metadata     map[string]interface{} `json:Metadata`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	if r.Method == http.MethodPost {
		//var hook WebHook
		var body = r.Body
		data, _ := io.ReadAll(body)
		var result map[string]interface{}
		json.Unmarshal(data, &result)
		//fmt.Print(hook)
		fmt.Print(result)
		//fmt.Print(hook.EventAccount)
	}
}
