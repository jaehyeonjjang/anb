<template>
  <Title title="참여 기술자" />

  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="primary" @click="clickInsert">추가</el-button>
  </div>


  <el-table :data="data.items" border :height="height(170)">
    <el-table-column prop="type" label="참여구분" align="center" width="120">
      <template #default="scope">
        {{Detailtechnician.getType(scope.row.type)}}
      </template>
    </el-table-column>
    <el-table-column prop="extra.technician.name" label="성명"  />
    <el-table-column prop="part" label="참여분야" />
    <el-table-column prop="part" label="참여 기간" align="center">
      <template #default="scope">
        {{util.viewDate(scope.row.signupstartdate)}} ~ {{util.viewDate(scope.row.signupenddate)}}
      </template>
    </el-table-column>
    <el-table-column label="기술등급" align="center" width="120">
      <template #default="scope">
        {{Technician.getGrade(scope.row.extra.technician.grade)}}
      </template>
    </el-table-column>
    <el-table-column prop="remark" label="비고" />
    <el-table-column label="" align="center" width="180">
      <template #default="scope">
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickUp(scope.row, scope.$index)"><el-icon><ArrowUp /></el-icon></el-button>
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickDown(scope.row, scope.$index)"><el-icon><ArrowDown /></el-icon></el-button>
        <el-button size="small" type="success" style="margin-right:-7px;" @click="clickUpdate(scope.$index, scope.row)">수정</el-button>
        <el-button size="small" type="danger" style="margin-right:0px;" @click="clickDelete(scope.row, scope.$index)">삭제</el-button>
      </template>
    </el-table-column>
  </el-table>


  <el-dialog v-model="data.visible" width="600px">
    <y-table>
      <y-tr>
        <y-th>참여구분</y-th>
        <y-td>
          <el-radio-group v-model.number="data.item.type">
            <el-radio-button size="small" label="1">책임기술자</el-radio-button>
            <el-radio-button size="small" label="2">참여기술자</el-radio-button>
          </el-radio-group>
        </y-td>
      </y-tr>

      <y-tr>
        <y-th>기술자</y-th>
        <y-td>

          <el-select v-model.number="data.item.technician" style="width:100%;">
            <el-option v-for="item in data.alltechnicians" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </y-td>
      </y-tr>
      <y-tr>
        <y-th>참여분야</y-th>
        <y-td>

          <el-input v-model="data.item.part" placeholder="" />
        </y-td>
      </y-tr>

      <y-tr>
        <y-th>참여 기간</y-th>


        <y-td><el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.signupstartdate" /> ~ <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.signupenddate" /> </y-td>

      </y-tr>

      <y-tr>
        <y-th>비고</y-th>
        <y-td>

          <el-input v-model="data.item.remark" placeholder="" />
        </y-td>
      </y-tr>

    </y-table>


    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
    </template>
  </el-dialog>

</template>

<script setup lang="ts">

import { reactive, onMounted, ref, watch } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Detailtechnician, Technician } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const item = {
  id: 0,
  type: 0,
  part: '',
  signupstartdate: null,
  signupenddate: null,
  remark: '',
  order: 0,
  technician: 0,
  detail: 0,
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  items: [],
  visible: false,
  alltechnicians: []
})

async function initData() {
  let res = await Technician.find({orderby: 'te_name'})
  if (res.items == undefined) {
    res.items = []
  }

  let items = [{id: 0, name: ' '}]

  res.items.forEach(item => {
    items.push({id: item.id, name: `${item.name} - ${Technician.getGrade(item.grade)}`, original: item.name, grade: item.grade});
  })
  data.alltechnicians = items
}

async function getItems() {
  let res = await Detailtechnician.find({detail: data.id, orderby: 'dt_order,dt_id'})

  if (res.items == undefined) {
    res.items = []
  }
  data.items = res.items
}

async function clickSubmit(type) {
  let item = util.clone(data.item)

  item.type = util.getInt(item.type)
  item.technician = util.getInt(item.technician)
  item.signupstartdate = util.convertDate(item.signupstartdate)
  item.signupenddate = util.convertDate(item.signupenddate)


  if (item.type == 0) {
    util.error('참여구분을 선택하세요')
    return
  }

  if (item.technician == 0) {
    util.error('기술자를 선택하세요')
    return
  }

  if (item.id > 0) {
    await Detailtechnician.update(item)
    util.info('수정되었습니다')
  } else {
    let max = 0
    for (let i = 0; i < data.items.length; i++) {
      if (data.items[i].order > max) {
        max = data.items[i].order
      }
    }

    item.order = max + 1
    await Detailtechnician.insert(item)
    util.info('등록되었습니다')
  }

  data.visible = false

  await getItems()
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

function clickInsert() {
  data.item = util.clone(item)
  data.item.detail = data.id
  data.visible = true
}

function clickUpdate(pos, item) {
  data.item = util.clone(item)
  data.visible = true
}

function clickDelete(item, index) {
  util.confirm('삭제하시겠습니까', async function() {
    let res = await Detailtechnician.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      await getItems(true)
    }
  })
}

async function clickUp(row, index) {
  if (index == 0) {
    return
  }

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index - 1].order
  items[index - 1].order = temp

  await Detailtechnician.update(items[index])
  await Detailtechnician.update(items[index - 1])

  await getItems()
}

async function clickDown(row, index) {
  if (index >= data.items.length - 1) {
    return
  }

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index + 1].order
  items[index + 1].order = temp

  await Detailtechnician.update(items[index])
  await Detailtechnician.update(items[index + 1])

  await getItems()

}

</script>
