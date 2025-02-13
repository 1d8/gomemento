package webmon

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/xFaraday/gomemento/alertmon"
	"github.com/xFaraday/gomemento/config"
)

type Beat struct {
	IP string
}

var (
	ssUserAgent = config.GetSerialScripterUserAgent()
	ssIP        = config.GetSerialScripterIP()
)

func GetIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ipaddr := localAddr.IP
	return ipaddr.String()
}

func HeartBeat() {
	m := Beat{IP: GetIP()}
	jsonStr, err := json.Marshal(m)
	if err != nil {
		println("error")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	bodyReader := bytes.NewReader(jsonStr)

	requestURL := fmt.Sprintf(ssIP)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		println("error")
	}

	req.Header.Set("User-Agent", ssUserAgent)
	resp, err := client.Do(req)
	if err != nil {
		println("error")
	} else {
		//data, _ := ioutil.ReadAll(resp.Body)
		//println(string(data))
	}

	defer resp.Body.Close()

}

func IncidentAlert(alert alertmon.Alert) {

	jsonStr, err := json.Marshal(alert)
	if err != nil {
		println("error")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	bodyReader := bytes.NewReader(jsonStr)

	requestURL := fmt.Sprintf(ssIP)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		println("error")
	}

	req.Header.Set("User-Agent", ssUserAgent)
	resp, err := client.Do(req)
	if err != nil {
		println("error")
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		println(string(data))
	}

	defer resp.Body.Close()
}
