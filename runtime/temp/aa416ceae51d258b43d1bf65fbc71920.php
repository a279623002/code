<?php if (!defined('THINK_PATH')) exit(); /*a:1:{s:62:"D:\phpStudy\WWW\code./application/index\view\test\getTest.html";i:1533882151;}*/ ?>
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <table  border="1" cellspacing="10">
        <tr>
            <th>id</th>
            <th>name</th>
            <th>addtime</th>
        </tr>
            <?php if(is_array($list) || $list instanceof \think\Collection || $list instanceof \think\Paginator): if( count($list)==0 ) : echo "" ;else: foreach($list as $k=>$vo): ?>
            <tr>
                <td align="center"><?php echo $vo['id']; ?></td>
                <td align="center"><?php echo $vo['name']; ?></td>
                <td align="center"><?php echo $vo['addtime']; ?></td>
            </tr>
            <?php endforeach; endif; else: echo "" ;endif; ?>
    </table>
</body>

</html>