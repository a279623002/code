<?php

interface animal{
    public function watch();
}

interface subject{
    public function register(animal $animal);
    public function notify();
}

// 主题
class Action implements subject{
    private $_animal = [];

    public function register($animal) {
        $this->_animal[] = $animal;
    }

    public function notify() {
        foreach ($this->_animal as $a) {
            $a->watch();
        }
    }
}

class Dog implements animal{
    public function watch() {
        echo 'Dog Watch'.PHP_EOL;
    }
}

class Cat implements animal{
    public function watch() {
        echo 'Cat Watch'.PHP_EOL;
    }
}

// 观察者模式可根据主题规定的方法调用所有注册的类里面的方法
$act = new Action();
$act->register(new Dog);
$act->register(new Cat);
$act->notify();