<template>
  <Popup :buttonTitle="buttonTitle">
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
  </Popup>
</template>

<script setup>
import { defineProps, toRefs, reactive, ref } from 'vue';
import { updateUserProfile } from '../api/userApi';
import Popup from './Popup.vue';

const buttonTitle = "Edit Profile"
const errorMessage = ref('');

const props =  defineProps({
  currentUserData: Map
});
const { currentUserData } =  toRefs(props);

const user = reactive({
  username: currentUserData.value.username,
  iconUrl: currentUserData.value.iconUrl,
  selfIntroduction: currentUserData.value.selfIntroduction,
  dateOfBirth: currentUserData.value.dateOfBirth,
  location: currentUserData.value.location,
});

const updateProfile = async () => {
  try {
    console.log(currentUserData)
    const userResult = await updateUserProfile(user);
    currentUserData.value = userResult.user;
  } catch (error) {
    errorMessage.value = error;
  }
};

</script>
