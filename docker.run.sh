#!/bin/bash
docker rm -f grammar-bot
docker run -d \
    --name grammar-bot \
    -e "TELEGRAM_BOT_TOKEN=<put_your_bot_token_here>" \
    -e "BING_SPELL_API_KEY=<put_your_api_key_here>" \
    grammar-bot

