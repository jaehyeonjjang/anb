<template>  
  <el-menu @select="clickMenu" :collapse="false" style="padding-top:10px;">
    <el-menu-item index="40">
      <template #title>
        <el-icon><Document /></el-icon>기존 자료 변환
      </template>
    </el-menu-item>
    
    <el-menu-item index="6">
      <template #title>
        <el-icon><Document /></el-icon>기본 정보
      </template>
    </el-menu-item>

    <el-menu-item index="5">
      <template #title>
        <el-icon><Picture /></el-icon>사진자료
      </template>
    </el-menu-item>

    

    <el-menu-item index="11">
      <template #title>
        <el-icon><Notebook /></el-icon>결함현황표
      </template>
    </el-menu-item>

    <el-menu-item index="12">
      <template #title>
        <el-icon><DocumentChecked /></el-icon>점검표
      </template>
    </el-menu-item>

    <el-menu-item index="13">
      <template #title>
        <el-icon><Files /></el-icon>변경사항
      </template>
    </el-menu-item>

    <el-menu-item index="21">
      <template #title>
        <el-icon><DataBoard /></el-icon>공중이 이용하는 부위
      </template>
    </el-menu-item>
    

    <el-menu-item index="15">
      <template #title>
        <el-icon><OfficeBuilding /></el-icon>외벽 마감제
      </template>
    </el-menu-item>

    <el-menu-item index="16">
      <template #title>
        <el-icon><Paperclip /></el-icon>부대 점검사항
      </template>
    </el-menu-item>

    <el-menu-item index="17">
      <template #title>
        <el-icon><Monitor /></el-icon>종합의견
      </template>
    </el-menu-item>

    <!--
    <el-menu-item index="18">
      <template #title>
        <el-icon><ReadingLamp /></el-icon>발생원인 분석
      </template>
    </el-menu-item>
    -->
    
    <el-menu-item index="19">
      <template #title>
        <el-icon><Tickets /></el-icon>시설물 관리대장
      </template>
    </el-menu-item>

    <el-menu-item index="20">
      <template #title>
        <el-icon><Coin /></el-icon>기타참고자료
      </template>
    </el-menu-item>

    <!-- <el-menu-item index="3">
         <template #title>
         <el-icon><OfficeBuilding /></el-icon>시설물 현황
         </template>
         </el-menu-item>

         <el-menu-item index="4">
         <template #title>
         <el-icon><Reading /></el-icon>개요
         </template>
         </el-menu-item>

         <el-menu-item index="45">
         <template #title>
         <el-icon><Calendar /></el-icon>점검 수행일정
         </template>
         </el-menu-item>

         <el-menu-item index="7">
         <template #title>
         <el-icon><SetUp /></el-icon>점검 일반사항
         </template>
         </el-menu-item>

         <el-menu-item index="48">
         <template #title>
         <el-icon><School /></el-icon>건축물 구조상태
         </template>
         </el-menu-item>

         <el-menu-item index="49">
         <template #title>
         <el-icon><DataBoard /></el-icon>용도현황 (층구분)
         </template>
         </el-menu-item>     -->
    
    <el-menu-item index="30">
      <template #title>
        <el-dropdown>      
          <span class="el-dropdown-link" style="color:#333333;font-size:12px;">
            <el-icon><Printer /></el-icon>출력물
            <el-icon class="el-icon--right" style="font-size:12px;">
              <arrow-down />
            </el-icon>
          </span>
      <template #dropdown>
        <el-dropdown-menu v-if="data.periodic.category == 1">          
          <el-dropdown-item style="font-size:12px;" @click="clickReport(-1)">전체 파일 다운로드 (ZIP 압축)</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" divided @click="clickReport(0)">00.본보고서</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReport(1)">01.과업지시서</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReport(2)">02.외관조사망도</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReport(3)">03.시설물관리대장</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReport(4)">04.사진 자료</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReport(5)">05.사전자료일체 및 기타참고자료</el-dropdown-item>
        </el-dropdown-menu>
        <el-dropdown-menu v-if="data.periodic.category == 2">          
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(-1)">전체 파일 다운로드 (ZIP 압축)</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" divided @click="clickReportDetail(0)">00.본보고서</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(1)">01.과업지시서</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(2)">02.외관조사망도</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(3)">03.측정시험결과표</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(4)">04.상대평가 결과자료</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(5)">05.시설물관리대장</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(6)">06.사진 자료</el-dropdown-item>
          <el-dropdown-item style="font-size:12px;" @click="clickReportDetail(7)">07.사전자료일체</el-dropdown-item>          
        </el-dropdown-menu>
      </template>
        </el-dropdown>
        
      </template>
    </el-menu-item>
    
  </el-menu>
  
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { Apt, Periodic } from "~/models"
import { util }  from "~/global"
import axios from 'axios'

const store = useStore()
const route = useRoute()

const data = reactive({
  id: 0,
  periodic: {
    category: 1
  }
})

function download(url: string, filename: string) {
  axios.get(import.meta.env.VITE_REPORT_URL + url, {
    responseType: 'blob',
    headers: {
      Authorization: 'Bearer ' + store.state.token
    }
  }).then(response => {
    const url = window.URL.createObjectURL(new Blob([response.data]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', filename);
    document.body.appendChild(link)
    link.click()
    
    util.info('다운로드가 완료되었습니다')
    util.loading(false)
  }).catch(exception => {
    util.error('다운로드가 실패하였습니다')
    util.loading(false)
  });        
}

const clickMenu = async (key: string, keyPath: string[]) => {
  const apt = route.params.apt
  const id = route.params.id

  if (key == '6') {
    router.push(`/${apt}/periodic/${id}/periodic`)  
  } else if (key == '3') {
    router.push(`/${apt}/periodic/${id}/dong`)  
  } else if (key == '5') {
    router.push(`/${apt}/periodic/${id}/image`)
  } else if (key == '11') {
    router.push(`/${apt}/periodic/${id}/data`)
  } else if (key == '12') {
    router.push(`/${apt}/periodic/${id}/check`)
  } else if (key == '13') {
    router.push(`/${apt}/periodic/${id}/result`)  
  } else if (key == '15') {
    router.push(`/${apt}/periodic/${id}/outerwall`)
  } else if (key == '7') {
    router.push(`/${apt}/periodic/${id}/vent`)    
  } else if (key == '16') {
    router.push(`/${apt}/periodic/${id}/incidental`)
  } else if (key == '17') {
    router.push(`/${apt}/periodic/${id}/opinion`)
  } else if (key == '18') {
    router.push(`/${apt}/periodic/${id}/cause`)
  } else if (key == '19') {
    router.push(`/${apt}/periodic/${id}/managebook`)
  } else if (key == '20') {
    router.push(`/${apt}/periodic/${id}/etc`)
  } else if (key == '21') {
    router.push(`/${apt}/periodic/${id}/public`)    
  } else if (key == '45') {
    router.push(`/${apt}/periodic/${id}/schedule`)    
  } else if (key == '48') {
    router.push(`/${apt}/periodic/${id}/struct`)
  } else if (key == '49') {
    router.push(`/${apt}/periodic/${id}/usagefloor`)
  } else if (key == '30') {
    util.loading(true)
    
    let res = await Apt.get(apt)
    let aptItem = res.item
    
    res = await Periodic.get(id)
    let periodic = res.item

    const url = '/api/download/periodic/' + id
    const filename = `${aptItem.name} ${periodic.name}.zip`    

    download(url, filename)
  } else if (key == '40') {
    router.push(`/${apt}/periodic/${id}/convert`)
  }
}

async function clickReport(key) {
  util.loading(true)
  
  const apt = route.params.apt
  const id = route.params.id  
  
  let res = await Apt.get(apt)
  let aptItem = res.item

  let url = ''
  let filename = ''
  if (key == -1) {
    res = await Periodic.get(id)
    let periodic = res.item
    
    url = '/api/download/periodic/' + id
    filename = `${aptItem.name} ${periodic.name}.zip`
  } else if (key == 0) {
    url = '/api/download/periodic0/' + id
    filename = `00.본보고서-${aptItem.name}.hml`    
  } else if (key == 1) {
    url = '/api/download/periodic1/' + id
    filename = `01.과업지시서-${aptItem.name}.hml`
  } else if (key == 2) {
    url = '/api/download/periodic2/' + id
    filename = `02.외관조사망도-${aptItem.name}.hml`
  } else if (key == 3) {
    url = '/api/download/periodic3/' + id
    filename = `03.시설물관리대장-${aptItem.name}.hml`
  } else if (key == 4) {
    url = '/api/download/periodic4/' + id
    filename = `04.사진자료-${aptItem.name}.hml`
  } else if (key == 5) {
    url = '/api/download/periodic5/' + id
    filename = `05.사전자료일체 및 기타참고자료-${aptItem.name}.hml`    
  }

  download( url, filename)
}

async function clickReportDetail(key) {
  util.loading(true)
  
  const apt = route.params.apt
  const id = route.params.id  
  
  let res = await Apt.get(apt)
  let aptItem = res.item

  let url = ''
  let filename = ''
  if (key == -1) {
    res = await Periodic.get(id)
    let periodic = res.item
    
    url = '/api/download/detail/' + id
    filename = `${aptItem.name} ${periodic.name}.zip`
  } else if (key == 0) {
    url = '/api/download/detail0/' + id
    filename = `00.본보고서-${aptItem.name}.hml`    
  } else if (key == 1) {
    url = '/api/download/detail1/' + id
    filename = `01.과업지시서-${aptItem.name}.hml`
  } else if (key == 2) {
    url = '/api/download/detail2/' + id
    filename = `02.외관조사망도-${aptItem.name}.hml`
  } else if (key == 3) {
    url = '/api/download/detail3/' + id
    filename = `03.시설물관리대장-${aptItem.name}.hml`
  } else if (key == 4) {
    url = '/api/download/detail4/' + id
    filename = `04.상대평가결과자료-${aptItem.name}.hml`
  } else if (key == 5) {
    url = '/api/download/detail5/' + id
    filename = `05.시설물관리대장-${aptItem.name}.hml`  
  } else if (key == 6) {
    url = '/api/download/deatil6/' + id
    filename = `06.사진자료-${aptItem.name}.hml`
  } else if (key == 7) {
    url = '/api/download/detail7/' + id
    filename = `07.사전자료일체-${aptItem.name}.hml`    
  }

  download( url, filename)
}

onMounted(async () => {
  data.id = util.getInt(route.params.id)

  const res = await Periodic.get(data.id)
  data.periodic = res.item  
})
</script>
<style>
.el-dropdown-menu__item {
  font-size:12px;
  font-weight: normal;
}
</style>
