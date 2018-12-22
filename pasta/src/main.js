import Vue from 'vue';
import App from './app';
import router from './router';
import VueLogger from 'vuejs-logger';

Vue.config.productionTip = false;
Vue.use(VueLogger, {
  logLevel: 'debug',
  stringifyArguments: false,
  showLogLevel: true,
  showMethodName: true,
  separator: '|',
  showConsoleColors: true
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App />'
});
