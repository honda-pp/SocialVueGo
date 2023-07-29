<template>
  <div>
    <button v-if="isLoggedIn" @click="logout">Logout</button>
    <div v-else>
      <button @click="showLoginPopup">Login</button>
      <button @click="showSignupPopup">Signup</button>
    </div>

    <teleport to="body">
      <div v-if="isLoginPopupVisible" class="popup-overlay">
        <div class="popup-content">
          <button class="close-button" @click="hideLoginPopup">&times;</button>

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

      <div v-if="isSignupPopupVisible" class="popup-overlay">
        <div class="popup-content">
          <button class="close-button" @click="hideSignupPopup">&times;</button>

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

          <p v-if="signupErrorMessage" class="error-message">{{ signupErrorMessage }}</p>
        </div>
      </div>
    </teleport>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useStore } from 'vuex';

const isLoginPopupVisible = ref(false);
const isSignupPopupVisible = ref(false);
const username = ref('');
const password = ref('');
const errorMessage = ref('');
const newUsername = ref('');
const email = ref('');
const newPassword = ref('');
const signupErrorMessage = ref('');

const store = useStore();

const isLoggedIn = computed(() => store.state.isLoggedIn);

const showLoginPopup = () => {
  isLoginPopupVisible.value = true;
  disableScroll();
  errorMessage.value = '';
};

const hideLoginPopup = () => {
  isLoginPopupVisible.value = false;
  enableScroll();
};

const showSignupPopup = () => {
  isSignupPopupVisible.value = true;
  disableScroll();
  signupErrorMessage.value = '';
};

const hideSignupPopup = () => {
  isSignupPopupVisible.value = false;
  enableScroll();
};

const login = async () => {
  try {
    const userData = {
      username: username.value,
      password: password.value,
    };
    store.dispatch('login', userData);
    hideLoginPopup();
  } catch (error) {
    console.error('Login failed:', error);
    errorMessage.value = 'Login failed. Please check your username and password.';
  }
};

const logout = async () => {
  try {
    store.dispatch('logout');
  } catch (error) {
    console.error('Logout failed:', error);
  }
};

const disableScroll = () => {
  document.body.style.overflow = 'hidden';
};

const enableScroll = () => {
  document.body.style.overflow = 'auto';
};

const signup = async () => {
  try {
    const userData = {
      username: newUsername.value,
      email: email.value,
      password: newPassword.value,
    };
    store.dispatch('signup', userData);
    hideSignupPopup();
  } catch (error) {
    console.error('Signup failed:', error);
    signupErrorMessage.value = 'Signup failed. Please check your information and try again.';
  }
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
  position: relative;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;
  cursor: pointer;
  background: none;
  border: none;
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
