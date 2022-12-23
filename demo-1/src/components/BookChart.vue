<template>
    <div class="row bg-white text-dark overflow-auto height" style="width: 1247px;">
        <div class="col-md-6 col-12 h-50 mt-2 rounded mb-5">
            <div class="mt-5 mb-5">
                <div class="col-10 mb-5 row-12">
                    <h3 class="text-center mb-3">Kitap Türleri & Popülariteleri</h3>
                    <canvas class="" id="book-chart-1"></canvas>
                </div>
            </div>
        </div>
        <div class="col-md-6 col-12 h-50 rounded mt-1 mb-5">
            <div class="ms-5 mt-5 mb-5">
                <div class="col-10 row-12 ms-2 mt-5 mb-5">
                    <h3 class="text-center">Kütüphanelerdeki Kitap Sayıları</h3>
                    <canvas class="ms-5 mb-4 mt-2 h-25" id="book-chart-2"></canvas>
                </div>
            </div>
        </div>
        <div class="col-md-6 col-12 rounded mt-5">
            <div class="mt-5">
                <div class="mt-5 col-10 row-12">
                    <h3 class="mt-5 text-center">Halk Kütüphanlerine Kayıtlı Üye Sayısı ve Yararlanan Kişi Sayısı</h3>
                    <canvas class="mt-5" id="book-chart-3"></canvas>
                </div>
            </div>
        </div>
        <div class="col-md-6 col-12 rounded mt-5">
            <div class="ms-5 mt-5">
                <div class="col-10 mt-5 row-12">
                    <h3 class="text-center mt-5">Halk Kütüphanlerinin Yıllara Göre Sayısı</h3>
                    <canvas class="mt-5" id="book-chart-4"></canvas>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
  import Chart from 'chart.js/auto'
  import { bookChartData_1, bookChartData_2, bookChartData_3, bookChartData_4 } from '../book-data'
  
  export default {
    name: 'BookChart',

    data() {
        return {
            bookChartData: [bookChartData_1, bookChartData_2, bookChartData_3, bookChartData_4],
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
    .height {
        height: 100vh;
    }
</style>