#!/bin/bash

username="$1"
ip="$2"

ssh -i ~/.ssh/PE_UAT.pem root@korea << COMMAND
    cd /data/sports
    ./add_ip.sh $username $ip
COMMAND
