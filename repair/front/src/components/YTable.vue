<template>
  <table class="table" :style="props.style">
    <slot />
  </table>
  
</template>
<script lang="ts">
export default {
  name: 'y-table',
  inheritAttrs: false,
  customOptions: {}
}
</script>
<script setup lang="ts">
const props = defineProps({
  style: Object,
})
</script>
<style>

.table {
  width: 100%;
  border: 1px solid #dddddd;
  border-collapse : collapse;
}
</style>
