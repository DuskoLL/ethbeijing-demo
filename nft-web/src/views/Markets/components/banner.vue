<template>
    <div class="shadow-nft-banner-container">
        <div class="shadow-nft-banner-box">
            <div class="shadow-nft-banner-left-box">
<!--                <img src="../../../assets/img/logo.png" alt="" class="logo">-->
                <p class="logo-desc-text">最新最全NFT市场情报，独创NFT推荐指数，更快一步获取收益</p>
                <div class="banner-nav-list">
                    <span>Support Public Chain</span>
                    <a href="#"><img width="16" height="16" src="../../../assets/img/logo_scroll.png" alt=""></a>
                    <a href="#"><img width="16" height="16" src="../../../assets/img/log_eth.png" alt=""></a>
                </div>
            </div>
            <div class="shadow-nft-banner-right-box">
<!--                <img src="https://dummyimage.com/404x165/eae0d0" alt="">-->
            </div>
        </div>
    </div>
</template>

<script setup>
</script>

<style lang="scss" scoped>
    .shadow-nft-banner-container{
        width: 100%;
        .shadow-nft-banner-box{
            width: 1360px;
            height: 280px;
            margin: 0 auto;
            display: flex;
            justify-content: space-between;
            box-sizing: border-box;
            padding: 67.5px 0;
            .shadow-nft-banner-left-box{
                .logo{
                    width: 280px;
                }
                .logo-desc-text{
                    font-style: normal;
                    font-weight: 400;
                    font-size: 16px;
                    line-height: 24px;
                    color: #121214;
                    margin-top: 10px;
                }
                .banner-nav-list{
                    display: flex;
                    align-items: center;
                    margin-top: 20px;
                    span{
                        font-style: normal;
                        font-weight: 500;
                        font-size: 16px;
                        line-height: 24px;
                        color: #81858C;
                    }
                    a{
                        width: 16px;
                        height: 16px;
                        cursor: pointer;
                        margin-left: 12px;
                    }
                }
            }
            .shadow-nft-banner-right-box{
                img{
                    width: 404px;
                }
            }
        }
    }
</style>