#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd "${SCRIPT_DIR}/mixed/"
if [ ! -f 100_input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100_input.jsonl
fi
if [ ! -f 1000_input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000_input.jsonl
fi
if [ ! -f 10000_input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000_input.jsonl
fi
if [ ! -f 100000_input.jsonl ]; then
    pimo --empty-input --repeat=100000 > 100000_input.jsonl
fi

echo "data for mixed : OK"
cd "${SCRIPT_DIR}/bool/"
if [ ! -f 100_input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100_input.jsonl
fi
if [ ! -f 1000_input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000_input.jsonl
fi
if [ ! -f 10000_input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000_input.jsonl
fi
echo "data for mixed : OK"

cd "${SCRIPT_DIR}/numeric/"
if [ ! -f 100_input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100_input.jsonl
fi
if [ ! -f 1000_input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000_input.jsonl
fi
if [ ! -f 10000_input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000_input.jsonl
fi
echo "data for numeric : OK"

cd "${SCRIPT_DIR}/text/"
if [ ! -f 100_input.jsonl ]; then
    pimo --empty-input --repeat=100 > 100_input.jsonl
fi
if [ ! -f 1000_input.jsonl ]; then
    pimo --empty-input --repeat=1000 > 1000_input.jsonl
fi
if [ ! -f 10000_input.jsonl ]; then
    pimo --empty-input --repeat=10000 > 10000_input.jsonl
fi
echo "data generated for text : OK"

echo "Done generated benchmark dataset"
