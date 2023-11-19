<template>
  <button v-if="!showPopup" @click="showPopup = true">{{ buttonTitle }}</button>
  <div v-if="showPopup" class="popup-overlay">
    <div class="popup-content">
      <button class="close-button" @click="showPopup = false">&times;</button>
      <slot></slot>
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref, defineProps, toRefs } from 'vue';

const showPopup = ref(false);
const props = defineProps({
  buttonTitle: String,
  errorMessage: String
});
const { buttonTitle, errorMessage } = toRefs(props);
</script>
