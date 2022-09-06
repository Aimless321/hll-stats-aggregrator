<script setup>
import {reactive, toRef, toRefs} from "vue";

const props = defineProps(['publicAvg', 'compAvg'])

const loaded = toRef(props, 'publicAvg')
const {publicAvg, compAvg} = toRefs(props);

let data = reactive({compSelected: true, selectedAvg: compAvg});

// sorry 🤮
function switchAvgs(toComp) {
  let evaluate = data.compSelected;
  if (toComp !== undefined) {
    evaluate = toComp;
  }

  if (evaluate) {
    data.selectedAvg = publicAvg;
  } else {
    data.selectedAvg = compAvg;
  }

  if (toComp !== undefined) {
    data.compSelected = !toComp;
  } else {
    data.compSelected = !data.compSelected;
  }
}
</script>

<template>
  <div v-if="loaded">
    <div class="pb-2">
      <div class="sm:hidden">
        <label for="tabs" class="sr-only">Select a tab</label>
        <select @change="switchAvgs()" id="tabs" name="tabs"
                class="block w-full rounded-md border-gray-300 focus:border-indigo-500 focus:ring-indigo-500">
          <option :selected="data.compSelected">Comp</option>
          <option :selected="!data.compSelected">Public</option>
        </select>
      </div>
      <div class="hidden sm:block">
        <nav class="flex space-x-4" aria-label="Tabs">
          <a @click="switchAvgs(false)"
             class="cursor-pointer"
             :class="[data.compSelected ? 'bg-gray-100 text-gray-700' : 'text-gray-500 hover:text-gray-700', 'px-3 py-2 font-medium text-sm rounded-md']"
             :aria-current="data.compSelected ? 'page' : undefined">
            Comp
          </a>
          <a @click="switchAvgs(true)"
             class="cursor-pointer"
             :class="[!data.compSelected ? 'bg-gray-100 text-gray-700' : 'text-gray-500 hover:text-gray-700', 'px-3 py-2 font-medium text-sm rounded-md']"
             :aria-current="!data.compSelected ? 'page' : undefined">
            Public
          </a>
        </nav>
      </div>
    </div>


    <dl class="rounded-lg bg-white shadow-lg sm:grid sm:grid-cols-5">
      <div class="flex flex-col border-b border-gray-100 p-6 text-center sm:border-0 sm:border-r">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Kills</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ data.selectedAvg.kills.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-b border-gray-100 p-6 text-center sm:border-0 sm:border-l sm:border-r">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Deaths</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ data.selectedAvg.deaths.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-gray-100 p-6 text-center sm:border-0 sm:border-l">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">K/D</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ data.selectedAvg.KDRatio.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-gray-100 p-6 text-center sm:border-0 sm:border-l">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Kills per minute</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ data.selectedAvg.killsPerMinute.toFixed(1) || '0.0' }}
        </dd>
      </div>
      <div class="flex flex-col border-t border-gray-100 p-6 text-center sm:border-0 sm:border-l">
        <dt class="order-2 mt-2 text-lg font-medium leading-6 text-gray-500">Deaths per minute</dt>
        <dd class="order-1 text-5xl font-bold tracking-tight text-indigo-600">
          {{ data.selectedAvg.deathsPerMinute.toFixed(1) || '0.0' }}
        </dd>
      </div>
    </dl>
  </div>
</template>