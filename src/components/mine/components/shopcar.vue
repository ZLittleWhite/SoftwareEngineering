<template>
    <div>
        <!--        头部-->
        <div>
            <van-nav-bar @click-left="back" title="❤ 购物车 ❤" left-text="返回"
                         left-arrow></van-nav-bar>
        </div>
        <!--        网络中生成的数据-->
        <div>
            <van-row>
                <van-col span="22">
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
                            <template #tag>
                                <van-checkbox-group v-model="result" ref="checkboxGroup">
                                    <van-checkbox
                                            ref="checkboxes"
                                            :name="index"
                                            style="margin: 50px 0 0 320px"
                                            @click="checkedArr()"
                                    ></van-checkbox>
                                </van-checkbox-group>
                            </template>
                        </van-card>
                    </div>
                </van-col>
            </van-row>

            <div>
                <van-action-sheet title="请输入密码" style="height: 400px" v-model="show">
                    <van-password-input
                            :value="value"
                            :focused="showKeyboard"
                            @focus="showKeyboard = true"
                    />
                    <van-row>
                        <van-col style="text-align: center;margin: 5px 0 0 0;color: gray;font-size: 12px" span="24">
                            总价：{{totalPrice}}
                        </van-col>
                    </van-row>
                    <!-- 数字键盘 -->
                    <van-number-keyboard
                            :show="showKeyboard"
                            @input="onInput"
                            @delete="onDelete"
                            @blur="showKeyboard = false"
                    />
                    <van-goods-action>
                        <van-goods-action-button @click="ConfirmBuy" color="#7232dd" type="danger" text="立即购买"/>
                    </van-goods-action>
                </van-action-sheet>
            </div>

            <div v-if="getLoveList.length!=0">
                <van-goods-action>
                    <van-goods-action-icon icon="chat-o" text="客服"/>
                    <van-goods-action-icon icon="shop-o" text="店铺"/>
                    <van-goods-action-button color="#be99ff" type="warning" :text="'总价：'+totalPrice"/>
                    <van-goods-action-button @click="toBuyNow" color="#7232dd" type="danger" text="立即购买"/>
                </van-goods-action>
            </div>
        </div>
        <!--        没有数据的显示样子-->
        <div v-if="getLoveList.length==0">
            <empty @click="reLoad"></empty>
        </div>
        <no-things></no-things>
    </div>
</template>
<script>
    import NoThings from "../../public/nothings";
    import Empty from "@/components/public/empty"
    import {Toast} from "vant"
    import qs from 'qs';

    export default {
        name: "shopcar",
        components: {
            "no-things": NoThings,
            "empty": Empty
        },
        data() {
            return {
                show: false,
                totalPrice: "0",
                checked: false,
                result: [],
                //所有数据
                showAllList: [],
                //收藏数据
                getLoveList: [],
                getIds: [],
                value: '',
                showKeyboard: true,
                //选中的数据
                checkedIdArr: []

            }
        },
        methods: {
            toDeleteChoosed() {
                console.log(this.checkedIdArr)
                let deleteParams = qs.stringify({
                    ids: this.checkedIdArr
                })
                console.log("params",deleteParams)

            },
            ConfirmBuy() {
                if (this.value.length == 6) {
                    Toast("购买成功")
                    this.$router.push("/mine")
                }
            },
            toBuyNow() {
                //this.toDeleteChoosed()
                if (this.totalPrice !== "0") {
                    this.value = ""
                    this.show = true
                } else {
                    Toast("请选择商品后支付！")
                }
            },
            checkedArr() {
                let _self = this
                let tprice = 0
                console.log(this.result);
                //被选中的id值
                this.checkedIdArr = this.result;
                this.result.forEach(function (item) {
                    //console.log(_self.getLoveList[item].price)
                    tprice += _self.getLoveList[item].price
                })

                //console.log("总价", tprice)
                this.totalPrice = tprice + ""
            },
            reLoad() {
                this.$router.push("/shopcar")
            },
            onInput(key) {
                this.value = (this.value + key).slice(0, 6);
            },
            onDelete() {
                this.value = this.value.slice(0, this.value.length - 1);
            },

            //获取购物车列表
            getShopCarList() {
                let _self = this
                let favorateURL = "http://localhost:9090/api/getshopcarlist"
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
        },
        created() {
            this.getAllData()
            this.getShopCarList()
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