package models

type Subscriber struct {
	MSISDN         int    `json:"msisdn"`
	IMSI           string `json:"imsi"`
	Name           string `json:"name"`
	Identification string `json:"identification"`
	Birthday       string `json:"birthday"`
	IpAllocate     string `json:"ipallocate"`
	UPF            string `json:"upf"`
	PCF            string `json:"pcf"`
	CHF            string `json:"chf"`
}
