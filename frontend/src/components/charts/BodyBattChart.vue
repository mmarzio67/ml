<template>
  <div>
    <canvas id="bodybatt-chart"></canvas>
  </div>
</template>

<script>
import Chart from 'chart.js'
import BodyBattChartData from './bodybatt-data'

export default {
  name: 'BodyBattChart',
  data() {
    return {
      bodyBattChartData: BodyBattChartData
    }
  },
  mounted() {
    const graphData= this.$store.getters["getGraphData"]
    console.log("[BodyBattChart] :" + graphData)
    //console.log("[BodyBattChart] :" + this.bodyBattChartData.data.datasets[0])

    this.bodyBattChartData.data.labels=graphData[0]
    this.bodyBattChartData.data.datasets[0].data=graphData[1]
    this.bodyBattChartData.data.datasets[1].data=graphData[2]
    this.bodyBattChartData.data.datasets[2].data=graphData[3]

    const ctx = document.getElementById('bodybatt-chart');
    new Chart(ctx, this.bodyBattChartData);
  }
}
</script>