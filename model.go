package main

import (
	"database/sql"

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
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-unittableGet-1")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM unittable")
	checkError(err, "model-unittableGet-2")

	var unit string
	for rows.Next() {
		err = rows.Scan(&unit)
		checkError(err, "model-unittableGet-3")
		units = append(units, unit)
	}
	return
}

func unittableReplace(units []string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-unittableReplace-1")
	defer db.Close()
	_, err = db.Exec("delete from unittable")
	checkError(err, "model-unittableReplace-2")

	stmt, err := db.Prepare("insert into unittable(unit) values(?)")
	checkError(err, "model-unittableReplace-3")
	defer stmt.Close()
	for _, unit := range units {
		_, err = stmt.Exec(unit)
		checkError(err, "model-unittableReplace-4")
	}
	return true
}

//item
func itemtableGet() (items []string, units []string, unitprices []float64, notes []string) {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-itemtableGet-1")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM itemtable")
	checkError(err, "model-itemtableGet-2")

	var unitprice float64
	var item, unit, note string
	for rows.Next() {
		err = rows.Scan(&item, &unit, &unitprice, &note)
		checkError(err, "model-itemtableGet-3")
		items = append(items, item)
		units = append(units, unit)
		unitprices = append(unitprices, unitprice)
		notes = append(notes, note)
	}
	return
}

func itemtableReplace(items []string, units []string, unitprices []float64, notes []string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-itemtableReplace-1")
	defer db.Close()
	_, err = db.Exec("delete from itemtable")
	checkError(err, "model-itemtableReplace-2")

	stmt, err := db.Prepare("insert into itemtable(item,unit,unitprice,notes) values(?,?,?,?)")
	checkError(err, "model-itemtableReplace-3")
	defer stmt.Close()
	for i, _ := range items {
		_, err = stmt.Exec(items[i], units[i], unitprices[i], notes[i])
		checkError(err, "model-itemtableReplace-4")
	}
	return true
}

//bldg
func bldgtableGet() (regdates []int, bldgs []string, addrs []string, zips []int, notes []string) {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-bldgtableGet-1")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM bldgtable")
	checkError(err, "model-bldgtableGet-2")

	var regdate, zip int
	var bldg, addr, note string
	for rows.Next() {
		err = rows.Scan(&regdate, &bldg, &addr, &zip, &note)
		checkError(err, "model-bldgtableGet-3")
		regdates = append(regdates, regdate)
		bldgs = append(bldgs, bldg)
		addrs = append(addrs, addr)
		zips = append(zips, zip)
		notes = append(notes, note)
	}
	return
}

func bldgtableReplace(regdates []int, bldgs []string, addrs []string, zips []int, notes []string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-bldgtableReplace-1")
	defer db.Close()
	_, err = db.Exec("delete from bldgtable")
	checkError(err, "model-bldgtableReplace-2")

	stmt, err := db.Prepare("insert into bldgtable(regdate, bldg, addr, zip, notes) values(?,?,?,?,?)")
	checkError(err, "model-bldgtableReplace-3")
	defer stmt.Close()
	for i, _ := range regdates {
		_, err = stmt.Exec(regdates[i], bldgs[i], addrs[i], zips[i], notes[i])
		checkError(err, "model-bldgtableReplace-4")
	}
	return true
}

//customer
func customertableGet() (ids []int, regdates []int, nicknames []string, phones []int64, bldgs []string, rooms []string, notes []string) {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-customertableGet-1")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM customertable")
	checkError(err, "model-customertableGet-2")

	var id, regdate int
	var phone int64
	var nickname, bldg, room, note string
	for rows.Next() {
		err = rows.Scan(&id, &regdate, &nickname, &phone, &bldg, &room, &note)
		checkError(err, "model-customertableGet-3")
		ids = append(ids, id)
		regdates = append(regdates, regdate)
		nicknames = append(nicknames, nickname)
		phones = append(phones, phone)
		bldgs = append(bldgs, bldg)
		rooms = append(rooms, room)
		notes = append(notes, note)
	}
	return
}

func customertableAppend(regdate int, nickname string, phone int64, bldg string, room string, note string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-customertableAppend-1")
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO customertable(regdate, nickname, phone, bldg, room, notes) values(?,?,?,?,?,?)")
	checkError(err, "model-customertableAppend-2")

	_, err = stmt.Exec(regdate, nickname, phone, bldg, room, note)
	checkError(err, "model-customertableAppend-3")
	return true
}

func customertableUpdate(id int, regdate int, nickname string, phone int64, bldg string, room string, note string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-customertableUpdate-1")
	defer db.Close()
	stmt, err := db.Prepare("update customertable set regdate=?, nickname=?, phone=?, bldg=?, room=?, notes=?  where id=?")
	checkError(err, "model-customertableUpdate-2")

	_, err = stmt.Exec(regdate, nickname, phone, bldg, room, note, id)
	checkError(err, "model-customertableUpdate-3")
	return true
}

func customertableDelete(id int) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-customertableDelete-1")
	defer db.Close()
	stmt, err := db.Prepare("delete from customertable where id=?")
	checkError(err, "model-customertableDelete-2")

	_, err = stmt.Exec(id)
	checkError(err, "model-customertableDelete-3")
	return true
}

//order
func ordertableGetDay(orderdate int) (ids []int, nicknames []string, orderlists []string) {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-ordertableGetDay-1")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM ordertable")
	checkError(err, "model-ordertableGetDay-2")

	var id, tempdate int
	var nickname, orderlist string
	for rows.Next() {
		err = rows.Scan(&id, &nickname, &tempdate, &orderlist)
		checkError(err, "model-ordertableGetDay-3")
		if tempdate == orderdate {
			ids = append(ids, id)
			nicknames = append(nicknames, nickname)
			orderlists = append(orderlists, orderlist)
		}
	}
	return
}

func ordertableAppend(nickname string, orderdate int, orderlist string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-ordertableAppend-1")
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO ordertable(nickname, orderdate, orderlist) values(?,?,?)")
	checkError(err, "model-ordertableAppend-2")

	_, err = stmt.Exec(nickname, orderdate, orderlist)
	checkError(err, "model-ordertableAppend-3")
	return true
}

func ordertableUpdate(id int, nickname string, orderdate int, orderlist string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-ordertableUpdate-1")
	defer db.Close()
	stmt, err := db.Prepare("update ordertable set nickname=?, orderdate=?, orderlist=?  where id=?")
	checkError(err, "model-ordertableUpdate-2")

	_, err = stmt.Exec(nickname, orderdate, orderlist, id)
	checkError(err, "model-ordertableUpdate-3")
	return true
}

func ordertableDelete(id int) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-ordertableDelete-1")
	defer db.Close()
	stmt, err := db.Prepare("delete from ordertable where id=?")
	checkError(err, "model-ordertableDelete-2")

	_, err = stmt.Exec(id)
	checkError(err, "model-ordertableDelete-3")
	return true
}

//purchase
func purchasetableGetDay(purchasedate int) (ids []int, purchaselists []string) {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-purchasetableGetDay-1")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM purchasetable")
	checkError(err, "model-purchasetableGetDay-2")

	var id, tempdate int
	var purchaselist string
	for rows.Next() {
		err = rows.Scan(&id, &tempdate, &purchaselist)
		checkError(err, "model-purchasetableGetDay-3")
		if tempdate == purchasedate {
			ids = append(ids, id)
			purchaselists = append(purchaselists, purchaselist)
		}
	}
	return
}
func purchasetableAppend(purchasedate int, purchaselist string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-purchasetableAppend-1")
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO purchasetable(purchasedate, purchaselist) values(?,?)")
	checkError(err, "model-purchasetableAppend-2")

	_, err = stmt.Exec(purchasedate, purchaselist)
	checkError(err, "model-purchasetableAppend-3")
	return true
}

func purchasetableUpdate(id int, purchasedate int, purchaselist string) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-purchasetableUpdate-1")
	defer db.Close()
	stmt, err := db.Prepare("update purchasetable set purchasedate=?, purchaselist=?  where id=?")
	checkError(err, "model-purchasetableUpdate-2")

	_, err = stmt.Exec(purchasedate, purchaselist, id)
	checkError(err, "model-purchasetableUpdate-3")
	return true
}

func purchasetableDelete(id int) bool {
	db, err := sql.Open("sqlite3", "./database/record.db")
	checkError(err, "model-purchasetableDelete-1")
	defer db.Close()
	stmt, err := db.Prepare("delete from purchasetable where id=?")
	checkError(err, "model-purchasetableDelete-2")

	_, err = stmt.Exec(id)
	checkError(err, "model-purchasetableDelete-3")
	return true
}
