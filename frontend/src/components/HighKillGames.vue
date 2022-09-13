<script setup>
import { ChevronRightIcon } from '@heroicons/vue/20/solid'
import {toRef, toRefs} from "vue";

const props = defineProps(['data'])
const loaded = toRef(props, 'data')

const {data} = toRefs(props);

</script>

<template>
  <div v-if="loaded">
    <h2 class="text-lg tracking-tight text-gray-900 sm:text-xl pb-4">
      High Kill Games (7 days)
    </h2>
    <div class="overflow-hidden bg-white shadow sm:rounded-md">
      <ul role="list" class="divide-y divide-gray-200">
        <li v-for="(row, index) in data" :key="index">
          <a :href="`/stats/${row.steamId}`" class="block hover:bg-gray-50">
            <div class="flex items-center px-4 py-4 sm:px-6">
              <div class="min-w-0 flex-1 sm:flex sm:items-center sm:justify-between">
                <div class="truncate text-sm">
                  <p class="truncate font-medium text-indigo-600">{{ row.name }}</p>
                  <div class="mt-2 flex text-gray-500">
                    <p>
                      {{ row.steamId }}
                    </p>
                  </div>
                </div>
                <div class="mt-4 flex-shrink-0 sm:mt-0 sm:ml-5 text-gray-700">
                  {{ row.kills }} kills
                </div>
              </div>
              <div class="ml-5 flex-shrink-0">
                <ChevronRightIcon class="h-5 w-5 text-gray-400" aria-hidden="true" />
              </div>
            </div>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>