<template>
  <div class="home-container">
    <p>
      You are currently logged in as {{ username }}.
    </p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getUserInfo } from '../api/userApi';

const username = ref('');
onMounted(async () => {
  try {
    await fetchUserInfo(localStorage.getItem('userID'))
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
});

const fetchUserInfo = async (userID) => {
  try {
    const response = await getUserInfo(userID)
    username.value = response.user.username
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
}


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

.router-link {
  display: block;
  margin-top: 10px;
  color: #007bff;
  text-decoration: underline;
  cursor: pointer;
}
</style>
