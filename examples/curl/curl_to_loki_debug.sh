#!/bin/sh

curl -X POST http://127.0.0.1:3100/loki/api/v1/push \
  -H "Content-Type: application/json" \
  -d '{
    "streams": [
      {
        "stream": {
          "job": "curl",
          "level": "debug"
        },
        "values": [
          ["'$(date +%s000000000)'", "This is a debug message sent to Loki using curl with word error"]
        ]
      }
    ]
  }'
