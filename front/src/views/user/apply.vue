<script>
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { onBeforeUnmount, ref, shallowRef, onMounted, reactive } from 'vue';
import { Editor, Toolbar } from '@wangeditor/editor-for-vue';
import { CreateHeritageInheritor, CreateHeritageProject, GetCateGory, GetLevel, uploadFile } from "@/api/heritage";
import { ElMessage, ElMessageBox } from "element-plus";
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import SectionMain from "@/components/SectionMain.vue";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import {mdiKeyboardReturn, mdiPageNext} from "@mdi/js";
import {CityList} from "@/api/auth";

export default {
  methods: {
    mdiKeyboardReturn() {
      return mdiKeyboardReturn
    },
    mdiPageNext() {
      return mdiPageNext
    }
  },
  components: {BaseButtons, BaseButton, SectionMain, LayoutAuthenticated, Editor, Toolbar},
  setup() {
    const editorRef = shallowRef();
    const showDialog = ref(false);
    const valueHtml = ref('');
    const heritageInheritor = reactive({
      inheritor: "",
      inheritorImg: "",
      project: "",
      cateGory: null,
      level: null,
      details: '',
      locate: ''
    });

    const heritageProject = reactive({
      name: '',
      category: null,
      level: null,
      locate: '',
      details: ''
    });
    const cateGoryList = reactive([]);
    const levelList = reactive([]);
    const toolbarConfig = {};
    const details = ref('');
    const editorConfig = {MENU_CONF: {}};
    editorConfig.MENU_CONF['uploadImage'] = {
      server: '/api/upload',
      timeout: 5 * 1000,
      fieldName: 'file',
      headers: {Accept: 'multipart/form-data'},
      onFailed: (file, result) => {
        const imageUrl = result.data[0];
        details.value = `<img src="${imageUrl}" style="width: 200px;height: 200px;display: block;"/>\n`;
      },
    };

    onBeforeUnmount(() => {
      const editor = editorRef.value;
      if (editor == null) return;
      editor.destroy();
    });

    onMounted(() => {
      GetCateGory().then((res) => {
        const data = res.data.data[0];
        Object.keys(data).forEach((key) => {
          const f = {
            label: key,
            value: data[key]
          };
          cateGoryList.push(f);
        });
      });
      GetLevel().then(res => {
        const data = res.data.data[0];
        Object.keys(data).forEach((key) => {
          const f = {
            label: key,
            value: data[key]
          };
          levelList.push(f);
        });
      });
    });

    const handleCreated = (editor) => {
      editorRef.value = editor;
    };

    const showHeritageInheritorForm = ref(false);
    const showHeritageProjectForm = ref(false);
    const saveApply = () => {
      ElMessageBox.confirm("选择你要申请的类型", 'Title', {
        confirmButtonText: '非遗人项目申请',
        cancelButtonText: '非遗项目申请',
        type: "success",
        beforeClose: (action, instance, done) => {
          if (action === 'confirm') {
            showHeritageInheritorForm.value = true;
            showHeritageProjectForm.value = false;
            heritageInheritor.details = details.value;
          } else if (action === 'cancel') {
            showHeritageInheritorForm.value = false;
            showHeritageProjectForm.value = true;
            heritageProject.details = details.value;
          }
          showDialog.value = true;
          done();
        }
      }).then(() => {
        // 这里可以放置一些在对话框关闭后执行的代码
      }).catch(() => {
      });
    };

    const options = CityList();
    const createHeritageInheritor = () => {
      if (heritageInheritor.inheritor === null || heritageInheritor.project === null ||
      heritageInheritor.cateGory === null || heritageInheritor.level === null ||
        heritageInheritor.details === null || heritageInheritor.locate === null) {
        ElMessage.warning("请填写完整表单")
        return
      }
      if (heritageInheritor.locate.length === 2) {
        heritageInheritor.locate = heritageInheritor.locate[0] + '-' + heritageInheritor.locate[1];
      } else {
        heritageInheritor.locate = heritageInheritor.locate[0];
      }
      CreateHeritageInheritor(heritageInheritor).then(res => {
        ElMessage.success(res.data.msg);
        showDialog.value = false;
        heritageInheritor.inheritor = "";
        heritageInheritor.inheritorImg = "";
        heritageInheritor.project = "";
        heritageInheritor.cateGory = null;
        heritageInheritor.level = null;
        heritageInheritor.details = '';
        heritageInheritor.locate = '';
      }).catch(err => {
        ElMessage.error(err);
      });
    };

    const createHeritageProject = () => {
      if (heritageProject.name === null || heritageProject.category === null || heritageProject.level === null || heritageProject.details === null || heritageProject.locate === null) {
        ElMessage.warning("请填写完整表单")
        return
      }
      if (heritageProject.locate.length === 2) {
        heritageProject.locate = heritageProject.locate[0] + '-' + heritageProject.locate[1];
      } else {
        heritageProject.locate = heritageProject.locate[0];
      }
      CreateHeritageProject(heritageProject).then(res => {
        ElMessage.success(res.data.msg);
        showDialog.value = false;
        heritageProject.name = "";
        heritageProject.category = null;
        heritageProject.level = null;
        heritageProject.details = '';
        heritageProject.locate = '';
      }).catch(err => {
        ElMessage.error(err);
      });
    };

    const handleMainPhoto = (file) => {
      uploadFile(file.raw).then((res) => {
        heritageInheritor.inheritorImg = res.data.data[0];
        ElMessage.success("上传成功");
      }).catch(err => {
        ElMessage.error(err);
      });
    };

    return {
      editorRef,
      valueHtml,
      mode: 'default',
      toolbarConfig,
      editorConfig,
      heritageInheritor,
      handleCreated,
      showDialog,
      saveApply,
      cateGoryList,
      levelList,
      options,
      createHeritageInheritor,
      handleMainPhoto,
      heritageProject,
      createHeritageProject,
      showHeritageProjectForm,
      showHeritageInheritorForm,
      details,
    };
  },
  watch: {
    'heritageInheritor.locate'(newValue) {
      console.log('Selected value changed to:', JSON.stringify(newValue));
    },
  },
};
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <div class="edit" style="border: 1px solid #ccc; border-radius: 20px; overflow: hidden;">
        <Toolbar
          style="border-bottom: 1px solid #ccc;"
          :editor="editorRef"
          :defaultConfig="toolbarConfig"
          :mode="mode"
        />
        <Editor
          style="height: 500px; overflow-y: hidden; border-bottom-left-radius: 20px; border-bottom-right-radius: 20px;"
          v-model="details"
          :defaultConfig="editorConfig"
          :mode="mode"
          @onCreated="handleCreated"
        />
      </div>
      <BaseButton
        label="Button"
        color="success"
        small
        @click="saveApply"
      />
      <el-dialog
        v-model="showDialog"
        title="提交申请"
        width="600"
        :lock-scroll="false"
        style="border-radius: 20px;"
      >
        <el-descriptions border column="1" v-if="showHeritageInheritorForm">
          <el-descriptions-item label="申请人姓名">
            <el-input v-model="heritageInheritor.inheritor" style="width: 200px;"
                      placeholder="请输入申请人姓名">

            </el-input>
          </el-descriptions-item>
          <el-descriptions-item label="申请人照片">
            <el-upload
              class="upload-demo"
              drag
              style="width: 200px;"
              :limit="1"
              :on-change="handleMainPhoto"
              :auto-upload="false"
            >
              <div class="el-upload__text">添加证件</div>
            </el-upload>
          </el-descriptions-item>
          <el-descriptions-item label="申请项目">
            <el-input v-model="heritageInheritor.project" style="width: 300px;" placeholder="请输入申请项目"></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="项目类型">
            <el-select filterable v-model="heritageInheritor.cateGory" placeholder="选择项目类型" style="width: 240px;">
              <el-option
                v-for="item in cateGoryList"
                :key="item.label"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-descriptions-item>
          <el-descriptions-item label="申请级别">
            <el-select filterable v-model="heritageInheritor.level" placeholder="选择申请级别" style="width: 240px;">
              <el-option
                v-for="item in levelList"
                :key="item.label"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-descriptions-item>
          <el-descriptions-item label="申请地点">
            <el-cascader filterable placeholder="选择省市" :options="options"
                         :props="{checkStrictly: true, value: 'label', label: 'label' }" clearable
                         v-model="heritageInheritor.locate"/>
          </el-descriptions-item>
          <el-descriptions-item>
            <el-button @click="showDialog = false">取消</el-button>
            <el-button @click="createHeritageInheritor">确认</el-button>
          </el-descriptions-item>
        </el-descriptions>

        <el-descriptions v-if="showHeritageProjectForm" border column="1">
          <el-descriptions-item label="项目名称">
            <el-input v-model="heritageProject.name" style="width: 300px;" placeholder="请输入申请项目"></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="项目类型">
            <el-select filterable v-model="heritageProject.category" placeholder="选择项目类型" style="width: 240px;">
              <el-option
                v-for="item in cateGoryList"
                :key="item.label"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-descriptions-item>
          <el-descriptions-item label="申请级别">
            <el-select filterable v-model="heritageProject.level" placeholder="选择申请级别" style="width: 240px;">
              <el-option
                v-for="item in levelList"
                :key="item.label"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-descriptions-item>
          <el-descriptions-item label="申请地点">
            <el-cascader filterable placeholder="选择省市" :options="options"
                         :props="{checkStrictly: true, value: 'label', label: 'label' }" clearable
                         v-model="heritageProject.locate"/>
          </el-descriptions-item>
          <el-descriptions-item>
            <el-button @click="showDialog = false">取消</el-button>
            <el-button @click="createHeritageProject">确认</el-button>
          </el-descriptions-item>
        </el-descriptions>
      </el-dialog>
    </SectionMain>
  </LayoutAuthenticated>
</template>

<style scoped>
.top-container {
  height: 100vh;
  overflow: hidden;
  background-image: linear-gradient(to right bottom, #604281, #bdb287);

}

.main {
  padding: 20px;
}

.edit {
  margin-bottom: 20px;
}


.upload-demo {
  border: 2px dashed #d9d9d9;
  border-radius: 6px;
  background-color: #fafafa;
  text-align: center;
  padding: 20px;
  cursor: pointer;
  transition: border-color 0.3s ease;
}

.upload-demo:hover {
  border-color: #409eff;
}

</style>
