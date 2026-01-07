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

        <el-form-item label-width="100px" label="유형명">
          <el-input ref="name" v-model="form.name" type="text" placeholder="유형명을 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="구분">
          <el-radio-group v-model="form.type" @input="changeType">
            <el-radio-button label="10">부위</el-radio-button>
            <el-radio-button label="1">부재 - 빨강</el-radio-button>
            <el-radio-button label="2">유형 - 빨강</el-radio-button>
            <el-radio-button label="3">폭</el-radio-button>
            <el-radio-button label="4">길이</el-radio-button>
            <el-radio-button label="5">개소</el-radio-button>
            <el-radio-button label="6">진행사항</el-radio-button>
            <el-radio-button label="7">비고</el-radio-button>
            <el-radio-button label="9">계단실 - 부위</el-radio-button>
            <el-radio-button label="8">계단실 - 부재</el-radio-button>
            <el-radio-button label="11">부재 - 파랑</el-radio-button>
            <el-radio-button label="12">유형 - 파랑</el-radio-button>
            <el-radio-button label="21">하자조사 - 부위</el-radio-button>
            <el-radio-button label="22">하자조사 - 부재</el-radio-button>
            <el-radio-button label="23">하자조사 - 하자</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="form.type == 2 || form.type == 9 || form.type == 8 || form.type == 12 || form.type == 10" label-width="100px" label="구분">
          <el-select v-model="form.statuscategory" style="float:left;margin-right:10px;">
            <el-option
              v-for="item in statuscategorys"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label-width="100px" label="이미지">
          <el-input ref="content" v-model="form.content" type="text" placeholder="" />
        </el-form-item>

        <el-form-item label-width="100px" label="내용">
          <el-input ref="etc" v-model="form.etc" type="text" placeholder="" />
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
  type: 10,
  content: '',
  etc: ''
}

export default {
  name: 'StatusDetail',
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
      statuscategorys: [],
      allstatuscategorys: []
    }
  },
  created: async function() {
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)
    }

    this.tempRoute = Object.assign({}, this.$route)
    this.changeType()

    this.loading = false
  },
  mounted() {
    this.$refs.name.focus()
  },
  methods: {
    fetchData: async function(id) {
      const response = await request({
        method: 'GET',
        url: '/api/status/' + id
      })

      this.form = response.data

      const statuscategory = this.form.statuscategory
      await this.changeType()
      this.form.statuscategory = statuscategory
    },
    clickSave: async function() {
      this.form.type = parseInt(this.form.type)
      this.form.company = this.$store.getters.company

      if (!this.form.name) {
        this.$alert('유형명을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.name.focus()
        return
      }

      if (this.isEdit) {
        await request({
          method: 'PUT',
          url: '/api/status/' + this.form.id,
          data: this.form
        })
      } else {
        await request({
          method: 'POST',
          url: '/api/status',
          data: this.form
        })
      }

      this.$router.push('/management/status')
    },
    clickCancel() {
      this.$router.go(-1)
    },
    changeType: async function() {
      const params = `type=${this.form.type}&company=` + this.$store.getters.company
      const response = await request({
        method: 'GET',
        url: '/api/statuscategory?' + params
      })

      this.form.statuscategory = 0
      this.statuscategorys = [{ id: 0, name: '유형분류' }, ...response.data]
    }
  }
}
</script>
