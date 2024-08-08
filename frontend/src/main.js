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

app.use(router);
app.mount("#app");
