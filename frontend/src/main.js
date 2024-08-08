import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

// 通用字体
import "vfonts/Lato.css";
// 等宽字体
import "vfonts/FiraCode.css";

// naive-ui
import naive from "naive-ui";
const app = createApp(App);
app.use(naive);

// axios
import axios from "axios";
axios.defaults.baseURL = "http://localhost:7779";
// axios请求拦截
// 请求拦截
axios.interceptors.request.use((config) => {
  let token = localStorage.getItem("jwtToken");
  if (token) {
    config.headers["Authorization"] = token;
  }
  return config;
});
// 响应拦截
axios.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    if (error.response.code === 401) {
      router.push("/Login");
    }
    return Promise.reject(error);
  }
);
app.config.globalProperties.$axios = axios;

app.use(router);
app.mount("#app");
