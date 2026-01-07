<script setup lang="ts">
import { reactive, onMounted } from "vue"
import { request, util } from "~/global"

const item = reactive({
  loginid: '',
  passwd: ''
})

function clickSignin() {
  if (item.loginid === '') {
    util.error('로그인 아이디를 입력하세요')
    return
  }

  const res = request({
    method: "GET",
    url: "api/user"
  });
}

</script>

<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="Loginid">
      <el-input v-model="item.loginid" />
    </el-form-item>
    <el-form-item label="Password">
      <el-input v-model="item.passwd" />
    </el-form-item>
    <el-button type="primary" @click="clickSignin">Sign In</el-button>
  </el-form>
</template>
