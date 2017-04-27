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
        "is_desc": True,
        "filter_oj": "xoj",
        "filter_p_status": 0,
    }

    r = session.get("http://127.0.0.1:8000/api/inline/problems", params=data)
    print(r.status_code)
    print(json.dumps(r.json(), indent=4))
    print(r.headers)

if __name__ == '__main__':
    # Login()
    List()
