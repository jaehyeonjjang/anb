<template>
  <Title title="사진자료" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert">등록</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :height="height(170)" @row-click="clickUpdate"  ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />    
    <el-table-column prop="floor" label="구분" align="center">
      <template #default="scope">
        <span v-if="scope.row.type==1">위치도</span>
        <span v-if="scope.row.type==2">전경</span>
        <span v-if="scope.row.type==3">부위별</span>
        <span v-if="scope.row.type==10">주변공사</span>        
        <span v-if="scope.row.type==4">과업수행계획서 승인</span>
        <span v-if="scope.row.type==5">전회점검 결과표</span>
      </template>
    </el-table-column>
    <el-table-column prop="name" label="명칭" align="center" />
    <el-table-column label="이미지" align="center" width="60">
      <template #default="scope">
        <el-image :src="util.getImagePath(scope.row.filename)" fit="cover" style="width:50px;height:50px;" @click="v3ImgPreviewFn(util.getImagePath(scope.row.filename))" />                                                               
      </template>
    </el-table-column>
    <el-table-column prop="floor" label="사용여부" align="center">
      <template #default="scope">
        <span v-if="scope.row.use==1">사용</span>
        <span v-if="scope.row.use==2">사용 안함</span>
      </template>
        </el-table-column>
        <el-table-column prop="order" label="순번" align="center" />
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    width="800px"
  >

      <y-table>
        <y-tr>
          <y-th>구분</y-th>
          <y-td>
            <el-radio-group v-model.number="data.item.type">
              <el-radio-button size="small" label="1">위치도</el-radio-button>
              <el-radio-button size="small" label="2">전경</el-radio-button>
              <el-radio-button size="small" label="3">부위별</el-radio-button>
              <el-radio-button size="small" label="10">주변공사</el-radio-button>
            </el-radio-group>
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>명칭</y-th>
          <y-td>
            <el-input v-model="data.item.name" />
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>사용 여부</y-th>
          <y-td>
            <el-radio-group v-model.number="data.item.use">
              <el-radio-button size="small" label="1">사용</el-radio-button>
              <el-radio-button size="small" label="2">사용 안함</el-radio-button>
            </el-radio-group>
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>순번</y-th>
          <y-td>
            <el-input v-model.number="data.item.order" />
          </y-td>
        </y-tr>
        <y-tr>
          <y-th>이미지</y-th>
          <y-td>

            <el-upload
              accept="image/jpeg,image/png"
              style="float:left;"
              class="upload-demo"
              ref="upload"
              :data="{path:'periodic'}"
              :action="data.upload"
              :headers="headers"
              :limit="1"
              :on-exceed="handleExceed"
              :on-success="handelSuccess"
              :show-file-list="false"
              :auto-upload="true"              
            >
              <el-button size="small" type="danger" @click="submitUpload">이미지 업로드</el-button>

            </el-upload>

          </y-td>
        </y-tr>
      </y-table>

      <template #footer>
        <el-button size="small" @click="clickCancel">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmit">등록</el-button>
      </template>
  </el-dialog>

</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Periodicimage } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodicimage

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const item = {
  id: 0,
  type: 2,
  filename: '',
  name: '',
  use: 1,
  order: 0,
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
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`
})

async function initData() {  
}

async function getItems() {
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    periodic: data.id,
    orderby: 'pi_type,pi_order,pi_id'
  })

  if (res.items == null) {
    res.items = []
   }

  let items = []
  
  for (let i = 0; i < res.items.length; i++) {
    let item = res.items[i]

    if (item.type > 3 && item.type != 10) {
      continue
    }
    
    item.index = i + 1
    items.push(item)
  }

  data.total = res.total
  data.items = items
}

function clickInsert() {  
  data.item = util.clone(item)
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0 || index.no == 3) {
    return
  }

  data.item = util.clone(item)
  data.visible = true  
}

onMounted(async () => {
  data.apt = parseInt(route.params.apt)
  data.id = parseInt(route.params.id)
  
  util.loading(true)
  
  await initData()
  await getItems()

  data.visible = false
  util.loading(false)
})

function clickCancel() {
  data.visible = false
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

  let item = util.clone(data.item)

  item.periodic = data.id
  item.type = util.getInt(item.type)
  item.use = util.getInt(item.use)
  item.order = util.getInt(item.order)

  if (item.id > 0) {
    await model.update(item)
  } else {
    await model.insert(item)
  }

  util.info('등록되었습니다')
  
  await getItems()

  data.visible = false  
  util.loading(false)  
}

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
}

async function handelSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  data.item.filename = response.filename
}

const submitUpload = () => {
  upload.value.clearFiles()
  upload.value!.submit()
}

</script>
