<template>
  <header>
    <h1>{{ title }}</h1>
    <nav>
      <router-link to="/">Home</router-link>
    </nav>
  </header>
</template>

<script setup>
const title = 'SocialVueGo App';
</script>

<style>
header {
  background-color: #333;
  color: #fff;
  padding: 1rem;
  text-align: center;
}

nav {
  margin-top: 1rem;
}

nav a {
  color: #fff;
  text-decoration: none;
  margin-right: 1rem;
}

nav a:hover {
  text-decoration: underline;
}
</style>
