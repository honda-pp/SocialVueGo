<template>
  <div class="home-container">
    <p>
      You are currently logged in as {{ userID }}.
    </p>
  </div>
</template>

<script setup>
const userID = localStorage.getItem('userID');
</script>

<style>
.home-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  text-align: center;
}

p {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

.router-link {
  display: block;
  margin-top: 10px;
  color: #007bff;
  text-decoration: underline;
  cursor: pointer;
}
</style>
