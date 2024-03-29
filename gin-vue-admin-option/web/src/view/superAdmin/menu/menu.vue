<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="addMenu('0')">新增根菜单</el-button>
      </div>

      <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" label="路由Name" show-overflow-tooltip min-width="160" prop="name" />
        <el-table-column align="left" label="路由Path" show-overflow-tooltip min-width="160" prop="path" />
        <el-table-column align="left" label="是否隐藏" min-width="100" prop="hidden">
          <template #default="scope">
            <span>{{ scope.row.hidden?"隐藏":"显示" }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="父节点" min-width="90" prop="parentId" />
        <el-table-column align="left" label="排序" min-width="70" prop="sort" />
        <el-table-column align="left" label="文件路径" min-width="360" prop="component" />
        <el-table-column align="left" label="展示名称" min-width="120" prop="authorityName">
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="图标" min-width="140" prop="authorityName">
          <template #default="scope">
            <div class="icon-column">
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" fixed="right" :lable="$t('general.operations')" width="300">
          <template #default="scope">
            <el-button
              size="mini"
              type="text"
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >添加子菜单</el-button>
            <el-button
              size="mini"
              type="text"
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >{{ $t('general.edit') }}</el-button>
            <el-button
              size="mini"
              type="text"
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >{{ $t('general.delete') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="handleClose" :title="dialogTitle">
      <warning-bar title="新增菜单，需要在角色管理内篇日志权限才可使用" />
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item label="路由Name" prop="path" style="width:30%">
          <el-input
            v-model="form.name"
            autocomplete="off"
            placeholder="唯一英文字符串"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item prop="path" style="width:30%">
          <template #label>
            <div style="display:inline-flex">
              路由Path
              <el-checkbox v-model="checkFlag" style="float:right;margin-left:20px;">添加参数</el-checkbox>
            </div>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="建议只在后方拼接参数"
          />
        </el-form-item>
        <el-form-item label="是否隐藏" style="width:30%">
          <el-select v-model="form.hidden" placeholder="是否在列表隐藏">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="父节点ID" style="width:30%">
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="文件路径" prop="component" style="width:60%">
          <el-input v-model="form.component" autocomplete="off" />
          <span style="font-size:12px;margin-right:12px;">如果菜单包含子菜单，请创建router-view二级路由页面或者</span><el-button size="mini" @click="form.component = 'view/routerHolder.vue'">点我设置</el-button>
        </el-form-item>
        <el-form-item label="展示名称" prop="meta.title" style="width:30%">
          <el-input v-model="form.meta.title" autocomplete="off" />
        </el-form-item>
        <el-form-item label="图标" prop="meta.icon" style="width:30%">
          <icon :meta="form.meta" style="width:100%" />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort" style="width:30%">
          <el-input v-model.number="form.sort" autocomplete="off" />
        </el-form-item>
        <el-form-item label="KeepAlive" prop="meta.keepAlive" style="width:30%">
          <el-select v-model="form.meta.keepAlive" style="width:100%" placeholder="是否keepAlive缓存页面">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
        <el-form-item label="CloseTab" prop="meta.closeTab" style="width:30%">
          <el-select v-model="form.meta.closeTab" style="width:100%" placeholder="是否自动关闭tab">
            <el-option :value="false" label="否" />
            <el-option :value="true" label="是" />
          </el-select>
        </el-form-item>
      </el-form>
      <div>
        <el-button
          size="small"
          type="primary"
          icon="edit"
          @click="addParameter(form)"
        >新增菜单参数</el-button>
        <el-table :data="form.parameters" style="width: 100%">
          <el-table-column align="left" prop="type" label="参数类型" width="180">
            <template #default="scope">
              <el-select v-model="scope.row.type" placeholder="请选择">
                <el-option key="query" value="query" label="query" />
                <el-option key="params" value="params" label="params" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="key" label="参数key" width="180">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.key" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left" prop="value" label="参数值">
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.value" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"
                  size="small"
                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >{{ $t('general.delete') }}</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
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
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成

import {
  updateBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from '@/api/menu'
import infoList from '@/mixins/infoList'
import icon from '@/view/superAdmin/menu/icon.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
export default {
  name: 'Menus',
  components: {
    icon,
    warningBar
  },
  mixins: [infoList],
  data() {
    return {
      checkFlag: false,
      listApi: getMenuList,
      dialogFormVisible: false,
      dialogTitle: '新增菜单',
      menuOption: [
        {
          ID: '0',
          title: '根菜单'
        }
      ],
      form: {
        ID: 0,
        path: '',
        name: '',
        hidden: '',
        parentId: '',
        component: '',
        meta: {
          title: '',
          icon: '',
          defaultMenu: false,
          closeTab: false,
          keepAlive: false
        },
        parameters: []
      },
      rules: {
        path: [{ required: true, message: '请输入菜单name', trigger: 'blur' }],
        component: [
          { required: true, message: '请输入文件路径', trigger: 'blur' }
        ],
        'meta.title': [
          { required: true, message: '请输入菜单展示名称', trigger: 'blur' }
        ]
      },
      isEdit: false,
      test: ''
    }
  },
  async created() {
    this.pageSize = 999
    await this.getTableData()
  },
  methods: {
    addParameter(form) {
      if (!form.parameters) {
        this.form.parameters = []
      }
      form.parameters.push({
        type: 'query',
        key: '',
        value: ''
      })
    },
    deleteParameter(parameters, index) {
      parameters.splice(index, 1)
    },
    changeName() {
      this.form.path = this.form.name
    },
    setOptions() {
      this.menuOption = [
        {
          ID: '0',
          title: '根目录'
        }
      ]
      this.setMenuOptions(this.tableData, this.menuOption, false)
    },
    setMenuOptions(menuData, optionsData, disabled) {
      menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === this.form.ID,
              children: []
            }
            this.setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === this.form.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === this.form.ID
            }
            optionsData.push(option)
          }
        })
    },
    handleClose(done) {
      this.initForm()
      done()
    },
    // 删除菜单
    deleteMenu(ID) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', this.$t('general.hint'), {
        confirmButtonText: this.$t('general.confirm'),
        cancelButtonText: this.$t('general.cancel'),
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteBaseMenu({ ID })
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            if (this.tableData.length === 1 && this.page > 1) {
              this.page--
            }
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    // 初始化弹窗内表格方法
    initForm() {
      this.checkFlag = false
      this.$refs.menuForm.resetFields()
      this.form = {
        ID: 0,
        path: '',
        name: '',
        hidden: '',
        parentId: '',
        component: '',
        meta: {
          title: '',
          icon: '',
          defaultMenu: false,
          keepAlive: ''
        }
      }
    },
    // 关闭弹窗
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    // 添加menu
    async enterDialog() {
      this.$refs.menuForm.validate(async valid => {
        if (valid) {
          let res
          if (this.isEdit) {
            res = await updateBaseMenu(this.form)
          } else {
            res = await addBaseMenu(this.form)
          }
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: this.isEdit ? '编辑成功' : '添加成功!'
            })
            this.getTableData()
          }
          this.initForm()
          this.dialogFormVisible = false
        }
      })
    },
    // 添加菜单方法，id为 0则为添加根菜单
    addMenu(id) {
      this.dialogTitle = '新增菜单'
      this.form.parentId = String(id)
      this.isEdit = false
      this.setOptions()
      this.dialogFormVisible = true
    },
    // 修改菜单方法
    async editMenu(id) {
      this.dialogTitle = '编辑菜单'
      const res = await getBaseMenuById({ id })
      this.form = res.data.menu
      this.isEdit = true
      this.setOptions()
      this.dialogFormVisible = true
    }
  }
}
</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
