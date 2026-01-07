<template>
  <Container :height="80">  
    <Title title="순찰 도면" />

    <div style="margin-bottom:10px;">
      <el-button style="display:block;float:left;" size="small" type="primary" @click="clickSave()">저장</el-button>
      <el-button style="display:block;float:left;margin-left:10px;" size="small" @click="clickCancel()">취소</el-button>
      <el-button style="display:block;float:right;" size="small" type="success" @click="clickImageInsert(0)"><el-icon><Plus /></el-icon></el-button>
      <div style="clear:both;"></div>
    </div>
  <div v-for="(item, index) in data.images" :key="item.id" class="block" :style="{marginLeft: (item.level - 1) * 40 + 'px'}">
    <input v-model="item.name" type="text" class="input">

    <!--
    <el-dropdown style="display:none;float:right;margin-left:10px;" @command="(command) => clickImageCommand(command, index)">
      <span class="el-dropdown-link">
        More<i class="el-icon-arrow-down el-icon--right" />
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item command="copy">복사</el-dropdown-item>
        <el-dropdown-item command="children">하위작업 생성</el-dropdown-item>
        <el-dropdown-item v-if="item.filename" command="imageDelete" divided>도면 삭제</el-dropdown-item>
        <el-dropdown-item v-if="item.filename" command="imageView">도면 보기</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
    -->

    <div style="width:100%;">

      <el-upload
        class="plan-uploader"
        accept="image/jpeg,image/png"
        ref="upload"
        :data="{path:'blueprint'}"
        style="margin:0px 10px 0px 0px;height:20px;padding;0px 0px;"
        :action="data.upload"
        :headers="headers"
        :limit="1"
        :on-exceed="(file, files) => handleExceed(file, files, index)"
        :on-success="(res, file, files) => handleSuccess(res, file, files, index)"
        :show-file-list="false"
        :auto-upload="true"
        :before-upload="beforeImageUpload"
      >
        
        <img v-if="item.filename" :src="getImagePath(item.filename)" class="plan">
        <el-icon v-else class="btn"><Picture /></el-icon>        
      </el-upload>
      
      <el-icon class="btn" @click="clickImageInsert(index+1)"><Plus /></el-icon>
      <el-icon class="btn" @click="clickImageDelete(index)"><Close /></el-icon>           
      <el-icon class="btn" @click="clickImageLeft(index)"><ArrowLeft /></el-icon>
      <el-icon class="btn" @click="clickImageRight(index)"><ArrowRight /></el-icon>
      <el-icon class="btn" @click="clickImageUp(index)"><ArrowUp /></el-icon>
      <el-icon class="btn" @click="clickImageDown(index)"><ArrowDown /></el-icon>

    </div>
    <div class="clear" />
  </div>
  <el-dialog v-model="data.dialogVisible">
    <img width="100%" :src="dialogImageUrl" alt="">
  </el-dialog>

  <el-dialog title="하위작업 생성" v-model="data.dialogFormVisible">
    <el-form :model="form">
      <el-form-item label="Prefix" label-width="100px">
        <el-input v-model="data.prefix" autocomplete="off" />
      </el-form-item>
      <el-form-item label="" label-width="100px">
        <el-input v-model="data.start" autocomplete="off" style="width:100;" />
        ~
        <el-input v-model="data.end" autocomplete="off" style="width:100;" />
      </el-form-item>
      <el-form-item label="Postfix" label-width="100px">
        <el-input v-model="data.postfix" autocomplete="off" />
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="data.dialogFormVisible = false">Cancel</el-button>
      <el-button type="primary" @click="clickChildren">Confirm</el-button>
    </span>
  </el-dialog>

  

  </Container>
</template>

<script setup lang="ts">

import { ref, reactive, onMounted, onUnmounted, computed, watch, getCurrentInstance } from "vue"
import { util, size }  from "~/global"
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import type { UploadInstance } from 'element-plus'
import { Blueprint } from "~/models"

const { width, height } = size();
const store = useStore()
const route = useRoute()

const headers = {
  Authorization: 'Bearer ' + store.state.token
}

const data = reactive({
  apt: 0,
  masters: [],
  summarytypes: [],
  allsubmasters: [],
  submasters: [],
  allusers: [],
  users: [],
  prefix: '',
  postfix: '',
  start: '',
  end: '',
  dialogPos: 0,
  imageId: -1,
  dialogImageUrl: '',
  dialogVisible: false,
  dialogFormVisible: false,
  loading: true,
  tempRoute: {},
  aptgroups: [],
  imageUrl: '',
  images: [],
  upload: `${import.meta.env.VITE_REPORT_URL}/api/upload/index`
})

function initData() {
  let images = []
  images.push(getImage())
  data.images = images
}

async function getItems() {
  let res = await Blueprint.find({apt: data.apt, category: 3, orderby: 'bp_order,bp_id'})
  if (res.items == null) {
    res.items = []
  }

  data.images = res.items
}

onMounted(async () => {
  data.apt = parseInt(route.params.apt)
  
  util.loading(true)
  
  await initData()
  await getItems()  
  
  util.loading(false)
})


function getImagePath(filename) {
  return import.meta.env.VITE_REPORT_URL + '/webdata/' + filename
}

async function clickSave() {
  util.loading(true)
  
  let res = await Blueprint.find({apt: data.apt, category: 3})
  let blueprints = res.items

  if (blueprints == null) {
    blueprints = []
  }
  
  const oldImages = blueprints

  for (let i = 0; i < data.images.length; i++) {
    for (let j = 0; j < oldImages.length; j++) {
      if (data.images[i].id === oldImages[j].id) {
        oldImages[j].find = true
      }
    }
  }

  for (let j = 0; j < oldImages.length; j++) {
    if (oldImages[j].find === undefined) {
      await Blueprint.remove({id: oldImages[j].id})        
    }
  }

  for (let i = 0; i < data.images.length; i++) {
    const item = data.images[i]

    for (let j = i + 1; j < data.images.length; j++) {
      const item2 = data.images[j]

      if (item2.level <= item.level) {
        break
      }

      data.images[j].parent = item.id
    }
  }

  for (let i = 0; i < data.images.length; i++) {
    const item = data.images[i]

    item.apt = data.apt
    item.category = 3
    item.order = i
    if (item.filename != '') {
      item.upload = 1
    } else {
      item.upload = 0
    }

    const oldId = item.id
    let res

    if (item.id < 0) {
      res = await Blueprint.insert(item)      
    } else {
      res = await Blueprint.update(item)      
    }

    const newId = res.id

    if (oldId < 0) {
      for (let j = 0; j < data.images.length; j++) {
        if (data.images[j].parent === oldId) {
          data.images[j].parent = newId
        }
      }
    }
  }

  util.loading(false)
  util.info('저장되었습니다.')  
}

function clickCancel() {
  this.$router.go(-1)
}

function getImage() {
  data.imageId--
  return {
    id: data.imageId,
    apt: 0,
    name: '',
    level: 1,
    parent: 0,
    last: 0,
    title: '',
    type: '0',
    filename: '',
    order: 0,
    date: ''
  }
}

function getImageEnd(pos) {
  const item = data.images[pos]
  let end = pos
  for (let i = pos + 1; i < data.images.length; i++) {
    if (data.images[i].level <= item.level) {
      break
    }

    end = i
  }

  return end
}

function clickImageInsert(pos) {
  const image = getImage()

  if (pos > 0) {
    image.level = data.images[pos - 1].level
  }
  let images = data.images
  images.splice(pos, 0, image)

  data.images = images
}

function clickImageDelete(pos) {
  util.confirm('삭제하시겠습니까', async function() {
    const end = getImageEnd(pos)
    data.images.splice(pos, end - pos + 1)
  })
}

function clickImageDown(pos) {  
  if (pos >= data.images.length - 1) {
    return
  }

  const temp = data.images[pos]
  data.images[pos] = data.images[pos + 1]
  data.images[pos + 1] = temp
}

function clickImageUp(pos) {
  if (pos === 0) {
    return
  }

  const temp = data.images[pos]
  data.images[pos] = data.images[pos - 1]
  data.images[pos - 1] = temp
}

function clickImageLeft(pos) {
  const item = data.images[pos]

  if (item.level <= 1) {
    return
  }

  const end = getImageEnd(pos)

  for (let i = pos; i <= end; i++) {
    data.images[i].level--
  }
}

function clickImageRight(pos) {
  const item = data.images[pos]

  const max = 4

  if (pos === 0) {
    return
  }

  if (item.level >= max) {
    return
  }

  if (pos > 0) {
    if (item.level - data.images[pos - 1].level >= 1) {
      return
    }
  }

  let end = pos
  for (let i = pos + 1; i < data.images.length; i++) {
    if (data.images[i].level <= item.level) {
      break
    }

    if (data.images[i].level >= max) {
      return
    }

    end = i
  }

  for (let i = pos; i <= end; i++) {
    data.images[i].level++
  }
}

const upload = ref<UploadInstance>()

const handleExceed: UploadProps['onExceed'] = (files, uploadFiles, index) => {
}

function handleSuccess(response: any, file: UploadFile, files: UploadFiles, pos: number) {
  const filename = response.filename

  for (let i = 0; i < upload.value.length; i++) {
    upload.value[i].clearFiles()
  }
  
  
  let images = util.clone(data.images)
  images[pos].filename = filename
  
  data.images = images
}

function beforeImageUpload(file) {
  const isImage = (file.type === 'image/jpeg' || file.type === 'image/png')

  if (!isImage) {
    util.error('이미지 파일만 업로드 가능합니다 (jpg, png)')
  }

  return isImage
}

const submitUpload = (index) => {    
}

function clickImageCommand(command, pos) {
  if (command === 'copy') {
    util.confirm('삭제하시겠습니까', async function() {
      const end = getImageEnd(pos)

      const items = []

      for (var i = pos; i <= end; i++) {
        const item = JSON.parse(JSON.stringify(data.images[i]))
        data.imageId--
                     item.id = data.imageId
        items.push(item)
      }

      data.images.splice(end + 1, 0, ...items)

    })
  } else if (command === 'children') {
    data.dialogFormVisible = true
    data.dialogPos = pos
    data.prefix = ''
    data.postfix = ''
    data.start = ''
    data.end = ''
  } else if (command === 'imageView') {
    const image = data.images[pos]
    const filename = this.getImagePath(image.filename)
    data.dialogImageUrl = filename
    data.dialogVisible = true
  } else if (command === 'imageDelete') {
    data.images[pos].filename = ''
  }
}

function clickChildren() {
  const item = data.images[data.dialogPos]

  const start = parseInt(data.start)
  const end = parseInt(data.end)

  const items = []
  for (var i = start; i <= end; i++) {
    const image = getImage()

    image.level = item.level + 1
    image.name = data.prefix + i + data.postfix
    items.push(image)
  }

  data.images.splice(data.dialogPos + 1, 0, ...items)

  data.dialogFormVisible = false
}

async function clickCopy() {
  const id = this.$route.params && this.$route.params.id
  util.confirm('삭제하시겠습니까', async function() {
    await request({
      method: 'POST',
      url: '/api/apt/' + id + '/copy'
    })

    util.info('복사되었습니다')

    getList()
  })
}

function clickTest() {
  console.log(data.images)
}

</script>

<style lang="scss" scoped>

.block {
  border: 1px solid #aaa;
  border-radius: 3px;
  padding: 10px 10px;
  margin-bottom: 10px;  
  display: flex;
  justify-content: space-between;

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
}

.plan-uploader {
  float:right;
  margin-top: 4px;
  margin-right: 20px;
  height: 34px;

  .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }

  .el-upload:hover {
    border-color: #409EFF;
  }
}

.plan-uploader-icon {
  font-size: 24px;
  color: #8c939d;
  width: 32px;
  height: 32px;
  line-height: 32px;
  text-align: center;
}
.plan {
  width: 24px;
  height: 24px;
  display: block;
}
</style>
