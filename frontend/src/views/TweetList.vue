<template>
  <div class="tweet-list">
    <h1>Tweet List</h1>
    <ul>
      <li v-for="tweet in tweetList" :key="tweet.tweet_id">
        <span>{{ tweet.content }}</span>
        <span>{{ tweet.createdAt }}</span>
        <span>{{ tweet.username }}</span>
      </li>
    </ul>

    <button @click="showTweetPopup">Create Tweet</button>

    <teleport to="body">
      <div v-if="isTweetPopupVisible" class="popup-overlay">
        <div class="popup-content">
          <button class="close-button" @click="hideTweetPopup">&times;</button>

          <div class="input-group">
            <label for="tweetContent">Tweet Content:</label>
            <textarea id="tweetContent" v-model="tweetContent"></textarea>
          </div>

          <button @click="createNewTweet">Create</button>

          <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
        </div>
      </div>
    </teleport>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getTweetList, createTweet } from '../api/tweetApi';

const tweetList = ref([]);
const isTweetPopupVisible = ref(false);
const tweetContent = ref('');
const errorMessage = ref('');

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

const showTweetPopup = () => {
  isTweetPopupVisible.value = true;
  disableScroll();
  errorMessage.value = '';
};

const hideTweetPopup = () => {
  tweetContent.value = ''
  isTweetPopupVisible.value = false;
  enableScroll();
};

const createNewTweet = async () => {
  try {
    const newTweet = {
      content: tweetContent.value,
    };
    await createTweet(newTweet);
    await fetchTweetList();
    hideTweetPopup();
  } catch (error) {
    errorMessage.value = error;
  }
};

const disableScroll = () => {
  document.body.style.overflow = 'hidden';
};

const enableScroll = () => {
  document.body.style.overflow = 'auto';
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

textarea {
  width: 100%;
  height: 100px;
  resize: none;
}
</style>
