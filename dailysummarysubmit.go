package main

import (
	//"fmt"
	"html/template"
	"net/http"
	"strings"
)

func dailysummarysubmit(w http.ResponseWriter, r *http.Request) {
	authCheck(w, r)
	url := r.URL.String()
	sIndx := strings.Index(url, "?")
	if sIndx == -1 { //if no selection, print something
		t, _ := template.New("").Parse(tpl_noSelection)
		_ = t.Execute(w, "")
	} else {
		url = url[sIndx+1:]
		url = strings.Replace(url, "%20", " ", -1)
		selections := strings.Split(url, "+")
		ids, nicknames, bldgs, orders := getOrders()
		uniqueBldgs := uniqueStrings(bldgs)
		if len(selections) != len(uniqueBldgs) {
			t, _ := template.New("").Parse(tpl_wrong)
			_ = t.Execute(w, "")
		} else {
			_ = ids
			_ = nicknames
			_ = orders
		}
	}
}

const tpl_wrong = `
<html>
<body>
<h1>Don't hack me!!!</h1>
</body>
</html>
`
