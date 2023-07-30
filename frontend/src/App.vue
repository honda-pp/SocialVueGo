<template>
  <div id="app">
    <header v-if="isLoggedIn">
      <nav>
        <ul>
          <li><router-link to="/">Home</router-link></li>
          <li><router-link to="/user-list">User List</router-link></li>
          <li><router-link to="/tweet-list">Tweet List</router-link></li>
        </ul>
      </nav>
      <logout-button />
    </header>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useStore } from 'vuex';
import LogoutButton from './components/LogoutButton.vue';

const store = useStore()

const isLoggedIn = computed(() => store.state.isLoggedIn);

</script>

<style>
header {
  background-color: #f2f2f2;
  padding: 10px;
  display: flex;
  justify-content: space-between;
}

nav ul {
  list-style: none;
  display: flex;
  gap: 10px;
}

nav ul li {
  display: inline;
}
</style>
