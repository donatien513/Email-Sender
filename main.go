package Handler

import (
  "os"
  "log"
  "net/smtp"
)

var username string = os.Getenv("USERNAME")
var password string = os.Getenv("PASSWORD")
var hostname string = os.Getenv("HOSTNAME")
var port string = os.Getenv("PORT")

func Handler(w http.ResponseWriter, r *http.Request) {
  SendEmail([]string{"donatiennambinintsoa@gmail.com"}, []byte("HEY THERE"))
}

func SendEmail(recipients []string, body []byte) {
  auth := smtp.PlainAuth(
    "",
    username,
    password,
    hostname,
  )
  
  err := smtp.SendMail(
    hostname + ":" + port,
    auth,
    username,
    recipients,
    body,
  )
  if err != nil {
    log.Fatal(err)
  }
}