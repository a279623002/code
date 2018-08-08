<?php
namespace app\index\model;
use think\Db;
use think\Model;

class Test extends Model
{
    public function getTest()
    {
        $result = Db::table('test')->order('id desc')->select();
        return $result;
    }

    public function addTest($data) {
        $Model = model('test');
        if(empty($data['name'])) return ['code'=>0, 'msg'=>'请输入名字！'];
        $data['addtime'] = time();
        $result = $Model->save($data);
        if($result) {
            return ['code'=>1, 'msg'=>'添加成功！', 'data'=>$result];
        } else {
            return ['code'=>0, 'msg'=>'添加失败！'];
        }
    }

    public function editTest($data) {
        $Model = model('test');
        if(empty($data['id'])) return ['code'=>0, 'msg'=>'缺少ID！'];
        if(empty($data['name'])) return ['code'=>0, 'msg'=>'缺少名字！'];
        $result = $Model->update($data);
        if($result) {
            return ['code'=>1, 'msg'=>'更新成功！', 'data'=>$result];
        } else {
            return ['code'=>0, 'msg'=>'更新失败！'];
        }
    }

    public function delTest($data) {
        $Model = model('test');
        if(empty($data['id'])) return ['code'=>0, 'msg'=>'缺少ID！'];
        $result = $Model->where('id', $data['id'])->delete();
        if($result) {
            return ['code'=>1, 'msg'=>'删除成功！', 'data'=>$result];
        } else {
            return ['code'=>0, 'msg'=>'删除失败！'];
        }
    }
}
