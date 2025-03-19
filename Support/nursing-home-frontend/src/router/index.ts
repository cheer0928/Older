import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/nursing-homes',
      name: 'nursing-homes',
      component: () => import('../views/NursingHomeList.vue')
    },
    {
      path: '/nursing-homes/:id',
      name: 'nursing-home-detail',
      component: () => import('../views/NursingHomeDetail.vue')
    }
  ]
})

export default router 