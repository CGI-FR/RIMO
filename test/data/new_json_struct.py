import json
from pathlib import Path
import os
# Read the JSONL data from the file
path = Path('test/data/testcase_data.jsonl')
print(os.getcwd())
with open(path, 'r') as f:
    data = [json.loads(line) for line in f]

# Convert the data into the desired structure
result = {}
for row in data:
    for col_name, col_value in row.items():
        if col_name not in result:
            result[col_name] = []
        result[col_name].append(col_value)

# Print the result
print(json.dumps(result))

output = Path("test/data/testcase_newstruct.json")
with open(output, 'w') as f:
    for col, data in result.items():
        f.write(json.dumps({col: data}) + '\n')