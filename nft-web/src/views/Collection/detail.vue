<template>
    <default-layout>
        <div class="collection-detail-container">
            <div class="mb-32 mt-40 flex">
                <div class="shadow-nfts-card">
                    <div class="shadow-nfts-img"></div>
                    <div class="shadow-nfts-card-content">
                        <div class="shadow-nfts-title">Creator <span class="c-121214">0x9aBa...6F9b</span>  0x8955...f86b</div>
                        <div>Description:</div>
                        <div class="c-81858C">Ai Makita combines digital and analog methods to express her view of the human world. She creates layers of painted and digital images that are piled</div>
                    </div>
                </div>
                <div class="overview-content">
                    <div class="title">SavourNFT</div>
                    <div class="sub-title">Holder <span class="c-121214"></span>  0x8955...f86b</div>
                    <el-divider />
                    <div class="overview-content-list">
                        <div class="overview-content-item" v-for="item in options" :key="item.title">
                            <div class="left">{{item.title}}</div>
                            <div class="right">{{item.value}}</div>
                        </div>
                    </div>
                </div>
            </div>
            <price-card :list="priceCardList" class="mb-32"/>
            <div class="card">
                <div class="title">项目活动</div>
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
    </default-layout>
</template>

<script setup>
import { ref } from 'vue';
import DefaultLayout from '@/layout/defaultLayout.vue';
import PriceCard from './components/priceCard.vue';
const tableData = [
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
  ],

  priceCardList = ref([{
    title: 'Price',
    price: '1000',
    percent: '2.48%'
  },{
    title: 'Total Holders',
    price: '12,121',
    percent: '2.48%'
  },{
    title: 'Total Whale Holders',
    price: '12,121',
    percent: '2.48%'
  },{
    title: 'Total TXNs',
    price: '2.68',
    percent: '2.48%'
  }
  ]),
  options = [
    {
      title: 'contract adress',
      value: '0x000000000000000000000'
    },{
      title: 'tokenId',
      value: '384111112323233'
    },{
      title: 'Content URL',
      value: 'QmPBS8UwZJfMDftLSxB946PgsDb1FQwnVcCYuXfWoAyxqSView'
    },{
      title: 'Minting Hash',
      value: '0x5346a5e70a8cc2a4deb85b0ea088a45e9131db7a499fe0bb755c5440a5f57d37'
    },{
      title: 'Minting Date',
      value: '2023-01-30 12:30:47'
    },{
      title: '其他字段',
      value: '2023-01-30 12:30:47'
    },{
      title: '其他字段',
      value: '2023-01-30 12:30:47'
    }
  ];

</script>

  <style lang="scss" scoped>
  .flex{
    display: flex;
  }
  .mt-40{
    margin-top: 40px;
  }
  .mb-32{
    margin-bottom: 32px;
  }
  .c-121214{
    color: #121214;
  }
  .c-81858C{
    color: #81858C;
  }
  .collection-detail-container{
    margin: 0 40px;
    .card{
      padding: 24px;
      background: #FFFFFF;
      border-radius: 24px;
      box-sizing: border-box;
      margin-bottom: 32px;
      .title{
        font-weight: 600;
        font-size: 20px;
        line-height: 28px;
        color: #121214;
        }
    }
    .shadow-nfts-card{
      width: 580px;
      overflow: hidden;
      background: #fff;
      border-radius: 24px;
      margin: 0 32px 38px 0;
      .shadow-nfts-img{
        width: 580px;
        height: 642px;
        background: linear-gradient(0deg, rgba(0, 0, 0, 0.4), rgba(0, 0, 0, 0.4)), url('~@/assets/img/index_bg.png');
      }
      .shadow-nfts-card-content{
        padding: 24px;
        font-weight: 400;
        font-size: 16px;
        line-height: 24px;
      }
      .shadow-nfts-title{
        font-weight: 600;
        font-size: 24px;
        line-height: 32px;
        color: #81858C;
        margin-bottom: 24px;
      }
    }
    .overview-content{
        padding: 24px;
        width: 748px;
        height: 640px;
        background: #FFFFFF;
        border-radius: 24px;
        .title{
            font-weight: 700;
            font-size: 40px;
            line-height: 48px;
            color: #121214;
            margin-bottom: 12px;
        }
        .sub-title{
            font-weight: 500;
            font-size: 20px;
            line-height: 28px;
            color: #ADB1B8;
        }
        .overview-content-list{
            .overview-content-item{
                display: flex;
                margin-bottom: 28px;
                .left{
                    width: 224px;
                    font-weight: 400;
                    font-size: 16px;
                    line-height: 24px;
                    color: #81858C;
                }
                .right{
                    font-weight: 500;
                    font-size: 16px;
                    line-height: 24px;
                    color: #121214;
                    flex: 1;
                    width: 0;
                    white-space: pre-wrap;
                    word-break: break-all;
                }

            }

        }
    }
  }
  </style>