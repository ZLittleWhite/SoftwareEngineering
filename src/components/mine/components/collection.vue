<template>
    <div>
        <!--        头部-->
        <div>
            <van-nav-bar @click-left="back" @click-right="share" title="❤ 我的收藏 ❤" left-text="返回" left-arrow>
                <template #right>
                    <van-icon name="cart-circle-o" size="18"/>
                </template>
            </van-nav-bar>
        </div>
        <!--        网络中生成的数据-->
        <div>
            <div class="cardCtn" v-for="(isSelectShow,index) in getLoveList" :key="isSelectShow.id">
                <van-card
                        style="background: white"
                        :price="isSelectShow.price"
                        :desc="isSelectShow.venue"
                        :title="isSelectShow.name"

                >
                    <template #thumb>
                        <img style="height: 140%;width: 100%" :src="isSelectShow.verticalPic">
                    </template>
                    <template #tags>
                        <p style="padding: 20px 0 0 0">{{isSelectShow.showtime}}</p>
                        <van-tag plain type="danger">{{isSelectShow.categoryname}}</van-tag>
                    </template>
                    <template #desc>
                        <van-icon style="margin: 0 0 0 220px" size="25px" name="like" color="#f11"/>
                    </template>
                </van-card>
            </div>
        </div>
        <!--        没有数据的显示样子-->
        <div v-if="getLoveList.length==0">
            <van-empty description="暂无数据">
                <van-button plain round type="info" class="bottom-button">
                    刷新
                </van-button>
            </van-empty>
        </div>
    </div>
</template>

<script>
    export default {
        name: "collection",
        data() {
            return {
                //所有数据
                showAllList: [],

                //收藏数据
                getLoveList: [],
                getIds: []
            }
        },
        methods: {
            //获取收藏列表
            getFavorateList() {
                let _self = this
                let favorateURL = "http://localhost:9090/api/getlovelist"
                this.$axios.get(favorateURL)
                    .then(res => {
                        console.log(res.data.data);
                        let tempIdList = []
                        tempIdList = res.data.data
                        tempIdList.forEach(function (idItem, ii) {
                            _self.getIds.push(idItem.SId)
                        })
                    }).catch(err => {
                    console.log(err);
                })
            },

            //获取所有的数据
            getAllData() {
                let _self = this
                let baseURL = "https://search.damai.cn/searchajax.html?keyword=&cty=&ctl=&sctl=&tsg=0&st=&et=&order=1&pageSize=30&currPage=1&tn="
                this.$axios.get(baseURL)
                    .then(res => {
                        //网络数据
                        //console.log("网络数据", res.data.pageData.resultData)
                        _self.showAllList = res.data.pageData.resultData

                        let tempList = []
                        tempList = _self.showAllList
                        tempList.forEach(function (showItem) {
                            //console.log(showItem.projectid)
                            //console.log(_self.getIds);
                            if (_self.getIds.indexOf(showItem.projectid) != -1) {
                                _self.getLoveList.push(showItem)
                            }
                        })

                        console.log("Love", this.getLoveList)
                    })
                    .catch(err => {
                        console.log(err);
                    })
            },
            back() {
                this.$router.push('/mine')
            },
            share() {
                this.showShare = true
                console.log("share");
            },
        },

        created() {
            this.getFavorateList()
            this.getAllData()

        }
    }
</script>

<style scoped>
    .bottom-button {
        width: 160px;
        height: 40px;
        margin: 200px 0 0px 0;
    }
    .cardCtn{
        width: 100%;
        height: 150px;
        background: white;
    }
</style>