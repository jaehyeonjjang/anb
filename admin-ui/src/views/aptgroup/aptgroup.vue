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

      <el-select v-if="$store.getters.isAdmin" v-model="searchcompany" style="float:left;margin-right:10px;">
        <el-option
          v-for="item in companys"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        />
      </el-select>

      <el-input v-model="search" placeholder="검색" style="width:300px;" class="filter-item" @keyup.enter.native="clickSearch" />

      <el-button class="filter-item" type="primary" style="margin-left:15px;" icon="el-icon-search" @click="clickSearch">
        검색
      </el-button>
      <el-button :loading="downloadLoading" style="float:right;" class="filter-item" type="success" icon="el-icon-circle-plus-outline" @click="$router.push('/aptgroup/aptgroup/insert')">
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

      <el-table-column min-width="150px" label="현장명" sortable>
        <template slot-scope="{row}">
          <span @click="clickAptlist(row.id)">{{ row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" width="100px" label="시설물 구분">
        <template slot-scope="{row}">
          <span>{{ getFacility(row.facility) }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" width="100px" label="종류">
        <template slot-scope="{row}">
          <span>{{ getType(row.type) }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" width="100px" label="작업현황">
        <template slot-scope="{row}">
          <router-link :to="{name: 'AptList', params: {aptgroup: row.id}}" class="link-type">
            <span>{{ row.aptcount }}건</span>
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

      <el-table-column width="80px" align="center" label="작성자">
        <template slot-scope="{row}">
          <span>{{ getUser(row.user) }}</span>
        </template>
      </el-table-column>

      <el-table-column width="80px" align="center" label="수정자">
        <template slot-scope="{row}">
          <span>{{ getUser(row.user) }}</span>
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="등록일" sortable>
        <template slot-scope="{row}">
          <span>{{ row.date | moment('YYYY-MM-DD HH:mm') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="" width="200px">
        <template slot-scope="{row}">
          <router-link :to="'/aptgroup/aptgroup/update/'+row.id" style="margin-right:10px;">
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

    <el-dialog
      title="작업 선택"
      :visible.sync="dialogVisible"
      width="30%">
      <span>

        <div v-for="item in apts"
             :key="item.id"
             :label="item.name"
             :value="item.id" style="margin-bottom:10px;">
          <el-button @click="$router.push('/apt/apt/update/' + item.id)" type="success" plain icon="el-icon-edit" style="width:100%;text-align:left;">
          {{item.name}}
        </el-button>
        </div>

        </span>
      <span slot="footer" class="dialog-footer">
        <el-button type="success" icon="el-icon-circle-plus-outline" @click="clickAptInsert">작업 등록</el-button>
        <el-button type="primary" icon="el-icon-close" @click="dialogVisible = false">Close</el-button>
      </span>
    </el-dialog>

  </div>


</template>

<script>
import request from '@/utils/request'
import Pagination from '@/components/Pagination'

export default {
  name: 'AptgroupList',
  components: { Pagination },
  filters: {
    statusFilter(value) {
      if (value === 1) {
        return 'primary'
      } else {
        return 'danger'
      }
    }
  },
  data() {
    return {
      orderby: 'dateDesc',
      dialogVisible: false,
      items: null,
      total: 0,
      downloadLoading: false,
      listLoading: true,
      page: 1,
      pagesize: 10,
      search: '',
      searchstatus: 0,
      searchcompany: 0,
      searchtype: 'loginid',
      statuss: [
        { name: '상태', id: 0 },
        { name: '사용', id: 1 },
        { name: '사용 안함', id: 2 }
      ],
      facilitys: [
        { name: '시설물 구분', id: 0 },
        { name: '1종', id: 1 },
        { name: '2종', id: 2 },
        { name: '3종', id: 3 }
      ],
      types: [
        { name: '종류', id: 0 },
        { name: '대형건축물', id: 1 },
        { name: '공동주택', id: 2 },
        { name: '단독주택', id: 3 },
        { name: '주상복합', id: 4 },
        { name: '업무시설', id: 5 },
        { name: '다중이용건축물', id: 6 }
      ],
      companys: [],
      users: [],
      apts: [],
      selectedId: 0
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

    const responseUser = await request({
      method: 'GET',
      url: '/api/user?company=' + this.searchcompany
    })

    this.users = responseUser.data

    this.getList()
  },
  methods: {
    getUser: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.users.length; i++) {
        if (value === this.users[i].id) {
          return this.users[i].name
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
    getFacility: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.facilitys.length; i++) {
        if (value === this.facilitys[i].id) {
          return this.facilitys[i].name
        }
      }

      return ''
    },
    getType: function(value) {
      if (value === 0) {
        return ''
      }

      for (var i = 0; i < this.types.length; i++) {
        if (value === this.types[i].id) {
          return this.types[i].name
        }
      }

      return ''
    },
    getList: async function() {
      this.listLoading = true

      const params = `page=${this.page - 1}&size=${this.pagesize}&orderby=${this.orderby}&status=${this.searchstatus}&company=${this.searchcompany}&name=${encodeURIComponent(this.search)}`
      console.log(params)
      const response = await request({
        method: 'GET',
        url: '/api/aptgroup?' + params
      })

      for (let i = 0; i < response.data.content.length; i++) {
        const item = response.data.content[i]
        const params = `aptgroup=${item.id}`
        const itemResponse = await request({
          method: 'GET',
          url: '/api/apt?' + params
        })

        response.data.content[i].aptcount = itemResponse.data.length
      }

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
          url: '/api/aptgroup/' + id
        })

        this.getList()
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제가 취소되었습니다'
        })
      })
    },
    clickAptlist: async function(id) {
      console.log(id)

      const params = `aptgroup=${id}`
      const item = await request({
        method: 'GET',
        url: '/api/apt?' + params
      })

      this.selectedId = id
      this.apts = item.data
      this.dialogVisible = true
      console.log(item.data)
    },
    clickAptInsert: function() {
      console.log(this.selectedId)
      this.$router.push({name: 'AptInsert', params: {id: this.selectedId}})
    },
    clickSort(sortProps) {
      this.clickHeader(sortProps.column,event)
    },
    clickHeader(column, event) {
      let orderby = ''
      if (column.label === '현장명') {
        orderby = 'name'
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

      this.orderby = orderby      

      this.getList()
    }
  }
}
</script>
