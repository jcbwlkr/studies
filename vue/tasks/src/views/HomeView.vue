<script setup>
import {ref, reactive} from 'vue';
import { storeToRefs } from 'pinia'

import { useCounterStore } from '@/stores/counter';

const store = useCounterStore();

const { obj } = storeToRefs(store);
const { increment, reset } = store;

const tetrisWins = reactive({
  'Jenn': 123,
  'Jacob': 39,
});

const players = ref(['Jenn', 'Jacob'])

</script>

<template>
  <header>
    <p>
      Score {{ obj.score }}
    </p>
    <button @click=increment>Add 1</button>
    <button @click=reset>Reset</button>
  </header>
  <main>
    <h1>Tetris Wins</h1>
    <pre class="scoreboard">{{ tetrisWins }}</pre>

    <label v-for="player in players" :key=player>
      {{ player }}

      <input type=number v-model="tetrisWins[player]" />
    </label>
  </main>
</template>

<style>
.scoreboard {
  font-size: 2em;
}

label {
  font-size: 1.5em;
  display:block;
}
input {
  font-size: 1em;
}
</style>
