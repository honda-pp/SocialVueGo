<template>
  <ul>
    <li v-for="user in userList" :key="user.id" class="user-list">
      <router-link :to="`/${user.id}`" class="user-link">
        <img :src="user.iconUrl" alt="User Icon" class="list-user-icon">
        <span class="list-user-name">{{ user.username }}</span>
      </router-link>
      <button @click="follow(user.id)" v-if="!user.followed && !isLoggedInUser(user.id)">Follow</button>
      <button @click="unfollow(user.id)" v-else-if="user.followed && !isLoggedInUser(user.id)">Unfollow</button>
      <span class="followed-label" v-if="user.followedByLoggedInUser || isLoggedInUser(user.id)">
        {{ isLoggedInUser(user.id) ? 'You' : 'Followed' }}
      </span>
    </li>
  </ul>
</template>

<script setup>
import { defineProps, toRefs } from 'vue';
import { followUser, unfollowUser } from '../api/followApi';


const props = defineProps({
  userList: Array
});
let { userList } = toRefs(props);
const loggedInUserId = parseInt(localStorage.getItem('userID'));


const follow = async (userId) => {
  try {
    await followUser(userId);
    const userIndex = userList.value.findIndex((user) => user.id === userId);
    if (userIndex !== -1) {
      userList.value[userIndex].followed = true;
    }
  } catch (error) {
    console.error('Failed to follow user:', error);
  }
};

const unfollow = async (userId) => {
  try {
    await unfollowUser(userId);
    const userIndex = userList.value.findIndex((user) => user.id === userId);
    if (userIndex !== -1) {
      userList.value[userIndex].followed = false;
    }
  } catch (error) {
    console.error('Failed to unfollow user:', error);
  }
};

const isLoggedInUser = (userId) => {
  return loggedInUserId === userId;
};
</script>

<style>
.user-list {
  display: flex;
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  align-items: center;
  height: 50px;
}

.user-link {
  text-decoration: none;
  color: #3498db;
  margin-right: 20px;
}

.list-user-icon {
  width: 50px;
}

.list-user-name {
  margin-left: 10px;
  vertical-align: 50%;
}

.followed-label {
  color: green;
  font-weight: bold;
}
</style>
