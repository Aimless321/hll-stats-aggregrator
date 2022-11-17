<template>
  <div v-if="data.game">
    <h1 class="text-xl text-gray-900 mb-6">{{ data.game.name }} - {{ data.game.date }}</h1>
    <div class="flex justify-between">
      <h2 class="mb-4">Squad composition</h2>
      <button type="button"
              @click="splitClans"
              class="inline-flex items-center rounded border border-transparent bg-indigo-600 px-2.5 py-1.5 text-xs font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
        Split Clans
      </button>
    </div>

    <div class="grid grid-cols-6 gap-2.5 mb-4">
      <div
          class="squad-block"
          @dragover="$event.preventDefault()"
          @dragenter="$event.preventDefault()"
          @dragleave="$event.preventDefault()"
          @drop="movePlayerTo($event, data)"
      >
        <label for="players" class="block text-sm font-medium text-gray-700">Players</label>
        <select :size="data.players.length+1" id="players"
                class="mt-1 block rounded-md border-gray-300 py-2 pl-3 pr-3 bg-none text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
          <option v-for="player in data.players" :key="player.steamId"
                  @dragstart="dragPlayerFrom($event, player, '_game')"
                  :value="player.steamId" draggable="true">
            {{ player.name }}
          </option>
        </select>
      </div>

      <div
          v-for="(squad, index) in data.game.squads" :key="index"
          class="squad-block sticky top-32"
          @dragover="$event.preventDefault()"
          @dragenter="$event.preventDefault()"
          @dragleave="$event.preventDefault()"
          @drop="movePlayerTo($event, squad)"
      >
        <XMarkIcon @click="removeSquad(squad)" class="cursor-pointer h-4 w-4 absolute top-0 right-0 m-2"/>
        <label class="block text-sm font-medium text-gray-700">{{ squad.name }}</label>
        <select :size="squad.players.length+1"

                class="mt-1 block rounded-md border-gray-300 py-2 pl-3 pr-3 bg-none text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
          <option v-for="player in squad.players" :key="player.steamId"
                  @dragstart="dragPlayerFrom($event, player, squad.name)"
                  :value="player.steamId" draggable="true">
            {{ player.name }}
          </option>
        </select>
      </div>

      <form class="squad-block">
        <label for="email" class="block text-sm font-medium text-gray-700">Name</label>
        <div class="mt-1">
          <input
              v-model="name" required
              type="email" name="email" id="email"
              class="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              placeholder="The Circle"/>
        </div>
        <button type="submit"
                @click="addSquad"
                class="items-center rounded border border-transparent bg-indigo-100 px-2.5 py-1.5 text-xs font-medium text-indigo-700 hover:bg-indigo-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
          Add squad
        </button>
      </form>
    </div>
    <button
        v-if="changesMade"
        type="button"
        @click="saveSquads"
        class="inline-flex w-full justify-center items-center rounded-full border border-transparent bg-green-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2">
      Save squads
    </button>

    <h2 class="mb-4 mt-6">Stats</h2>
    <div class="flex gap-2.5">
      <div
          v-for="(squad, index) in data.game.squads" :key="index"
          class="squad-block"
      >
        <label class="block text-sm font-medium text-gray-700">{{ squad.name }}</label>
        <table class="min-w-full divide-y divide-gray-300">
          <thead class="">
          <tr>
            <th scope="col" class="pr-2 text-left text-sm font-semibold text-gray-900">Name</th>
            <th scope="col" class="px-2 py-3.5 text-center text-sm font-semibold text-gray-900">K</th>
            <th scope="col" class="px-2 text-center text-sm font-semibold text-gray-900">D</th>
            <th scope="col" class="px-1 text-center text-sm font-semibold text-gray-900">K/D</th>
            <th scope="col" class="px-1 text-center text-sm font-semibold text-gray-900">KPM</th>
          </tr>
          </thead>
          <tbody class="divide-y divide-gray-200 bg-white">
          <tr v-for="player in squad.players" :key="player.steamId">
            <td class="whitespace-nowrap pr-2 text-sm font-medium text-gray-900">{{ player.name }}</td>
            <td class="whitespace-nowrap px-1 py-3 text-sm text-gray-500 text-center">{{ player.kills }}</td>
            <td class="whitespace-nowrap text-sm text-gray-500 text-center">{{ player.deaths }}</td>
            <td class="whitespace-nowrap text-sm text-gray-500 text-center">{{ player.KDRatio }}</td>
            <td class="whitespace-nowrap text-sm text-gray-500 text-center">{{ player.killsPerMinute }}</td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import {XMarkIcon} from '@heroicons/vue/20/solid'
import {reactive, ref} from "vue";
import {useRoute, useRouter} from "vue-router";

const apiUrl = import.meta.env.VITE_API_BASE_URL;
const router = useRouter();
const route = useRoute();

const statsUrl = `${apiUrl}/external/${route.params.gameId}`;


const data = reactive({
  game: null,
  players: null,
});
const name = ref("");
const changesMade = ref(false);

fetch(statsUrl, {credentials: "include"})
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
    .then(apiData => {
      Object.assign(data, apiData);

      const assignedPlayers = [];
      data.game.squads.forEach((squad, squadId) => {
        squad.players.forEach((player, playerIndex) => {
          const allPlayerData = data.players.find(dataPlayer => dataPlayer.steamId === player.steamId);

          data.game.squads[squadId].players[playerIndex] = {
            ...allPlayerData
          };

          assignedPlayers.push(player.steamId);
        })
      });

      data.players = data.players.filter(player => !assignedPlayers.includes(player.steamId))
    })
    .catch((error) => {
      console.error('Error:', error);
    });

function splitClans() {
  const knownClans = {
    'The Circle': ['◯ |', 'Ⓡ |'],
    'GOF': ['//'],
    'DC': ['[DC]', '[DC*]'],
    'Phx': ['phx.', 'phx*.'],
    '82AD': ['[82AD]', '[82ADr]'],
    'ARC': ['[ARC]'],
    '126': ['[126]'],
    'Raptors': ['|||™'],
    'HLL.PL': ['HLL.PL |'],
    'StDb': ['[StDb]'],
    'RED': ['[RED]'],
    'TL': ['-TL-'],
    'CoRe': ['CoRe |'],
    'WTH': ['(WTH)', '[WTH]'],
    '116': ['l|l™', '🅶 I|I'],
    '501.ES': ['501.es |'],
    'BC': ['[BC]', '[29SD]'],
    'CM': ['CM |'],
  }

  for (const [clanName, tags] of Object.entries(knownClans)) {
    const clanPlayers = data.players.filter(player => tags.some(tag => player.name.startsWith(tag)));

    if (clanPlayers.length === 0) {
      continue;
    }

    changesMade.value = true;

    data.game.squads.push({
      name: clanName,
      players: clanPlayers,
    });

    data.players = data.players.filter(player => !clanPlayers.some(clanPlayer => clanPlayer.steamId === player.steamId))
  }
}

function addSquad() {
  if (!name.value || name.value === "_game") {
    return;
  }

  changesMade.value = true;

  data.game.squads.push({
    name: name.value,
    players: [],
  });

  name.value = "";
}

function dragPlayerFrom(event, player, from) {
  event.dataTransfer.effectAllowed = 'move';
  event.dataTransfer.setData('from', from);
  event.dataTransfer.setData('player', JSON.stringify(player));
}

function movePlayerTo(event, to) {
  event.stopPropagation();

  changesMade.value = true;

  const fromTarget = event.dataTransfer.getData('from');
  const from = fromTarget === '_game' ? data : data.game.squads.find(squad => squad.name === fromTarget);

  const player = JSON.parse(event.dataTransfer.getData('player'));

  from.players = from.players.filter((gamePlayer) => gamePlayer.steamId !== player.steamId);
  from.players = from.players.sort((a, b) => a.kills < b.kills)

  to.players.push(player);
  to.players = to.players.sort((a, b) => a.kills < b.kills)
}

function removeSquad(deleteSquad) {
  data.players.push(...deleteSquad.players);
  data.players = data.players.sort((a, b) => a.kills < b.kills)

  data.game.squads = data.game.squads.filter(squad => squad.name !== deleteSquad.name);
}

function saveSquads() {
  fetch(`${apiUrl}/external/${route.params.gameId}`, {
    credentials: 'include',
    method: 'PUT',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data.game)
  }).then((resp) => {
    if (resp.status === 201) {
      changesMade.value = false;
      return;
    }

    if (resp.status === 401) {
      localStorage.clear();
      router.replace('/');
      return;
    }

    throw new Error('Something went wrong');
  })
      .catch((error) => {
        console.error('Error:', error);
      });
}
</script>

<style scoped>
.squad-block {
  @apply flex flex-col gap-2.5 justify-center shadow rounded-md py-2 px-4 h-fit;
}
</style>