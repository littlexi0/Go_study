import Vue from 'vue';
import VueRouter from 'vue-router';
// import store from '@/store';
import Home from '../views/Home.vue';



Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },

  {
    path:'/login',
    name:'login',
    component:()=>import('../views/login/Login.vue')
  },
  {
    path:'/register',
    name:'register',
    component:()=>import('../views/register/Register.vue')
  },
//   ...userRoutes,
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

// router.beforeEach((to, from, next) => {
//   if (to.meta.auth) { // 判断是否需要登录
//     // 判断用户是否登录
//     if (store.state.userModule.token) {
//       // 这里还要判断token的有效性，比如有没有过期，需要后台发放token的时候携带的有效期
//       // 如果token无效，需要请求token
//       next();
//     } else {
//       // 跳转登录
//       router.push({ name: 'login' });
//     }
//   } else {
//     next();
//   }
// });

export default router