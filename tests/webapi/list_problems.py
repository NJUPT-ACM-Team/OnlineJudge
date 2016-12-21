import requests
import json

session = requests.session()

def Login():
    data = {
        "username" : "kevince",
        "password": "abc",
    }
    r = session.post("http://127.0.0.1:8000/api/inline/login/auth", json={"login_auth_request":data})
    print(r.json())

def List():
    data = {
        "per_page": 2,
        "current_page": 1,
        "order_by": 0,
        "isDesc": True,
        "filter": {
            "oj": "zoj",
            "p_status": 0,
        },
    }
    send = {
        "list_problems_request": data
    }

    r = session.get("http://127.0.0.1:8000/api/inline/problems", json=send)
    print(r.status_code)
    print(json.dumps(r.json(), indent=4))

if __name__ == '__main__':
    Login()
    List()
