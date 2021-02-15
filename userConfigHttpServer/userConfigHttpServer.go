package userConfigHttpServer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dhcgn/gohomematicmqttplugin/server"
	"github.com/dhcgn/gohomematicmqttplugin/shared"
)

const forms = `
<!DOCTYPE html>
<html lang="en">
<body>
{{if .Success}}
	<meta http-equiv="refresh" content="1"/>
    <h1>Saved!</h1>
	<p>Page will be reloaded in one second.</p>
{{else}}
    <h1>HomeMatic MQTT Plugin</h1>
	<p>More information: <a href="https://github.com/dhcgn/GoHomeMaticMqttPlugin">https://github.com/dhcgn/GoHomeMaticMqttPlugin</a></p>
    <form method="POST">
        <label>ListenerPort:</label><br />
        <input type="text" name="ListenerPort" value="%ListenerPort%"><br />
		<small>Port which will be used to listen local to rpc callbacks form the ccu, e.g. 8777</small><br /><br />
        <label>InterfaceId:</label><br />
        <input type="text" name="InterfaceId" value="%InterfaceId%"><br />
		<small>InterfaceId which will be used to subscribe to rpc callbacks, e.g. 1 </small><br /><br />
        <label>HomematicUrl:</label><br />
        <input type="text" name="HomematicUrl" value="%HomematicUrl%"><br />
		<small>Url to access the XML-RPC-Server, e.g. http://127.0.0.1:2001/ </small><br /><br />
        <label>BrokerUrl:</label><br />
        <input type="text" name="BrokerUrl" value="%BrokerUrl%"><br />
		<small>Url to access your mqtt broker, e.g. tcp://192.168.1.100:1883 </small><br /><br />
        <input type="submit" value="Save">
    </form>
{{end}}
</body>
</html>
`

func createTemplate() string {
	r := strings.ReplaceAll(forms, "%ListenerPort%", strconv.Itoa(shared.Config.ListenerPort))
	r = strings.ReplaceAll(r, "%InterfaceId%", strconv.Itoa(shared.Config.InterfaceId))
	r = strings.ReplaceAll(r, "%HomematicUrl%", shared.Config.HomematicUrl)
	r = strings.ReplaceAll(r, "%BrokerUrl%", shared.Config.BrokerUrl)

	return r
}

func StartWebService() {
	tmpl, _ := template.New("foo").Parse(createTemplate())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		port, err := strconv.Atoi(r.FormValue("ListenerPort"))
		if err != nil {
			fmt.Fprintf(w, "Port must be a number, was: %s", r.FormValue("ListenerPort"))
			return
		}
		interfaceId, err := strconv.Atoi(r.FormValue("InterfaceId"))
		if err != nil {
			fmt.Fprintf(w, "InterfaceId must be a number, was: %s", r.FormValue("InterfaceId"))
			return
		}

		c := shared.Configuration{
			ListenerPort: port,
			InterfaceId:  interfaceId,
			HomematicUrl: r.FormValue("HomematicUrl"),
			BrokerUrl:    r.FormValue("BrokerUrl"),
		}

		shared.UpdateConfiguration(c)

		tmpl.Execute(w, struct{ Success bool }{true})

		tmpl, _ = template.New("foo").Parse(createTemplate())
	})

	port := 8070
	srv := server.New(mux, port)
	log.Println("Starting http server for user configuration in port", port)
	srv.ListenAndServe()
}
