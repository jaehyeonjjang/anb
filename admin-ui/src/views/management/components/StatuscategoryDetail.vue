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

        <el-form-item label-width="100px" label="유형분류명">
          <el-input ref="name" v-model="form.name" type="text" placeholder="유형분류명을 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="구분">
          <el-radio-group v-model="form.type">
            <el-radio-button label="2">유형 - 빨강</el-radio-button>
            <el-radio-button label="12">유형 - 파랑</el-radio-button>
            <el-radio-button label="9">계단실 - 부위</el-radio-button>
            <el-radio-button label="8">계단실 - 부재</el-radio-button>
            <el-radio-button label="10">부위</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label-width="100px" label="정렬">
          <el-input ref="order" v-model="form.order" type="text" placeholder="" />
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
  order: '',
  type: 10
}

export default {
  name: 'StatuscategoryDetail',
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
      tempRoute: {}
    }
  },
  created() {
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
        url: '/api/statuscategory/' + id
      })

      this.form = response.data
    },
    clickSave: async function() {
      this.form.type = parseInt(this.form.type)
      this.form.company = this.$store.getters.company

      if (!this.form.name) {
        this.$alert('유형분류명을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.name.focus()
        return
      }

      if (this.isEdit) {
        const response = await request({
          method: 'PUT',
          url: '/api/statuscategory/' + this.form.id,
          data: this.form
        })

        if (response.code === 'CONFLICT') {
          this.$alert('이미 등록된 유형분류명입니다', '오류', { confirmButtonText: '확인' })
          this.$refs.name.focus()
          return
        }
      } else {
        const response = await request({
          method: 'POST',
          url: '/api/statuscategory',
          data: this.form
        })

        if (response.code === 'CONFLICT') {
          this.$alert('이미 등록된 유형분류명입니다', '오류', { confirmButtonText: '확인' })
          this.$refs.name.focus()
          return
        }
      }

      this.$router.push('/management/statuscategory')
    },
    clickCancel() {
      this.$router.go(-1)
    }
  }
}
</script>
