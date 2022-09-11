# Queue
data yang pertama masuk akan keluar pertama juga (elemen data yang masuk pertama akan keluar lebih dulu). Biasanya kita sering tahu istrila FIFO (First In First Out)<br>
## pembuatan 
queue dpt dibuat memakai:<br>
* Slices 
* Structures
* Link List
#
adapun tipe datanya bisa berupa array, slice, interface (if want dinamis)
## penggambaran queue adlh orang yang mengantri
barisan orang yang menunggu untuk membeli tiket di gedung bioskop. Orang yang baru datang akan bergabung dengan barisan dari ujung dan orang yang berdiri di depan akan menjadi yang pertama mendapatkan tiket dan meninggalkan barisan

## Jenis jenis
* Simple Queue
* Circular Queue
* Priority Queue
* Double-Ended Queue (Dequeue)
## operasi
* Enqueue: Menambahkan elemen ke akhir antrian
* Dequeue: Menghapus elemen dari depan antrian
* IsEmpty: Memeriksa apakah antrian kosong
* IsFull: Memeriksa apakah antrian sudah penuh
* Peek: Mengetahui nilai bagian depan antrian
* Initialize: Membuat antrian baru tanpa elemen data (kosong)
## sumber materi
https://www.trivusi.web.id/2022/07/struktur-data-queue.html#:~:text=Queue%20adalah%20struktur%20data%20linier,First%20In%2C%20First%20Out).<br>
https://santekno.com/cara-implementasi-queue-antrian-golang/<br>
## sumber code
https://santekno.com/cara-implementasi-queue-antrian-golang/<br>
https://www.geeksforgeeks.org/queue-in-go-language/<br>
#### operasi lengkap ada di directory struct
#### jenis jenis queue pada golang belum ada, hanya simple queue