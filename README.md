# trader-slack-notifier-go

a simple service to send messages to slack.

it accepts requests from a REST API and from a Kafka Stream

REST API
```
POST http://localhost:8080/message
```



request:
```
{
    "token": "<token>",
    "sender": "<user_id>",
    "channel": "<channel_id>",
    "receiver": "<user_id>",
    "ephemeral": false,
    "type": "<Transaction|Execution|Text>",
    "message": <Transaction|Execution|Text JSON>
}
```



