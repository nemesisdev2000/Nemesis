import requests

def sendRequest():
    url = "http://localhost:8000/listen"
    headers = {
            'Content-Type': 'application/json'
            }
    data = {
            'type': 'TcpListener',
            'port': '1331'
            }
    try:
        requests.post(url = url, json = data, headers = headers, timeout=0.3)
    except Exception as e:
        return

sendRequest()
