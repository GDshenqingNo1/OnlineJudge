<template>
  <div>
    <el-card class="box-card">
      <el-form ref="loginForm" :model="form" :rules="rules" label-width="80px" class="login-box">
        <h3 class="login-title">登录</h3>
        <el-form-item label=" 题名" prop="name">
          <el-input type="text" placeholder="请输入题名" v-model="form.name"/>
        </el-form-item>
        <el-form-item label=" 描述" prop="description">
          <el-input type="text" placeholder=" 请输入描述" v-model="form.description"/>
        </el-form-item>
        <h3>示例</h3>>
        <el-form-item label=" 输入" prop="testIn">
          <el-input type="text" placeholder=" 请输入输入案例" v-model="form.testIn"/>
        </el-form-item>
        <el-form-item label=" 输出" prop="testOut">
          <el-input type="text" placeholder=" 请输入输出案例" v-model="form.testOut"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" v-on:click="onSubmit( 'loginForm' )">登录</el-button>
          <el-button plain type="primary" v-on:click="cancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "CreateProblem",
  data() {
    return {
      form: {
        name: '',
        description: '',
        testIn:'',
        testOut:''
      },
      //表单验证，需要在el-form-item 元素中增加prop 属性
      rules: {
        name: [
          {required: true, message: " 题名不可为空", trigger: 'blur'}
        ],
      description: [
          {required: true, message: " 问题描述不可为空 ", trigger: 'blur'}
        ],
        testIn: [
          {required: true, message: " 输入不可为空", trigger: 'blur'}
        ],
        testOut: [
          {required: true, message: " 输出不可为空", trigger: 'blur'}
        ],
      }
    }
  },
  methods: {
    onSubmit(formName) {

    },
    cancel(){
      this.$router.push("/")
    }
  }
}
</script>

<style scoped>
.box-card {
  width: 480px;margin: auto;
}
</style>

