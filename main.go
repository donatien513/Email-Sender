package Handler

import (
  "os"
  "net/http"
  "net/smtp"
)

var username string = os.Getenv("username")
var password string = os.Getenv("password")
var hostname string = os.Getenv("hostname")
var port string = os.Getenv("port")

func Handler(w http.ResponseWriter, r *http.Request) {
  err := sendEmail([]string{"donatiennambinintsoa@gmail.com"}, []byte("HEY THERE"))
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  } else {
  	w.WriteHeader(http.StatusOK)
  }
}

func sendEmail(recipients []string, body []byte) error {
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
  return err
}
