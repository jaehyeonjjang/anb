<template>
  <Title title="기본 정보" />

  <el-tabs v-model="data.menu">


    <el-tab-pane label="기본정보" name="basic">
      <BasicInsert />
    </el-tab-pane>
    <el-tab-pane label="참여 기술자" name="technician">  
      <TechnicianInsert />
    </el-tab-pane>
    <el-tab-pane label="일반 현황" name="normal">  
      <NormalInsert />
    </el-tab-pane>
    <el-tab-pane label="개요" name="outline">  
      <OutlineInsert />
    </el-tab-pane>
    <el-tab-pane label="용도 현황" name="usagefloor">  
      <UsagefloorInsert />
    </el-tab-pane>
    <el-tab-pane label="기존 점검" name="past">  
      <PastInsert />
    </el-tab-pane>
    <el-tab-pane label="정기/정밀 전환" name="change">  
      <ChangeInsert />
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
  menu: 'basic'  
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
