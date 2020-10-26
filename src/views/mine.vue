<template>
    <div>
        <app-mine></app-mine>
        <tab-bar></tab-bar>
    </div>
</template>

<script>
    import AppMine from "@/components/mine/index"
    import TabBar from "@/components/public/tabbar";

    export default {
        name: "mine",
        components: {
            "app-mine": AppMine,
            "tab-bar":TabBar
        }
    }
</script>

<style scoped>

</style>