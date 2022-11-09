<template>
  <header class="shadow fixed w-full z-10 bg-white">
    <nav class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8" aria-label="Top">
      <div class="flex w-full items-center justify-between border-b border-indigo-500 py-6 lg:border-none">
        <div class="flex items-center">
          <img class="h-16 w-auto rounded" src="https://cdn.discordapp.com/attachments/1008165425386897533/1012145165856682024/the-circle-large-rotating.gif" alt="" />
          <div class="ml-10 hidden space-x-8 lg:block">
            <router-link v-for="link in navigation" :key="link.name" :to="link.href" class="text-base font-medium text-gray-900 hover:text-gray-400">{{ link.name }}</router-link>
          </div>
        </div>
      </div>
      <div class="flex flex-wrap justify-center space-x-6 py-4 lg:hidden">
        <a v-for="link in navigation" :key="link.name" :href="link.href" class="text-base font-medium text-white hover:text-indigo-50">{{ link.name }}</a>
      </div>
    </nav>
  </header>
</template>

<script setup>
const navigation = [
  { name: 'Stats', href: '/' },
  { name: 'External Events', href: '/import' },
]
</script>