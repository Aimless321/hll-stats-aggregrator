<script setup>
import {ArrowTopRightOnSquareIcon} from '@heroicons/vue/20/solid'
import {toRef} from "vue";

const props = defineProps(['games']);
const games = toRef(props, 'games');
const rconUrl = import.meta.env.VITE_APP_RCON_URL;

function getAllWeapons(game) {
  const weapons = Object.entries(game.weapons);
  weapons.sort();

  return weapons;
}
</script>

<template>
  <div v-if="games" class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 sm:-mx-6 md:mx-0 md:rounded-lg">
    <table class="min-w-full divide-y divide-gray-300">
      <thead class="bg-gray-50">
      <tr>
        <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">Date</th>
        <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">Playtime</th>
        <th scope="col" class="hidden px-3 py-3.5 text-left text-sm font-semibold text-gray-900 lg:table-cell">
          Kills
        </th>
        <th scope="col" class="hidden px-3 py-3.5 text-left text-sm font-semibold text-gray-900 sm:table-cell">
          Deaths
        </th>
        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">K/D</th>
        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Kill Streak</th>
        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">KPM</th>
        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">DPM</th>
        <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Weapons</th>
      </tr>
      </thead>
      <tbody class="divide-y divide-gray-200 bg-white">
      <tr v-for="game in games" :key="game.date">
        <td class="px-3 py-4 text-sm text-gray-900">
          {{
            new Date(game.date).toLocaleDateString()
          }}
        </td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ (game.gameLength / 60).toFixed(1) }} minutes</td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ game.kills }}</td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ game.deaths }}</td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ game.KDRatio }}</td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ game.killStreak }}</td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ game.killsPerMinute }}</td>
        <td class="px-3 py-4 text-sm text-gray-900">{{ game.deathsPerMinute }}</td>
        <td class="px-3 py-4 text-sm text-gray-900">
          <ol>
            <li v-for="weapon in getAllWeapons(game)">
              {{ weapon[0] }}: {{ weapon[1] }}
            </li>
          </ol>
        </td>
        <td>
          <a :href="`${rconUrl}/#/gamescoreboard/${game.gameId}`" target="_blank"
             class="text-blue-500 flex items-center gap-1">
            Full stats
            <ArrowTopRightOnSquareIcon class="w-4 h-4"/>
          </a>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
th, td {
  @apply text-center
}
</style>