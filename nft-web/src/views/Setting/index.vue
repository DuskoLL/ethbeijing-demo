<template>
  <default-layout>
    <div class="shadow-setting-container">
        <customize />
        <what-shadow-score />
    </div>
  </default-layout>
</template>

<script setup>
import DefaultLayout from '@/layout/defaultLayout.vue';
import Customize from './components/customize.vue';
import WhatShadowScore from '@/components/whatShadowScore.vue';
</script>

<style lang="scss" scoped>
</style>