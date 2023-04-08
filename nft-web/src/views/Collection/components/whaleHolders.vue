<template>
    <div>
        <div class="flex justify-between mb-46">
        <div class="flex">
            <div class="mr-32 holders-type">
                <div :class="holdersType === 'total' ? 'holders-type-item active': 'holders-type-item'">Total</div>
                <div :class="holdersType === 'average' ? 'holders-type-item active': 'holders-type-item'">Average</div>
            </div>
            <el-select v-model="holdersTime" class="search-select"  placeholder="Select">
            <el-option
                v-for="item in hourRange"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
            </el-select>
        </div>
        <el-input :prefix-icon="Search" class="search-input" v-model="holdersStr" placeholder="Please input" />
        </div>
        <div class="sub-title mb-32">Holders</div>
        <div class="holders-container">
        <div class="title">Whale Holder</div>
        <el-table
            :data="tableData"
            :default-sort="{ prop: 'date', order: 'descending' }"
            style="width: 100%"
        >
            <el-table-column
            prop="address"
            label="Address"
            sortable
            width="330"
            />
            <el-table-column prop="tokenValue" label="Token Value" sortable>
            <template #default="{ row }">
                <div class="vllume-box">
                <p>{{ row.tokenValue }}</p>
                <span
                    >0.9%<img src="https://dummyimage.com/8x12/52CCA3"
                /></span>
                </div>
            </template>
            </el-table-column>
            <el-table-column prop="owned" label="Owned" sortable />
            <el-table-column prop="realizedPnL" label="Realized PnL" sortable/>
            <el-table-column prop="label" label="Label" sortable />
        </el-table>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { Search } from '@element-plus/icons-vue';

const holdersType = ref('total'),
  holdersTime = ref('lastHour'),
  holdersStr= ref(''),
  hourRange = [
    {
      value: 'lastHour',
      label: 'Last Hour'
    },
    {
      value: 'last24Hour',
      label: 'Last 24 Hour'
    },
    {
      value: 'lastDay',
      label: 'Last Day'
    },
    {
      value: 'last7Days',
      label: 'Last 7 Days'
    },
    {
      value: 'last30Days',
      label: 'Last 30 Days'
    }
  ],
  tableData = [
    {
      address: '0x8955...f86b',
      tokenValue: 1000,
      owned: 200,
      realizedPnL: 'Tom',
      label: 12300
    },
    {
      address: '0x8955...f86b',
      tokenValue: 1000,
      owned: 200,
      realizedPnL: 'Tom',
      label: 12300
    },
    {
      address: '0x8955...f86b',
      tokenValue: 1000,
      owned: 200,
      realizedPnL: 'Tom',
      label: 12300
    },
    {
      address: '0x8955...f86b',
      tokenValue: 1000,
      owned: 200,
      realizedPnL: 'Tom',
      label: 12300
    },
    {
      address: '0x8955...f86b',
      tokenValue: 1000,
      owned: 200,
      realizedPnL: 'Tom',
      label: 12300
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
  .holders-container{
    background: #FFFFFF;
    border-radius: 16px;
    padding: 24px;
    .title{
      font-weight: 600;
      font-size: 20px;
      line-height: 28px;
      color: #121214;
      margin-bottom: 28px;
    }
  }
</style>