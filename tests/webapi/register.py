import requests

def login(sess):
    data = {
        "login_auth_request": {
            "username" : "kevince",
            "password": "abc",
        }
    }

    r = sess.post("http://127.0.0.1:8000/api/inline/login/auth", json=data)
    print(r.status_code)
    print(r.json())

data = {
        "register_request": {
    "username" : "hong",
    "password": "abc",
    "email": "123@123.com",
    }
}

s = requests.session()
# login(s)
r = s.get("http://127.0.0.1:8000/api/inline/captcha")
# r = s.get("http://35.189.170.28:8000/api/inline/captcha")
print(r.headers)
with open('captcha.png', 'wb') as f:
    f.write(r.content)

captcha = raw_input()
data["captcha"] = captcha

r = s.post("http://127.0.0.1:8000/api/inline/register", json=data)
# r = s.post("http://35.189.170.28:8000/api/inline/register", json=data)
print(r.status_code)
print(r.json())
