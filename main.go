/*

1. Monday 7am message on Twilio, calculates weekly spending budget + summarises spend in previous week
2. Thursday 7am message, update of spend during week so far, + message from message queue

3. Monthly message calculating forecast of bills and income in following month

4. Signup process

5. Reconnection process


##

1. Create short url linking page and customer
2. When that page is accessed present it

Database of customers
{
	id: 1
	name: Nick
	mobile: 0787665757
	bank: barclays
}
{
	id: 2
	name: bubu
	mobile: 0787665756
	bank: Stanchart
}


Database of web pages

BankPage: <p>Hi [name]</p><p>Click <a href>here</a> to go to bank</p>

Page 2: <p>Hi [name]</p><p>Click <a href>here</a> to go to your Utility page</p>

func renderBankPage(customer Customer){}
func renderUtilityPage(customer Customer){}

Types of pages
BankPage: Page 1
UtilityPage: Page 2

if page = BankPage {
	renderBankPage(customer)
}

Database of short urls
Page : typeofPage, CustomerId
abc12 > Page
abc13 > Page

> abc12

https://spendable.uk/abc12 > web page '123' for customer 1001
https://spendable.uk/abc13 > web page 2 for Bubu

web page 123

<p>Hi [name]</p>


//Models

Customer
	- Id
	- Name
	- PhoneNumber
	- Bank

PageTypes:
- Bank
- Utility
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	BankPage    = 0
	UtilityPage = 1
)

type Page struct {
	Type       int
	CustomerId string
}

var pageMap = map[string]Page{
	"abc12": Page{BankPage, "1"},
	"abc13": Page{UtilityPage, "2"},
}

func JsonResponse(w http.ResponseWriter, p interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Println(err)
	}
}
func ErrorResponse(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/members", CreateMember).Methods("POST")
	r.HandleFunc("/members/{memberId}", GetMember).Methods("GET")
	r.HandleFunc("/members", GetMembers).Methods("GET")
	r.HandleFunc("/{shortCode}", HandleShortCode).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func createPageContent(p Page) string {

	c := MemberMap[p.CustomerId]

	if p.Type == BankPage {
		return bankPage(c.Name)
	}

	if p.Type == UtilityPage {
		return utilityPage(c.Name)
	}

	return "<p>Page Not Found</p>"
}

func bankPage(name string) string {
	a := fmt.Sprintf("<p>Hi %s</p><p>Click <a href>here</a> to go to bank</p>", name)
	return a
}

func utilityPage(name string) string {
	a := fmt.Sprintf("<p>Hi %s</p><p>Click <a href>here</a> to go to your Utility page</p>", name)
	return a
}

func HandleShortCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]
	if p, ok := pageMap[shortCode]; !ok {
		fmt.Fprintf(w, "this page doesnt exist\n")
	} else {
		fmt.Fprintf(w, createPageContent(p))
	}
}
