<template>
  <div class="whale-holder-container">
    <div class="whale-holder-box">
      <div class="whale-holder-title">
        <h6>WhaleHolder</h6>
        <p>Top100</p>
      </div>
      <div class="whale-holer-table">
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
                  <el-rate :model-value="3.5" disabled text-color="#fffff" />
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="vllume" label="24H Vllume(USDT)" sortable>
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
          <el-table-column prop="suggestion" label="Suggestion" sortable />
        </el-table>
        <!-- view more -->
        <div class="view-more-box" @click="handleToViewMore">{{ `View More` }}</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, getCurrentInstance } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter(),
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
  ],
  handleToViewMore = () => {
    router.push({
      path: '/ViewMore'
    });
  };

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
.whale-holder-container {
  width: 100%;
  .whale-holder-box {
    padding: 32px 0 40px;
    background: #ffffff;
    border-radius: 24px;
    margin: 0 auto 32px;
    .whale-holder-title {
      padding: 0 24px 0;
      h6 {
        font-size: 24px;
        font-family: PingFangSC-Semibold, PingFang SC;
        font-weight: 600;
        color: #222222;
        line-height: 33px;
        margin-bottom: 10px;
      }
      p {
        font-size: 14px;
        font-family: PingFangSC-Semibold, PingFang SC;
        font-weight: 600;
        color: #222222;
        line-height: 20px;
      }
    }
    .whale-holer-table {
      :deep(.el-tabs__nav-scroll) {
        padding: 0 24px;
      }
      :deep(.el-tabs__nav-wrap::after) {
        height: 0;
      }
      :deep(.el-tabs__active-bar) {
        height: 0;
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
}
</style>