#!/bin/bash

T_HOST=$1
T_NS=$2

dig a $T_HOST @$T_NS +norec | grep "$T_HOST" | grep -v ";" | grep "A" | cut -f 6 | head -1 

