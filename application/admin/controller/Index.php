<?php
namespace app\admin\controller;
use think\View;
use think\Db;
use think\Model;
use think\Request; 
use think\Session;
use think\Controller;
class Index extends Controller
{
    public function index()
    {   
        $user = model('Index');
        $view = new \think\View();
        return $view->fetch('index');
    }
    public function demo(){
        
    }
    public function check_login(){
            $data=input('post.');  
            $user = model('Index');
            $arr = $user->getAdmin($data['name'],$data['pass']);
            if($arr){
                $result['status'] = 1;
                Session::set('name',$arr['name']); 
            }else{
                $result['status'] = 0; 
            }
            // 返回JSON数据格式到客户端 包含状态信息
            return json($result);
    }
    public function is_login(){
        $name = Session::get('name');
        if(empty($name)){
            echo "<script>window.location.href = 'index';</script>";

            exit();
        }
    }
    public function content()
    {
        $this->is_login();
        $view = new \think\View();
        $name = Session::get('name');
        $view->name = $name;
        return $view->fetch('content');
    }
    public function column()
    {
        $this->is_login();

        $view = new \think\View();
        return $view->fetch('column');
    }
    public function userList()
    {
        $this->is_login();
        $user = model('Index');
        $arr = $user->getUserAll();
        $page = $arr->render();
        $view = new \think\View();
        $view->user = $arr;
        $view->page = $page;
        return $view->fetch('userList');
    }
    public function userDel()
    {
        $data=input('post.id');  
        $user = model('Index');
        $arr = $user->delUser($data);
        $result['see'] = $data;
        if($arr){
            $result['status'] = 1;
        }else{
            $result['status'] = 0; 
        }
        // 返回JSON数据格式到客户端 包含状态信息
        return json($result);
    }
    public function introduce()
    {
        $this->is_login();

        $view = new \think\View();
        return $view->fetch('introduce');
    }
    public function main()
    {
        $this->is_login();

        $view = new \think\View();
        return $view->fetch('main');
    }
    public function webset()
    {
        $this->is_login();

        $view = new \think\View();
        return $view->fetch('webset');
    }
}
