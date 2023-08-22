# Place PIMO binary in this folder and launch this script to generate data for benchmark.

cd /mixed/
./pimo --empty-input --repeat=100 > 100_input.jsonl
./pimo --empty-input --repeat=1000 > 1000_input.jsonl
./pimo --empty-input --repeat=10000 > 10000_input.jsonl
./pimo --empty-input --repeat=100000 > 100000_input.jsonl
cd ../bool/
./pimo --empty-input --repeat=100 > 100_input.jsonl
./pimo --empty-input --repeat=1000 > 1000_input.jsonl
./pimo --empty-input --repeat=10000 > 10000_input.jsonl
cd ../numeric/
./pimo --empty-input --repeat=100 > 100_input.jsonl
./pimo --empty-input --repeat=1000 > 1000_input.jsonl
./pimo --empty-input --repeat=10000 > 10000_input.jsonl
cd ../string/
./pimo --empty-input --repeat=100 > 100_input.jsonl
./pimo --empty-input --repeat=1000 > 1000_input.jsonl
./pimo --empty-input --repeat=10000 > 10000_input.jsonl
echo "Done"