package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WebAccess interface {
	Post(json_bytes []byte, href_slug string) ([]byte, error)
	Get(href_slug string) ([]byte, error)
}

type ConnectionParameters_t struct {
	BaseUrl    string
	Username   string
	Password   string
	MIQToken   string
	verify_ssl bool
	GUID       string
}

func (params ConnectionParameters_t) Post(json_bytes []byte, href_slug string) ([]byte, error) {
	req, err := http.NewRequest("POST", params.BaseUrl+href_slug, bytes.NewBuffer(json_bytes))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	if params.MIQToken != "" {
		req.Header.Set("X-Auth-Token", params.MIQToken)
	} else if params.Username != "" && params.Password != "" {
		req.SetBasicAuth(params.Username, params.Password)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do: ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return body, nil
}

func (params ConnectionParameters_t) Get(href_slug string) ([]byte, error) {
	fmt.Println("URL is ", params.BaseUrl+href_slug)
	req, err := http.NewRequest("GET", params.BaseUrl+href_slug, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	if params.MIQToken != "" {
		req.Header.Set("X-Auth-Token", params.MIQToken)
	} else if params.Username != "" && params.Password != "" {
		req.SetBasicAuth(params.Username, params.Password)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do: ", err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return body, nil
}
