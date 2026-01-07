<template>
  <Title title="기타참고자료" />

  <el-tabs v-model="data.menu">


    <el-tab-pane label="자료 보유 현황" name="keep">
      <KeepInsert />
    </el-tab-pane>

    <el-tab-pane label="사진 자료" name="image">  
      <PastimageInsert />
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
  menu: 'keep'  
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
