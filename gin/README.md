## gin

> go编写的web框架，有很好的性能比 martini，速度提高了40 倍[httprouter](https://github.com/julienschmidt/httprouter)



## HttpRouter

1. 轻量的高性能http请求路由器
2. 对高性能和较小的内存占用进行了优化

## gin 问题
1. go get gin失败解决方式

   ```
   ➜  pika git:(master) ✗ go get -u github.com/gin-gonic/gin
   go: golang.org/x/sys@v0.0.0-20190222072716-a9d3bda3a223: unrecognized import path "golang.org/x/sys" (https fetch: Get https://golang.org/x/sys?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
   go: golang.org/x/net@v0.0.0-20190503192946-f4e77d36d62c: unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
   go: error loading module requirements
   ➜  pika git:(master) ✗ export GOPROXY=https://goproxy.io
   ➜  pika git:(master) ✗ go get -u github.com/gin-gonic/gin
   
   
   ```

2. 