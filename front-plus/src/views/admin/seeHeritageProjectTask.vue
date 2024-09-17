<script>
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import SectionMain from "@/components/SectionMain.vue";
import CardBox from "@/components/CardBox.vue";
import SectionTitleLineWithButton from "@/components/SectionTitleLineWithButton.vue";
import {mdiBallotOutline, mdiKeyboardReturn, mdiPageNext} from "@mdi/js";
import BaseLevel from "@/components/BaseLevel.vue";
import {AuditHeritageTask} from "@/api/heritage";
import router from "@/router";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";

export default {
  components: {
    BaseButtons,
    BaseButton, BaseLevel,
    SectionTitleLineWithButton, CardBox, SectionMain, LayoutAuthenticated},
  data() {
    return {
      fieldForm:{
        name:"",
        category:0,
        level:0,
        locate:"",
        details:""
      },
      id:"",
      cateGory:{
        0:"民间文学",
        1:"传统音乐",
        2:"传统舞蹈",
        3:"传统戏剧",
        4:"曲艺",
        5:"传统体育",
        6:"传统美术",
        7:"传统技艺",
        8:"传统医药",
        9:"民俗"
      },
      rank:{
        0:"国家级",
        1:"省级",
        2:"市级",
        3:"人类非物质文化遗产"
      },
    }
  },
  mounted() {
    this.fieldForm = JSON.parse(this.$route.query.field)
    this.id = this.$route.query.id
  },
  methods: {
    mdiKeyboardReturn() {
      return mdiKeyboardReturn
    },
    mdiPageNext() {
      return mdiPageNext
    },
    mdiBallotOutline() {
      return mdiBallotOutline
    },
    handleAudit() {
      this.$messageBox.confirm('是否确认审核', '提示', {
        confirmButtonText: 'pass',
        cancelButtonText: 'refuse',
        type: "success",
        beforeClose: (action, instance, done) => {
          let res = false;
          if (action === 'confirm') {
            res = true;
          }
          instance.confirmButtonLoading = true;
          instance.confirmButtonText = '执行中...';
          AuditHeritageTask(this.id, res).then(() => {
            instance.confirmButtonLoading = false;
            done();
            this.$message({
              type: 'success',
              message: '审核成功!'
            });
            router.back();
          }).catch(() => {
            instance.confirmButtonLoading = false;
            done();
            this.$message({
              type: 'error',
              message: '审核失败!'
            });
          })
        }
      }).catch(() => {

      })
    },
    back() {
      router.back();
    },
  }
}
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <SectionTitleLineWithButton :icon="mdiBallotOutline()" title="Heritage Inheritor Apply Message" main >
        <BaseButtons>
          <BaseButton
            :icon="mdiKeyboardReturn()"
            label="返回"
            color="info"
            rounded-full
            small
            @click="back"
          />
          <BaseButton
            target="_blank"
            :icon="mdiPageNext()"
            label="审核"
            color="success"
            rounded-full
            small
            @click="handleAudit"
          />
        </BaseButtons>
      </SectionTitleLineWithButton>
      <CardBox form >
        <CardBox>
          <BaseLevel type="justify-around lg:justify-center">
            <div class="space-y-3 text-center md:text-left lg:mx-12">
              <div class="flex  md:block">
                <p><b>Name: </b>{{fieldForm.name}}</p>
                <p><b>CateGory: </b>{{cateGory[fieldForm.category]}}</p>
                <p><b>Level: </b>{{rank[fieldForm.level]}}</p>
                <p><b>Locate: </b>{{fieldForm.locate}}</p>
                <p><b>Details:</b><span v-html="fieldForm.details"></span></p>
              </div>
            </div>
          </BaseLevel>
        </CardBox>
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>
