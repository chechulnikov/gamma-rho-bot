# γρ-bot
Simple spell check
Telegram bot. It just reply on invalid text messages and pass a revised version.

## How to use
1. Create a bot in Telegram via BotFather bot. You should get bot token.
2. Register in Bing Spell API and receive API Key.
3. Put token and API key to `.env` file.
4. Exports env vars to shell by `export $(egrep -v '^#' .env | xargs)`.
5. Start `start.sh`.
8. It should work.
