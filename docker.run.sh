#!/bin/bash
docker rm -f grammar-bot
docker run -d \
    --name grammar-bot \
    -e "TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN" \
    -e "BING_SPELL_API_KEY=$BING_SPELL_API_KEY" \
    -e "CHATS_IDS_CSV=$CHATS_IDS_CSV" \
    $GRAMMAR_BOT_IMAGE

