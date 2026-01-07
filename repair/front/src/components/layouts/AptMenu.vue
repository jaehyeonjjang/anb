<template>  
  <el-menu @select="clickMenu" :collapse="false" style="padding-top:10px;">
    <el-menu-item index="1">
      <template #title>
        <el-icon><Document /></el-icon>기본 정보
      </template>
    </el-menu-item>

    <el-menu-item index="3">
      <template #title>
        <el-icon><User /></el-icon>사용자
      </template>
    </el-menu-item>
    
    <el-menu-item index="4">
      <template #title>
        <el-icon><OfficeBuilding /></el-icon>시설물
      </template>
    </el-menu-item>    

    <el-menu-item index="2">
      <template #title>
        <el-icon><Files /></el-icon>도면
      </template>
    </el-menu-item>    


    <el-menu-item index="5">
      <template #title>
        <el-icon><Files /></el-icon>순찰 도면
      </template>
    </el-menu-item>

    <el-menu-item index="6">
      <template #title>
        <el-icon><Coin /></el-icon>수량
      </template>
    </el-menu-item>    
    

  </el-menu>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import axios from 'axios'

const store = useStore()
const route = useRoute()

const clickMenu = async (key: string, keyPath: string[]) => {
  const apt = route.params.apt
  
  if (key == '1') {
    router.push(`/${apt}/apt/apt`)
  } else if (key == '2') {
    router.push(`/${apt}/apt/blueprint`)    
  } else if (key == '3') {
    router.push(`/${apt}/apt/user`)
  } else if (key == '4') {
    router.push(`/${apt}/apt/dong`)
  } else if (key == '5') {
    router.push(`/${apt}/apt/patrolblueprint`)
  } else if (key == '6') {
    router.push(`/${apt}/apt/area`)    
  }
}

</script>
