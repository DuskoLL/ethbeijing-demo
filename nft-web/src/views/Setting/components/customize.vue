<template>
  <div class="customize-container">
    <div class="customize-box">
      <h6 class="customize-title">Customize My Shadow Score</h6>
      <div class="customize-content-box">
        <el-radio-group v-model="radio">
          <el-radio :label="3">
            <slot>
              <p class="customize-content-desc">
                <span>稳健型</span>80%的人选择
              </p>
              <p class="customize-content-desc">
                稳健型的方案会怎么调整，解释一下
              </p>
            </slot>
          </el-radio>
          <el-radio :label="6">
            <slot>
              <p class="customize-content-desc">
                <span>进取型</span>12%的人选择
              </p>
              <p class="customize-content-desc">
                进取型的方案会怎么调整，解释一下
              </p>
            </slot></el-radio
          >
          <el-radio :label="9">
            <slot>
              <p class="customize-content-desc">
                <span>保守型</span>8%的人选择
              </p>
              <p class="customize-content-desc">
                保守型的方案会怎么调整，解释一下
              </p>
            </slot>
          </el-radio>
        </el-radio-group>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const radio = ref(3);

</script>

<style lang="scss" scoped>
.customize-container {
  width: 100%;
  .customize-box {
    width: 1360px;
    padding: 32px 24px;
    background: #ffffff;
    border-radius: 24px;
    margin: 32px auto;
    box-sizing: border-box;
    .customize-title {
      font-family: "IBM Plex Sans";
      font-style: normal;
      font-weight: 700;
      font-size: 32px;
      line-height: 40px;
      /* identical to box height, or 125% */
      text-transform: capitalize;
      /* Gray [T]/T1_Title */
      color: #121214;
    }
    .customize-content-box {
        :deep(.el-radio-group){
            flex-direction: column;
            .el-radio{
                margin-right: 0;
                margin: 40px 0;
            }
        }
      .customize-content-desc {
        /* Regular/16 */
        font-family: "IBM Plex Sans";
        font-style: normal;
        font-weight: 400;
        font-size: 16px;
        line-height: 24px;
        /* Gray [T]/T2 */
        color: #81858c;
        span {
          /* Semibold/20 */
          font-family: "IBM Plex Sans";
          font-style: normal;
          font-weight: 600;
          font-size: 20px;
          line-height: 28px;
          /* identical to box height, or 140% */
          /* Gray [T]/T1_Title */
          color: #121214;
          margin-right: 5px;
        }
      }
    }
  }
}
</style>