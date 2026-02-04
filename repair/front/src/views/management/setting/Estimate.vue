<template>
  <Title title="견적 기본값 관리" />

  <y-table>
    <y-tr>
      <y-th>구분</y-th>
      <y-th>{{ getCurrentYearLabel() }}</y-th>
      <y-th>{{ getPreviousYearLabel() }}</y-th>
    </y-tr>
    <y-tr>
      <y-th>기술사</y-th>
      <y-td>{{util.money(data.currentYear?.person1 || 0)}}</y-td>
      <y-td>{{util.money(data.previousYear?.person1 || 0)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>특급기술자</y-th>
      <y-td>{{util.money(data.currentYear?.person2 || 0)}}</y-td>
      <y-td>{{util.money(data.previousYear?.person2 || 0)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>고급기술자</y-th>
      <y-td>{{util.money(data.currentYear?.person3 || 0)}}</y-td>
      <y-td>{{util.money(data.previousYear?.person3 || 0)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>중급기술자</y-th>
      <y-td>{{util.money(data.currentYear?.person4 || 0)}}</y-td>
      <y-td>{{util.money(data.previousYear?.person4 || 0)}}</y-td>
    </y-tr>
    <y-tr>
      <y-th>초급기술자</y-th>
      <y-td>{{util.money(data.currentYear?.person5 || 0)}}</y-td>
      <y-td>{{util.money(data.previousYear?.person5 || 0)}}</y-td>
    </y-tr>
  </y-table>
  
  <p v-if="!data.previousYear" style="color:#999;font-size:12px;margin-top:10px;">※ 전년도 단가 데이터가 없습니다</p>
  
  <el-dialog
    v-model="data.visible"
    title="사용자 등록/수정"
    width="600px"
    :before-close="handleClose"
  >
    <el-form :model="data.item" label-width="80px">
      <el-form-item label="ID" v-show="data.item.id != 0">
        {{ data.item.id }}
      </el-form-item>
      <el-form-item label="아이디">
        <el-input v-model="data.item.loginid" />
      </el-form-item>
      <el-form-item label="비밀번호">
        <el-input v-model="data.item.passwd" />
      </el-form-item>
      <el-form-item label="이름">
        <el-input v-model="data.item.name" />
      </el-form-item>
      <el-form-item label="이메일">
        <el-input v-model="data.item.email" />
      </el-form-item>
      
      <el-form-item label="권한">
        <el-select v-model.number="data.item.level" class="m-2" placeholder="권한">
          <el-option
            v-for="item in data.levels"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>

      </el-form-item>      
    </el-form>

    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>  
</template>

<script setup lang="ts">

import { reactive, onMounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Standardwage } from "~/models"

const { width, height } = size()

const model = Standardwage

const search = reactive({
  text: ''
})

function clickSearch() {
  getItems(true)
}

const item = {    
}

const data = reactive({
  items: [],
  total: 0,  
  item: util.clone(item),
  currentYear: null,
  previousYear: null,
  visible: false  
})

async function initData() {  
}

async function getItems(reset) {
  let res = await Standardwage.find({orderby: 'sw_date DESC'})
  
  if (res.items && res.items.length > 0) {
    data.currentYear = res.items[0]
    // 두번째가 전년도 (있으면)
    if (res.items.length > 1) {
      data.previousYear = res.items[1]
    } else {
      // 전년도 단가 (2024년도)
      data.previousYear = {
        ...res.items[0],
        id: 999,
        person1: 452718,  // 기술사
        person2: 358273,  // 특급기술자
        person3: 300980,  // 고급기술자
        person4: 284046,  // 중급기술자
        person5: 223644,  // 초급기술자
        person6: 452718,
        person7: 358273,
        person8: 300980,
        person9: 284046,
        person10: 223644,
        date: '2024-01-01 00:00:00'
      }
    }
  } else {
    // 없으면 ID=1 조회 (기존 방식)
    res = await model.get(1)
    data.currentYear = res.item
    data.previousYear = null
  }
  
  data.item = data.currentYear
}

function getCurrentYearLabel() {
  if (!data.currentYear?.date) return '현년도 단가'
  const year = data.currentYear.date.substring(0, 4)
  return `${year}년도 단가`
}

function getPreviousYearLabel() {
  return '전년도 단가'
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true
}

function clickUpdate(pos, item) {
  data.item = util.clone(item)
  data.visible = true
}

async function clickSubmit() {
  const item = data.item
  if (item.loginid === '') {
    util.error('아이디를 입력하세요')
    return    
  }

  if (item.passwd === '') {
    util.error('비밀번호를 입력하세요')
    return
  }

  if (item.passwd === '') {
    util.error('이름을 입력하세요')
    return
  }

  if (item.level === 0) {
    util.error('권한을 선택하세요')
    return
  }
  
  let res;

  let count = await User.countByLoginid(item.loginid)
  if (count > 0) {
    util.error('이미 등록된 아이디입니다. 다른 아이디를 입력하세요')
    return
  }

  if (item.id === 0) {
    res = await User.insert(item)
  } else {
    res = await User.update(item)
  }

  if (res.code === 'ok') {
    util.info('등록되었습니다')
    getItems(true)
    data.visible = false
  } else {
    util.error('오류가 발생했습니다')
  }
}

const handleClose = (done: () => void) => {
  util.confirm('팝업창을 닫으시겠습니까', function() {
    done()
  })  
}

onMounted(async () => {
  util.loading(true)
  
  await initData()
  await getItems()

  util.loading(false)
})

</script>
