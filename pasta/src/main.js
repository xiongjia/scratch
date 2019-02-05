import Vue from 'vue';

import VueLogger from 'vuejs-logger';
import BootstrapVue from 'bootstrap-vue';
import iView from 'iview';

import 'iview/dist/styles/iview.css';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import App from './app';
import router from './router';
import store from './store';

/* TODO loading vue logger config from build arguments */
Vue.use(VueLogger, {
  logLevel: 'debug',
  stringifyArguments: false,
  showLogLevel: true,
  showMethodName: true,
  separator: '|',
  showConsoleColors: true
});
Vue.use(BootstrapVue);
Vue.use(iView);

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  components: { App },
  template: '<App />'
});
