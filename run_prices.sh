#!/bin/bash


echo 'point your browser at


http://127.0.0.1:13579/api/v2/tickers?expire_at=259200

'
sleep 1s

pkill -f mm2_tools_server_bin
sleep 5
./mm2_tools_server_bin -only_price_service=true
