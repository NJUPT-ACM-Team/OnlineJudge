import requests
import json

data = {
    "about_request": {
    "need_ojs_list" : True,
    "need_languages_list": True,
    }
}

r = requests.get("http://127.0.0.1:8000/api/inline/about", json=data)
print(r.status_code)
print(json.dumps(r.json(), indent=4))
