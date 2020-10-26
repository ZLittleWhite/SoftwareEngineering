<template>
    <div>
        <div class="noThings">
            <!--            头部分享个朋友-->
            <div>
                <van-nav-bar @click-left="backToMine()" title="用户注册" left-text="返回" left-arrow>
                </van-nav-bar>
            </div>
        </div>
        <!--        注册-->
        <div>
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
                <van-field
                        v-model="cfpassword"
                        type="password"
                        name="cfpwd"
                        label="确认密码"
                        placeholder="请再次填写密码"
                        :rules="[{ required: true, message: '请再次填写密码' }]"
                />
                <div style="margin: 16px;">
                    <van-button round block type="info" native-type="submit">
                        提交
                    </van-button>
                </div>
            </van-form>
        </div>
    </div>
</template>

<script>
    import {Toast} from "vant";

    export default {
        name: "register",
        data() {
            return {
                username: '',
                password: '',
                cfpassword: '',
                goToLogin: false
            };
        },
        methods: {
            backToMine() {
                this.$router.push("/mine")
            },
            onSubmit(values) {
                console.log('submit', values);
                if (values.pwd != values.cfpwd) {
                    Toast("两次密码不一致")
                } else {
                    //http://webforah:90/fdmapi/register?name=zodggud11&pwd=123488
                    let registerURL = "http://webforah:90/fdmapi/register"
                    this.$axios.get(registerURL, {
                        params: {
                            name: values.name,
                            pwd: values.pwd
                        }
                    }).then(res => {
                        console.log(res);
                        if (res.data.code === 1) {
                            Toast("注册成功！")
                            if (this.goToLogin) {
                                this.$router.push("/login")
                            } else {
                                this.$router.push("/mine")
                            }
                        } else if (res.data.code === 0) {
                            Toast("重复注册！")
                            this.username = '';
                        }
                    }).catch(err => {
                        console.log(err);
                    })

                }
            },
            getShouldLoginValue() {
                let isShouldvalue = localStorage.getItem("shouldLogin");
                console.log("login", isShouldvalue);
                if (isShouldvalue === undefined || isShouldvalue === null) {
                    this.goToLogin = true;
                }
            },
        },
        created() {
            this.getShouldLoginValue()
        }
    }
</script>

<style scoped>
    .noThings {
        width: 100%;
        height: 80px;
    }
</style>