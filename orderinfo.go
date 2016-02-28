package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type OrderInfo struct {
	NicknameList []string
	UnitList     []string
	ItemList     []string
}

func orderinfo(w http.ResponseWriter, r *http.Request) {
	TempNicknames := []string{"", "John-Harrison-301", "Jane-StMary-226"}
	TempUnits := []string{"lb", "ea"}
	TempItems := []string{"Apple", "Banana", "Orange", "Cabbage"}

	oinfo := OrderInfo{
		NicknameList: TempNicknames,
		UnitList:     TempUnits,
		ItemList:     TempItems,
	}

	t, err := template.New("").Parse(tpl_order)
	checkError(err, "orderinfo-orderinfo-1")
	err = r.ParseForm()
	checkError(err, "orderinfo-orderinfo-2")
	err = t.ExecuteTemplate(w, "t_top", oinfo)
	checkError(err, "orderinfo-orderinfo-3")

	for i := 0; i < 15; i++ {
		t, err = template.New("").Parse(tpl_order)
		checkError(err, "orderinfo-orderinfo-4")
		err = r.ParseForm()
		checkError(err, "orderinfo-orderinfo-5")
		err = t.ExecuteTemplate(w, "t_mid", oinfo)
		checkError(err, "orderinfo-orderinfo-6")
	}

	t, err = template.New("").Parse(tpl_order)
	checkError(err, "orderinfo-orderinfo-7")
	err = t.ExecuteTemplate(w, "t_bot", oinfo)
	checkError(err, "orderinfo-orderinfo-8")

	if r.Method == "POST" {
		fmt.Println(r.Form)
	}
}

const tpl_order = `
{{define "t_top"}}
<html>
<head>
<title></title>
</head>
<body>

<form action="/OrderInfo" method="post">

  <p><span>&nbsp</span>Nickname</p>
  <span>&nbsp</span>
  <select name="Nickname">
	{{range .NicknameList}}
	  <option>{{.}}</option>
	{{end}}
  </select>
  <br><br>

  <table>

    <tr>
      <td>*Amount</td>
      <td>*Unit</td>
      <td>*Item</td>
	  <td>Notes</td>
    </tr>
{{end}}

{{define "t_mid"}}
    <tr>
      <td><input type="number" name="Amount"></td>
      <span>&nbsp</span>
      <td>
        <select name="Unit">
		  {{range .UnitList}}
		    <option>{{.}}</option>
		  {{end}}          
        </select>
      </td>
      <span>&nbsp</span>
      <td>
        <select name="Item">
          {{range .ItemList}}
		    <option>{{.}}</option>
		  {{end}}
        </select>
      </td>
	  <span>&nbsp</span>
	  <td><input type="text" name="Notes"></td>
    </tr>
{{end}}

{{define "t_bot"}}
  </table>

  <br>
  <span>&nbsp</span>
  <input type="submit" value="Save & Proceed">
</form>

</body>
</html>
{{end}}
`
