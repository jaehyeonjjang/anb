<template>
  <Title title="개요" />

  <el-tabs v-model="data.menu">


    <el-tab-pane label="기본 정보" name="basic">

      <div :style="{'height': height(186), 'overflow': 'auto'}">
        <div style="display:flex;justify-content: space-between;margin-top:0px;">
        <PageTitle title="리포트 정보" />
        <el-button size="small" style="margin-top:15px;" type="success" @click="clickInsert(4)">수정</el-button>    
      </div>

            
      <el-table :data="data.report" border style="width: 100%;" :show-header="false" :cell-class-name="cellClassName">
        <el-table-column prop="label" label="" width="240"/>
        <el-table-column prop="value" label="" />        
      </el-table>

      
      <div style="display:flex;justify-content: space-between;margin-top:0px;">
        <PageTitle title="단지 개요" />
        <el-button size="small" style="margin-top:15px;" type="success" @click="clickInsert(1)">수정</el-button>    
      </div>
      
      <el-table :data="data.complex" border style="width: 100%;" :show-header="false" :cell-class-name="cellClassName">
        <el-table-column prop="label" label="" width="240"/>
        <el-table-column prop="value" label="" />              
      </el-table>

      <div style="display:flex;justify-content: space-between;">
        <PageTitle title="일반 개요" />
        <el-button size="small" style="margin-top:15px;" type="success" @click="clickInsert(2)">수정</el-button>    
      </div>
      
      <el-table :data="data.info" border style="width: 100%;" :show-header="false" :cell-class-name="cellClassName">
        <el-table-column prop="label" label="" width="240" />
        <el-table-column prop="value" label="" />              
      </el-table>

      <div style="display:flex;justify-content: space-between;">
        <PageTitle title="건축물 기본현황" />
        <el-button size="small" style="margin-top:15px;" type="success" @click="clickInsert(3)">수정</el-button>    
      </div>
      
      <el-table :data="data.structure" border style="width:100%;margin-bottom:20px;" :show-header="false" :span-method="spanMethod" :cell-class-name="cellClassNameStructure">
        <el-table-column prop="label" label="" width="120" />
        <el-table-column prop="sublabel" label="" width="120" />
        <el-table-column prop="value" label="" />              
      </el-table>
      
      </div>
    </el-tab-pane>


    <el-tab-pane label="엑셀 업로드" name="upload">  
      <div style="text-align:left;margin-bottom:50px;margin-top:10px;">
        <el-upload
          style="float:left;"
          class="upload-demo"
          ref="upload"
          :action="data.upload"
          :headers="headers"
          :limit="1"
          :on-exceed="handleExceed"
          :on-success="handelSuccess"
          :show-file-list="false"
          :auto-upload="true"
        >
          <el-button size="small" type="danger" @click="submitUpload">엑셀 등록</el-button>
          
        </el-upload>

        <div style="margin:4px 0px 0px 10px;float:left;color:red;font-weight:bold;">* 엑셀 등록을 하면 기존에 등록된 모든 자료가 삭제됩니다. 주의해서 사용하세요</div>
        
        <div style="clear:both;"></div>
        <div style="margin-top:10px;">
          <el-checkbox v-model="data.historydel" label="기존 입력한 사용현황 삭제" size="small" style="margin:0px 0px;" />
        </div>
        <div>
          <el-checkbox v-model="data.breakdowndel" label="기존 입력한 세부내역 삭제" size="small" style="margin:0px 0px;" />
        </div>
      </div>

    </el-tab-pane>

    <el-tab-pane label="보조부원장 업로드" name="assistance">
      <div style="text-align:left;margin-bottom:50px;margin-top:10px;">
        <el-upload
          class="upload-demo"
          ref="uploadAssistance"
          :drag="true"
          :action="data.upload"
          :headers="headers"
          :on-exceed="handleExceedAssistance"
          :on-success="handelSuccessAssistance"
          :show-file-list="true"
          :auto-upload="true"
          :multiple="true"
          v-model:file-list="data.assistances"
        >
          
          <el-icon class="el-icon--upload"><upload-filled /></el-icon>
          <div class="el-upload__text" style="font-size:12px;">
            파일을 드래그 하시거나 <em>여기를 클릭하세요</em>
          </div>          
        </el-upload>

      </div>
      <div style="text-align:left;">
        <el-button style="display:block;float:left;" size="small" type="danger" @click="submitUploadAssistance">보조부원장 등록</el-button>

        <div style="margin:4px 0px 0px 10px;float:left;color:red;font-weight:bold;">* 보조부원장 등록을 하면 기존에 등록된 사용내역 및 충당금적립 자료가 삭제됩니다. 주의해서 사용하세요</div>
        
        <div style="clear:both;"></div>
        <div style="margin-top:10px;">
          <el-checkbox v-model="data.historydelassistance" label="기존 입력한 사용현황 삭제" size="small" style="margin:0px 0px;" />
        </div>
      </div>

    </el-tab-pane>
    
    <el-tab-pane label="조정 이력" name="history">
      <RepairhistoryInsert :id="data.aptid" />
    </el-tab-pane>

    <el-tab-pane label="공사종별 변환" name="change">
        <el-button style="display:block;float:left;" size="small" type="danger" @click="clickChange">공사종별 변환</el-button>
        <div style="margin:4px 0px 0px 10px;float:left;color:red;font-weight:bold;">* 공사종별 변환을 하면 신규 수립기준에 맞춰 기존 데이터가 변경되면 복원이 불가능 합니다. 주의해서 사용하세요</div>
    </el-tab-pane>
  </el-tabs>

  <el-dialog
    v-model="data.visibleReport"
    title="리포트 정보 수정"

    :before-close="handleClose"
  >
    <el-form :model="[1]" label-width="100px">
      <el-form-item label="단지명">
        <div style="font-size:16px;font-weight:bold;">{{data.complex[0].value}}</div>
      </el-form-item>

      <el-form-item label="리포트 작성 일자">
        <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.reportdate" placeholder="" />
      </el-form-item>

      <el-form-item label="법조문 종류">        
        <el-radio-group v-model="data.item.provision">
          <el-radio :label="1" size="small">공동주택관리법</el-radio>
          <el-radio :label="2" size="small">집합건물의 소유 및 관리에 관한 법률</el-radio>
        </el-radio-group>        
      </el-form-item>
      
      <el-form-item label="분양 비율">        
        <el-input v-model="data.item.parcelrate" style="width:100px;" />
      </el-form-item>

      <el-form-item label="총 계획년도">        
        <el-input v-model="data.item.planyears" style="width:100px;" />
      </el-form-item>

      <el-form-item label="부분수선 수선율 적용 여부" v-show="false">
        <el-radio-group v-model.number="data.item.calculatetype">
          <el-radio :label="1" size="small">수선율 적용</el-radio>
          <el-radio :label="2" size="small">수선율 적용 안함 (전면교체와 동일한 금액으로 표시)</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="상태" v-if="data.level > 1">        
        <el-radio-group v-model="data.item.status">
          <el-radio :label="1" size="small">진행</el-radio>
          <el-radio :label="2" size="small">마감</el-radio>
        </el-radio-group>        
      </el-form-item>
      
    </el-form>

    <template #footer>
      <el-button size="small" @click="data.visibleReport = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit(4)">저장</el-button>
    </template>
  </el-dialog>

  
  <el-dialog
    v-model="data.visible"
    title="단지 개요 수정"

    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="100px">
      <el-form-item label="단지명">
        <el-input v-model="data.item.name" />
      </el-form-item>

      <el-form-item label="세대수">
        <el-input v-model="data.item.familycount" />
      </el-form-item>

      <el-form-item label="동수">
        <el-input v-model="data.item.flatcount" />
      </el-form-item>

      <el-form-item label="소재지">
        <el-input v-model="data.item.address" />
      </el-form-item>

      <el-form-item :label="complexTitle[4]">
        <el-input v-model="data.item.complex1" />
      </el-form-item>

      <el-form-item :label="complexTitle[5]">
        <el-input v-model="data.item.complex2" /> 
      </el-form-item>

      <el-form-item label="사용검사 일자">
        <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.completiondate" placeholder="" />
      </el-form-item>

      <!--
      <el-form-item label="사용검사 년도">
        <el-input v-model="data.item.completionyear" />
      </el-form-item>

      <el-form-item label="사용검사 월">
        <el-input v-model="data.item.completionmonth" />
      </el-form-item>

      <el-form-item label="사용검사 일">
        <el-input v-model="data.item.completionday" />
      </el-form-item>
      -->
    </el-form>

    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit(1)">저장</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleInfo"
    title="일반 개요 수정"

    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="100px">
      <el-form-item label="단지명">
        <div style="font-size:16px;font-weight:bold;">{{data.complex[0].value}}</div>
      </el-form-item>

      <el-form-item :label="infoTitle[0]">
        <el-input v-model="data.item.info1"  />
      </el-form-item>

      <el-form-item :label="infoTitle[1]">
        <el-input v-model="data.item.info2"  />
      </el-form-item>

      <el-form-item :label="infoTitle[2]">
        <el-input v-model="data.item.info3"  />
      </el-form-item>
      
      <el-form-item :label="infoTitle[3]">
        <el-input v-model="data.item.info4"  />
      </el-form-item>

      <el-form-item :label="infoTitle[4]">
        <el-input v-model="data.item.info5"  />
      </el-form-item>

      <el-form-item :label="infoTitle[5]">
        <el-input v-model="data.item.info6"  />
      </el-form-item>
      
      <el-form-item :label="infoTitle[6]">
        <el-input v-model="data.item.info7"  />
      </el-form-item>

      <el-form-item :label="infoTitle[7]">
        <el-input v-model="data.item.info8"  />
      </el-form-item>

      <el-form-item :label="infoTitle[8]">
        <el-input v-model="data.item.info9"  />
      </el-form-item>
      
      <el-form-item :label="infoTitle[9]">
        <el-input v-model="data.item.info10"  />
      </el-form-item>

      <el-form-item :label="infoTitle[10]">
        <el-input v-model="data.item.info11"  />
      </el-form-item>      
      
    </el-form>

    <template #footer>
      <el-button size="small" @click="data.visibleInfo = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit(2)">저장</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleStructure"
    title="건축물 기본현황 수정"

    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="100px">
      <el-form-item label="단지명">
        <div style="font-size:16px;font-weight:bold;">{{data.complex[0].value}}</div>
      </el-form-item>

      
      <el-form-item :label="structureTitle[0]">
        <el-input v-model="data.item.structure1"  />
      </el-form-item>

      <el-form-item :label="structureTitle[1]">
        <el-input v-model="data.item.structure2"  />
      </el-form-item>

      <el-form-item :label="structureTitle[2]">
        <el-input v-model="data.item.structure3"  />
      </el-form-item>
      
      <el-form-item :label="structureTitle[3]">
        <el-input v-model="data.item.structure4"  />
      </el-form-item>

      <el-form-item :label="structureTitle[4]">
        <el-input v-model="data.item.structure5"  />
      </el-form-item>

      <el-form-item :label="structureTitle[5]">
        <el-input v-model="data.item.structure6"  />
      </el-form-item>
      
      <el-form-item :label="structureTitle[6]">
        <el-input v-model="data.item.structure7"  />
      </el-form-item>

      <el-form-item :label="structureTitle[7]">
        <el-input v-model="data.item.structure8"  />
      </el-form-item>

      <el-form-item :label="structureTitle[8]">
        <el-input v-model="data.item.structure9"  />
      </el-form-item>
      
      <el-form-item :label="structureTitle[9]">
        <el-input v-model="data.item.structure10"  />
      </el-form-item>

      <el-form-item :label="structureTitle[10]">
        <el-input v-model="data.item.structure11"  />
      </el-form-item>

      <el-form-item :label="structureTitle[11]">
        <el-input v-model="data.item.structure12"  />
      </el-form-item>
      
      <el-form-item :label="structureTitle[12]">
        <el-input v-model="data.item.structure13"  />
      </el-form-item>

      <el-form-item :label="structureTitle[13]">
        <el-input v-model="data.item.structure14"  />
      </el-form-item>            

    </el-form>
    <template #footer>
      <el-button size="small" @click="data.visibleStructure = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit(3)">저장</el-button>
    </template>
  </el-dialog>  
</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Repair, Upload } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import type { UploadInstance } from 'element-plus'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems()
}

function paging(page) {
  data.page = page
  getItems()
}

const item = {
  id: 0,
  apt: 0,
  type: 1,
  status: 1,
  provision: 1,
  parcelrate: 0,
  planyears: 50,
  complex1: '',
  complex2: '',
  completiondate: '',
  completionyear: '',
  completionmonth: '',
  completionday: '',
  info1: '',
  info2: '',
  info3: '',
  info4: '',
  info5: '',
  info6: '',
  info7: '',
  info8: '',
  info9: '',
  info10: '',
  info11: '',
  structure1: '',
  structure2: '',
  structure3: '',
  structure4: '',
  structure5: '',
  structure6: '',
  structure7: '',
  structure8: '',
  structure9: '',
  structure10: '',
  structure11: '',
  structure12: '',
  structure13: '',
  structure14: '',
  name: '',
  familycount: '',
  flatcount: '',
  address: '',
  calculatetype: 1,
  date: ''
}

const data = reactive({
  id: 0,
  originalItem: null,
  total: 0,
  page: 1,
  pagesize: 10,
  item: util.clone(item),
  visible: false,
  visibleInfo: false,
  visibleStructure: false,
  visibleReport: false,
  complex: [],
  info: [],
  struct: [],
  apt: null,
  historydel: true,  
  breakdowndel: true,
  historydelassistance: true,  
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`,
  menu: 'basic',
  level: 0,
  aptid: 0,
  assistances: []
})

watch(() => route.params.id, async () => {
  data.id = util.getInt(route.params.id)
  item.apt = data.id

  await initData()
  await getItems()
})

async function initData() {
}

const complexTitle = ['단지명', '세대수', '동수', '소재지', '사업주체(시행사)', '사업계획승인 일자', '사용검사 일자']
const infoTitle = ['지역ㆍ지구', '대지 면적', '건축 면적', '연 면적', '건폐율', '용적율', '용도', '건물구조', '난방방식', '출입구수', '유형']

const structureTitle = ['관리사무소', '입주자집회소', '경로당', '경비실', '어린이놀이터', '주민운동시설', '주차대수', '가스공급시설', 'CCTV 모니터', 'CCTV 녹화기', 'CCTV 카메라', '승강기 제조사', '승강기 대수', '공동저수시설']

const structureTitle1 = ['부대복리시설', '입주자집회소', '경로당', '경비실', '어린이놀이터', '주민운동시설', '주차대수', '가스공급시설', 'CCTV 설치', '녹화기', '카메라', '승강기', '승강기 대수', '공동저수시설']
const structureTitle2 = ['관리사무소', '', '', '', '', '', '', '', '모니터', '', '', '제조사', '', '']

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const spanMethod = ({
  row,
  column,
  rowIndex,
  columnIndex,
}: SpanMethodProps) => {
  if (columnIndex === 0) {
    if (rowIndex == 0) {
      return {rowspan: 4, colspan: 1}
    } else if (rowIndex == 4) {
      return {rowspan: 1, colspan: 2}
    } else if (rowIndex == 5) {
      return {rowspan: 1, colspan: 2}
    } else if (rowIndex == 6) {
      return {rowspan: 1, colspan: 2}
    } else if (rowIndex == 7) {
      return {rowspan: 1, colspan: 2}  
    } else if (rowIndex == 8) {
      return {rowspan: 3, colspan: 1}
    } else if (rowIndex == 11) {
      return {rowspan: 2, colspan: 1}
    } else if (rowIndex == 13) {
      return {rowspan: 1, colspan: 2}
    } else {
      return {rowspan: 1, colspan: 1}  
    }
  } else if  (columnIndex === 1) {
    if (rowIndex == 1) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 2) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 3) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 4) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 5) {
      return {rowspan: 0, colspan: 0}            
    } else if (rowIndex == 6) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 7) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 9) {
      return {rowspan: 0, colspan: 0}
    } else if (rowIndex == 10) {
      return {rowspan: 0, colspan: 0}      
    } else if (rowIndex == 12) {
      return {rowspan: 0, colspan: 0}      
    } else if (rowIndex == 13) {      
      return {rowspan: 0, colspan: 0}                
    }      
  }

  return {rowspan: 1, colspan: 1}
}


function makeData(label, value) {
  let items = []
  for (let i = 0; i < label.length; i++) {
    items.push({
      label: label[i],
      value: value[i]
    })
  }

  return items
}

async function getItems() {  
  let res = await Repair.get(data.id)
  const item = res.item
  data.originalItem = item

  res = await Apt.get(item.apt)
  const apt = res.item
  data.apt = apt

  const complex = [apt.name, apt.familycount, apt.flatcount, apt.address, item.complex1, item.complex2, `${item.completionyear}년 ${item.completionmonth}월 ${item.completionday}일`]
  data.complex = makeData(complexTitle, complex)

  let provision

  if (item.provision == 2) {
    provision = '집합건물의 소유 및 관리에 관한 법률'
  } else {
    provision = '공동주택관리법'
  }

  let status
  if (item.status == 2) {
    status = '마감'
  } else {
    status = '진행'
  }

  let calculatetype
  if (item.calculatetype == 1) {
    calculatetype = '수선율 적용'
  } else {
    calculatetype = '수선율 적용 안함 (전면교체와 동일한 금액으로 표시)'
  }

  /*
  if (data.level == 2) {
    data.report = [{label: '리포트 작성 일자', value: item.reportdate}, {label: '법조문 종류', value: provision}, {label: '분양 비율', value: item.parcelrate}, {label: '총 계획년도', value: item.planyears}, {label: '부분수선 수선율 적용', value: calculatetype}, {label: '상태', value: status}]
  } else {
    data.report = [{label: '리포트 작성 일자', value: item.reportdate}, {label: '법조문 종류', value: provision}, {label: '분양 비율', value: item.parcelrate}, {label: '총 계획년도', value: item.planyears}, {label: '부분수선 수선율 적용', value: calculatetype}]
  }
  */
  if (data.level > 1) {
    data.report = [{label: '리포트 작성 일자', value: item.reportdate}, {label: '법조문 종류', value: provision}, {label: '분양 비율', value: item.parcelrate}, {label: '총 계획년도', value: item.planyears}, {label: '상태', value: status}]
  } else {
    data.report = [{label: '리포트 작성 일자', value: item.reportdate}, {label: '법조문 종류', value: provision}, {label: '분양 비율', value: item.parcelrate}, {label: '총 계획년도', value: item.planyears}]
  }  

  let info = [];
  for (let i = 1; i <= 11; i++) {
    info.push(item[`info${i}`])
  }
  data.info = makeData(infoTitle, info)

  let structure = [];
  for (let i = 1; i <= 14; i++) {
    structure.push(item[`structure${i}`])
  }


  let items = []
  for (let i = 0; i < structureTitle1.length; i++) {
    items.push({
      label: structureTitle1[i], 
      sublabel: structureTitle2[i],
      value: structure[i]
    })
  }
  
  data.structure = items

}

function clickInsert(pos) {  
  data.item = util.clone(data.originalItem)

  if (pos == 1) {
    data.visible = true
    data.item.name = data.apt.name
    data.item.familycount = data.apt.familycount
    data.item.flatcount = data.apt.flatcount
    data.item.address = data.apt.address

    if (data.item.completionmonth == 0) {
      data.item.completionmonth = 1
    }

    if (data.item.completionday == 0) {
      data.item.completionday = 1
    }
    
    data.item.completiondate = util.makeDate(data.item.completionyear, data.item.completionmonth, data.item.completionday)    
  } else if (pos == 2) {
    data.visibleInfo = true
  } else if (pos == 3) {
    data.visibleStructure = true
  } else if (pos == 4) {
    data.visibleReport = true    
  }
}

async function clickSubmit(type) {
  let res = await Repair.get(data.id)
  let item = res.item

  if (type == 1) {
    item.complex1 = data.item.complex1
    item.complex2 = data.item.complex2

    let date = data.item.completiondate
    if (typeof date == 'string') {
      let d = date.split('-')
      
      data.item.completionyear = d[0]
      data.item.completionmonth = d[1]
      data.item.completionday = d[2]      
    } else if (date == null || date == undefined || date == '') {
      util.error('사용검사일자를 입력하세요')
      return
    } else {
      data.item.completionyear = date.getFullYear()
      data.item.completionmonth = date.getMonth()+1
      data.item.completionday = date.getDate()          
    }
        
    item.completionyear = util.getInt(data.item.completionyear)
    item.completionmonth = util.getInt(data.item.completionmonth)
    item.completionday = util.getInt(data.item.completionday)

    if (data.item.name == '') {
      util.error('단지명을 입력하세요')
      return
    }

    res = await Apt.get(data.originalItem.apt)
    res.item.name = data.item.name
    res.item.familycount = data.item.familycount
    res.item.flatcount = data.item.flatcount
    res.item.address = data.item.address

    await Apt.update(res.item)
  } else if (type == 2) {
    item.info1 = data.item.info1
    item.info2 = data.item.info2
    item.info3 = data.item.info3
    item.info4 = data.item.info4
    item.info5 = data.item.info5
    item.info6 = data.item.info6
    item.info7 = data.item.info7
    item.info8 = data.item.info8
    item.info9 = data.item.info9
    item.info10 = data.item.info10
    item.info11 = data.item.info11
  } else if (type == 3) {
    item.structure1 = data.item.structure1
    item.structure2 = data.item.structure2
    item.structure3 = data.item.structure3
    item.structure4 = data.item.structure4
    item.structure5 = data.item.structure5
    item.structure6 = data.item.structure6
    item.structure7 = data.item.structure7
    item.structure8 = data.item.structure8
    item.structure9 = data.item.structure9
    item.structure10 = data.item.structure10
    item.structure11 = data.item.structure11
    item.structure12 = data.item.structure12
    item.structure13 = data.item.structure13
    item.structure14 = data.item.structure14
  } else if (type == 4) {
    let date = data.item.reportdate
    if (typeof date == 'string') {
      item.reportate = date
    } else if (date == null || date == undefined || date == '') {
      util.error('리포트 작성일을 입력하세요')
      return
    } else {
      item.reportdate = date.getFullYear() + '-' + util.pad(date.getMonth()+1, 2) + '-' + util.pad(date.getDate(), 2)      
    }
    item.provision = data.item.provision
    item.parcelrate = util.getFloat(data.item.parcelrate)
    item.planyears = util.getInt(data.item.planyears)

    item.calculatetype = util.getInt(data.item.calculatetype)

    if (data.level > 1) {
      item.status = util.getInt(data.item.status)
    }
  }

  res = await Repair.update(item)

  data.complex[4] = data.item.complex1
  data.complex[5] = data.item.complex2
  data.complex[6] = `${data.item.completionyear}년 ${data.item.completionmonth}월 ${data.item.completionday}일`

  if (res.code === 'ok') {
    util.info('등록되었습니다')
    getItems()
    data.visible = false
    data.visibleInfo = false
    data.visibleStructure = false
    data.visibleReport = false
  } else {
    util.error('오류가 발생했습니다')
  }
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
  item.apt = id
  data.aptid = apt

  if (store.getters['getUser'] != null) {
    data.level = store.getters['getUser'].level
  }
  
  await initData()
  await getItems()
})

function cellClassName({columnIndex}) {
  if (columnIndex == 0) {
    return 'title'
  } else {
    return 'value'    
  }
}

function cellClassNameStructure({columnIndex}) {
  if (columnIndex < 2) {
    return 'title'
  } else {
    return 'value'    
  }
}

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {  
}

async function handelSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  util.loading(true)

  let historydel = 0
  let breakdowndel = 0
  
  if (data.historydel == true) {
    historydel = 1
  }

  if (data.breakdowndel == true) {
    breakdowndel = 1
  }
  
  let params = {
    id: item.apt,
    historydel: historydel,
    breakdowndel: breakdowndel,
    filename: response.filename
  }

  await Upload.excel(item.apt, params)
  /*
     const res = await request({
     method: 'GET',
     url: 'api/upload/excel/' + item.apt,
     params: params
     })
   */

  getItems()

  util.loading(false)
  util.info('등록되었습니다')
}

const submitUpload = () => {
  //util.confirm('엑셀파일을 등록하면 기존 데이터는 모두 삭제됩니다. 실행하시겠습니까', async function() {
  upload.value.clearFiles()
  upload.value!.submit()
  //})
}

const uploadAssistance = ref<UploadInstance>()

const handleExceedAssistance: UploadProps['onExceed'] = (files, uploadFiles) => {  
}

async function handelSuccessAssistance(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
}

async function uploadAssistanceProcess() {
  util.loading(true)
  
  let historydel = 0
  
  if (data.historydelassistance == true) {
    historydel = 1
  }

  let filenames = []
  data.assistances.forEach(v => filenames.push(v.response.filename))
  
  let params = {
    id: item.apt,
    historydel: historydel,
    filenames: filenames
  }

  await Upload.assistance(params)
  
  getItems()

  util.loading(false)
  util.info('등록되었습니다')
  
  uploadAssistance.value.clearFiles()
}
const submitUploadAssistance = () => {
  if (data.assistances == null) {
    util.error('업로드 된 파일이 없습니다')
    return
    
  }

  if (data.assistances.length == 0) {
    util.error('업로드 된 파일이 없습니다')
    return
  }
  
  if (data.historydelassistance == true) {
    util.confirm('보조부원장 파일을 등록하면 기존 충당금 적립 데이터는 모두 삭제됩니다. 실행하시겠습니까', function() {
      uploadAssistanceProcess()
    })
  } else {
    uploadAssistanceProcess()
  }
}

async function submitChange() {
  util.loading(true)
  let res = await Repair.change(data.id)
  console.log(res.item)
  data.menu = 'basic'
  util.info('변환되었습니다')
  util.loading(false)
}

function clickChange() {
    util.confirm('공사종별 변환을 실행하면 기존 데이터가 변경됩니다. 실행은 5분 정도 소요됩니다. 실행하시겠습니까', function() {
      submitChange()
    })
}

</script>
<style>
.title {
  background-color: #fafafa;
}

.value {
  background-color: #FFF;  
}

.el-upload-list__item-file-name {
  font-size: 12px;
}
</style>  
