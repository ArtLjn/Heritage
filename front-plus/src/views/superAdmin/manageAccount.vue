<script>
import {mdiGithub, mdiTableBorder} from "@mdi/js";
import SectionMain from '@/components/SectionMain.vue'
import CardBox from '@/components/CardBox.vue'
import LayoutAuthenticated from '@/layouts/LayoutAuthenticated.vue'
import SectionTitleLineWithButton from '@/components/SectionTitleLineWithButton.vue'
import AccountTable from "@/views/superAdmin/accountTable.vue";
export default {
  components:{
    AccountTable,
    SectionMain,CardBox,
    LayoutAuthenticated,SectionTitleLineWithButton
  },
  methods: {
    mdiGithub() {
      return mdiGithub
    },
    mdiTableBorder() {
      return mdiTableBorder
    }
  }
}
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <SectionTitleLineWithButton :icon="mdiTableBorder()" title="Account Manage" main>
      </SectionTitleLineWithButton>
      <CardBox class="mb-6" has-table>
        <account-table checkable  />
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>

<style scoped>

</style>
