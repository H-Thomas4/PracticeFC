package main

import (
	"PracticeFC/Handle"
	"PracticeFC/Repo"
	"PracticeFC/Service"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main(){
fileName := "FcDb.json"
ext := filepath.Ext(fileName)

if ext != ".json" {
log.Fatalln("incorrect file extension")
}


Repo := Repo.NewFcRepo(fileName)
Svc := Service.NewFcService(Repo)
handler := Handle.NewFcHandler(Svc)
router := Handle.NewServer(handler)



server := &http.Server{
Handler: router,
Addr:    "127.0.0.1:8080",
}

fmt.Println("Successfully running server 8080")

log.Fatal(server.ListenAndServe())

}