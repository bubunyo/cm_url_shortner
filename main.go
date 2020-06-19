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
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID          int
	Name        string
	PhoneNumber string
	Bank        string
}

var customerMap = map[int]Customer{
	1: Customer{1, "Nick", "0787665757", "Barclays"},
	2: Customer{2, "Bubu", "244792164", "Stanchart"},
}

const (
	BankPage    = 0
	UtilityPage = 1
)

type Page struct {
	Type       int
	CustomerId int
}

var pageMap = map[string]Page{
	"abc12": Page{BankPage, 1},
	"abc13": Page{UtilityPage, 2},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{shortCode}", HandleShortCode)

	http.ListenAndServe(":8080", r)
}

func createPageContent(p Page) string {

	c := customerMap[p.CustomerId]

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
