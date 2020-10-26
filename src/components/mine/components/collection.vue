<template>
    <div>
        <!--        头部-->
        <div>
            <van-nav-bar @click-left="back" @click-right="toShopCar" title="❤ 我的收藏 ❤" left-text="返回" left-arrow>
                <template #right>
                    <van-icon name="cart-circle-o" style="margin: 0 8px 0 0" size="20"/>
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
                        <van-icon @click="removeILove(isSelectShow.projectid)" style="margin: 0 0 0 220px" size="20px"
                                  name="like" color="#f11"/>
                    </template>
                </van-card>
            </div>
        </div>
        <!--        没有数据的显示样子-->
        <div v-if="getLoveList.length==0">
            <empty @click="reLoad"></empty>
        </div>
    </div>
</template>
<script>
    import Empty from "@/components/public/empty"
    import {Toast} from "vant"

    export default {
        name: "collection",
        components: {
            "empty": Empty
        },
        data() {
            return {
                //所有数据
                showAllList: [],
                //收藏数据
                getLoveList: [],
                getIds: [],

            }
        },
        methods: {
            toShopCar(){
              this.$router.push("/shopcar")
            },
            reLoad() {
                this.$router.push("/collection")
            },
            removeILove(remove_id) {
                //http://localhost:9090/api/removealove?s_id=626250108387
                let removeURL = "http://localhost:9090/api/removealove"
                this.$axios.get(removeURL, {
                    params: {
                        s_id: remove_id
                    }
                }).then(res => {
                    console.log(res.data.code);
                    if (res.data.code == 1) {
                        Toast("取消成功")
                        console.log("movedata", res.data.data)
                        this.getFavorateList()
                        this.$forceUpdate()
                    }
                }).catch(err => {
                    console.log(err);
                })
            },
            //获取收藏列表
            getFavorateList() {
                let _self = this
                let favorateURL = "http://localhost:9090/api/getlovelist"
                this.$axios.post(favorateURL)
                    .then(res => {
                        let tempIdList = []
                        //获取到喜欢的数组
                        tempIdList = res.data.data
                        console.log(tempIdList)
                        tempIdList.forEach(function (idItem) {
                            _self.getIds.push(idItem.SId)
                        })

                    }).catch(err => {
                    console.log(err);
                })
            },
            //获取所有的数据
            getAllData() {
                let _self = this
                let localDataURL = "internetData.json"
                this.$axios.get(localDataURL)
                    .then(res => {
                        _self.showAllList = res.data.data
                        //console.log(res.data.data)
                        let tempList = []
                        tempList = _self.showAllList
                        tempList.forEach(function (showItem) {
                            if (_self.getIds.indexOf(showItem.projectid) !== -1) {
                                _self.getLoveList.push(showItem)
                            }
                        })
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
            this.getAllData()
            this.getFavorateList()
        }
    }
</script>

<style scoped>
    .cardCtn {
        width: 100%;
        height: 150px;
        background: white;
    }
</style>