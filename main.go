package main

import (
	"log"
	"net/http"
	"strings"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	if mux := subdomains[domainParts[0]]; mux != nil {
		mux.ServeHTTP(w, r)
	} else { // Handle 404
		http.Error(w, "Not found", 404)
	}
}

func main() {
	//subdomain "manage"
	manageMux := http.NewServeMux()
	manageMux.HandleFunc("/login", login)
	manageMux.HandleFunc("/", manage)
	manageMux.HandleFunc("/OrderInfo", orderinfo)                     //Log Orders
	manageMux.HandleFunc("/PurchaseInfo", purchaseinfo)               //Log Purchases
	manageMux.HandleFunc("/UnitsList", unitslist)                     //lb, ea, etc.
	manageMux.HandleFunc("/ItemList", itemlist)                       //List of items to serve
	manageMux.HandleFunc("/CustomerList", customerlist)               //List of customers
	manageMux.HandleFunc("/BuildingList", buildinglist)               //Serving buildings
	manageMux.HandleFunc("/DailySummary", dailysummary)               //Daily Summary selection page
	manageMux.HandleFunc("/DailySummaryRecords", dailysummaryrecords) //Daily Summary summary page
	//manageMux.HandleFunc("/DailySummaryPrint", dailysummaryprint)     //Daily Summary printout page

	//domain (for future)

	//default settings
	subdomains := make(Subdomains)
	subdomains["manage"] = manageMux

	err := http.ListenAndServe(":8080", subdomains)
	checkError(err, "main-main")
}

func checkError(err error, loc string) {
	//format of loc:  scriptName-functionName-anyOtherComment
	//e.g. loc for error in current function = main-checkError
	if err != nil {
		log.Fatal(loc, ":", err)
	}
}
