#!/bin/bash


echo 'point your browser at

this should be null because the data is merged into port 1717 for market
http://127.0.0.1:13579/api/v2/symbols?expire_at=259200

http://127.0.0.1:1717/market.json

'
#sleep 1s

pkill -f mm2_tools_server_bin
sleep 5
./mm2_tools_server_bin -only_price_service=true &

./server-stock &


while :
do
find ./stocks/stocks.json -type f  -exec sed -i s/regularMarketPrice/last_price/g {} +
find ./stocks/crypto.json -type f  -exec sed -i s/lastPrice/last_price/g {} +

cat ./stocks/stocks.json | jq . > ./stocks/stock.json
cp -rf ./stocks/stock.json ./stocks/stocks.json

cat ./stocks/crypto.json | jq . > ./stocks/cryptos.json
cp ./stocks/cryptos.json ./stocks/crypto.json
jq -s '.[0] * .[1]' stocks/crypto.json stocks/stocks.json > stocks/markets.json
#find ./stocks -type f  -exec sed -i s/BIT/BITN/g {} +
#find ./stocks -type f  -exec sed -i s/F/FORD/g {} +
sleep 10s


echo '


!!!!!!!!!!!!!!!!!!!!!!!
updating the market api
!!!!!!!!!!!!!!!!!!!!!!!

'
done

#find ./stocks -type f  -exec sed -i s/ticker/symbol/g {} +
#find ./stocks -type f  -exec sed -i s/ticker/symbol/g {} +
