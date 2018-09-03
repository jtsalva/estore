import requests
import json

HOST = "http://localhost:8080/items/"

mock_items = json.loads(open("./mock_items.json", "r").read())

for item in mock_items:  
    r = requests.post(HOST, json=item)

    if r.status_code != 201:
        print(r.status_code)

print(r.status_code)

# print(r.status_code)

# if r.status_code == 200:
#     print(r.json())