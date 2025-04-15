##### skill
1. initAct
 * hmset(act_id, actInfo)
 * set(act_id, stock)
2. skill
 * limit_user
 * check_user_order -> hget(user_id, order)
 * check_stock
   ```
    stock = redis.Get(act_id)
    if stock > 0 {
        // 扣减库存
        // redis->lua->decr stock
    }else {
        // 获取所有订单orders
        // orders已支付状态
    if len(orders) < stock {
        // 被占领
    }else {
        // 已抢完
    }}
  ```
* create orde
  ```
    // 1.异步生成订单
    order[order_id] = 雪花id
    go createOrder(order)
    // 2. 刷到redis
    redis.ZSet(act_id, order_id)
    redis.HSet(order_id, order)
  ```
3. 获取所有订单
 ```
    order_ids = redis.ZRange(act_id)
    for order_id = range order_ids {
        order = redis.HGet(order_id)
        orders = append(orders, order)
    }
 ```