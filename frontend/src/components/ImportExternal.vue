<template>
  <div class="mx-auto max-w-md sm:max-w-3xl min-h-screen grid place-content-center">
    <div>
      <div class="text-center">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"
             xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
        </svg>
        <h2 class="mt-2 text-lg font-medium text-gray-900">Import external game</h2>
      </div>
      <p v-if="error" class="text-red-500">Cannot import game</p>
      <form class="mt-6 sm:flex sm:items-center">
        <div class="relative rounded-md shadow-sm sm:min-w-0 sm:flex-1">
          <input v-model="name" type="text"
                 class="block w-full rounded-md border-gray-300 pr-32 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm mb-1"
                 placeholder="Name"/>
          <input v-model="statsURL" type="text"
                 class="block w-full rounded-md border-gray-300 pr-32 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                 placeholder="Stats URL"/>
        </div>
        <div class="mt-3 sm:mt-0 sm:ml-4 sm:flex-shrink-0">
          <button @click.prevent="importGame"
                  type="submit"
                  class="block w-full rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-center text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
            Import
          </button>
        </div>
      </form>
    </div>
    <div class="mt-10">
      <h3 class="text-sm font-medium text-gray-500">Already imported games</h3>
      <ul role="list" class="mt-4 grid grid-cols-2 gap-4 sm:grid-cols-3">
        <li v-for="(game, id) in games" :key="id">
          <button @click="router.push(`/external/${game.id}`)"
                  type="button"
                  class="group flex w-full items-center justify-between space-x-3 rounded-full border border-gray-300 p-2 text-left shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
            <span class="flex min-w-0 flex-1 items-center space-x-3">
              <span class="block flex-shrink-0"></span>
              <span class="block min-w-0 flex-1">
                <span class="block truncate text-sm font-medium text-gray-900">{{ game.name }}</span>
                <span class="block truncate text-sm font-medium text-gray-500">
                  {{
                    new Date(game.date).toLocaleDateString()
                  }}
                </span>
              </span>
            </span>
            <span class="inline-flex h-10 w-10 flex-shrink-0 items-center justify-center">
              <ChevronRightIcon class="h-5 w-5 text-gray-400 group-hover:text-gray-500" aria-hidden="true"/>
            </span>
          </button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import {ChevronRightIcon} from '@heroicons/vue/20/solid'
import {useRouter} from "vue-router";
import {ref} from "vue";

const apiUrl = import.meta.env.VITE_API_BASE_URL;

const error = ref(false);
const statsURL = ref("");
const name = ref("");

const router = useRouter();
const games = ref({});

fetch(`${apiUrl}/external`, {credentials: 'include'})
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
      games.value = data;
    })
    .catch((error) => {
      console.error('Error:', error);
    });


function importGame() {
  fetch(`${apiUrl}/external/import`, {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    credentials: 'include',
    body: JSON.stringify({
      name: name.value,
      statsURL: statsURL.value,
    })
  })
      .then((resp) => {
        if (resp.status === 200) {
          return resp.json();
        }

        if (resp.status === 401) {
          localStorage.clear();
          router.replace('/');
          return;
        }

        throw new Error('Something went wrong');
      })
      .then(data => {
        router.push(`/external/${data.eventId}`)
      })
      .catch((e) => {
        error.value = true;
        console.error('Error:', e);
      });
}
</script>
