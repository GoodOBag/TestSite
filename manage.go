package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) { //login for manage subdomain
	if r.Method == "GET" {
		t, err := template.ParseFiles("manage/login.html")
		checkError(err, "manage-login")
		t.Execute(w, nil)
	} else { //POST
		r.ParseForm()
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
		//fmt.Println(r.Form)

		//add authentication verication function here
		if r.Form["username"][0] == "lalala" && r.Form["password"][0] == "bababa" {
			authIssue(w, r.Form["username"][0])
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	}
}

func manage(w http.ResponseWriter, r *http.Request) {
	authCheck(w, r)
	t, err := template.ParseFiles("manage/manage.html")
	checkError(err, "manage-manage")
	t.Execute(w, nil)
	r.ParseForm()
}
