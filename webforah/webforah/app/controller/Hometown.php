<?php

namespace app\controller;

use app\BaseController;
use think\facade\Db;
use think\facade\Request;

class Hometown extends BaseController
{

    public function login() {
        $name = Request::param("username");
        $psw = Request::param("password");
        //查询数据
        $data = Db::table("login")->select();
        //检查数据是否存在，存在的话，可以登录，否则不可以
        //print_r($data);
        // -1 表示没有找到，1 代表找到

        $code = -1;
        $mes = "不存在该用户";
        foreach ($data as $k => $v) {
            if ($v["username"] == $name && $v["password"] == $psw) {
                $code = 1;
                $mes = "OK";
            }
        }

        return json(["mes" => $mes, "code" => $code]);

    }

    public function userRegister() {

        $name = Request::param("name");
        $psw = Request::param("pwd");
        $isFind = Db::table("register")->where("username", "=", $name)->find();
        $code = 0;
        $mes = "重复注册";
        if (!$isFind) {
            Db::table("register")->save([
                'username' => $name,
                'password' => $psw,
            ]);
            Db::table("login")->save([
                'username' => $name,
                'password' => $psw,
            ]);
            $code = 1;
            $mes = "注册成功";
        }
        return json(["code" => $code, "mes" => $mes]);
        //return json(["data" => $isFind, "null" => empty($data)]);

    }

    public function userlist() {

        $data = Db::table("login")->select();
        return json(["name" => $data]);

    }


}
