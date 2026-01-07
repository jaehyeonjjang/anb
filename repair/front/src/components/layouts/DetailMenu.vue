<template>  
  <el-menu @select="clickMenu" :collapse="false" style="padding-top:10px;">
    <el-menu-item index="6">
      <template #title>
        <el-icon><Document /></el-icon>기본 정보
      </template>
    </el-menu-item>

    <el-menu-item index="1">
      <template #title>
        <el-icon><User /></el-icon>참여 기술자
      </template>
    </el-menu-item>

    <el-menu-item index="2">
      <template #title>
        <el-icon><Tickets /></el-icon>일반 현황
      </template>
    </el-menu-item>

    <el-menu-item index="3">
      <template #title>
        <el-icon><OfficeBuilding /></el-icon>시설물 현황
      </template>
    </el-menu-item>

    <el-menu-item index="4">
      <template #title>
        <el-icon><Reading /></el-icon>개요
      </template>
    </el-menu-item>

    <el-menu-item index="5">
      <template #title>
        <el-icon><Calendar /></el-icon>점검 수행일정
      </template>
    </el-menu-item>

    <el-menu-item index="7">
      <template #title>
        <el-icon><SetUp /></el-icon>점검 일반사항
      </template>
    </el-menu-item>

    <el-menu-item index="8">
      <template #title>
        <el-icon><School /></el-icon>건축물 구조상태
      </template>
    </el-menu-item>

    <el-menu-item index="9">
      <template #title>
        <el-icon><DataBoard /></el-icon>용도현황 (층구분)
      </template>
        </el-menu-item>    
    


  </el-menu>
  
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { Apt, Detail } from "~/models"
import { util }  from "~/global"

const store = useStore()
const route = useRoute()

const clickMenu = async (key: string, keyPath: string[]) => {
  const apt = route.params.apt
  const id = route.params.id

  if (key == '6') {
    router.push(`/${apt}/detail/${id}/detail`)
  } else if (key == '1') {
    router.push(`/${apt}/detail/${id}/technician`)
  } else if (key == '2') {
    router.push(`/${apt}/detail/${id}/normal`)
  } else if (key == '3') {
      router.push(`/${apt}/detail/${id}/current`)
  } else if (key == '4') {
    router.push(`/${apt}/detail/${id}/outline`)
  } else if (key == '5') {
    router.push(`/${apt}/detail/${id}/schedule`)
  } else if (key == '7') {
    router.push(`/${apt}/detail/${id}/blueprint`)
  } else if (key == '8') {
    router.push(`/${apt}/detail/${id}/struct`)
  } else if (key == '9') {
    router.push(`/${apt}/detail/${id}/usagefloor`)
  }
  /*
     } else if (key == '9') {
     let res = await Detail.get(id)
     let detail = res.item

     res = await Apt.get(detail.apt)
     let apt = res.item
     
     const url = '/api/download/detail/' + id
     const filename = `정밀점검-${detail.name}.xlsx`

     util.download(store, url, filename)      
     }
   */
}

</script>
