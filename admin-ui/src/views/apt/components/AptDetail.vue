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

        <el-button v-loading="loading" type="primary" icon="el-icon-document-copy" style="margin-left:20px;" @click="clickCopy">
          복사
        </el-button>
      </sticky>

      <div class="createPost-main-container">

        <el-form-item v-if="isEdit" label-width="100px" label="ID">
          <el-input v-model="form.id" type="text" disabled />
        </el-form-item>

        <el-form-item label-width="100px" label="현장">
          <el-select ref="aptgroup" v-model="form.aptgroup" style="float:left;margin-right:10px;">
            <el-option
              v-for="item in aptgroups"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>

        <el-form-item label-width="100px" label="구분">
          <el-radio-group v-model="form.type">
            <el-radio-button label="1">정밀안전점검</el-radio-button>
            <el-radio-button label="2">정기안전점검</el-radio-button>
            <el-radio-button label="3">하자조사</el-radio-button>
            <el-radio-button label="4">장기수선</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label-width="100px" label="작업명">
          <el-input ref="name" v-model="form.name" type="text" placeholder="작업명을 입력하세요" />
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

        <el-form-item label-width="100px" label="진행상태">
          <el-radio-group v-model="form.status">
            <el-radio-button label="0">준비</el-radio-button>
            <el-radio-button label="1">착수</el-radio-button>
            <el-radio-button label="2">완료</el-radio-button>
            <el-radio-button label="3">중단</el-radio-button>
          </el-radio-group>
        </el-form-item>

        <el-form-item label-width="100px" label="시작일">
          <el-date-picker
            v-model="form.startdate"
            type="date"
            placeholder="시작일을 선택하세요"
          />
        </el-form-item>

        <el-form-item label-width="100px" label="종료일">
          <el-date-picker
            v-model="form.enddate"
            type="date"
            placeholder="종료일을 선택하세요"
          />
        </el-form-item>

        <el-form-item label-width="100px" label="책임기술자">
          <el-select ref="aptgroup" v-model="form.master" style="float:left;margin-right:10px;">
            <el-option
              v-for="item in masters"
              :key="item.key"
              :label="item.label"
              :value="item.key"
            />
          </el-select>
        </el-form-item>

        <el-form-item label-width="100px" label="참여기술자">
          <el-transfer
            v-model="submasters"
            :titles="['직원 목록', '참여기술자 목록']"
            :data="allsubmasters"
          />
        </el-form-item>

        <el-form-item label-width="100px" label="작업자">
          <el-transfer
            v-model="users"
            :titles="['직원 목록', '작업자 목록']"
            :data="allusers"
          />
        </el-form-item>

        <el-form-item label-width="100px" label="도면 구분">
          <el-select ref="aptgroup" v-model="form.summarytype" style="float:left;margin-right:10px;width:450px;">
            <el-option
              v-for="item in summarytypes"
              :key="item.key"
              :label="item.label"
              :value="item.key"
            />
          </el-select>
        </el-form-item>

        <el-form-item label-width="100px" label="도면">
          <i class="el-icon-plus" style="cursor:hand;cursor:pointer;font-size:20px;float:right;display:block;" @click="clickImageInsert(0)" />
          <div style="clear:both;" />

          <div v-for="(item, index) in images" :key="item.id" class="block" :style="{marginLeft: item.level * 40 + 'px'}">
            <input v-model="item.name" type="text" class="input">

            <el-dropdown style="float:right;margin-left:10px;" @command="(command) => clickImageCommand(command, index)">
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

            <i class="el-icon-arrow-down btn" @click="clickImageDown(index)" />
            <i class="el-icon-arrow-up btn" @click="clickImageUp(index)" />
            <i class="el-icon-arrow-right btn" @click="clickImageRight(index)" />
            <i class="el-icon-arrow-left btn" @click="clickImageLeft(index)" />
            <i class="el-icon-close btn" @click="clickImageDelete(index)" />
            <i class="el-icon-plus btn" @click="clickImageInsert(index+1)" />

            <el-upload
              class="plan-uploader"
              :action="uploadUrl"
              accept="image/jpeg,image/png"
              :show-file-list="false"
              :on-success="(res, file) => handleImageSuccess(res, file, index)"
              :before-upload="beforeImageUpload"
            >
              <img v-if="item.filename" :src="getImagePath(item.filename)" class="plan">
              <i v-else class="el-icon-picture-outline plan-uploader-icon" />
            </el-upload>

            <el-select v-model="item.floortype" style="float:right;margin-right:10px;">
              <el-option label=" " value="0" />
              <el-option label="지하" value="1" />
              <el-option label="지상" value="2" />
              <el-option label="옥상" value="3" />
            </el-select>

            <el-select v-model="item.type" style="float:right;margin-right:10px;">
              <el-option label=" " value="0" />
              <el-option label="결함도" value="1" />
              <el-option label="계단실" value="5" />
              <el-option label="계단실 항목" value="8" />
              <el-option label="부재" value="4" />
              <el-option label="기울기" value="2" />
              <el-option label="강도 탄산화" value="3" />
              <el-option label="길이/면적" value="6" />
            </el-select>

            <div class="clear" />
          </div>

        </el-form-item>
      </div>
    </el-form>

    <el-dialog :visible.sync="dialogVisible">
      <img width="100%" :src="dialogImageUrl" alt="">
    </el-dialog>

    <el-dialog title="하위작업 생성" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="Prefix" label-width="100px">
          <el-input v-model="prefix" autocomplete="off" />
        </el-form-item>
        <el-form-item label="" label-width="100px">
          <el-input v-model="start" autocomplete="off" style="width:100;" />
          ~
          <el-input v-model="end" autocomplete="off" style="width:100;" />
        </el-form-item>
        <el-form-item label="Postfix" label-width="100px">
          <el-input v-model="postfix" autocomplete="off" />
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">Cancel</el-button>
        <el-button type="primary" @click="clickChildren">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import Sticky from '@/components/Sticky'
import request from '@/utils/request'

const defaultForm = {
  id: undefined,
  name: '',
  company: 0,
  aptgroup: 0,
  status: 1,
  type: 1,
  master: 0,
  startdate: '',
  enddate: '',
  report: 1,
  report1: 1,
  report2: 1,
  report3: 1,
  report4: 1,
  report5: 1,
  report6: 1,
  summarytype: 1
}

export default {
  name: 'AptDetail',
  components: { Sticky },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
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
      form: Object.assign({}, defaultForm),
      loading: true,
      tempRoute: {},
      companys: [],
      aptgroups: [],
      imageUrl: '',
      images: [],
      uploadUrl: process.env.VUE_APP_BASE_API + '/api/upload'
    }
  },
  created: async function() {
    this.form.aptgroup = this.$route.params.id

    const response = await request({
      method: 'GET',
      url: '/api/company'
    })

    this.companys = [{ id: 0, name: '업체' }, ...response.data]

    const responseAptgroup = await request({
      method: 'GET',
      url: '/api/aptgroup?company=' + this.$store.getters.company
    })

    this.aptgroups = [{ id: 0, name: '현장' }, ...responseAptgroup.data]

    const responseAlluser = await request({
      method: 'GET',
      url: '/api/user?company=' + this.$store.getters.company
    })

    const allusers = []
    const users = []

    const allsubmasters = []
    const submasters = []

    for (let i = 0; i < responseAlluser.data.length; i++) {
      allsubmasters.push({
        key: responseAlluser.data[i].id,
        label: responseAlluser.data[i].name
      })

      if (responseAlluser.data[i].level === 3) {
        continue
      }

      allusers.push({
        key: responseAlluser.data[i].id,
        label: responseAlluser.data[i].name
      })
    }

    this.allusers = allusers
    this.allsubmasters = allsubmasters

    this.masters = [{ key: 0, label: '책임기술자' }, ...allsubmasters]

    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)

      const response = await request({
        method: 'GET',
        url: '/api/image?apt=' + id
      })

      this.images = response.data

      for (let i = 0; i < this.images.length; i++) {
        this.images[i].type = '' + this.images[i].type
        this.images[i].floortype = '' + this.images[i].floortype
      }

      const responseUser = await request({
        method: 'GET',
        url: '/api/aptuser?apt=' + id
      })

      for (let i = 0; i < responseUser.data.length; i++) {
        users.push(responseUser.data[i].user)
      }

      const responseSubmaster = await request({
        method: 'GET',
        url: '/api/aptsubmaster?apt=' + id
      })

      for (let i = 0; i < responseSubmaster.data.length; i++) {
        submasters.push(responseSubmaster.data[i].user)
      }
    }

    this.summarytypes = [{ key: 1, label: '단지 (여러동으로 구성된 형태)' }, { key: 2, label: '단일동' }, { key: 3, label: '단일동 위치 구분 (지붕, 지상, 지하 등으로 구분지어 놓은 형태)' }]

    this.users = users
    this.submasters = submasters

    this.tempRoute = Object.assign({}, this.$route)

    this.loading = false
  },
  mounted() {

  },
  methods: {
    getImagePath: function(filename) {
      return process.env.VUE_APP_BASE_API + '/api/download?filename=' + filename
    },
    fetchData: async function(id) {
      const response = await request({
        method: 'GET',
        url: '/api/apt/' + id
      })

      this.form = response.data
    },
    clickSave: async function() {
      this.form.aptgroup = parseInt(this.form.aptgroup)
      this.form.company = parseInt(this.form.company)
      this.form.master = parseInt(this.form.master)

      this.form.updateuser = this.$store.getters.id

      if (!this.$store.getters.isAdmin) {
        this.form.company = this.$store.getters.company
      }

      if (this.form.aptgroup === 0) {
        this.$alert('현장을 선택하세요', '', { confirmButtonText: '확인' })
        this.$refs.aptgroup.focus()
        return
      }

      if (!this.form.name) {
        this.$alert('작업명을 입력하세요', '', { confirmButtonText: '확인' })
        this.$refs.name.focus()
        return
      }

      this.$message({
        type: 'success',
        message: '저장중입니다. 잠시만 기다려주세요'
      })

      let apt

      if (this.isEdit) {
        const response = await request({
          method: 'PUT',
          url: '/api/apt/' + this.form.id,
          data: this.form
        })

        apt = response.data

        await request({
          method: 'DELETE',
          url: '/api/aptuser/apt/' + this.form.id
        })

        await request({
          method: 'DELETE',
          url: '/api/aptsubmaster/apt/' + this.form.id
        })
      } else {
        this.form.user = this.$store.getters.id

        const response = await request({
          method: 'POST',
          url: '/api/apt',
          data: this.form
        })

        apt = response.data
      }

      for (let i = 0; i < this.users.length; i++) {
        const item = {
          apt: apt.id,
          user: this.users[i],
          company: this.$store.getters.company,
          level: 2
        }

        await request({
          method: 'POST',
          url: '/api/aptuser',
          data: item
        })
      }

      for (let i = 0; i < this.submasters.length; i++) {
        const item = {
          apt: apt.id,
          user: this.submasters[i],
          company: this.$store.getters.company,
          level: 2
        }

        await request({
          method: 'POST',
          url: '/api/aptsubmaster',
          data: item
        })
      }

      const response = await request({
        method: 'GET',
        url: '/api/image?apt=' + apt.id
      })

      const oldImages = response.data

      for (let i = 0; i < this.images.length; i++) {
        for (let j = 0; j < oldImages.length; j++) {
          if (this.images[i].id === oldImages[j].id) {
            oldImages[j].find = true
          }
        }
      }

      for (let j = 0; j < oldImages.length; j++) {
        if (oldImages[j].find === undefined) {
          await request({
            method: 'DELETE',
            url: '/api/image/' + oldImages[j].id
          })
        }
      }

      for (let i = 0; i < this.images.length; i++) {
        const item = this.images[i]

        for (let j = i + 1; j < this.images.length; j++) {
          const item2 = this.images[j]

          if (item2.level <= item.level) {
            break
          }

          this.images[j].parent = item.id
        }
      }

      for (let i = 0; i < this.images.length; i++) {
        const item = this.images[i]

        item.apt = apt.id
        item.company = this.$store.getters.company
        item.order = i

        const oldId = item.id
        let response

        if (item.id < 0) {
          response = await request({
            method: 'POST',
            url: '/api/image',
            data: item
          })
        } else {
          response = await request({
            method: 'PUT',
            url: '/api/image/' + item.id,
            data: item
          })
        }

        const newId = response.data.id

        if (oldId < 0) {
          for (let j = 0; j < this.images.length; j++) {
            if (this.images[j].parent === oldId) {
              this.images[j].parent = newId
            }
          }
        }
      }

      // this.$router.push('/apt/apt')
      this.$router.go(-1)
    },
    clickCancel() {
      this.$router.go(-1)
    },
    getImage() {
      this.imageId--
      return {
        id: this.imageId,
        apt: 0,
        name: '',
        level: 0,
        parent: 0,
        last: 0,
        title: '',
        type: '0',
        floortype: '0',
        filename: '',
        order: 0,
        date: ''
      }
    },
    getImageEnd(pos) {
      const item = this.images[pos]
      let end = pos
      for (let i = pos + 1; i < this.images.length; i++) {
        if (this.images[i].level <= item.level) {
          break
        }

        end = i
      }

      return end
    },
    clickImageInsert(pos) {
      const image = this.getImage()

      if (pos > 0) {
        image.level = this.images[pos - 1].level
      }
      this.images.splice(pos, 0, image)

      this.$forceUpdate()
    },
    clickImageDelete(pos) {
      this.$confirm('삭제하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        const end = this.getImageEnd(pos)

        this.images.splice(pos, end - pos + 1)

        this.$forceUpdate()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제가 취소되었습니다'
        })

        return
      })
    },
    clickImageDown(pos) {
      /*
      const blocks = []

      let current = 0
      let end = 0
      let start = 0

      let j = 0
      while (end >= this.images.length) {
        end = this.getImageEnd(start)

        blocks.push({ start: start, end: end })

        if (pos >= start && pos <= end) {
          current = j
        }

        start = end + 1

        j++
      }

      console.log(blocks)

      if (current === blocks.length - 1) {
        return
      }

      const count = blocks[current].end - blocks[current].start + 1
      const nextCount = blocks[current + 1].end - blocks[current + 1].start + 1

      console.log(count)
      console.log(nextCount)

      const items = []

      for (var i = blocks[current].start; i <= blocks[current].end; i++) {
        items.push(this.images[i])
      }

      this.images.splice(blocks[current].start, count)
      this.images.splice(blocks[current].start + nextCount, 0, ...items)
      */

      if (pos >= this.images.length - 1) {
        return
      }

      const temp = this.images[pos]
      this.images[pos] = this.images[pos + 1]
      this.images[pos + 1] = temp

      this.$forceUpdate()
    },
    clickImageUp(pos) {
      /*
      const blocks = []

      let current = 0
      let end = 0
      let start = 0

      let j = 0
      while (end >= this.images.length) {
        end = this.getImageEnd(start)

        blocks.push({ start: start, end: end })

        if (pos >= start && pos <= end) {
          current = j
        }

        start = end + 1

        j++
      }

      console.log(blocks)

      if (current === 0) {
        return
      }

      const count = blocks[current].end - blocks[current].start + 1
      const prevCount = blocks[current - 1].end - blocks[current - 1].start + 1

      console.log(count)
      console.log(prevCount)

      const items = []

      for (var i = blocks[current].start; i <= blocks[current].end; i++) {
        items.push(this.images[i])
      }

      this.images.splice(blocks[current].start, count)
      this.images.splice(blocks[current].start - prevCount, 0, ...items)
      */

      if (pos === 0) {
        return
      }

      const temp = this.images[pos]
      this.images[pos] = this.images[pos - 1]
      this.images[pos - 1] = temp

      this.$forceUpdate()
    },
    clickImageLeft(pos) {
      const item = this.images[pos]

      if (item.level === 0) {
        return
      }

      const end = this.getImageEnd(pos)

      for (let i = pos; i <= end; i++) {
        this.images[i].level--
      }

      this.$forceUpdate()
    },
    clickImageRight(pos) {
      const item = this.images[pos]

      const max = 4

      if (pos === 0) {
        return
      }

      if (item.level >= max) {
        return
      }

      let end = pos
      for (let i = pos + 1; i < this.images.length; i++) {
        if (this.images[i].level <= item.level) {
          break
        }

        if (this.images[i].level >= max) {
          return
        }

        end = i
      }

      for (let i = pos; i <= end; i++) {
        this.images[i].level++
      }

      this.$forceUpdate()
    },
    handleImageSuccess(res, file, pos) {
      const filename = res.data
      this.images[pos].filename = filename

      this.imageUrl = URL.createObjectURL(file.raw)

      this.$forceUpdate()
    },
    beforeImageUpload(file) {
      const isImage = (file.type === 'image/jpeg' || file.type === 'image/png')

      if (!isImage) {
        this.$message.error('이미지 파일만 업로드 가능합니다 (jpg, png)')
      }

      return isImage
    },
    clickImageCommand(command, pos) {
      if (command === 'copy') {
        this.$confirm('복사하시겠습니까', '', {
          confirmButtonText: '확인',
          cancelButtonText: '취소',
          type: 'warning'
        }).then(async() => {
          const end = this.getImageEnd(pos)

          console.log(pos)
          console.log(end)
          const items = []

          for (var i = pos; i <= end; i++) {
            const item = JSON.parse(JSON.stringify(this.images[i]))
            this.imageId--
            item.id = this.imageId
            items.push(item)
          }

          console.log(items)

          this.images.splice(end + 1, 0, ...items)

          this.$forceUpdate()
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '복사가 취소되었습니다'
          })

          return
        })
      } else if (command === 'children') {
        this.dialogFormVisible = true
        this.dialogPos = pos
        this.prefix = ''
        this.postfix = ''
        this.start = ''
        this.end = ''
      } else if (command === 'imageView') {
        const image = this.images[pos]
        const filename = this.getImagePath(image.filename)
        this.dialogImageUrl = filename
        this.dialogVisible = true
      } else if (command === 'imageDelete') {
        this.images[pos].filename = ''
        this.$forceUpdate()
      }
    },
    clickChildren() {
      const item = this.images[this.dialogPos]

      const start = parseInt(this.start)
      const end = parseInt(this.end)

      const items = []
      for (var i = start; i <= end; i++) {
        const image = this.getImage()

        image.level = item.level + 1
        image.name = this.prefix + i + this.postfix
        items.push(image)
      }

      this.images.splice(this.dialogPos + 1, 0, ...items)

      this.dialogFormVisible = false
    },
    clickCopy: async function() {
      const id = this.$route.params && this.$route.params.id
      this.$confirm('해당 작업을 복사하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        await request({
          method: 'POST',
          url: '/api/apt/' + id + '/copy'
        })

        this.$message({
          type: 'success',
          message: '복사되었습니다'
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '복사가 취소되었습니다'
        })
      })
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
    width: 32px;
    height: 32px;
    display: block;
}
</style>
