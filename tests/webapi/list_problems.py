import requests

session = requests.session()

def Login():
    data = {
        "username" : "kevince",
        "password": "abc",
    }
    r = session.post("http://127.0.0.1:8000/api/inline/login/auth", json=data)
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

    r = session.get("http://127.0.0.1:8000/api/inline/problems", json=data)
    print(r.status_code)
    print(r.content)

if __name__ == '__main__':
    Login()
    List()
