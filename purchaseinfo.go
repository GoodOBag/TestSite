package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PinfoDefault struct {
	ItemList []string
	UnitList []string
}

func purchaseinfo(w http.ResponseWriter, r *http.Request) {
	TempItems, _, _, _ := getItems()
	TempUnits := getUnits()

	pinfo := PinfoDefault{
		ItemList: TempItems,
		UnitList: TempUnits,
	}

	t, err := template.New("").Parse(tpl_purchase)
	checkError(err, "purchaseinfo-purchaseinfo-1")
	err = t.ExecuteTemplate(w, "t_top", pinfo)
	checkError(err, "purchaseinfo-purchaseinfo-2")

	for i := 0; i < 10; i++ {
		t, err = template.New("").Parse(tpl_purchase)
		checkError(err, "purchaseinfo-purchaseinfo-3")
		err = r.ParseForm()
		checkError(err, "purchaseinfo-purchaseinfo-4")
		err = t.ExecuteTemplate(w, "t_mid", pinfo)
		checkError(err, "purchaseinfo-purchaseinfo-5")
	}

	t, err = template.New("").Parse(tpl_purchase)
	checkError(err, "purchaseinfo-purchaseinfo-6")
	err = t.ExecuteTemplate(w, "t_bot", pinfo)
	checkError(err, "purchaseinfo-purchaseinfo-7")

	if r.Method == "POST" {
		fmt.Println(r.Form)
		_ = logPurchases(r.Form["Item"], r.Form["Unit"], r.Form["Amount"])
	}
}

const tpl_purchase = `
{{define "t_top"}}
<html>
<head>
<title>Update Purchase</title>
</head>
<body>
<h2>Purchases</h2>

<form action="/PurchaseInfo" method="post">

  <table>

    <tr>
      <td>Item</td>
      <td>Unit</td>
      <td>Amount</td>
    </tr>
{{end}}

{{define "t_mid"}}
    <tr>
      <td>
        <select name="Item">
          {{range .ItemList}}
		    <option>{{.}}</option>
		  {{end}}
        </select>
      </td>
      
      <span>&nbsp</span>
      <td>
        <select name="Unit">
		  {{range .UnitList}}
		    <option>{{.}}</option>
		  {{end}}          
        </select>
      </td>
      <span>&nbsp</span>
      <td><input type="number" step="0.01" min="0" name="Amount"></td>
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
