package main

import (
	"./school"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func loadView(w http.ResponseWriter, name string) {
	bytes, err := ioutil.ReadFile("views/" + name + ".html")
	if err == nil {
		w.Write(bytes)
	}
}

func main() {
	sc := school.School{ Students: make(map[string]map[string]float64) };

	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
		loadView(w, "index")
	})

	http.HandleFunc("/addgrade", func (w http.ResponseWriter, req *http.Request) {
		var g school.Grade

		err := json.NewDecoder(req.Body).Decode(&g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
      return
		} 
	
		sc.AddGrade(g)
		w.Header().Set("Content-type", "text/plain")
		fmt.Fprintf(w, "ok")
	})

	http.HandleFunc("/grades", func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		fmt.Fprintf(w, "%f", sc.GetGeneralAverage())
	})

	http.HandleFunc("/grades/student", func (w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		w.Header().Set("Content-type", "text/plain")
		fmt.Fprintf(w, "%f", sc.GetStudentAverage(string(body)))
	})

	http.HandleFunc("/grades/class", func (w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		w.Header().Set("Content-type", "text/plain")
		fmt.Fprintf(w, "%f", sc.GetClassAverage(string(body)))
	})

	http.ListenAndServe(":3000", nil)
}