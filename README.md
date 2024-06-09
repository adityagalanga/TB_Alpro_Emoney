# E-Money Application

Aplikasi ini digunakan untuk mencatat saldo uang yang dimiliki dalam setiap akun. Pengguna aplikasi ini adalah admin e-money dan pemilik akun e-money.

## Spesifikasi

### 1. Registrasi Akun

- Pengguna dapat melakukan registrasi akun baru.
- Admin akan memverifikasi registrasi akun dengan opsi untuk menyetujui atau menolak pendaftaran.

### 2. Persetujuan/Penolakan Registrasi Akun oleh Admin

- Admin dapat melakukan persetujuan atau penolakan registrasi akun yang baru didaftarkan.
- Admin dapat mencetak daftar akun yang terdaftar.

### 3. Transaksi Uang

- Pengguna dapat mengirimkan uang ke akun lain.
- Pengguna dapat menerima kiriman uang dari akun lain.

### 4. Pembayaran

- Pengguna dapat melakukan pembayaran untuk berbagai keperluan seperti pembelian makanan, pulsa, listrik, BPJS, dan lain-lain.

### 5. Riwayat Transaksi

- Pengguna dapat mencetak riwayat transaksi mereka.

## Instalasi

1. Clone repositori ini:
    ```bash
    git clone https://github.com/adityagalanga/TB_Alpro_Emoney.git
    ```

2. Pindah ke direktori proyek:
    ```bash
    cd e-money-app
    ```

3. Install dependensi:
    ```bash
    go mod tidy
    ```

4. Jalankan aplikasi:
    ```bash
    go run main.go
    ```

## Penggunaan

Setelah aplikasi berjalan, pengguna dapat mengakses endpoint untuk melakukan berbagai operasi seperti registrasi, verifikasi, transaksi, pembayaran, dan pencetakan riwayat transaksi. Dokumentasi API lengkap tersedia di dalam repositori.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repositori ini dan buat pull request dengan perubahan yang Anda usulkan. Kami menyambut baik kontribusi dari siapa pun.

## version
v0.0.1