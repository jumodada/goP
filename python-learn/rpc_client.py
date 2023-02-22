import requests

res = requests.get('http://127.0.0.1:8003/?a=1&b=2')

print(res.text)

