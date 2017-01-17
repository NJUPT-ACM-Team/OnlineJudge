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

def Submit():
    data = {
        "submit_request": {
            "problem_sid" : "zoj#1000",
            "code":"""
#include <iostream>
using namespace std;
int main() {
    int a, b;
    cin >> a >> b;
    cout << a + b << endl;
    return 0;
}
                """,
            "language_id": 1,
            "is_shared" : True,
        },
    }

    r = session.post("http://127.0.0.1:8000/api/inline/submit", json=data)
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
    # Logout() 
    Submit()
