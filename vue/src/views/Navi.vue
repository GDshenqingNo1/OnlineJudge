<template>

  <div>
    <el-menu  class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <el-menu-item index="1">题单</el-menu-item>
      <el-menu-item index="2" >发布题目</el-menu-item>
      <el-menu-item index="3" >排行</el-menu-item>
      <el-menu-item index="4">个人中心</el-menu-item>
      <el-button round>登录</el-button>
      <el-button round>注册</el-button>
    </el-menu>

  </div>

</template>


<script>

export default {
  name:"Navi",
  data() {
    return {
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },

    }

}
</script>
<style scoped>
.main-container{/*因为有两个相同的标签,所以要自定义标签*/
  height: 100%
}
.el-header{/*可以直接使用elementui的标签作为css选择器*/
  background-color: #188aaf;
  display: flex;
  justify-content: space-between;
  padding-left: 0;
  align-items: center;
  color: #fafafa;
  font-size:60px;
}
.el-aside{
  background-color: #f6f6f8;
}
.el-main{

  background-color: #eaedf1;
}
</style>

