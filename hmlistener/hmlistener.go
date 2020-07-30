package hmlistener

import (
	"fmt"
	"github.com/dhcgn/gohomematicmqttplugin/server"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//StartServer start a http listener and write all incoming request to the chan
func StartServer(messages chan<- string, port int) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {

		bodyBytes, _ := ioutil.ReadAll(request.Body)
		bodyString := string(bodyBytes)

		var r string

		if request.Method == "POST" && strings.HasPrefix(bodyString, `<?xml`) {
			messages <- bodyString
			log.Println("Moved incoming message to channel")
			r = fmt.Sprintf("Got %d bytes", len(bodyBytes))
		} else {
			log.Println("Invalid message will be droped")
			log.Println(bodyString[0:64])
			r = fmt.Sprintf("Invalid message")
		}

		w.Header().Set("Content-Type", "text/plan; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(r))
	})

	srv := server.New(mux, port)
	log.Print("Starting Server at: ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
