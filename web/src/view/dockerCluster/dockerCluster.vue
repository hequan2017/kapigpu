<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="集群名称">
          <el-input v-model="searchInfo.clusterName" placeholder="集群名称" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="searchInfo.remark" placeholder="备注" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="集群名称" prop="clusterName" min-width="150" />
        <el-table-column align="left" label="备注" prop="remark" min-width="200" />
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="200">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateDockerClusterFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link icon="View" size="small" @click="viewCredentials(scope.row)">查看凭证</el-button>
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

    <!-- 新增/编辑对话框 -->
    <el-dialog v-model="dialogFormVisible" :title="type === 'create' ? '新增Docker集群' : '编辑Docker集群'" destroy-on-close width="700px">
      <el-form ref="dialogForm" :model="formData" :rules="rules" label-width="120px">
        <el-form-item label="集群名称" prop="clusterName">
          <el-input v-model="formData.clusterName" placeholder="请输入集群名称" clearable :style="{ width: '100%' }" />
        </el-form-item>
        <el-form-item label="CA证书" prop="caCert">
          <el-input
            v-model="formData.caCert"
            type="textarea"
            placeholder="请输入CA证书内容（PEM格式）"
            :rows="6"
            clearable
            :style="{ width: '100%' }"
          />
          <div class="form-tip">请输入CA证书的PEM格式内容，将被加密存储</div>
        </el-form-item>
        <el-form-item label="客户端证书" prop="clientCert">
          <el-input
            v-model="formData.clientCert"
            type="textarea"
            placeholder="请输入客户端证书内容（PEM格式）"
            :rows="6"
            clearable
            :style="{ width: '100%' }"
          />
          <div class="form-tip">请输入客户端证书的PEM格式内容，将被加密存储</div>
        </el-form-item>
        <el-form-item label="客户端私钥" prop="clientKey">
          <el-input
            v-model="formData.clientKey"
            type="textarea"
            placeholder="请输入客户端私钥内容（PEM格式）"
            :rows="6"
            clearable
            :style="{ width: '100%' }"
          />
          <div class="form-tip">请输入客户端私钥的PEM格式内容，将被加密存储</div>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" placeholder="请输入备注" clearable :style="{ width: '100%' }" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 查看凭证对话框 -->
    <el-dialog v-model="credentialDialogVisible" title="集群凭证" destroy-on-close width="700px">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="集群名称">{{ credentialData.clusterName }}</el-descriptions-item>
        <el-descriptions-item label="CA证书">
          <el-input v-model="credentialData.caCert" type="textarea" :rows="6" readonly />
          <el-button type="primary" link size="small" @click="copyToClipboard(credentialData.caCert)">复制</el-button>
        </el-descriptions-item>
        <el-descriptions-item label="客户端证书">
          <el-input v-model="credentialData.clientCert" type="textarea" :rows="6" readonly />
          <el-button type="primary" link size="small" @click="copyToClipboard(credentialData.clientCert)">复制</el-button>
        </el-descriptions-item>
        <el-descriptions-item label="客户端私钥">
          <el-input v-model="credentialData.clientKey" type="textarea" :rows="6" readonly />
          <el-button type="primary" link size="small" @click="copyToClipboard(credentialData.clientKey)">复制</el-button>
        </el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="credentialDialogVisible = false">关 闭</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  createDockerCluster,
  deleteDockerCluster,
  deleteDockerClusterByIds,
  updateDockerCluster,
  findDockerCluster,
  getDockerClusterList,
  getDockerClusterCredentials
} from '@/api/dockerCluster'

// 搜索相关
const searchInfo = reactive({
  clusterName: '',
  remark: ''
})

const onReset = () => {
  searchInfo.clusterName = ''
  searchInfo.remark = ''
}

// 分页相关
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchForm = ref(null)

// 查询
const getTableData = async () => {
  const table = await getDockerClusterList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 提交搜索
const onSubmit = () => {
  page.value = 1
  getTableData()
}

// 分页
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 多选
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const deleteVisible = ref(false)
const onDelete = async () => {
  const ids = multipleSelection.value.map((item) => item.ID)
  const res = await deleteDockerClusterByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteDockerCluster(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 弹窗相关
const dialogFormVisible = ref(false)
const type = ref('')
const dialogForm = ref(null)

const formData = reactive({
  clusterName: '',
  caCert: '',
  clientCert: '',
  clientKey: '',
  remark: ''
})

const rules = reactive({
  clusterName: [{ required: true, message: '请输入集群名称', trigger: 'blur' }],
  caCert: [{ required: true, message: '请输入CA证书', trigger: 'blur' }],
  clientCert: [{ required: true, message: '请输入客户端证书', trigger: 'blur' }],
  clientKey: [{ required: true, message: '请输入客户端私钥', trigger: 'blur' }]
})

const initForm = () => {
  dialogForm.value?.resetFields()
  formData.clusterName = ''
  formData.caCert = ''
  formData.clientCert = ''
  formData.clientKey = ''
  formData.remark = ''
}

const openDialog = () => {
  type.value = 'create'
  initForm()
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  initForm()
}

const updateDockerClusterFunc = async (row) => {
  const res = await findDockerCluster({ ID: row.ID })
  if (res.code === 0) {
    type.value = 'update'
    Object.assign(formData, res.data)
    dialogFormVisible.value = true
  }
}

const enterDialog = async () => {
  dialogForm.value.validate(async (valid) => {
    if (valid) {
      let res
      switch (type.value) {
        case 'create':
          res = await createDockerCluster(formData)
          break
        case 'update':
          res = await updateDockerCluster(formData)
          break
        default:
          res = await createDockerCluster(formData)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: type.value === 'create' ? '创建成功' : '修改成功'
        })
        closeDialog()
        getTableData()
      }
    }
  })
}

// 凭证查看
const credentialDialogVisible = ref(false)
const credentialData = reactive({
  clusterName: '',
  caCert: '',
  clientCert: '',
  clientKey: ''
})

const viewCredentials = async (row) => {
  const res = await getDockerClusterCredentials({ id: row.ID })
  if (res.code === 0) {
    Object.assign(credentialData, res.data)
    credentialDialogVisible.value = true
  }
}

// 复制到剪贴板
const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    ElMessage({
      type: 'success',
      message: '复制成功'
    })
  }).catch(() => {
    ElMessage({
      type: 'error',
      message: '复制失败'
    })
  })
}

// 格式化日期
const formatDate = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleString()
}
</script>

<style scoped>
.form-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}
</style>
