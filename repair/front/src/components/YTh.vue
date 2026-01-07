<template>
  <th class="th" :style="props.style" :rowspan="props.rowspan" :colspan="props.colspan">
    <slot />
  </th>
  
</template>
<script lang="ts">
export default {
  name: 'y-th',
  inheritAttrs: false,
  customOptions: {}
}
</script>
<script setup lang="ts">
const props = defineProps({
  style: Object,
  rowspan: String,
  colspan: String
})
</script>
<style>

.th {
  border: 1px solid #dddddd;
  border-collapse : collapse;
  font-size: 12px;
  padding:5px 10px;
  text-align:left;
}
</style>
