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
        "username" : "test_root",
        "password": "abc123",
        }
    }
    r = session.post(url+"/api/inline/login/auth", json=data)
    dump(r.json())

def Save():
    data = {
        "save_problem_request": {
            "oj_name": "hdu",
            "oj_pid": "5698",
            "title": "瞬间移动",
            "description":'''<script type='text/x-mathjax-config'>MathJax.Hub.Config({tex2jax: { inlineMath: [['$','$'],['\[','\]']] } }); </script>
            <script type='text/javascript' src='https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.1/MathJax.js?config=TeX-AMS-MML_HTMLorMML'></script>
            <script type='text/javascript'>setTimeout(function(){MathJax.Hub.Queue(['Typeset', MathJax.Hub, 'left_view']);}, 2000);</script>
            <div class="panel_content">
              有一个无限大的矩形，初始时你在左上角（即第一行第一列），每次你都可以选择一个右下方格子，并瞬移过去（如从下图中的红色格子能直接瞬移到蓝色格子），求到第$n$行第$m$列的格子有几种方案，答案对$1000000007$取模。 
               <br> 
                <br> 
                 <center> 
                   <img style="max-width:100%;" src="https://odzkskevi.qnssl.com/6be27c7c8f0609e142553195f49b804c?v=1495902953" SRC="https://odzkskevi.qnssl.com/6be27c7c8f0609e142553195f49b804c?v=1495902953"> 
                    </center> 
                    </div>''',
            "input":'''多组测试数据。 
                    <br> 
                    <br>两个整数$n,m(2\leq n,m\leq 100000)$ 
                    <br>''',
            "output":'''一个整数表示答案''',
            "sample_in": '''<pre>4 5</pre>''',
            "sample_out":'''<pre>10</pre>''',
            "source": "HDU",
            "limits": [
                {"lanuage":"java", "time_limit":4000, "memory_limit":65536}, 
                {"language":"others", "time_limit":2000, "memory_limit":65536}]
        }

    }

    r = session.post(url+"/api/inline/save_problem", json=data)
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

    r = session.post(url+"/api/inline/save_problem", json=data)
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
    # Logout() 
    # Submit()
