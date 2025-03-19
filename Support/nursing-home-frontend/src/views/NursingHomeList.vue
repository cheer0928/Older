<template>
  <div class="nursing-home-list">
    <h1>养老院列表</h1>
    <div class="search-bar">
      <input v-model="searchQuery" placeholder="搜索养老院..." />
    </div>
    <div class="homes-grid">
      <div v-for="home in filteredHomes" :key="home.id" class="home-card">
        <h2>{{ home.name }}</h2>
        <div class="info">
          <p><strong>地址：</strong>{{ home.address }}</p>
          <p><strong>床位数：</strong>{{ home.beds }}</p>
          <p><strong>价格：</strong>¥{{ home.price }}/月</p>
          <p><strong>评分：</strong>{{ home.rating }}/5</p>
        </div>
        <router-link :to="`/nursing-homes/${home.id}`" class="view-detail">
          查看详情
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { nursingHomeApi } from '@/services/api'
import type { NursingHome } from '@/types/nursing-home'

const homes = ref<NursingHome[]>([])
const searchQuery = ref('')

const filteredHomes = computed(() => {
  return homes.value.filter(home => 
    home.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
    home.address.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

onMounted(async () => {
  try {
    const response = await nursingHomeApi.getAll()
    homes.value = response.data
  } catch (error) {
    console.error('获取养老院列表失败:', error)
  }
})
</script>

<style scoped>
.nursing-home-list {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.search-bar {
  margin-bottom: 20px;
}

.search-bar input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.homes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.home-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.view-detail {
  display: inline-block;
  margin-top: 10px;
  padding: 8px 16px;
  background: #4CAF50;
  color: white;
  text-decoration: none;
  border-radius: 4px;
}
</style> 