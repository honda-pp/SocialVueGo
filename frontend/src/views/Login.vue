<template>
  <div class="home-container">
    <div>
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
import { computed } from 'vue';
import { useStore } from 'vuex';

const store = useStore();

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

</style>
