<template>
  <div style="text-align:left;font-size:16px;font-weight:bold;margin-top:20px;margin-bottom:5px;margin-left:3px;">
    {{title}}
  </div>
  
</template>

<script setup lang="ts">

import { reactive } from "vue"

const props = defineProps({
  title: String
})

</script>
