<template>
  <div>
    <warning-bar title="在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog">{{ $t('general.add') }}</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="接入日期" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="姓名" prop="customerName" width="120" />
        <el-table-column align="left" label="电话" prop="customerPhoneData" width="120" />
        <el-table-column align="left" label="接入人ID" prop="sysUserId" width="120" />
        <el-table-column align="left" label="按钮组" min-width="160">
          <template #default="scope">
            <el-button size="small" type="text" icon="edit" @click="updateCustomer(scope.row)">变更</el-button>
            <el-popover :visible="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">{{ $t('general.cancel') }}</el-button>
                <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">{{ $t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button type="text" icon="delete" size="mini">{{ $t('general.delete') }}</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="客户">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="客户名">
          <el-input v-model="form.customerName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="客户电话">
          <el-input v-model="form.customerPhoneData" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">{{ $t('general.close') }}</el-button>
          <el-button size="small" type="primary" @click="enterDialog">{{ $t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createExaCustomer,
  updateExaCustomer,
  deleteExaCustomer,
  getExaCustomer,
  getExaCustomerList
} from '@/api/customer'
import infoList from '@/mixins/infoList'
import warningBar from '@/components/warningBar/warningBar.vue'

export default {
  name: 'Customer',
  components: { warningBar },
  mixins: [infoList],
  data() {
    return {
      listApi: getExaCustomerList,
      dialogFormVisible: false,
      type: '',
      form: {
        customerName: '',
        customerPhoneData: ''
      }
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    async updateCustomer(row) {
      const res = await getExaCustomer({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.form = res.data.customer
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.form = {
        customerName: '',
        customerPhoneData: ''
      }
    },
    async deleteCustomer(row) {
      row.visible = false
      const res = await deleteExaCustomer({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: this.$t('general.deleteSuccess')
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createExaCustomer(this.form)
          break
        case 'update':
          res = await updateExaCustomer(this.form)
          break
        default:
          res = await createExaCustomer(this.form)
          break
      }

      if (res.code === 0) {
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  }
}
</script>

<style></style>
