<template>
  <n-split
    direction="horizontal"
    style="height: 100%; width: 100%"
    default-size="22%"
  >
    <template #1>
      <div style="margin-top: 10px; margin-bottom: 10px">
        <h1>Guard-卡密管理系统</h1>
      </div>
      <div>
        <Menu page="user" style="height: 100%" />
      </div>
    </template>
    <template #2>
      <n-split direction="vertical" style="height: 100%" default-size="8%">
        <template #1> <div id="header"></div> </template>
        <template #2>
          <div id="main">
            <n-data-table
              :bordered="false"
              :single-line="false"
              :columns="columns"
              :data="data"
              :pagination="pagination"
            />
            <div style="margin-top: 100px; margin-right: 10%; float: right">
              <n-button type="primary" @click="createStatus = true">
                创建用户
              </n-button>
            </div>
          </div>
        </template>
      </n-split>
    </template>
  </n-split>
  <n-drawer v-model:show="createStatus" :width="502">
    <n-drawer-content>
      <template #header> <h3>创建用户</h3> </template>
      <n-input v-model:value="name" placeholder="用户名" />
      <n-input
        style="margin-top: 20px"
        type="password"
        show-password-on="mousedown"
        v-model:value="password"
        placeholder="密码"
      />
      <div style="margin-top: 50px">
        <n-button @click="createStatus = false">取消</n-button>
        <n-button style="margin-left: 20px" @click="createUser">创建</n-button>
      </div>
    </n-drawer-content>
  </n-drawer>
</template>

<script setup>
import Menu from "@/components/Menu.vue";
import { ref, getCurrentInstance, h } from "vue";
import { NButton, useNotification, useDialog } from "naive-ui";

const createStatus = ref(false);

// 引入axios
const { proxy } = getCurrentInstance();
const axios = proxy.$axios;

// 列名称
const columns = ref([
  { title: "用户名", key: "Name" },
  { title: "创建时间", key: "CreatedAt" },
  {
    title: "操作",
    key: "actions",
    render(row) {
      return h(
        NButton,
        {
          size: "small",
          onClick: () => deleteUserWarn(row),
        },
        { default: () => "删除" }
      );
    },
  },
]);
const deleteUser = (row) => {
  axios
    .delete("/api/user/delete", {
      data: {
        id: row.ID,
      },
    })
    .then((res) => {
      notify("success", "信息", "删除" + row.Name + "用户成功");
      getUserList();
    })
    .catch((err) => {
      notify("error", "错误", err.request.response);
    });
};
// 表数据
const data = ref([]);

// 获取所有用户
const getUserList = () => {
  axios
    .get("/api/user/getList")
    .then((res) => {
      console.log(res);
      res.data.forEach((item) => {
        // 使用 toISOString() 将日期转换为 ISO 8601 格式字符串，然后截取所需部分
        if (item.EndDate) {
          item.EndDate = new Date(item.EndDate).toISOString().substring(0, 19); // 截取到秒的部分
        }
        if (item.CreatedAt) {
          item.CreatedAt = new Date(item.CreatedAt)
            .toISOString()
            .substring(0, 19); // 截取到秒的部分
        }
        if (item.UpdatedAt) {
          item.UpdatedAt = new Date(item.UpdatedAt)
            .toISOString()
            .substring(0, 19); // 截取到秒的部分
        }
      });
      data.value = res.data;
    })
    .catch((err) => {
      console.log(err);
    });
};
getUserList();

const name = ref("");
const password = ref("");
const createUser = () => {
  axios
    .post("/api/user/create", {
      name: name.value,
      pass: password.value,
    })
    .then((res) => {
      notify("success", "信息", "创建用户成功");
      createStatus.value = false;
      name.value = "";
      password.value = "";
      getUserList();
    })
    .catch((err) => {
      notify("error", "错误", err.request.response);
    });
};

const dialog = useDialog();

// 删除提示
const deleteUserWarn = (row) => {
  dialog.warning({
    title: "警告",
    content: "你确定？",
    positiveText: "确定",
    negativeText: "不确定",
    onPositiveClick: () => {
      deleteUser(row);
    },
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
#header {
  /* background-color: aqua; */
  height: 100%;
  width: 100%;
}
#main {
  /* background-color: antiquewhite; */
  height: 100%;
  width: 100%;
}
/* #create {
  height: 40%;
  width: 40%;
  border: 1px solid black;
  background-color: aliceblue;
  position: absolute;
  top: 20%;
  left: 30%;
} */
</style>
