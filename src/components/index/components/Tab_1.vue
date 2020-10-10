<template>
    <div>
        <!--        轮播图-->
        <div class="Swipe">
            <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white">
                <van-swipe-item>
                    <van-image style="width: 100%;height: 200px" :src="require('@/assets/images/index/swipe/swipe_0.png')"></van-image>
                </van-swipe-item>
                <van-swipe-item>
                    <van-image style="width: 100%;height: 200px" :src="require('@/assets/images/index/swipe/swipe_1.png')"></van-image>
                </van-swipe-item>
                <van-swipe-item>
                    <van-image style="width: 100%;height: 200px" :src="require('@/assets/images/index/swipe/swipe_2.png')"></van-image>
                </van-swipe-item>
                <van-swipe-item>
                    <van-image style="width: 100%;height: 200px" :src="require('@/assets/images/index/swipe/swipe_3.png')"></van-image>
                </van-swipe-item>
            </van-swipe>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Tab_1",
        data() {
            return {
                images: [
                    'https://img.yzcdn.cn/vant/apple-1.jpg',
                    'https://img.yzcdn.cn/vant/apple-2.jpg',
                ],
            }
        }
    }
</script>

<style scoped>
    .Swipe {
        margin: 10px 0 0 0;
        width: 100%;
        overflow: hidden;
        height: 200px;
        float: left;
    }

    .my-swipe .van-swipe-item {
        color: #fff;
        font-size: 20px;
        line-height: 150px;
        text-align: center;
        background-color: #39a9ed;
    }
</style>