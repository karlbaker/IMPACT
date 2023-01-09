#!/bin/bash
## Name:    Server Registration Script
## Purpose: Send initial server information payload to the device regisration API.

UUIDVAR=$(dmidecode -t system | grep UUID | awk 'BEGIN {FS=": "}; {print $2}')
MFGVAR=$(dmidecode -t system | grep Manufacturer | awk 'BEGIN {FS=": "}; {print $2}')
MDLVAR=$(dmidecode -t system | grep "Product Name" | awk 'BEGIN {FS=": "}; {print $2}')
SNVAR=$(dmidecode -t system | grep "Serial Number" | awk 'BEGIN {FS=": "}; {print $2}')
IPVAR=$(ip addr | grep 192.168.50 | awk 'BEGIN {FS=" "};{print $2}' | awk 'BEGIN {FS="/"};{print $1}')
MACVAR=$(ip addr | grep -B 1 192.168.50 | head -1 | awk 'BEGIN {FS=" "};{print $2}')

echo $UUIDVAR
echo $MFGVAR
echo $MDLVAR
echo $SNVAR
echo $IPVAR
echo $MACVAR

payload='{"ID": "8", "WallPort": "'"$UUIDVAR"'", "VlanName": "'"$MFGVAR"'", "Change": false}'



curl -X POST -H "Content-Type: application/json" -d "$payload" http://192.168.50.119:9090/ports