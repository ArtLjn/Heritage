<script>
import {mdiGithub, mdiPageNext, mdiTableBorder} from "@mdi/js";
import SectionMain from '@/components/SectionMain.vue'
import CardBox from '@/components/CardBox.vue'
import LayoutAuthenticated from '@/layouts/LayoutAuthenticated.vue'
import SectionTitleLineWithButton from '@/components/SectionTitleLineWithButton.vue'
import BaseButton from "@/components/BaseButton.vue";
import ApplyHeritageInheritorTable from "@/views/admin/applyHeritageInheritorTable.vue";
import ApplyHeritageProjectTable from "@/views/admin/applyHeritageProjectTable.vue";
export default {
  components:{
    ApplyHeritageProjectTable,
    ApplyHeritageInheritorTable,
    BaseButton,
    SectionMain,CardBox,
    LayoutAuthenticated,SectionTitleLineWithButton
  },
  data() {
    return{
      city: localStorage.getItem("city"),
      msg:"Next Heritage Project Task",
      showInheritor:true,
      showProject: false,
      title:"Heritage Inheritor Task"
    }
  },
  methods: {
    mdiPageNext() {
      return mdiPageNext
    },
    mdiGithub() {
      return mdiGithub
    },
    mdiTableBorder() {
      return mdiTableBorder
    },
    handleSwitch() {
      if (this.showInheritor) {
        this.showInheritor = false
        this.showProject = true
        this.msg = "Back Heritage Inheritor Task"
        this.title = "Heritage Project Task"
      } else if (this.showProject) {
        this.showProject = false
        this.showInheritor = true
        this.msg = "Next Heritage Project Task"
        this.title = "Heritage Inheritor Task"
      }
    }
  }
}
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <SectionTitleLineWithButton :icon="mdiTableBorder()" :title="title" main>
        <BaseButton
          target="_blank"
          :icon="mdiPageNext()"
          :label="msg"
          color="contrast"
          rounded-full
          small
          @click="handleSwitch"
        />
      </SectionTitleLineWithButton>
      <CardBox class="mb-6" has-table>
        <apply-heritage-inheritor-table v-if="showInheritor"/>
        <apply-heritage-project-table v-if="showProject"/>
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>

<style scoped>

</style>
