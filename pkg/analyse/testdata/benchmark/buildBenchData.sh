#!/usr/bin/env bash

# Place PIMO binary in this folder and launch this script to generate data for benchmark.

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


cd "${SCRIPT_DIR}/mixed/"
pimo --empty-input --repeat=100 > 100_input.jsonl
pimo --empty-input --repeat=1000 > 1000_input.jsonl
pimo --empty-input --repeat=10000 > 10000_input.jsonl
pimo --empty-input --repeat=100000 > 100000_input.jsonl
cd "${SCRIPT_DIR}/bool/"
pimo --empty-input --repeat=100 > 100_input.jsonl
pimo --empty-input --repeat=1000 > 1000_input.jsonl
pimo --empty-input --repeat=10000 > 10000_input.jsonl
cd "${SCRIPT_DIR}/numeric/"
pimo --empty-input --repeat=100 > 100_input.jsonl
pimo --empty-input --repeat=1000 > 1000_input.jsonl
pimo --empty-input --repeat=10000 > 10000_input.jsonl
cd "${SCRIPT_DIR}/text/"
pimo --empty-input --repeat=100 > 100_input.jsonl
pimo --empty-input --repeat=1000 > 1000_input.jsonl
pimo --empty-input --repeat=10000 > 10000_input.jsonl
echo "Done"
