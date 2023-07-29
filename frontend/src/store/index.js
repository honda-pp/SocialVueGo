import { createStore } from 'vuex';

const store = createStore({
  state: {
    isLoggedIn: false,
    userID: null,
  },
  mutations: {
    setLoggedIn(state, loggedIn) {
      state.isLoggedIn = loggedIn;
    },
    setUserID(state, userID) {
      state.userID = userID;
    },
  },
  actions: {
    login({ commit }, userID) {
      commit('setLoggedIn', true);
      commit('setUserID', userID);
    },
    logout({ commit }) {
      commit('setLoggedIn', false);
      commit('setUserID', null);
    },
  },
  getters: {
    isLoggedIn: (state) => state.isLoggedIn,
    userID: (state) => state.userID,
  },
});

export default store;
