package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var hostname string = ""
var token string = ""
var secret string = ""

type IPResponse struct {
	IP string `json:"ip"`
}

func update() bool {
	ip := getIP()
	if ip == "" {
		log.Println("Could not get IP-address")
		return false
	}

	url := fmt.Sprintf("https://api.domeneshop.no/v0/dyndns/update?hostname=%s&myip=%s", hostname, ip)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.SetBasicAuth(token, secret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	return res.StatusCode == 204
}

func getIP() (ip string) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Println(err)
		return ""
	}
	if resp.StatusCode != 200 {
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	var ipResponse IPResponse
	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		log.Println(err)
		return ""
	}
	return ipResponse.IP
}

func init() {
	hostname = os.Getenv("HOSTNAME")
	if hostname == "" {
		log.Fatalln("Missing environment variable: HOSTNAME")
	}

	token = os.Getenv("TOKEN")
	if token == "" {
		log.Fatalln("Missing environment variable: TOKEN")
	}

	secret = os.Getenv("SECRET")
	if secret == "" {
		log.Fatalln("Missing environment variable: SECRET")
	}
}

func main() {
	for {
		res := update()
		if !res {
			log.Println("Http request failed for api.domeneshop.no")
		}

		time.Sleep(time.Minute * 5)
	}
}
