<template>
    <div>
        <!--        头部搜索search-->
        <div class="indexTop">
            <van-row>
                <van-col span="5" style="text-align: center">
                    <div class="mtop">
                        <b style="font-size: 15px;">景德镇</b>
                        <van-icon name="arrow-down"/>
                    </div>
                </van-col>
                <van-col span="17" style="text-align: center">
                    <van-search v-model="value" placeholder="请输入搜索关键词"/>
                </van-col>
                <van-col span="2" style="text-align: center">
                    <van-icon class="mtop" size="20px" name="chat-o"/>
                </van-col>
            </van-row>
        </div>
        <!--        Tab-->
        <div class="indexTab">
            <van-tabs v-model="active">
                <van-tab title="❤ 特别推荐">
                    <tab-fir></tab-fir>
                </van-tab>
                <van-tab title="Livehouse">内容 2</van-tab>
                <van-tab title="音乐节">内容 3</van-tab>
                <van-tab title="话剧音乐剧">内容 4</van-tab>
            </van-tabs>
        </div>

    </div>
</template>

<script>
    import TabFir from "./components/Tab_1"
    import {Grid, GridItem} from 'vant';

    export default {
        name: "index",
        components: {
            "tab-fir": TabFir
        },
        data() {
            return {
                value: "",
                active: 0,
            }
        },
        created() {

        }
    }
</script>

<style scoped>
    .indexTop {
        float: left;
        width: 100%;
        height: 8vh;
    }

    .indexTab {
        float: left;
        width: 100%;
        height: 30px;
    }

    .mtop {
        margin: 15px 0 0 0;
    }

</style>
