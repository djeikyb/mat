Mats your photo if needed. Assumes you want a 4x6 print.

This script is equivalent for landscape oriented photos (uses imagemagick):

	#!/bin/sh

	input="$1"
	file="$(basename $input)"
	name="${file%.*}"

	calc_dim() {
		echo "$1"
		magick identify -format '%[w] %[fx:4*w/6]' "$1"
	}

	read width height <<< $(calc_dim "$input")

	magick convert "$input" \
		-background white \
		-gravity center \
		-extent "${width}x${height}" \
		"${name}_mat_4x6.jpg"
