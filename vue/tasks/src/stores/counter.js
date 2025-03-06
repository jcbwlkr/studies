import { ref, reactive, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {

  const obj = reactive({
    name: 'Jacob',
    score: 1,
  })

  function reset() {
    console.log("resetting", obj)
    Object.assign(obj, {})
  }

  function increment() {
    obj.score++
  }

  return { obj, reset, increment }
})
