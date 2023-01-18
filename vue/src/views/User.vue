<template>
  <Navi></Navi>
  <el-descriptions class="margin-top" title="用户信息" :column="2"  border>
    <el-descriptions-item label="用户名">
lwj
</el-descriptions-item>
    <el-descriptions-item label="正确次数">
liangweijian666@outlook.com
</el-descriptions-item>
    <el-descriptions-item label="提交次数">
0
</el-descriptions-item>
<el-descriptions-item label="正确次数">
0
</el-descriptions-item>
</el-descriptions>
  </template>

<script>
import Navi from "./Navi.vue";

export default {
  name:"User",
  components: {Navi}
}

</script>
