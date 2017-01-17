import requests

data = {
        "login_auth_request": {
    "username" : "kevince",
    "password": "abc",
    }
}

r = requests.post("http://127.0.0.1:8000/api/inline/login/auth", json=data)
print(r.status_code)
print(r.json())
