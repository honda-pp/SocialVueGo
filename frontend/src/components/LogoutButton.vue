<template>
  <button @click="logout">Logout</button>
</template>

<script setup>
import { useStore } from 'vuex';

const store = useStore();

const logout = async () => {
  try {
    await store.dispatch('logout');
  } catch (error) {
    console.error('Logout failed:', error);
  }
};
</script>

<style>
button {
  padding: 10px 20px;
  margin: 0 10px;
  border: none;
  background-color: #007bff;
  color: #fff;
  font-size: 1.2rem;
  cursor: pointer;
  border-radius: 5px;
}

button:hover {
  background-color: #0056b3;
}
</style>
