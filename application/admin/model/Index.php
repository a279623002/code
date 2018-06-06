<?php

namespace app\admin\model;
use think\Model;
use think\Db;
class Index extends Model{
    public function getAdmin($name,$pass){
        $pass = md5($pass);
        $user = Db::table('Admin')->where(array('name'=>$name,'pass'=>$pass))->find();
        return $user;
    }
    public function getUserAll(){
        $user = Db::table('user')->paginate(10);
        return $user;
    }
    public function delUser($id){
        $data = Db::table('user')->where(array('Id'=>$id))->delete();
        return $data;
    }
}