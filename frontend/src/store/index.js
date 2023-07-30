import { createStore } from 'vuex';
import { checkLoggedIn } from '../api/userApi';

export default createStore({
  state: {
    isLoginPopupVisible: false,
    isSignupPopupVisible: false,
  },
  mutations: {
    SET_LOGIN_POPUP_VISIBILITY(state, isVisible) {
      state.isLoginPopupVisible = isVisible;
    },
    SET_SIGNUP_POPUP_VISIBILITY(state, isVisible) {
      state.isSignupPopupVisible = isVisible;
    },
  },
  actions: {
    async checkLoginStatus() {
      try {
        const response = await checkLoggedIn();
        if (response.isLoggedIn) {
          localStorage.setItem('isLoggedIn', response.isLoggedIn == true);
          localStorage.setItem('userID', response.userID);
        } else {
          localStorage.removeItem('isLoggedIn');
          localStorage.removeItem('userID');
        }
        } catch (error) {
        console.error('Error checking login status:', error);
      }
    },
  },
});
