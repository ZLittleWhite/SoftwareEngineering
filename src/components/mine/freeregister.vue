<template>
  <div>

    <div class="tips" style="margin: 40px 0 15px 0">

      <h2 class="h2_tip">免费注册</h2>
    </div>

    <van-form @submit="onSubmit" >
      <van-field
        v-model="cellphone"
        name="手机号"
        label="中国+86"
        placeholder="请输入手机号"
      />
      <van-field
        v-model="code"
        name="验证码"
        placeholder="请输入6位验证码"
      />
      <div>
        <span class="code">获取验证码</span>
      </div>
      <van-field
        v-model="password"
        type="password"
        name="密码"
        placeholder="请设置8-20位(数字+字符)密码"
      />
      <div style="position: fixed;top: 180px;right: 52px">
        <van-icon name="closed-eye"/>
      </div>
      <van-field
        v-model="password_again"
        type="password"
        name="确认密码"
        placeholder="请再次输入密码"
      />
      <div style="position: fixed;top: 225px;right: 52px">
        <van-icon name="closed-eye"/>
      </div>
      <div style="margin: 16px;">
        <van-button square block style="background: #30e084;color: white" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>


    <!--    <div style="width: 90%;margin: 0 auto;font-size: 10px;color: blue">-->
    <!--      &lt;!&ndash; 两端对齐 &ndash;&gt;-->
    <!--      <van-row type="flex" justify="space-between">-->
    <!--        <van-col span="８" @click="toPasswordLogin">密码登录</van-col>-->
    <!--        <van-col span="8">-->
    <!--          <div style="float:right;">免费注册</div>-->
    <!--        </van-col>-->
    <!--      </van-row>-->
    <!--    </div>-->


    <!--    同意条款-->

    <div class="follow">
      <van-radio-group v-model="radio">
        <diV style="float:left;">
          <van-radio
            style="float:left;"
            name="1"
            checked-color="#07c160"
            icon-size="12px">
          </van-radio>
          <div style="float:left;margin-left: 5px">
            我已阅读并同意<b @click="toUserAgree" style="text-decoration: underline">《用户服务协议》</b>
          </div>
        </diV>
      </van-radio-group>
    </div>


  </div>
</template>

<script>
  export default {
    name: "freeregister",
    data() {
      return {
        cellphone: '',
        password: '',
        password_again: '',
        code: '',
        radio: '0',
      }
    },
    methods: {
      onSubmit(values) {
        console.log('submit', values);
      },
      toPasswordLogin() {
        this.$router.push('/mine')
      },
      toUserAgree() {
        this.$router.push('/useragree')
      }
    }
  }
</script>

<style scoped>
  .h2_tip, .h4_tip {
    margin: 20px 0 0 15px;
  }

  .code {
    color: blue;
    cursor: pointer;
    float: right;
    position: fixed;
    right: 22px;
    top: 140px;
  }

  .follow {
    font-size: 10px;
    margin: 0 auto;
    width: 230px;
    height: 20px;

  }

</style>

