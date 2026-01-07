<template>
  <div class="flex-container">
    <el-card class="box-card">
      <el-form label-width="120px">
        <el-form-item label="Loginid">
          <el-input v-model="item.loginid" />
        </el-form-item>
        <el-form-item label="Password">
          <el-input v-model="item.passwd" show-password @keypress.enter.native="clickSignin" />
        </el-form-item>
        <el-button type="primary" @click="clickSignin">Sign In</el-button>
      </el-form>
    </el-card>    
  </div>

  <div style="margin-top:-80px;" @click="clickDownload">
    프로그램 다운 로드
  </div>


  
  <el-dialog v-model="data.visibleDownload" width="800px" title="프로그램 다운로드">

    <y-table>
      <y-tr>
        <y-th style="text-align:center;">구분</y-th>
        <y-th style="text-align:center;">Version</y-th>
        <y-th style="text-align:center;">다운로드</y-th>
      </y-tr>
      <y-tr>
        <y-td style="text-align:center;">정기점검 프로그램</y-td>
        <y-td style="text-align:center;">V{{data.periodicProgram}}</y-td>
        <y-td style="text-align:center;"><el-button size="small" @click="clickDownloadPeriodic">다운로드</el-button></y-td>
      </y-tr>
      <y-tr>
        <y-td style="text-align:center;">순찰 프로그램</y-td>
        <y-td style="text-align:center;">V{{data.patrolProgram}}</y-td>
        <y-td style="text-align:center;"><el-button size="small" @click="clickDownloadPatrol">다운로드</el-button></y-td>
      </y-tr>
    </y-table>
    
    <template #footer>
      <el-button size="small" @click="data.visibleDownload = false">닫기</el-button>
    </template>
  </el-dialog>
  
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useStore } from 'vuex'
import { util } from '~/global'
import { Login, Program } from "~/models"
import router from '~/router'

const store = useStore()

const item = reactive({
  loginid: '',
  passwd: ''
})

const data = reactive({
  visibleDownload: false  
})



async function clickSignin() {
  if (item.loginid === '') {
    util.error('로그인 아이디를 입력하세요')
    return
  }

  if (item.passwd === '') {
    util.error('패스워드를 입력하세요')
    return
  }  

  const res = await Login.login(item)
  if (res.code === 'ok') {
    store.commit('setRepair', null)
    util.login(store, res)
    router.push('/')
  } else {
    console.log(res);
    util.error('로그인 정보가 정확하지 않습니다')
  }
}

async function clickDownload() {
  let res = await Program.find({orderby: 'p_id desc'})

  let items = res.items

  for (let i = 0; i < items.length; i++) {
    let item = items[i]

    if (item.type == 1) {
      data.periodicProgram = item.version
    }

    if (item.type == 3) {
      data.patrolProgram = item.version
    }
  }

  data.visibleDownload = true
}

function clickDownloadPeriodic() {
  let version = data.periodicProgram
  const url = `/webdata/apk/periodic-V${version}.apk`
  const filename = `ANB-정기점검프로그램-V${version}.apk`

  util.download(store, url, filename)  
}

function clickDownloadPatrol() {
  let version = data.patrolProgram
  const url = `/webdata/apk/patrol-V${version}.apk`
  const filename = `ANB-순찰프로그램-V${version}.apk`

  util.download(store, url, filename)  
}
</script>
