<template>
  <default-layout>
    <div class="shadow-setting-container">
        <banner />
        <div  class="main">
          <basic-info />
          <tabs-table />
        </div>
    </div>
  </default-layout>
</template>

<script setup>
import { onMounted } from 'vue';
import DefaultLayout from '@/layout/defaultLayout.vue';
import Banner from './components/banner.vue';
import BasicInfo from './components/basicInfo.vue';
import TabsTable from './components/tabsTable.vue';
import { getHotCollectionDetail } from '@/assets/js/http.js';

onMounted(async () => {
  const res = await getHotCollectionDetail({
    // eslint-disable-next-line
    collection_id: 1,
    page: 1,
    // eslint-disable-next-line
    page_size: 10
  });

  console.log(1111111, res);
});
</script>

<style lang="scss" scoped>
.main {
  position: relative;
  top: -68px;
  margin: 0 40px;
}
</style>