#!/bin/bash
echo "Post request on orders"
echo "Enter cookies "
read -p "jwt_token : " JWT_TOKEN
read -p "foodopia-session : " FOODOPIA_SESSION

ab -n 10000 -c 1000 \
  -p payload/demo_order.json \
  -T application/json \
  -H "Cookie: jwt_token=$JWT_TOKEN; foodopia-session=$FOODOPIA_SESSION" \
  http://localhost:9005/api/order