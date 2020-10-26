<template>
    <div>
        <van-empty description="暂无数据">
            <van-button plain round type="info" class="bottom-button">
                刷新
            </van-button>
        </van-empty>
    </div>
</template>

<script>
    export default {
        name: "empty"
    }
</script>

<style scoped>
    .bottom-button {
        width: 160px;
        height: 40px;
        margin: 200px 0 0px 0;
    }
</style>