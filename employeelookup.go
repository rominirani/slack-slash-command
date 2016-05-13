package employeelookup

import (
	"fmt"
	"net/http"
)

var employees = map[string]string{
	"E1":  "Romin Irani, x1000",
	"E2":  "Neil Irani, x1001",
	"E3":  "Aryan Irani, x1002",
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	//Read the Request Parameter "command". This should be /lookup
	command := r.FormValue("command")

	//Read the Request Parameter "text". This will contain the Employee Id
	empid := r.FormValue("text")

	//Ideally do other checks for tokens/username/etc

	if command == "/lookup" {
		emp,err := employees[empid]
		if err != false {
		    fmt.Fprint(w, emp)
		} else {
		    fmt.Fprint(w,"No such employee.")
		}
	} else {
		fmt.Fprint(w, "Something's seem wrong! Try Again.")
	}
}
