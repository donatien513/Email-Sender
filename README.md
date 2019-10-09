# Email Sender
Rest API email sender

[![Go Report Card](https://goreportcard.com/badge/github.com/donatien513/Email-Sender)](https://goreportcard.com/report/github.com/donatien513/Email-Sender)

### Usage

There is no route. Url path is ignored.

- Method `GET` is used for checking Email server's health
- Method `POST` is used for sending Emails

Parameters are send as JSON

- **recipients**: Array of the email addresses of the recipients
- **body**: String containing the body of the email

### Example

```
{
  "recipients": ["someone@example.com"],
  "body": "Hello Someone, This is Me"
}
```
