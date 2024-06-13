package web

import (
	_ "embed"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
)

var (
	//go:embed web.html
	webPage string
)

type Config struct {
	TargetServer string
}

type MusicSocketMessage struct {
	Client string `json:"client"`
	Type   string `json:"type"`
	Value  string `json:"value"`
	Time   string `json:"time"`
}

func (c Config) PageHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.New("web").Parse(webPage))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to render web page", http.StatusInternalServerError)
		return
	}

}

func (c Config) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	client := r.FormValue("client")
	inputType := r.FormValue("type")
	value := r.FormValue("value")
	if value == "" {
		fmt.Fprintf(w, "Invalid form data")
		return
	}

	var unixTimeString string
	clientTime := r.FormValue("currentTime")
	if clientTime == "" {
		unixTimeString = fmt.Sprintf("%d", time.Now().UnixMilli())
	} else {
		// convert the time to unix time
		unixTimeString = clientTime
	}
	// Dial the WebSocket server
	ws, _, err := websocket.DefaultDialer.Dial(c.TargetServer, nil)
	if err != nil {
		fmt.Println("Dial failed:", err)
		return
	}
	defer ws.Close()

	// Send a message to the server
	var message MusicSocketMessage
	if inputType == "light" {
		message = MusicSocketMessage{
			Client: client,
			Type:   "light",
			Value:  value,
			Time:   unixTimeString,
		}
	} else if inputType == "sound" {
		message = MusicSocketMessage{
			Client: client,
			Type:   "sound",
			Value:  value,
			Time:   unixTimeString,
		}
	}
	err = ws.WriteJSON(message)
	fmt.Printf("Sending message: %v\n", message)
	if err != nil {
		fmt.Println("Write error:", err)
	}

	fmt.Fprintf(w, `<div id="log">Message %s sent to %s</div>`, message, c.TargetServer)
}
