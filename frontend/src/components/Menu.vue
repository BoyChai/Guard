<template>
  <div class="menu-container">
    <n-menu
      style="width: 300px; height: 40%"
      :options="menuOptions"
      :default-value="page"
    />
  </div>
</template>

<script>
import { defineComponent, h } from "vue";
import { NIcon } from "naive-ui";
import { RouterLink } from "vue-router";

import {
  SupervisedUserCircleRound as UserIcon,
  AddCardFilled as CardIcon,
  HomeOutlined as HomeIcon,
} from "@vicons/material";

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}
const menuOptions = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "home",
          },
        },
        { default: () => "首页" }
      ),

    key: "home",
    icon: renderIcon(HomeIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: "User",
          },
        },
        { default: () => "用户管理" }
      ),
    key: "user",
    icon: renderIcon(UserIcon),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: "Card",
          },
        },
        { default: () => "卡密管理" }
      ),
    key: "card",
    icon: renderIcon(CardIcon),
  },
];
export default defineComponent({
  props: {
    page: String,
  },
  setup() {
    return {
      menuOptions,
    };
  },
});
</script>

<style>
.menu-container {
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  height: 100vh; /* 使用视窗高度作为容器高度 */
}

/* 菜单项容器 */
.n-menu-item-content {
  display: flex; /* 使用 flex 布局 */
  align-items: center; /* 垂直居中 */
  gap: 10px; /* 图标和文字之间的间距 */
  padding: 8px 12px; /* 调整内边距 */
  border-radius: 8px; /* 圆角边框 */
  cursor: pointer; /* 鼠标悬停时显示指针 */
}

/* 图标样式 */
.n-menu-item-content-icon {
  width: 20px; /* 图标宽度 */
  height: 20px; /* 图标高度 */
}

/* 文字样式 */
.n-menu-item-content-label {
  flex: 1; /* 文字占据剩余空间 */
}
</style>
