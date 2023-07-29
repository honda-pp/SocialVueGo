import { createStore } from 'vuex';
import { loginUser, logoutUser, isLoggedIn } from '../api/userApi';

export default createStore({
  state: {
    isLoggedIn: false,
    userID: null,
  },
  mutations: {
    SET_LOGIN_STATUS(state, status) {
      state.isLoggedIn = status;
    },
    SET_USER_ID(state, userID) {
      state.userID = userID;
    },
  },
  actions: {
    async login({ commit }, userData) {
      const response = await loginUser(userData);
      commit('SET_LOGIN_STATUS', true);
      commit('SET_USER_ID', response.userID);
      return response.userID;
    },
    async logout({ commit }) {
      await logoutUser();
      commit('SET_LOGIN_STATUS', false);
      commit('SET_USER_ID', null);
    },
    async checkLoginStatus({ commit }) {
      const response = await isLoggedIn();
      commit('SET_LOGIN_STATUS', response.isLoggedIn);
      if (response.isLoggedIn) {
        commit('SET_USER_ID', response.userID);
      } else {
        commit('SET_USER_ID', null);
      }
    },
  },
});
