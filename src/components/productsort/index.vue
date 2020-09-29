<template>
  <div>
    <div style="margin-top: 20px;width: 100%; text-align: center">
      全部商品
    </div>
    <!--    搜索框-->
    <van-search id="search" v-model="value" placeholder="请输入搜索关键词"/>

    <nav-bar></nav-bar>

  </div>
</template>

<script>

  import NavBar from "@/components/public/navbar"

  export default {
    name: "index",
    components: {
      "nav-bar": NavBar
    },
    data() {
      return {
        value: ""
      }
    },
    // beforeCreate() {
    //   document.getElementById("search").blur()
    // }
  }
</script>

<style scoped>

</style>
