<template>
    <div>
        <!--        头部-->
        <div>
            <van-row>
                <van-col class="center" span="18"></van-col>
                <van-col class="center" span="6">
                    <van-icon size="20px" class="icons" name="scan"/>
                    <van-icon @click="toRegister" size="20px" class="icons" name="user-o"/>
                </van-col>
            </van-row>
            <van-row>
                <!--                头像-->
                <van-col class="center" span="6">
                    <van-image
                            @click="toLogin"
                            round
                            width="4rem"
                            height="4rem"
                            src="https://img.yzcdn.cn/vant/cat.jpeg"
                    />
                </van-col>
                <van-col span="18">
                    <van-cell-group :border="false">
                        <van-cell>
                            <template #title>
                                <h4>{{name}} </h4>
                            </template>
                            <template #label>
                                <h4>go shopping ~</h4>
                            </template>
                        </van-cell>
                    </van-cell-group>
                </van-col>
                <van-col span="24">
                    <van-cell-group>
                        <van-cell class="vcSpecial" is-link icon="shopping-cart-o">
                            <template #title>
                                <h4 class="mleft" @click="toShopCar">购物车</h4>
                            </template>
                        </van-cell>
                        <div v-for="item in mineList" :key="item.id">
                            <van-cell is-link :icon="item.icon" @click="toMyPath(item.path)">
                                <template #title>
                                    <h4 class="mleft">{{item.title}}</h4>
                                </template>
                            </van-cell>
                        </div>
                    </van-cell-group>
                </van-col>
            </van-row>
        </div>
    </div>
</template>

<script>
    export default {
        name: "index",
        data() {
            return {
                name: "",
                mineList: [
                    {id: 0, title: "已付款", icon: "paid", path: "/"},
                    {id: 1, title: "卡包", icon: "balance-o", path: "/buytickets"},
                    {id: 2, title: "收藏", icon: "goods-collect-o", path: "/collection"},
                    {id: 3, title: "设置", icon: "setting-o", path: "/ticketsdetails"},
                    {id: 4, title: "钱包", icon: "refund-o", path: "/"},
                ],
            }
        },
        methods: {
            toShopCar() {
                this.$router.push("/shopcar")
            },
            toRegister() {
                this.$router.push("/register")
            },
            toLogin() {
                this.$router.push("/login")
            },
            toMyPath(path) {
                console.log(path);
                this.$router.push(path)
            },
            getUsernameFromSession() {
                let getUsername = localStorage.getItem("username")
                if (getUsername === undefined || getUsername === null) {
                    this.name = "" + "HI~ 早上好"
                } else {
                    getUsername = getUsername.toUpperCase()
                    this.name = "Vip : " + "[ " + getUsername + " ]" + " 已登录 ";
                }
            }
        },
        created() {
            this.getUsernameFromSession()
        }
    }
</script>

<style scoped>
    .icons {
        margin: 15px 15px 15px 0px;
    }

    .center {
        text-align: center;
    }

    .mleft {
        margin-left: 10px;
    }

    .vcSpecial {
        margin: 20px 0 20px 0;
    }
</style>