<template>
  <default-layout>
    <div class="view-more-container">
      <div class="banner-container">
        <img src="https://dummyimage.com/465x56/eae0d0" alt="" />
        <p>Profit faster with real time NFT market insight</p>
      </div>
      <div class="table-data-container">
        <div class="table-data-box">
          <el-tabs v-model="activeTab">
            <el-tab-pane label="Hot Collections" name="hotCollections">
              <!-- Hot Collections Table -->
              <el-table
                :data="tableData"
                :default-sort="{ prop: 'date', order: 'descending' }"
                style="width: 100%;border-radius: 24px;"
              >
                <el-table-column type="index" />
                <el-table-column
                  prop="shadowScore"
                  label="Shadow Score"
                  sortable
                  width="330"
                >
                  <template #default="{ row }">
                    <div class="shadow-score-box">
                      <img
                        class="shadow-score-logo"
                        :src="row.shadowScore.logo"
                        alt="logo"
                      />
                      <div class="shadow-score-info-box">
                        <p class="name">
                          <span>{{ row.shadowScore.name }}</span
                          ><span
                            ><img
                              src="https://dummyimage.com/12x12/eae0d0"
                              alt="logo"
                          /></span>
                        </p>
                        <p class="holder">
                          {{
                            `holder:${row.shadowScore.holder} / whale address:${row.shadowScore.whaleAddress}`
                          }}
                        </p>
                        <el-rate
                          :model-value="3.5"
                          disabled
                          text-color="#fffff"
                        />
                      </div>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column
                  prop="vllume"
                  label="24H Vllume(USDT)"
                  sortable
                >
                  <template #default="{ row }">
                    <div class="vllume-box">
                      <p>{{ row.vllume }}</p>
                      <span
                        >0.9%<img src="https://dummyimage.com/8x12/52CCA3"
                      /></span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="price" label="Price" sortable />
                <el-table-column prop="mint" label="Mint" sortable>
                  <div class="mint-chart" ref="mintChart"></div>
                </el-table-column>
                <el-table-column
                  prop="suggestion"
                  label="Suggestion"
                  sortable
                />
              </el-table>
            </el-tab-pane>
            <el-tab-pane label="Live Mint" name="liveMint">
              <!-- Hot Collections Table -->
              <el-table
                :data="tableData"
                :default-sort="{ prop: 'date', order: 'descending' }"
                style="width: 100%"
              >
                <el-table-column type="index" />
                <el-table-column
                  prop="shadowScore"
                  label="Shadow Score"
                  sortable
                  width="330"
                >
                  <template #default="{ row }">
                    <div class="shadow-score-box">
                      <img
                        class="shadow-score-logo"
                        :src="row.shadowScore.logo"
                        alt="logo"
                      />
                      <div class="shadow-score-info-box">
                        <p class="name">
                          <span>{{ row.shadowScore.name }}</span
                          ><span
                            ><img
                              src="https://dummyimage.com/12x12/eae0d0"
                              alt="logo"
                          /></span>
                        </p>
                        <p class="holder">
                          {{
                            `holder:${row.shadowScore.holder} / whale address:${row.shadowScore.whaleAddress}`
                          }}
                        </p>
                        <el-rate
                          :model-value="3.5"
                          disabled
                          text-color="#fffff"
                        />
                      </div>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column
                  prop="vllume"
                  label="24H Vllume(USDT)"
                  sortable
                >
                  <template #default="{ row }">
                    <div class="vllume-box">
                      <p>{{ row.vllume }}</p>
                      <span
                        >0.9%<img src="https://dummyimage.com/8x12/52CCA3"
                      /></span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="price" label="Price" sortable />
                <el-table-column prop="mint" label="Mint" sortable>
                  <div class="mint-chart" ref="mintChart"></div>
                </el-table-column>
                <el-table-column
                  prop="suggestion"
                  label="Suggestion"
                  sortable
                />
              </el-table>
            </el-tab-pane>
            <el-tab-pane label="Hot NFT" name="hotNft">
              <!-- Hot Collections Table -->
              <el-table
                :data="tableData"
                :default-sort="{ prop: 'date', order: 'descending' }"
                style="width: 100%"
              >
                <el-table-column type="index" />
                <el-table-column
                  prop="shadowScore"
                  label="Shadow Score"
                  sortable
                  width="330"
                >
                  <template #default="{ row }">
                    <div class="shadow-score-box">
                      <img
                        class="shadow-score-logo"
                        :src="row.shadowScore.logo"
                        alt="logo"
                      />
                      <div class="shadow-score-info-box">
                        <p class="name">
                          <span>{{ row.shadowScore.name }}</span
                          ><span
                            ><img
                              src="https://dummyimage.com/12x12/eae0d0"
                              alt="logo"
                          /></span>
                        </p>
                        <p class="holder">
                          {{
                            `holder:${row.shadowScore.holder} / whale address:${row.shadowScore.whaleAddress}`
                          }}
                        </p>
                        <el-rate
                          :model-value="3.5"
                          disabled
                          text-color="#fffff"
                        />
                      </div>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column
                  prop="vllume"
                  label="24H Vllume(USDT)"
                  sortable
                >
                  <template #default="{ row }">
                    <div class="vllume-box">
                      <p>{{ row.vllume }}</p>
                      <span
                        >0.9%<img src="https://dummyimage.com/8x12/52CCA3"
                      /></span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column prop="price" label="Price" sortable />
                <el-table-column prop="mint" label="Mint" sortable>
                  <div class="mint-chart" ref="mintChart"></div>
                </el-table-column>
                <el-table-column
                  prop="suggestion"
                  label="Suggestion"
                  sortable
                />
              </el-table>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
      <div class="load-more">
          <img src="https://dummyimage.com/30x30/eae0d0" alt="">
      </div>
    </div>
  </default-layout>
</template>

<script setup>
import DefaultLayout from '@/layout/defaultLayout.vue';
import { ref, onMounted, getCurrentInstance } from 'vue';
const activeTab = ref('hotCollections'),
  mintChart = ref(null),
  { proxy } = getCurrentInstance(),
  tableData = [
    {
      shadowScore: {
        name: 'savour groups',
        holder: 1000,
        whaleAddress: 100,
        logo: 'https://dummyimage.com/60x60/eae0d0',
        mark: 4.5
      },
      vllume: 1000,
      price: 200,
      mint: 'Tom',
      suggestion: 12300
    },
    {
      shadowScore: {
        name: 'savour groups',
        holder: 1000,
        whaleAddress: 100,
        logo: 'https://dummyimage.com/60x60/eae0d0',
        mark: 4.5
      },
      vllume: 1000,
      price: 200,
      mint: 'Tom',
      suggestion: 12300
    },
    {
      shadowScore: {
        name: 'savour groups',
        holder: 1000,
        whaleAddress: 100,
        logo: 'https://dummyimage.com/60x60/eae0d0',
        mark: 4.5
      },
      vllume: 1000,
      price: 200,
      mint: 'Tom',
      suggestion: 12300
    },
    {
      shadowScore: {
        name: 'savour groups',
        holder: 1000,
        whaleAddress: 100,
        logo: 'https://dummyimage.com/60x60/eae0d0',
        mark: 4.5
      },
      vllume: 1000,
      price: 200,
      mint: 'Tom',
      suggestion: 12300
    },
    {
      shadowScore: {
        name: 'savour groups',
        holder: 1000,
        whaleAddress: 100,
        logo: 'https://dummyimage.com/60x60/eae0d0',
        mark: 4.5
      },
      vllume: 1000,
      price: 200,
      mint: 'Tom',
      suggestion: 12300
    }
  ];

function initMyChart() {
  const myChart = proxy.$echarts.init(mintChart.value),
    option = {
      xAxis: {
        show: false,
        type: 'category',
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
      },
      yAxis: {
        type: 'value',
        show: false
      },
      color: ['#20B26C'],
      series: [
        {
          data: [150, 230, 224, 218, 135, 147, 260],
          type: 'line',
          symbol: 'none'
        }
      ],
      grid: {
        // 让图表占满容器
        top: '0px',
        left: '0px',
        right: '0px',
        bottom: '0px'
      }
    };

  myChart.setOption(option);
}

onMounted(() => {
  setTimeout(() => {
    initMyChart();
  }, 1000);
  console.dir(mintChart, mintChart.value, '2222');
});
</script>

<style lang="scss" scoped>
.banner-container {
  width: 100%;
  height: 234px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  img {
    width: 465px;
    height: 56px;
  }
  p {
    /* Regular/16 */
    font-family: "IBM Plex Sans";
    font-style: normal;
    font-weight: 400;
    font-size: 16px;
    line-height: 24px;
    /* identical to box height, or 150% */
    /* Gray [T]/T1_Title */
    color: #121214;
    margin-top: 16px;
  }
}
.table-data-container {
  width: 100%;
  .table-data-box {
    width: 1360px;
    margin: 0 auto;
    box-sizing: border-box;
    :deep(.el-tabs__nav-scroll) {
      padding: 0 24px;
    }
    :deep(.el-tabs__nav-wrap::after) {
      height: 0;
    }
    :deep(.el-tabs__active-bar) {
        background: linear-gradient(90deg, #C65EFF 0%, #9372FF 26%, #03ABFF 100%);
    }
    :deep(.el-tabs__item) {
      font-size: 24px;
      font-family: PingFangSC-Semibold, PingFang SC;
      font-weight: 600;
      line-height: 33px;
      color: #81858c;
    }
    :deep(.el-tabs__item.is-active) {
      color: #222;
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
  }
}
.load-more{
    width: 1360px;
    height: 100px;
    margin: 0 auto;
    display: flex;
    justify-content: center;
    align-items: center;
    img{
        cursor: pointer;
    }
}
</style>