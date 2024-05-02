package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

//go:embed template
var tplFolder embed.FS

func main() {

	API_KEY := os.Getenv("API_KEY")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get(API_KEY)
		if err != nil {
			log.Print(err.Error())
			os.Exit(1)
		}

		defer response.Body.Close()

		var responseData Response
		err = json.NewDecoder(response.Body).Decode(&responseData)

		if err != nil {
			log.Fatal(err)
		}

		tmpl, err := template.ParseFS(tplFolder, "template/template.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, responseData)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
