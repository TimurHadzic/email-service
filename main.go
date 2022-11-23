package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var router *http.ServeMux
var db *sql.DB
var err error

type Email struct {
	FromEmail  string `json:FromEmail`
	ToEmail    string `json:ToEmail`
	CCEmail    string `json:CCEmail`
	Subject    string `json:Subject`
	Importance string `json:Importance`
	Content    string `json:Content`
}

// func SendMail(w http.ResponseWriter, r *http.Request) {
// 	var NewEmail Email
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.Unmarshal(reqBody, &NewEmail)

// 	fmt.Println(NewEmail.FromEmail)
// 	db.Exec("INSERT INTO email (from_email, to_email, cc_email, subject, importance, content) VALUES (?, ?, ?, ?, ?, ?)", NewEmail.FromEmail, NewEmail.ToEmail, NewEmail.CCEmail, NewEmail.Subject, NewEmail.Importance, NewEmail.Content)
// }

func SendMail(w http.ResponseWriter, r *http.Request) {
	var NewMail Email
	NewMail.FromEmail = r.FormValue("from")
	NewMail.ToEmail = r.FormValue("to")
	NewMail.CCEmail = r.FormValue("cc")
	NewMail.Subject = r.FormValue("subject")
	NewMail.Importance = r.FormValue("importance")
	NewMail.Content = r.FormValue("content")
	db.Exec("INSERT INTO email (from_email, to_email, cc_email, subject, importance, content) VALUES (?, ?, ?, ?, ?, ?)", NewMail.FromEmail, NewMail.ToEmail, NewMail.CCEmail, NewMail.Subject, NewMail.Importance, NewMail.Content)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	db, _ = sql.Open("mysql", "root:a84628462@tcp(localhost:3306)/email_service")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router = http.NewServeMux()

	router.HandleFunc("/", Welcome)
	router.HandleFunc("/send-email", SendMail)
	log.Fatal(http.ListenAndServe(":8080", router))

}
