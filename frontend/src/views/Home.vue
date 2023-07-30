<template>
  <div class="home-container">
    <p v-if="isLoggedIn">
      You are currently logged in as {{ userID }}.
      <router-link to="/user-list">User List</router-link>
      <router-link to="/tweet-list">Tweet List</router-link>
      <logout-button></logout-button>
    </p>
    <div v-else>
      <p>Ready to get started?</p>
      <div class="cta-buttons">
        <button @click="showLoginPopup">Login</button>
        <button @click="showSignupPopup">Signup</button>
      </div>
    </div>

    <teleport to="body">
      <login-popup v-if="isLoginPopupVisible"></login-popup>
      <signup-popup v-if="isSignupPopupVisible"></signup-popup>
    </teleport>
  </div>
</template>

<script setup>
import LoginPopup from '../components/LoginPopup.vue';
import SignupPopup from '../components/SignupPopup.vue';
import LogoutButton from '../components/LogoutButton.vue';
import { computed } from 'vue';
import { useStore } from 'vuex';

const store = useStore();

const isLoggedIn = computed(() => store.state.isLoggedIn);
const userID = computed(() => store.state.userID);
const isLoginPopupVisible = computed(() => store.state.isLoginPopupVisible);
const isSignupPopupVisible = computed(() => store.state.isSignupPopupVisible);

const showLoginPopup = () => {
  store.commit('SET_LOGIN_POPUP_VISIBILITY', true);
};

const showSignupPopup = () => {
  store.commit('SET_SIGNUP_POPUP_VISIBILITY', true);
};

</script>

<style>
.home-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  text-align: center;
}

p {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

.cta-buttons {
  margin-top: 20px;
}

.router-link {
  display: block;
  margin-top: 10px;
  color: #007bff;
  text-decoration: underline;
  cursor: pointer;
}
</style>
