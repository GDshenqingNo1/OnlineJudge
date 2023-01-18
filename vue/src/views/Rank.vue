<template>
  <Navi></Navi>
  <el-table
    :data="tableData"
    stripe
    style="width: 100%">
    <el-table-column
      prop="id"
      label="id"
      width="180">
    </el-table-column>
    <el-table-column
      prop="username"
      label="用户名"
      width="180">
    </el-table-column>
    <el-table-column
      prop="submit_time"
      label="提交次数"
      width="180">
    </el-table-column>
    <el-table-column
      prop="accept_time"
      label="正确次数"
      width="180">
    </el-table-column>
  </el-table>
</template>
<script>
import Navi from "./Navi.vue";
export default {
  name :"Rank",
  components: {Navi}
}
</script>

