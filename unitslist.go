package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UnitInfo struct {
	Units []string
}

func unitslist(w http.ResponseWriter, r *http.Request) {
	ulist := UnitInfo{
		Units: getUnits(),
	}

	for i := 0; i < 5; i++ { //leave emply slots for new units
		ulist.Units = append(ulist.Units, "")
	}

	t, err := template.New("").Parse(tpl_unit)
	checkError(err, "unitslist-unitslist-1")
	err = r.ParseForm()
	checkError(err, "unitslist-unitslist-2")
	err = t.Execute(w, ulist)
	checkError(err, "unitslist-unitslist-3")

	if r.Method == "POST" {
		fmt.Println(r.Form["Unit"])
		_ = updateUnits(r.Form["Unit"])
	}
}

const tpl_unit = `
<html>
<head>
<title>Update Units</title>
</head>
<body>
<h2>Units</h2>
<form action="/UnitsList" method="post">

  <table>

    <tr>
      <td>Units</td>
    </tr>

{{range .Units}}
    <tr>
      <td><input type="text" name="Unit" value="{{.}}"></td>
    </tr>
{{end}}
  </table>

  <br>
  <span>&nbsp</span>
  <input type="submit" value="Update Units">
</form>

</body>
</html>

`
