<template>
  <div>
    <app-index></app-index>
    <tab-bar></tab-bar>
  </div>
</template>

<script>

  import AppIndex from "@/components/index"
  import TabBar from "@/components/public/tabbar";

  export default {
    name: "index",
    components: {
      "app-index": AppIndex,
      "tab-bar":TabBar
    }
  }
</script>

<style scoped>

</style>
