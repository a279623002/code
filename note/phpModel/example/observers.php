<?php
// 观察者模式

/**
 * 主题接口
 */
interface subject{
    public function register(observer $observer);
    public function notify();
}

/**
 * 观察者接口
 */
interface observer{
    public function watch();
}

/**
 * 主题
 */
class action implements subject{
    public $_observers = [];

    function register(observer $observer){
        $this->_observers[] = $observer;
    }

    function notify() {
        foreach ($this->_observers as $observer) {
            $observer->watch();
        }
    }
}

/**
 * 观察者
 */
class cat implements observer{
    function watch() {
        echo 'cat watched tv'.PHP_EOL;
    }
}

class dog implements observer{
    function watch() {
        echo 'dog watched tv'.PHP_EOL;
    }
}

$action = new action();
$action->register(new cat());
$action->register(new dog());
$action->notify();