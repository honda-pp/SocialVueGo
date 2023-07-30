<template>
  <div class="tweet-list">
    <h1>Tweet List</h1>
    <ul>
      <li v-for="tweet in tweetList" :key="tweet.tweet_id" class="tweet-item">
        <div class="tweet-content">{{ tweet.content }}</div>
        <div class="tweet-info">
          <span class="tweet-username">{{ tweet.username }}</span>
          <span class="tweet-date">{{ formatDate(tweet.createdAt) }}</span>
        </div>
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

const formatDate = (dateString) => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  return `${year}-${month}-${day} ${hours}:${minutes}`;
};
</script>

<style>
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
