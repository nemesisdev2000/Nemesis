import requests

def sendRequest():
    url = input("Enter URL : ")
    data = {
            'Content-Type':'application/json',
            'type': 'TcpListener',
            'port': '1331'
            }
    r = requests.post(url = url, data = data)
    print(r.text)

sendRequest()
