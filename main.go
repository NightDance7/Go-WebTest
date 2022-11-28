package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Reasons               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s He is %d age old. He have %d "+
		"money on count and %f level of happiness", u.Name, u.Age, u.Money, u.Happiness)

}
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 27, -50, 4.2, 0.8, []string{"Есть профессиональное обучение", "Опыт работы с C,C++,Java", "Люблю получать новые знания", "Мало опыта по вебу, но не дурачок:3)"}}
	//	fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
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
