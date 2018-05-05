```bash
go get -u github.com/anthonynsimon/bild/...
go get github.com/betacraft/yaag/...
go get -v github.com/kataras/iris
go get -u github.com/iris-contrib/httpexpect
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/labstack/echo
go get -u github.com/facebookgo/grace/gracehttp
go get -u github.com/tylerb/graceful
go get -v -u github.com/disintegration/imaging
go get -v -u github.com/parnurzeal/gorequest
go get -v -u github.com/Kodeworks/golang-image-ico
go get -v -u github.com/dkua/go-ico
go get -v -u github.com/google/go-querystring
```


```
protoc --go_out=plugins=grpc,paths=source_relative:. im/*.proto
```