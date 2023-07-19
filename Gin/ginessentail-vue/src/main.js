import Vue from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import Vuelidate from 'vuelidate';

import App from './App.vue';
import router from './router';
import axios from 'axios';
import VueAxios from 'vue-axios';

// scss style
import './assets/css/index.css';

Vue.config.productionTip = false;






Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.use(Vuelidate);
Vue.use(VueAxios, axios);


new Vue({
  router,
  // store,
  render: (h) => h(App),
}).$mount('#app');
