#!/bin/bash

echo "----- Pre-Requisites -----"

if ! command -v go >/dev/null 2>&1
then
    echo "Golang is not installed on your device."
    echo "Golang installation is required to run this Application."
    echo "Install on macOS : brew install go"
    echo "Install on Ubuntu : sudo apt-get update && sudo apt-get install golang-go"
    exit 1
else
    echo "Golang is Installed !"
fi

if ! command -v mysql >/dev/null 2>&1
then
    echo "MySQL is not installed on your device."
    echo "MySQL installation is required to run this Application."
    echo "Install on macOS : brew install mysql"
    echo "Install on Ubuntu : sudo apt-get update && sudo apt-get install mysql-server"
    exit 1
else
    echo "MySQL is Installed !"
fi


echo ""
echo "----- Environment Variables -----"
echo "Please input valid values. Press Enter to use the default value."
echo ""
echo "" > .env

read -p "Enter SERVER_PORT [default: 9005]: " SERVER_PORT
SERVER_PORT=${SERVER_PORT:-9005}

read -p "Enter MYSQL_HOST [default: localhost]: " MYSQL_HOST
MYSQL_HOST=${MYSQL_HOST:-"localhost"}

read -p "Enter MYSQL_USER [default: root]: " MYSQL_USER
MYSQL_USER=${MYSQL_USER:-"root"}

while [[ "$MYSQL_PASSWORD" == "" ]]; do
    read -p "Enter MySQL password (mandatory): " MYSQL_PASSWORD
    if [[ "$MYSQL_PASSWORD" == "" ]]; then
        echo "Password cannot be empty. Please try again."
    fi
done

read -p "Enter MYSQL_PORT [default: 3306]: " MYSQL_PORT
MYSQL_PORT=${MYSQL_PORT:-3306}

read -p "Enter MYSQL_DATABASE [default: karma_mvc_foodopia]: " MYSQL_DATABASE
MYSQL_DATABASE=${MYSQL_DATABASE:-"karma_mvc_foodopia"}

while true; do
    read -p "Enter JWT Secret (mandatory, min 32 characters): " JWT_SECRET
    if [[ ${#JWT_SECRET} -ge 32 ]]; then
        break
    else
        echo "JWT Secret must be at least 32 characters long. Please try again."
    fi
done

while true; do
    read -p "Enter Sessions Secret (mandatory, min 32 characters): " SESSIONS_SECRET
    if [[ ${#SESSIONS_SECRET} -ge 32 ]]; then
        break
    else
        echo "Sessions Secret must be at least 32 characters long. Please try again."
    fi
done


echo "!!!--- IMPORTANT : Pay Attention ---!!!"
echo "Enter the Admin Account details with UTMOST CARE"
echo ""
sleep 2s
read -p "Enter ADMIN_USERNAME [default: haldi]: " ADMIN_USERNAME
ADMIN_USERNAME=${ADMIN_USERNAME:-"haldi"}

read -p "Enter ADMIN_NAME [default: Mr. Admin]: " ADMIN_NAME
ADMIN_NAME=${ADMIN_NAME:-"Mr. Admin"}

read -p "Enter ADMIN_PASSWORD [default: Admin@123]: " ADMIN_PASSWORD
ADMIN_PASSWORD=${ADMIN_PASSWORD:-"Admin@123"}

echo "SERVER_PORT=$SERVER_PORT" >> .env
echo "MYSQL_HOST=\"$MYSQL_HOST\"" >> .env
echo "MYSQL_USER=\"$MYSQL_USER\"" >> .env
echo "MYSQL_PASSWORD=\"$MYSQL_PASSWORD\"" >> .env
echo "MYSQL_PORT=$MYSQL_PORT" >> .env
echo "MYSQL_DATABASE=\"$MYSQL_DATABASE\"" >> .env
echo "JWT_SECRET=\"$JWT_SECRET\"" >> .env
echo "SESSIONS_SECRET=\"$SESSIONS_SECRET\"" >> .env
echo "ADMIN_USERNAME=\"$ADMIN_USERNAME\"" >> .env
echo "ADMIN_NAME=\"$ADMIN_NAME\"" >> .env
echo "ADMIN_PASSWORD=\"$ADMIN_PASSWORD\"" >> .env

echo ""
echo "----- .env file created Successfully !. -----"

echo ""
echo "----- Database Initialisations -----"
mysql -u $MYSQL_USER -p$MYSQL_PASSWORD -h $MYSQL_HOST -P $MYSQL_PORT -e "CREATE DATABASE IF NOT EXISTS $MYSQL_DATABASE;"
migrate -path database/schema_migrate -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DATABASE" up
echo "Schema Migrations Applied !"
echo ""
sleep 0.5s
echo "We also provide with Dummy Data to get started with the app... "
read -p "Enter Y to accept the Gift !" response
if [[ $response == "y" ]]; then
    echo "Seeding dummy data into your database."
    mysql -h $MYSQL_HOST -P $MYSQL_PORT -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE < dump.sql
else
    echo "Sure! No worries — all the best on your journey!"
fi

sleep 2s
echo "----- Starting the application -----"
go mod tidy
go run cmd/main.go
