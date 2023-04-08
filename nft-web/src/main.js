import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

// 引入ElementPlus
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';

import * as echarts from 'echarts';

const app = createApp(App);

app.config.globalProperties.$echarts = echarts; // 全局挂载echarts

app.use(ElementPlus); // 挂载ElementPlus

app.use(store).use(router).mount('#app');
