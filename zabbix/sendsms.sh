#!/bin/bash

if [ $# -ne 4 ]; then
  echo "Usage: $0 <api_token> <sender_identity> <recipient_numbers> <message_text>"
  exit 1
fi

api_token="$1"
sender_identity="$2"
recipient_numbers="$3"
message_text="$4"

curl -s http://api.sparrowsms.com/v2/sms/ \
    -F token="$api_token" \
    -F from="$sender_identity" \
    -F to="$recipient_numbers" \
    -F text="$message_text"
