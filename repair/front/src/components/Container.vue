<template>
  <div :style="{height: height(props.height)}">
    <slot></slot>
  </div>
  
</template>

<script setup lang="ts">

import { reactive } from "vue"
import { util, size }  from "~/global"

const props = defineProps({
  height: Number
})

const { width, height } = size();

</script>
