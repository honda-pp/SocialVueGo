<template>
  <div>
    <button v-if="!showPopup" @click="showPopup = true">Edit Profile</button>
    <div v-if="showPopup" class="popup-overlay">
      <div class="popup-content">
        <button class="close-button" @click="showPopup = false">&times;</button>

        <div class="input-group">
          <label for="username">Username:</label>
          <input id="username" v-model="user.username" />
        </div>

        <div class="input-group">
          <label for="selfIntroduction">Self Introduction:</label>
          <textarea id="selfIntroduction" v-model="user.selfIntroduction"></textarea>
        </div>

        <div class="input-group">
            <label for="dateOfBirth">Date of Birth:</label>
            <input id="dateOfBirth" v-model="user.dateOfBirth" type="date" />
        </div>

        <div class="input-group">
          <label for="location">Location:</label>
          <input id="location" v-model="user.location" />
        </div>

        <button @click="updateProfile">Update</button>

        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, toRefs, reactive, ref } from 'vue';

const props = defineProps({
  currentUserData: Map
});
let { currentUserData } = toRefs(props);
const user = reactive({
  username: currentUserData.value.username,
  iconUrl: currentUserData.value.iconUrl,
  selfIntroduction: currentUserData.value.selfIntroduction,
  dateOfBirth: currentUserData.value.dateOfBirth,
  location: currentUserData.value.location,
});

const showPopup = ref(false);

const updateProfile = () => {
  // Add your update profile logic here
};

</script>
