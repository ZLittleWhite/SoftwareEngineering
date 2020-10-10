<template>
    <div>
        <van-tabbar v-model="active" route>
            <van-tabbar-item to="/" icon="shop">首页</van-tabbar-item>
            <van-tabbar-item to="/buytickets" icon="cart">购票</van-tabbar-item>
            <van-tabbar-item to="/mine" icon="manager">我的</van-tabbar-item>
        </van-tabbar>
    </div>
</template>

<script>
    export default {
        name: "tabbar",
        data() {
            return {
                active: 0,
            }
        }
    }
</script>

<style scoped>

</style>