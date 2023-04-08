<template>
    <div class="price-card">
        <div class="price-card-item" v-for="item in list" :key="item.title">
            <div class="title">{{item.title}}</div>
            <div class="des">{{item.price}}</div>
            <div class="sub-des green">{{item.percent}}</div>
        </div>
    </div>
</template>

<script setup>
import { toRefs } from 'vue';

const props = defineProps({
  list: Array
});
const {list} = toRefs(props);

</script>

<style lang="scss" scoped>
  .price-card{
    background: #FFFFFF;
    border-radius: 24px;
    padding: 32px 24px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    &-item{
      .title{
        font-weight: 500;
        font-size: 16px;
        line-height: 24px;
        color: #81858C;
      }
      .des{
        font-weight: 700;
        font-size: 32px;
        line-height: 40px;
        color: #121214;
        margin: 12px 0 4px;
      }
      .sub-des{
        font-weight: 400;
        font-size: 14px;
        line-height: 22px;
        
      }
      .green{
        color: #24AE64;
      }
      .red{
        color: #EF454A;
      }
    }

  }
</style>