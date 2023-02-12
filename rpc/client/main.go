package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func directCall() {

	url := "http://app-rpc-s.lab.local/v1.0/invoke/app-rpc-s/method/tick"
	// url := "http://app-rpc-s.lab.local:81/tick"
	rsp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Errorf("%v: %s", rsp.StatusCode, rsp.Body)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("response body: %s", body)
	rspBody := &Response{}
	if err := json.Unmarshal(body, rspBody); err != nil {
		log.Fatal(err)
	}
	log.Infof("response body struct: %v", rspBody)
}

func main() {

	directCall()

	os.Setenv("DAPR_CLIENT_TIMEOUT_SECONDS", "3")
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	for {
		ctx := context.Background()
		result, err := client.InvokeMethod(ctx, "app-rpc-s", "tick", "get")
		if err != nil {
			log.Error(err)
		}
		log.Infof("%s", result)
		time.Sleep(3 * time.Second)
	}

}
