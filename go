#!/bin/bash

DATA_DIR=/home/ate/bin/data/setcurr
if [[ ! -e $DATA_DIR/$1 ]]; then
	echo Error: directory not set.
fi

if [[ -e $DATA_DIR/$1 ]]; then
	cd `cat $DATA_DIR/$1`
fi
