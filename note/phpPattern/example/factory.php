<?php
// 工厂模式

/**
 * 定义抽象类，让子类继承实现它
 */
abstract class Operation{
    abstract function getValue($num1, $num2);
}

/**
 * 加法类
 */
class OperationAdd extends Operation{
    function getValue($num1, $num2) {
        return $num1 + $num2;
    }
}

/**
 * 减法类
 */
class OperationSub extends Operation{
    function getValue($num1, $num2) {
        return $num1 - $num2;
    }
}

/**
 * 定义工厂类，为了创建对象
 */
class Factory{
    public static function createObj($operation) {
        switch($operation) {
            case '+':
                return new OperationAdd;
            case '-':
                return new OperationSub;
        }
    }
}
$test = Factory::createObj('+');
echo $test->getValue(1, 2).PHP_EOL;