# Notify Service

This service provides the ability to send notifications using multiple services:

- Mail
- Telegram
- Slack
- Webhooks
- SMS
- Whatsapp

### Configs

The service is configured using the following environment variables:

| Token            | Description                           |
|------------------|---------------------------------------|
| `AUTH_TOKEN`     | token to make request to api          |
| `MAIL_TOKEN`     | token to access mail service          |
| `TELEGRAM_TOKEN` | token to send messages using telegram |
| `SLACK_TOKEN`    | token for Slack                       |
| `WEBHOOK_TOKEN`  | token to send request to webhook      |
| `SMS_TOKEN`      | token for sending sms                 |
| `WHATSAPP_TOKEN` | token for whatsapp                    |

All tokens are optional, but if you want to use a service you must provide the token.
You need create configs/config.env file and set the environment variables.