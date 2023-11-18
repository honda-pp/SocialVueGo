<template>
  <div class="user-list">
    <h1>{{ title }}</h1>
      <UserList :userList="userList" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getFollowingIDs, getFollowerIDs } from '../api/followApi';
import { getUserList } from '../api/userApi';
import UserList from '../components/UserList.vue';

const title = ref(`User List`);
const userList = ref([]);

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

</script>

<style>

h1 {
  font-size: 2.5rem;
  margin-bottom: 20px;
}

ul {
  list-style: none;
  padding: 0;
}

</style>
