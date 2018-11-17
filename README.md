# γρ-bot
Simple spell check and grammar expert bot for Telegram via Bing Spell API.

## Status
[![Build Status](https://travis-ci.com/sunloving/gamma-rho-bot.svg?branch=master)](https://travis-ci.com/sunloving/gamma-rho-bot)

## How to install
1. Create a bot in Telegram via BotFather bot. You should get bot token.
2. Register in Bing Spell API and receive API Key.
3. Put token and API key to `.env` file.
4. Start `build/start.sh`.
5. It should work.

## How to use
* As a spell checker, the bot just replies on invalid text messages and sends a revised version of source message.
* On command `/iv learn` the bot returns a message with all forms of *irregular verbs* if operand was one and error if it wasn't.
