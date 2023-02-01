#!/bin/sh

ps aux | grep -e 'urlapi'
ps aux | grep -e 'urldisp'
ps aux | grep -e 'urlapi' | awk '{print $1}' | xargs kill -9
ps aux | grep -e 'urldisp' | awk '{print $1}' | xargs kill -9

./urlapi > logapi.log 2>&1 &
./urldisp > logdisp.log 2>&1 &

tail -F logdisp.log

# mysql -u checker -p checker -h localhost work