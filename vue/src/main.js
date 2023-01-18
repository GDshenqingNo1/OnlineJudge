
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'

import {createApp} from 'vue'

import App from './App'
import router from './router'
import ElementPlus from "element-plus";
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import store from "@/store";

/* eslint-disable no-new */
createApp(App).use(router).use(ElementPlus,{
  locale: zhCn,
}).use(store).mount('#app')
