<template>
  <div class="shadow-nft-tabs-table-container">
    <div class="shadow-nft-tabs-table-box">
      <el-tabs v-model="activeTab">
        <el-tab-pane label="Overview" name="overview">
          <overview/>
        </el-tab-pane>
        <el-tab-pane label="NFTs" name="nfts">
          <nfts/>
        </el-tab-pane>
        <el-tab-pane label="Whale Holders" name="whaleHolders">
          <whale-holders/>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
import Overview from './overview.vue';
import Nfts from './nfts.vue';
import whaleHolders from './whaleHolders.vue';
import { ref } from 'vue';

const activeTab = ref('overview');

</script>

<style lang="scss" scoped>
.shadow-nft-tabs-table-container {
  width: 100%;

  .holders-type{
    padding: 4px;
    background: #E9EDF2;
    border-radius: 12px;
    display: flex;
    .holders-type-item{
      width: 150px;
      height: 44px;
      border-radius: 4px;
      font-weight: 600;
      font-size: 16px;
      line-height: 24px;
      color: #ADB1B8;
      display: flex;
      justify-content: center;
      align-items: center;
      &.active{
        background: #FFFFFF;
        color: #121214;
      }
    }
  }
  .sub-title{
    font-weight: 600;
    font-size: 24px;
    line-height: 32px;
    color: #121214;
  }

  .shadow-nft-tabs-table-box {
    padding: 32px 0 40px;
    :deep(.el-tabs__nav-wrap::after) {
      height: 0;
    }
    :deep(.el-tabs__active-bar) {
      height: 0;
    }
    :deep(.el-tabs__item) {
      font-size: 24px;
      font-weight: 600;
      line-height: 33px;
      color: #81858c;
    }
    :deep(.el-tabs__item.is-active) {
      color: #222;
    }
    .vllume-box {
      display: flex;
      align-items: center;
      justify-content: flex-start;
      p {
        font-size: 18px;
        font-family: PingFangSC-Medium, PingFang SC;
        font-weight: 500;
        color: #222222;
        line-height: 25px;
        margin-right: 5px;
      }
      span {
        font-size: 14px;
        font-family: Helvetica;
        color: #52cca3;
        line-height: 12px;
        display: flex;
        align-items: center;
        img {
          margin-left: 2px;
        }
      }
    }
    .mint-chart {
      width: 115px;
      height: 44px;
    }
    .view-more-box {
      width: 1312px;
      height: 52px;
      background: #f5f5fc;
      border-radius: 10px;
      margin: 24px auto 0;
      /* Regular/16 */
      font-family: "IBM Plex Sans";
      font-style: normal;
      font-weight: 400;
      font-size: 16px;
      line-height: 52px;
      /* identical to box height, or 150% */
      text-align: center;
      /* Gray [T]/T1_Title */
      color: #121214;
      cursor: pointer;
    }
  }
}
</style>