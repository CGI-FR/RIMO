#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd "${SCRIPT_DIR}/mixed/"
mkdir 100 1000 10000 100000
if [ ! -f 100/input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100/input.jsonl
fi
if [ ! -f 1000/input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000/input.jsonl
fi
if [ ! -f 10000/input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000/input.jsonl
fi
if [ ! -f 100000/input.jsonl ]; then
    pimo --empty-input --repeat=100000 > 100000/input.jsonl
fi

echo "data for mixed : OK"
cd "${SCRIPT_DIR}/bool/"
mkdir 100 1000 10000 100000
if [ ! -f 100/input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100/input.jsonl
fi
if [ ! -f 1000/input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000/input.jsonl
fi
if [ ! -f 10000/input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000/input.jsonl
fi
echo "data for mixed : OK"

cd "${SCRIPT_DIR}/numeric/"
mkdir 100 1000 10000 100000
if [ ! -f 100/input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100/input.jsonl
fi
if [ ! -f 1000/input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000/input.jsonl
fi
if [ ! -f 10000/input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000/input.jsonl
fi
echo "data for numeric : OK"

cd "${SCRIPT_DIR}/text/"
mkdir 100 1000 10000 100000
if [ ! -f 100/input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100/input.jsonl
fi
if [ ! -f 1000/input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000/input.jsonl
fi
if [ ! -f 10000/input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000/input.jsonl
fi
echo "data generated for text : OK"

echo "Done generated benchmark dataset"
