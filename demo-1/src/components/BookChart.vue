<template>
    <div class="row w-100">
        <div class="col-6 w-50 h-50 bg-light bg-gradient rounded">
            <div class="ms-5 mt-2">
                <div class="col-10 row-12 ms-2">
                    <h1 class="text-center ms-5 mb-3 text-dark">Kitap Türleri & Popülariteleri</h1>
                    <canvas class="ms-5 h-25" id="book-chart-1"></canvas>
                </div>
            </div>
        </div>
        <div class="ms-2 col-6 w-50 h-50 bg-light bg-gradient rounded">
            <div class="ms-5 mt-2">
                <div class="col-10 row-12 ms-2">
                    <h1 class="text-center text-dark">Kütüphanelerdeki Kitap Sayıları</h1>
                    <canvas class="ms-5 mb-4 mt-2 h-25" id="book-chart-2"></canvas>
                </div>
            </div>
        </div>
        <div class="ms-2 col-6 w-50 h-50 bg-light bg-gradient rounded">
            <div class="ms-5 mt-2">
                <div class="col-10 row-12 ms-2">
                    <h1 class="text-center text-dark">Halk Kütüphanlerine Kayıtlı Üye Sayısı ve Yararlanan Kişi Sayısı</h1>
                    <canvas class="ms-5 mb-5 mt-5 h-25" id="book-chart-3"></canvas>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
  import Chart from 'chart.js/auto'
  import { bookChartData_1, bookChartData_2, bookChartData_3 } from '../book-data'
  
  export default {
    name: 'BookChart',

    data() {
        return {
            bookChartData: [bookChartData_1, bookChartData_2, bookChartData_3],
            chart : [],
        }
    },

    mounted() {
        var canvas_length = document.getElementsByTagName('canvas').length

        for (let i = 0; i < canvas_length; i++) {
            this.chart.push(`book-chart-${i+1}`)
        }

        for (let i = 0; i < this.chart.length; i++) {
            new Chart(document.getElementById(this.chart[i]), this.bookChartData[i])
        }
    }
}
</script>

<style scoped>

</style>