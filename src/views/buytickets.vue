<template>
    <div>
        <buy-tickets></buy-tickets>
        <tab-bar></tab-bar>
    </div>
</template>

<script>
    import BuyTickets from "@/components/buytickets"
    import TabBar from "@/components/public/tabbar";
    export default {
        name: "buytickets",
        components:{
            "buy-tickets":BuyTickets,
            "tab-bar":TabBar
        }
    }
</script>

<style scoped>

</style>