<template>
  <div style="margin-bottom:20px;display:flex;">
    <div style="font-size:20px;font-weight:bold;margin-left:10px;">{{title}}</div>

    <div v-if="subtitle != undefined && subtitle != null && subtitle != ''" style="font-size:14px;color:#409eff;font-weight:bold;margin: 5px 0px 0px 10px;">
      {{subtitle}}
    </div>
  </div>
  
</template>

<script setup lang="ts">

import { reactive } from "vue"
import { ArrowRight } from '@element-plus/icons-vue'

const props = defineProps({
  title: String,
  subtitle: String
})

</script>
