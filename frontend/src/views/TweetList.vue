<template>
  <h1>Tweet List</h1>
  <button @click="showTweetPopup">Create Tweet</button>

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
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getTweetList, createTweet } from '../api/tweetApi';
import { formatDate } from '../utils/dateUtil';

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

.user-link {
  text-decoration: none;
  color: #3498db;
  margin-right: 20px;
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
