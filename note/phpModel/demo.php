<?php

$str = '****xyz23333***';
$preg = '/(.*xyz\d)/';

// echo preg_match($preg, $str);

preg_match_all($preg, $str, $res);
print_r($res);

function a() {
    include 'index1.html';
    echo 2233;
}
@a();

//require 'index.html';
echo 233;