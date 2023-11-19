<template>
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
</template>

<script setup>
import { defineProps, toRefs, ref } from 'vue';
import { createTweet } from '../api/tweetApi';

const isTweetPopupVisible = ref(false);
const tweetContent = ref('');
const errorMessage = ref('');
const props = defineProps({
  tweetList: Array
});
let { tweetList } = toRefs(props);

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
    const tweetResult = await createTweet(newTweet);
    tweetList.value.unshift(tweetResult.tweet);
    hideTweetPopup();
  } catch (error) {
    errorMessage.value = error;
  }
};

const showTweetPopup = () => {
  isTweetPopupVisible.value = true;
  disableScroll();
  errorMessage.value = '';
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
