<template>
  <div>

    <!--    轮播图-->
    <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white">
      <van-swipe-item>1</van-swipe-item>
      <van-swipe-item>2</van-swipe-item>
      <van-swipe-item>3</van-swipe-item>
      <van-swipe-item>4</van-swipe-item>
    </van-swipe>

    <!--    搜索框-->
    <van-search v-model="value" placeholder="请输入搜索关键词"/>

    <!--    宫格-->


    <van-grid>
      <div v-for="(item,index) in gridList" style="width: 25%">
        <router-link :to="{path:'detail?id='+index,query:{text:item.text}}">
          <van-grid-item :icon="item.icon" :text="item.text"/>
        </router-link>
      </div>
    </van-grid>


  </div>
</template>

<script>
  export default {
    name: "index",
    data() {
      return {
        value: "",
        gridList: [
          {
            icon: "photo-o",
            text: "我是周敌"
          },
          {
            icon: "photo-o",
            text: "亲爱的阿亨"
          },
          {
            icon: "photo-o",
            text: "文字111"
          },
          {
            icon: "photo-o",
            text: "文字111"
          },
          {
            icon: "photo-o",
            text: "文字"
          }
        ]
      }
    }
  }
</script>

<style scoped>
  .my-swipe .van-swipe-item {
    color: #fff;
    font-size: 20px;
    line-height: 200px;
    text-align: center;
    background-color: #39a9ed;
  }
</style>
