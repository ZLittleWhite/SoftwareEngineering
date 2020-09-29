<template>
  <div>

    <van-tabbar v-model="active" style="font-size: 24px;color: black" route>
      <van-tabbar-item url="/" icon="wap-home-o"></van-tabbar-item>
      <van-tabbar-item url="/"><b style="font-size: 14px;">首页</b></van-tabbar-item>
      <van-tabbar-item to="/mine"><b style="font-size: 14px;">个人中心</b></van-tabbar-item>
      <van-tabbar-item to="/product"><b style="font-size: 14px;">产品分类</b></van-tabbar-item>
    </van-tabbar>

  </div>
</template>

<script>
  export default {
    name: "tabbar",
    data() {
      return {
        active: -1
      }
    }
  }
</script>

<style scoped>

</style>
