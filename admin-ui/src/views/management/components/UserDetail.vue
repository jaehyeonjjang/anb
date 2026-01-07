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

        <el-form-item label-width="100px" label="이메일">
          <el-input ref="loginid" v-model="form.loginid" type="text" placeholder="이메일을 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="비밀번호">
          <el-input ref="passwd" v-model="form.passwd" type="password" placeholder="비밀번호를 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="비밀번호 확인">
          <el-input ref="passwd2" v-model="form.passwd2" type="password" placeholder="비밀번호를 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="이름">
          <el-input ref="name" v-model="form.name" type="text" placeholder="이름을 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="핸드폰">
          <el-input ref="hp" v-model="form.hp" type="text" placeholder="핸드폰 번호를 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="기술등급">
          <el-radio-group v-model="form.grade">
            <el-radio-button label="1">없음</el-radio-button>
            <el-radio-button label="2">건축초급기술자</el-radio-button>
            <el-radio-button label="3">건축중급기술자</el-radio-button>
            <el-radio-button label="4">건축고급기술자</el-radio-button>
            <el-radio-button label="5">건축특급기술자</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label-width="100px" label="레벨">
          <el-radio-group v-model="form.level">
            <el-radio-button label="1">작업자</el-radio-button>
            <el-radio-button label="2">매니저</el-radio-button>
            <el-radio-button label="3">관리자</el-radio-button>
            <el-radio-button v-if="$store.getters.isAdmin" label="4">총 관리자</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="form.level != 4 && $store.getters.isAdmin" label-width="100px" label="업체">
          <el-select v-model="form.company" style="float:left;margin-right:10px;">
            <el-option
              v-for="item in companys"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label-width="100px" label="상태">
          <el-radio-group v-model="form.status">
            <el-radio-button label="1">사용</el-radio-button>
            <el-radio-button label="2">사용 안함</el-radio-button>
          </el-radio-group>
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
  level: 3,
  company: 0,
  status: 1,
  grade: 1
}

export default {
  name: 'UserDetail',
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
        url: '/api/user/' + id
      })

      this.form = response.data
    },
    clickSave: async function() {
      this.form.grade = parseInt(this.form.grade)
      this.form.level = parseInt(this.form.level)
      this.form.status = parseInt(this.form.status)
      this.form.company = parseInt(this.form.company)
      this.form.email = this.form.loginid

      if (!this.$store.getters.isAdmin) {
        this.form.company = this.$store.getters.company
      }

      if (!this.form.loginid) {
        this.$alert('이메일을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.loginid.focus()
        return
      }

      if (!this.isEdit) {
        if (!this.form.passwd) {
          this.$alert('비밀번호를 입력하세요', '', { confirmButtonText: '확인' })
          this.$refs.passwd.focus()
          return
        }

        if (!this.form.passwd2) {
          this.$alert('비밀번호를 입력하세요', '', { confirmButtonText: '확인' })
          this.$refs.passwd2.focus()
          return
        }

        if (this.form.passwd !== this.form.passwd2) {
          this.$alert('비밀번호를 정확하게 입력하세요', '', { confirmButtonText: '확인' })
          this.$refs.passwd.focus()
          return
        }
      }

      if (!this.form.name) {
        this.$alert('이름을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.name.focus()
        return
      }

      if (!this.form.hp) {
        this.$alert('핸드폰을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.hp.focus()
        return
      }

      if (this.isEdit) {
        const response = await request({
          method: 'PUT',
          url: '/api/user/' + this.form.id,
          data: this.form
        })

        if (response.code === 'CONFLICT') {
          this.$alert('이미 등록된 로그인ID입니다', '오류', { confirmButtonText: '확인' })
          this.$refs.loginid.focus()
          return
        }
      } else {
        const response = await request({
          method: 'POST',
          url: '/api/user',
          data: this.form
        })

        if (response.code === 'CONFLICT') {
          this.$alert('이미 등록된 로그인ID입니다', '오류', { confirmButtonText: '확인' })
          this.$refs.loginid.focus()
          return
        }
      }

      this.$router.push('/management/user')
    },
    clickCancel() {
      this.$router.go(-1)
    }
  }
}
</script>
