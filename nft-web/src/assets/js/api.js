import axios from 'axios';

/* 添加在页面中切换接口环境 start */
const NODE_ENV =
  process.env.NODE_ENV !== 'production' ? 'dev' : 'prod',

  baseURL = {
    dev: '/dev',
    prod: window.location.origin || 'https://www.gomefinance.com.cn'
  }[NODE_ENV],

  // 创建新的axios实例
  httpRequest = axios.create({
    baseURL, // 默认路径
    timeout: 60000 // 默认超时时间
  });

const customErrorHandler = ({ data, config })=> {
  notification.error({
    message: `请求错误 (Code: ${data.code})`,
    description: `${data.msg} (${config.url})`,
    className: 'break-all'
  });
};

httpRequest.interceptors.response.use(
  async (response) => {
    if (response.data.code === 2000) {
      return Promise.resolve(response.data.data);
    }
    customErrorHandler(response);
    return Promise.reject(response);
  },
  (error) => {
    // 过滤主动取消请求导致的错误
    if (!isCancel(error)) {
      standardErrorHandler(error);
    }

    return Promise.reject(error);
  }
);

export default httpRequest;