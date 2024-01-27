#!/bin/bash



while :
do
cat crypto.json | jq . > cryptos.json
cp cryptos.json crypto.json

sed -n '3p' crypto.json > crypto.txt
sed -i s/$/','/ crypto.txt
sed -i s/usdValue/last_price/ crypto.txt

sed -n '5p' stock.json > ford.txt
 sed -i s/','/'",'/  ford.txt
 sed -i s/': '/': "'/  ford.txt


sed -n '9p' stock.json > ibm.txt
 sed -i s/','/'",'/  ibm.txt
 sed -i s/': '/': "'/  ibm.txt

sleep 1s
echo 'slept 5s'
cat ford.txt
cat ibm.txt
cat crypto.txt

file="ford.txt"
ford=$(cat "$file")

awk -v x="$ford" 'NR==17 {$0=x} 1' default.json > def.json
sleep 2s


file2="ibm.txt"
ibm=$(cat "$file2")


awk -v x="$ibm" 'NR==30 {$0=x} 1' def.json > marketss.json
sleep 2s



file3="crypto.txt"
ccc=$(cat "$file3")


awk -v x="$ccc" 'NR==4 {$0=x} 1' marketss.json > mar.json
cp mar.json market.json

sleep 1s



#sed -n '3p' mar.json > m1.txt
#file="m1.txt"
#m1=$(cat "$file")

#awk -v x="$m1" 'NR==4 {$0=x} 1' mar.json > mars.json


#sed -n '17p' mars.json > m2.txt
#file="m2.txt"
#m2=$(cat "$file")

#awk -v x="$m2" 'NR==17 {$0=x} 1' mars.json > marss.json

#cp marss.json market.json

#sed -n '30p' marss.json > m3.txt
#file="m3.txt"
#m3=$(cat "$file")

#awk -v x="$m3" 'NR==30 {$0=x} 1' marss.json > marsss.json

#cp marsss.json market.json








done
