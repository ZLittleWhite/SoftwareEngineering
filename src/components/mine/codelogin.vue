<template>
  <div>

    <div class="tips" style="margin: 40px 0 15px 0">

      <h2 class="h2_tip">验证码登录</h2>
      <h5 class="h4_tip" style="color: #b4b9ba">为了你的账号安全，请用手机号登录</h5>

    </div>

    <van-form @submit="onSubmit">
      <van-field
        v-model="cellphone"
        name="手机号"
        label="中国+86　|"
        placeholder="请输入手机号"
      />
      <van-field
        v-model="code"
        name="密码"
        placeholder="请输入6位验证码"
      />
      <div>
        <span class="code">获取验证码</span>
      </div>
      <div style="margin: 16px;">
        <van-button square block style="background: #30e084;color: white" native-type="submit">
          登录
        </van-button>
      </div>
    </van-form>


    <div style="width: 90%;margin: 0 auto;font-size: 10px;color: blue">
      <!-- 两端对齐 -->
      <van-row type="flex" justify="space-between">
        <van-col span="８" @click="toPasswordLogin">密码登录</van-col>
        <van-col span="8">
          <div style="float:right;" @click="toFreeRegister">免费注册</div>
        </van-col>
      </van-row>
    </div>


  </div>
</template>

<script>
  export default {
    name: "codelogin",
    data() {
      return {
        cellphone: '',
        code: '',
      }
    },
    methods: {
      onSubmit(values) {
        console.log('submit', values);
      },
      toPasswordLogin(){
        this.$router.push('/mine')
      },
      toFreeRegister() {
        this.$router.push('/mine/freeregister')
      }
    }
  }
</script>

<style scoped>
  .h2_tip, .h4_tip {
    margin: 20px 0 0 15px;
  }
  .code{
    color: blue;
    cursor: pointer;
    float: right;
    position: fixed;
    right: 22px;
    top: 170px;
  }

</style>
