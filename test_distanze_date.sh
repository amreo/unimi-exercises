#!/bin/sh
echo "Dovrebbe restituire 79987"
echo -e "1-1-1800 31-12-2018\n" | go run distanza_date.go 
echo "Dovrebbe restituire 79259"
echo -e "31-12-1800 1-1-2018\n" | go run distanza_date.go 
echo "Dovrebbe restituire 79672"
echo -e "18-5-1800 6-7-2018\n" | go run distanza_date.go 
