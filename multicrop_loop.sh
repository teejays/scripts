#!/bin/bash

dir=$(pwd)
cnt=1
for f in raw/*.jpg; do
	echo "Processing file $f"
	$dir/multicrop "$f" cropped/picture_$cnt.jpg
	cnt=$((cnt+1))
done
