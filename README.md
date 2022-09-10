# Linked list
linked list terdiri dr list (semua data)<br>
list terdiri dari kumpulan node (semacam kontainer nya lah)<br>
node sendiri terdiri dari value & penunjuk/reference/konektor<br>
value mrpkn data nya sedangkan reference/konektor itu ada 2 yaitu prev (penunjuk ke node sebelumnya) & next (penunjuk/konektor ke node setelahnya)<br>
selain itu, linked list memiliki bagian kepala (head) atau bagian paling awal (kalau slice [0]) dan ekor (tail) atau bagian paling akhir (kalau slice ini [-1])<br> 
ex 4->5->6->7 : dr sana dpt dik bahwa: <br>
value 4 (node ke 1) => head, value 7 (node ke 4) => tail, length nya ada 4 node, if posisi kita ada di node ke 2 (value=5) maka kita punya konektor prev ke node 1 (value=4), sedangkan konektor next ke node 3 (value=6)<br>
untuk node ke 1 (value=4) atau head maka konektor prev ke node NULL sedangkan koenktor next ke node 2 (value=5)<br>
tail/node ke 4 (value=7) maka konektor prev ke node 3 (value=6) sedangkan konektor next ke NULL<br>
## pembuatan node & list
node dan list dapat dibuat dalam bentuk struct (go, c, dll) atau class (python, java, dll)
## penggambaran linked list sprti kereta api
linked list itu sprti kereta api (untuk yg double) dimana terdpt gerbong/lokomotif (dlm double linked list ini disbt node)<br>
gerbong tempat masinis/gerbong mesin untuk menggerakan kereta (dlm double linked list ini disbt head)<br>
dan kereta api yg merupakan kumpulan gerbong (dlm double linked list ini disbt list)<br>
gerbong-gerbong akan saling tersambung satu dg yg lainnya melalui sambungan (dlm double linked list ini disbt prev dan next yg tipe datanya adlh node)<br>
isi dari gerbong (bisa berupa barang/org) dlm linked list itu disbt dg value/data, tipe nya dpt berupa string,int, interface(if di golang), dll (pokoknya tipe data nya untuk menyimpan data sesuai kebthn)<br>
gerbong dibagian terakhir ada sambungan tp tdk digunakan (dlm double linked list ini disbt tail dan isinya NULL)<br>
## Jenis jenis
single : 1 konektor<br>
double : 2 konektor<br>
circular single/double : tail terkoneksi ke head (terjd siklus)<br>
## operasi
traversal : mengunjungi node 1 persatu sampai akhir<br>
hapus node : posisinya bisa hapus head, tail, atau tentukan posisi node yg ingin dihapus<br>
tambah node : posisinya head, tail, atau tentukan posisi node baru nya<br>
edit : posisinya bisa hapus head, tail, atau tentukan posisi node yg ingin diedit<br>
length of node : berapa panjang linked list nya (hitung saja jumlah node nya)<br>
pencarian (sorting)<br>
## sumber materi
https://www.simplilearn.com/tutorials/data-structure-tutorial/types-of-linked-list#:~:text=A%20linked%20list%20is%20a,list%20is%20called%20the%20tail.

## sumber code
https://www.developer.com/languages/linked-list-go/<br>
https://www.sibis.dev/linked-lists-in-golang-explained-with-real-world-application