import Vue from 'vue';
import Vuex from 'vuex';
import createLogger from 'vuex/dist/logger';

import configData from './modules/config-data';

Vue.use(Vuex);

const logger = createLogger({
  collapsed: true,
  logger: console
});

export default new Vuex.Store({
  modules: {configData},
  strict: true,
  /* TODO disable the logger plugin in production */
  plugins: [ logger ]
});
