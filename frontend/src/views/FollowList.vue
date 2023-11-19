<template>
  <h1>{{ title }}</h1>
  <UserList :userList="userList" />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getFollowingIDs, getFollowerIDs } from '../api/followApi';
import { getUserList } from '../api/userApi';
import UserList from '../components/UserList.vue';

const route = useRoute();
const relationshipType = ref(route.name);
const title = ref(`${relationshipType.value} List`);
const userList = ref([]);

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

</script>

<style>

</style>
