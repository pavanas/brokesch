require("expose-loader?$!expose-loader?jQuery!jquery");

import Vue from "vue";
import VueRouter from "router";
Vue.use(VueRouter);

import HomeComponent from "./components/home.vue";
import RaceComponent from "./components/race.vue";
import Countdown from "./components/countdown.vue";

Vue.component('countdown', Countdown);

const routes = [
  {path: "/race/:id", component: RaceComponent, name: "showRace"},
  {path: "/", component: HomeComponent}
];

const router = new VueRouter({
  mode: "history",
  routes
});

const app = new Vue({
  router
}).$mount("#app");
