<template>
  <div class="nursing-home-detail">
    <div v-if="home" class="detail-card">
      <div class="header">
        <h1>{{ home.name }}</h1>
        <div class="rating">
          <span class="stars">★★★★★</span>
          <span class="rating-number">{{ home.rating }}/5</span>
        </div>
      </div>

      <div class="info-grid">
        <div class="info-item">
          <h3>基本信息</h3>
          <p><strong>地址：</strong>{{ home.address }}</p>
          <p><strong>联系电话：</strong>{{ home.phone }}</p>
          <p><strong>床位数：</strong>{{ home.beds }}</p>
          <p><strong>月费用：</strong>¥{{ home.price }}</p>
        </div>

        <div class="info-item">
          <h3>提供服务</h3>
          <div class="services">
            <span v-for="service in servicesList" 
                  :key="service" 
                  class="service-tag">
              {{ service }}
            </span>
          </div>
        </div>
      </div>

      <div class="description">
        <h3>机构介绍</h3>
        <p>{{ home.description }}</p>
      </div>

      <div class="actions">
        <button class="primary-btn" @click="handleContact">联系养老院</button>
        <button class="secondary-btn" @click="handleReserve">预约参观</button>
      </div>
    </div>

    <div v-else class="loading">
      加载中...
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { nursingHomeApi } from '@/services/api'
import type { NursingHome } from '@/types/nursing-home'

const route = useRoute()
const home = ref<NursingHome | null>(null)

const servicesList = computed(() => {
  if (!home.value?.services) return []
  return home.value.services.split(',').map(s => s.trim())
})

const handleContact = () => {
  if (home.value) {
    window.location.href = `tel:${home.value.phone}`
  }
}

const handleReserve = () => {
  // 实现预约功能
  console.log('预约参观')
}

onMounted(async () => {
  try {
    const id = route.params.id as string
    const response = await nursingHomeApi.getById(parseInt(id))
    home.value = response.data
  } catch (error) {
    console.error('获取养老院详情失败:', error)
  }
})
</script>

<style scoped>
.nursing-home-detail {
  padding: 20px;
  max-width: 1000px;
  margin: 0 auto;
}

.detail-card {
  background: white;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.rating {
  text-align: right;
}

.stars {
  color: #ffd700;
  letter-spacing: 2px;
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30px;
  margin-bottom: 30px;
}

.info-item {
  background: #f8f8f8;
  padding: 20px;
  border-radius: 8px;
}

.services {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.service-tag {
  background: #e8f5e9;
  color: #4caf50;
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 14px;
}

.description {
  margin-bottom: 30px;
}

.actions {
  display: flex;
  gap: 20px;
}

.primary-btn, .secondary-btn {
  padding: 12px 24px;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  font-weight: bold;
}

.primary-btn {
  background: #4caf50;
  color: white;
}

.secondary-btn {
  background: #e8f5e9;
  color: #4caf50;
}

.loading {
  text-align: center;
  padding: 40px;
  font-size: 18px;
  color: #666;
}
</style> 