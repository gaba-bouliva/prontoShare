<script setup lang="ts">
</script>

<template>
  <div class="max-h-screen flex items-center justify-center bg-gray-100 text-center">
    <div>
      <h1 class="text-4xl font-bold text-blue-600">🎉 ProntoShare is working!</h1>
      <a href="https://vite.dev" target="_blank">
        <img src="/logo.png" class="logo" alt="Vite logo" />
      </a>
    </div>
    <div>
      <router-view/>
    </div>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
