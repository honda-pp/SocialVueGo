import { createStore } from 'vuex';
import { loginUser, logoutUser, isLoggedIn } from '../api/userApi';

export default createStore({
  state: {
    isLoggedIn: false,
  },
  mutations: {
    SET_LOGIN_STATUS(state, status) {
      state.isLoggedIn = status;
    },
  },
  actions: {
    async login({ commit }, userData) {
      const response = await loginUser(userData);
      commit('SET_LOGIN_STATUS', true);
      return response.userID;
    },
    async logout({ commit }) {
      await logoutUser();
      commit('SET_LOGIN_STATUS', false);
    },
    async checkLoginStatus({ commit }) {
      const response = await isLoggedIn();
      commit('SET_LOGIN_STATUS', response.isLoggedIn);
    },
  },
});
