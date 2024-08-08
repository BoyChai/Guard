<template>
  <div class="login_wrapper">
    <div class="login_content">
      <h1>Guard-卡密校验</h1>
      <n-input
        class="fillet input"
        v-model:value="name"
        type="text"
        placeholder="账号"
      />
      <n-input
        class="fillet input"
        v-model:value="pass"
        style="margin-top: 12px"
        type="password"
        show-password-on="mousedown"
        placeholder="密码"
      />
      <n-button class="btn" type="info" style="margin-top: 20px" @click="login">
        登入
      </n-button>
    </div>
  </div>
</template>

<script setup>
import { ref, getCurrentInstance } from "vue";
import { useNotification } from "naive-ui";
import router from "@/router";

// 引入axios
const { proxy } = getCurrentInstance();
const axios = proxy.$axios;

const name = ref("");
const pass = ref("");
const login = () => {
  var raw = JSON.stringify({
    name: name.value,
    pass: pass.value,
  });

  axios
    .post("/api/user/login", raw)
    .then((res) => {
      notify("success", "信息", "登录成功");
      localStorage.setItem("jwtToken", res.data);
      router.push("/");
    })
    .catch((err) => {
      notify("error", "错误", err.request.response);
      name.value = "";
      pass.value = "";
    });
};
const notification = useNotification();
const notify = (type, title, text) => {
  notification[type]({
    content: title,
    meta: text,
    duration: 2500,
    keepAliveOnHover: true,
  });
};
</script>

<style>
.login_wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  margin-top: -5%;
}

.login_content {
  width: 20%;
  height: 19%;
  text-align: center;
  background-color: #f0f0f0;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.fillet {
  border-radius: 6px;
}
.input {
  text-align: left;
  height: 45px;
  font-size: 20px;
  line-height: 45px;
}
.btn {
  height: 40px;
  width: 100%;
}
</style>
