package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s He is %d age old. He have %d "+
		"money on count and %f level of happiness", u.name, u.age, u.money, u.happiness)

}
func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
