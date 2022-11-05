<?php

abstract class animal{
    abstract public function watch();
}

abstract class subject{
    abstract public function register(animal $animal);
    abstract public function notify();
}

class Action extends subject{
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

class Dog extends animal{
    public function watch() {
        echo 'Dog Watch'.PHP_EOL;
    }
}

class Cat extends animal{
    public function watch() {
        echo 'Cat Watch'.PHP_EOL;
    }
}

$act = new Action();
$act->register(new Dog);
$act->register(new Cat);
$act->notify();