<template>
    <div>
        <!--        详细信息-->
        <div class="ticketdetails">
            <!--            头部分享个朋友-->
            <div>
                <van-nav-bar @click-left="back" @click-right="share" title="" left-text="返回" left-arrow>
                    <template #right>
                        <van-icon name="share" size="18"/>
                    </template>
                </van-nav-bar>
                <van-share-sheet
                        v-model="showShare"
                        title="立即分享给好友"
                        :options="options"
                        @select="onSelect"
                />
            </div>
            <!--            主体中间部-->
            <div>
                <div class="cardCtn" v-for="(show,index) in chooseTicket" :key="show.id">
                    <!--            详细信息-->
                    <van-card
                            :title="show.name"
                            style="background: #ffffff;height: 200px"
                    >
                        <template #title>
                            <h2 style="margin-left: 20%">【{{show.cityname}}】{{show.name}}</h2>

                        </template>
                        <template #thumb>
                            <img style="height: 200%;width: 150%" :src="show.verticalPic">
                        </template>
                        <template #tags>
                            <p style="padding: 75px 0 10px 0;margin-left: 20%" v-limit="12">{{show.actors}}</p>
                            <van-tag style="margin-left: 20%" plain type="danger">{{show.showstatus}}</van-tag>
                            <van-tag style="margin-left: 1%" plain type="danger">￥{{show.price}}</van-tag>
                        </template>
                        <template #footer>
                            <van-tag round type="primary" @click="addToMyFavorate(show.projectid,show.categoryname)">
                                收藏
                            </van-tag>
                        </template>
                    </van-card>
                    <!--        细胞单元，不同的功能-->
                    <div>
                        <div>
                            <van-notice-bar left-icon="volume-o" :scrollable="false">
                                <van-swipe
                                        vertical
                                        class="notice-swipe"
                                        :autoplay="3000"
                                        :show-indicators="false"
                                >
                                    <van-swipe-item>价格可能存在一定的波动 ~~</van-swipe-item>
                                    <van-swipe-item>目前最低价格：{{show.price}} 元</van-swipe-item>
                                    <van-swipe-item>目前最高价格：{{show.pricehigh}} 元</van-swipe-item>
                                </van-swipe>
                            </van-notice-bar>
                        </div>
                    </div>
                    <div>
                        <van-cell is-link :title="show.showtime" :label="show.venue"/>
                        <van-cell is-link :title="show.nameNoHtml"/>
                        <!--                        <van-cell is-link :title="'类型：'+show.categoryname" :label="'表演者：'+show.categoryname"/>-->

                        <van-cell>
                            <h4 style="color: #efa7be;padding: 0 0 10px 0">我的评分</h4>
                            <van-rate v-model="value" icon="like" void-icon="like-o"/>
                        </van-cell>
                        <van-cell>
                            <template #title>
                                <h4>网友评价</h4>
                                <van-tag class="mright" type="danger">点赞</van-tag>
                                <van-tag class="mright" type="danger">踩他</van-tag>
                                <van-tag class="mright" type="danger">下次一定</van-tag>
                            </template>
                        </van-cell>
                        <van-grid :border="false">
                            <van-grid-item icon="photo-o" text="喜欢"/>
                            <van-grid-item icon="photo-o" text="超级好看"/>
                            <van-grid-item icon="photo-o" text="下次一定"/>
                            <van-grid-item icon="photo-o" text="前方高能"/>
                        </van-grid>
                    </div>
                </div>
            </div>
        </div>
        <!--        底部购买商品-->
        <div>
            <van-goods-action>
                <van-goods-action-icon icon="service-o" text="帮助" @click="onClickIcon"/>
                <van-goods-action-icon icon="like-o" text="想看" @click="onClickIcon"/>
                <van-goods-action-button
                        type="danger"
                        text="立即购买"
                        @click="onClickButton"
                />
            </van-goods-action>
        </div>
        <!--        跳转到购买界面-->
        <div>
            <!-- 自定义图标 -->
            <van-popup
                    v-model="toBuyShow"
                    round
                    closeable
                    close-icon="close"
                    position="bottom"
                    :style="{ height: '30%' }"
            >
                <van-submit-bar :price="ticketPrice" button-text="提交订单" @submit="onSubmit">
                    <template #tip>
                        你的收货地址不支持同城送, <span @click="onClickEditAddress">修改地址</span>
                    </template>
                    <template #top>
                        <van-button plain block style="color: orange" @click="addToShopCar">加入购物车</van-button>
                    </template>
                </van-submit-bar>
            </van-popup>
        </div>
        <!--        跳转到支付界面-->
        <div>
            <!-- 自定义图标 -->
            <van-popup
                    v-model="pay"
                    round
                    closeable
                    close-icon="close"
                    position="bottom"
                    :style="{ height: '60%' }"
            >
                <van-submit-bar button-text="确认密码" button-type="info" @submit="confirmPay">
                    <template #top>
                        <div>
                            <!-- 密码输入框 -->
                            <van-password-input

                                    style="margin: 0 0 220px 0"
                                    :value="pwdvalue"
                                    :focused="showKeyboard"
                                    @focus="showKeyboard = true"
                            />
                            <!-- 数字键盘 -->
                            <van-number-keyboard
                                    :show="showKeyboard"
                                    @input="onInput"
                                    @delete="onDelete"
                                    @blur="showKeyboard = false"
                            />
                        </div>
                    </template>
                </van-submit-bar>
            </van-popup>
        </div>
    </div>
</template>

<script>

    import {Toast} from 'vant';

    export default {
        name: "index",
        components: {},
        data() {
            return {
                pwdvalue: '',
                showKeyboard: true,
                pay: false,
                ticketPrice: 0,
                toBuyShow: false,
                value: 3,
                chooseId: -1,
                chooseTicket: [],
                showShare: false,
                options: [
                    {name: '微信', icon: 'wechat'},
                    {name: '微博', icon: 'weibo'},
                    {name: '复制链接', icon: 'link'},
                    {name: '二维码', icon: 'qrcode'},
                ],
            }
        },
        created() {
            this.getData()
            this.chooseId = this.$route.query.id
            console.log("chooseId", this.chooseId)

        },
        methods: {
            onDelete() {
                this.pwdvalue = this.pwdvalue.slice(0, this.pwdvalue.length - 1);
            },
            addToShopCar() {
                console.log(this.chooseTicket[0].projectid)
                let SId = this.chooseTicket[0].projectid
                //addToShopCarURL  http://localhost:9090/api/addtoshopcar?sid=11114
                let addToShopCarURL = "http://localhost:9090/api/addtoshopcar"
                this.$axios.get(addToShopCarURL, {
                    params: {
                        sid: SId
                    }
                }).then(res => {
                    console.log(res);
                    if (res.data.code === 1) {
                        Toast("已成功添加到购物车")
                    } else if (res.data.code === 0) {
                        Toast("请勿重复添加")
                    }
                }).catch(err => {
                    console.log(err);
                })
            },
            onClickEditAddress() {
            },
            onSubmit() {
                this.pay = true;
            },
            onClickIcon() {
                Toast('点击图标');
            },
            onClickButton() {
                this.toBuyShow = true;
            },
            onSelect(option) {
                Toast(option.name);
                this.showShare = false;
            },
            back() {
                this.$router.push('/buytickets')
            },
            share() {
                this.showShare = true
                console.log("share");
            },
            confirmPay() {
                if (this.pwdvalue.length === 6) {
                    console.log("pay");
                    Toast("支付成功");
                    this.$router.push("/buytickets")
                }
            },
            onInput(key) {
                this.pwdvalue = (this.pwdvalue + key).slice(0, 6);
                console.log(this.pwdvalue)
            },
            addToMyFavorate(id, name) {
                //添加请求地址 -- http://localhost:9090/api/add?sid=1154&cname=zhadsaoudi
                let addURL = "http://localhost:9090/api/add"
                console.log(id, name)
                id = id.toString()
                //添加数据库
                this.$axios.get(addURL, {
                    params: {
                        sid: id,
                        cname: name
                    }
                }).then(res => {
                    console.log(res);
                    if (res.data.code === 1) {
                        Toast("收藏成功");
                    }
                }).then(err => {
                    console.log(err);
                })
            },
            getData() {
                let _self = this
                let localDataURL = "internetData.json"
                this.$axios.get(localDataURL)
                    .then(res => {
                        //console.log("网络数据", res.data.data)
                        let internetData = []
                        internetData = res.data.data
                        internetData.forEach(function (item, index) {
                            if (item.id == _self.chooseId) {
                                _self.chooseTicket.push(item)
                                _self.ticketPrice = item.price * 100;
                            }
                        })
                    })
                    .catch(err => {
                        console.log(err);
                    })
            }
        }
    }
</script>

<style scoped>
    .notice-swipe {
        height: 40px;
        line-height: 40px;
    }

    .ticketdetails {
        width: 100%;
        height: 100%;
        margin-bottom: 150px;
        /*background: red;*/
    }

    .mright {
        margin-right: 5px;
    }
</style>