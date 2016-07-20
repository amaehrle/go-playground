# Vagrant box for golang development with some aws elements

## aws tools installed
* aws console tools
* local redis cache
* local DynamoDb

## left to do:
* start up DynamoDb as a service so it doesn't stop on halt
* make home_dir: "/home/vagrant" a global var and use it in awstools
* clean up

## starting up DynamoDb after a halt

    ./.dynamodbService.sh start

DynamoDB console can be called via

    http://127.0.0.1:8000/shell/

## some useful commands

### redis

    redis-cli ping
    redis-cli
        KEYS *
        GET 'key'
        SET 'key' 'some value'
        DEL 'key'
        TTL 'key'

### DynamoDB

    aws dynamodb list-tables --endpoint-url http://127.0.0.1:8000

    aws dynamodb create-table --table-name dev-table --attribute-definitions AttributeName=id,AttributeType=N --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 --endpoint-url http://127.0.0.1:8000

    aws dynamodb scan --table-name dev-table --endpoint-url http://127.0.0.1:8000

    aws dynamodb get-item --table-name dev-table --key "{ \"id\" : {\"N\":1}}" --endpoint-url http://127.0.0.1:8000

    aws dynamodb delete-item --table-name dev-table --key "{ \"id\" : {\"N\":2}}" --endpoint-url http://127.0.0.1:8000
