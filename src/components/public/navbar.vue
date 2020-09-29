<template>
  <div>
    <!-- 居中 -->
    <van-row class="rowsList" v-for="(item,index) in navList" :key="index">
      <b @click="toActive(index)" :class="index==activeIndex?'active':''">{{item.name}}</b>
      <van-icon :name="item.iconName"></van-icon>
    </van-row>

  </div>
</template>

<script>
  export default {
    name: "navbar",
    data() {
      return {
        //默认不显示
        activeIndex: -1,

        navList: [
          {
            name: "综合",
          },
          {
            name: "销量",
          }
          ,
          {
            name: "价格",
            iconName: "descending"
          }
          ,
          {
            name: "上新",
          }
        ]
      }

    },
    methods: {
      toActive(index) {
        console.log(index)
        this.activeIndex = index
      }
    }
  }
</script>

<style scoped>
  .rowsList {
    width: 25%;
    text-align: center;
    margin-top: 15px;
    float: left;
  }

  .colList {
    text-align: center;
    float: left;
  }

  .active {
    color: red;
  }
</style>
