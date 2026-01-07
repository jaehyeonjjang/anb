<template>
  <div class="createPost-container">
    <el-form ref="form" :model="form" class="form-container">

      <sticky :z-index="10" :class-name="'sub-navbar '">
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" icon="el-icon-check" @click="clickSave">
          저장
        </el-button>
        <el-button v-loading="loading" type="warning" icon="el-icon-close" @click="clickCancel">
          취소
        </el-button>
      </sticky>

      <div class="createPost-main-container">

        <el-form-item v-if="isEdit" label-width="100px" label="ID">
          <el-input v-model="form.id" type="text" disabled />
        </el-form-item>

        <el-form-item label-width="100px" label="현장명">
          <el-input ref="name" v-model="form.name" type="text" placeholder="현장명을 입력하세요" />
        </el-form-item>

        <el-form-item v-if="$store.getters.isAdmin" label-width="100px" label="업체">
          <el-select v-model="form.company" style="float:left;margin-right:10px;">
            <el-option
              v-for="item in companys"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label-width="100px" label="시설물 구분">
          <el-radio-group v-model="form.facility">
            <el-radio-button label="1">1종</el-radio-button>
            <el-radio-button label="2">2종</el-radio-button>
            <el-radio-button label="3">3종</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label-width="100px" label="종류">
          <el-radio-group v-model="form.type">
            <el-radio-button label="1">대형건축물</el-radio-button>
            <el-radio-button label="2">공동주택</el-radio-button>
            <el-radio-button label="3">단독주택</el-radio-button>
            <el-radio-button label="4">주상복합</el-radio-button>
            <el-radio-button label="5">업무시설</el-radio-button>
            <el-radio-button label="6">다중이용건축물</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label-width="100px" label="상태">
          <el-radio-group v-model="form.status">
            <el-radio-button label="1">사용</el-radio-button>
            <el-radio-button label="2">사용 안함</el-radio-button>
          </el-radio-group>
        </el-form-item>


        <el-form-item label-width="100px" label="사진 분류">
          <el-input
            type="textarea"
            :rows="10"
            placeholder="사진 분류를 입력하세요"
            v-model="form.imagecategory">
          </el-input>
        </el-form-item>

      </div>
    </el-form>
  </div>
</template>

<script>
import Sticky from '@/components/Sticky'
import request from '@/utils/request'

const defaultForm = {
  id: undefined,
  name: '',
  company: 0,
  status: 1,
  facility: 1,
  type: 1,
  imagecategory: ''
}

export default {
  name: 'AptgroupDetail',
  components: { Sticky },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      form: Object.assign({}, defaultForm),
      loading: true,
      tempRoute: {},
      companys: []
    }
  },
  created: async function() {
    const response = await request({
      method: 'GET',
      url: '/api/company'
    })

    this.companys = [{ id: 0, name: '업체' }, ...response.data]

    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)
    }

    this.tempRoute = Object.assign({}, this.$route)

    this.loading = false
  },
  mounted() {
    this.$refs.name.focus()
  },
  methods: {
    fetchData: async function(id) {
      const response = await request({
        method: 'GET',
        url: '/api/aptgroup/' + id
      })

      this.form = response.data
    },
    clickSave: async function() {
      this.form.status = parseInt(this.form.status)
      this.form.company = parseInt(this.form.company)

      this.form.updateuser = this.$store.getters.id

      if (!this.$store.getters.isAdmin) {
        this.form.company = this.$store.getters.company
      }

      if (!this.form.name) {
        this.$alert('현장명을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.name.focus()
        return
      }

      if (this.isEdit) {
        const response = await request({
          method: 'PUT',
          url: '/api/aptgroup/' + this.form.id,
          data: this.form
        })

        if (response.code === 'CONFLICT') {
          this.$alert('이미 등록된 현장입니다', '오류', { confirmButtonText: '확인' })
          this.$refs.loginid.focus()
          return
        }
      } else {
        this.form.user = this.$store.getters.id

        const response = await request({
          method: 'POST',
          url: '/api/aptgroup',
          data: this.form
        })

        if (response.code === 'CONFLICT') {
          this.$alert('이미 등록된 현장입니다', '오류', { confirmButtonText: '확인' })
          this.$refs.loginid.focus()
          return
        }
      }

      this.$router.push('/aptgroup/aptgroup')
    },
    clickCancel() {
      this.$router.go(-1)
    }
  }
}
</script>
