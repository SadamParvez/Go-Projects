package main

import (
	"fmt"
	"log"
	"net/http"
)

func StudentInfo(flag string) (string, string, string) {
	if flag == "student" {
		name := "Sadam Parvez"
		email := "s@gmail.com"
		phoneNo := "123321"
		return name, email, phoneNo
	} else {
		return "unknown", "unknown", "unknown"
	}
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getinfo" {
		http.Error(w, "404 --  Not Found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
	}
	name, email, phoneno := StudentInfo("student")
	fmt.Fprintf(w, "Name = %s\nEmail = %s\nPhone No = %s", name, email, phoneno)

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// if r.Method == http.MethodGet {
	// 	http.ServeFile(w, r, "./static/form.html")
	// 	return
	// }
	// if r.Method == http.MethodPost {
	// 	if err := r.ParseForm(); err != nil {
	// 		fmt.Fprintf(w, "FormParseError: %v", err)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, "POST Request Successful \n")
	// 	name := r.FormValue("name")
	// 	email := r.FormValue("email")
	// 	phoneno := r.FormValue("phoneno")
	// 	fmt.Fprintf(w, "Name = %s\nEmail = %s\nPhoneNumber = %s \n", name, email, phoneno)
	// }

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./static/form.html")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "FormParseError: %v", err)
			return
		}
		fmt.Fprintf(w, "POST Request Successful \n")
		name := r.FormValue("name")
		email := r.FormValue("email")
		phoneno := r.FormValue("phoneno")
		fmt.Fprintf(w, "Name = %s\nEmail = %s\nPhoneNumber = %s \n", name, email, phoneno)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // it will look for index.html in static dir
	http.Handle("/", fileServer)
	http.HandleFunc("/getinfo", handleInfo)
	http.HandleFunc("/form", handleForm)
	log.Println("Starting Server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

