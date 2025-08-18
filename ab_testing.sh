# go run cmd/main.go &

read -p "Server Port : " SERVER_PORT
read -p "fool cookie (both jwt and foodopia-session cookie)" COOKIES

echo "Get request benchmark"
ab -C jwt_token=$COOKIES -c 1000 -n 100000 localhost:$SERVER_PORT/api/item  

echo ""
echo "Post request benchmark"
ab -n 10000 -c 1000 \
  -p stress_testing/payload/demo_order.json \
  -T application/json \
  -H "Cookie: $COOKIES" \
  http://localhost:$SERVER_PORT/api/order