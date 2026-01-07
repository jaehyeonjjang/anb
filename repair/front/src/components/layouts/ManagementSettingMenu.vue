<template>  
  <el-menu @select="clickMenu" :collapse="false" style="padding-top:10px;">
    <el-menu-item index="1">
      <template #title>
        <el-icon><User /></el-icon>사용자 관리
      </template>
    </el-menu-item>
    
    <el-menu-item index="2">
      <template #title>
        <el-icon><Cpu /></el-icon>기술자 관리
      </template>
    </el-menu-item>

    <el-menu-item index="3">
      <template #title>
        <el-icon><Document /></el-icon>입력항목 관리
      </template>
    </el-menu-item>

    <el-menu-item index="4">
      <template #title>
        <el-icon><Money /></el-icon>견적 기본값 관리
      </template>
    </el-menu-item>    

    <el-menu-item index="5">
      <template #title>
        <el-icon><Reading /></el-icon>비교견적업체 관리
      </template>
    </el-menu-item>    

  </el-menu>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'

const store = useStore()
const route = useRoute()

const clickMenu = async (key: string, keyPath: string[]) => {
  if (key == '1') {
    router.push('/management/setting/user')
  } else if (key == '2') {
    router.push('/management/setting/technician')
  } else if (key == '3') {
    router.push('/management/setting/datacategory')
  } else if (key == '4') {
    router.push('/management/setting/estimate')    
  } else if (key == '5') {
    router.push('/management/setting/comparecompany')    
  }
}

</script>
