package service

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// xu ly request tao moi
func HandleCreateSub(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)

	}
	a := bytes.NewBuffer(data)
	log.Println(a)

	client := &http.Client{}
	req, errReq := http.NewRequest("POST", "http://localhost:8888/logic-create", a)

	if errReq != nil {
		log.Println(errReq)

	}

	resFromLogic, errRes := client.Do(req)
	if errRes != nil {
		log.Println(errRes)

	}
	if resFromLogic.StatusCode != 200 {
		log.Println("Logic Error")
		c.Data(500, "text", []byte("Logic Panic"))
	}
	c.Data(200, "text", []byte("Forward Create Request Successful"))
}

// xu ly request tim kiem
func HandleViewSub(c *gin.Context) {
	client := &http.Client{}
	req, errReq := http.NewRequest("GET", "http://localhost:8888/logic-view/:id", nil)

	if errReq != nil {
		log.Println(errReq)

	}

	resFromLogic, errRes := client.Do(req)
	if errRes != nil {
		log.Println(errRes)

	}
	if resFromLogic.StatusCode != 200 {
		log.Println("Logic Error")
		c.Data(500, "text", []byte("Logic Panic"))
	}
	c.Data(200, "text", []byte("Forward Search Request Successful"))
}

// xu ly request update
func HandleUpdateSub(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)

	}
	a := bytes.NewBuffer(data)
	log.Println(a)

	client := &http.Client{}
	req, errReq := http.NewRequest("PUT", "http://localhost:8888/logic-update", a)

	if errReq != nil {
		log.Println(errReq)

	}

	resFromLogic, errRes := client.Do(req)
	if errRes != nil {
		log.Println(errRes)
	}

	if resFromLogic.StatusCode != 200 {
		log.Println("Logic Error")
		c.Data(500, "text", []byte("Logic Panic"))
	}
	c.Data(200, "text", []byte("Forward Update Request Successful"))
}

// xu ly request delete
func HandleDeleteSub(c *gin.Context) {
	client := &http.Client{}
	req, errReq := http.NewRequest("DELETE", "http://localhost:8888/logic-delete/:id", nil)

	if errReq != nil {
		log.Println(errReq)

	}

	resFromLogic, errRes := client.Do(req)
	if errRes != nil {
		log.Println(errRes)

	}
	if resFromLogic.StatusCode != 200 {
		log.Println("Logic Error")
		c.Data(500, "text", []byte("Logic Panic"))
	}
	c.Data(200, "text", []byte("Forward Delete Request Successful"))
}
