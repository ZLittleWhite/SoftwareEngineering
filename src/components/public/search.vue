<template>
  <div>

    <div style="float:left;width: 85%">
      <!--    搜索框-->
      <van-search v-model="value" placeholder="请输入搜索关键词"/>
    </div>
    <div
      style="float: right;width: 15%;margin: 20px 0 0 0"
      @click="this.$router.go(-1)"
    >
      取消
    </div>

  </div>
</template>

<script>
  export default {
    name: "search",
    data() {
      return {
        value: "",
      }
    }
  }
</script>

<style scoped>

</style>
