package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SlackResponse struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	ImageUrl string `json:"image_url"`
	Color    string `json:"color"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	//attachement := Attachment{
	//	"Connect 4!!!",
	//	"Player 1 played",
	//	"https://media.giphy.com/media/yoJC2i7kpPtZrvAZR6/giphy.gif",
	//	"#00ff00",
	//}

	response := SlackResponse{}
	response.Text = ":red_circle: :red_circle: :red_circle:"
	//response.Attachments = append(response.Attachments, attachement)

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/webhook", http.HandlerFunc(webhookHandler))

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
