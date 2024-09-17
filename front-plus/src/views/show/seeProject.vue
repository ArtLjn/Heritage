<script>
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import SectionMain from "@/components/SectionMain.vue";
import CardBox from "@/components/CardBox.vue";
import SectionTitleLineWithButton from "@/components/SectionTitleLineWithButton.vue";
import {mdiBallotOutline, mdiKeyboardReturn} from "@mdi/js";
import BaseLevel from "@/components/BaseLevel.vue";
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
      fieldForm: {
        number: '',
        name: '',
        category: '',
        locate: '',
        rank: '',
        details: '',
      },
    }
  },
  mounted() {
    this.fieldForm = this.$route.query
  },
  methods: {
    mdiKeyboardReturn() {
      return mdiKeyboardReturn
    },
    mdiBallotOutline() {
      return mdiBallotOutline
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
      <SectionTitleLineWithButton :icon="mdiBallotOutline()" title="Heritage Inheritor Message" main>
        <BaseButtons>
          <BaseButton
            :icon="mdiKeyboardReturn()"
            label="返回"
            color="info"
            rounded-full
            small
            @click="back"
          />
        </BaseButtons>
      </SectionTitleLineWithButton>
      <CardBox form>
        <CardBox>
          <BaseLevel type="justify-around lg:justify-center">
            <div class="space-y-3 text-center md:text-left lg:mx-12">
              <div class="flex justify-center md:block">
                <p><b>Number: </b>{{ fieldForm.number }}</p>
                <p><b>Name: </b>{{ fieldForm.name }}</p>
                <p><b>CateGory: </b>{{fieldForm.category}}</p>
                <p><b>Rank: </b>{{fieldForm.rank}}</p>
                <p><b>Locate: </b>{{ fieldForm.locate }}</p>
                <p><b>Details:</b><span v-html="fieldForm.details"></span></p>
              </div>
            </div>
          </BaseLevel>
        </CardBox>
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>
