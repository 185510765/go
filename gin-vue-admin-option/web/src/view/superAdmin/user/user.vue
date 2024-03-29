<template>
  <div>
    <warning-bar :title="$t('authority.authorityNote')" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="addUser">{{ $t('user.addUser') }}</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column align="left" :label="$t('user.avatar')" min-width="50">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="UUID" min-width="250" prop="uuid" />
        <el-table-column align="left" :label="$t('user.userName')" min-width="150" prop="userName" />
        <el-table-column align="left" :label="$t('user.nickName')" min-width="100" prop="nickName">
          <template #default="scope">
            <p v-if="!scope.row.editFlag" class="nickName">{{ scope.row.nickName }}
              <el-icon class="pointer" color="#66b1ff" @click="openEidt(scope.row)">
                <edit />
              </el-icon>
            </p>
            <p v-if="scope.row.editFlag" class="nickName">
              <el-input v-model="scope.row.nickName" />
              <el-icon class="pointer" color="#67c23a" @click="enterEdit(scope.row)">
                <check />
              </el-icon>
              <el-icon class="pointer" color="#f23c3c" @click="closeEdit(scope.row)">
                <close />
              </el-icon>
            </p>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="$t('user.userRole')" min-width="150">
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeAuthority(scope.row,flag)}"
              @remove-tag="()=>{changeAuthority(scope.row,false)}"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" :lable="$t('general.operations')" min-width="150">
          <template #default="scope">
            <el-popover :visible="scope.row.visible" placement="top" width="160">
              <p>{{ $t('user.deleteUserConfrim') }}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">{{ $t('general.cancel') }}</el-button>
                <el-button type="primary" size="mini" @click="deleteUser(scope.row)">{{ $t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button type="text" icon="delete" size="mini">{{ $t('general.delete') }}</el-button>
              </template>
            </el-popover>
            <el-button type="text" icon="magic-stick" size="mini" @click="resetPassword(scope.row)">{{ $t('user.resetPassword') }}</el-button>
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
    <el-dialog v-model="addUserDialog" custom-class="user-dialog" :title="$t('user.addUser')">
      <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="90px">
        <el-form-item :label="$t('user.userName')" prop="username">
          <el-input v-model="userInfo.username" />
        </el-form-item>
        <el-form-item :label="$t('user.password')" prop="password">
          <el-input v-model="userInfo.password" />
        </el-form-item>
        <el-form-item :label="$t('user.nickName')" prop="nickName">
          <el-input v-model="userInfo.nickName" />
        </el-form-item>
        <el-form-item :label="$t('user.userRole')" prop="authorityId">
          <el-cascader
            v-model="userInfo.authorityIds"
            style="width:100%"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="" label-width="80px">
          <div style="display:inline-block" @click="openHeaderChange">
            <img v-if="userInfo.headerImg" class="header-img-box" :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg">
            <div v-else class="header-img-box">{{ $t('user.mediaLibrary') }}</div>
          </div>
        </el-form-item>

      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeAddUserDialog">{{ $t('general.close') }}</el-button>
          <el-button size="small" type="primary" @click="enterAddUserDialog">{{ $t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = import.meta.env.VITE_BASE_API
import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import infoList from '@/mixins/infoList'
import { mapGetters } from 'vuex'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'
export default {
  name: 'Api',
  components: { CustomPic, ChooseImg, warningBar },
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      backNickName: '',
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: '',
        authorityIds: []
      },
      rules: {
        username: [
          { required: true, message: this.$t('user.userNameNote'), trigger: 'blur' },
          { min: 5, message: this.$t('user.userNameLenNote'), trigger: 'blur' }
        ],
        password: [
          { required: true, message: this.$t('user.passwordNote'), trigger: 'blur' },
          { min: 6, message: this.$t('user.passwordLenNote'), trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: this.$t('user.nickNameNote'), trigger: 'blur' }
        ],
        authorityId: [
          { required: true, message: this.$t('user.userRoleNote'), trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  watch: {
    tableData() {
      this.setAuthorityIds()
    }
  },
  async created() {
    await this.getTableData()
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    this.setOptions(res.data.list)
  },
  methods: {
    resetPassword(row) {
      this.$confirm(
        this.$t('user.resetPasswordConfrim'),
        this.$t('general.warning'),
        {
          confirmButtonText: this.$t('general.confirm'),
          cancelButtonText: this.$t('general.cancel'),
          type: 'warning',
        }
      ).then(async() => {
        const res = await resetPassword({
          ID: row.ID,
        })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: res.msg,
          })
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          })
        }
      })
    },
    setAuthorityIds() {
      this.tableData && this.tableData.forEach((user) => {
        const authorityIds = user.authorities && user.authorities.map(i => {
          return i.authorityId
        })
        user.authorityIds = authorityIds
      })
    },
    openHeaderChange() {
      this.$refs.chooseImg.open()
    },
    setOptions(authData) {
      this.authOptions = []
      this.setAuthorityOptions(authData, this.authOptions)
    },
    openEidt(row) {
      if (this.tableData.some(item => item.editFlag)) {
        this.$message(this.$t('user.anotherUserEdit'))
        return
      }
      this.backNickName = row.nickName
      row.editFlag = true
    },
    async enterEdit(row) {
      const res = await setUserInfo({ nickName: row.nickName, ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: this.$t('user.setUserInfoNote')
        })
      }
      this.backNickName = ''
      row.editFlag = false
    },
    closeEdit(row) {
      row.nickName = this.backNickName
      this.backNickName = ''
      row.editFlag = false
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            this.setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.ID })
      if (res.code === 0) {
        this.$message.success(this.$t('general.deleteSuccess'))
        await this.getTableData()
        row.visible = false
      }
    },
    async enterAddUserDialog() {
      this.userInfo.authorityId = this.userInfo.authorityIds[0]
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo)
          if (res.code === 0) {
            this.$message({ type: 'success', message: this.$t('user.userAddedNote') })
          }
          await this.getTableData()
          this.closeAddUserDialog()
        }
      })
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields()
      this.userInfo.headerImg = ''
      this.userInfo.authorityIds = []
      this.addUserDialog = false
    },
    addUser() {
      this.addUserDialog = true
    },
    async changeAuthority(row, flag) {
      if (flag) {
        return
      }
      this.$nextTick(async() => {
        const res = await setUserAuthorities({
          ID: row.ID,
          authorityIds: row.authorityIds
        })
        if (res.code === 0) {
          this.$message({ type: 'success', message: this.$t('user.roleSetNote') })
        }
      })
    },
  }
}
</script>

<style lang="scss">
.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
.nickName{
  display: flex;
  justify-content: flex-start;
  align-items: center;
}
.pointer{
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
