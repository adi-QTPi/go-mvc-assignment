source ../.env
SERVER_PORT=$SERVER_PORT
echo $SERVER_PORT
filename="cookies.txt"

echo "Storing cookies in a file"

read -p "UserName : " $username
read -p "Password : " $password

curl -c $filename -X POST "http://localhost:$SERVER_PORT/account/login" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "user_name=$username&password=$password"