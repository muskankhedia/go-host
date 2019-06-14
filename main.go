package main
import (
  "net/http"
  "strings"
  "os"
  "log"
)

type jsonResponseQuery struct {
	Result []string `json:"result"`
}

func fetchData(w http.ResponseWriter, r *http.Request) {

  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message

  w.Write([]byte(message))
} 

func main() {

  http.HandleFunc("/", fetchData)
  port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
  }

  port = ":" + port

  if err := http.ListenAndServe(port, nil); err != nil {
    panic(err)
  }

}