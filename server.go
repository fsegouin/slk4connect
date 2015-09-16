package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	url = "https://hooks.slack.com/services/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"
)

type SlackResponse struct {
	Channel     string       `json:"channel"`
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
	r.ParseForm()

	response := SlackResponse{}
	fmt.Println("path", r.Form)
	response.Channel = r.Form["channel_name"][0]
	response.Text = ":red_circle: :red_circle: :red_circle:"
	//response.Attachments = append(response.Attachments, attachement)

	jsonResponse, _ := json.Marshal(response)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonResponse))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Fprintf(w, "")
}

func main() {
	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/webhook", http.HandlerFunc(webhookHandler))

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
