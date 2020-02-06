#!/bin/bash

# echo usage ./run.sh query

if ! [ -f config.sh ]; then
	echo "please create a config.sh that sets ORG, BUCKET and TOKEN"
	exit 1;
fi

unset ORG BUCKET TOKEN
. config.sh

IFQL=`mktemp tmp.XXXXX`
FLUX=`mktemp tmp.XXXXX`

echo "$1;" > $IFQL

if ! ./ifql2flux $BUCKET < $IFQL >> $FLUX; then
	rm $IFQL $FLUX
	exit 1;
fi

echo sending query:
cat $FLUX

curl \
	-H "Authorization: Token $TOKEN" \
	-H "Content-Type: application/vnd.flux" \
	"http://localhost:9999/api/v2/query?org=$ORG" \
	--data-binary "@$FLUX"
echo
rm $IFQL $FLUX


