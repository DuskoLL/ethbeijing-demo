<template>
    <div class="overview-shadow">
        <div class="overview-shadow-score-box">
        <div class="overview-shadow-score-box-main">
            <div class="what-shadow-score-left-box">
            <h6 class="score-title">Shadow Score</h6>
            <div class="score-chart" ref="scoreChart"></div>
            </div>
            <div class="what-shadow-score-right-box">
              <p class="shadow-score-desc-title">概要</p>
              <p class="shadow-score-desc">蓝筹度: 蓝筹度越高，说明该 collection 被巨鲸持有量越大, 集合里面的 NFT 越有价值</p>
              <p class="shadow-score-desc">流动性: 流动性越高，说明该 collection 挂单深度越深并且换手率越高，集合里面的 NFT 越有价值</p>
              <p class="shadow-score-desc">潜在收益: 潜在收益越高，说明该 collection 均价上升速率快, 地板价和你的成本差距大, 集合里面的 NFT 越有价值</p>
              <p class="shadow-score-desc">热度: 热度越高，说明该 collection Mint 或者交易次数比较多，集合里面的 NFT 越有价值</p>
              <p class="shadow-score-desc">社区活跃度: 社区活跃度越高，说明关注该 NFT 的人比较多，集合里面的 NFT 越有价值</p>
              <p class="shadow-score-desc">风险: 风险与收益是并存的，风险越高，收益也就越大 </p>
            </div>
        </div>
        <div class="flex justify-center">
            <el-icon><Warning /></el-icon>  <span class="c-81858C">推荐指数仅供参考，请注意投资风险</span>
        </div>
        </div>
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
        <price-card :list="priceCardList" class="mb-32"/>
        <div class="sub-title mb-32">Trading</div>
        <div class="card-container mb-32">
          <div class="card">
            <div class="title">成交价格</div>
          </div>
          <div class="card">
            <div class="title">Volume</div>
          </div>
          <div class="card">
            <div class="title">List</div>
          </div>
          <div class="card">
            <div class="title">Floor Price</div>
          </div>
        </div>
        <div class="sub-title mb-32">Holders</div>
        <whale-holder/>
        <div class="sub-title mb-32">Social Trends</div>
        <div class="card-container mb-32">
          <div class="card">
            <div class="title">Twitter Explore</div>
          </div>
          <div class="card">
            <div class="title">搜索热度指数</div>
          </div>
        </div>
    </div>
</template>

<script setup>
import WhaleHolder from './whaleHolder.vue';
import PriceCard from './priceCard.vue';

import { ref, onMounted, getCurrentInstance } from 'vue';
import { Search, Warning } from '@element-plus/icons-vue';

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
  scoreChart = ref(null);

const priceCardList = ref([{
  title: 'Total Price',
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
]);

const { proxy } = getCurrentInstance();

let data = [80, 70, 30, 85, 25],
  indicatorname = ['蓝筹度', '流动性', '热度', '社区活跃度', '风险'],
  maxdata = [100, 100, 100, 100, 100],
  optionData = null,
  indicator = [];

const contains = (arrays, obj) => {
  let i = arrays.length;

  while (i--) {
    if (arrays[i] === obj) {
      return i;
    }
  }
  return res;
};
const innerdata = (i) => {
  let innerdata2 = [];

  for (let j = 0; j < data.length; j++) {
    innerdata2.push(100 - 20 * i);
  }
  return innerdata2;
};
const getData = () => {
  let res = {
    series: [{
      type: 'radar',
      symbolSize: 10,
      symbol: 'circle',
      areaStyle: {
        color: '#39B2FF',
        opacity: 0.3
      },
      lineStyle: {
        width: 3
      },
      itemStyle: {
        color: '#fff ',
        borderWidth: 4,
        opacity: 1
      },
      label: {
        show: false
      },
      data: [{
        value: data
      }],
      z: 100
    }]
  };

  for (let i = 0; i < data.length; i++) {
    res.series.push({
      type: 'radar',
      data: [{
        value: innerdata(i)
      }],
      symbol: 'none',
      lineStyle: {
        width: 0
      },
      itemStyle: {
        color: '#fff'
      },
      areaStyle: {
        color: '#fff',
        shadowColor: 'rgba(14,122,191,0.15)',
        shadowBlur: 30,
        shadowOffsetY: 20
      }
    });
  }
  return res;
};

const initScoreChart = () => {
  optionData = getData();
  for (let i = 0; i < indicatorname.length; i++) {
    indicator.push({
      name: indicatorname[i],
      max: maxdata[i]
    });
  }
  const myChart = proxy.$echarts.init(scoreChart.value),
    option = {
      backgroundColor: '#fff',
      tooltip: {
        formatter: function() {
          let html = '';

          for (let i = 0; i < data.length; i++) {
            html += indicatorname[i] + ' : ' + data[i] + '%<br>';
          }
          return html;
        }
      },
      radar: {
        indicator: indicator,
        splitArea: {
          show: true,
          areaStyle: {
            color: '#fff',
            shadowColor: 'rgba(14,122,191,0.19)',
            shadowBlur: 30,
            shadowOffsetY: 20
          }
        },
        splitLine: {
          show: false

        },
        axisLine: {
          show: false
        },
        axisLabel: {
          show: false
        },
        name: {
          textStyle: {
            rich: {
              a: {
                fontSize: '17',
                color: '#333',
                align: 'left',
                lineHeight: '30',
                fontWeight: 'bold'
              },
              b: {
                fontSize: '15',
                color: '#666',
                align: 'left'
              }
            }
          },

          formatter: function(params) {
            let i = contains(indicatorname, params),
              percent = data[i] / 100 * 100;

            return '{a|' + percent + '%}\n' + '{b|' + params + '}';
          }
        }
      },
      series: optionData.series
    };

  myChart.setOption(option);
};

onMounted(() => {
  setTimeout(() => {
    initScoreChart();
  }, 1000);
});

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

  .card-container{
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    .card{
      width: 664px;
      height: 468px;
      padding: 24px;
      background: #FFFFFF;
      border-radius: 24px;
      box-sizing: border-box;
      margin-bottom: 32px
    }
    .title{
      font-weight: 600;
      font-size: 20px;
      line-height: 28px;
      color: #121214;
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
  .overview-shadow-score-box {
      padding: 24px;
      background: #ffffff;
      border-radius: 24px;
      margin: 0 auto 32px;
      box-sizing: border-box;

      .overview-shadow-score-box-main{
        display: flex;
        justify-content: space-between;
      }
      .what-shadow-score-left-box{
          width: 500px;
          .score-title{
            font-size: 24px;
            font-family: PingFangSC-Semibold, PingFang SC;
            font-weight: 600;
            color: #222222;
            line-height: 33px;
            margin-bottom: 10px;
          }
          .score-chart{
              width: 500px;
              height: 500px;
          }
      }
      .what-shadow-score-right-box{
          width: 657px;
          .shadow-score-desc{
            font-weight: 400;
            font-size: 18px;
            line-height: 26px;
            color: #121214;
          }
          .shadow-score-desc-title{
            font-weight: 600;
            font-size: 18px;
            line-height: 26px;
            color: #121214;
            line-height: 33px;
            margin: 0 0 32px;
          }
      }
      .search-box{
        width: 410px;
        height: 46px;
        background: #F3F5F7;
        border-radius: 8px;
        box-sizing: border-box;
        padding: 0 14px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        .search-icon{
            width: 16px;
            height: 46px;
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .search-ipt{
            width: 380px;
            height: 46px;
            box-sizing: border-box;
            padding: 0 8px;
            border: none;
            background: #F3F5F7;
            // color: #ADB1B8;
            font-size: 16px;
            line-height: 46px;
        }
      }
    }
</style>