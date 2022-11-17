package main

import (
	"html/template"
	"log"
	"net/http"
)

type device struct {
	ErrorCode int
	Name      string
	Count     int
	Info      info
}

type info struct {
	ProductName string
	ProductCode int
	RevisionNo  int
	VendorId    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/scan", scan)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}

}

func scan(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		info := info{"MADLN05BE", 4885, 10000, "66FF"}
		device := device{0, "Device 0", 1, info}

		err := tpl.ExecuteTemplate(w, "scanProcess.gohtml", device)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
		return
	}

	err := tpl.ExecuteTemplate(w, "scan.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}

}
