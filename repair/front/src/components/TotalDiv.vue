<template>
  <div style="flex: 1; align-self: end; text-align: right;margin-right:5px;">
    Total {{util.money(total)}}
  </div>  
</template>

<script setup lang="ts">

import { util }  from "~/global"

const props = defineProps({
  total: Number
})

</script>
