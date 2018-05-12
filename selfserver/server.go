package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("app listen on 8080")

	mux := initlisteners(http.NewServeMux())

	http.ListenAndServe(":8080", mux)
}

func initlisteners(mux *http.ServeMux) *http.ServeMux {

	mux.Handle("/save", handleSave())
	mux.Handle("/get", handleGet())

	return mux
}

func handleSave() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			log.Println(r.Header)

			defer r.Body.Close()

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				rw.WriteHeader(500)
				return
			}

			file, err := os.OpenFile("file.zip", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
			if err != nil {
				log.Println(err)
				rw.WriteHeader(500)
				return
			}

			defer file.Close()

			io.Copy(file, bytes.NewReader(body))
			rw.WriteHeader(200)

		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func handleGet() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			file, err := os.Open("file.zip")
			if err != nil {
				log.Println(err)
				rw.WriteHeader(500)
				return
			}

			defer file.Close()

			io.Copy(rw, file)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
