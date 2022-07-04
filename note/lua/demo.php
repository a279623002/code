<?php

$script = <<<EOF
    local _key = KEYS[1]

    local _val = ARGV[1]

    local result = redis.call('GET', _key);
    result = result and result or ""

    local text = ''
    if result == '' then
        return text
    else
        text = result .. _val
        redis.call('SET', _key, text)
    end

    return text
EOF;

// 获取传过来的变量
$text = isset($argv[1]) ? $argv[1] : '';
$redis = new \Redis();
$redis->connect('127.0.0.1', 6379);
$redis->select(1);
$result = $redis->eval($script, array("lua:test", $text), 1);
echo $result;