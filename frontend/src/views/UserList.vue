<template>
  <div class="user-list">
    <h1>User List</h1>
    <ul>
      <li v-for="user in userList" :key="user.id">
        <router-link :to="`/${user.id}`">
        <span>{{ user.username }}</span>
      </router-link>
        <button @click="follow(user.id)" v-if="!user.followed && !isLoggedInUser(user.id)">Follow</button>
        <button @click="unfollow(user.id)" v-else-if="user.followed && !isLoggedInUser(user.id)">Unfollow</button>
        <span class="followed-label" v-if="user.followedByLoggedInUser || isLoggedInUser(user.id)">
          {{ isLoggedInUser(user.id) ? 'You' : 'Followed' }}
        </span>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { followUser, unfollowUser, getUserList, getFollowingIDs, getFollowerIDs } from '../api/userApi';

const userList = ref([]);
const loggedInUserId = parseInt(localStorage.getItem('userID'));

onMounted(async () => {
  try {
    await Promise.all([fetchUserList(), fetchFollowingUsers(), fetchFollowers()]);
  } catch (error) {
    console.error('Error fetching data:', error);
  }
});

const fetchUserList = async () => {
  try {
    const response = await getUserList();
    userList.value = response.userList;
  } catch (error) {
    console.error('Failed to fetch user list:', error);
  }
};

const fetchFollowingUsers = async () => {
  try {
    const response = await getFollowingIDs();
    const followingIDs = response.followingIDs;
    userList.value.forEach((user) => {
      user.followed = followingIDs.includes(user.id);
    });
  } catch (error) {
    console.error('Failed to fetch following users:', error);
  }
};

const fetchFollowers = async () => {
  try {
    const response = await getFollowerIDs();
    const followerIDs = response.followerIDs;
    userList.value.forEach((user) => {
      user.followedByLoggedInUser = followerIDs.includes(user.id);
    });
  } catch (error) {
    console.error('Failed to fetch followers:', error);
  }
};

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

button {
  margin-right: 10px;
}

.followed-label {
  color: green;
  font-weight: bold;
}
</style>
