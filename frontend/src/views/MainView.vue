<template>
  <div v-if="!store.loggedIn" class="grid place-content-center h-screen">
    <DiscordButton/>
  </div>
  <EmptyStats v-else />
</template>

<script setup>
import DiscordButton from "../components/DiscordButton.vue";
import EmptyStats from "../components/EmptyStats.vue";
import {useStore} from '../store'

const store = useStore();
</script>