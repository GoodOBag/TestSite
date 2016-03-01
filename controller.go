package main

import (
	"strconv"
)

func getUnits() []string {
	units := unittableGet()
	return units
}

func updateUnits(units []string) bool {
	unitsMod := make([]string, 0)
	for _, unit := range units {
		if unit != "" {
			unitsMod = append(unitsMod, unit)
		}
	}
	_ = unittableReplace(unitsMod)
	return true
}

func getItems() (items []string, units []string, unitPrices []float64, notes []string) {
	items, units, unitPrices, notes = itemtableGet()
	return
}

func updateItems(itemsRaw []string, unitsRaw []string, unitPricesRaw []string, notesRaw []string) bool { //unitprice read from map as string
	items := make([]string, 0)
	units := make([]string, 0)
	unitPrices := make([]float64, 0)
	notes := make([]string, 0)

	for i, _ := range itemsRaw {
		if itemsRaw[i] != "" && unitsRaw[i] != "" && unitPricesRaw[i] != "" && notesRaw[i] != "" {
			items = append(items, itemsRaw[i])
			units = append(units, unitsRaw[i])
			tempPrice, err := strconv.ParseFloat(unitPricesRaw[i], 64)
			checkError(err, "controller-updateItems")
			unitPrices = append(unitPrices, tempPrice)
			notes = append(notes, notesRaw[i])
		}
	}
	_ = itemtableReplace(items, units, unitPrices, notes)
	return true
}

func getBldgs() (regdates []int, bldgs []string, addrs []string, zips []string, notes []string) {
	regdates, bldgs, addrs, zipsInt, notes := bldgtableGet()
	for _, val := range zipsInt { //convert zips to string as the form returns string
		zips = append(zips, strconv.Itoa(val))
	}
	return
}

func updateBldgs(bldgsRaw []string, addrsRaw []string, zipsRaw []string, notesRaw []string) bool {
	bldgs := make([]string, 0)
	addrs := make([]string, 0)
	zips := make([]string, 0)
	notes := make([]string, 0)
	for i, _ := range bldgsRaw {
		if bldgsRaw[i] != "" && addrsRaw[i] != "" && zipsRaw[i] != "" {
			bldgs = append(bldgs, bldgsRaw[i])
			addrs = append(addrs, addrsRaw[i])
			zips = append(zips, zipsRaw[i])
			notes = append(notes, notesRaw[i])
		}
	}

	refDates, refBldgs, _, _, _ := getBldgs()
	regdates := make([]int, 0)
	for _, val := range bldgs {
		tempIndx := findStrInSlice(val, refBldgs)
		if tempIndx != -1 { //if bldg name exist, use existing date
			regdates = append(regdates, refDates[tempIndx])
		} else { //if new bldg name, use current date
			regdates = append(regdates, getCurrentDate())
		}
	}
	intZips := make([]int, 0)
	for _, val := range zips {
		intZip, err := strconv.Atoi(val)
		checkError(err, "controller-updateBldgs")
		intZips = append(intZips, intZip)
	}
	_ = bldgtableReplace(regdates, bldgs, addrs, intZips, notes)
	return true
}

func getCustomers() (nicknames []string, phones []string, bldgs []string, rooms []string, notes []string) {
	nicknames = []string{"John", "Rick", "Sam", "Ali"}
	phones = []string{"2346", "235", "236", "247"}
	bldgs = []string{"E", "C", "A", "B"}
	rooms = []string{"361q", "23y", "56w", "qa354"}
	notes = []string{"54yw", "q34y", "qw34", "e5i"}
	return nicknames, phones, bldgs, rooms, notes
}

func updateCustomers(nicknames []string, phones []string, bldgs []string, rooms []string, notes []string) bool {
	_, _, _, _, _ = nicknames, phones, bldgs, rooms, notes
	return true
}

func logOrders(nickname string, items []string, units []string, amounts []float64) bool {
	_, _, _, _ = nickname, items, units, amounts
	return true
}

func logPurchases(items []string, units []string, amounts []string) bool {
	_, _, _ = items, units, amounts
	return true
}

func getAllOrders() {
	//for daily reports
}

func updateAllOrders() {
	//for daily reports
}

func getAllPurchases() {
	//for daily reports
}

func updateAllPurchases() {
	//for daily reports
}
