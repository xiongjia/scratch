export default {
  namespaced: true,
  state: {
    settings: [{id: 1, name: 'test1'}, {id: 2, name: 'test2'}]
  },
  getters: {
    getItem: (state) => (id) => {
      console.log('get item...');
      return state.settings[id];
    },
    getData: (state) => 'test data'
  },
  actions: {
    loadConfig ({ commit }, { item }) {
      console.log('loading config', item);
      commit('setItems', [
        {id: 1, name: 'config-test1'},
        {id: 2, name: 'config-test2'}]);
    }
  },
  mutations: {
    setItems (state, items) {
      state.settings = items;
    }
  }
};
