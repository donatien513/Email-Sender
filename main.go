package Handler

import (
  "os"
  "net/http"
  "net/smtp"
)

// Get all env vars
var username string = os.Getenv("username")
var password string = os.Getenv("password")
var hostname string = os.Getenv("hostname")
var port string = os.Getenv("port")

// Type of JSON HTTP payload
type EmailSendRequest struct {
  Recipients []string
  Body string
}

// HTTP Entry point
func Handler(w http.ResponseWriter, r *http.Request) {
  // Parse request body
  if req.Body == nil {
    // Tell the client that his request payload is missing
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
  }

  if r.Method != http.MethodGet {
    // Tell the client that everything is fine
    w.WriteHeader(http.StatusOK)
    w.Write("I am working :)")
    return
  }

  if r.Method != http.MethodPost {
    // Tell the client that his request method is not implemented
    http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
    return
  }

  decoder := json.NewDecoder(req.Body)
  var emailSendRequest EmailSendRequest
  err := decoder.Decode(&emailSendRequest)
  if err != nil {
    // Tell the client that his request payload is not ok
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  err := sendEmail(emailSendRequest.Recipients, []byte(emailSendRequest.Body))
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  } else {
  	w.WriteHeader(http.StatusOK)
  }
}

// Main action : Sending email by SMTP
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
