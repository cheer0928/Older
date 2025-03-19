<template>
  <div class="app">
    <nav-bar />
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import NavBar from '@/components/NavBar.vue'
</script>

<style scoped>
.app {
  width: 100%;
  min-height: 100vh;
  background-color: #f5f5f5;
}
</style> 