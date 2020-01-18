package main

import (
	"backery/router"
	"backery/structs"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	testServerListener, _ := net.Listen("tcp", ":5678")
	testServerSv := &http.Server{Handler: router.Init()}
	go func() { log.Fatal(testServerSv.Serve(testServerListener)) }()
}

func TestOrder(t *testing.T) {
	Convey("Check for health status", t, func() {
		var orderResp *structs.OrderResp
		jsonBody, _ := json.Marshal(structs.OrderReq{
			Code:     "VS5",
			Quantity: 10,
		})
		resp, err := getHttpClient().Post("http://127.0.0.1:5678/health", "application/json", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal("Cannot get health status because of", err.Error())
		}
		So(resp.StatusCode, ShouldEqual, http.StatusOK)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("cannot read response")
		}
		err = json.Unmarshal(body, &orderResp)
		if err != nil {
			t.Fatal("cannot unmarshal response")
		}
		So(orderResp.TotalPrice, ShouldEqual, 17.98)
	})
}

func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
}
