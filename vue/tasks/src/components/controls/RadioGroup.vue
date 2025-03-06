<script setup>
defineProps(['options', 'modelValue']);
defineEmits(['update:modelValue']);
</script>

<template>
  <fieldset>
    <label
      v-for="option in options"
      :key="option.id"
      :for="'radio-'+option.id"
    >
      <input
        type="radio"
        :id="'radio-' + option.id"
        :value="option.id"
        :checked="modelValue.id === option.id"
        @change="$emit('update:modelValue', option);"
      />
      {{ option.name }}
    </label>
  </fieldset>
</template>

<style scoped>
label {
  min-width: 100px;
  display: inline-block;
}
</style>
