import Vue from 'vue';

import VueLogger from 'vuejs-logger';
import BootstrapVue from 'bootstrap-vue';

import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import App from './app';
import router from './router';

import Pin from './common/pin.js';
Vue.use(Pin);

Vue.config.productionTip = false;
Vue.use(VueLogger, {
  logLevel: 'debug',
  stringifyArguments: false,
  showLogLevel: true,
  showMethodName: true,
  separator: '|',
  showConsoleColors: true
});
Vue.use(BootstrapVue);

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App />'
});
