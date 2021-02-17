# page-ss

## 简介

page screenshot 页面截图工具,可以使用page-ss给任意web界面截图


## 调试程序 

```bash 
fresh 
```

## 可携带参数


|参数|参数变量|参数类型| 默认值 |
|---|-------|------|----|
|延迟截图      |__parDelay |int|  0|
| 强制刷新     |__forceUpdate |bool|  false|
| 是用使用缓存: |__useCache |bool | false|
| 是否自动刷新  |__autoFlush |bool | true|
| 有效期       |__validityDays |int | 90 |

## 单元测试

```bash
 $GOPATH/bin/goconvey
```


## todo 

- [ ] 支持css selector 部分截图


## 参考资料 

* [单元测试框架](https://github.com/smartystreets/goconvey/wiki/Documentation)
* [htttp服务框架](https://go-macaron.com/)