<?php
$script = <<<EOF
    local _key = KEYS[1]

    local _val = tonumber(ARGV[1])
    local stock = tonumber(redis.call('get', _key))

    if stock > 0 then
        if stock >= _val then
            redis.call('decrby', _key, _val)
            return 1
        else
            return -1
        end
    end

    return 0

EOF;

// 获取传过来的变量

$stock = 1;
$redis = new \Redis();
$redis->connect('127.0.0.1', 6379);
$redis->select(1);
$result = $redis->eval($script, ['lua:test', $stock], 1);

if ($result == 1) {
    $servername = "172.31.96.1";
    $username = "root";
    $password = "root";
    $db = 'test';

    // 创建连接
    $conn = new mysqli($servername, $username, $password, $db);

    // 检测连接
    if ($conn->connect_error) {
        die("连接失败: " . $conn->connect_error);
    }

    /*
    $sql = "SELECT * FROM test";
    $result = $conn->query($sql);

    if ($result->num_rows > 0) {
        // 输出数据
        while($row = $result->fetch_assoc()) {
            print_r($row);
        }
    } else {
        echo "0 结果";
    }
    */
    $sql = 'update test set stock = stock -1 where id = 1';
    $conn->query($sql);

    $conn->close();
}
