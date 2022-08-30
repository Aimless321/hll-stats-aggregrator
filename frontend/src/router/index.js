import { createRouter, createWebHistory } from 'vue-router'
import MainView from '../views/MainView.vue'
import StatsView from '../views/StatsView.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'main',
      component: MainView
    },
    {
      path: '/stats/:steamId',
      name: 'player-stats',
      component: StatsView
    },
  ]
})

export default router
