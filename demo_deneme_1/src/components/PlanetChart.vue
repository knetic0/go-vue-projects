<template>
    <div class="d-inline-flex flex-wrap ms-5 justify-content-center">
        <div class="ms-2 mt-5 w-50 h-50 bg-light bg-gradient rounded">
            <div class="ms-5 mt-2">
                <div class="col-10 row-12 ms-2">
                    <h1 class="text-center text-dark">About Planets</h1>
                    <canvas class="ms-5 h-25" id="planet-chart-1"></canvas>
                </div>
            </div>
        </div>
        <div class="ms-2 mt-5 w-50 h-50 bg-light bg-gradient rounded">
            <div class="ms-5 mt-2">
                <div class="col-10 row-12 ms-2">
                    <h1 class="text-center text-dark">About Planets</h1>
                    <canvas class="ms-5 h-25" id="planet-chart-2"></canvas>
                </div>
            </div>
        </div>
        <div class="ms-2 mt-5 w-50 h-50 bg-light bg-gradient rounded">
            <div class="ms-5 mt-2">
                <div class="col-10 row-12 ms-2">
                    <h1 class="text-center text-dark">About Planets</h1>
                    <canvas class="ms-5 h-25" id="planet-chart-3"></canvas>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
  import Chart from 'chart.js/auto'
  import planetChartData from '../planet-data'
  
  export default {
    name: 'PlanetChart',

    data() {
        return {
            planetChartData: planetChartData,
            chart : [],
        }
    },

    mounted() {
        var canvas_length = document.getElementsByTagName('canvas').length

        for (let i = 0; i < canvas_length; i++) {
            this.chart.push(`planet-chart-${i+1}`)
        }

        for (let i = 0; i < this.chart.length; i++) {
            new Chart(document.getElementById(this.chart[i]), this.planetChartData)
        }
    }
}
</script>