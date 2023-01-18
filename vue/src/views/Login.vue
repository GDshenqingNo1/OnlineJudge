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
        <el-form-item>
          <el-button type="primary" v-on:click="onSubmit( 'loginForm' )">登录</el-button>
          <el-button plain type="primary" v-on:click="cancel">取消</el-button>
        </el-form-item>
      </el-form>

      <el-dialog
        title="温馨提示"
        :visible.sync="dialogVisible"
        width="30%"
        :before-close="handLeClose">
        <span>请输入账号和密码</span>
        <span slot="footer" class="dialog- footer">
        <el-button type="primary" @click="dialogVisible = false">确定</el-button>
      </span>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
import api from "../api/api";
import {ElMessage} from "element-plus";
import store from "../store";
export default {
  name: "Login",
  data() {
    return {
      form: {
        username: '',
        password: ''
      },
      //表单验证，需要在el-form-item 元素中增加prop 属性
      rules: {
        username: [
          {required: true, message: " 账号不可为空", trigger: 'blur'}
        ],
        password: [
          {required: true, message: " 密码不可为空 ", trigger: 'blur'}
        ]
      },
//对话框显示和隐藏
      dialogVisible: false
    }
  },
  methods: {
    onSubmit() {
         api.login(JSON.stringify(this.form)).then(res=>{
           if (res.data.code == 200) {
             ElMessage.success('登录成功')
             localStorage.setItem("token", res.data.data.token);
             store.commit("loginSucc", res.data.data.token);
             store.commit("setUser", {username: this.form.username, is_admin: res.data.data.is_admin});
             localStorage.setItem('is_admin', res.data.data.is_admin)
             localStorage.setItem('username', this.form.username)
           } else {
             ElMessage.error(res.data.msg)
           }
         })
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

