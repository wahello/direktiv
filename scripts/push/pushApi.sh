#!/bin/bash

make docker-api && docker tag direktiv-api localhost:5000/api

docker push localhost:5000/api
