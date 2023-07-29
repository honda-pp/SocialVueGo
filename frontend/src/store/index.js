import { createStore } from 'vuex';
import { loginUser, logoutUser, signupUser, checkLoggedIn } from '../api/userApi';

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
      try {
        const response = await loginUser(userData);
        commit('SET_LOGIN_STATUS', true);
        commit('SET_USER_ID', response.userID);
      } catch (error) {
        console.error('Login failed:', error);
        throw new Error('Login failed. Please check your username and password.');
      }
    },
    async logout({ commit }) {
      try {
        await logoutUser();
        commit('SET_LOGIN_STATUS', false);
        commit('SET_USER_ID', null);
      } catch (error) {
        console.error('Logout failed:', error);
      }
    },
    async checkLoginStatus({ commit }) {
      try {
        const response = await checkLoggedIn();
        commit('SET_LOGIN_STATUS', response.isLoggedIn);
        if (response.isLoggedIn) {
          commit('SET_USER_ID', response.userID);
        } else {
          commit('SET_USER_ID', null);
        }
      } catch (error) {
        console.error('Error checking login status:', error);
      }
    },
    async signup({ commit }, userData) {
      try {
        await signupUser(userData);
        commit('SET_LOGIN_STATUS', true);
        commit('SET_USER_ID', userData.userID);
      } catch (error) {
        console.error('Signup failed:', error);
        throw new Error('Signup failed. Please check your information and try again.');
      }
    },
  },
});
