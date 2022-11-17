package main

import (
	"html/template"
	"log"
	"net/http"
)

type profile struct {
	ProfileNo  int
	ErrorCode  int
	Count      int
	DeviceList []device
}

type device struct {
	Name string
	Info info
}

type info struct {
	VendorName  string
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
		d0 := device{
			Name: "Device0",
			Info: info{
				VendorName:  "eMotionTek",
				ProductName: "UNI_ECAT60B_2X",
				ProductCode: 2003,
				RevisionNo:  1,
				VendorId:    "AAA555",
			},
		}

		d1 := device{
			Name: "Device1",
			Info: info{
				VendorName:  "eMotionTek",
				ProductName: "UNI_ECAT60B_1X",
				ProductCode: 2002,
				RevisionNo:  1,
				VendorId:    "AAA555",
			},
		}

		d2 := device{
			Name: "Device2",
			Info: info{
				VendorName:  "FASTECH",
				ProductName: "Ezi-SERVO2 EtherCAT",
				ProductCode: 1002,
				RevisionNo:  1,
				VendorId:    "0FA0",
			},
		}

		p := profile{
			ProfileNo:  402,
			ErrorCode:  0,
			Count:      3,
			DeviceList: []device{d0, d1, d2},
		}

		err := tpl.ExecuteTemplate(w, "scanProcess.gohtml", p)
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
