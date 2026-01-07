<template>
  <Title :title="data.title" subtitle="장기수선계획" />


  <el-tabs v-model="data.menu">
    <el-tab-pane label="작업 목록" name="list">

      <el-table :data="data.items" border :width="data.width" :height="height(566)" @row-click="clickRepair" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
        <el-table-column prop="index" label="NO" align="center" width="60" />
        <el-table-column prop="type" label="구분" align="center" width="100">
          <template #default="scope">
            <el-tag :type="Repair.getTypeType(scope.row.type)">{{Repair.getType(scope.row.type)}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="reportdate" label="리포트 작성일" align="center" />
        <el-table-column label="조정" width="60" align="center">
          <template #default="scope">
            <span v-if="scope.row.periodtype == 2">수시</span>
            <span v-else>정기</span>
          </template>
        </el-table-column>
        <el-table-column label="설명">
          <template #default="scope">
            {{scope.row.remark}}
          </template>
        </el-table-column>
        <el-table-column prop="date" label="등록일" align="center" width="140" />
        <el-table-column label="상태" align="center" width="60">
          <template #default="scope">
            <span v-if="scope.row.status == 2 || scope.row.type == 3">마감</span>
            <span v-if="scope.row.type != 3 && scope.row.status == 1 && scope.$index < data.items.length - 1">진행</span>
            <span v-if="scope.row.type != 3 && scope.row.status == 1 && scope.$index == data.items.length - 1" style="color:#af2020;">현재진행</span>            
          </template>
        </el-table-column>
        <el-table-column label="" align="center" width="120">
          <template #default="scope">
            <el-button v-if="scope.row.type != 3 && scope.row.status == 1" size="small" type="success" style="margin-right:0px;">장기수선</el-button>
          </template>
        </el-table-column>
        
      </el-table>
      
    </el-tab-pane>
    <el-tab-pane label="관리" name="management">


          
    <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">
      <div style="flex:1;text-align:right;gap:5;">
        <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;float:left;">삭제</el-button>
        <el-button v-if="data.level > 1" size="small" type="warning" @click="clickChangeStatus(1)" style="margin-right:0px;">마감 해제</el-button>
        <el-button v-if="data.level > 1" size="small" type="warning" @click="clickChangeStatus(2)" style="margin-right:0px;">마감 처리</el-button>
        <el-button size="small" type="info" @click="clickOther" style="margin-right:0px;">타업체 작업</el-button>
        <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">재수립 생성</el-button>
      </div>
    </div>  

    
  <el-table :data="data.items" border :width="data.width" :height="height(600)" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" />
    <el-table-column prop="type" label="구분" align="center" width="100">
      <template #default="scope">
        <el-tag :type="Repair.getTypeType(scope.row.type)">{{Repair.getType(scope.row.type)}}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="reportdate" label="리포트 작성일" align="center" />
    <el-table-column label="조정" width="60" align="center">
          <template #default="scope">
            <span v-if="scope.row.periodtype == 2">수시</span>
            <span v-else>정기</span>
          </template>
    </el-table-column>
    <el-table-column label="설명">
          <template #default="scope">
            {{scope.row.remark}}
          </template>
    </el-table-column>
    <el-table-column prop="date" label="등록일" align="center" width="140" />
    <el-table-column label="상태" align="center" width="60">
          <template #default="scope">
            <span v-if="scope.row.status == 2 || scope.row.type == 3">마감</span>
            <span v-if="scope.row.type != 3 && scope.row.status == 1 && scope.$index < data.items.length - 1">진행</span>
            <span v-if="scope.row.type != 3 && scope.row.status == 1 && scope.$index == data.items.length - 1" style="color:#af2020;">현재진행</span>            
          </template>
    </el-table-column>
    <el-table-column label="" align="center" width="120">
      <template #default="scope">
        <el-button v-if="scope.row.type != 3" size="small" type="warning" @click="clickDuplication(scope.row)" style="margin-right:0px;">검토조정 생성</el-button>
      </template>
    </el-table-column>
  </el-table>  

  
    </el-tab-pane>
  </el-tabs>





  <el-dialog
    v-model="data.visible"
    width="340px"
  >

    <el-form :model="[1]" label-width="140px">      
      <el-form-item label="리포트 작성 일자">
        <el-date-picker style="margin: 0px 0px;height: 24px;width:150px;" v-model="data.item.reportdate" placeholder="" />
      </el-form-item>

      <el-form-item label="설명">
        <el-radio-group v-model="data.item.periodtype">
          <el-radio :label="1" size="small">정기</el-radio>
          <el-radio :label="2" size="small">수시</el-radio>
        </el-radio-group>        
      </el-form-item>

      <el-form-item label="설명">
        <el-input v-model="data.item.remark" placeholder="" />
      </el-form-item>
    </el-form>

    
    <template #footer>
      <el-button size="small" @click="data.visible = false">취소</el-button>
      <el-button size="small" type="primary" @click="clickSubmitOther">등록</el-button>
    </template>
  </el-dialog>


</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted, computed, watch, watchEffect } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Apt, Repair, Category } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'

const { width, height } = size()

const props = defineProps({
  id: Number,
  close: Function
})

defineExpose({
  reset
})

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
  reportdate: null
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
  search: '',
  title: '',
  menu: 'list',
  level: 0
})

watch(() => props.id, () => {
  data.menu = 'list'
});

watchEffect(() => {
  data.apt = props.id
  getItems()
});

async function initData() {  
}

async function getItems() {
  if (data.apt == 0) {
    return
  }

  let res = await Apt.get(data.apt)
  data.title = res.item.name

  res = await Repair.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'r_reportdate,r_id'
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

function makeItems(items) {
  return items
}

function clickInsert() {
  util.confirm(data.title + '의 장기수선계획 재수립을 생성하시겠습니까', async function() {
    util.loading(true)
    
    await Category.init(data.apt)

    await getItems()

    util.loading(false)
  })
}

async function clickUpdate(item, index) {
  if (index.no == 0 || index.no == 8) {
    return
  }

  let res = await Repair.get(item.id)
  let repair = res.item


  console.log(data.menu)
  if (data.menu == 'management') {
    data.item = util.clone(item)
    data.visible = true
    return
  }

  if (repair.status == 2) {
    return
  }

  /*
  res = await Apt.get(item.apt)

  repair.info = res.item
  store.commit('setRepair', repair)
  */
  router.push(`/${repair.apt}/repair/${repair.id}/breakdown`)
}

function clickDelete() {
  let item = data.batchs[0]  
  util.confirm('삭제하시겠습니까', async function() {
    util.confirm('삭제한 데이터는 복구가 불가능합니다. 삭제하시겠습니까', async function() {
      let res = await Repair.remove(item)
      if (res.code === 'ok') {
        util.info('삭제되었습니다')
        getItems()
        
      }
    })
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

  data.menu = 'list'
}

function setWindowSize() {
  data.width = (window.innerWidth - 500) + 'px'
  data.height = (window.innerHeight - 170 - 300) + 'px'
}

onMounted(async () => {
  if (store.getters["getUser"] != null) {
    data.level = store.getters["getUser"].level
  }
  
  util.loading(true)
  
  await initData()
  await getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  util.loading(false)
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
})

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
  if (listSelection.value.length == 0) {
    util.error('선택된 항목이 없습니다')
    return
  }
  
  util.confirm('삭제하시겠습니까', async function() {
    util.confirm('삭제한 데이터는 복구가 불가능합니다. 삭제하시겠습니까', async function() {
      let items = []
      for (let i = 0; i < listSelection.value.length; i++) {
        let value = listSelection.value[i]

        let item = {
          id: value.id
        }

        items.push(item)      
      }

      await Repair.removebatch(items)
      util.info('삭제되었습니다')
      await getItems()
    })
  })
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

function clickDuplication(item) {
  util.confirm(data.title + '의 장기수선계획 검토조정을 생성하시겠습니까<BR/>생성시 약 5분 정도의 시간이 소요됩니다.<BR/>브라우져 창을 닫지 말고 기다려주세요', async function() {
    util.loading(true)
    
    await Category.duplication(item.id)
    await getItems()

    util.loading(false)
  })
}

function clickOther() {
  data.item = util.clone(item)
  data.item.type = 3
  data.visible = true
}

async function clickSubmitOther() {
  if (data.item.reportdate == null || data.item.reportdate == '') {
    util.error('리포트 작성일을 입력하세요')
    return
  }

  let item = util.clone(data.item)
  
  let year = item.reportdate  
  if (typeof year == 'string' || typeof year == 'number') {
    let d = new Date(year)
    item.reportdate = util.getDate(d)
  } else if (year == null || year == undefined || year == 0) {
    util.error('리포트 작성일을 입력하세요')
    return
  } else {
    item.reportdate = util.getDate(year)
  }
  
  item.periodtype = util.getInt(item.periodtype)
  item.apt = data.apt  

  if (item.id > 0) {
    await Repair.update(item)
  } else {
    item.status = 2
    await Repair.insert(item)
  }
  
  util.info('등록되었습니다')
  
  getItems()
  data.visible = false  
}

async function clickRepair(item) {
  let res = await Repair.get(item.id)
  let repair = res.item

  if (repair.type == 3) {
    return
  }

  if (repair.status == 2) {
    return
  }

  if (!util.isNull(props.close)) {
    props.close()
  }

  router.push(`/${repair.apt}/repair/${repair.id}/breakdown`)
}

function clickChangeStatus(status) {
  if (listSelection.value.length == 0) {
    util.error('선택된 항목이 없습니다')
    return
  }

  let title = ''

  if (status == 1) {
    title = '마감 해제'
  } else {
    title = '마감 처리'
  }
  util.confirm(title + ' 하시겠습니까', async function() {
    let items = []
    for (let i = 0; i < listSelection.value.length; i++) {
      let value = listSelection.value[i]

      await Repair.updateStatusById(status, value.id)
    }

    util.info(title + ' 되었습니다')
    getItems()
  })  
}

function reset() {
  data.menu = 'list';
}

</script>
