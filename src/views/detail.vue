<template>
  <div>

    <van-row type="flex" justify="center">
      <van-col span="6" style="text-align: center;margin: 20px 0 10px 0"> {{title}}</van-col>
    </van-row>

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
        title: ""
      }
    },

    created() {
      console.log(this.$route.query.text)
      this.title = this.$route.query.text
    }

  }
</script>

<style scoped>

</style>
