<template>
  <div class="user-profile">
    <div class="user-info">
      <h1 class="username">{{ user.username }}</h1>
      <div class="follower-info">
        <router-link :to="`/${user.id}/following`" class="follower-link">
          {{ user.following_num }} Following
        </router-link>
        <router-link :to="`/${user.id}/follower`" class="follower-link">
          {{ user.follower_num }} Followers
        </router-link>
      </div>
    </div>
    <ul class="tweet-list">
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

<style scoped>
.user-profile {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.username {
  font-size: 2.5rem;
  margin-bottom: 10px;
}

.follower-info {
  font-size: 1.2rem;
  color: #777;
}

.follower-link {
  text-decoration: none;
  color: #3498db;
  margin-right: 20px;
}

.tweet-list {
  margin-top: 20px;
  padding: 0;
  list-style: none;
}

.tweet-item {
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 15px;
  margin-bottom: 15px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.tweet-content {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

.tweet-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tweet-username {
  font-weight: bold;
}

.tweet-date {
  color: #888;
}
</style>
