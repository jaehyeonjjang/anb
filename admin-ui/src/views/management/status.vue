<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="searchtype" style="float:left;margin-right:10px;" @change="changeSearchtype">
        <el-option
          v-for="item in types"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>

      <el-select v-if="searchtype == 2 || searchtype == 9 || searchtype == 8" v-model="searchstatuscategory" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in statuscategorys"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-input v-model="search" placeholder="검색" style="width:300px;" class="filter-item" @keyup.enter.native="clickSearch" />

      <el-button class="filter-item" type="primary" style="margin-left:15px;" icon="el-icon-search" @click="clickSearch">
        검색
      </el-button>
      <el-button :loading="downloadLoading" style="float:right;" class="filter-item" type="success" icon="el-icon-circle-plus-outline" @click="$router.push('/management/status/insert')">
        등록
      </el-button>
    </div>

    <el-table v-loading="listLoading" :data="items" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="구분" width="150px">
        <template slot-scope="{row}">
          <el-tag type="primary">
            {{ getType(row.type) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="유형분류" width="150px">
        <template slot-scope="{row}">
          <el-tag v-if="row.statuscategory > 0" type="primary">
            {{ getStatuscategory(row.statuscategory) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column min-width="300px" label="유형명">
        <template slot-scope="{row}">
          <router-link :to="'/management/status/update/'+row.id" class="link-type">
            <span>{{ row.name }}</span>
          </router-link>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="정렬" width="150px">
        <template slot-scope="{row}">
          {{ row.order }}
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="등록일">
        <template slot-scope="{row}">
          <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="" width="200px">
        <template slot-scope="{row}">
          <router-link :to="'/management/status/update/'+row.id" style="margin-right:10px;">
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
  name: 'StatusList',
  components: { Pagination },
  data() {
    return {
      items: null,
      total: 0,
      downloadLoading: false,
      listLoading: true,
      page: 1,
      pagesize: 10,
      search: '',
      searchtype: 0,
      searchstatuscategory: 0,
      types: [
        { label: '구분', value: 0 },
        { label: '부위', value: 10 },
        { label: '부재 - 빨강', value: 1 },
        { label: '유형 - 빨강', value: 2 },
        { label: '폭', value: 3 },
        { label: '길이', value: 4 },
        { label: '개소', value: 5 },
        { label: '진행사항', value: 6 },
        { label: '비고', value: 7 },
        { label: '계단실 - 부재', value: 8 },
        { label: '계단실 - 부위', value: 9 },
        { label: '부재 - 파랑', value: 11 },
        { label: '유형 - 파랑', value: 12 },
        { label: '하자조사 - 부위', value: 21 },
        { label: '하자조사 - 부재', value: 22 },
        { label: '하자조사 - 하자', value: 23 }
      ],
      statuscategorys: [],
      allstatuscategorys: []
    }
  },
  created: async function() {
    this.getList()

    const params = `company=0`
    const response = await request({
      method: 'GET',
      url: '/api/statuscategory?' + params
    })

    this.allstatuscategorys = response.data
  },
  methods: {
    getType: function(value) {
      for (let i = 0; i < this.types.length; i++) {
        if (value === this.types[i].value) {
          return this.types[i].label
        }
      }

      return ''
    },
    getStatuscategory: function(value) {
      for (let i = 0; i < this.allstatuscategorys.length; i++) {
        if (value === this.allstatuscategorys[i].id) {
          return this.allstatuscategorys[i].name
        }
      }

      return ''
    },
    getList: async function() {
      this.listLoading = true

      const params = `page=${this.page - 1}&size=${this.pagesize}&company=${this.$store.getters.company}&type=${this.searchtype}&statuscategory=${this.searchstatuscategory}&name=${encodeURIComponent(this.search)}`
      const response = await request({
        method: 'GET',
        url: '/api/status?' + params
      })

      console.log(response)

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
          url: '/api/status/' + id
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제가 취소되었습니다'
        })
      })
    },
    changeSearchtype: async function() {
      const params = `type=${this.searchtype}`
      const response = await request({
        method: 'GET',
        url: '/api/statuscategory?' + params
      })

      this.searchstatuscategory = 0
      this.statuscategorys = [{ id: 0, name: '유형분류' }, ...response.data]
    }
  }
}
</script>
