<template>
  <Title title="점검 수행일정" />


  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <y-table>
    <y-tr>
      <y-th>사전조사</y-th>
      <y-td>{{viewDuration(data.detail.prestartdate, data.detail.preenddate)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>현장조사</y-th>
      <y-td>{{viewDuration(data.detail.researchstartdate, data.detail.researchenddate)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>수집자료 및 결과 분석</y-th>
      <y-td>{{viewDuration(data.detail.analyzestartdate, data.detail.analyzeenddate)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>상태평가</y-th>
      <y-td>{{viewDuration(data.detail.ratingstartdate, data.detail.ratingenddate)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>보고서 작성</y-th>
      <y-td>{{viewDuration(data.detail.writestartdate, data.detail.writeenddate)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>보고서 인쇄 및 제출</y-th>
      <y-td>{{viewDuration(data.detail.printstartdate, data.detail.printenddate)}}</y-td>
    </y-tr>
  </y-table>

  <el-dialog
    v-model="data.visible"
    title="점검 수행일정 수정"
    :before-close="handleClose"
  >


  <y-table>
    <y-tr>
      <y-th>사전조사</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.prestartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.preenddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>현장조사</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.researchstartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.researchenddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>수집자료 및 결과 분석</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.analyzestartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.analyzeenddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>상태평가</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.ratingstartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.ratingenddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>보고서 작성</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.writestartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.writeenddate" /> </y-td>
    </y-tr>
    <y-tr>
      <y-th>보고서 인쇄 및 제출</y-th>
      <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.printstartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.printenddate" /> </y-td>
    </y-tr>
    
    
  </y-table>

  <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">저장</el-button>
    </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Detail } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,  
  prestartdate: '',
  preenddate: '',
  researchstartdate: '',
  researchenddate: '',
  analyzestartdate: '',
  analyzeenddate: '',
  ratingstartdate: '',
  ratingenddate: '',
  writestartdate: '',
  writeenddate: '',
  printstartdate: '',
  printenddate: '',
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  detail: util.clone(item),
  visible: false  
})

async function initData() {
}

async function getItems() {
  let res = await Detail.get(data.id)
  data.detail = res.item  
}

async function clickUpdate(pos) {
  let res = await Detail.get(data.id)
  const item = res.item
  
  data.item = item
  data.visible = true
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.prestartdate = util.convertDBDate(item.prestartdate)
  item.preenddate = util.convertDBDate(item.preenddate)
  item.researchstartdate = util.convertDBDate(item.researchstartdate)
  item.researchenddate = util.convertDBDate(item.researchenddate)
  item.analyzestartdate = util.convertDBDate(item.analyzestartdate)
  item.analyzeenddate = util.convertDBDate(item.analyzeenddate)
  item.ratingstartdate = util.convertDBDate(item.ratingstartdate)
  item.ratingenddate = util.convertDBDate(item.ratingenddate)
  item.writestartdate = util.convertDBDate(item.writestartdate)
  item.writeenddate = util.convertDBDate(item.writeenddate)
  item.printstartdate = util.convertDBDate(item.printstartdate)
  item.printenddate = util.convertDBDate(item.printenddate)
  
  await Detail.update(item)

  util.info('수정되었습니다')

  data.detail = item
  data.visible = false
}

const handleClose = (done: () => void) => {
  /*
     util.confirm('팝업창을 닫으시겠습니까', function() {
     done()
     })
   */

  done()
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

function viewDuration(startdate, enddate) {
  let flag1 = true
  let flag2 = true
  if (util.isNull(startdate) || startdate == '' || startdate == '0000-00-00' || startdate == '1000-01-01') {
    flag1 = false
  }

  if (util.isNull(enddate) || enddate == '' || enddate == '0000-00-00' || enddate == '1000-01-01') {
    flag2 = false
  }

  if (flag1 == true && flag2 == true) {
    return util.viewDate(startdate) + ' ~ ' + util.viewDate(enddate) 
  }

  if (flag1 == true) {
    return util.viewDate(startdate)
  }

  if (flag2 == true) {
    return util.viewDate(enddate)
  }
}

</script>
