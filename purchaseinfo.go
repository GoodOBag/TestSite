package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

type PinfoDefault struct {
	ItemList []string
	UnitList []string
	Item     string
	Unit     string
	Amount   float64
}

func purchaseinfo(w http.ResponseWriter, r *http.Request) {
	authCheck(w, r)
	TempItems, _, _, _ := getItems()
	TempUnits := getUnits()

	pinfo := PinfoDefault{
		ItemList: TempItems,
		UnitList: TempUnits,
		Item:     "",
		Unit:     "",
		Amount:   0,
	}

	_, _, items, units, amounts := getPurchases()

	t, err := template.New("").Funcs(template.FuncMap{
		"isMatching":   isMatching,
		"isValidPrice": isValidPrice,
	}).Parse(tpl_purchase)
	checkError(err, "purchaseinfo-purchaseinfo-1")
	err = t.ExecuteTemplate(w, "t_top", pinfo)
	checkError(err, "purchaseinfo-purchaseinfo-2")
	for i, _ := range items {
		pinfo.Item = items[i]
		pinfo.Unit = units[i]
		pinfo.Amount = amounts[i]

		t, err = template.New("").Funcs(template.FuncMap{
			"isMatching":   isMatching,
			"isValidPrice": isValidPrice,
		}).Parse(tpl_purchase)
		checkError(err, "purchaseinfo-purchaseinfo-3")
		err = r.ParseForm()
		checkError(err, "purchaseinfo-purchaseinfo-4")
		err = t.ExecuteTemplate(w, "t_mid", pinfo)
		checkError(err, "purchaseinfo-purchaseinfo-5")
	}

	for i := 0; i < 10; i++ {
		pinfo.Item = ""
		pinfo.Unit = ""
		pinfo.Amount = 0

		t, err = template.New("").Funcs(template.FuncMap{
			"isMatching":   isMatching,
			"isValidPrice": isValidPrice,
		}).Parse(tpl_purchase)
		checkError(err, "purchaseinfo-purchaseinfo-3")
		err = r.ParseForm()
		checkError(err, "purchaseinfo-purchaseinfo-4")
		err = t.ExecuteTemplate(w, "t_mid", pinfo)
		checkError(err, "purchaseinfo-purchaseinfo-5")
	}

	t, err = template.New("").Funcs(template.FuncMap{
		"isMatching":   isMatching,
		"isValidPrice": isValidPrice,
	}).Parse(tpl_purchase)
	checkError(err, "purchaseinfo-purchaseinfo-6")
	err = t.ExecuteTemplate(w, "t_bot", pinfo)
	checkError(err, "purchaseinfo-purchaseinfo-7")

	if r.Method == "POST" {
		//fmt.Println(r.Form)
		_ = logPurchases(r.Form["Item"], r.Form["Unit"], r.Form["Amount"])
	}
}

const tpl_purchase = `
{{define "t_top"}}
<html>
<head>
<title>Update Purchase</title>
<script src="http://code.jquery.com/jquery-1.9.1.js"></script>
<script>
  $(function () {
    $('form').on('submit', function (e) {
      e.preventDefault();
      $.ajax({
        type: 'post',
        data: $('form').serialize(),
      });
    });
  });
</script>
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
		  {{$item := .Item}}
          {{range .ItemList}}
		    <option {{if isMatching $item .}}selected{{end}}>{{.}}</option>
		  {{end}}
        </select>
      </td>
      
      <span>&nbsp</span>
      <td>
        <select name="Unit">
		  {{$unit := .Unit}}
		  {{range .UnitList}}
		    <option {{if isMatching $unit .}}selected{{end}}>{{.}}</option>
		  {{end}}          
        </select>
      </td>
      <span>&nbsp</span>
      <td><input type="number" step="0.01" min="0" name="Amount" value="{{if isValidPrice .Amount}}{{.Amount}}{{end}}"></td>
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
