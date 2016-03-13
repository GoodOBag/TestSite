package main

import (
	"html/template"
	"net/http"
)

type SelectInfo struct {
	List []string
}

func dailysummary(w http.ResponseWriter, r *http.Request) {
	_, _, items, _, _ := getPurchases()
	_, _, bldgs, _ := getOrders()
	sinfo := SelectInfo{}

	if r.Method == "GET" {
		t, err := template.New("").Parse(tpl_ds)
		checkError(err, "dailysummary-dailysummary-1")
		err = t.ExecuteTemplate(w, "t_top", "")
		checkError(err, "dailysummary-dailysummary-2")

		if len(items) == 0 { //no purchase
			t, _ := template.New("").Parse(tpl_ds)
			_ = r.ParseForm()
			_ = t.ExecuteTemplate(w, "t_emptyP", "")
		}

		if len(bldgs) == 0 { //no order
			t, _ := template.New("").Parse(tpl_ds)
			_ = r.ParseForm()
			_ = t.ExecuteTemplate(w, "t_emptyO", "")
		}

		if len(items) != 0 || len(bldgs) != 0 {
			t, _ := template.New("").Parse(tpl_ds)
			_ = r.ParseForm()
			_ = t.ExecuteTemplate(w, "t_mid_top", sinfo)

			if len(items) != 0 {
				tempList := []string{"Purchase List"}
				sinfo = SelectInfo{
					List: tempList,
				}

				t, err := template.New("").Parse(tpl_ds)
				checkError(err, "dailysummary-dailysummary-3")
				err = t.ExecuteTemplate(w, "t_mid", sinfo)
				checkError(err, "dailysummary-dailysummary-4")
			}

			if len(bldgs) != 0 {
				sinfo = SelectInfo{
					List: uniqueStrings(bldgs),
				}

				t, err := template.New("").Parse(tpl_ds)
				checkError(err, "dailysummary-dailysummary-5")
				err = t.ExecuteTemplate(w, "t_mid", sinfo)
				checkError(err, "dailysummary-dailysummary-6")
			}

			t, _ = template.New("").Parse(tpl_ds)
			_ = r.ParseForm()
			_ = t.ExecuteTemplate(w, "t_mid_bot", "")
		}

		t, err = template.New("").Parse(tpl_ds)
		checkError(err, "dailysummary-dailysummary-7")
		err = t.ExecuteTemplate(w, "t_end", sinfo)
		checkError(err, "dailysummary-dailysummary-8")
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		checkError(err, "dailysummary-dailysummary-9")
		tempUrl := ""
		if r.Form["type"][0] == "P" {
			tempUrl = "/DailySummaryPrint"
		} else {
			tempUrl = "/DailySummaryRecords"
		}

		if len(r.Form["choices"]) > 0 {
			tempUrl = tempUrl + "?"
			for _, val := range r.Form["choices"] {
				tempUrl = tempUrl + val + "+"
			}
			tempUrl = tempUrl[:len(tempUrl)-1]
		}
		http.Redirect(w, r, tempUrl, http.StatusSeeOther)
	}
}

const tpl_ds = `
{{define "t_top"}}
<html>
<style>
input[type=checkbox] {
  width: 20px; 
  height: 20px;
}
span {
  font-size: 170%;
}
p {
  font-size: 200%;
}
</style>
<body>
{{end}}

{{define "t_emptyP"}}
<br><br>
<p>No purchase is made</p>
<br><br>
{{end}}

{{define "t_emptyO"}}
<br><br>
<p>No order is made</p>
<br><br>
{{end}}

{{define "t_mid_top"}}
<form method="POST">
<br>
{{end}}

{{define "t_mid"}}
  {{range .List}}
    <input type="checkbox" name="choices" value="{{.}}"> <span>{{.}}</span><br>
  {{end}}
{{end}}

{{define "t_mid_bot"}}
  <br>
  <input type="radio" name="type" value="P" checked>For print&#160&#160&#160&#160
  <input type="radio" name="type" value="E">For edit
  <br><br>
  <input type="submit" value="Submit" size="20">
</form>
{{end}}

{{define "t_end"}}
</body>
</html>
{{end}}
`
