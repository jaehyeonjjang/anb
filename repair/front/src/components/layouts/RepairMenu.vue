<template>  
  <el-menu @select="clickMenu" :collapse="false" style="padding-top:10px;">
    <el-menu-item index="1">
      <template #title>
        <el-icon><Document /></el-icon>개요
      </template>
    </el-menu-item>

    <el-menu-item index="19">
      <template #title>    
        <el-icon><Key /></el-icon>결재
      </template>
    </el-menu-item>

    <el-menu-item index="17">
      <template #title>
        <el-icon><Files /></el-icon>파일
      </template>
    </el-menu-item>
    
    <el-menu-item index="7">
      <template #title>
        <el-icon><House /></el-icon>공동주택
      </template>
    </el-menu-item>

    <el-menu-item index="2">
      <template #title>
        <el-icon><OfficeBuilding /></el-icon>시설물
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

    <el-menu-item index="16">
      <template #title>
        <el-icon><Notebook /></el-icon>총론
      </template>
    </el-menu-item>

    <el-menu-item index="18">
      <template #title>
        <el-icon><Paperclip /></el-icon>예외적 집행
      </template>
    </el-menu-item>    
    
    <el-menu-item index="5">
      <template #title>
        <el-icon><Tickets /></el-icon>사용현황
      </template>
    </el-menu-item>

    <el-menu-item index="6">
      <template #title>
        <el-icon><Reading /></el-icon>세부내역
      </template>
    </el-menu-item>

    <el-menu-item index="12">
      <template #title>
        <el-icon><ReadingLamp /></el-icon>검토사항
      </template>
    </el-menu-item>

    <el-menu-item index="14">
      <template #title>
        <el-icon><Monitor /></el-icon>항목검토
      </template>
    </el-menu-item>

    <el-menu-item index="15">
      <template #title>
        <el-icon><Coin /></el-icon>충당금적립
      </template>
    </el-menu-item>

    <el-menu-item index="20">
      <template #title>
        <el-icon><Money /></el-icon>단가 조정
      </template>
    </el-menu-item>
    
    <el-menu-item index="10">
      <template #title>
        <el-icon><ChatSquare /></el-icon> 상담이력
      </template>
    </el-menu-item>

    <el-menu-item index="8">
      <template #title>
        <el-icon><PieChart /></el-icon>리포트
      </template>
    </el-menu-item>

    
    <el-menu-item index="9">
      <template #title>
        <el-icon><Printer /></el-icon>출력물
      </template>
    </el-menu-item>
  </el-menu>  
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { Apt, Repair } from "~/models"
import { util }  from "~/global"

const store = useStore()
const route = useRoute()

const clickMenu = async (key: string, keyPath: string[]) => {
  const apt = route.params.apt
  const id = route.params.id
  
  if (key == '1') {
    router.push(`/${apt}/repair/${id}/repair`)
  } else if (key == '2') {
    router.push(`/${apt}/repair/${id}/dong`)    
  } else if (key == '3') {
    router.push(`/${apt}/repair/${id}/category`)      
  } else if (key == '4') {    
    router.push(`/${apt}/repair/${id}/standard`)
  } else if (key == '5') {
    router.push(`/${apt}/repair/${id}/history`)
  } else if (key == '6') {
    router.push(`/${apt}/repair/${id}/breakdown`)
  } else if (key == '7') {
    router.push(`/${apt}/repair/${id}/area`)
  } else if (key == '8') {
    router.push(`/${apt}/repair/${id}/report`)
  } else if (key == '10') {
    router.push(`/${apt}/repair/${id}/advice`)  
  } else if (key == '12') {
    router.push(`/${apt}/repair/${id}/reviewcontent`)  
  } else if (key == '14') {
    router.push(`/${apt}/repair/${id}/review`)
  } else if (key == '15') {
    router.push(`/${apt}/repair/${id}/saving`)
  } else if (key == '16') {
    router.push(`/${apt}/repair/${id}/outline`)
  } else if (key == '17') {
    router.push(`/${apt}/repair/${id}/file`)
  } else if (key == '18') {
    router.push(`/${apt}/repair/${id}/exception`)
  } else if (key == '19') {
    router.push(`/${apt}/repair/${id}/approval`)
  } else if (key == '20') {
    router.push(`/${apt}/repair/${id}/adjust`)    
  } else if (key == '9') {
    let res = await Apt.get(apt)
    let aptItem = res.item

    
    const url = '/api/download/report/' + id
    let filename = `장기수선계획-${aptItem.name}.xlsx`
    res = await Repair.get(id)
    if (res.provision == 2) {
      filename = `수선계획-${aptItem.name}.xlsx`
    }

    util.download(store, url, filename)      
  }
}

</script>
