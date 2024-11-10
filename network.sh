#!/bin/bash

docker build -t gone-broadcast .


gone-cli node -- docker run -d --rm --network gone_net --name gone-1 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-2 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-3 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-4 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-5 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-6 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-7 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-8 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-9 gone-broadcast
gone-cli node -- docker run -d --rm --network gone_net --name gone-10 gone-broadcast


gone-cli bridge bridge1

gone-cli bridge bridge2

gone-cli bridge bridge3
gone-cli router router1

gone-cli connect node gone-1 bridge1
gone-cli connect node gone-2 bridge1
gone-cli connect node gone-3 bridge1
gone-cli connect node gone-4 bridge1
gone-cli connect node gone-5 bridge1

gone-cli connect node gone-6 bridge2
gone-cli connect node gone-7 bridge2
gone-cli connect node gone-8 bridge2

gone-cli connect node gone-9 bridge3
gone-cli connect node gone-10 bridge3

gone-cli connect bridge bridge1 router1
gone-cli connect bridge bridge2 router1
gone-cli connect bridge bridge3 router1

gone-cli unpause -a
