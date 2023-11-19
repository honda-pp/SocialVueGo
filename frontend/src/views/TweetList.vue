<template>
  <h1>Tweet List</h1>
  <button @click="showTweetPopup">Create Tweet</button>

  <TweetList :tweetList="tweetList" />

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
import TweetList from '../components/TweetList.vue';

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
span {
  margin-right: 10px;
}

textarea {
  width: 100%;
  height: 100px;
  resize: none;
}
</style>
