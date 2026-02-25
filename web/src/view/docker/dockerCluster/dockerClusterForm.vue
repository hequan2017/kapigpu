
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="集群名字:" prop="clusterName">
    <el-input v-model="formData.clusterName" :clearable="true" placeholder="请输入集群名字" />
</el-form-item>
        <el-form-item label="CA证书:" prop="caCert">
    <el-input v-model="formData.caCert" type="textarea" :rows="4" :clearable="true" placeholder="请输入CA证书" />
</el-form-item>
        <el-form-item label="客户端证书:" prop="clientCert">
    <el-input v-model="formData.clientCert" type="textarea" :rows="4" :clearable="true" placeholder="请输入客户端证书" />
</el-form-item>
        <el-form-item label="客户端私钥:" prop="clientKey">
    <el-input v-model="formData.clientKey" type="textarea" :rows="4" :clearable="true" placeholder="请输入客户端私钥" />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDockerCluster,
  updateDockerCluster,
  findDockerCluster
} from '@/api/docker/dockerCluster'

defineOptions({
    name: 'DockerClusterForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            clusterName: '',
            caCert: '',
            clientCert: '',
            clientKey: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               clusterName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDockerCluster({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createDockerCluster(formData.value)
               break
             case 'update':
               res = await updateDockerCluster(formData.value)
               break
             default:
               res = await createDockerCluster(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
