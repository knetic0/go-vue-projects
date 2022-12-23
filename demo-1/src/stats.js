export const stats_1 = {
    type: "line",
    data: {
        labels: ["Roman", "Aşk", "Tarih", "Polisiye"],
        datasets: [
            {
                label: "Okuduğunuz Kitap Türleri",
                data: ["12", "8", "20", "4"],
                backgroundColor: "rgba(79, 65, 101, 0.8)",
                borderColor: "rgba(79, 65, 101, 0.8)",
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
}

export default { stats_1 }; 