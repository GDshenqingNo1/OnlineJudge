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
      prop="name"
      label="题名"
      width="180">
    </el-table-column>
    <el-table-column
      prop="author"
      label="作者"
      width="180">
    </el-table-column>
    <el-table-column
      prop="creat_time"
      label="创建时间"
      width="220">
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
  name :"ProblemList",
  components: {Navi}
}
</script>

