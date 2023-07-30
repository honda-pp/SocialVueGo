<template>
  <div v-if="isVisible" class="popup-overlay">
    <div class="popup-content">
      <button class="close-button" @click="closePopup">&times;</button>

      <div class="input-group">
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" />
      </div>

      <div class="input-group">
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" />
      </div>

      <button @click="login">Login</button>

      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

const username = ref('');
const password = ref('');
const errorMessage = ref('');

const store = useStore();
const router = useRouter();

const isVisible = ref(store.state.isLoginPopupVisible);

const closePopup = () => {
  store.commit('SET_LOGIN_POPUP_VISIBILITY', false);
};

const login = async () => {
  try {
    const userData = {
      username: username.value,
      password: password.value,
    };
    await store.dispatch('login', userData);
    router.push('/')
  } catch (error) {
    errorMessage.value = error;
  }
};
</script>

<style>
@import '../common-styles.css';
</style>
