import http from "./http";
export default {
  //获取题单
  getProblemList(param){
    return http.get(`/problem/list`,param)
  },
  //获取问题详情
  getProblemInfo(param,id){
    return http.get(`problem/info?id=${id}`,param)
   },
   //注册
  register(param){
    return http.postJson(`/user/register`,param)
  },
  //登录
  login(param){
    return http.postJson(`/user/register`,param)
  },
  //提交代码
  submit(param,id) {
    return http.postJson(`/submit?problem_id=${id}`,param)
  },
  //修改题目
  modifyProblem(param,id){
    return http.postJson(`/problem/modify?id=${id}`,param)
  },
  //发布题目
  createProblem(param){
    return http.postJson(`/problem/create`,param)
  }
}
