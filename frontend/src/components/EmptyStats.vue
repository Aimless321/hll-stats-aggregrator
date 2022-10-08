<template>
  <div class="mx-auto max-w-md sm:max-w-3xl min-h-screen grid place-content-center">
    <div>
      <div class="text-center">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 48 48"
             aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M34 40h10v-4a6 6 0 00-10.712-3.714M34 40H14m20 0v-4a9.971 9.971 0 00-.712-3.714M14 40H4v-4a6 6 0 0110.713-3.714M14 40v-4c0-1.313.253-2.566.713-3.714m0 0A10.003 10.003 0 0124 26c4.21 0 7.813 2.602 9.288 6.286M30 14a6 6 0 11-12 0 6 6 0 0112 0zm12 6a4 4 0 11-8 0 4 4 0 018 0zm-28 0a4 4 0 11-8 0 4 4 0 018 0z"/>
        </svg>
        <h2 class="mt-2 text-lg font-medium text-gray-900">Search player stats</h2>
      </div>
      <p v-if="error" class="text-red-500">Invalid steamId entered, or player did not play any games on our server</p>
      <form class="mt-6 sm:flex sm:items-center">
        <div class="relative rounded-md shadow-sm sm:min-w-0 sm:flex-1">
          <input v-model="query" type="text"
                 class="block w-full rounded-md border-gray-300 pr-32 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                 placeholder="Enter a steamid"/>
        </div>
        <div class="mt-3 sm:mt-0 sm:ml-4 sm:flex-shrink-0">
          <button @click="search(query)"
                  type="submit"
                  class="block w-full rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-center text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
            Search
          </button>
        </div>
      </form>
    </div>
    <div class="mt-10">
      <h3 class="text-sm font-medium text-gray-500">Recent searches</h3>
      <ul role="list" class="mt-4 grid grid-cols-2 gap-4 sm:grid-cols-3">
        <li v-for="(query, id) in recent" :key="id">
          <button @click="search(query.steamId)"
                  type="button"
                  class="group flex w-full items-center justify-between space-x-3 rounded-full border border-gray-300 p-2 text-left shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
            <span class="flex min-w-0 flex-1 items-center space-x-3">
              <span class="block flex-shrink-0"></span>
              <span class="block min-w-0 flex-1">
                <span class="block truncate text-sm font-medium text-gray-900">{{ query.name }}</span>
                <span class="block truncate text-sm font-medium text-gray-500">{{ query.steamId }}</span>
              </span>
            </span>
            <span class="inline-flex h-10 w-10 flex-shrink-0 items-center justify-center">
              <ChevronRightIcon class="h-5 w-5 text-gray-400 group-hover:text-gray-500" aria-hidden="true"/>
            </span>
          </button>
        </li>
      </ul>
    </div>
    <div class="mt-10 grid grid-cols-2 gap-3">
      <HighKillGames :data="stats.highKillGames" />
      <HighKPMPlayers :data="stats.highKPMGamers" />
    </div>
  </div>
</template>

<script setup>
import {ChevronRightIcon} from '@heroicons/vue/20/solid'
import {ref, reactive} from "vue";
import {useRouter} from "vue-router";
import HighKillGames from "./HighKillGames.vue";
import HighKPMPlayers from "./HighKPMPlayers.vue";

const router = useRouter();
const apiUrl = import.meta.env.VITE_API_BASE_URL;
const recruitmentUrl = `${apiUrl}/stats/recruitment`;
const stats = reactive({
  highKillGames: null,
  highKPMGamers: null,
});

fetch(recruitmentUrl, {credentials: "include"})
    .then((resp) => {
      if (resp.status === 200) {
        return resp.json()
      }

      if (resp.status === 401) {
        localStorage.clear();
        router.replace('/');
        return;
      }

      throw new Error('Something went wrong');
    })
    .then(data => {
      Object.assign(stats, data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });

const props = defineProps({
  error: {
    type: Boolean,
    default: false
  }
});
const query = ref("");

const recent = JSON.parse(localStorage.getItem('recentSearches')) || [];

function search(steamId) {
  router.push(`/stats/${steamId}`);
}
</script>
