<template>
  <p>
    You are currently logged in as {{ username }}.
  </p>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getSessionInfo } from '../api/userApi';

const username = ref('');
onMounted(async () => {
  try {
    await fetchUserInfo()
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
});

const fetchUserInfo = async () => {
  try {
    const response = await getSessionInfo()
    username.value = response.username
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
}


</script>

<style>

p {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

.router-link {
  display: block;
  margin-top: 10px;
  color: #007bff;
  text-decoration: underline;
  cursor: pointer;
}
</style>
