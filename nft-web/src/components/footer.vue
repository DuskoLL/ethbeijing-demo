<template>
  <div class="shadow-nft-footer-container">
    <div class="shadow-nft-footer-box">
      <div class="footer-left-box">
        <img src="../assets/img/logo.png" alt="footer-logo">
        <p>最新最全NFT市场情报，独创NFT推荐指数，更快一步获取收益</p>
      </div>
      <div class="footer-right-box">
        <div class="about-us-box">
          <h6 class="title">About us</h6>
          <a href="#">Platform Introduction</a>
          <a href="#">Service Agreements</a>
          <a href="#">Contact  Us</a>
        </div>
        <div class="support-box">
          <h6 class="title">Support</h6>
          <a href="#">Help Center</a>
          <a href="#">Docs</a>
          <a href="#">API</a>
        </div>
        <div class="join-us-box">
          <h6 class="title">Join us</h6>
          <a href="https://twitter.com/0xsavourlabs"><img src="../assets/img/twitter.svg" alt=""></a>
          <a href="https://t.me/+qqhy1i-_xnU1M2Jl"><img src="../assets/img/telegram.svg" alt=""></a>
          <a href="https://discord.gg/9N62gJgMVW"><img src="../assets/img/discord.svg" alt=""></a>
          <a href="https://discord.gg/9N62gJgMVW"><img src="../assets/img/github.svg" alt=""></a>

        </div>
      </div>
    </div>
    <p class="shadow-ntf-footer-hint-box">
      推荐指数仅供参考，请注意投资风险
    </p>
  </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
  .shadow-nft-footer-container{
    width: 100%;
    background-color: #fff;
    .shadow-nft-footer-box{
      width: 1360px;
      box-sizing: border-box;
      margin: 0 auto;
      padding: 40px 40px 0;
      display: flex;
      justify-content: space-between;
      .footer-left-box{
        img{
          width: 280px;
        }
        p{
          font-style: normal;
          font-weight: 400;
          font-size: 14px;
          line-height: 22px;
          color: #ADB1B8;
          margin-top: 11px;
        }
      }
      .footer-right-box{
        display: flex;
        .title{
          font-style: normal;
          font-weight: 600;
          font-size: 18px;
          line-height: 26px;
          color: #121214;
          margin-bottom: 24px;
        }
        .about-us-box,.support-box{
          display: flex;
          flex-direction: column;
          margin-left: 80px;
          a{
            font-style: normal;
            font-weight: 500;
            font-size: 14px;
            line-height: 22px;
            color: #ADB1B8;
            margin-bottom: 24px;
          }
        }
        .join-us-box{
          margin-left: 80px;
          a{
            display: inline-block;
            width: 40px;
            height: 40px;
            margin-right: 16px;
          }
        }
      }
    }
    .shadow-ntf-footer-hint-box{
      width: 100%;
      border-top: 1px solid #E9EDF2;
      font-style: normal;
      font-weight: 400;
      font-size: 16px;
      line-height: 24px;
      padding: 24px 0;
      text-align: center;
    }
  }
</style>