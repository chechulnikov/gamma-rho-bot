#!/bin/bash
export $(egrep -v '^#' ../.env | xargs)
docker rm -f gamma-rho-bot
docker run -d \
    --name gamma-rho-bot \
    -e "TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN" \
    -e "BING_SPELL_API_KEY=$BING_SPELL_API_KEY" \
    -e "CHATS_IDS_CSV=$CHATS_IDS_CSV" \
    $GRAMMAR_BOT_IMAGE

