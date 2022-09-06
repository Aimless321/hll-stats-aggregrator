<script setup>
import {reactive, ref} from "vue";
import {useRoute, useRouter} from "vue-router";
import {MagnifyingGlassIcon} from '@heroicons/vue/20/solid'
import AvgStats from "../components/AvgStats.vue";
import LastGames from "../components/LastGames.vue";
import EmptyStats from "../components/EmptyStats.vue";

const router = useRouter();
const route = useRoute();

const notFound = ref(false);
const query = ref("");
const apiUrl = import.meta.env.VITE_API_BASE_URL;
const steamId = route.params.steamId;
const statsUrl = `${apiUrl}/stats/${steamId}`;
const stats = reactive({
  steamInfo: null,
  publicAvg: null,
  compAvg: null,
  publicGames: null,
  compGames: null
});

fetch(statsUrl, {credentials: "include"})
    .then((resp) => {
      if (resp.status === 200) {
        return resp.json()
      }

      if (resp.status === 404) {
        notFound.value = true;
      }

      throw new Error('Something went wrong');
    })
    .then(data => {
      // Remove existing entry and prepend new entry to the front
      const recent = JSON.parse(localStorage.getItem('recentSearches')) || [];
      const name = data.steamInfo.PersonaName;
      const newRecent = recent.filter(query => query.steamId !== steamId);
      newRecent.unshift({steamId, name});
      localStorage.setItem('recentSearches', JSON.stringify(newRecent.slice(0, 6)));

      Object.assign(stats, data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
</script>

<template>
  <div v-if="!notFound">
    <div v-if="stats.steamInfo" class="flex items-center gap-x-3 pt-8">
      <img :src="stats.steamInfo.avatarmedium" class="rounded-full">
      <h1 class="text-xl text-gray-900">{{ stats.steamInfo.PersonaName }}</h1>
      <div class="flex flex-1 lg:justify-end">
        <div class="w-full max-w-lg lg:max-w-xs">
          <label for="search" class="sr-only">Enter steamid</label>
          <div class="relative">
            <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
              <MagnifyingGlassIcon class="h-5 w-5 text-gray-400" aria-hidden="true"/>
            </div>
            <form @submit.prevent="router.push(`/stats/${query}`)">
              <input
                  v-model="query"
                  id="search" name="search"
                  class="block w-full rounded-md border border-gray-300 bg-white py-2 pl-10 pr-3 leading-5 placeholder-gray-500 focus:border-indigo-500 focus:placeholder-gray-400 focus:outline-none focus:ring-1 focus:ring-indigo-500 sm:text-sm"
                  placeholder="Enter steamid" type="search"/>
              <input type="submit" hidden/>
            </form>
          </div>
        </div>
      </div>
    </div>

    <h2 class="text-xl tracking-tight text-gray-900 sm:text-3xl py-4 sm:py-6 lg:py-8">
      Averages over last 30 games
    </h2>
    <AvgStats :public-avg="stats.publicAvg" :comp-avg="stats.compAvg"/>

    <h2 class="text-xl tracking-tight text-gray-900 sm:text-3xl py-4 sm:py-6 lg:py-8">
      Last 30 comp games
    </h2>
    <LastGames :games="stats?.compGames"/>

    <h2 class="text-xl tracking-tight text-gray-900 sm:text-3xl py-4 sm:py-6 lg:py-8">
      Last 30 public games
    </h2>
    <LastGames :games="stats?.publicGames"/>
  </div>
  <EmptyStats v-else :error="true"/>
</template>

