<template>
  <td class="td" :style="props.style" :rowspan="props.rowspan" :colspan="props.colspan">    
    <slot />
  </td>
  
</template>
<script lang="ts">
export default {
  name: 'y-td',
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

.td {
  border: 1px solid #dddddd;
  border-collapse : collapse;
  background: #fafafa;
  font-size: 12px;
  padding:5px 5px;
}
</style>
