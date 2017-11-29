```
curl -v \
  --header "Connection: Upgrade" \
  --header "Upgrade: websocket" \
  --header "Host: 127.0.0.1:8080" \
  --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
   --header "Sec-WebSocket-Version: 13" \
127.0.0.1:8080/chat
```
