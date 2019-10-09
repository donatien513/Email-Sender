# Email Sender
Rest API email sender

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
