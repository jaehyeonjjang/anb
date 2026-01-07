<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="searchcompany" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in companys"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-select v-model="searchstatus" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in statuss"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-input v-model="search" placeholder="검색" style="width:300px;" class="filter-item" @keyup.enter.native="clickSearch" />

      <el-button class="filter-item" type="primary" style="margin-left:15px;" icon="el-icon-search" @click="clickSearch">
        검색
      </el-button>

    </div>

    <el-table v-loading="listLoading" :data="items" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="300px" label="업체명">
        <template slot-scope="{row}">
          <span>{{ getCompany(row.company) }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="상태" width="150px">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ getStatus(row.status) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="등록일">
        <template slot-scope="{row}">
          <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="" width="200px">
        <template slot-scope="{row}">
          <el-button v-if="row.status === 1" type="primary" size="small" icon="el-icon-edit" @click="clickApply(row)">
            승인
          </el-button>

          <el-button v-if="row.status === 1" type="danger" size="small" icon="el-icon-delete" @click="clickReject(row)">
            거부
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="page" :limit.sync="pagesize" @pagination="getList" />
  </div>
</template>

<script>
import request from '@/utils/request'
import Pagination from '@/components/Pagination'

export default {
  name: 'ContractList',
  components: { Pagination },
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
  data() {
    return {
      items: null,
      total: 0,
      downloadLoading: false,
      listLoading: true,
      page: 1,
      pagesize: 10,
      search: '',
      searchcompany: 0,
      searchstatus: 0,
      companys: [],
      statuss: [
        { name: '상태', id: 0 },
        { name: '대기중', id: 1 },
        { name: '승인', id: 2 },
        { name: '거부', id: 3 }
      ]
    }
  },
  created: async function() {
    this.getList()
  },
  methods: {
    getCompany: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.companys.length; i++) {
        if (value === this.companys[i].id) {
          return this.companys[i].name
        }
      }

      return ''
    },
    getStatus: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.statuss.length; i++) {
        if (value === this.statuss[i].id) {
          return this.statuss[i].name
        }
      }
    },
    getList: async function() {
      this.listLoading = true

      const companyResponse = await request({
        method: 'GET',
        url: '/api/company'
      })

      this.companys = [{ id: 0, name: '업체' }, ...companyResponse.data]

      const params = `page=${this.page - 1}&size=${this.pagesize}&status=${this.searchstatus}&company=${this.searchcompany}`

      const response = await request({
        method: 'GET',
        url: '/api/contract?' + params
      })

      this.items = response.data.content
      this.total = response.data.totalElements
      this.pagesize = response.data.pageable.pageSize

      this.listLoading = false
    },
    clickSearch: function() {
      this.getList()
    },
    clickApply: async function(item) {
      this.$confirm('승인시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        this.$message({
          type: 'success',
          message: '승인되었습니다'
        })

        item.status = 2
        await request({
          method: 'PUT',
          url: '/api/contract/' + item.id,
          data: item
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '승인이 취소되었습니다'
        })
      })
    },
    clickReject: async function(item) {
      this.$confirm('거부하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        this.$message({
          type: 'success',
          message: '거부되었습니다'
        })

        item.status = 3
        await request({
          method: 'PUT',
          url: '/api/contract/' + item.id,
          data: item
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '거부가 취소되었습니다'
        })
      })
    }
  }
}
</script>
