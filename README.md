# gin-demo-app

A demo app that implemented after the tutorial:
https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin

Powered on Gin framework: https://github.com/gin-gonic/gin

## HTML/JSON/XML support

```bash
curl -X GET -H "Accept: application/json" http://localhost:8080/

[{"id":1,"title":"Article 1","content":"Article 1 body"},{"id":2,"title":"Article 2","content":"Article 2 body"}]
```

```bash
curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1

<article><ID>1</ID><Title>Article 1</Title><Content>Article 1 body</Content></article>%  
```
