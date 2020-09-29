<template>
  <div>

    <div class="tips" style="margin: 40px 0 15px 0">

      <h2 class="h2_tip">账号密码登录</h2>
      <h5 class="h4_tip" style="color: #b4b9ba">为了你的账号安全，请用手机号登录</h5>

    </div>

    <van-form @submit="onSubmit">
      <van-field
        v-model="username"
        name="cellphone"
        label="中国+86　|"
        placeholder="请输入手机号"
      />
      <van-field
        v-model="password"
        type="password"
        name="密码"
        placeholder="请输入密码"
      />
      <div style="margin: 16px;">
        <van-button square block style="background: #30e084;color: white" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>


    <div style="width: 90%;margin: 0 auto;font-size: 10px;color: blue">
      <!-- 两端对齐 -->
      <van-row type="flex" justify="space-between">
        <van-col span="８" @click="toCodeLogin">验证码登录</van-col>
        <van-col span="9">
          <div style="float:left;" @click="toForgetPassword">忘记密码</div>
          <div style="float:left;margin-left: 11%;">|</div>
          <div style="float:right;" @click="toFreeRegister">免费注册</div>
        </van-col>
      </van-row>
    </div>


  </div>
</template>

<script>
  export default {
    name: "index",
    data() {
      return {
        username: '',
        password: '',
        value: ''
      }
    },
    methods: {
      onSubmit(values) {
        console.log('submit', values);
      },
      toCodeLogin() {
        this.$router.push('/mine/codelogin')
      },
      toFreeRegister() {
        this.$router.push('/mine/freeregister')
      },
      toForgetPassword() {
        this.$router.push('/mine/forgetpassword')
      }
    }
  }
</script>

<style scoped>
  .h2_tip, .h4_tip {
    margin: 20px 0 0 15px;
  }

</style>
