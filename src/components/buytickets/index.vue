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
                <van-tab v-for="(item,index) in tabBarList" :title="item.title" :key="index">
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
            <empty></empty>
        </div>
        <!--        空白-->
        <no-things></no-things>

    </div>
</template>

<script>
    import Empty from "../public/empty"
    import NoThings from "../public/nothings";
    import {Toast} from 'vant';

    export default {
        name: "index",
        components: {
            "no-things": NoThings,
            "empty": Empty
        },
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
                        title: "音乐会"
                    },
                    {
                        id: 4,
                        title: "话剧歌剧"
                    },
                    {
                        id: 5,
                        title: "舞蹈芭蕾"
                    },
                    {
                        id: 6,
                        title: "曲艺杂坛"
                    },
                    {
                        id: 7,
                        title: "体育"
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
            },
            getData() {
                let localDataURL = "internetData.json"
                this.$axios.get(localDataURL)
                    .then(res => {
                        console.log("网络数据", res.data.data)
                        let internetData = []
                        internetData = res.data.data
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
    .nulTab {
        width: 100%;
        height: 50px;
        /*background: red;*/
    }

</style>