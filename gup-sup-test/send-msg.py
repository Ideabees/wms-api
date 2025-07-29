import requests

url = "https://api.gupshup.io/wa/api/v1/msg"

payload = {
    "channel": "whatsapp",
    "source": 917472850482,
    "destination": 918748133759,
    "message": "{\"type\":\"text\",\"text\":\"Hello user, how are you?\"}",
    "src.name": "convoEngage",
    "disablePreview": False,
    "encode": False
}
headers = {
    "accept": "application/json",
    "Content-Type": "application/x-www-form-urlencoded",
    "apikey": "sk_73a6f1b38e1b4cb5bbdae9ce323b03d1"
}

response = requests.post(url, data=payload, headers=headers)

print(response.text)