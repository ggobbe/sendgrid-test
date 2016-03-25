# sendgrid-test
Send a test email via SMTP through SendGrid

## Usage

```
Usage of sendgrid-test:
  -apiUser string
        Username to connect to the SendGrid API (REQUIRED)
  -apiKey string
        Password to connect to the SendGrid API (REQUIRED)
  -fromEmail string
        Sender of the test email (REQUIRED)
  -toEmail string
        Recipient of the test email (REQUIRED)
  -smtpPort string
        SendGrid SMTP port (default "587")
```

**Example**
```
sendgrid-test -apiUser="sendgrid-username" -apiKey="sendgrid-password" -fromEmail="test@yourdomain.com" -toEmail="you@yourdomain.com"
```
