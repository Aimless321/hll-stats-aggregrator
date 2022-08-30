<script setup>
const apiUrl = import.meta.env.VITE_API_BASE_URL;
const loginUrl = `${apiUrl}/auth/discord`;
</script>

<template>
  <a :href="loginUrl" type="button" class="inline-flex items-center px-6 py-3 border border-transparent shadow-sm text-base font-medium rounded-md text-white bg-blurple focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blurple">
    <img src="../assets/icon_clyde_white_RGB.svg" class="-ml-1 mr-3 h-5 w-5" aria-hidden="true" />
    Sign in with Discord
  </a>
</template>