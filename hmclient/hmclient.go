package hmclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GetDevices request all device from the CCU
func GetDevices() string {
	body := `<?xml version="1.0"?>
<methodCall><methodName>listDevices</methodName>`

	req, _ := http.NewRequest("POST", "http://192.168.10.23:2001/", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}

//Init send a request to the CCU to subscribe http posts
func Init(port int, interfaceId int, homematicUrl string) string {
	body :=
		`<?xml version="1.0"?>
<methodCall>
  <methodName>init</methodName>
  <params>
    <param>
      <value>
        <string>http://127.0.0.1:%d</string>
      </value>
    </param>
    <param>
      <value>
        <string>%d</string>
      </value>
    </param>
  </params>
</methodCall>`

	body = fmt.Sprintf(body, port, interfaceId)
	req, e := http.NewRequest("POST", homematicUrl, bytes.NewReader([]byte(body)))
	if e != nil {
		log.Println("Init ERROR:", e)
		return ""
	}
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, e := client.Do(req)

	if e != nil {
		log.Println("Init ERROR:", e)
		return ""
	}

	b, _ := ioutil.ReadAll(resp.Body)
	log.Println("Init OK")

	return string(b)
}
