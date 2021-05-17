import globalAxios from "axios";
import store from "../../store";
import router from "../../router";

export default {
  state: {
    lastHealthEntry: [],
    allHealthEntries:[]
  },
  mutations: {
    SET_LASTRECORD(state, lr) {
        state.lastHealthEntry = lr;
    },
    SET_ALLHEALTHRECORDS(state, ahrs) {
        state.allHealthEntries = ahrs;
    },
  },
  actions: {
    dailyhealthmon(dailymonData) {
      const regDate = new Date();
      if (!store.state.auth.idToken) {
        return;
      }
      console.log("[state activity.js]:", store.state.auth.idToken);
      globalAxios
        .post("/dailyhealthmon.json" + "?auth=" + store.state.auth.idToken, {
          weight: dailymonData.weight,
          bfi: dailymonData.bfi,
          imc: dailymonData.imc,
          waist: dailymonData.waist,
          spo2: dailymonData.spo2,
          breathrest: dailymonData.breathrest,
          breathactive: dailymonData.breathactive,
          hourssleep: dailymonData.timeslept,
          pulserest: dailymonData.pulserest,
          pulseactive: dailymonData.pulseactive,
          stress: dailymonData.stress,
          bodybattU: dailymonData.bodybattU,
          bodybattD: dailymonData.bodybattD,
          steps: dailymonData.steps,
          registerDate: regDate,
          userId: store.state.auth.userId,
        })
        .then((res) => {
          console.log(res);
          router.replace("/dashboard");
        })
        .catch((error) => console.log(error));
    },
    lastHealthEntry({commit}) {
      console.log("[action:lastHealthEntry]")
      // use the allHealthEntries state instead to call all the records again
      globalAxios
        //.get('/dailyhealthmon.json?orderBy="bodybatt"&auth=' + store.state.auth.idToken
        .get('/dailyhealthmon.json' + '?auth=' + store.state.auth.idToken
        ).then((res) => {   
          console.log(res);
          const data = res.data;
          const hrs = [];
          for (let key in data) {
            const hr = data[key];
            hr.id = key;
            hrs.push(hr);
          }
          console.log(hrs);
          // sort by registerDate ascending
          const lhr= hrs.sort(function(a, b){
            return a.registerDate - b.registerDate;
          });
          // count the number of records
          const nr=lhr.length
          console.log(nr)
          // update the state with the newest record
          commit("SET_LASTRECORD", lhr[nr-1]);         
      })
      .catch((error) => console.log(error));
    },
    allHealthEntries({commit}) {
      console.log("[action: allHealthEntries]")
      globalAxios
        .get('/dailyhealthmon.json' + '?auth=' + store.state.auth.idToken
        ).then((res) => {   
          //console.log(res);
          const data = res.data;
          const hrs = [];
          for (let key in data) {
            const hr = data[key];
            hr.id = key;
            hrs.push(hr);
          }
          // sort by registerDate ascending
          const ahrs= hrs.sort(function(a, b){
            return a.registerDate - b.registerDate;
          });
          console.log("[allHealthEntries]:" + JSON.stringify(ahrs))
          commit("SET_ALLHEALTHRECORDS", ahrs);         
      })
      .catch((error) => console.log(error));
    }
  },
  getters: {
    getLastHealthEntry: (state) => {
      return state.lastHealthEntry;
    },
    getAllHealthEntries: (state) => {
      return state.allHealthEntries;
    },
  }
}
