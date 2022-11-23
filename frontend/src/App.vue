<script setup>
import DiscordButton from "./components/DiscordButton.vue";
import Header from "./components/Header.vue";
import {RouterView} from 'vue-router'
import {useStore} from './store'

const store = useStore();
</script>

<template>
  <div v-if="!store.loggedIn" class="grid place-content-center h-screen">
    <DiscordButton/>
  </div>
  <div v-else class="bg-gray-50">
    <Header/>
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-8">
      <div class="max-w-6xl mx-auto">
        <RouterView class="pt-32" :key="$route.fullPath"/>
      </div>
    </div>
  </div>
</template>