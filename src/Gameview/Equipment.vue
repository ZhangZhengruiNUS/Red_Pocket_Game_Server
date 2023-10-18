<template>
    <n-card title="带封面的卡片">
    <template #cover>
      <img src="https://static.vecteezy.com/system/resources/previews/015/281/932/original/cannon-icon-cartoon-style-vector.jpg">
    </template>
    技能描述 
    <n-checkbox size="small" label="带上装备" />
    </n-card>

    <n-card title="带封面的卡片">
    <template #cover>
      <img src="https://s3-eu-central-1.amazonaws.com/cartoons-s3/styles/product_detail_image/s3/tsar%20cannon.jpg?itok=GxjsPcuQ">
    </template>
    技能描述 
    <n-checkbox size="small" label="带上装备" />
    </n-card>

  <router-link class="router_button" to="/GameInterface">GO!</router-link>
</template>

<script>
    let GameInterface = "/GameInterface"
</script>

<style scoped>
.n-card {
    max-width: 300px;
}
.router_button{
    height: 60px;
    width: 400px;
    line-height: 36px;
    font-weight: bold;
    font-size: 40px;
    background: #4df719;
    color: #2c3e50;
    border: 1px solid transparent;
    padding: 0 10px; margin: 20px;
}

</style>