**Insert 1 Juta Data dari File CSV Ke Database Server, Menggunakan Teknik Worker Pool, Database Connection Pool, dan Mekanisme Failover.**

- Pada chapter ini kita akan praktek penerapan salah satu teknik concurrent programming di Go yaitu worker pool, dikombinasikan dengan database connection pool, untuk membaca 1 juta rows data dari sebuah file csv untuk kemudian di-insert-kan ke mysql server.
- Pada bagian insert data kita terapkan mekanisme failover, jadi ketika ada operasi insert gagal, maka akan otomatis di recover dan di retry. Jadi idealnya di akhir, semua data, sejumlah satu juta, akan berhasil di-insert.
1. Worker Pool
   - Worker pool adalah teknik manajemen goroutine dalam concurrent programming pada Go. Sejumlah worker dijalankan dan masing-masing memiliki tugas yang sama yaitu menyelesaikan sejumlah jobs.
   - Dengan metode worker pool ini, maka penggunaan memory dan performansi program akan bisa optimal.
2. Database Connection Pool
   - Connection pool adalah metode untuk manajemen sejumlah koneksi database, agar bisa digunakan secara optimal.
   - Connection pool sangat penting dalam kasus operasi data yang berhubungan dengan database yang di mana concurrent programming diterapkan.
   - Karena pada concurrent programming, beberapa proses akan berjalan bersamaan, maka penggunaan 1 koneksi db akan menghambat proses tersebut. Perlu ada beberapa koneksi database, agar goroutine tidak rebutan objek koneksi database.
3. Failover 
   - Failover merupakan mekanisme backup ketika sebuah proses gagal. Pada konteks ini, failover mengarah ke proses untuk me-retry operasi insert ketika gagal. 
