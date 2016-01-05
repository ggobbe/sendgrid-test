# sendgrid-test
Send a test email via SMTP through SendGrid

## Usage

```
Usage of sendgrid-activities:
  -apiUser="REQUIRED": Username to connect to the SendGrid API
  -apiKey="all": Password to connect to the SendGrid API
  -fromEmail="REQUIRED": Sender's email address
  -toEmail="REQUIRED": Recipient's email address
```

**Example**
```
sendgrid-test -apiUser="sendgrid-username" -apiKey="sendgrid-password" -fromEmail="test@yourdomain.com" -toEmail="you@yourdomain.com"
```
