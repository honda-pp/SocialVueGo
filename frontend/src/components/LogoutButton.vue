<template>
  <button id="logout-button" @click="logout">Logout</button>
</template>

<script setup>
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

const store = useStore();
const router = useRouter();

const logout = async () => {
  try {
    await store.dispatch('logout');
    router.push('/login')
  } catch (error) {
    console.error('Logout failed:', error);
  }
};
</script>

<style>

#logout-button {
  height: 40px;
  width: 100px;
}
</style>
