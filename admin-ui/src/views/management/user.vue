<template>
  <div class="app-container">
    <div class="filter-container">
      <el-select v-model="searchstatus" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in statuss"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-select v-model="searchlevel" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in levels"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-select v-if="$store.getters.isAdmin" v-model="searchcompany" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in companys"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-select v-model="searchtype" style="float:left;margin-right:10px;">
        <el-option label="이메일" value="loginid" />
        <el-option label="이름" value="name" />
      </el-select>

      <el-input v-model="search" placeholder="검색" style="width:300px;" class="filter-item" @keyup.enter.native="clickSearch" />

      <el-button class="filter-item" type="primary" style="margin-left:15px;" icon="el-icon-search" @click="clickSearch">
        검색
      </el-button>
      <el-button :loading="downloadLoading" style="float:right;" class="filter-item" type="success" icon="el-icon-circle-plus-outline" @click="$router.push('/management/user/insert')">
        등록
      </el-button>
    </div>

    <el-table v-loading="listLoading" :data="items" border fit highlight-current-row style="width: 100%" @header-click="clickHeader">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column v-if="$store.getters.isAdmin" min-width="150px" label="업체" sortable>
        <template slot-scope="{row}">
          <span>{{ getCompany(row.company) }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="300px" label="이메일" sortable>
        <template slot-scope="{row}">
          <router-link :to="'/management/user/update/'+row.id" class="link-type">
            <span>{{ row.loginid }}</span>
          </router-link>
        </template>
      </el-table-column>

      <el-table-column min-width="150px" label="이름" sortable>
        <template slot-scope="{row}">
          <router-link :to="'/management/user/update/'+row.id" class="link-type">
            <span>{{ row.name }}</span>
          </router-link>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="상태" width="100px">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ row.status == 1 ? '사용' : '사용 안함' }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="레벨" width="100px">
        <template slot-scope="{row}">
          <el-tag :type="row.level | levelFilter">
            {{ getLevel(row.level) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="기술등급" width="120px" sortable>
        <template slot-scope="{row}">
          {{ getGrade(row.grade) }}
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="등록일" sortable>
        <template slot-scope="{row}">
          <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="" width="200px">
        <template slot-scope="{row}">
          <router-link :to="'/management/user/update/'+row.id" style="margin-right:10px;">
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
  name: 'UserList',
  components: { Pagination },
  filters: {
    statusFilter(value) {
      if (value === 1) {
        return 'primary'
      } else {
        return 'danger'
      }
    },
    levelFilter(value) {
      const values = ['info', 'primary', 'success', 'warning', 'danger']

      return values[value]
    }
  },
  data() {
    return {
      orderby: '',
      items: null,
      total: 0,
      downloadLoading: false,
      listLoading: true,
      page: 1,
      pagesize: 10,
      search: '',
      searchstatus: 0,
      searchlevel: 0,
      searchcompany: 0,
      searchtype: 'loginid',
      statuss: [
        { name: '상태', id: 0 },
        { name: '사용', id: 1 },
        { name: '사용 안함', id: 2 }
      ],
      levels: [
        { name: '레벨', id: 0 },
        { name: '작업자', id: 1 },
        { name: '매니저', id: 2 },
        { name: '관리자', id: 3 },
        { name: '총관리자', id: 4 }
      ],
      grades: [
        { name: '기술등급', id: 0 },
        { name: '없음', id: 1 },
        { name: '건축초급기술자', id: 2 },
        { name: '건축중급기술자', id: 3 },
        { name: '건축고급기술자', id: 4 }
      ],
      companys: []
    }
  },
  created: async function() {
    const response = await request({
      method: 'GET',
      url: '/api/company'
    })

    this.companys = [{ id: 0, name: '업체' }, ...response.data]

    if (!this.$store.getters.isAdmin) {
      this.searchcompany = this.$store.getters.company
    }

    this.getList()
  },
  methods: {
    getLevel: function(value) {
      for (var i = 0; i < this.levels.length; i++) {
        if (value === this.levels[i].id) {
          return this.levels[i].name
        }
      }

      return ''
    },
    getGrade: function(value) {
      for (var i = 0; i < this.grades.length; i++) {
        if (value === this.grades[i].id) {
          return this.grades[i].name
        }
      }

      return ''
    },
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
    getList: async function() {
      this.listLoading = true

      const params = `page=${this.page - 1}&size=${this.pagesize}&orderby=${this.orderby}&status=${this.searchstatus}&level=${this.searchlevel}&company=${this.searchcompany}&${this.searchtype}=${encodeURIComponent(this.search)}`
      const response = await request({
        method: 'GET',
        url: '/api/user?' + params
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
          url: '/api/user/' + id
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제가 취소되었습니다'
        })
      })
    },
    clickHeader(column, event) {
      let orderby = ''
      if (column.label === '이메일') {
        orderby = 'loginid'
      } else if (column.label === '이름') {
        orderby = 'name'
      } else if (column.label === '기술등급') {
        orderby = 'grade'
      } else if (column.label === '등록일') {
        orderby = 'date'
      } else {
        return
      }

      if (column.order === 'ascending') {
      } else if (column.order === 'descending') {
        orderby += 'Desc'
      } else {
        orderby = ''
      }

      console.log(orderby)
      this.orderby = orderby

      this.getList()
    }
  }
}
</script>
