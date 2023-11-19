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
    <footer></footer>
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

h1 {
  font-size: 2.5rem;
  margin-bottom: 20px;
}

ul {
  list-style: none;
  padding: 0;
}

p {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

li {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

nav ul {
  list-style: none;
  display: flex;
  gap: 10px;
}

nav ul li {
  display: inline;
}

main {
  max-width: 800px;
  margin: 0 auto;
  margin-bottom: 40px;
  padding: 10px;
}

footer {
  background-color: #f2f2f2;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 50px;
  padding: 10px;
}

</style>
