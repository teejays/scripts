#!/bin/bash

home=$(pwd)
cnt=1
for d in extract/*/; do
	echo "Home is $home"
	echo "Processing director $d..."
	echo "Count is $cnt"
	cd "$d"
	echo "Current directory is $(pwd)"
	for f in *.jpg; do
		echo " - file $f..."
		echo " - changing name from $f to ${f%.jpg}_$cnt.jpg"
		mv -- "$f" "${f%.jpg}_$cnt.jpg"
	done
	cd "$home"
	cnt=$((cnt + 1))
done
