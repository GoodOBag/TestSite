package main

func getUnits() []string {
	units := []string{"lb", "ea"}
	return units
}

func updateUnits(units []string) bool {
	_ = units
	return true
}

func getItems() (items []string, units []string, unitPrices []float64, notes []string) {
	items = []string{"Apple", "Banana", "Orange", "Cabbage"}
	units = []string{"ea", "ea", "lb", "ea"}
	unitPrices = []float64{1.99, 0.79, 1.49, 2.09}
	notes = []string{"ww", "aaa", "g", "xxx"}
	return items, units, unitPrices, notes
}

func updateItems(items []string, units []string, unitPrices []string, notes []string) bool { //unitprice read from map as string
	_, _, _, _ = items, units, unitPrices, notes
	return true
}

func getBldgs() (bldgs []string, addrs []string, zips []string, notes []string) {
	bldgs = []string{"St. Mary", "Alice St.", "Harrison St.", "Hilltop"}
	addrs = []string{"130 A st.", "200 Alice St.", "30 Harrison St.", "150 G st."}
	zips = []string{"12300", "45006", "70089", "05008"}
	notes = []string{"", "aaa", "", "xxx"}
	return bldgs, addrs, zips, notes
}

func updateBldgs(bldgs []string, addrs []string, zips []string, notes []string) bool {
	_, _, _, _ = bldgs, addrs, zips, notes
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
