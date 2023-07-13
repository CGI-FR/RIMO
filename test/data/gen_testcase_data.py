from faker import Faker
import json
from datetime import datetime


""" Generate fake data for testing purposes in 'analyse.py' """
fake = Faker()

data = []
for i in range(10):
    data.append({
        "address": ' '.join(fake.address().splitlines()),
        "age": fake.random_int(min=18, max=100, step=1),
        "date": fake.date_this_century().strftime('%Y-%m-%d'),
        "phone": fake.phone_number()
    })


print(data[0]['address'])
# Write to file in JSON line format

# Do not uncomment: mean, median, min, max, etc. will need to be recalculated
# with open('data/testcase_data.jsonl', 'w') as outfile:
#     for entry in data:
#         json.dump(entry, outfile)
#         outfile.write('\n')