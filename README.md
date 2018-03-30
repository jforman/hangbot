# Hangbot

Hangbot is a bot in the style of many other IRC/Slack bots, but for use on Google's new [Hangouts Chat](https://gsuite.google.com/products/chat/) product.

## Features

*  Receive messages via Pub/Sub topic subscription.
*  Asynchronous responses via HTTP API.

## Requirements

From the [Hangouts Chat](https://developers.google.com/hangouts/chat/) documentation.

*  ServiceAccount for the bot. To be used for receiving messages via Pub/Sub.
*  Pub/Sub topic created to be used for your Bot to receive messages.

## Screenshot of typical interaction

![Screenshot](https://github.com/jforman/homebot/blob/master/hangouts-chat-screenshot.png)

## Execution

```bash
$ ./hangbot.go -credentialsFile account.json \
-project fooProject \
-topic hangouts-chat.homebot \
-subscription HomeSub
```

## Reference Material

*  https://godoc.org/google.golang.org/api/chat/v1
*  https://github.com/google/google-api-go-client/tree/master/chat/v1
*  Basically everything under https://developers.google.com/hangouts/chat/concepts
*  Pub/Sub client library quick start: https://developers.google.com/hangouts/chat/how-tos/bots-publish?authuser=0
*  https://github.com/google/google-api-go-client/blob/master/GettingStarted.md

## Future Development Ideas

* Use the [cards](https://developers.google.com/hangouts/chat/concepts/cards) concept to make more rich interactions with the Bot.
* Interactions with [Nestmon](https://github.com/jforman/nestmon), some Go code I wrote to stream Nest thermostat updates from my house. First just a way to see the current cooling/heating status of my home. Eventually maybe some sort of write capabilities to modify temperatures inside my home.
