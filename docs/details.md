# NGINX代理办法

## 使用


* 启动容器  

```bash 
# 创建网络  
docker network ls
docker network create -d bridge mybridge 

docker run -d  --name grafana \
    --network mybridge\
    -p 3000:3000 grafana/grafana:latest


docker run -d --name pss -p 8080:8080   \
    --network mybridge \
  -e PPS_DBCONFIG_HOST=192.168.3.156 \
  -e PPS_DBCONFIG_USER=root \
  -e PPS_DBCONFIG_PASSWORD=123456 \
  -e PPS_DBCONFIG_PORT=3306 \
  -e PPS_DBCONFIG_DBNAME=db1 \
  -e PPS_PROXYURL=http://grafana:3000 \
  qiujiahong/page-ss:latest
  


docker run --name nginx \
    --network mybridge  \
    -v   $(pwd)/nginx/nginx.conf:/etc/nginx/nginx.conf:ro\
    -d -p 80:80 nginx:latest


```


* 访问截图

```bash 
# 截图，使用缓冲，延迟1S
http://localhost/pss/renderWithHeader/d/Y-8g9QPMk/test1?orgId=1&kiosk&__parDelay=1&__useCache=true
# 截图，使用缓存、延迟1S、强制刷新
http://localhost/pss/renderWithHeader/d/Y-8g9QPMk/test1?orgId=1&kiosk&__parDelay=1&__useCache=true&__forceUpdate=true

```

* 更多可配置参数参考[参数详细说明](config.md)

```bash 
docker stop nginx
docker stop pss
docker stop grafana

docker rm nginx
docker rm pss
docker rm grafana

```