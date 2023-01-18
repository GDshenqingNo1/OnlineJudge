import { createRouter,createWebHashHistory } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import User from "../views/User.vue";
import ProblemList from "../views/ProblemList.vue";
import Rank from "../views/Rank.vue";
import Register from "../views/Register.vue";
import CreateProblem from "../views/CreateProblem.vue";


  const routes=[
    {
      path: '/',
      name: 'main',
      component: Home
    },
    {
      path:'/login',
      name:'login',
      component:Login
    },
    {
      path:'/register',
      name:'register',
      component:Register
    },
    {
      path:'/user',
      name:'user',
      component:User
    },
    {
      path:'/problem-list',
      name:'problemList',
      component:ProblemList
    },
    {
      path:'/rank',
      name:'rank',
      component:Rank
    },
    {
      path:'/create-problem',
      name:'createProblem',
      component:CreateProblem
    },
  ]
const Router=createRouter({
  history:createWebHashHistory(),
  routes
})
export default Router
