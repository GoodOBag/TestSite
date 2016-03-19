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
	Unit2    string
	Amount2  float64
	IsFirst  bool
}

type BldgName struct {
	BldgName string
}

type Collection struct {
	Nicknames []string //for debug purposes
	Items     []string
	Units     []string
	Amounts   []string
	Notes     []string
	Units2    []string
	Amounts2  []string
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

		_, refNicknames, refBldgs, refOrders := getOrders()

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
				tempList = uniqueStrings(tempList)
				sort.Strings(tempList) //sort unique nicknames
				for _, tempNickname := range tempList {
					for i, refNickname := range refNicknames {
						if strings.EqualFold(tempNickname, refNickname) { //write out orders sorted by nickname per bldg
							tempOrderList := strings.Split(refOrders[i], ";")
							for indx, val := range tempOrderList {
								tempItemList := strings.Split(val, ",")
								tempFloat, _ := strconv.ParseFloat(tempItemList[2], 64)
								tempFloat2, _ := strconv.ParseFloat(tempItemList[5], 64)

								iInfo := OrderStuc{
									Nickname: tempNickname,
									Span:     1,
									Item:     tempItemList[0],
									Unit:     tempItemList[1],
									Amount:   tempFloat,
									Note:     tempItemList[3],
									Unit2:    tempItemList[4],
									Amount2:  tempFloat2,
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
			err := r.ParseForm()
			checkError(err, "dailysummaryrecords-dailysummaryrecords-15")
			status := checkboxParser(r.Form["Delete"])
			nicknames := r.Form["Nickname"]
			spansStr := r.Form["Span"]
			spans := make([]int, 0)
			for _, val := range spansStr {
				tempInt, err := strconv.ParseInt(val, 10, 64)
				checkError(err, "dailysummaryrecords-dailysummaryrecords-a")
				spans = append(spans, int(tempInt))
			}

			cnt := 0
			sCnt := 0
			collected := Collection{
				Nicknames: make([]string, 0),
				Items:     make([]string, 0),
				Units:     make([]string, 0),
				Amounts:   make([]string, 0),
				Notes:     make([]string, 0),
				Units2:    make([]string, 0),
				Amounts2:  make([]string, 0),
			}
			nMap := make(map[string][]int)
			uniqueNicknames := make([]string, 0)
			for i, n := range nicknames {
				for j := 0; j < spans[i]; j++ {
					if status[sCnt] == 0 {
						collected.Nicknames = append(collected.Nicknames, n)
						collected.Items = append(collected.Items, r.Form["Item"][sCnt])
						collected.Units = append(collected.Units, r.Form["Unit"][sCnt])
						if r.Form["Amount"][sCnt] != "" {
							collected.Amounts = append(collected.Amounts, r.Form["Amount"][sCnt])
						} else {
							collected.Amounts = append(collected.Amounts, "0")
						}
						collected.Notes = append(collected.Notes, r.Form["Notes"][sCnt])
						collected.Units2 = append(collected.Units2, r.Form["Unit2"][sCnt])
						if r.Form["Amount2"][sCnt] != "" {
							collected.Amounts2 = append(collected.Amounts2, r.Form["Amount2"][sCnt])
						} else {
							collected.Amounts2 = append(collected.Amounts2, "0")
						}

						if nMap[n] != nil {
							nMap[n] = append(nMap[n], cnt)
						} else {
							nMap[n] = []int{cnt}
							uniqueNicknames = append(uniqueNicknames, n)
						}
						cnt += 1
					}
					sCnt += 1
				}
			}

			for _, n := range uniqueNicknames {
				wItems := make([]string, 0)
				wUnits := make([]string, 0)
				wAmounts := make([]float64, 0)
				wNotes := make([]string, 0)
				wUnits2 := make([]string, 0)
				wAmounts2 := make([]float64, 0)

				tempLocs := nMap[n]
				tempCheck := make([]bool, len(tempLocs))
				for z, _ := range tempCheck {
					tempCheck[z] = false
				}
				for z, _ := range tempLocs {
					if tempCheck[z] == true { //if previously selected
						//Do nothing
					} else if z == len(tempLocs)-1 && tempCheck[z] == false { //if it's the last one & not selected before
						wItems = append(wItems, collected.Items[tempLocs[z]])
						wUnits = append(wUnits, collected.Units[tempLocs[z]])
						wNotes = append(wNotes, collected.Notes[tempLocs[z]])
						wUnits2 = append(wUnits2, collected.Units2[tempLocs[z]])

						tempAmount, err := strconv.ParseFloat(collected.Amounts[tempLocs[z]], 64)
						checkError(err, "dailysummaryrecords-dailysummaryrecords-16")
						tempAmount2, err := strconv.ParseFloat(collected.Amounts2[tempLocs[z]], 64)
						checkError(err, "dailysummaryrecords-dailysummaryrecords-17")
						wAmounts = append(wAmounts, tempAmount)
						wAmounts2 = append(wAmounts2, tempAmount2)
					} else { //if not selected & not the last one
						wAmount, err := strconv.ParseFloat(collected.Amounts[tempLocs[z]], 64)
						checkError(err, "dailysummaryrecords-dailysummaryrecords-18")
						wAmount2, err := strconv.ParseFloat(collected.Amounts2[tempLocs[z]], 64)
						checkError(err, "dailysummaryrecords-dailysummaryrecords-19")
						tempNotes := collected.Notes[tempLocs[z]]
						for x := z + 1; x < len(tempLocs); x++ {
							if tempCheck[x] == false {
								//see if Item, Unit and Unit2 are matching
								if strings.EqualFold(collected.Items[tempLocs[z]], collected.Items[tempLocs[x]]) &&
									strings.EqualFold(collected.Units[tempLocs[z]], collected.Units[tempLocs[x]]) &&
									strings.EqualFold(collected.Units2[tempLocs[z]], collected.Units2[tempLocs[x]]) {
									tempAmount, err := strconv.ParseFloat(collected.Amounts[tempLocs[x]], 64)
									checkError(err, "dailysummaryrecords-dailysummaryrecords-20")
									tempAmount2, err := strconv.ParseFloat(collected.Amounts2[tempLocs[x]], 64)
									checkError(err, "dailysummaryrecords-dailysummaryrecords-21")
									wAmount += tempAmount
									wAmount2 += tempAmount2
									if collected.Notes[tempLocs[x]] != "" {
										tempNotes = tempNotes + " | " + collected.Notes[tempLocs[x]]
									}
									tempCheck[x] = true
								}
							}
						}
						wItems = append(wItems, collected.Items[tempLocs[z]])
						wUnits = append(wUnits, collected.Units[tempLocs[z]])
						wUnits2 = append(wUnits2, collected.Units2[tempLocs[z]])
						wAmounts = append(wAmounts, wAmount)
						wAmounts2 = append(wAmounts2, wAmount2)
						wNotes = append(wNotes, tempNotes)
					}
				}
				updateOrders(n, wItems, wUnits, wAmounts, wNotes, wUnits2, wAmounts2)
			}

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
    {{if .IsFirst}}
	  <th rowspan="{{.Span}}">{{.Nickname}}</th>
	  <input type="hidden" name="Nickname" value="{{.Nickname}}" />
	  <input type="hidden" name="Span" value="{{.Span}}" />
	{{end}}
    <td><input type="text" name="Item" value="{{.Item}}" /></td>
    <td><input type="text" name="Unit" value="{{.Unit}}" /></td>
    <td><input type="text" name="Amount" value="{{.Amount}}" /></td>
	<td><input type="text" name="Notes" value="{{.Note}}" /></td>
    <td><input type="text" name="Unit2" value="{{.Unit2}}" /></td>
    <td><input type="text" name="Amount2" value="{{.Amount2}}" /></td>
    <td><input type="hidden" name="Delete" value="0" /><input type="checkbox" name="Delete" value="1" /></td>
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