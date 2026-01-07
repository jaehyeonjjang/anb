<template>
  <div style="display:flex;justify-content:space-between;margin-bottom:10px;">
    <div></div>
    <el-button size="small" type="success" @click="clickUpdate">수정</el-button>
  </div>

  <el-table :data="data.items" border>
    <el-table-column prop="position" label="점검내용" width="180" />
    <el-table-column label="상태" align="left">
      <template #default="scope">
        <div style="display:flex;">
          <div style="margin-right:20px;" v-for="item in getList(scope.row.name, scope.row.status)"><span v-if="item.checked == true">▣</span><span v-else>□</span> {{item.name}}</div>
        </div>
      </template>
    </el-table-column>      
    <el-table-column label="이미지" align="center" width="150">
      <template #default="scope">
        <div style="height:30px;">
          <div v-if="scope.row.filename != ''">
            <el-image v-for="(item, index) in scope.row.filename.split(',')" style="width: 20px; height: 20px; top:4px;left:0px;position:relative;margin-right:5px;"
                      :src="util.getImagePath(item)"                    
                      fit="cover"
                      @click="clickPreviews(scope.row.filename, index)"
            />
          </div>          
        </div>
      </template>
    </el-table-column>        
  </el-table>  

  
  <el-dialog
    v-model="data.visible"
    width="1000px"
  >


    <y-table>
      <y-tr>
        <y-th style="text-align:center;width:150px;">점검내용</y-th>
        <y-th>상태</y-th>        
        <y-th style="text-align:center;width:120px;">이미지</y-th>
      </y-tr>
      <y-tr v-for="(item, index) in data.batchs" :key="item.id">
        <y-td>{{data.batchs[index].position}}</y-td>
              
        <y-td>
          <div style="display:flex;" v-if="item.type != 1">
            <div style="margin-right:20px;" v-for="(item, pos) in getInput(item.name, item.status)"><el-checkbox size="small" v-model="data.batchs[index][`check${pos}`]" /> {{item.name}}</div>
          </div>
          <div style="display:flex;" v-else>
            <el-radio-group v-model="data.batchs[index].status">
              <el-radio  v-for="(item, pos) in getInput(item.name, item.status)" :label="item.name">{{item.name}}</el-radio>                
            </el-radio-group>
          </div>
        </y-td>
        
        <y-td style="text-align:center;">
          <div style="height:30px;">
            <el-upload
              accept="image/jpeg,image/png"
              ref="upload"
              style="display:block;float:left;margin:5px 10px 0px 5px;"
              :data="{path:'periodic'}"
              :action="data.upload"
              :headers="headers"
              :limit="1"
              :on-exceed="(file, files) => handleExceed(file, files, index)"
              :on-success="(res, file, files) => handleSuccess(res, file, files, index)"
              :show-file-list="false"
              :auto-upload="true"
              :before-upload="beforeImageUpload"
            >

              
              <el-icon style="font-size:20px;"><Plus /></el-icon>
            </el-upload>

            
            <div v-if="data.batchs[index].filename != ''" style="float:left;">
              <el-image v-for="(item, pos) in data.batchs[index].filename.split(',')" style="width: 20px; height: 20px; top:4px;left:0px;position:relative;margin-right:5px;"
                        :src="util.getImagePath(item)"                    
                        fit="cover"
                        @click="clickImageDelete(data.batchs[index].filename, index, pos)"
              />
            </div>          
          </div>
        </y-td>
        
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
import { Periodicother } from "~/models"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import request from '~/global/request'
import type { UploadInstance } from 'element-plus'
import { v3ImgPreviewFn } from 'v3-img-preview'

const { width, height } = size()

const store = useStore()
const route = useRoute()

const model = Periodicother

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const item = {
  id: 0,
  name: '',
  type: 0,
  result: 0,
  status: '',
  position: '',
  filename: '',
  offlinefilename: '',
  category: 0,
  order: 0,
  periodic: 0,
  date: ''
}

const data = reactive({
  id: 0,
  item: util.clone(item),
  visible: false,
  items: [],
  batchs: [],
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`
})

async function initData() {
}

async function getItems() {
  let res = await model.find({
    periodic: data.id,
    category: 11,
    orderby: 'po_order,po_id'    
  })

  if (res.items == null) {
    res.items = []
  }
  
  data.items = res.items

  data.uploads = new Array(res.items.length)    
}

async function clickUpdate(pos) {  
  let items = util.clone(data.items)

  for (let k = 0; k < items.length; k++) {
    let names = items[k].name.split(',')
    let values = items[k].status.split(',')
  
    for (let i = 0; i < names.length; i++) {
      let flag = false
      for (let j = 0; j < values.length; j++) {
        if (names[i] == values[j]) {
          flag = true
          break
        }
      }

      items[k][`check${i}`] = flag
    }    
  }

  data.batchs = items
  data.visible = true
}

async function clickSubmit(type) {  
  util.loading(true)

  for (let i = 0; i < data.batchs.length; i++) {
    let item = util.clone(data.batchs[i])

    if (item.type != 1) {
      let names = item.name.split(',')

      let values = [] 
      for (let k = 0; k < names.length; k++) {
        if (item[`check${k}`] == true) {
          values.push(names[k])
        }    
      }

      item.status = values.join(',')
    }

    item.result = util.getInt(item.result)
    await model.update(item)

  }
  
  data.visible = false

  await getItems(true)
  util.info('수정되었습니다')
  util.loading(false)
}

onMounted(async () => {
  util.loading(true)
  
  const apt = parseInt(route.params.apt)
  const id = parseInt(route.params.id)
  
  data.id = id
  data.apt = apt

  await initData()
  await getItems()

  util.loading(false)
})

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles, index) => {  
}

function handleSuccess(response: any, file: UploadFile, files: UploadFiles, index: number) {
  const filename = response.filename
  
  let items = util.clone(data.batchs)

  if (items[index].filename == '') {
    items[index].filename = filename
  } else {
    items[index].filename = items[index].filename + ',' + filename
  }
  
  data.batchs = items

  for (let i = 0; i < upload.value!.length; i++) {
    upload.value![i].clearFiles(['success'])
  }  
}

function beforeImageUpload(file) {
  const isImage = (file.type === 'image/jpeg' || file.type === 'image/png')

  if (!isImage) {
    util.error('이미지 파일만 업로드 가능합니다 (jpg, png)')
  }

  return isImage
}

function clickPreviews(str, index) {
  const imgs = str.split(',').map(item => util.getImagePath(item)) 
  v3ImgPreviewFn({images:imgs, index: index})  
}

function clickImageDelete(item, index, pos) {
  util.confirm('삭제하시겠습니까', function() {
    let items = util.clone(data.batchs)
    let temp = items[index].filename.split(',')
    temp.splice(pos, 1)

    items[index].filename = temp.join(',')
    data.batchs = items
    util.info('삭제되었습니다')
  })
}

function getList(name, value) {
  let names = name.split(',')
  let values = value.split(',')

  let items = []
  for (let i = 0; i < names.length; i++) {
    let flag = false
    for (let j = 0; j < values.length; j++) {
      if (names[i] == values[j]) {
        flag = true
        break
      }
    }

    items.push({
      name: names[i],
      checked: flag
    })

  }

  return items
}

function getInput(name, value) {
  let names = name.split(',')
  let values = value.split(',')

  let items = []
  for (let i = 0; i < names.length; i++) {
    let flag = false
    for (let j = 0; j < values.length; j++) {
      if (names[i] == values[j]) {
        flag = true
        break
      }
    }

    items.push({
      name: names[i],
      checked: flag
    })

  }

  return items
}

</script>
