package client

import (
	ioutil "io/ioutil"
	"log"
	"net/http"
)

func Client() {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
