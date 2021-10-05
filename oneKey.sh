#!/bin/sh
FROM=$1
TO=$2

ADDRESS=a0f412d26fa65d49999fa9b5bd1c184423e557c3
AMOUNT=888888888000000000000000000

gsed -i "20016s/Address:.*$/Address: $ADDRESS,/" ./core/genesis/init.go
gsed -i "20017s/Amount: .*$/Amount: $AMOUNT,/" ./core/genesis/init.go 

echo "change" $FROM "=>" $TO
grep -rl $FROM */* | xargs sed -i "" "s/${FROM}/${TO}/g"
grep --color=always -r "$TO" */*
echo "==============================================="
