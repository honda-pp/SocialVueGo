<template>
  <div class="user-list">
    <h1>User List</h1>
    <ul>
      <li v-for="user in userList" :key="user.id">
        <span>{{ user.username }}</span>
        <button @click="followUser(user.id)" v-if="!user.followed">Follow</button>
        <button @click="unfollowUser(user.id)" v-else>Unfollow</button>
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

const followUser = async (userId) => {
  try {
    await store.dispatch('followUser', userId);
    const userIndex = userList.value.findIndex((user) => user.id === userId);
    if (userIndex !== -1) {
      userList.value[userIndex].followed = true;
    }
  } catch (error) {
    console.error('Failed to follow user:', error);
  }
};

const unfollowUser = async (userId) => {
  try {
    await store.dispatch('unfollowUser', userId);
    const userIndex = userList.value.findIndex((user) => user.id === userId);
    if (userIndex !== -1) {
      userList.value[userIndex].followed = false;
    }
  } catch (error) {
    console.error('Failed to unfollow user:', error);
  }
};
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
