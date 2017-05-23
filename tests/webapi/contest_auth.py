import requests

session = requests.Session()

def login():
    data = {
    "login_auth_request": {
        "username" : "kevince",
        "password": "abc",
    }
    }

    r = session.post("http://127.0.0.1:8000/api/inline/login/auth", json=data)
    print(r.headers)
    print(r.status_code)
    print(r.json())

def auth():
    data = {
        "contest_auth_request": {
            "contest_id": 8,
            "password": "123456",
        }
    }
    r = session.post("http://127.0.0.1:8000/api/inline/contest_auth", json=data)
    print(r.json())

login()
auth()
