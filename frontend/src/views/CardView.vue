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
        <Menu page="card" style="height: 100%" />
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
                创建卡密
              </n-button>
            </div>
          </div>
        </template>
      </n-split>
    </template>
  </n-split>
  <n-drawer v-model:show="createStatus" :width="502">
    <n-drawer-content>
      <template #header> <h3>创建卡密</h3> </template>
      <n-input v-model:value="key" placeholder="Key,留空则自动生成" />

      <n-date-picker
        style="margin-top: 20px"
        v-model:value="datetime"
        type="datetime"
      />

      <div style="margin-top: 50px">
        <n-button @click="createStatus = false">取消</n-button>
        <n-button style="margin-left: 20px" @click="createCard">创建</n-button>
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
  { title: "卡密", key: "Key" },
  { title: "创建时间", key: "CreatedAt" },
  { title: "过期时间", key: "EndDate" },
  { title: "更新时间", key: "UpdatedAt" },
  {
    title: "操作",
    key: "actions",
    render(row) {
      return h(
        NButton,
        {
          size: "small",
          onClick: () => deleteCardWarn(row),
        },
        { default: () => "删除" }
      );
    },
  },
]);
const deleteCard = (row) => {
  axios
    .delete("/api/card/delete", {
      data: {
        id: row.ID,
      },
    })
    .then((res) => {
      notify("success", "信息", "删除" + row.Key + "卡密成功");
      getCardList();
    })
    .catch((err) => {
      notify("error", "错误", err.request.response);
    });
};
// 表数据
const data = ref([]);

// 获取所有用户
const getCardList = () => {
  axios
    .get("/api/card/getList")
    .then((res) => {
      console.log(res);
      // 遍历数组，处理 EndDate 字段
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
getCardList();
const dialog = useDialog();

const key = ref("");
const datetime = ref(undefined);

const createCard = () => {
  axios
    .post("/api/card/create", {
      key: key.value,
      time: datetime.value / 1000,
    })
    .then((res) => {
      notify("success", "信息", "创建卡密成功");
      createStatus.value = false;
      key.value = "";
      datetime.value = 0;
      getCardList();
    })
    .catch((err) => {
      notify("error", "错误", err.request.response);
    });
};
// 删除提示
const deleteCardWarn = (row) => {
  dialog.warning({
    title: "警告",
    content: "你确定？",
    positiveText: "确定",
    negativeText: "不确定",
    onPositiveClick: () => {
      deleteCard(row);
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
#create {
  height: 40%;
  width: 40%;
  border: 1px solid black;
  background-color: aliceblue;
  position: absolute;
  top: 20%;
  left: 30%;
}
</style>
