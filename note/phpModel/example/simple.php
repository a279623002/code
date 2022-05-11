<?php
// 单例模式
class main {
    //保存实例在此属性中
    private static $_instance;

    //构造方法设置为private，防止直接创建对象
    private function __construct() {}

    //防止用户复制对象实例
    private function __clone() {
        trigger_error('Clone is not allow', E_USER_ERROR);
    }

    //单例方法
    public static function getInstance() {
        if (!isset(self::$_instance)) {
            self::$_instance = new self();
        }
        return self::$_instance;
    }

    function test() {
        echo 'test'.PHP_EOL;
    }
}

$test = main::getInstance();
$test->test();