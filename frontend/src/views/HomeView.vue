<template>
  <!-- <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div> -->
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
        <Menu page="home" style="height: 100%" />
      </div>
    </template>
    <template #2>
      <n-split direction="vertical" style="height: 100%" default-size="8%">
        <template #1> <div id="header"></div> </template>
        <template #2> <div id="main"></div> </template>
      </n-split>
    </template>
  </n-split>
</template>

<script>
import Menu from "@/components/Menu.vue";

export default {
  name: "HomeView",
  components: {
    Menu,
  },
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
</style>
