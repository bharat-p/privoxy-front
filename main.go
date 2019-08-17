//go:generate statik -src=./static
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"privoxy-front/apis"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"

	"github.com/rakyll/statik/fs"
	_  "privoxy-front/statik"

)



var configFile = "/etc/privoxy/whitelist.action"

type saveRequest struct {
	Data string `json:"data"`
}

type errorResponse struct {
	Message string `json:message`
}

func main() {
	r := mux.NewRouter()

	// Handle API routes
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/getConfig", func(w http.ResponseWriter, r *http.Request) {
		json := simplejson.New()

		w.Header().Set("Content-Type", "application/json")
		config, err := apis.LoadConfig(configFile)
		if err != nil {
			json.Set("error", err.Error())
			payload, err := json.MarshalJSON()
			if err != nil {
				log.Println(err)
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(payload)

			return
		}


		json.Set("data", string(config))

		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)

	})

	api.HandleFunc("/saveConfig", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		fmt.Printf("got request\n")

		var req saveRequest
		err := decoder.Decode(&req)

		if err != nil {
			fmt.Errorf("error\n")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("Data = %s\n", req.Data)
		err = apis.SaveConfig(configFile, req.Data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error saving file. %s", err), http.StatusInternalServerError)
			return
		}


		w.Write([]byte("Success"))


	})

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Serve index page on all unhandled routes
	r.PathPrefix("/").Handler(http.FileServer(statikFS))

	fmt.Println("http://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", r))

}