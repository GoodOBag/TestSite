package main

import (
	"strconv"
	"strings"
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
		if itemsRaw[i] != "" && unitsRaw[i] != "" && unitPricesRaw[i] != "" {
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

func getCustomers() (ids []int, nicknames []string, phones []string, bldgs []string, rooms []string, notes []string) {
	ids, _, nicknames, phonesInt64, bldgs, rooms, notes := customertableGet()
	phonesRaw := make([]string, len(phonesInt64))
	for i, val := range phonesInt64 {
		phonesRaw[i] = strconv.FormatInt(val, 10)
		phones = append(phones, "("+phonesRaw[i][0:3]+") "+phonesRaw[i][3:6]+"-"+phonesRaw[i][6:])
	}
	return
}

func updateCustomers(nicknames []string, phones []string, bldgs []string, rooms []string, notes []string) bool {
	refIds, refDates, refNicknames, refPhones, refBldg, refRooms, refNotes := customertableGet()

	//check any nickname in ref but not in input, remove this row as user changed nickname to ""
	for i, val := range refNicknames {
		if findStrInSlice(val, nicknames) == -1 {
			customertableDelete(refIds[i])
		}
	}

	for i, _ := range phones {
		phones[i] = strings.Replace(phones[i], "(", "", -1) //remove "()- " from []phones
		phones[i] = strings.Replace(phones[i], ")", "", -1)
		phones[i] = strings.Replace(phones[i], "-", "", -1)
		phones[i] = strings.Replace(phones[i], " ", "", -1)
		if len(phones[i]) == 11 { //remove the leading "1" if exists
			phones[i] = phones[i][1:]
		}
		if len(phones[i]) != 10 { //remove invalid phone numbers with length != 10
			phones[i] = ""
		}
		_, err := strconv.ParseInt(phones[i], 10, 64) //remove phone numbers contains non-numbers
		if err != nil {
			phones[i] = ""
		}
		if phones[i] != "" && string(phones[i][0]) == "-" { //remove negative phone numbers
			phones[i] = ""
		}

		if nicknames[i] != "" && phones[i] != "" && bldgs[i] != "" && rooms[i] != "" { //ignore inputs with any of those fields empty
			tempPhone, _ := strconv.ParseInt(phones[i], 10, 64) //guaranteed no error, checked previously
			tempIndx := -1
			for j, _ := range refNicknames {
				if strings.EqualFold(nicknames[i], refNicknames[j]) {
					tempIndx = j
					break
				}
			}

			if tempIndx != -1 { //existing record
				if !(tempPhone == refPhones[tempIndx] && strings.EqualFold(bldgs[i], refBldg[tempIndx]) &&
					strings.EqualFold(rooms[i], refRooms[tempIndx]) && strings.EqualFold(notes[i], refNotes[tempIndx])) {
					//if not exactly the same, update the row but keep original ID and registration date
					_ = customertableUpdate(refIds[tempIndx], refDates[tempIndx], refNicknames[tempIndx], tempPhone, bldgs[i], rooms[i], notes[i])
				}
			} else { //new record
				customertableAppend(getCurrentDate(), nicknames[i], tempPhone, bldgs[i], rooms[i], notes[i])
			}

		} else if nicknames[i] != "" { //delete record if exists
			for j, _ := range refNicknames {
				if strings.EqualFold(nicknames[i], refNicknames[j]) {
					customertableDelete(refIds[j])
					break
				}
			}
		}
	}

	return true
}

func logOrders(nicknames []string, items []string, units []string, amounts []string, notes []string) bool {
	nickname := nicknames[0]
	if nickname != "" {
		orderList := ""
		for i, _ := range items {
			if items[i] != "" && units[i] != "" && amounts[i] != "" {
				orderList = orderList + items[i] + "/" + units[i] + "/" + amounts[i] + "/" + notes[i] + ";"
			}
		}
		if orderList != "" {
			orderList = orderList[:len(orderList)-1] //remove the last delimiter ;
		}
		_ = ordertableAppend(nickname, getCurrentDate(), orderList)
	}
	return true
}

func logPurchases(items []string, units []string, amounts []string) bool {
	for i, _ := range items {
		if items[i] != "" && units[i] != "" && amounts[i] != "" {
			amount, err := strconv.ParseFloat(amounts[i], 64)
			checkError(err, "controller-logPurchases-1")
			_ = purchasetableAppend(getCurrentDate(), items[i], units[i], amount)
		}
	}

	return true
}

func getOrders() {
	//for daily reports
}

func updateOrders() {
	//for daily reports
}

func getPurchases() (ids []int, dates []int, items []string, units []string, amounts []float64) {
	ids, dates, items, units, amounts = purchasetableGetActive()
	return
}

func updatePurchases() {
	//for daily reports
}
