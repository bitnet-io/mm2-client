#!/bin/bash


echo 'point your browser at

this should be null because the data is merged into port 1717 for market
http://127.0.0.1:13579/api/v2/tickers?expire_at=259200

http://127.0.0.1:1717/market.json

'
#sleep 1s

pkill -f mm2_tools_server_bin
sleep 5
./mm2_tools_server_bin -only_price_service=true &

./server-stock &

while :
do
jq -s '.[0] * .[1]' stocks/crypto.json stocks/stocks.json > stocks/market.json
sleep 60s


echo '


!!!!!!!!!!!!!!!!!!!!!!!
updating the market api
!!!!!!!!!!!!!!!!!!!!!!!

'
done

