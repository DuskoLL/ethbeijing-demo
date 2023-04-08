<template>
    <div>
        <div class="flex">
            <el-input  :prefix-icon="Search" class="search-input mr-32" v-model="nftsStr" placeholder="Please input" />
            <el-select v-model="nftsType" class="search-select"  placeholder="Select">
            <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
            </el-select>
        </div>
        <div class="shadow-nfts-card-container">
            <router-link to="/CollectionDetail" class="shadow-nfts-card" v-for="i in cardList" :key="i">
                <div class="shadow-nfts-img"></div>
                <div class="shadow-nfts-card-content">
                    <di class="shadow-nfts-title">Picture Of Colorful Nice YeP Illustration #1509-287</di>
                    <div class="flex justify-between">
                    <span class="c-81858C">Price</span>
                    <span>0.144 ETH</span>
                    </div>
                    <div class="text-right c-81858C">
                    ≈ $ 662.50
                    </div>
                </div>
            </router-link>
        </div>
        </div>
</template>

<script setup>
import { ref } from 'vue';

const nftsStr= ref(''),
  nftsType= ref(''),
  cardList = ref(15);
const options = [
  {
    value: 'Option1',
    label: 'Option1'
  },
  {
    value: 'Option2',
    label: 'Option2'
  },
  {
    value: 'Option3',
    label: 'Option3'
  },
  {
    value: 'Option4',
    label: 'Option4'
  },
  {
    value: 'Option5',
    label: 'Option5'
  }
];

</script>

<style lang="scss" scoped>
  .flex{
    display: flex;
    align-items: center;
  }
  .justify-center {
    justify-content: center;
  }
  .justify-between {
    justify-content: space-between;
  }
  
  .text-right{
    text-align: right;
  }
  .c-81858C{
    color: #81858C;
  }
  .mr-32{
    margin-right: 32px;
  }
  .mb-32{
    margin-bottom: 32px;
  }
  .mb-46{
    margin-bottom: 46px;
  }
  .search-input{
    width: 410px;
    height: 52px;
    background: #FFFFFF;
    border-radius: 8px;
    :deep(.el-input__wrapper){
      box-shadow: unset;
    }
    :deep(.el-input__inner){
      color: #121214;
      font-size: 16px;
      line-height: 24px;
    }
    
  }
  .search-select{
    :deep(.el-input){
      width: 280px;
      height: 52px;
    }
    :deep(.el-input__wrapper){
      border-radius: 8px;
      box-shadow: unset;
    }
    :deep(.el-input__inner){
      color: #121214;
      font-size: 16px;
      line-height: 24px;
    }
  }
  .shadow-nfts-card-container{
      display: flex;
      margin-top: 48px;
      flex-wrap: wrap;
      justify-content: space-between;
    }
    .shadow-nfts-card{
      width: 234px;
      overflow: hidden;
      /* background: #F3F5F7; */
      background-image: url('~@/assets/img/index_bg.png');
      border-radius: 16px;
      margin-bottom: 38px;
      .shadow-nfts-img{
        height: 234px;
        width: 234px;
        background: rgba(16, 16, 20, 0.7);
      }
      .shadow-nfts-card-content{
        padding: 16px;
      }
      .shadow-nfts-title{
        font-weight: 700;
        font-size: 16px;
        line-height: 19px;
        color: #121214;
      }
    }
    .shadow-score-box {
      display: flex;
      align-items: center;
      .shadow-score-logo {
        width: 60px;
        height: 60px;
        background: #d8d8d8;
        border-radius: 16px;
        margin-right: 11px;
      }
      .shadow-score-info-box {
        .name {
          display: flex;
          align-items: center;
          span {
            font-size: 18px;
            font-family: PingFangSC-Semibold, PingFang SC;
            font-weight: 600;
            color: #222222;
            line-height: 25px;
          }
          img {
            width: 12px;
            height: 12px;
            margin-top: 10px;
            margin-left: 5px;
          }
        }
        .holder {
          font-size: 10px;
          font-family: PingFangSC-Regular, PingFang SC;
          font-weight: 400;
          color: #999999;
          line-height: 14px;
        }
      }
    }
</style>