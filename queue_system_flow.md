# Queue System Flow (Go + Supervisor)

Dokumentasi ini menjelaskan arsitektur dan flow queue system yang akan kita bangun di project Go API.  
Tujuannya agar alur async job (email, OTP, notification, dll) tetap jelas saat implementasi, debugging, maupun scaling.

---

# Tujuan Queue System

Queue digunakan untuk memproses pekerjaan asynchronous (background jobs) agar request API tetap cepat dan tidak blocking.

Contoh kasus:

- kirim OTP email
- kirim welcome email
- kirim notifikasi
- generate report
- resize image
- cleanup data
- retry failed job

Tanpa queue, semua proses berat dijalankan langsung saat request masuk.

Contoh:

```text
POST /register
   └── create user
   └── create otp
   └── send email
   └── response
```

Masalahnya:

- request lambat
- rawan timeout
- SMTP lambat = API ikut lambat
- retry susah
- gagal email bisa merusak flow utama

Queue memisahkan proses berat dari request utama.

---

# Konsep Dasar Queue

Queue system dibagi menjadi 3 komponen utama:

1. **Producer**
2. **Worker**
3. **Supervisor**

---

# 1. Producer

Producer adalah bagian yang membuat job dan memasukkannya ke queue.

Dalam konteks project ini, producer biasanya berasal dari API.

Contoh:

- user register
- user request forgot password
- system trigger notification

Producer tidak mengerjakan job langsung.  
Producer hanya menyimpan job ke queue.

Contoh flow:

```text
POST /register
   └── create user
   └── create otp
   └── insert job(send_otp_email)
   └── return success
```

Jadi API tetap cepat karena hanya insert data job.

---

# 2. Worker

Worker adalah process background yang terus berjalan untuk mengambil job dari queue lalu mengeksekusinya.

Worker tidak menerima HTTP request.

Worker hanya:

- ambil job
- proses job
- tandai selesai / gagal
- ulangi lagi

Flow worker:

```text
loop forever:
   ambil job
   proses
   mark success / fail
   ulang
```

Contoh:

```text
ambil job: send_otp_email
   └── parse payload
   └── send email
   └── mark done
```

Worker adalah inti queue processing.

---

# 3. Supervisor

Supervisor adalah process manager untuk menjaga worker tetap hidup.

Supervisor **bukan queue engine**.  
Supervisor **tidak memproses job**.

Supervisor hanya bertugas:

- menjalankan worker
- restart worker jika crash
- start otomatis saat server boot
- simpan log worker
- monitor status worker

Jadi Supervisor hanya “mandor” process worker.

---

# Arsitektur Queue

```text
Client
   ↓
Gin API
   └── insert job ke database

Queue Worker
   └── ambil job dari database
   └── eksekusi job

Supervisor
   └── menjaga Queue Worker tetap hidup
```

---

# Responsibility Tiap Layer

## API (Producer)

Tugas API:

- menerima request
- validasi request
- simpan data utama
- insert job ke queue
- return response

API tidak boleh menjalankan job berat secara langsung.

---

## Queue Worker

Tugas worker:

- polling job
- lock job
- eksekusi handler
- retry jika gagal
- mark selesai

Worker fokus ke background processing.

---

## Supervisor

Tugas supervisor:

- start worker
- restart worker jika panic / crash
- auto start saat reboot
- monitoring status
- manage log worker

Supervisor tidak tahu business logic queue.

---

# Flow Register + OTP (Async)

## Tanpa Queue

```text
POST /register
   └── create user
   └── create otp
   └── send OTP email
   └── response
```

Masalah:

- response lambat
- SMTP delay ikut memperlambat request
- kalau email gagal, flow kacau

---

## Dengan Queue

```text
POST /register
   └── create user
   └── create otp
   └── insert job(send_otp_email)
   └── response cepat
```

Worker:

```text
Worker loop
   └── ambil send_otp_email
   └── kirim email
   └── mark done
```

Hasil:

- API cepat
- email async
- retry mudah
- failure terisolasi

---

# Kenapa Queue Penting

Queue menyelesaikan beberapa masalah besar:

- request tidak blocking
- job berat dipindah ke background
- retry lebih mudah
- failure tidak merusak request utama
- scalable (bisa tambah worker)
- observability lebih baik
- log job lebih rapi

---

# Kenapa Worker Harus Process Terpisah

Worker tidak boleh digabung dengan HTTP server.

Kenapa:

- lifecycle berbeda
- failure domain berbeda
- scaling berbeda
- resource usage berbeda

API dan worker harus dipisah.

Contoh executable:

```text
cmd/api/main.go
cmd/worker/main.go
```

---

# Kenapa Supervisor Cocok

Worker adalah long-running process.

Pattern worker:

```text
for {
   ambil job
   proses
   sleep
}
```

Process seperti ini cocok dijaga Supervisor karena:

- long-running
- bisa crash
- perlu auto restart
- perlu auto boot
- perlu logging terpisah

Ini use case ideal Supervisor.

---

# Supervisor Flow

```text
Supervisor start
   └── jalankan worker

Worker panic / crash
   └── Supervisor detect
   └── restart worker

Server reboot
   └── Supervisor auto start worker
```

Jadi worker tetap hidup tanpa intervensi manual.

---

# Benefit Arsitektur Ini

Dengan arsitektur ini kita dapat:

- API cepat
- worker stabil
- job async
- retryable
- scalable
- fault isolated
- mudah dimonitor
- production friendly

---

# Prinsip Penting

1. API hanya produce job  
2. Worker hanya consume job  
3. Supervisor hanya jaga process  

Jangan campur responsibility.

Kalau dipisah rapi, sistem lebih mudah:

- dibangun
- di-debug
- di-maintain
- di-scale

---

# Ringkasan

Queue system kita akan dibangun dengan pola:

- **Gin API** → producer
- **Go Worker** → consumer
- **Supervisor** → process manager

Flow utamanya:

```text
Request masuk
   └── API simpan job
   └── response cepat

Worker ambil job
   └── proses async
   └── mark selesai

Supervisor jaga worker
   └── restart jika crash
```

Ini adalah fondasi queue system yang sehat, scalable, dan aman untuk production.

