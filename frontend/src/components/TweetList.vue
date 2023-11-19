<template>
  <ul class="tweet-list">
    <li v-for="tweet in tweetList" :key="tweet.tweet_id" class="tweet-item">
      <img :src="tweet.iconUrl" alt="User Icon" class="tweet-user-icon">
      <span class="tweet-username">{{ tweet.username }}</span>
      <span class="tweet-date">{{ formatDate(tweet.createdAt) }}</span>
      <div class="tweet-content">{{ tweet.content }}</div>
    </li>
  </ul>
</template>

<script setup>
import { defineProps, toRefs } from 'vue';
import { formatDate } from '../utils/dateUtil';

const props = defineProps({
  tweetList: Array
});
let { tweetList } = toRefs(props);
</script>

<style>
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

.tweet-user-icon {
  float: left;
  width: 50px;
}

.tweet-username {
  padding: 10px;
  vertical-align: top;
  font-weight: bold;
}

.tweet-link {
  text-decoration: none;
  color: #3498db;
}

.tweet-date {
  color: #888;
  vertical-align: top;
}

.tweet-content {
  margin-left: 60px;
  vertical-align: top;
}
</style>