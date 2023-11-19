<template>
  <ul class="tweet-list">
    <li v-for="tweet in tweetList" :key="tweet.tweet_id" class="tweet-item">
      <div class="tweet-content">{{ tweet.content }}</div>
      <div class="tweet-info">
        <router-link :to="`/${tweet.userID}`" class="user-link">
          <span class="tweet-username">{{ tweet.username }}</span>
        </router-link>
        <span class="tweet-date">{{ formatDate(tweet.createdAt) }}</span>
      </div>
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
  color: #3498db;
}

.tweet-date {
  color: #888;
}
</style>