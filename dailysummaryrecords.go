package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func dailysummaryrecords(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	sIndx := strings.Index(url, "?")
	if sIndx == -1 { //if no selection, print something
		t, _ := template.New("").Parse(tpl_noSelection)
		_ = t.Execute(w, "")
	} else {
		url = url[sIndx+1:]
		url = strings.Replace(url, "%20", " ", -1)
		selections := strings.Split(url, "+")
		fmt.Println(selections)
	}

}

const tpl_noSelection = `
<html>
<body>
<h2>No data is available</h2>
</body>
</html>
`
