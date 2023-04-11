package main

import (
	"Client_5G/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	demoSubs := models.Subscriber{}

	byteData, err := json.Marshal(demoSubs)
	if err != nil {
		fmt.Println(err)
	}

	//request tao moi thue bao
	clientCreate := &http.Client{}
	reqCreate, errCreate := http.NewRequest("POST", "http://localhost:8080/create-subscriber", bytes.NewBuffer(byteData))

	if errCreate != nil {
		log.Fatal(errCreate)
		return
	}

	resCreate, errResCreate := clientCreate.Do(reqCreate)

	if errResCreate != nil {
		log.Fatal(errResCreate)
		return
	}
	log.Println(resCreate)
	defer resCreate.Body.Close()

	//request tim kiem thue bao
	clientView := &http.Client{}
	reqView, errView := http.NewRequest("GET", "http://localhost:8080/view-subscriber", nil)

	if errView != nil {
		log.Fatal(errView)
		return
	}

	resView, errResView := clientView.Do(reqView)
	if errResView != nil {
		log.Fatal(errResView)
		return
	}
	log.Println(resView)
	defer resView.Body.Close()

	//request update thue bao
	clientUpdate := &http.Client{}
	reqUpdate, errUpdate := http.NewRequest("PUT", "http://localhost:8080/update-subscriber", bytes.NewBuffer(byteData))

	if errUpdate != nil {
		log.Fatal(errUpdate)
		return
	}

	resUpdate, errResUpdate := clientUpdate.Do(reqUpdate)
	if errResUpdate != nil {
		log.Fatal(errResUpdate)
		return
	}
	log.Println(resUpdate)
	defer resView.Body.Close()
}
