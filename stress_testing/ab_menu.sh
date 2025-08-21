#!/bin/bash

echo "Enter cookies "
read -p "jwt_token : " JWT_TOKEN

echo $JWT_TOKEN

ab -C jwt_token=$JWT_TOKEN -n 100000 -c 1000 localhost:9005/api/item