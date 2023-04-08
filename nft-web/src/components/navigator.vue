<template>
  <div class="shadow-nft-nav-container">
      <div class="shadow-nft-nav-box">
          <div class="nav-logo-box">
              <img src="../assets/img/logo.png" alt="nav-logo" class="nav-logo">
              <div class="search-box">
                  <span class="search-icon"><img width="12" height="12" src="../assets/img/icon.png" alt="search-icon"></span>
                  <input placeholder="Search by collection, NFT" type="text" class="search-ipt">
              </div>
          </div>
          <ul class="nav-list">
              <li class="nav-item"><router-link to="/Markets" class="nav-word activted">Index</router-link></li>
              <li class="nav-item"><router-link to="/Collection" class="nav-word">Collection</router-link></li>
              <li class="nav-item"><a href="#" class="nav-word">Docs</a></li>
              <li class="nav-item"><a href="#" class="nav-word">Open API</a></li>
              <li class="nav-item"><router-link to="/Setting" class="nav-word">Setting</router-link></li>
          </ul>
      </div>
  </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
    @supports(background-clip: text) {
        .activted{
            background-image: linear-gradient(270.2deg, #C25FFF 42.99%, #0BA7FF 84.13%);
            background-clip: text !important;
            color: transparent !important;
        }
    }
    .shadow-nft-nav-container{
        width: 100%;
        height: 78px;
        background-color: #fff;
        .shadow-nft-nav-box{
            width: 1360px;
            height: 78px;
            margin: 0 auto;
            box-sizing: border-box;
            padding: 16px 40px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .nav-logo-box{
                width: 714px;
                display: flex;
                justify-content: space-between;
                align-items: center;
                .nav-logo{
                    width: 280px;
                    margin-right: 24px;
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
            .nav-list{
                display: flex;
                align-items: center;
                .nav-item{
                    margin: 0 20px;
                    .nav-word{
                        cursor: pointer;
                        font-style: normal;
                        font-weight: 500;
                        font-size: 16px;
                        line-height: 24px;
                        color: #121214;
                        &.activted{
                            color: #0BA7FF;
                        }
                    }
                }
            }
        }
    }
</style>