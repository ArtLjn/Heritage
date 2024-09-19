<script>
import { mdiGithub, mdiPageNext, mdiTableBorder } from "@mdi/js";
import SectionMain from '@/components/SectionMain.vue';
import CardBox from '@/components/CardBox.vue';
import LayoutAuthenticated from '@/layouts/LayoutAuthenticated.vue';
import SectionTitleLineWithButton from '@/components/SectionTitleLineWithButton.vue';
import BaseButton from "@/components/BaseButton.vue";
import HeritageInheritorTable from "@/views/show/heritageInheritorTable.vue";
import HeritageProjectTable from "@/views/show/heritageProjectTable.vue";
import { CityList } from "@/api/auth";

export default {
  components: {
    HeritageProjectTable,
    HeritageInheritorTable,
    BaseButton,
    SectionMain,
    CardBox,
    LayoutAuthenticated,
    SectionTitleLineWithButton
  },
  data() {
    return {
      province: '',
      city: localStorage.getItem("city") || "长春市",
      msg: "Next Heritage Project",
      showInheritor: true,
      showProject: false,
      title: "Heritage Inheritor",
      options: CityList(),
    };
  },
  computed: {
    cities() {
      const selectedProvince = this.options.find(option => option.label === this.province);
      return selectedProvince ? selectedProvince.children : [];
    }
  },
  methods: {
    mdiPageNext() {
      return mdiPageNext;
    },
    mdiGithub() {
      return mdiGithub;
    },
    mdiTableBorder() {
      return mdiTableBorder;
    },
    handleSwitch() {
      if (this.showInheritor) {
        this.showInheritor = false;
        this.showProject = true;
        this.msg = "Back Heritage Inheritor";
        this.title = "Heritage Project";
      } else if (this.showProject) {
        this.showProject = false;
        this.showInheritor = true;
        this.msg = "Next Heritage Project";
        this.title = "Heritage Inheritor";
      }
    },
    handleProvinceChange(event) {
      this.province = event.target.value;
      this.city = ''; // Reset city when province changes
    },
    handleCityChange(event) {
      this.city = event.target.value;
    }
  }
};
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
        <!-- 省选择框 -->
        <div class="mb-4">
          <label for="province" class="block text-sm font-medium text-gray-700">选择省份</label>
          <select
            id="province"
            v-model="province"
            @change="handleProvinceChange"
            class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
          >
            <option value="" disabled>请选择省份</option>
            <option v-for="option in options" :key="option.value" :value="option.label">{{ option.label }}</option>
          </select>
        </div>

        <!-- 市选择框，仅当有省份选择时展示 -->
        <div v-if="province" class="mb-4">
          <label for="city" class="block text-sm font-medium text-gray-700">选择城市</label>
          <select
            id="city"
            v-model="city"
            @change="handleCityChange"
            class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
          >
            <option value="" disabled>请选择城市</option>
            <option v-for="cityOption in cities" :key="cityOption.value" :value="cityOption.label">{{ cityOption.label }}</option>
          </select>
        </div>

        <!-- 表格展示 -->
        <heritage-inheritor-table v-if="showInheritor && city" :city="city" />
        <heritage-project-table v-if="showProject && city" :city="city" />
      </CardBox>
    </SectionMain>
  </LayoutAuthenticated>
</template>

<style scoped>
/* Add your custom styles here if needed */
</style>
