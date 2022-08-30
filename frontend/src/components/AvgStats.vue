<script setup>
import {toRef} from "vue";

const props = defineProps(['avg'])

const avg = toRef(props, 'avg')
</script>

<template>
  <div v-if="avg">
    <dl class="rounded-lg bg-white shadow-lg sm:grid sm:grid-cols-5">
      <div class="flex flex-col border-b border-gray-100 p-6 text-center sm:border-0 sm:border-r">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Kills</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ avg.kills.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-b border-gray-100 p-6 text-center sm:border-0 sm:border-l sm:border-r">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Deaths</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ avg.deaths.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-gray-100 p-6 text-center sm:border-0 sm:border-l">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">K/D</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ avg.KDRatio.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-gray-100 p-6 text-center sm:border-0 sm:border-l">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Kills per minute</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ avg.killsPerMinute.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-gray-100 p-6 text-center sm:border-0 sm:border-l">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Deaths per minute</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ avg.deathsPerMinute.toFixed(1) || '0.0' }}
        </dd>
      </div>
    </dl>
  </div>
</template>