<template>
  <Title title="시설물" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄처리</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :height="height(170)" @row-click="clickUpdate"  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />    
    <el-table-column prop="dong" label="동별" align="center" />
    <el-table-column prop="type" label="종별" align="center" />
    <el-table-column prop="parkingcount" label="주차장" align="center" width="80" />
    <el-table-column prop="undergroundcount" label="지하" align="center" width="80" />
    <el-table-column prop="groundcount" label="지상" align="center" width="80" />
    <el-table-column prop="topcount" label="옥탑" align="center" width="80" />
    <el-table-column prop="roofcount" label="지붕" align="center" width="80" />    
    <el-table-column prop="familycount" label="세대수" align="center" />
    <el-table-column prop="area" label="연면적" align="center" />        
    <el-table-column prop="remark" label="용도" align="center" />
    <el-table-column label="구분" align="center" width="120">
      <template #default="scope">
        {{data.privates[scope.row.private]}}
      </template>
    </el-table-column>
    <el-table-column label="" align="center" width="150">
      <template #default="scope">
        <el-button size="small" style="margin-right:-7px;" @click="clickFloor(scope.row, scope.$index)">층 추가</el-button>
        
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickUp(scope.row, scope.$index)"><el-icon><ArrowUp /></el-icon></el-button>
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickDown(scope.row, scope.$index)"><el-icon><ArrowDown /></el-icon></el-button>        
      </template>
    </el-table-column>    
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1000px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="동별" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].dong" />
          </template>
        </el-table-column>
        <el-table-column label="종별" align="center" width="50">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].type" />
          </template>
        </el-table-column>
        <el-table-column label="주차장" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].parkingcount" />
          </template>
        </el-table-column>
        <el-table-column label="지하" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].undergroundcount" />
          </template>
        </el-table-column>
        <el-table-column label="지상" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].groundcount" />
          </template>
        </el-table-column>
        <el-table-column label="옥탑" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].topcount" />
          </template>
        </el-table-column>
        <el-table-column label="지붕" align="center" width="60">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].roofcount" />
          </template>
        </el-table-column>
        <el-table-column label="세대수" align="center" width="80">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].familycount" />
          </template>
        </el-table-column>
        <el-table-column label="연면적" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].area" />
          </template>
        </el-table-column>
        <el-table-column label="용도" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].remark" />
          </template>
        </el-table-column>
        <el-table-column label="구분" align="center" width="130">
          <template #default="scope">
            <el-select v-model.number="data.batchs[scope.$index].private" style="width:100%;">
              <el-option v-for="(item, index) in data.privates" :key="index" :label="item" :value="index" />
            </el-select>
          </template>
        </el-table-column>

      </el-table>


    </el-form>

      <template #footer>
        <el-button size="small" type="danger" v-if="data.mode != 'batch' && (data.batchs.length > 0 && data.batchs[0].id > 0)" style="float:left;" @click="clickDelete">삭제</el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(1)"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" v-if="data.mode == 'batch'" style="float:left;" @click="clickAdd(10)"><el-icon><Plus /></el-icon> &nbsp;10</el-button>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleFloor"
    width="600px"
  >

    <div style="overflow:auto;padding:10px 10px;" :style="{height: height(500)}">
      <div v-for="(item, index) in data.floors" :key="item.id" class="block">
        <div style="flex:1;height:19px;padding-top:5px;" @click="clickPreview(item)">{{item.name}}</div>
        
        <el-button size="small" @click="clickAddFloor(item, index)"><el-icon><Plus /></el-icon></el-button>
        <el-button size="small" @click="clickRemoveFloor(item, index)" :disabled="item.type == 1"><el-icon><Minus /></el-icon></el-button>
      </div>
    </div>
    
    <template #footer>        
      <el-button size="small" @click="clickCancelFloor">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitFloor">등록</el-button>
    </template>
  </el-dialog>

  <el-dialog
    v-model="data.visibleAdd"
    width="400px"
  >

    <div style="text-align:left;margin-bottom:10px;">추가할 층 명칭을 입력하세요</div>
    <el-radio-group v-model="data.floortype">
      <el-radio :label="1" size="small">주하장</el-radio>
      <el-radio :label="2" size="small">지하</el-radio>
      <el-radio :label="3" size="small">지상</el-radio>
      <el-radio :label="4" size="small">옥탑</el-radio>
      <el-radio :label="5" size="small">지붕</el-radio>
    </el-radio-group>        
    <div style="height:10px;"></div>
    <el-input v-model="data.name" />
    
    <template #footer>        
      <el-button size="small" @click="clickCancelAdd">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitAdd">추가</el-button>
    </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Aptdong, Aptdongetc } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Aptdong

const item = {
  id: 0,
  dong: '',
  type: '',
  groundcount: 0,
  undergroundcount: 0,
  parkingcount: 0,
  topcount: 0,
  roofcount: 0,
  familycount: 0,
  area: '',
  remark: '',
  order: 0,
  private: '',
  date: ''
}

const data = reactive({
  apt: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  visibleFloor: false,
  visibleAdd: false,
  privates: [' ', '개별 동', '기타 공용시설', '공통층 (대형건축물)', '아파트 주차장'],
  floors: [],
  name: '',
  floortype: 3,
  index: -1
})

async function initData() {  
}

async function getItems() {
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'au_order,au_id'    
  })

  if (res.items == null) {
    res.items = []
  }
  
  
  for (let i = 0; i < res.items.length; i++) {
    res.items[i].index = i + 1
  }

  data.total = res.total
  data.items = res.items
}

function clickInsert() {  
  data.item = util.clone(item)

  let items = [data.item]
  
  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0 || index.no == 12) {
    return
  }

  let items = [util.clone(item)]

  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await model.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      await getItems()
    }
  })
}

const handleClose = (done: () => void) => {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      done()
    })
  } else {
    done()
  }
}

onMounted(async () => {
  data.apt = parseInt(route.params.apt)  
  
  util.loading(true)
  
  await initData()
  await getItems()

  data.visible = false
  util.loading(false)
})

function clickBatch() {
  let items = util.clone(data.items)

  if (items == null) {
    items = []
  }

  if (items.length == 0) {
    for (let i = 0; i < 5; i++) {
      items.push(util.clone(data.item))
    }
  }

  data.mode = 'batch'
  data.batchs = items
  data.visible = true  
}

function clickCancel() {
  if (data.mode == 'batch') {
    util.confirm('팝업창을 닫으시겠습니까', function() {
      data.visible = false
    })
  } else {
    data.visible = false
  }
}

const listRef = ref<InstanceType<typeof ElTable>>()
const listSelection = ref([])
const toggleListSelection = (rows) => {
  if (rows) {
    rows.forEach((row) => {
      listRef.value!.toggleRowSelection(row, undefined)
    })
  } else {
    listRef.value!.clearSelection()
  }
}
const changeList = (val) => {
  listSelection.value = val
}

function clickDeleteMulti() {
  util.confirm('삭제하시겠습니까', async function() {
    util.loading(true)
    
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      let item = {
        id: value.id,
        apt: data.apt
      }

      await model.remove(item)
    }

    util.info('삭제되었습니다')
    await getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  util.loading(true)

  for (let i = 0; i < data.items.length; i++) {
    let item = data.items[i]
    
    for (let j = 0; j < data.batchs.length; j++) {
      let target = data.batchs[j]
      if (item.id == target.id) {
        if (item.parkingcount != target.parkingcount ||
            item.undergroundcount != target.undergroundcount ||
            item.groundcount != target.groundcount ||
            item.topcount != target.topcount ||
            item.roofcount != target.roofcount) {          
          util.confirm('시설물 정보를 수정하면 해당 시설물에 추가된 층 정보가 모두 삭제됩니다. 수정하시겠습니까', function() {
            clickSubmitProcess()
          }, function() {
            util.loading(false)
          })
          return
        }
      }
    }
  }

  clickSubmitProcess()
}    

async function clickSubmitProcess() {
  if (data.mode == 'batch') {
    for (let i = 0; i < data.items.length; i++) {
      let item = data.items[i]
      let flag = false;
      for (let j = 0; j < data.batchs.length; j++) {
        if (data.items[i].id == data.batchs[j].id) {
          flag = true
          break
        }
      }

      if (flag == false) {      
        await model.remove(item)
      }
    }
  } else {
    if (data.batchs[0].id == 0) {      
      let max = 0
      for (let i = 0; i < data.items.length; i++) {
        let item = data.items[i]

        if (item.order > max) {
          max = item.order
        }
      }

      max++                
      data.batchs[0].order = max
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    if (item.dong == '') {
      continue
    }
    
    item.apt = data.apt
    item.dong = String(item.dong)
    
    item.type = String(item.type)
    item.groundcount = util.getInt(item.groundcount)
    item.undergroundcount = util.getInt(item.undergroundcount)
    item.parkingcount = util.getInt(item.parkingcount)
    item.topcount = util.getInt(item.topcount)
    item.roofcount = util.getInt(item.roofcount)
    item.familycount = util.getInt(item.familycount)
    item.private = util.getInt(item.private)
    item.order = util.getInt(item.order)
    if (item.private == 0) {
      item.private = 1
    }
    
    item.area = String(item.area)
    item.remark = String(item.remark)
    
    if (data.mode == 'batch') {
      item.order = i + 1
    }
    
    if (item.id > 0) {
      await model.update(item)
    } else { 
      await model.insert(item)
    }
  }

  util.info('등록되었습니다')
  
  await getItems()
  data.visible = false  
  util.loading(false)  
}

function clickRegistDelete(index) {
  data.batchs.splice(index, 1)
}

function clickAdd(count) {
  let items = []
  for (let i = 0; i < count; i++) {
    items.push(util.clone(item))
  }

  data.batchs = data.batchs.concat(items)
}

async function clickUp(row, index) {
  if (index == 0) {
    return
  }

  util.loading(true)

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index - 1].order
  items[index - 1].order = temp

  await model.update(items[index])
  await model.update(items[index - 1])

  await getItems()

  util.loading(false)
}

async function clickDown(row, index) {
  if (index >= data.items.length - 1) {
    return
  }

  util.loading(true)

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index + 1].order
  items[index + 1].order = temp

  await model.update(items[index])
  await model.update(items[index + 1])

  await getItems()

  util.loading(false)
}

async function clickFloor(item2, index) {
  let item = util.clone(item2)
  if (item.groundcount > 0 && item.topcount == 0) {
    item.roofcount = 1
  } else {
    item.roofcount = 0    
  }

  let items = []
  
  for (let i = 0; i < item.parkingcount; i++) {
    items.push({
      floortype: 1,
      name: `주차장 지하${item.undergroundcount+item.parkingcount-i}층`,
      pos: (item.undergroundcount+item.parkingcount-i) * -10,
      type: 1,
      parent: 0
    })
  }

  for (let i = 0; i < item.undergroundcount; i++) {
    items.push({
      floortype: 2,
      name: `지하${item.undergroundcount-i}층`,
      pos: (item.undergroundcount-i) * -10,
      type: 1,
      parent: 0
    })
  }

  for (let i = 0; i < item.groundcount; i++) {
    items.push({
      floortype: 3,
      name: `${i+1}층`,
      pos : (i + 1) * 10,
      type: 1,
      parent: 0
    })
  }

  for (let i = 0; i < item.roofcount; i++) {
    items.push({
      floortype: 4,
      name: `지붕층`,
      pos : (item.groundcount + 1) * 10,
      type: 1,
      parent: 0
    })
  }

  for (let i = 0; i < item.topcount; i++) {
    if (i == 0) {
      items.push({
        floortype: 5,
        name: `옥탑/지붕층`,
        pos : (item.groundcount + i + 1) * 10,
        type: 1,
        parent: 0
      })
    } else {
      items.push({
        floortype: 5,
        name: `옥탑${i+1}층`,
        pos : (item.groundcount + i + 1) * 10,
        type: 1,
        parent: 0
      })
    }
  }

  data.item = item
  items = items.reverse()

  let res = await Aptdongetc.find({apt:data.apt, aptdong:data.item.id, orderby: 'ae_order,ae_id'})  
  if (res.items == null) {
    res.items = []
  }
  
  for (let i = 0; i < res.items.length; i++) {
    res.items[i].pos = res.items[i].order 
    for (let k = 0; k < items.length; k++) {
      if (res.items[i].parent == items[k].pos) {
        items.splice(k + 1, 0, res.items[i])
        break
      }
    }
  }

  data.floors = items

  data.visibleFloor = true  
}

function clickCancelFloor() {
  data.visibleFloor = false
}

async function clickSubmitFloor() {
  util.loading(true)
  let items = util.clone(data.floors)

  let inserts = []
  
  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    if (item.type == 1) {
      continue
    }

    let n = {
      name: item.name,
      parent: item.parent,
      order: item.pos,
      aptdong: data.item.id,
      floortype: item.floortype,
      apt: data.apt
    }

    inserts.push(n)
  }
  
  await Aptdong.blueprint(data.apt, data.item.id, inserts)
  
  data.visibleFloor = false
  util.info('등록되었습니다')
  util.loading(false)
}

function clickAddFloor(item, index) {
  data.name = ''
  data.floortype = 3
  data.index = index
  data.visibleAdd = true
}

function clickRemoveFloor(item, index) {
  let items = util.clone(data.floors)

  for (let i = index + 1; i < items.length; i++) {
    let item = items[i]

    if (item.type == 1) {
      break
    }

    items[i].pos--
  }
  
  items.splice(index, 1)
  data.floors = items
}

function clickCancelAdd() {
  dta.visibleAdd = false
}

function clickSubmitAdd() {
  let items = util.clone(data.floors)

  let parent = 0
  for (let i = data.index; i >= 0; i--) {
    let item = items[i]

    if (item.type == 1) {
      parent = item.pos
      break
    }

  }

  for (let i = data.index + 1; i < items.length; i++) {
    let item = items[i]

    if (item.type == 1) {
      break
    }

    items[i].pos++
  }


  let pos = items[data.index].pos - 1
  
  let item = {
    floortype: data.floortype,
    name: data.name,
    pos : pos,
    type: 2,
    parent: parent
  }
  items.splice(data.index + 1, 0, item)
  data.floors = items
  
  data.visibleAdd = false
}

</script>
<style>
.block {
  border: 1px solid #aaa;
  border-radius: 3px;
  padding: 5px 10px;
  margin-bottom: 5px;
  text-align: left;
  font-weight: bold;  
  display:flex;
  flex-direction:row;
  justify-content: space-between;
}

.input {
  display:block;
  float:left;
  width:400px;
  margin-right: 10px;
  padding: 2px 5px;
}

.btn {
  cursor: hand;
  cursor: pointer;
  display:block;
  float:right;
  font-size:20px;
  margin: 2px 5px 0px 5px;
}

.clear {
  clear:both;
}

</style>
