package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

/*Available Actions:
Get - get entire table
GetDay - get rows for the latest day
Update - update a row
Append - append a new row to the end
Delete - delete a row
Replace - flush table and input new data into table
*/

//unit
func unittableGet() (units []string) {
	rows, err := db.Query("SELECT * FROM unittable")
	checkError(err, "model-unittableGet-1")

	var unit string
	for rows.NEXT() {
		err = rows.Scan(&unit)
		checkError(err, "model-unittableGet-2")
		units = append(units, unit)
	}
	return
}

func unittableReplace(units []string) bool {
	_, err := db.Exec("delete from unittable")
	checkError(err, "model-unittableReplace-1")

	stmt, err := db.Prepare("insert into unittable(unit) values(?)")
	checkError(err, "model-unittableReplace-2")
	defer stmt.Close()
	for _, unit := range units {
		_, err = stmt.Exec(unit)
		checkError(err, "model-unittableReplace-3")
	}
	return true
}

//item
func itemtableGet() (items []string, notes []string) {
	rows, err := db.Query("SELECT * FROM itemtable")
	checkError(err, "model-itemtableGet-1")

	var item, note string
	for rows.NEXT() {
		err = rows.Scan(&item, &note)
		checkError(err, "model-itemtableGet-2")
		items = append(items, item)
		notes = append(notes, note)
	}
	return
}

func itemtableReplace(items []string, units []string, unitprice []float64, notes []string) bool {
	_, err := db.Exec("delete from itemtable")
	checkError(err, "model-itemtableReplace-1")

	stmt, err := db.Prepare("insert into itemtable(item,unit,unitprice,notes) values(?,?,?,?)")
	checkError(err, "model-itemtableReplace-2")
	defer stmt.Close()
	for i, _ := range items {
		_, err = stmt.Exec(items[i], units[i], unitprice[i], notes[i])
		checkError(err, "model-unittableReplace-3")
	}
	return true
}

//bldg
func bldgtableGet() (regdates []int, bldgs []string, addrs []string, zips []int, notes []string) {
	rows, err := db.Query("SELECT * FROM bldgtable")
	checkError(err, "model-bldgtableGet-1")

	var regdate, zip int
	var bldg, addr, note string
	for rows.NEXT() {
		err = rows.Scan(&regdate, &bldg, &addr, &zip, &note)
		checkError(err, "model-bldgtableGet-2")
		regdates = append(regdates, regdate)
		bldgs = append(bldgs, bldg)
		addrs = append(addrs, addr)
		zips = append(zips, zip)
		notes = append(notes, note)
	}
	return
}

func bldgtableReplace() bool {
	return true
}

//customer
func customertableGet() (id []int, regdates []int, nicknames []string, phones []int64, bldg []string, room []string, notes []string) {

}

func customertableAppend(regdate int, nickname string, phone int64, bldg string, room string, note string) bool {

}

func customertableUpdate(id int, regdate int, nickname string, phone int64, bldg string, room string, note string) bool {

}

func customertableDelete(id int) bool(){
	
}

//order
func ordertableGetDay()(){
	
}

func ordertableAppend() bool{
	
}

func ordertableUpdate() bool{
	
}

func ordertableDelete() bool{
	
}

//purchase
func purchasetableGetDay()(){
	
}
func purchasetableAppend() bool{
	
}

func purchasetableUpdate() bool{
	
}

func purchasetableDelete() bool{
	
}