<template>
  <div class="user-profile">
    <h1>User Profile: {{ user.username }}</h1>
    <div>
      <span>{{ user.following_num }}</span>
      following
      <span>{{ user.follower_num }}</span>
      follower
    </div>
    <ul>
      <li v-for="tweet in tweetList" :key="tweet.tweet_id" class="tweet-item">
        <div class="tweet-content">{{ tweet.content }}</div>
        <div class="tweet-info">
          <span class="tweet-username">{{ tweet.username }}</span>
          <span class="tweet-date">{{ formatDate(tweet.createdAt) }}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { getUserInfo } from '../api/userApi';
import { getTweetListByUserID } from '../api/tweetApi';
import { formatDate } from '../utils/dateUtil';

const user = ref({});
const tweetList = ref([]);
const route = useRoute();

onMounted(async () => {
  try {
    await fetchUserInfo(route.params.userID)
    await fetchTweetList(route.params.userID)
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
.user-profile {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  font-size: 2.5rem;
  margin-bottom: 20px;
}

.tweet-list {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.tweet-item {
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
  margin-bottom: 10px;
}

.tweet-content {
  font-size: 1.2rem;
  margin-bottom: 5px;
}

.tweet-info {
  display: flex;
  justify-content: space-between;
}

.tweet-username {
  font-weight: bold;
}

.tweet-date {
  color: #888;
}
</style>
