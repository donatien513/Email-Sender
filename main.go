package Handler

import (
  "os"
  "net/http"
  "net/smtp"
  "encoding/json"
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
  if r.Body == nil {
    // Tell the client that his request payload is missing
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
  }

  if r.Method == http.MethodGet {
    // Tell the client that everything is fine
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("I am working :)"))
    return
  }

  if r.Method != http.MethodPost {
    // Tell the client that his request method is not implemented
    http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
    return
  }

  decoder := json.NewDecoder(r.Body)
  var emailSendRequest EmailSendRequest
  decodeErr := decoder.Decode(&emailSendRequest)
  if decodeErr != nil {
    // Tell the client that his request payload is not ok
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  sendErr := sendEmail(emailSendRequest.Recipients, []byte(emailSendRequest.Body))
  if sendErr != nil {
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
