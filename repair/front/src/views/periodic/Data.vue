<template>
  <Title title="결함현황표" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:-5px;">등록</el-button>
      <el-button size="small" type="warning" @click="clickBatch">일괄처리</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :height="height(170)" @row-click="clickUpdate"  ref="listRef" @selection-change="changeList" v-show="data.preview == false">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column label="위치" align="center">
      <template #default="scope">
        {{getBlueprint(scope.row)}}
      </template>
    </el-table-column>
    <el-table-column prop="group" label="순번" align="center" width="50" />
    <el-table-column prop="part" label="부위" align="center" />
    <el-table-column prop="member" label="부재" align="center" />
    <el-table-column prop="shape" label="유형 및 형상" align="center" />    
    <el-table-column prop="width" label="폭" align="center" width="80" />
    <el-table-column prop="length" label="길이" align="center" width="80" />        
    <el-table-column prop="count" label="개수" align="center" width="80" />
    <el-table-column label="진행사항" align="center" width="80">
      <template #default="scope">
        {{scope.row.progress == 2 ? 'X' : 'O'}}
      </template>
    </el-table-column>
    <el-table-column prop="remark" label="비고" align="center" />
    <el-table-column label="이미지" align="center">
      <template #default="scope">
        <div v-if="scope.row.filename != ''">
          <el-image v-for="(item, index) in scope.row.filename.split(',')" style="width: 20px; height: 20px; top:4px;left:0px;position:relative;margin-right:5px;"
                    :src="util.getImagePath(item)"                    
                    fit="cover"
                    @click="clickPreviews(scope.row.filename, index)"
          />
        </div>
      </template>
    </el-table-column>
    <el-table-column label="결함도" align="center" width="50">
      <template #default="scope">
        <el-image v-if="scope.row.extra.resultimage != ''"          
                  style="width: 20px; height: 20px; top:4px;left:0px;position:relative;z-index:999999999999999 !important;"
                  :src="util.getImagePath(scope.row.extra.resultimage)"                  
                  fit="cover"
                  @click="clickPreview(scope.row.extra.resultimage)"
        />
        
      </template>
    </el-table-column>
    <!--
    <el-table-column label="" align="center" width="105">
      <template #default="scope">

        
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickUp(scope.row, scope.$index)"><el-icon><ArrowUp /></el-icon></el-button>
        <el-button size="small" style="width:30px;margin-right:-7px;" @click="clickDown(scope.row, scope.$index)"><el-icon><ArrowDown /></el-icon></el-button>        
      </template>
    </el-table-column>
    -->
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="1050px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="위치" align="center" width="180">
          <template #default="scope">
            <el-tree-select style="width:170px;" v-model="data.batchs[scope.$index].blueprint" :data="data.blueprints" :default-expand-all="true" :render-after-expand="false" placeholder="위치" />
          </template>
        </el-table-column>
        <el-table-column label="순번" align="center" width="50">
          <template #default="scope">
            <el-input v-model.number="data.batchs[scope.$index].group" />
          </template>
        </el-table-column>
        <el-table-column label="부위" align="center" width="90">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].part" />
          </template>
        </el-table-column>
        <el-table-column label="부재" align="center" width="100">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].member" />
          </template>
        </el-table-column>
        <el-table-column label="유형 및 형상" align="center" width="170">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].shape" />
          </template>
        </el-table-column>
        <el-table-column label="폭" align="center" width="60">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].width" />
          </template>
        </el-table-column>
        <el-table-column label="길이" align="center" width="60">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index]['length']" />
          </template>
        </el-table-column>
        <el-table-column label="개소" align="center" width="50">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].count" />
          </template>
        </el-table-column>
        <el-table-column label="진행" align="center" width="70">
          <template #default="scope">
            <el-select v-model.number="data.batchs[scope.$index].progress" style="width:100%;">
              <el-option v-for="item in data.progresss" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column label="비고" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].remark" />
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

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Periodicdata, Blueprint } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodicdata

const item = {
  id: 0,
  group: 0,
  part: '',
  member: '',
  shape: '',
  width: '',
  length: '',
  count: 0,
  progress: '',
  remark: '',
  order: 0,
  content: '',
  blueprint: 0,
  periodic: 0,
  date: ''
}

const data = reactive({
  apt: 0,
  id: 0,
  mode: 'normal',
  items: [],
  total: 0,
  page: 1,
  pagesize: 0,
  item: util.clone(item),
  visible: false,
  allblueprints: [],
  blueprints: [],
  progresss: [{id: 1, name: 'O'}, {id: 2, name: 'X'}],
  preview: false
})

async function initData() {
  let res = await Blueprint.find({apt: data.apt, orderby: 'bp_parentorder,bp_order desc,bp_id'})

  data.allblueprints = res.items
  let blueprints = [{label:'위치', value: 0}]
  let items = res.items
  if (items == null) {
    items = []
  }
  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    if (item.level != 1) {
      continue
    }

    let children = []
    for (let i = 0; i < items.length; i++) {
      let item2 = items[i]

      if (item2.parent != item.id) {
        continue
      }

      children.push({label: `${item.name} ${item2.name}`, value: item2.id})
    }
    
    blueprints.push({label: item.name, value: item.id, children: children})
  }

  data.blueprints = blueprints
}

async function getItems() {
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    periodic: data.id,
    orderby: 'bp_parentorder,bp_order desc,bp_id,pd_order,pd_id'    
  })

  let items = []

  if (res.items == null) {
    res.items = []
  }

  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    if (item.group == 0) {
      continue
    }

    if (item.type >= 200) {
      continue
    }
    
    item.index = i + 1
    items.push(item)
  }  

  data.total = res.total  
  data.items = items

  console.log(data.items)
}

function clickInsert() {  
  data.item = util.clone(item)

  let items = [data.item]
  
  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index == undefined) {
    return
  }

  if (index.no == 0 || index.no > 10) {
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
  data.id = parseInt(route.params.id)
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
        id: value.id
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

  let blueprint = 0
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    
    
    if (item.blueprint == 0) {
      item.blueprint = blueprint
    } else {
      blueprint = util.getInt(item.blueprint)
    }
    
    item.periodic = data.id

    item.progress = util.getInt(item.progress)
    item.group = util.getInt(item.group)
    item.count = util.getInt(item.count)

    if (item.blueprint == 0 && item.group == 0) {
      continue
    }
    
    if (data.mode == 'batch' && item.id == 0) {
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

  let items = util.clone(data.items)

  let temp = items[index].order
  items[index].order = items[index - 1].order
  items[index - 1].order = temp

  await model.update(items[index])
  await model.update(items[index - 1])

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

  await model.update(items[index])
  await model.update(items[index + 1])

  await getItems()
}

function getBlueprint(item) {
  for (let i = 0; i < data.allblueprints.length; i++) {
    if (item.blueprint == data.allblueprints[i].id) {
      for (let j = 0; j < data.allblueprints.length; j++) {
        if (data.allblueprints[j].id == data.allblueprints[i].parent) {
          return data.allblueprints[j].name + ' ' + data.allblueprints[i].name 
        }
      }

      return data.allblueprints[i].name
    }
  }
}

function getMainImage(item) {  
  let temp = item.filename.split(',')
  
  return util.getImagePath(temp[0])
}

function getImageList(item) {
  if (item.filename == '') {
    return []
  }
  
  let items = []
  
  let temp = item.filename.split(',')

  console.log(temp)

  for(var i = 0; i < temp.length; i++) {
    console.log(temp[i])
    items.push(util.getImagePath(temp[i]))
  }

  console.log('______________')
  console.log(items)
  return items
}

function clickPreview(url) {
  const img = util.getImagePath(url)
  v3ImgPreviewFn(img)  
}

function clickPreviews(str, index) {
  const imgs = str.split(',').map(item => util.getImagePath(item)) 
  v3ImgPreviewFn({images:imgs, index: index})  
}
</script>
