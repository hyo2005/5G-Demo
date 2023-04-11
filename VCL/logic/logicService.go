package logic

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyo2005/demo-5g/initializer"
	"github.com/hyo2005/demo-5g/models"
)

// xu ly request tao moi
func CreateSubscriber(c *gin.Context) {
	//get data of req body
	var demoSub models.Subscriber
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)

	}
	errDecode := json.Unmarshal(data, &demoSub)

	if errDecode != nil {
		log.Println(errDecode)
	}

	//send request to IpAllocate
	var ipSub models.Subscriber
	clientToIpAllocate := &http.Client{}
	reqIpAllocate, errReqIpAllocate := http.NewRequest("GET", "http://localhost:8081/ip-allocate", nil)

	if errReqIpAllocate != nil {
		log.Println(errReqIpAllocate)

	}
	resFromIpAllocate, errResFromIpAllocate := clientToIpAllocate.Do(reqIpAllocate)
	if errResFromIpAllocate != nil {
		log.Println(errResFromIpAllocate)
	}
	//lay thong tin tu response
	resBody, errResBodyIpAllocate := ioutil.ReadAll(resFromIpAllocate.Body)
	if errResBodyIpAllocate != nil {
		log.Println(errResBodyIpAllocate)
	}
	errIp := json.Unmarshal(resBody, &ipSub)
	if errIp != nil {
		log.Println(errIp)
	}
	if resFromIpAllocate.StatusCode != 200 {
		log.Println("Error request to Ip Allocate")
		c.Data(501, "text", []byte("Panic request to IP Allocate"))
	}

	//

	//send request to UPF
	var uPF models.Subscriber
	clientToUPF := &http.Client{}
	reqUPF, errReqUPF := http.NewRequest("GET", "http://localhost:8082/upf-gw", nil)

	if errReqUPF != nil {
		log.Println(errReqUPF)

	}
	resFromUPF, errResFromUPF := clientToUPF.Do(reqUPF)
	if errResFromUPF != nil {
		log.Println(errResFromUPF)
	}
	//lay thong tin tu response
	resBodyUPF, errResBodyUPF := ioutil.ReadAll(resFromUPF.Body)
	if errResBodyUPF != nil {
		log.Println(errResBodyUPF)
	}
	errIpUPF := json.Unmarshal(resBodyUPF, &uPF)
	if errIpUPF != nil {
		log.Println(errIpUPF)
	}
	if resFromUPF.StatusCode != 200 {
		log.Println("Error request to UPF")
		c.Data(501, "text", []byte("Panic request to UPF"))
	}

	//

	//send request to PCF
	var pCF models.Subscriber
	clientToPCF := &http.Client{}
	reqPCF, errReqPCF := http.NewRequest("GET", "http://localhost:8083/pcf-gw", nil)

	if errReqPCF != nil {
		log.Println(errReqPCF)

	}
	resFromPCF, errResFromPCF := clientToPCF.Do(reqPCF)
	if errResFromPCF != nil {
		log.Println(errResFromPCF)
	}
	//lay thong tin tu response
	resBodyPCF, errResBodyPCF := ioutil.ReadAll(resFromPCF.Body)
	if errResBodyPCF != nil {
		log.Println(errResBodyPCF)
	}
	errIpPCF := json.Unmarshal(resBodyPCF, &pCF)
	if errIpPCF != nil {
		log.Println(errIpPCF)
	}
	if resFromPCF.StatusCode != 200 {
		log.Println("Error request to PCF")
		c.Data(501, "text", []byte("Panic request to PCF"))
	}

	//

	//send request to CHF
	var cHF models.Subscriber
	clientToCHF := &http.Client{}
	reqCHF, errReqCHF := http.NewRequest("GET", "http://localhost:8084/chf-gw", nil)

	if errReqCHF != nil {
		log.Println(errReqCHF)

	}
	resFromCHF, errResFromCHF := clientToCHF.Do(reqCHF)
	if errResFromCHF != nil {
		log.Println(errResFromCHF)
	}
	//lay thong tin tu response
	resBodyCHF, errResBodyCHF := ioutil.ReadAll(resFromCHF.Body)
	if errResBodyCHF != nil {
		log.Println(errResBodyCHF)
	}
	errIpCHF := json.Unmarshal(resBodyCHF, &cHF)
	if errIpCHF != nil {
		log.Println(errIpCHF)
	}
	if resFromCHF.StatusCode != 200 {
		log.Println("Error request to CHF")
		c.Data(501, "text", []byte("Panic request to CHF"))
	}

	//

	//tao 1 bien subscriber de parse data
	subscriber := models.Subscriber{
		MSISDN:         demoSub.MSISDN,
		IMSI:           demoSub.IMSI,
		Name:           demoSub.Name,
		Identification: demoSub.Identification,
		Birthday:       demoSub.Birthday,
		IpAllocate:     ipSub.IpAllocate,
		UPF:            uPF.UPF,
		PCF:            pCF.PCF,
		CHF:            cHF.CHF,
	}
	result := initializer.DB.Create(&subscriber)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//return
	c.JSON(200, gin.H{
		"subscriber": subscriber,
	})
}

// xu ly request tim kiem
func ViewSubscriber(c *gin.Context) {
	id := c.Param("msisdn")
	var demoSub models.Subscriber
	initializer.DB.First(&demoSub, id)

	c.JSON(200, gin.H{
		"subscriber": demoSub,
	})
}

// xu ly request update
func UpdateSubscriber(c *gin.Context) {
	var demoSub models.Subscriber
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
	}
	errDecode := json.Unmarshal(data, &demoSub)
	if errDecode != nil {
		log.Println(errDecode)
	}

	// tim kiem thue bao muon update
	id := demoSub.MSISDN
	var tempSub models.Subscriber
	initializer.DB.First(&tempSub, id)
	// update thong tin thue bao
	initializer.DB.Model(&tempSub).Where("MSISDN = ?", id).Updates(models.Subscriber{
		Name:           demoSub.Name,
		Identification: demoSub.Identification,
		Birthday:       demoSub.Birthday,
	})
	c.JSON(200, gin.H{
		"subscriber": demoSub,
	})
}

// xu ly request delete
func DeleteSubscriber(c *gin.Context) {
	id := c.Param("id")
	initializer.DB.Where("MSISDN = ?", id).Delete(&models.Subscriber{})
	c.Status(200)
}
