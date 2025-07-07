# srv_mqtt_demo

`srv_mqtt_demo` adalah aplikasi Go yang menghubungkan ke broker MQTT, mempublikasikan data sensor secara berkala, serta menyimpan data yang diterima ke MongoDB. Data sensor juga dapat diakses melalui endpoint HTTP berbasis Fiber.

---

## 🚀 Fitur

- Menghasilkan dan mengirim data sensor (suhu air, suhu, pH, PPM) ke topik MQTT secara periodik.
- Menerima pesan dari MQTT dan menyimpan ke MongoDB.
- API endpoint untuk mengambil seluruh data sensor dari database.
- Menggunakan TLS dan autentikasi saat koneksi ke broker MQTT.

---

## 🛠️ Teknologi

- [Go (Golang)](https://golang.org/)
- [Fiber (Web Framework)](https://gofiber.io/)
- [Eclipse Paho MQTT](https://github.com/eclipse/paho.mqtt.golang)
- [MongoDB](https://www.mongodb.com/)
- MQTT Broker: HiveMQ (atau yang kamu pilih)

---

## 📦 Struktur Proyek

```mqtt-demo/
├── main.go # Entry point aplikasi
├── main/Types # Struct & tipe data
│ └── sensor.go
├── main/constant # Konstanta (MQTT URL, username, password)
├── main/database # Inisialisasi koneksi MongoDB
├── main/handlers # Handler untuk simpan & ambil data
└── go.mod / go.sum # Dependency
```

## 🧪 Menjalankan Aplikasi
- Pastikan MongoDB aktif di localhost (default port 27017).
- Jalankan aplikasi:

```
go run main.go

```

## ⚙️ Konfigurasi MQTT

Pastikan kamu mengatur variabel berikut di file `main/constant/constant.go`:

```go
const (
	MQTT_URL      = "your-broker-url"
	MQTT_USERNAME = "your-username"
	MQTT_PASSWORD = "your-password"
)
```

🌐 Endpoint HTTP
GET /sensor
Mengembalikan seluruh data sensor dari MongoDB dalam format JSON.

Contoh response:
```
[
  {
    "senorSuhuAir": "27.5",
    "senorSuhu": "28.1",
    "sensorPPM": "400",
    "sensorPh": "6.8",
    "lastUpdate": "2025-07-07T20:20:34+07:00"
  }
]
```

🧪 Unit Test
Untuk menjalankan unit test:

```
go test ./...
```
