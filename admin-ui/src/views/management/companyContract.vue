<template>
  <div class="createPost-container">
    <el-form ref="form" class="form-container">

      <div class="createPost-main-container">

        <el-form-item label-width="100px" label="계약기간">
          <el-input v-model="contractdate" type="text" style="width:200px;float:left;" disabled />

          <el-button v-loading="loading" style="margin-left:10px;float:left;" type="warning" icon="el-icon-date" @click="clickApply">
            기간연장 신청
          </el-button>
        </el-form-item>

        <el-table v-loading="listLoading" :data="items" border fit highlight-current-row style="margin-top:50px;width: 100%">
          <el-table-column align="center" label="ID" width="80">
            <template slot-scope="{row}">
              <span>{{ row.id }}</span>
            </template>
          </el-table-column>

          <el-table-column min-width="300px" label="시작일자">
            <template slot-scope="{row}">
              <span>{{ row.contractstartdate | moment('YYYY-MM-DD') }}</span>
            </template>
          </el-table-column>

          <el-table-column min-width="300px" label="종료일자">
            <template slot-scope="{row}">
              <span>{{ row.contractenddate | moment('YYYY-MM-DD') }}</span>
            </template>
          </el-table-column>

          <el-table-column class-name="status-col" label="상태" width="150px">
            <template slot-scope="{row}">
              <el-tag :type="row.status | statusFilter">
                {{ getStatus(row.status) }}
              </el-tag>
            </template>
          </el-table-column>

          <el-table-column width="180px" align="center" label="신청일">
            <template slot-scope="{row}">
              <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
            </template>
          </el-table-column>

        </el-table>

      </div>

    </el-form>
  </div>
</template>

<script>
import request from '@/utils/request'

export default {
  name: 'CompanyContract',
  filters: {
    statusFilter(status) {
      if (status === 1) {
        return 'info'
      } else if (status === 2) {
        return 'primary'
      } else {
        return 'danger'
      }
    }
  },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      loading: true,
      listLoading: true,
      contractdate: '',
      items: [],
      statuss: [
        { name: '상태', id: 0 },
        { name: '신청중', id: 1 },
        { name: '승인', id: 2 },
        { name: '거부', id: 3 }
      ]
    }
  },
  created() {
    const id = this.$store.getters.company
    this.fetchData(id)
    this.fetchList(id)

    this.loading = false
  },
  mounted() {
  },
  methods: {
    getStatus: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.statuss.length; i++) {
        if (value === this.statuss[i].id) {
          return this.statuss[i].name
        }
      }

      return ''
    },
    fetchList: async function(id) {
      this.listLoading = true

      const params = `company=${id}`
      const response = await request({
        method: 'GET',
        url: '/api/contract/search?' + params
      })

      this.items = response.data

      this.listLoading = false
    },
    fetchData: async function(id) {
      const response = await request({
        method: 'GET',
        url: '/api/company/' + id
      })

      const item = response.data

      if (item.contractstartdate !== null && item.contractenddate !== null) {
        this.contractdate = this.$moment(item.contractstartdate).format('YYYY-MM-DD') + ' ~ ' + this.$moment(item.contractenddate).format('YYYY-MM-DD')
      }
    },
    clickApply: async function() {
      const params = `company=${this.$store.getters.company}&status=1`
      const response = await request({
        method: 'GET',
        url: '/api/contract/search?' + params
      })

      if (response.data.length > 0) {
        this.$message({
          type: 'warning',
          message: '이미 신청중입니다'
        })

        return
      }

      this.$confirm('기간연장 신청하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        const contractstartdate = this.$moment().format()
        const contractenddate = this.$moment().add(6, 'months').format()

        const item = {
          company: this.$store.getters.company,
          status: 1,
          contractstartdate: contractstartdate,
          contractenddate: contractenddate
        }

        console.log(item)

        await request({
          method: 'POST',
          url: '/api/contract',
          data: item
        })

        this.$message({
          type: 'info',
          message: '기간연장 신청되었습니다'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '기간연장 신청이 취소되었습니다'
        })
      })
    }
  }
}
</script>
