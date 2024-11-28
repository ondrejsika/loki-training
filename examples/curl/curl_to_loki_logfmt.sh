#!/bin/sh

curl -X POST http://127.0.0.1:3100/loki/api/v1/push \
  -H "Content-Type: application/json" \
  -d '{
    "streams": [
      {
        "stream": {
          "job": "curl"
        },
        "values": [
          ["'$(date +%s000000000)'", "cluster=us host=aaa-1 app=foo level=info message=hello"],
          ["'$(date +%s000000000)'", "cluster=eu host=bbb-1 app=bar level=info message=hello"],
          ["'$(date +%s000000000)'", "cluster=us host=aaa-2 app=foo level=warn message=hello"],
          ["'$(date +%s000000000)'", "cluster=cn host=ccc-1 app=baz level=info message=\"hello world from china\""]
        ]
      }
    ]
  }'
