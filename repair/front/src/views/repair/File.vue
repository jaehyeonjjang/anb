<template>
  <Title title="파일" />

  <div style="display:flex;justify-content: space-between;gap:5px;margin-bottom:10px;">    
    <div style="flex:1;text-align:right;gap:5;">
      <el-button size="small" type="danger" @click="clickDeleteMulti" style="margin-right:-5px;">삭제</el-button>
      <el-button size="small" type="success" @click="clickInsert" style="margin-right:0px;">등록</el-button>
    </div>
  </div>  

  
  <el-table :data="data.items" border :width="data.width" :height="data.height" @row-click="clickUpdate" :key="data.width+''+data.height" ref="listRef" @selection-change="changeList">
    <el-table-column type="selection" width="40" align="center" />
    <el-table-column prop="index" label="NO" align="center" width="60" />
    <el-table-column prop="title" label="제목" />
    <el-table-column prop="originalfilename" label="파일명" />
    <el-table-column label="다운로드" align="center" width="80">
      <template #default="scope">
        <el-button size="small" type="warning" @click="clickDownload(scope.row)">다운로드</el-button>
      </template>
    </el-table-column>
    <el-table-column prop="date" label="등록일" align="center" width="140" />        
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    :before-close="handleClose"
    width="800px"
  >

    <el-form label-width="100px">      
      <el-table :data="data.batchs" border :max-height="data.height" :key="data.width+''+data.height" style="margin-top:15px;">
        <el-table-column label="" align="center" width="35" v-if="data.mode == 'batch'">
          <template #default="scope">
            <el-icon @click="clickRegistDelete(scope.$index)"><Delete /></el-icon>
          </template>
        </el-table-column>
        <el-table-column label="제목" align="center">
          <template #default="scope">
            <el-input v-model="data.batchs[scope.$index].title" />
          </template>
        </el-table-column>
        <el-table-column label="파일" align="center" width="120">
          <template #default="scope">

            <el-upload
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
              <el-button size="small" type="danger" @click="submitUpload">파일 업로드</el-button>
              
            </el-upload>                
                
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
import { util }  from "~/global"
import { File } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import { ElTable, UploadInstance } from 'element-plus'
import axios from 'axios'

const store = useStore()
const route = useRoute()

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

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
  title: '',
  filename: '',
  originalfilensme: ''  
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
  filename: '',
  originalfilename: '',
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`                      
})

async function initData() {  
}

async function getItems() {
  let res = await File.find({
    page: data.page,
    pagesize: data.pagesize,
    apt: data.apt,
    orderby: 'f_id'    
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

function makeItems(items) {
  return items
}
function clickInsert() {  
  data.item = util.clone(item)

  let items = makeItems([data.item])

  data.filename = ''
  data.originalfilename = ''
  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickUpdate(item, index) {
  if (index.no == 0) {
    return
  }

  if (index.no == 4) {
    return
  }

  let items = makeItems([util.clone(item)])

  data.filename = ''
  data.originalfilename = ''
  data.mode = 'normal'
  data.batchs = items
  data.visible = true  
}

function clickDelete() {
  let item = data.batchs[0]
  
  util.confirm('삭제하시겠습니까', async function() {
    let res = await File.remove(item)
    if (res.code === 'ok') {
      util.info('삭제되었습니다')
      data.visible = false
      getItems()
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

function setWindowSize() {
  data.width = (window.innerWidth - 500) + 'px'
  data.height = (window.innerHeight - 170) + 'px'
}

onMounted(async () => {
  data.apt = parseInt(route.params.id)  
  
  util.loading(true)
  
  await initData()
  await getItems()

  setWindowSize()

  window.addEventListener('resize', setWindowSize)

  data.visible = false
  util.loading(false)
})

onUnmounted(() => {
  window.removeEventListener('resize', setWindowSize)
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

  items = makeItems(items)
  
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

      await File.remove(item)
    }

    util.info('삭제되었습니다')
    getItems()

    util.loading(false)
  })
}

async function clickSubmit() {
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]
    
    if (item.title == '') {
      util.error('제목을 입력하세요')
      return    
    }
  }
  
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
        await File.remove(item)
      }
    }
  }
  
  for (let i = 0; i < data.batchs.length; i++) {
    let item = data.batchs[i]

    item.apt = data.apt
    

    if (item.id > 0) {
      if (data.filename != '') {
        item.filename = data.filename
        item.originalfilename = data.originalfilename
      }      
      await File.update(item)
    } else {
      item.filename = data.filename
      item.originalfilename = data.originalfilename
      await File.insert(item)
    }
  }

  util.info('등록되었습니다')
  
  getItems()
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

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles) => {  
}

async function handelSuccess(response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) {
  console.log(response.filename)
  console.log(response.originalfilename)
  data.filename = response.filename
  data.originalfilename = response.originalfilename
}

const submitUpload = () => {
  upload.value.clearFiles()
  upload.value!.submit()
}

function clickDownload(item) {
  axios.get(import.meta.env.VITE_REPORT_URL + '/api/download/file/' + item.id, {
    responseType: 'blob',
    headers: {
      Authorization: 'Bearer ' + store.state.token
    }
  }).then(response => {
    const url = window.URL.createObjectURL(new Blob([response.data]))
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', item.originalfilename);
    document.body.appendChild(link)
    link.click()
  }).catch(exception => {
    alert("파일 다운로드 실패");
  });
}

</script>
