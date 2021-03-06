package main

import (
	"time"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"regexp"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
	"html/template"
)

var validPath = regexp.MustCompile("^(open|close|status)$")

var templates = template.Must(template.ParseFiles("web/button.html"))

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates.Execute(w, r)
}


func initGobot() (*gobot.Gobot, httprouter.Handle) {
	_gobot := gobot.NewGobot()

	raspiAdapter := raspi.NewRaspiAdaptor("raspi")

	remoteButton := gpio.NewRelayDriver(raspiAdapter, "remoteButton", "26")

	robot := gobot.NewRobot("remoteButtonBot",
		[]gobot.Connection{raspiAdapter},
		[]gobot.Device{remoteButton},
	)

	_gobot.AddRobot(robot)

	gobotHandler := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		command := ps.ByName("command")

		m := validPath.FindStringSubmatch(command)
		if m == nil {
			http.Error(w, "Invalid command: "+command, http.StatusBadRequest)
			return
		}

		log.Printf("Command [%s] accepted.", command)

		if (command=="open" || command=="close") {
			remoteButton.On();
			time.Sleep(2000 * time.Millisecond)
			log.Println("[+] Relay going LOW")
			remoteButton.Off();

			log.Printf("state: %v",remoteButton.State())
		}

		templates.Execute(w, r)
	}

	return _gobot, gobotHandler;
}

func main() {
	_gobot, gobotHandler := initGobot()

	router := httprouter.New()
	router.GET("/garer", indexHandler)

	router.GET("/garer/v1/cmd/:command", gobotHandler)


	log.Printf("Blown will a port:8080")

	log.Fatal(http.ListenAndServe(":8080", router))

	_gobot.Start()
}