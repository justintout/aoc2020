#!/usr/bin/env bash

die() {
    echo "$*" 1>&2
    exit 1
}

[[ $(date +"%Y-%m-%d %H:%M:%S") > "2020-12-01 00:00:00" ]] || die "It's not December 2020 yet!"

for i in $(seq 1 $(date +"%-d")); do
    curl -sS --cookie "session=$1" "https://adventofcode.com/2020/day/$i/input" > "./input/day$i"
done