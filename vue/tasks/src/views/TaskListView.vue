<script setup>
import {ref, watch} from 'vue';
import {TaskService} from '@/services/TaskService'
import RadioGroup from '@/components/controls/RadioGroup.vue'

const categories = ref([]);
const selectedCategory = ref({});

const tasks = ref([]);

TaskService.getCategories().then(result => {
  categories.value = result;
  selectedCategory.value = result[0];
})

watch(selectedCategory, (newCategory) => {
  if (!newCategory.id) {
    return;
  }

  TaskService.getTasksByCategory(newCategory).then(result => {
    tasks.value = result;
  })
});

</script>

<template>
  <main>
    <h1>Task List</h1>
    <h2>Showing tasks in category: {{ selectedCategory.name }}</h2>

    <RadioGroup :options="categories" v-model="selectedCategory" />

    <ul v-if="tasks">
      <li
        v-for="task in tasks"
        :key="task.id"
      >
        {{ task.name }}
      </li>
    </ul>
  </main>
</template>
