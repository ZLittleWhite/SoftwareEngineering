<template>
    <div>
        <!--        登录框-->
        <div v-show="shouldLogin">
            <div class="noThings">
                <!--            头部分享个朋友-->
                <div>
                    <van-nav-bar @click-left="backToMine(shouldLogin)" title="用户登录" left-text="返回" left-arrow>
                    </van-nav-bar>
                </div>
            </div>
            <van-form @submit="onSubmit">
                <van-field
                        v-model="username"
                        name="name"
                        label="用户名"
                        placeholder="用户名"
                        :rules="[{ required: true, message: '请填写用户名' }]"
                />
                <van-field
                        v-model="password"
                        type="password"
                        name="pwd"
                        label="密码"
                        placeholder="密码"
                        :rules="[{ required: true, message: '请填写密码' }]"
                />
                <div style="margin: 16px;">
                    <van-button round block type="info" native-type="submit">
                        提交
                    </van-button>
                </div>
            </van-form>
        </div>
        <!--        更换头像-->
        <div v-show="!shouldLogin">
            <!--            头部分享个朋友-->
            <div>
                <van-nav-bar @click-left="backToMine()" title="个人信息" left-text="返回" left-arrow>
                </van-nav-bar>
            </div>
            <div>
                <van-cell-group>
                    <van-cell title="">
                        <template #label>
                            <div>
                                <br/>
                                <b style="font-size: 15px;color: #575757;margin: 20px 0 0 0"> 头像</b>
                            </div>
                        </template>
                        <template #extra>
                            <div>
                                <van-image
                                        width="3rem"
                                        height="3rem"
                                        src="https://img.yzcdn.cn/vant/cat.jpeg"
                                />
                            </div>
                        </template>
                    </van-cell>
                </van-cell-group>
                <van-cell-group v-for="(item,index) in userList" :key="index">
                    <van-cell :title="item.title" :value="item.value" is-link></van-cell>
                </van-cell-group>
            </div>
            <!--            提示-->
            <div>
                <van-notice-bar
                        color="#1989fa" background="#ecf9ff"
                        scrollable
                        left-icon="volume-o"
                        text="返回可能将导致数据丢失，可能需要重新登录"
                />
            </div>
            <!--            注销登录-->
            <div>
                <van-row>
                    <van-col class="cancelLogin" span="24">
                        <van-button @click="removeLogin" round icon="wap-home-o" type="info">
                            注销登录
                        </van-button>
                    </van-col>
                </van-row>
            </div>
        </div>
    </div>
</template>

<script>
    import Vue from 'vue';
    import {Toast} from 'vant';

    Vue.use(Toast);
    export default {
        name: "login",
        data() {
            return {
                username: '',
                password: '',
                shouldLogin: true,
                userList: [
                    {title: "二维码名片", value: ""},
                    {title: "昵称", value: "周敌"},
                    {title: "手机号", value: "17379896567"},
                    {title: "更多", value: ""},
                    {title: "我的地址", value: ""},
                ]
            };
        },
        methods: {
            removeLogin() {
                localStorage.setItem("shouldLogin", 1)
                localStorage.removeItem("username")
                this.$router.push("/mine")
                // this.$router.push("/login")
            },
            //兄弟间传值，把这个false值传递到sessionStore界面，当再次点击login界面时，现获取这个传递过来的false值，可以辨别是否登录过
            backToMine() {
                this.$router.push("/mine")
            },
            onSubmit(values) {
                console.log('submit', values);
                //判断是否登录
                // http://webforah:90/api/login?username=zhoudi&password=zhoudi
                let loginURL = "http://webforah:90/fdmapi/login"
                this.$axios.post(loginURL, {
                    username: values.name,
                    password: values.pwd,
                }).then(res => {
                    console.log(res.data.code);
                    if (res.data.code == 1) {
                        //传递这个值给sessionStorage，当再次点击login时，在created 中获取，v-if显示时，就不必再次登录
                        localStorage.setItem("shouldLogin", 0)
                        localStorage.setItem("username", this.username)
                        //设置昵称
                        this.userList[1].value = this.username
                        const toast = Toast.loading({
                            duration: 0, // 持续展示 toast
                            forbidClick: true,
                            message: '稍等片刻',
                        });
                        let second = 1;
                        const timer = setInterval(() => {
                            second--;
                            if (second) {
                                toast.message = "";
                            } else {
                                this.shouldLogin = false;
                                clearInterval(timer);
                                // 手动清除 Toast
                                Toast.clear();
                            }
                        }, 1000);
                    } else {
                        Toast.fail('用户不存在');
                        this.username = "";
                        this.password = "";
                    }
                }).catch(err => {
                    console.log(err);
                })
            },
            getShouldLoginValue() {
                let isShouldvalue = localStorage.getItem("shouldLogin");
                console.log("login", isShouldvalue);
                if (isShouldvalue == 0) {
                    this.shouldLogin = false;
                }
                if (isShouldvalue === undefined || isShouldvalue === null) {
                    this.shouldLogin = true;
                }
            },
            getUsernameFromSession() {
                let nameValue = localStorage.getItem("username");
                if (nameValue === undefined || nameValue === null) {
                    console.log("未登陆");
                } else {
                    console.log(nameValue);
                    this.userList[1].value = nameValue
                    //this.nameValue = true;
                }
            }
        },
        created() {
            this.getUsernameFromSession()
        },
        mounted() {
            this.getShouldLoginValue()
        },

    }
</script>

<style scoped>
    .noThings {
        width: 100%;
        height: 80px;
    }

    .cancelLogin {
        margin: 20px 0 0 0;
        text-align: center
    }
</style>