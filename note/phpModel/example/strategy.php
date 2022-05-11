<?php
// 策略模式

/**
 * 定义策略抽象类，让子类继承并实现它
 */
abstract class Strategy{
    abstract function goSchool();
}


/**
 *实现接口
 */
class walk extends Strategy{
    function goSchool() {
        echo '走路去学校'.PHP_EOL;
    }
}

class subway extends Strategy{
    function goSchool() {
        echo '坐地铁去学校'.PHP_EOL;
    }
}

class bike extends Strategy{
    function goSchool() {
        echo '骑自行车去学校'.PHP_EOL;
    }
}

/**
 * 主类
 */
class main{
    protected $_stratege;

    function goSchool($stratege) {
        $this->_stratege = $stratege;
        $this->_stratege->goSchool();
    }
}

$main = new main();
$main->goSchool(new bike());