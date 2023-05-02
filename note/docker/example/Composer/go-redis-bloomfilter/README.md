#### 布隆过滤器
##### api
```
// 查询键是否存在
// get
// query key string
http://127.0.0.1:6573/bloomFilter/isExists
```

```
// 存入
// POST
// query key string
// query val string
http://127.0.0.1:6573/bloomFilter/set
```

```
// 获取
// func get
// query key string
http://127.0.0.1:6573/bloomFilter/get
```