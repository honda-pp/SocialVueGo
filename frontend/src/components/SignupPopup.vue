<template>
  <div v-if="isVisible" class="popup-overlay">
    <div class="popup-content">
      <button class="close-button" @click="closePopup">&times;</button>

      <div class="input-group">
        <label for="newUsername">Username:</label>
        <input type="text" id="newUsername" v-model="newUsername" />
      </div>

      <div class="input-group">
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" />
      </div>

      <div class="input-group">
        <label for="newPassword">Password:</label>
        <input type="password" id="newPassword" v-model="newPassword" />
      </div>

      <button @click="signup">Signup</button>

      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { signupUser } from '../api/userApi';
const router = useRouter();

const newUsername = ref('');
const email = ref('');
const newPassword = ref('');
const errorMessage = ref('');

const store = useStore();
const isVisible = ref(store.state.isSignupPopupVisible);

const closePopup = () => {
  store.commit('SET_SIGNUP_POPUP_VISIBILITY', false);
};

const signup = async () => {
  try {
    const userData = {
      username: newUsername.value,
      email: email.value,
      password: newPassword.value,
    };
    const response = await signupUser(userData);
    localStorage.setItem('isLoggedIn', 'true');
    localStorage.setItem('userID', response.userID);
    router.push('/');
  } catch (error) {
    console.error('Signup failed:', error);
    errorMessage.value = 'Signup failed. ' + error.error;
  }
};
</script>

<style>
@import '../common-styles.css';
</style>
