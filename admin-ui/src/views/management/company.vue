<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="searchstatus" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in statuss"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

      <el-input v-model="search" placeholder="검색" style="width:300px;" class="filter-item" @keyup.enter.native="clickSearch" />

      <el-button class="filter-item" type="primary" style="margin-left:15px;" icon="el-icon-search" @click="clickSearch">
        검색
      </el-button>
      <el-button :loading="downloadLoading" style="float:right;" class="filter-item" type="success" icon="el-icon-circle-plus-outline" @click="$router.push('/management/company/insert')">
        등록
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
          <router-link :to="'/management/company/update/'+row.id" class="link-type">
            <span>{{ row.name }}</span>
          </router-link>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="상태" width="150px">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ row.status == 1 ? '사용' : '사용 안함' }}
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
          <router-link :to="'/management/company/update/'+row.id" style="margin-right:10px;">
            <el-button type="primary" size="small" icon="el-icon-edit">
              수정
            </el-button>
          </router-link>

          <el-button type="warning" size="small" icon="el-icon-delete" @click="clickDelete(row.id)">
            삭제
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
  name: 'CompanyList',
  components: { Pagination },
  filters: {
    statusFilter(status) {
      if (status === 1) {
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
      searchstatus: 0,
      statuss: [
        { label: '상태', value: 0 },
        { label: '사용', value: 1 },
        { label: '사용 안함', value: 2 }
      ]
    }
  },
  created: async function() {
    this.getList()
  },
  methods: {
    getList: async function() {
      this.listLoading = true

      const params = `page=${this.page - 1}&size=${this.pagesize}&status=${this.searchstatus}&name=${encodeURIComponent(this.search)}`
      console.log(params)
      const response = await request({
        method: 'GET',
        url: '/api/company?' + params
      })

      this.items = response.data.content
      this.total = response.data.totalElements
      this.pagesize = response.data.pageable.pageSize

      this.listLoading = false
    },
    clickSearch: function() {
      this.getList()
    },
    clickDelete: async function(id) {
      this.$confirm('삭제하시겠습니까', '', {
        confirmButtonText: '확인',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(async() => {
        this.$message({
          type: 'success',
          message: '삭제되었습니다'
        })

        await request({
          method: 'DELETE',
          url: '/api/company/' + id
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제가 취소되었습니다'
        })
      })
    }
  }
}
</script>
