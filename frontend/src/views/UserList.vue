<template>
  <div class="user-list">
    <h1>User List</h1>
    <ul>
      <li v-for="user in userList" :key="user.id">
        {{ user.username }}
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';

const store = useStore();
const userList = ref([]);

onMounted(async () => {
  try {
    await store.dispatch('getUserList');
    userList.value = store.state.userList;
  } catch (error) {
    console.error('Failed to fetch user list:', error);
  }
});
</script>

<style>
.user-list {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  font-size: 2.5rem;
  margin-bottom: 20px;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  font-size: 1.2rem;
  margin-bottom: 10px;
}
</style>
