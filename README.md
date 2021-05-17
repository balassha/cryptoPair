# cryptoPair

## Goal
Create a service that collect data from cryptocompare.com using its API and stores it
in a database (MySQL or PostgreSQL)
Example API request: GET
https://min-api.cryptocompare.com/data/pricemultifull?fsyms=BTC&tsyms=USD,EUR

## Description

* This is a Scalable HTTP based Application that accepts a Currency pair and gives the predefined set of values about the currency pairs.
* It has a Load balancer in the front which can be configured easily to forward data to one/more CryptoPair servers.
* The Load balancer can be easily configured as a Reverse Proxy with SSL termination using self signed/CA signed certificates.
* Both Load balancer and CryptoPair server are containerized which make it easy to deploy it to the cloud.
* Load balancer is defined to listen on Port 8111.
* This application hosts a HTTP server that listens on Port 8011 and has two endpoints defined.

First API accepts a GET request which needs to have the Currency Pair as Query params. It sends a http request to the remote API and 
constructs the response. It also schedules an update to the Database with the latest values. If the remote API is down, the API checks
if the data is already available in the DB and send the response with data from DB.

Second API fetches the data directly from the the Database.

## Pre-requisites
1. Docker
2. Git
3. mysql setup

## Getting started

DB Configuration - DatabaseParams.csv & DatabaseConfig.csv inside /Config. Update the files with DB Configurations and params.

To start this application, we need to start CryptoPair server and Load Balancer use the following commands.
Directory - CryptoPair/api
Build : sudo docker build -t crypto-currencies -f Dockerfile .
Run   : sudo docker run --rm --network=host --name crypto -d crypto-currencies

Directory - CryptoPair/nginx
Build : sudo docker build -t nginx -f Dockerfile .
Run   : sudo docker run --rm --network=host --name nginx -d nginx

## API Reference
http://<host-address>:8111/v1/service/price?fsyms=BTC&tsyms=EUR
http://<host-address>:8111/v1/service/db?fsyms=ETH&tsyms=USD

## Improvements (Couldn't do these due to the time limit)
1. Better Test coverage
2. Create a Docker compose file
3. Better definition of constants
4. Accept multiple currencies in a single request
5. Websocket implementation
