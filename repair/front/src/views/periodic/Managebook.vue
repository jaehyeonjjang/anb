<template>
  <Title title="시설물 관리대장" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert">등록</el-button>
      <el-button size="small" type="primary" @click="clickInsertMulti">일괄 등록</el-button>      
    </div>

  </div>  

  
  <el-table :data="data.items" border :height="height(170)" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" width="40" align="center" />
    <el-table-column prop="extra.managebookcategory.name" label="명칭" align="center" />    
    <el-table-column label="이미지" align="center">
      <template #default="scope">
        <el-image :src="util.getImagePath(scope.row.filename, `periodicresult/${data.id}`)" fit="cover" style="position:relative;top:5px;width:75px;height:50px;" @click="v3ImgPreviewFn(util.getImagePath(scope.row.filename, `periodicresult/${data.id}`))" />                                                               
      </template>
    </el-table-column>
  </el-table>


  <el-dialog
    v-model="data.visible"
    width="800px"
  >

      <y-table>
        <y-tr>
          <y-th>명칭</y-th>
          <y-td>
            <el-input v-model="data.item.name" />
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
              :data="{path:'periodic'}"
              :accept="'application/pdf'"
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


<el-dialog
    v-model="multi.visible"
    width="800px"
  >

      <el-upload
          class="upload-demo"
          ref="upload"
          :drag="true"
          :action="data.upload"
          :headers="headers"
          :on-exceed="handleExceed"
          :on-success="handelSuccess"
          :show-file-list="true"
          :auto-upload="true"
          :multiple="true"
          v-model:file-list="multi.files"
        >
          
          <el-icon class="el-icon--upload"><upload-filled /></el-icon>
          <div class="el-upload__text" style="font-size:12px;">
            파일을 드래그 하시거나 <em>여기를 클릭하세요</em>
          </div>          
      </el-upload>
      
      <template #footer>
        <el-button size="small" @click="clickCancelMulti">취소</el-button>
        <el-button size="small" type="primary" @click="clickSubmitMulti">등록</el-button>
      </template>
</el-dialog>



</template>


<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted } from "vue"
import router from '~/router'
import { util, size }  from "~/global"
import { Managebook } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Managebook

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const item = {
  id: 0,
  name: '',
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
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`,
  filename: ''
})

async function initData() {  
}

async function getItems() {
  let res = await model.find({
    page: data.page,
    pagesize: data.pagesize,
    periodic: data.id,
    orderby: 'mc_order,mc_name,mb_order,mb_id'
  })

  if (res.items != null) {   
    for (let i = 0; i < res.items.length; i++) {
      res.items[i].index = i + 1
    }
  }

  data.total = res.total
  if (res.items == null) {
    res.items = []
  }

  data.items = res.items
}

function clickInsert() {  
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

    let items = []
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

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {
}

async function handelSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  const filename = response.filename
  data.filename = filename    
}

const submitUpload = () => {
  upload.value.clearFiles()
  upload.value!.submit()
}

async function clickSubmit() {
  util.loading(true)
  await model.process(data.id, data.item.name, data.item.order, data.filename)

  await getItems()
  util.info('등록되었습니다')
  util.loading(false)

  data.visible = false
}

const multi = reactive({
  visible: false,
  files: []
})

function clickInsertMulti() {
  multi.visible = true
}

function clickCancelMulti() {  
  upload.value.clearFiles()
  
  multi.files = []
  multi.visible = false
}

async function clickSubmitMulti() {
  let filenames = []
  let originalfilenames = []
  
  for (let i = 0; i < multi.files.length; i++) {
    let item = multi.files[i];

    filenames.push(item.response.filename)
    originalfilenames.push(item.response.originalfilename)
  }

  util.loading(true)
  await model.multiprocess(data.id, filenames.join(','), originalfilenames.join(','))

  await getItems()
  util.info('등록되었습니다')
  util.loading(false)

  multi.visible = false  
}
</script>
