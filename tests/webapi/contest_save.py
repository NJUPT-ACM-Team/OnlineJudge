#!coding:utf8
import requests
import json

session = requests.session()

url = "http://127.0.0.1:8000"
# url = "http://35.189.170.28:8000"

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
    r = session.post(url+"/api/inline/login/auth", json=data)
    dump(r.json())

def Save():
    data = {
        "contest_save_request": {
            "title": "有密码的比赛10",
            "description": "这是第n次比赛",
            "is_virtual": True,
            "contest_type": "icpc",
            "password": "123456",
            "start_time": "Fri, 25 May 2017 16:25:00 GMT",
            "end_time": "Fri, 30 May 2017 20:25:00 GMT",
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

    r = session.post(url+"/api/inline/contest/save", json=data)
    print(r.status_code)
    dump(r.json())

def Update():
    data = {
        "contest_save_request": {
            "contest_id": 5,
            "title": "第三次比赛",
            "description": "这是第四次比赛",
            "is_virtual": True,
            "contest_type": "icpc",
            "start_time": "Fri, 30 May 2017 09:25:00 GMT",
            "end_time": "Fri, 30 May 2017 12:25:00 GMT",
            "problems": [
                {
                    "problem_sid": "local-1000",
                    "alias": "No.2"
                },
                {
                    "problem_sid": "zoj-1001", "alias":"fuck"
                },
            ]
        }

    }

    r = session.post(url+"/api/inline/contest/save", json=data)
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
    # Save()
    Update()
    # Logout() 
    # Submit()
