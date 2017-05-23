#!coding:utf8
import requests
import json

session = requests.session()

def dump(js):
    print json.dumps(js, indent=4)

def Login():
    data = {
        "login_auth_request":
        {
        "username" : "kevince",
        "password": "abc",
        }
    }
    r = session.post("http://127.0.0.1:8000/api/inline/login/auth", json=data)
    dump(r.json())

def Save():
    data = {
        "save_contest_request": {
            "title": "有密码的比赛",
            "description": "这是第四次比赛",
            "is_virtual": True,
            "contest_type": "icpc",
            "password": "123456",
            "problems": [
                {
                    "problem_sid": "zoj-1000",
                },
                {
                    "problem_sid": "local-1000",
                }
            ]
        }

    }

    r = session.post("http://127.0.0.1:8000/api/inline/save_contest", json=data)
    print(r.status_code)
    dump(r.json())

def Update():
    data = {
        "save_contest_request": {
            "contest_id": 6,
            "title": "不是瞎搞的比赛",
            "description": "这是第四次比赛",
            "is_virtual": True,
            "contest_type": "icpc",
            "problems": [
                {
                    "problem_sid": "zoj-1000",
                },
                {
                    "problem_sid": "zoj-1001",
                }
            ]
        }

    }

    r = session.post("http://127.0.0.1:8000/api/inline/save_contest", json=data)
    print(r.status_code)
    dump(r.json())


def Logout():
    data = {
            "logout_request": {}
    }
    r = session.post("http://127.0.0.1:8000/api/inline/logout", json=data)
    print r.status_code
    dump(r.json())


if __name__ == '__main__':
    Login()
    Save()
    # Update()
    # Logout() 
    # Submit()
