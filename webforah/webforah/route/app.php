<?php
// +----------------------------------------------------------------------
// | ThinkPHP [ WE CAN DO IT JUST THINK ]
// +----------------------------------------------------------------------
// | Copyright (c) 2006~2018 http://thinkphp.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: liu21st <liu21st@gmail.com>
// +----------------------------------------------------------------------
use think\facade\Route;

Route::get('think', function () {
    return 'hello,ThinkPHP6!';
});

Route::get('hello/:name', 'index/hello');


//Hometown属于阿亨的接口
Route::post('api/login', 'Hometown/login')->allowCrossDomain();

Route::get('api/register', 'Hometown/userRegister')->allowCrossDomain();

Route::get('api/userlist', 'Hometown/userlist');


//Fdm属于自己的的接口
Route::post('fdmapi/login', 'Fdm/login')->allowCrossDomain();

Route::get('fdmapi/register', 'Fdm/fdmregister')->allowCrossDomain();
