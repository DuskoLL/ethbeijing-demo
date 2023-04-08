<template>
  <div class="shadow-nft-overview-container">
      <div class="shadow-nft-overview-img">
      </div>
      <div class="shadow-nft-overview-content">
        <div class="title">SavourNFT</div>
        <div class="sub-title">
          <div class="sub-title-item">Creator <span class="sub-title-black">0xe9d5...f6ef</span> </div>
          <div class="sub-title-item">Address <span class="sub-title-black">0x8955...f86b</span> </div>
          <div class="sub-title-item">Chain <span class="sub-title-black">Ethereum</span> </div>
        </div>
        <div class="desc">Ai Makita combines digital and analog methods to express her view of the human world. She creates layers of painted and digital images that are piled </div>
      </div>
      <div class="shadow-nft-overview-score">
        <div class="title">Shadow Score</div>
        <div class="flex">
          <span class="number">8</span>
          <span class="start">★</span>
        </div>
        <div ref="shadowTotalChart" style="height: 50%;"></div>
      </div>
  </div>
</template>

<script setup>
import { ref, onMounted, getCurrentInstance } from 'vue';

const shadowTotalChart = ref([]),

  { proxy } = getCurrentInstance(), // 获取全局配置项
  initMyChart = () => {
    const myChart = proxy.$echarts.init(shadowTotalChart.value),
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
            symbol: 'none',
            areaStyle: {
              color: {
                type: 'linear',
                x: 0,
                y: 0,
                x2: 0,
                y2: 1,
                colorStops: [
                  {
                    offset: 0, color: '#24AE64' // 0% 处的颜色
                  },
                  // {
                  //   offset: 0.5, color: '#FFFFFF' // 0% 处的颜色
                  // },
                  {
                    offset: 1, color: '#FFFFFF' // 100% 处的颜色
                  }
                ],
                global: false // 缺省为 false
              }
            }
          }
        ],
        grid:{ // 让图表占满容器
          top:'0px',
          left:'0px',
          right:'0px',
          bottom:'0px'
        }

      };

    myChart.setOption(option);
  };

onMounted(() => {
  initMyChart();
});

</script>

<style lang="scss" scoped>
    .shadow-nft-overview-container{
      display: flex;
      align-items: center;
      padding: 24px;
      background: #FFFFFF;
      border-radius: 24px;
      .shadow-nft-overview-img{
        width: 240px;
        height: 240px;
        background: linear-gradient(320.01deg, rgba(255, 105, 58, 0.3) 1.17%, rgba(255, 105, 58, 0) 37.61%), linear-gradient(137.73deg, rgba(247, 166, 0, 0.3) 3.85%, rgba(247, 166, 0, 0) 29.08%);
        border-radius: 32px;
        margin-right: 24px;
      }
      .title{
        color: #121214;
        font-weight: 700;
        font-size: 40px;
        line-height: 48px;
      }
      .sub-title{
        display: flex;
        font-weight: 500;
        font-size: 20px;
        line-height: 28px;
        color: #ADB1B8;
      }
      .sub-title-item{
        margin-right: 24px;
      }
      .sub-title-black{
        color: #121214;
      }
      .desc{
        font-weight: 400;
        font-size: 16px;
        line-height: 24px;
        color: #81858C;
      }
      .shadow-nft-overview-content{
        flex: 1;
        margin-right: 60px;
      }
      .shadow-nft-overview-score{
        width: 200px;
        height: 200px;
        padding: 24px 0px 0px;
        background: #FFFFFF;
        box-shadow: 2px 2px 16px rgba(18, 18, 20, 0.1);
        border-radius: 8px;
        text-align: center;
        box-sizing: border-box;
        .flex{
          display: flex;
          align-items: center;
          justify-content: center;
        }
        .title{
          font-weight: 400;
          font-size: 14px;
          line-height: 22px;
        }
        .number{
          font-weight: 700;
          font-size: 32px;
          line-height: 40px;
          background-image: linear-gradient(270.2deg, #C25FFF 42.99%, #0BA7FF 84.13%);
          background-clip: text ;
          color: transparent ;
        }
        .start{
          font-size: 21px;
          color: #16A3FF;
        }
      }
    }
</style>