<template>
    <div>
        {{msg}}
    </div>
</template>

<!--加上scoped能够防止样式之间的冲突-->
<style scoped>
    body {
        background-color: #ff0000;
    }
</style>

<script>
export default {
  data () {
    return { msg: '这个是Vue模板页' }
  }
}
</script>

