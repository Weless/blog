import Vue from "vue";
import Vuex from "vuex";
import userModule from "./module/user";

Vue.use(Vuex);

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== "production", // 开启严格模式
  state: {},
  mutations: {},
  actions: {},
  modules: {
    userModule
  }
});
