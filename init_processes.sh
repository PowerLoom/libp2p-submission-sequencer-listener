#!/bin/bash

echo 'starting pm2...';

pm2 start pm2.config.js

pm2 logs --lines 1000
#docker-compose -f docker-compose.yaml up
#docker run -it proto-snapshot-collector:latest
#go run cmd/main.go