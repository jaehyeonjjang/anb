<template>  
  <el-menu @select="clickMenu" :collapse="false" style="padding-top:10px;">
    <el-menu-item index="1">
      <template #title>
        <el-icon><OfficeBuilding /></el-icon>장기수선
      </template>
    </el-menu-item>

    <el-menu-item index="3">
      <template #title>
        <el-icon><CollectionTag /></el-icon>공사종별
      </template>
    </el-menu-item>

    <el-menu-item index="4">
      <template #title>
        <el-icon><Filter /></el-icon>규격
      </template>
    </el-menu-item>

    <el-menu-item index="11">
      <template #title>
        <el-icon><ReadingLamp /></el-icon>검토사항
      </template>
    </el-menu-item>

    <el-menu-item index="13">
      <template #title>
        <el-icon><Monitor /></el-icon>항목검토
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
import axios from 'axios'

const store = useStore()

const data = reactive({  
  height: 0
})

const totalRef = ref<InstanceType<typeof Total>>()

function setWindowSize() {
  data.height = (window.innerHeight - 800) + 'px'
}

onMounted(() => {  
  setWindowSize()

  window.addEventListener('resize', setWindowSize)  
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
})


const clickMenu = async (key: string, keyPath: string[]) => {
  if (key == '1') {
    router.push('/management/repair/repair')
  } else if (key == '3') {
    router.push('/management/repair/category')
  } else if (key == '4') {
      router.push('/management/repair/standard')
  } else if (key == '11') {
    router.push('/management/repair/reviewcontentbasic')    
  } else if (key == '13') {
    router.push('/management/repair/reviewbasic')
  } else if (key == '6') {
    router.push('/management/repair/area')    
  }
}

</script>
