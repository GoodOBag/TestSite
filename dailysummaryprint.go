package main

import (
	//"fmt"
	"html/template"
	"net/http"
	"strings"
)

func dailysummaryprint(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	sIndx := strings.Index(url, "?")
	if sIndx == -1 { //if no selection, print something
		t, _ := template.New("").Parse(tpl_noSelection)
		_ = t.Execute(w, "")
	} else {
		if r.Method == "GET" {
			t, err := template.New("").Parse(tpl_print)
			checkError(err, "dailysummaryprint-dailysummaryprint-1")
			err = t.Execute(w, "")
			checkError(err, "dailysummaryprint-dailysummaryprint-2")
		}
		if r.Method == "POST" {
			err := r.ParseForm()
			checkError(err, "dailysummaryprint-dailysummaryprint-3")
			if r.Form["print"][0] == "Print Checklist" {
				http.Redirect(w, r, "/DailySummaryChecklist"+url[sIndx:], http.StatusSeeOther)
			} else if r.Form["print"][0] == "Print Receipt" {
				http.Redirect(w, r, "/DailySummaryReceipt"+url[sIndx:], http.StatusSeeOther)
			} else if r.Form["print"][0] == "Submit Daily Report" {
				http.Redirect(w, r, "/DailySummarySubmit"+url[sIndx:], http.StatusSeeOther)
			}

		}
	}
}

const tpl_print = `
<html>
<head>
<style>
input[type=submit] {
  width: 200px; 
  height: 50px;
  font-size: 20px;
}
</style>
</head>
<body>
<form method="post">
<input type="submit" name="print" value="Print Checklist">
<br><br>
<input type="submit" name="print" value="Print Receipt">
<br><br>
<input type="submit" name="print" value="Submit Daily Report">
</form>
</body>
</html>
`
