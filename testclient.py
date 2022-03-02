import requests

def startListener():
    token = input("Enter Token : ")
    url = "http://127.0.0.1:5555/api/listen"
    headers = {
            "Content-Type": "application/json"
            }
    data = {
            "type": "TcpListener",
            "port": "1331",
            "token": token
            }
    try:
        res = requests.post(url = url, json = data, headers = headers)
        print(res.text)
    except Exception as e:
        return

def stopListener():
    token = input("Enter Token : ")
    ID = input("Enter Listener ID : ")
    url = "http://127.0.0.1:5555/api/stoplistener"
    header = {
            "Content-Type": "application/json"
            }
    data = {
            "type": "TcpListener",
            "port": "1331",
            "token": token,
            "id": ID
            }
    try:
        res = requests.post(url = url, json = data, headers = header)
        print(res.text)
    except Exception as e:
        print(e)
        return

def checksignin():
    url = "http://localhost:5555/auth/signin"
    headers = {
            'Username': 'admin',
            'Passwordhash': 'passw0rd'
            }
    try:
        res = requests.get(url = url, headers = headers)
        print(res.text)
    except Exception as e:
        return

def checksignup():
    url = "http://localhost:5555/auth/signup"
    headers = {
            'Username': 'ritaban',
            'Passwordhash': 'password1'
            }
    try:
        res = requests.post(url = url, headers = headers)
        print(res.text)
    except Exception as e:
        return

def checksignout():
    url = "http://localhost:5555/auth/signout"
    token = input("Enter the token : ")
    header = {
            'Username' : 'admin',
            'Token': token
            }
    try:
        res = requests.get(url = url, headers = header)
        print(res.text)
    except Exception as e:
        print(e)
        return

checksignin()
startListener()
stopListener()
checksignout()
