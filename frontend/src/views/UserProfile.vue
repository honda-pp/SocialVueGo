<template>
  <div class="user-profile">
    <h1>User Profile: {{ username }}</h1>
    <!-- ここにユーザーページのコンテンツを追加してください -->
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getUserInfo } from '../api/userApi';

const username = ref('');
const route = useRoute();

onMounted(async () => {
  try {
    const response = await getUserInfo(route.params.userID)
    username.value = response.user.username
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
});
</script>

<style>
.user-profile {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  font-size: 2.5rem;
  margin-bottom: 20px;
}
</style>
