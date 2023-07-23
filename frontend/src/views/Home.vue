<template>
  <div>
    <button @click="showLoginPopup">Login</button>

    <teleport to="body">
      <div v-if="isLoginPopupVisible" class="popup-overlay">
        <div class="popup-content">
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
    </teleport>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { loginUser } from '../api/userApi';

const isLoginPopupVisible = ref(false);
const username = ref('');
const password = ref('');
const errorMessage = ref('');

const showLoginPopup = () => {
  isLoginPopupVisible.value = true;
  disableScroll();
  errorMessage.value = '';
};

const hideLoginPopup = () => {
  isLoginPopupVisible.value = false;
  enableScroll();
};

const login = async () => {
  try {
    const userData = {
      username: username.value,
      password: password.value,
    };
    const response = await loginUser(userData);
    console.log('Login successful:', response);
    hideLoginPopup();
  } catch (error) {
    console.error('Login failed:', error);
    errorMessage.value = 'Login failed. Please check your username and password.';
  }
};

const disableScroll = () => {
  document.body.style.overflow = 'hidden';
};

const enableScroll = () => {
  document.body.style.overflow = 'auto';
};
</script>

<style>
.popup-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 9999;
  display: flex;
  justify-content: center;
  align-items: center;
}

.popup-content {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
}

.input-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
}

.error-message {
  color: red;
}
</style>
