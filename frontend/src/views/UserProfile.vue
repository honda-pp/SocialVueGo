<template>
  <div class="user-info">
    <img :src="user.iconUrl" alt="User Icon" class="user-icon">
    <EditProfile v-if="mounted && showEditButon" :currentUserData="user" />
    <div class="username">{{ user.username }}</div>
    <div class="self-introduction">{{ user.selfIntroduction }}</div>
    <div class="user-info-row">
      <span class="date-of-birth">{{ user.dateOfBirth ? user.dateOfBirth.slice(0, 10) : '' }}</span>
      <span class="user-location">{{ user.location }}</span>
    </div>
    <div class="follower-info">
      <router-link :to="`/${user.id}/following`" class="follower-link">
        {{ user.followingNum }} Following
      </router-link>
      <router-link :to="`/${user.id}/follower`" class="follower-link">
        {{ user.followerNum }} Followers
      </router-link>
    </div>
  </div>
  <TweetList :tweetList="tweetList" />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getUserInfo } from '../api/userApi';
import { getTweetListByUserID } from '../api/tweetApi';
import TweetList from '../components/TweetList.vue';
import EditProfile from '../components/EditProfile.vue';

const user = ref({});
const tweetList = ref([]);
const mounted = ref(false);
const route = useRoute();
const showEditButon = route.params.userID === localStorage.getItem('userID');

onMounted(async () => {
  try {
    await fetchUserInfo(route.params.userID)
    await fetchTweetList(route.params.userID)
    mounted.value = true;
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
});

const fetchUserInfo = async (userID) => {
  try {
    const response = await getUserInfo(userID)
    user.value = response.user
  } catch (error) {
    console.error('Error fetching user info:', error);
  }
}

const fetchTweetList = async (userID) => {
  try {
    const response = await getTweetListByUserID(userID);
    tweetList.value = response.tweetList;
  } catch (error) {
    console.error('Failed to fetch tweet list:', error);
  }
};
</script>

<style>
.user-icon {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  margin-bottom: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.username {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 10px;
}

.self-introduction {
  font-size: 1.2rem;
  margin-bottom: 20px;
}

.user-info-row {
  color: #777;
  margin-right: 20px;
}

.follower-info {
  font-size: 1.2rem;
  color: #777;
  margin-bottom: 20px;
}

.follower-link {
  text-decoration: none;
  color: #3498db;
  margin-right: 20px;
}

</style>
