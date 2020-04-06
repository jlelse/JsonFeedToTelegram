# JSON Feed to Telegram Bot

This is a simple Telegram bot, written in Go, that is able to check a JSON Feed for updates and sends the latest article to a Telegram channel.

## How to use

You should configure this bot to run behind a reverse proxy. It is listening to port 8080 by default. It is recommended to use Docker to run this bot. That let's your remap the port, easily set environment variables etc.

## How to trigger posting to Telegram

A recheck of the JSON Feed is triggered when the path `/hook` is called using a HTTP POST request.

## Configuration

Some environment variables are required:

`LAST_ARTICLE_FILE`: file path to the file where the URL of the last sent article should be saved

`FEED`: URL to the JSON Feed

`BOT_TOKEN`: The token for the Telegram bot

`CHANNEL`: The channel or user ID to which notifications should be send

## License

This project is licensed under the MIT license, so you can do basically everything with it, but nevertheless, please contribute your improvements to make it better for everyone. See the LICENSE file.