<template>
    <div>
        <!--        轮播图-->
        <div class="Swipe">
            <van-swipe class="my-swipe" :autoplay="3000" indicator-color="white">
                <van-swipe-item>
                    <van-image class="imgCtn"
                               :src="require('@/assets/images/index/swipe/swipe_0.png')"></van-image>
                </van-swipe-item>
                <van-swipe-item>
                    <van-image class="imgCtn"
                               :src="require('@/assets/images/index/swipe/swipe_1.png')"></van-image>
                </van-swipe-item>
                <van-swipe-item>
                    <van-image class="imgCtn"
                               :src="require('@/assets/images/index/swipe/swipe_2.png')"></van-image>
                </van-swipe-item>
                <van-swipe-item>
                    <van-image class="imgCtn"
                               :src="require('@/assets/images/index/swipe/swipe_3.png')"></van-image>
                </van-swipe-item>
            </van-swipe>
        </div>
        <!--        宫格-->
        <div class="Grid">
            <van-grid>
                <van-grid-item icon="photo-o" text="榜单广场"/>
                <van-grid-item icon="photo-o" text="演出日历"/>
                <van-grid-item icon="photo-o" text="折扣专区"/>
                <van-grid-item icon="photo-o" text="会员专区"/>
                <van-grid-item icon="photo-o" text="榜单广场"/>
                <van-grid-item icon="photo-o" text="演出日历"/>
                <van-grid-item icon="photo-o" text="折扣专区"/>
                <van-grid-item icon="photo-o" text="会员专区"/>
            </van-grid>
        </div>

        <!--        特别推荐-->
        <div>
            <van-row>
                <van-col span="8" style="font-size: 20px;text-align: center;margin: 20px 0 10px 0">
                    <b>必看排行榜</b>
                </van-col>
                <van-col span="12"></van-col>
            </van-row>
        </div>

        <!--        滚动推荐-->
        <div>
            <van-row gutter="10">
                <van-col style="text-align: center" span="8" v-for="(item,index) in showList">
                    <div class="imgTitleCtn">
                        <img class="AImg" :src="item.verticalPic">
                        <div style="font-size: 14px;" v-limit="15">{{item.name}}</div>
                        <span>
                    <b style="color: red">
                        ￥{{item.price}}
                    </b>
                    <b style="font-weight: lighter;font-size: 10px;margin-left: 5px">
                        起
                    </b></span>
                    </div>
                </van-col>
            </van-row>
        </div>

        <!--        空组件-->
        <no-things></no-things>
    </div>
</template>
<script>
    import NoThings from "../../public/nothings";

    export default {
        name: "Tab_1",
        components: {
            "no-things": NoThings
        },
        data() {
            return {
                images: [
                    'https://img.yzcdn.cn/vant/apple-1.jpg',
                    'https://img.yzcdn.cn/vant/apple-2.jpg',
                ],
                showList: []
            }
        },
        methods: {
            getData() {
                let localDataURL = "internetData.json"
                this.$axios.get(localDataURL)
                    .then(res => {
                        console.log("网络数据", res.data.data)
                        this.showList = res.data.data
                        let tempArray = []
                        tempArray = res.data.data
                        this.showList = tempArray.slice(5, 13)
                    })
                    .catch(err => {
                        console.log(err);
                    })
            }
        },
        created() {
            this.getData()
        }
    }
</script>

<style scoped>
    .Swipe {
        margin: 10px 0 0 0;
        width: 100%;
        overflow: hidden;
        height: 100px;
    }

    .Grid {
        width: 100%;
    }

    .my-swipe .van-swipe-item {
        color: #fff;
        font-size: 20px;
        line-height: 0px;
        text-align: center;
        background-color: #39a9ed;
    }

    .imgCtn {
        width: 100%;
    }

    .imgTitleCtn {
        margin: 0px 0 30px 0;
    }

    .AImg {
        border: 1px #d4d3d3 solid;
        border-radius: 5px;
        height: 150px;
        overflow: hidden;
        width: 80%;
    }

    .VanRowCtn {

    }
</style>