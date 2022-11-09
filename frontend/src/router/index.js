import { createRouter, createWebHistory } from 'vue-router';
import MainView from '../views/MainView.vue';
import StatsView from '../views/StatsView.vue';
import ImportView from '../views/ImportView.vue';
import GameView from '../views/GameView.vue';


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
    {
      path: '/import',
      name: 'import-external',
      component: ImportView
    },
    {
      path: '/external/:gameId',
      name: 'external-game',
      component: GameView
    },
  ]
})

export default router
