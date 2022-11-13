#!/usr/bin/env bash

ME=$(basename "$0")
MY_DIR=$(dirname "$0")
DAY=$1
#YEAR=${2:-$(date "+%Y")}
YEAR=${2:-2018}

if [ "${DAY}" == "" ]; then
    echo "usage: ${ME} day [year]" 1>&2
    exit 1
fi

DAY_2D=$(printf '%02d' ${DAY})

PROJECT="aoc-${YEAR}-${DAY_2D}"
PROJECT_DIR="${MY_DIR}/${PROJECT}"

mkdir -p "${PROJECT_DIR}"
"${MY_DIR}/aoc-input.sh" "${DAY}" "${YEAR}" > "${PROJECT_DIR}/input.txt"
cd "${PROJECT_DIR}"
go mod init "rdj/${PROJECT}"
emacsclient --no-wait go.mod

