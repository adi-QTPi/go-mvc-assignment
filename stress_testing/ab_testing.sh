# docker compose -f ../docker/docker-compose.yml up&

# sleep 10s

read -p "Server Port : " SERVER_PORT
read -p "Customer JWT jwt_token=" JWT_TOKEN

ab -C jwt_token=$JWT_TOKEN -c 1000 -n 100000 localhost:$SERVER_PORT/api/item  