# OnlineJudge
A new generation online judge for NJUPT ACM TEAM

* The structure of all modules of the backend.
```
$GOPATH
  |- bin/
  |- pkg/
  └- src/
     |- OnlineJudge/
     |  |- <some common libraries>
     |  |- <config files etc.>
     |  |- WebBackend/
     |  |- Daemon/
     |  └- <other parts>
     |- LocalJudger/
     |- VirtualJudger/
     |- ...
     |- github.com/
     |  └- ...
     └- ...
```

## Design Doc
[Design Doc Wiki](https://github.com/NJUPT-ACM-Team/OnlineJudge/wiki/Design-Doc)
