#!/bin/bash

# Check if docker compose V2 is available
if docker compose version &> /dev/null; then
    docker compose build
else
    docker-compose build
fi