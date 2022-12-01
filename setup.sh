#!/usr/bin/env bash

ME=$(basename "$0")
MY_DIR=$(dirname "$0")
DAY=$1
#YEAR=${2:-$(date "+%Y")}
YEAR=${2:-2022}

if [ "${DAY}" == "" ]; then
    echo "usage: ${ME} day [year]" 1>&2
    exit 1
fi

DAY_2D=$(printf '%02d' ${DAY})

PROJECT="aoc_${YEAR}_${DAY_2D}"
PROJECT_DIR="${MY_DIR}/${PROJECT}"

mkdir -p "${PROJECT_DIR}"
"${MY_DIR}/aoc-input.sh" "${DAY}" "${YEAR}" > "${PROJECT_DIR}/input.txt"

for f in "${MY_DIR}/template/"*.go; do
    dst=$(basename $f)
    dst="${PROJECT_DIR}/${dst//aoc_yyyy_mm/${PROJECT}}"
    perl -p -e "s{aoc_yyyy_mm}{${PROJECT}}" < "$f" > "$dst"
done
cd "${PROJECT_DIR}"
go mod init "rdj/${PROJECT}"
emacsclient --no-wait *.go

