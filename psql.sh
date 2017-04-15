#! /bin/sh

export POSTGRES_CONTAINER=postgres
export POSTGRES_HOST=localhost
export POSTGRES_PASSWORD=postgres
export POSTGRES_PORT=5432
export POSTGRES_DATABASE=liam
export POSTGRES_USER=postgres
export POSTGRES_URL=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DATABASE?sslmode=disable

# start postgres inside docker
docker rm -f $POSTGRES_CONTAINER
docker run --name $POSTGRES_CONTAINER -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -d  -p $POSTGRES_PORT:$POSTGRES_PORT postgres
sleep 5

# create database
docker exec -it postgres psql -U postgres -c "CREATE DATABASE liam"

# migrate db
go get github.com/mattes/migrate
migrate -url $POSTGRES_URL -path ./migrations up

