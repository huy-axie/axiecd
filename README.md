# axiecd

```
Helper

go run main.go -h                                                                                              
Usage of /var/folders/z8/dgpdbww538x3gbfs0_f0wbg00000gn/T/go-build3834845186/b001/exe/main:
  -c string
        Container name
  -e string
        Env file
  -i string
        Container image
  -m string
        Mount volume locapath:containerpath
  -p string
        Port open

example:
    ➜  axieCD git:(main) ✗ go run main.go -c "test" -e "/Users/huyt/go/src/github/axieCD/env" -i "nginx:latest" -m "/root/go  run main.go -c "test" -e "/Users/huyt/go/src/github/axieCD/env" -i "nginx:latest" -m "/root/volume" -p "443" -m "/Users/huyt/go/src/github/axieCD/env:/env"
      ----- Deploy new container 
      ----- Name: test
      ----- Image: nginx:latest
      ----- Port: 443
      ----- Env file: /Users/huyt/go/src/github/axieCD/env
      ----- Volume mount: /Users/huyt/go/src/github/axieCD/env:/env
      ----> Host: MacBook-Air-cua-huyt.local
      2021/06/24 21:25:11 Container test has been stopped & removed !!!
      2021/06/24 21:25:13 Container 18421f8d8cb414dc601e2d5083b805d5faa00d7cb7eb0a5b2c36c2a6d0fed596 is created
```