import Vue from "vue";
import Vuex from "vuex";
import auth from "./state/modules/auth";
import activity from "./state/modules/activity";
import charts from "./state/modules/charts";

Vue.use(Vuex);

const Store = new Vuex.Store({
  modules: {
    auth,
    activity,
    charts
  },

  actions: {
    setLogoutTimer({ commit }, expirationTime) {
      setTimeout(() => {
        commit("clearAuthData");
      }, expirationTime * 1000);
    },
  },
});

export default Store;
