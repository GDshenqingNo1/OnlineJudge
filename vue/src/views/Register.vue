<template>
  <div>
    <el-card class="box-card">
      <el-form ref="loginForm" :model="form" :rules="rules" label-width="80px" class="login-box">
        <h3 class="login-title">登录</h3>
        <el-form-item label=" 用户名" prop="username">
          <el-input type="text" placeholder="请输入用户名" v-model="form.username"/>
        </el-form-item>
        <el-form-item label=" 密码" prop="password">
          <el-input type="password" placeholder=" 请输入密码" v-model="form.password"/>
        </el-form-item>
        <el-form-item label=" 邮箱" prop="mail">
          <el-input style="width: 100px" type="text" placeholder=" 请输入邮箱" v-model="form.mail"/>
          <el-button type="primary" >发送验证码</el-button>
        </el-form-item>
        <el-form-item label=" 验证码" prop="code">
          <el-input type="password" placeholder=" 请输入验证码" v-model="form.code"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" v-on:click="onSubmit">登录</el-button>
          <el-button plain type="primary" v-on:click="cancel">取消</el-button>
        </el-form-item>
      </el-form>


    </el-card>
  </div>
</template>

<script>
export default {
  name: "Register",
  data() {
    return {
      form: {
        username: '',
        password: '',
        mail:'',
        code:''
      },
      //表单验证，需要在el-form-item 元素中增加prop 属性
      rules: {
        username: [
          {required: true, message: " 账号不可为空", trigger: 'blur'}
        ],
        password: [
          {required: true, message: " 密码不可为空 ", trigger: 'blur'}
        ],
        mail: [
          {required: true, message: " 邮箱不可为空 ", trigger: 'blur'}
        ],
        code: [
          {required: true, message: " 验证码不可为空 ", trigger: 'blur'}
        ]
      },
//对话框显示和隐藏
      dialogVisible: false
    }
  },
  methods: {

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
