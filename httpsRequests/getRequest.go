package main

import (
	"io"
	"log"
	"net/http"
)


func main(){
//in net/http package we will import a Get function used for 
// making Get requests

//The GEt function takes in a URL and 
// returns a response of type pointer to a struct and an error.
//When the error is nil, the response will contain a response body.

resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

if err != nil{
	log.Fatalln(err)
}

//we will read the response body in here
//the reason we read it and not print the resp
//is because we will get a bunch of incoherent data 
// such as the headers used to make the request etc

//so we have to use io.ReadAll to read resp body


body, err := io.ReadAll((resp.Body))
if err != nil{
	log.Fatalln(err)
}

//then conver that body to type string
newBody := string(body)
log.Print((newBody))
}