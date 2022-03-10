package main

import (
	"encoding/json"
	"fmt"
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

type Metadata struct {
	AddedAt              int    `json:addedAt`
	GrandparentGuid      string `json:grandparentGuid`
	GrandparentKey       string `json:grandparentKey`
	GrandparentRatingKey int    `json:grandparentRatingKey`
	GrandparentThumb     string `json:grandparentThumb`
	GrandparentTitle     string `json:grandparentTitle`
	Guid                 string `json:guid`
	ParentGuid           string `json:parentGuid`
	LibrarySectionTitle  string `json:librarySectionTitle`
	LibrarySectionType   string `json:librarySectionType`
	ParentTitle          string `json:parentTitle`
	Title                string `json:title`
	ParentYear           string `json:parentYear`
	MdType               string `json:type`
}

type WebHook struct {
	Event   string  `json:event`
	User    bool    `json:user`
	Owner   bool    `json:owner`
	Account Account `json:Account`
	Server  Server  `json:Server`
	Player  Player  `json:Player`
	//Metadata map[string]interface{} `json:Metadata`
	Metadata Metadata `json:Metadata`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Success")
	//for k, v := range r.Header {
	//	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//}
	//fmt.Fprintf(w, "Host = %q\n", r.Host)
	//fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(4 * 1024); err != nil {
			log.Print(err)
		}
		for k, v := range r.PostForm {
			//fmt.Printf("Form [%s] = %s\n", k, v)
			if k == "payload" {
				//fmt.Printf("len v: %d\n", len(v))
				//fmt.Printf("v[0]: %q\n", v[0])
				if len(v) == 1 {
					b := []byte(v[0])
					//var m map[string]interface{}
					var w WebHook
					err := json.Unmarshal(b, &w)

					//////////////////////////////////////////
					if err != nil {
						fmt.Println("error unmarshalling json")
					}
					//fmt.Println("WebHook")
					//fmt.Println(w)
					fmt.Println("Event")
					fmt.Println(w.Event)
					fmt.Println("Account")
					fmt.Println(w.Account)
					fmt.Println("User")
					fmt.Println(w.User)
					fmt.Println("Server")
					fmt.Println(w.Server)
					fmt.Println("Player")
					fmt.Println(w.Player)
					fmt.Println("Metadata")
					fmt.Println(w.Metadata)
					//m := w.Metadata
					//fmt.Println("gpTitle")
					//fmt.Println(m.grandparentTitle)
					//fmt.Println("parentTitle")
					//fmt.Println(m.parentTitle)
					//fmt.Println("title")
					//fmt.Println(m.title)

					//event := m["event"]
					//metadata := m["Metadata"]
					//fmt.Printf("event: %q\n", event)
					//if event == "media.scrobble" {
					//	fmt.Printf("Got Scrobble\n")
					//} else if event == "media.play" {
					//	fmt.Printf("Got Play\n")
					//} else if event == "media.stop" {
					//	fmt.Printf("Got Stop\n")
					//} else if event == "media.pause" {
					//	fmt.Printf("Got Pause\n")
					//}

					//md := metadata.(map[string]interface{})
					//scrobble := make(map[string]string)
					//for ki, vi := range md {
					//	switch vv := vi.(type) {
					//	case string:
					//		//fmt.Println(ki, "is string", vv)
					//		scrobble[ki] = vv
					//	case float64:
					//		//fmt.Println(ki, "is float64", vv)
					//	case []interface{}:
					//		//fmt.Println(ki, "is an array:")
					//		//for i, u := range vv {
					//		//	//fmt.Println(i, u)
					//		//}
					//	default:
					//		fmt.Println(k, "is of a type I don't know how to handle")
					//	}
					//}
					//fmt.Printf("%s\n", scrobble["grandparentTitle"])
					//fmt.Printf("%s\n", scrobble["parentTitle"])
					//fmt.Printf("%s\n", scrobble["title"])
					//////////////////////////////////////////

					//for k, _ := range metadata {
					//	fmt.Printf("k: [%q]\n", k)
					//}
					//fmt.Printf("%s\n", metadata["grandParentTitle"])
					//fmt.Printf("%s\n", metadata["parentTitle"])
					//fmt.Printf("%s\n", metadata["title"])
					//fmt.Printf("%s\n", metadata["parentYear"])
				}
			}
		}
		//var hook WebHook
		//var body = r.Body
		//data, _ := io.ReadAll(body)
		//var result map[string]interface{}
		//json.Unmarshal(data, &result)
		////fmt.Print(hook)
		//fmt.Println(result)
		//fmt.Println("r")
		//fmt.Println(r)
		//fmt.Print(hook.EventAccount)
	}
}
