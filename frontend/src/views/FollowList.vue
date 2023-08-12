<template>
  <div class="user-list">
    <h1>{{ title }}</h1>
    <ul>
      <li v-for="user in userList" :key="user.id">
        <router-link :to="`/${user.id}`" class="user-link">
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
import { useRoute } from 'vue-router';
import { followUser, unfollowUser, getFollowingIDs, getFollowerIDs } from '../api/followApi';
import { getUserList } from '../api/userApi';

const route = useRoute();
const relationshipType = ref(route.name);
const title = ref(`${relationshipType.value} List`);
const userList = ref([]);
const loggedInUserId = parseInt(localStorage.getItem('userID'));

onMounted(async () => {
  try {
    await Promise.all([fetchUserList(), fetchFollowingUsers(), fetchFollowers()]);
  } catch (error) {
    console.error(`Error fetching ${relationshipType.value} list:`, error);
  }
});

const fetchUserList = async () => {
  try {
    const response = await getUserList(route.params.userID, relationshipType.value);
    userList.value = response.userList;
  } catch (error) {
    console.error(`Failed to fetch ${relationshipType.value} list:`, error);
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

.user-link {
  text-decoration: none;
  color: #3498db;
  margin-right: 20px;
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
