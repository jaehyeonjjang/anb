<template>
  <div class="createPost-container">
    <el-form ref="form" :model="form" class="form-container">

      <sticky :z-index="10" :class-name="'sub-navbar '">
        <el-button v-loading="loading" style="margin-left: 10px;" type="success" icon="el-icon-check" @click="clickSave">
          저장
        </el-button>
      </sticky>

      <div class="createPost-main-container">

        <el-form-item v-if="isEdit" label-width="100px" label="ID">
          <el-input v-model="form.id" type="text" disabled />
        </el-form-item>

        <el-form-item label-width="100px" label="업체명">
          <el-input ref="name" v-model="form.name" type="text" placeholder="업체명을 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="대표자명">
          <el-input ref="ceo" v-model="form.ceo" type="text" placeholder="대표자명을 입력하세요" />
        </el-form-item>

        <el-form-item label-width="100px" label="로고">
          <el-upload
            class="plan-uploader"
            :action="uploadUrl"
            accept="image/jpeg,image/png"
            :show-file-list="false"
            :on-success="(res, file) => handleImageSuccess(res, file, 'logo')"
            :before-upload="beforeImageUpload"
          >
            <img v-if="form.logo" :src="getImagePath(form.logo)" class="plan">
            <i v-else class="el-icon-picture-outline plan-uploader-icon" />
          </el-upload>
        </el-form-item>

        <el-form-item label-width="100px" label="직인">
          <el-upload
            class="plan-uploader"
            :action="uploadUrl"
            accept="image/jpeg,image/png"
            :show-file-list="false"
            :on-success="(res, file) => handleImageSuccess(res, file, 'stamp')"
            :before-upload="beforeImageUpload"
          >
            <img v-if="form.stamp" :src="getImagePath(form.stamp)" class="plan">
            <i v-else class="el-icon-picture-outline plan-uploader-icon" />
          </el-upload>
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
  ceo: '',
  status: 1,
  logo: '',
  stamp: ''
}

export default {
  name: 'CompanyBasic',
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
      uploadUrl: process.env.VUE_APP_BASE_API + '/api/upload/management'
    }
  },
  created() {
    const id = this.$store.getters.company
    this.fetchData(id)

    this.tempRoute = Object.assign({}, this.$route)

    this.loading = false
  },
  mounted() {

  },
  methods: {
    getImagePath: function(filename) {
      return process.env.VUE_APP_BASE_API + '/api/download/management?filename=' + filename
    },
    fetchData: async function(id) {
      const response = await request({
        method: 'GET',
        url: '/api/company/' + id
      })

      this.form = response.data
    },
    clickSave: async function() {
      if (!this.form.name) {
        this.$alert('업체명을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.name.focus()
        return
      }

      await request({
        method: 'PUT',
        url: '/api/company/basic/' + this.form.id,
        data: this.form
      })

      this.$alert('저장되었습니다.', '', { confirmButtonText: '확인' })
    },
    clickCancel() {
      this.$router.go(-1)
    },
    handleImageSuccess(res, file, mode) {
      const filename = res.data

      if (mode === 'logo') {
        this.form.logo = filename
      } else {
        this.form.stamp = filename
      }

      this.$forceUpdate()
    },
    beforeImageUpload(file) {
      const isImage = (file.type === 'image/jpeg' || file.type === 'image/png')

      if (!isImage) {
        this.$message.error('이미지 파일만 업로드 가능합니다 (jpg, png)')
      }

      return isImage
    }
  }
}
</script>

<style lang="scss" scoped>

.block {
  border: 1px solid #aaa;
  border-radius: 3px;
  padding: 10px 10px;
  margin-bottom: 10px;

  .input {
    display:block;
    float:left;
    width:400px;
    padding:10px 10px;
    border:none;
    margin-right: 10px;
  }

  .btn {
    cursor: hand;
    cursor: pointer;
    display:block;
    float:right;
    font-size:20px;
    margin: 10px 10px;
  }

  .clear {
    clear:both;
  }
}

.plan-uploader {
  margin-top: 4px;
  margin-right: 20px;
  margin-bottom: 20px;
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
  font-size: 48px;
  color: #8c939d;
  width: 64px;
  height: 64px;
  line-height: 64px;
  text-align: center;
}
.plan {
  width: 64px;
  height: 64px;
  display: block;
}
</style>
