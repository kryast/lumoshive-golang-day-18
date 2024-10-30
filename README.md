catatan harian Ahmad Syarifuddin
reference By Lumoshive Academy

Alur Repository pattern
1. Deklarasi Model Struct
2. Deklarasi repository struct
3. Deklarasi repository interface untuk menampung struct func
4. buat func repository interface untuk menampung data ke dalam database
5. buat struct func create berisikan query insert data ke dalam database
6. buat func create yg memiliki parameter (db, dan struct) lalu panggil func repository interface agar parameter tersebut tersimpan di dalam database.
7. buat struct func yg berisikan query select untuk mengambil data di dalam database.
8. buat func untuk mengambil data di dalam database
9. buat func untuk menampilkan data yg telah diambil di dalam database
10. panggil func menampilkan data di main.go