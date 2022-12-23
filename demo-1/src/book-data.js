export const bookChartData_1 = {
    type: "line",
    data: {
      labels: ["Roman", "Psikoloji", "Korku-Gerilim", "Aşk", "Tarih", "Kişisel Gelişim", "Polisiye", "Öykü"],
      datasets: [
        {
          label: "Popülarite",
          data: [417000, 416000, 158000, 598000, 231000, 210000, 164000, 130000],
          backgroundColor: "rgba(79, 65, 101, 0.8)",
          borderColor: "rgba(79, 65, 101, 0.8)",
          borderWidth: 3
        },
        {
          label: "Kitap Sayısı",
          data: [25400, 2100, 865, 3339, 11600, 3292, 2907, 9500],
          backgroundColor: "rgba(71, 183,132,.5)",
          borderColor: "#47b784",
          borderWidth: 3
        }
      ]
    },
    options: {
      responsive: true,
      lineTension: 1,
      scales: {
        yAxes: [
          {
            stacked: true,
            gridLines : {
              display: true,
            },
            ticks: {
              beginAtZero: true,
              padding: 25
            }
          }
        ],
        xAxes: [
          {
            gridLines: {
              display: true,
            }
          }
        ]
      }
    }
};

export const bookChartData_2 = {
  type: "pie",
  data: {
    labels: ["Milli Kütüphane", "Üniversite Kütüphaneleri", "Halk Kütüphaneleri", "Örgün ve Yaygın Eğitim Kütüphaneleri"],
    datasets: [
      {
        label: "Kütüphanelerdeki Kitap Sayıları (Milyon)",
        data: [1.5, 20.2, 22.4, 34.3],
        backgroundColor: ["blue", "red", "orange", "green"],
        borderColor: "#36495d",
        borderWidth: 3
      },
    ]
  },
  options: {
    responsive: true,
    lineTension: 1,
    scales: {
      yAxes: [
        {
          ticks: {
            beginAtZero: true,
            padding: 25
          }
        }
      ]
    }
  }
};

export const bookChartData_3 = {
  type: "bar",
  data: {
    labels: ["Halk Kütüphanelerine Kayıtlı Üye ve Yararlanan Kişi Sayısı (Milyon)"],
    datasets: [
      {
        label: "Kayıtlı Üye Sayısı (Milyon)",
        data: [4.9],
        backgroundColor: "rgba(79, 65, 101, 0.8)",
        borderColor: "rgba(79, 65, 101, 0.8)",
        borderWidth: 3
      },
      {
        label: "Yararlanan Kişi Sayısı (Milyon)",
        data: [15.7],
        backgroundColor: "rgba(71, 183,132,.5)",
        borderColor: "#47b784",
        borderWidth: 3
      }
    ]
  },
  options: {
    responsive: true,
    lineTension: 1,
    scales: {
      yAxes: [
        {
          ticks: {
            beginAtZero: true,
            padding: 25
          }
        }
      ]
    }
  }
};

export const bookChartData_4 = {
  type: "line",
  data: {
    labels: ["2016", "2017", "2018", "2019", "2020"],
    datasets: [
      {
        label: "Halk Kütüphaneleri",
        data: [1140, 1160, 1180, 1200, 1220],
        backgroundColor: "rgba(79, 65, 101, 0.8)",
        borderColor: "rgba(79, 65, 101, 0.8)",
        borderWidth: 3
      },
    ]
  },
  options: {
    responsive: true,
    lineTension: 1,
    scales: {
      yAxes: [
        {
          stacked: true,
          gridLines : {
            display: true,
          },
          ticks: {
            beginAtZero: true,
            padding: 25
          }
        }
      ],
      xAxes: [
        {
          gridLines: {
            display: true,
          }
        }
      ]
    }
  }
};
  
export default { bookChartData_1, bookChartData_2, bookChartData_3, bookChartData_4 };