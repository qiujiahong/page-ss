# page-ss

## 简介

page screenshot 页面截图工具,可以使用page-ss给任意web界面截图。


## 使用举例

### 直接截图

1. 使用浏览器访问page-ss
2. page-ss反向代理访问web server，并截图
3. 截图返回给前端

```md
     ----------
    |          |
    |web server|
    |          |
     ----------
         ^
        (2)
         |
     ----------
    |          |
    | page-ss  |
    |          |
     ----------
         ^
         |
        (1)
```


* docker运行page-ss，代理截图百度

```bash 
docker run -p 8080:8080   \
  -e PPS_DBCONFIG_HOST=192.168.3.156 \
  -e PPS_DBCONFIG_USER=root \
  -e PPS_DBCONFIG_PASSWORD=123456 \
  -e PPS_DBCONFIG_PORT=3306 \
  -e PPS_DBCONFIG_DBNAME=db1 \
  -e PPS_PROXYURL=https://www.baidu.com \
  qiujiahong/page-ss:latest

```

* 访问服务``http://localhost:8080/pss/render`` ,浏览器将会返回百度页面的截图



### NGINX代理

如下图使用nginx代理截图，该场景下，可以配置page-ss转发header和cookie，这样在很多场景下就可以解决截图工具与被截图服务的鉴权问题。

```md
     ----------     ---------- 
    |          |(3)|          |
    |web server|<->|  page-ss |
    |          |   |          |
     ----------     ----------
                 ^
                (2)
                 |
             ----------
            |          |
            |   NGINX  |
            |          |
             ----------
                 ^
                 |
                (1)
```




## 调试程序 

```bash 
fresh 
```

## [配置参数](./docs/config.md)



## 单元测试

```bash
 $GOPATH/bin/goconvey
```


## todo 

- [ ] 支持css selector 部分截图


## 参考资料 

* [单元测试框架](https://github.com/smartystreets/goconvey/wiki/Documentation)
* [htttp服务框架](https://go-macaron.com/)