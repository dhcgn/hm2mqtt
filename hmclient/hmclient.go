package hmclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//GetDevices request all device from the CCU
func GetDevices() string {
	body := `<?xml version="1.0"?>
<methodCall><methodName>listDevices</methodName>`

	// TODO url must be configured. But should be used for debugging purpose too.
	req, _ := http.NewRequest("POST", "http://192.168.10.23:2001/", bytes.NewReader([]byte(body)))
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}

// SetState sends a RPC call to homematic to set a value
func SetState(address string, valueKey string, value string, homematicURL string) string {
	body := `<?xml version="1.0"?>
<methodCall>
    <methodName>setValue</methodName>
    <params>
        <param>
            <value>%v</value>
        </param>
        <param>
            <value>%v</value>
        </param>
        <param>
            <value>%v</value>
        </param>
    </params>
</methodCall>`

	body = fmt.Sprintf(body, address, valueKey, value)

	fmt.Println(body)

	req, e := http.NewRequest("POST", homematicURL, bytes.NewReader([]byte(body)))
	if e != nil {
		log.Println("SetState ERROR:", e)
		return ""
	}
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, e := client.Do(req)

	if e != nil {
		log.Println("SetState ERROR:", e)
		return ""
	}

	b, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(b), "faultCode") {
		log.Println("SetState Response ERROR")
	} else {
		log.Println("SetState Response OK")
	}

	return string(b)
}

//Init send a request to the CCU to subscribe http posts
func Init(port int, interfaceID int, homematicURL string) string {
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

	body = fmt.Sprintf(body, port, interfaceID)
	req, e := http.NewRequest("POST", homematicURL, bytes.NewReader([]byte(body)))
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
