<script setup>
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-8">
    <div class="max-w-6xl mx-auto">
      <RouterView :key="$route.fullPath" />
    </div>
  </div>
</template>