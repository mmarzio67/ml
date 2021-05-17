<template>
  <div id="dashboard">
    <h1>That's the Health Dashboard!</h1>
    <p>You are here because you are authenticated with:</p>
    <p v-if="email">email address: {{ email }}</p>
    <div>
      <BodyBattChart />
    </div>
  </div>
</template>

<script>

//import PlanetChart from '../charts/PlanetChart.vue'
import BodyBattChart from '../charts/BodyBattChart.vue'
export default {
  components: {
    BodyBattChart
  },
  data() {
    return {
      bodybatt: null
    }  
  },
  computed: {
    email() {
      return !this.$store.getters.user ? false : this.$store.getters.user.email;
    },
  },
  created() {
    this.$store.dispatch("fetchUser"); 
  },
  mounted() {
    this.setChartHealthTrend()
  },

  methods: {
    setChartHealthTrend(){
      this.$store.dispatch("prepHealthGraph")
      
      /*
      const ahr = this.$store.getters["getAllHealthEntries"];
      console.log("[setChartHealthTrend, register date]: " + ahr)
      this.bodybatt= ahr[0].bodybatt
      console.log("[bodybatt-data] normal: " + ahr)
      // reduce the ahr fields (rahr) to registerDate and BodyBatt fields
      const rahr = ahr.map(( {registerDate, bodybatt} ) => ({registerDate, bodybatt}))
      console.log("[bodybatt-data] reduced: " + rahr[0])
      */
    }
  }
};
</script>

<style scoped>
h1,
p {
  text-align: center;
}

p {
  color: red;
}
</style>
