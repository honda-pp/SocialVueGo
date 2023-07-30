<template>
  <div class="tweet-list">
    <h1>Tweet List</h1>
    <ul>
      <li v-for="tweet in tweetList" :key="tweet.tweet_id">
        <span>{{ tweet.content }}</span>
        <span>{{ tweet.created_at }}</span>
        <span>{{ tweet.username }}</span>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getTweetList } from '../api/tweetApi';

const tweetList = ref([]);

onMounted(async () => {
  try {
    await fetchTweetList();
  } catch (error) {
    console.error('Failed to fetch tweet list:', error);
  }
});

const fetchTweetList = async () => {
  try {
    const response = await getTweetList();
    tweetList.value = response.tweetList;
  } catch (error) {
    console.error('Failed to fetch tweet list:', error);
  }
};
</script>

<style>
.tweet-list {
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

span {
  margin-right: 10px;
}
</style>
