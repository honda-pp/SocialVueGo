<template>
  <h1>Tweet List</h1>
  <tweetPopup :tweetList="tweetList" />
  <TweetList :tweetList="tweetList" />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getTweetList } from '../api/tweetApi';
import TweetList from '../components/TweetList.vue';
import tweetPopup from '../components/tweetPopup.vue';

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
</style>