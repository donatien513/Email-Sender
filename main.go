package Handler

import (
  "os"
  "log"
  "net/http"
  "net/smtp"
)

var username string = os.Getenv("username")
var password string = os.Getenv("password")
var hostname string = os.Getenv("hostname")
var port string = os.Getenv("port")

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