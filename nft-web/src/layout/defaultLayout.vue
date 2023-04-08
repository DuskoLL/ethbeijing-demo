<template>
  <div class="shadow-default-layout-container">
    <Navigator></Navigator>
    <slot></slot>
    <Footer></Footer>
  </div>
</template>

<script setup>
import Navigator from '@/components/navigator.vue';
import Footer from '@/components/footer.vue';
</script>

<style lang="scss" scoped>

</style>