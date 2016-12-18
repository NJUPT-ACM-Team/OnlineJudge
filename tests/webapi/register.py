import requests

data = {
        "register_request": {
    "username" : "kevince",
    "password": "abc",
    "email": "123@123.com",
    }
}

r = requests.post("http://127.0.0.1:8000/api/inline/register", json=data)
print(r.status_code)
print(r.json())
