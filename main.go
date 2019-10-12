package Handler

import (
  "os"
  "net/http"
  "net/smtp"
  "encoding/json"
)

// Get all env vars
var requestAllowedToken string = os.Getenv("request_allowed_token")
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
  AuthToken := r.Header.Get("Authorization")
  if AuthToken != requestAllowedToken {
    httpFailure(w, http.StatusUnauthorized)
    return
  }
  if r.Body == nil {
    httpFailure(w, http.StatusBadRequest)
    return
  }
  defer r.Body.Close()
  if r.Method == http.MethodGet {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("I am working :)"))
    return
  }

  if r.Method != http.MethodPost {
    httpFailure(w, http.StatusNotImplemented)
    return
  }

  decoder := json.NewDecoder(r.Body)
  var emailSendRequest EmailSendRequest
  decodeErr := decoder.Decode(&emailSendRequest)
  if decodeErr != nil {
    httpFailure(w, http.StatusBadRequest)
    return
  }

  if len(emailSendRequest.Recipients) == 0 || emailSendRequest.Body == "" {
    httpFailure(w, http.StatusBadRequest)
    return
  }

  sendErr := sendEmail(emailSendRequest.Recipients, []byte(emailSendRequest.Body))

  if sendErr != nil {
    httpFailure(w, http.StatusInternalServerError)
  } else {
  	w.WriteHeader(http.StatusOK)
  }
}

// Print error to Http socket
func httpFailure(w http.ResponseWriter, httpStatusCode int) {
  http.Error(w, http.StatusText(httpStatusCode), httpStatusCode)
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
