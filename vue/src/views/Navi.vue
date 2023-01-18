<template>

  <div>
    <el-menu  router class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <el-menu-item index="/problem-list">题单</el-menu-item>
      <el-menu-item index="/create-problem" >发布题目</el-menu-item>
      <el-menu-item index="/rank" >排行</el-menu-item>
      <el-menu-item index="/user">个人中心</el-menu-item>
      <el-button round @click="loginInfo">登录</el-button>
      <el-button round @click="registerInfo">注册</el-button>
    </el-menu>

  </div>

</template>


<script>

import register from "./Register.vue";

export default {

  name:"navi",
  computed: {
    register() {
      return register
    }
  },
  data() {
    return {
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    loginInfo(){
      this.$router.push("/login")
    },
    registerInfo(){
      this.$router.push("/register")
    }
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

