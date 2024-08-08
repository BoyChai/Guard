import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/User",
    name: "user",
    component: () => import("../views/UserView.vue"),
  },
  {
    path: "/Card",
    name: "card",
    component: () => import("../views/CardView.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
