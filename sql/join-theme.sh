#!/bin/sh

sed 's/password)/password, dark_mode)/' \
 | awk '
/users/ {
	insert = $0
	getline

	if ($0 ~ /true/) {
		sub(/);$/, ", true);", insert)
	} else {
		sub(/);$/, ", false);", insert)
	}
	print insert
}'
