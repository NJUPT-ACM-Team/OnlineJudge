import requests
import json

session = requests.session()

def List():
    data = {
        "list_submissions_request": {
            "per_page": 1,
            "current_page": 1,
            "filter": {

            },
        }        
    }
    r = session.get("http://127.0.0.1:8000/api/inline/status", json=data)
    print(r.status_code)
    print json.dumps(r.json(), indent=4)

def main():
    List()

if __name__ == '__main__':
    main()

