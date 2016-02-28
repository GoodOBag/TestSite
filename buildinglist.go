package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type BuildingsDefault struct {
	BldgName string
	Address  string
	ZipCode  string
	Notes    string
}

func buildinglist(w http.ResponseWriter, r *http.Request) {
	TempBldg, TempAddr, TempZip, TempNotes := getBldgs()

	t, err := template.New("").Parse(tpl_building)
	checkError(err, "buildinglist-buildinglist-1")
	err = t.ExecuteTemplate(w, "t_top", "")
	checkError(err, "buildinglist-buildinglist-2")

	for i, _ := range TempBldg {
		bList := BuildingsDefault{
			BldgName: TempBldg[i],
			Address:  TempAddr[i],
			ZipCode:  TempZip[i],
			Notes:    TempNotes[i],
		}
		t, err := template.New("").Parse(tpl_building)
		checkError(err, "buildinglist-buildinglist-3")
		err = r.ParseForm()
		checkError(err, "buildinglist-buildinglist-4")
		err = t.ExecuteTemplate(w, "t_info", bList)
		checkError(err, "buildinglist-buildinglist-5")
	}

	for i := 0; i < 5; i++ { //add empty columns
		bList := BuildingsDefault{
			BldgName: "",
			Address:  "",
			ZipCode:  "",
			Notes:    "",
		}
		t, err := template.New("").Parse(tpl_building)
		checkError(err, "buildinglist-buildinglist-6")
		err = r.ParseForm()
		checkError(err, "buildinglist-buildinglist-7")
		err = t.ExecuteTemplate(w, "t_info", bList)
		checkError(err, "buildinglist-buildinglist-8")
	}

	t, err = template.New("").Parse(tpl_building)
	checkError(err, "buildinglist-buildinglist-9")
	err = t.ExecuteTemplate(w, "t_bot", "")
	checkError(err, "buildinglist-buildinglist-10")

	if r.Method == "POST" {
		fmt.Println(r.Form)
		_ = updateBldgs(r.Form["Building Name"], r.Form["Address"], r.Form["Zip Code"], r.Form["Notes"])
	}
}

const tpl_building = `
{{define "t_top"}}
<html>
<head>
<title></title>
</head>
<body>

<h2>Building List</h2>

<form action="/BuildingList" method="post">
  <table>
    <tr>
      <td>*Building Name</td>
      <td>*Address</td>
      <td>*Zip Code</td>
      <td>Notes</td>
    </tr>
{{end}}

{{define "t_info"}}
    <tr>
      <td><input type="text" name="Building Name" value="{{.BldgName}}"></td>

      <td><input type="text" name="Address" value="{{.Address}}"></td>  

      <td><input type="text" name="Zip Code" value="{{.ZipCode}}"></td>

      <td><input type="text" name="Notes" value="{{.Notes}}"></td>
    </tr>
{{end}}

{{define "t_bot"}}
  </table>
  <br>
  <span>&nbsp</span>
  <input type="submit" value="Update">
</form>
</body>
</html>
{{end}}
`
