<template>
    <div class="bg-white w-100 text-dark justify-content-center">
        <div class="ms-5">
            <div class="ms-5">
                <div class="ms-5">
                    <h1 class="ms-5 text-center">Sayın Mehmet Solak, Verileriniz Aşağıdaki Gibidir: </h1>
                </div>
                <div class="ms-5 w-50 h-50 mt-5">
                    <canvas id="stats-1" class="ms-5"></canvas>
                    <canvas id="stats-2" class="ms-5 mt-5"></canvas>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import { Chart } from "chart.js/auto"
    import { stats_1 } from "../stats"

    export default {
        name : "Stats",

        data() {
            return {
                stats : [stats_1],
                canvas : []
            }
        },

        mounted() {
            var len = document.getElementsByTagName("canvas").length
            for (let i = 0; i < len; i++) {
                this.canvas.push(`stats-${i+1}`)
            }

            for(let i =0; i < this.canvas.length; i++) {
                new Chart(document.getElementById(this.canvas[i]), this.stats[0])
            }
        }
    }
</script>

