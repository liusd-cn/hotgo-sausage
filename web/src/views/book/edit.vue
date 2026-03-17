<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑书籍 #' + formValue.id : '添加书籍'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid
              cols="1 s:1 m:1 l:1 xl:1 2xl:1"
              responsive="screen"
            >
              <n-gi span="1">
                <n-form-item label="name" path="name">
                  <n-input
                    placeholder="请输入name"
                    v-model:value="formValue.name"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="type" path="type">
                  <n-input-number
                    placeholder="请输入type"
                    v-model:value="formValue.type"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="sale_time" path="saleTime">
                  <DatePicker v-model:formValue="formValue.saleTime" type="datetime" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" :disabled="!isFormValid" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>
<script lang="ts" setup>
  import { Edit, View } from '@/api/book';
  import { State, newState } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });
  const isFormValid = ref(true);

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    
    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      
      return;
    }

    // 编辑
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>