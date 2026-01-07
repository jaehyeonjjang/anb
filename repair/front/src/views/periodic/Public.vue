<template>
  <Title title="공중이 이용하는 부위" />

  <el-tabs v-model="data.menu">


    <el-tab-pane label="추락방지시설" name="fall">
      <FallInsert />
    </el-tab-pane>

    <el-tab-pane label="도로포장" name="road">  
      <RoadInsert />
    </el-tab-pane>

    <el-tab-pane label="도로부 신축 이음부" name="join">  
      <JointInsert />
    </el-tab-pane>

    <el-tab-pane label="환기구 등의 덮개" name="vent">  
      <VentInsert />      
    </el-tab-pane>
    
  </el-tabs>


</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const data = reactive({
  id: 0,
  apt: 0,
  menu: 'fall'  
})

async function initData() {
}

async function getItems() {
}

onMounted(async () => {
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  await initData()
  await getItems()
})

</script>
