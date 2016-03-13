package main

import (
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type OrderStuc struct {
	Nickname string
	Span     int
	Item     string
	Unit     string
	Amount   float64
	Note     string
	IsFirst  bool
}

type BldgName struct {
	BldgName string
}

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

		_, _, items, units, amounts := getPurchases()
		tempPurchases := ""
		for i, _ := range items {
			tempPurchases = tempPurchases + items[i] + "/" + units[i] + "/" + strconv.FormatFloat(amounts[i], 'f', 2, 64) + "/;"
		}
		if tempPurchases != "" { //remove last ;
			tempPurchases = tempPurchases[:len(tempPurchases)-1]
		}

		_, refNicknames, refBldgs, refOrders := getOrders()
		refNicknames = append(refNicknames, "Purchase")
		refBldgs = append(refBldgs, "Purchase List")
		refOrders = append(refOrders, tempPurchases)

		if r.Method == "GET" {

			t, err := template.New("").Parse(tpl_record)
			checkError(err, "dailysummaryrecords-dailysummaryrecords-1")
			err = t.ExecuteTemplate(w, "t_top", "")
			checkError(err, "dailysummaryrecords-dailysummaryrecords-2")

			//first sort bldg list & nickname from orders
			for _, bldg := range selections { //rotate over selections (bldgs)
				//sort nicknames per selection (bldg)
				tempList := make([]string, 0) //take nicknames per selection (bldg)
				for i, refBldg := range refBldgs {
					if strings.EqualFold(refBldg, bldg) {
						tempList = append(tempList, refNicknames[i])
					}
				}
				bldgName := BldgName{
					BldgName: bldg,
				}
				t, err = template.New("").Parse(tpl_record)
				checkError(err, "dailysummaryrecords-dailysummaryrecords-3")
				err = t.ExecuteTemplate(w, "t_bldgTop", bldgName)
				checkError(err, "dailysummaryrecords-dailysummaryrecords-4")

				sort.Strings(tempList) //sort nicknames
				for _, tempNickname := range tempList {
					for i, refNickname := range refNicknames {

						if strings.EqualFold(tempNickname, refNickname) { //write out orders sorted by nickname per bldg
							tempOrderList := strings.Split(refOrders[i], ";")
							for indx, val := range tempOrderList {
								tempItemList := strings.Split(val, "/")
								tempFloat, _ := strconv.ParseFloat(tempItemList[2], 64)

								iInfo := OrderStuc{
									Nickname: tempNickname,
									Span:     1,
									Item:     tempItemList[0],
									Unit:     tempItemList[1],
									Amount:   tempFloat,
									Note:     tempItemList[3],
									IsFirst:  false,
								}

								if indx == 0 {
									iInfo.IsFirst = true
									iInfo.Span = len(tempOrderList)
								}

								t, err := template.New("").Parse(tpl_record)
								checkError(err, "dailysummaryrecords-dailysummaryrecords-7")
								err = t.ExecuteTemplate(w, "t_loop", iInfo)
								checkError(err, "dailysummaryrecords-dailysummaryrecords-8")
							}
						}
					}
				}
				t, err = template.New("").Parse(tpl_record)
				checkError(err, "dailysummaryrecords-dailysummaryrecords-11")
				err = t.ExecuteTemplate(w, "t_bldgBot", "")
				checkError(err, "dailysummaryrecords-dailysummaryrecords-12")
			}
			t, err = template.New("").Parse(tpl_record)
			checkError(err, "dailysummaryrecords-dailysummaryrecords-13")
			err = t.ExecuteTemplate(w, "t_bot", "")
			checkError(err, "dailysummaryrecords-dailysummaryrecords-14")
		}
		if r.Method == "POST" {
		}
	}
}

const tpl_noSelection = `
<html>
<body>
<h2>No data is available</h2>
</body>
</html>
`

const tpl_record = `
{{define "t_top"}}
<html>
<head>
<style>
table, th, td {
    border: 1px solid black;
    border-collapse: collapse;
}
th, td {
    padding: 5px;
    text-align: left;    
}
input[type=checkbox] {
  width: 30px; 
  height: 30px;
}
input[type=submit] {
  width: 200px; 
  height: 50px;
font-size: 20px;
}
</style>
</head>
<body>

<h1>Summary</h1>
<form method="post">
{{end}}

{{define "t_bldgTop"}}
<table>
<h3>{{.BldgName}}</h3>
  <tr>
    <th>Nickname</th>
    <td>Item</td>
    <td>Unit</td>
    <td>Amount</td>
	<td>Notes</td>
    <td>UnitAssigned</td>
    <td>AmountAssigned</td>
    <td>Delete</td>
  </tr>
{{end}}

{{define "t_loop"}}
<tr>
    {{if .IsFirst}}<th rowspan="{{.Span}}">{{.Nickname}}</th>{{end}}
    <td>{{.Item}}</td>
    <td><input type="text" name="{{.Nickname}}Unit" value="{{.Unit}}"></td>
    <td><input type="text" name="{{.Nickname}}Amount" value="{{.Amount}}"></td>
	<td><input type="text" name="{{.Nickname}}Notes" value="{{.Note}}"></td>
    <td><input type="text" name="{{.Nickname}}Unit2" value=""></td>
    <td><input type="text" name="{{.Nickname}}Amount2" value=""></td>
    <td><input type="checkbox" name="{{.Nickname}}Delete"></td>
</tr>
{{end}}

{{define "t_bldgBot"}}
</table>
<br>
{{end}}

{{define "t_bot"}}
<input type="submit" value="Submit">
</form>
<br><br><br><br><br><br>
</body>
</html>
{{end}}
`
