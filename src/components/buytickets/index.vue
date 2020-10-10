<template>
    <div>
        <div>
            <!-- 居中 -->
            <van-row type="flex" justify="center">
                <van-col style="text-align: center;margin: 10px 0 0 0" span="24">❤在线购票❤</van-col>
            </van-row>
            <van-divider dashed :style="{ color: '#1989fa', borderColor: '#f55546', padding: '0 16px' }">
            </van-divider>
        </div>

        <!--        网络中生成的数据-->
        <div>
            <van-tabs title-active-color="red" color="#fff" @click="onClick">
                <van-tab v-for="(item,index) in tabBarList" :title="item.title">
                    <div class="cardCtn" v-for="(isSelectShow,index) in selectedShowList" :key="isSelectShow.id">
                        <van-card
                                style="background: white"
                                :price="isSelectShow.price"
                                :desc="isSelectShow.venue"
                                :title="isSelectShow.name"
                                @click="toShowDetails(isSelectShow.id)"
                        >
                            <template #thumb>
                                <img style="height: 130%;width: 100%" :src="isSelectShow.verticalPic">
                            </template>
                            <template #tags>
                                <p style="padding: 25px 0 0 0">{{isSelectShow.showtime}}</p>
                                <van-tag plain type="danger">{{isSelectShow.categoryname}}</van-tag>
                            </template>
                            <template #footer>
                                <van-icon name="like-o"/>
                            </template>
                        </van-card>
                    </div>
                </van-tab>
            </van-tabs>
        </div>
        <!--        没有数据的显示样子-->
        <div v-if="selectedShowList.length==0">
            <van-empty description="暂无数据">
                <van-button plain round type="info" class="bottom-button">
                    刷新
                </van-button>
            </van-empty>
        </div>

        <!--        空白-->
        <div class="nulTab"></div>

    </div>
</template>

<script>
    import {Toast} from 'vant';

    export default {
        name: "index",
        data() {
            return {
                tabBarList: [
                    {
                        id: 0,
                        title: "全部"
                    },
                    {
                        id: 1,
                        title: "Livehouse"
                    },
                    {
                        id: 2,
                        title: "演唱会"
                    },
                    {
                        id: 3,
                        title: "体育"
                    },
                    {
                        id: 4,
                        title: "话剧歌剧"
                    },
                    {
                        id: 5,
                        title: "体育"
                    },
                    {
                        id: 6,
                        title: "曲艺杂坛"
                    },
                    {
                        id: 7,
                        title: "其他"
                    },

                ],
                showList: [],
                selectedShowList: []
            }
        },
        created() {
            this.getData()
        },

        methods: {
            toShowDetails(s_id) {
                console.log(s_id)
                this.$router.push({
                    path: '/ticketsdetails',
                    query: {
                        id: s_id
                    }
                })
            },

            onClick(name, title) {
                this.selectedShowList = []
                let _selt = this
                Toast(title);
                //console.log(name)
                this.showList.forEach(function (currentValue, index) {

                        if (currentValue.categoryid == (name - 1)) {
                            _selt.selectedShowList.push(currentValue)
                        }
                        if (name == 0) {
                            _selt.selectedShowList = _selt.showList
                        }
                    }
                )
                //console.log(this.selectedShowList)
            },
            getData() {
                let baseURL = "https://search.damai.cn/searchajax.html?keyword=&cty=&ctl=&sctl=&tsg=0&st=&et=&order=1&pageSize=30&currPage=1&tn="
                let localDataURL = "data.json"
                this.$axios.get(baseURL)
                    .then(res => {
                        // console.log("本地数据", res.data.data)
                        // this.showList = res.data.data
                        // this.selectedShowList = this.showList


                        //网络数据
                        console.log("网络数据", res.data.pageData.resultData)
                        let internetData = []
                        internetData = res.data.pageData.resultData
                        this.showList = internetData
                        this.selectedShowList = this.showList
                    })
                    .catch(err => {
                        console.log(err);
                    })
            }
        }
    }
</script>

<style scoped>
    .bottom-button {
        width: 160px;
        height: 40px;
        margin: 200px 0 0px 0;
    }

    .nulTab {
        width: 100%;
        height: 50px;
        /*background: red;*/
    }

</style>