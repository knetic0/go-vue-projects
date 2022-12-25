<template>
  <div class="bg-white text-dark overflow-auto" style="width : 100%; height: 100vh">
    <div class="col-md-4 col-12 w-50 h-50 rounded">
      <div class="ms-5 mt-2">
        <div class="col-12 row-12 ms-2 mt-4">
          <canvas class="ms-5 mb-4 mt-2 h-25" id="myChart"></canvas>
        </div>
      </div>
    </div>
    <div class="col-md-4 col-12 w-50 h-50 rounded">
      <div class="ms-5">
        <div class="col-12 row-12 ms-2">
          <canvas class="ms-5" id="myChart-2"></canvas>
        </div>
      </div>
    </div>
    <div class="col-md-4 col-12 w-50 h-50 rounded">
      <div class="ms-5">
        <div class="col-12 row-12 ms-2">
          <canvas class="ms-5" id="myChart-3"></canvas>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Chart from "chart.js/auto";

export default {
  name: "Dashboard",

  data() {
    return {
      socket: "",
      books: [],
      books_type: [],
      popularity: [],
      total_book: [],
      suggestion_ratio : [],
    };
  },

  created() {
    this.socket = new WebSocket("ws://localhost:9100/dashboard");
    this.socket.onmessage = (event) => {
      var message = JSON.parse(event.data);
      console.log(message)
      for (let i = 0; i < message.length; i++) {
        this.books_type.push(message[i]["BookType"]);
        this.popularity.push(message[i]["Popularity"]);
        this.total_book.push(message[i]["TotalBook"]);
        this.suggestion_ratio.push(message[i].popularity % message[i].totalbook)
      }
      console.log(this.books_type);
    };
  },

  mounted() {
    setTimeout(() => {
      const ctx = document.getElementById("myChart").getContext('2d');
      const ctx_2 = document.getElementById("myChart-2").getContext('2d');
      const ctx_3 = document.getElementById("myChart-3").getContext('2d')

      new Chart(ctx, {
        type: "bar",
        data: {
          labels: this.books_type,
          datasets: [
            {
              label: "Kitap Türleri ve Popülariteleri",
              data: this.popularity,
              backgroundColor: [
                "blue",
                "red",
                "orange",
                "purple",
                "pink",
                "yellow",
                "darkblue",
                "green",
              ],
              borderWidth: 1,
            },
          ],
        },
        options: {
          responsive: true,
          scales: {
            y: {
              beginAtZero: true,
            }
          },
        },
      });
      new Chart(ctx_2, {
        type: "line",
        data: {
          labels: this.books_type,
          datasets: [
            {
              label: "Toplam Kitap Sayıları",
              backgroundColor: "rgba(126, 6, 47, 0.8)",
              borderColor: "rgba(126, 6, 47, 0.8)",
              data: this.total_book,
              borderWidth: 1,
            },
          ],
        },
        options: {
          responsive: true,
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
      new Chart(ctx_3, {
        type : "polarArea",
        data : {
          labels : this.books_type,
          datasets : [
            {
              label : "Tavsiye Eden Okuyucu Sayısı",
              data : this.suggestion_ratio,
              backgroundColor : [
                 "rgba(2, 192, 240, 0.8)",
                 "rgba(255, 178, 0, 0.8)",
                 "rgba(98, 255, 0, 0.8)",
                 "rgba(255, 0, 0, 0.8)",
                 "rgba(168, 0, 255, 0.8)",
                 "rgba(0, 39, 255, 0.8)",
                 "rgba(204, 255, 0, 0.8)",
                 "rgba(99, 0, 255, 0.8)",
              ],
            }
          ],
          options : {
            plugins : {
              legend : false,
              tooltip : false,
            }
          }
        }

      });
    }, 1000);
  },
};
</script>

<style scoped>
.height {
  height: 100vh;
}
</style>
