# Tree (binary tree)
disbt binary tree karena bisa ada 2 cabang yaitu ke kanan, ke kiri
## pembuatan 
setiap cabang & root disimpan dlm node yg dpt berupa struct, konektor antar cabang/root bisa berupa node dlm bntk struct.
## penggambaran tree adlh pohon keturunan (silsilah) yg bercabang ke kiri, kanan
perlu dicatat pada binary tree bisa dianggap bahwa setiap node (ortu) paling banyak hanya memiliki 2 cabang/anak (ikut program kb) tp urutan kanan lbh tua dr kiri itu tdk berlaku dulu (karena ini akan tergantung algoritma yg digunakan, dimana penelusuran bisa dari kanan dulu ke kiri/sebaliknya)<br> 
root = akar, child right = cabangnya ke kanan, child left = cabangnya ke kiri<br>
suatu child dpt berubah menjd root dari if child tsb memiliki child lagi<br>
setiap node (child) / dlm silsilah itu org-nya memiliki level atau dlm silsilah disbt dg generasi, ex generasi ke 1 (paling tua, ex kakek) berarti level 1, generasi ke n (paling muda, ex kita) berarti level n
## Jenis jenis

## operasi
dalam melakukan pengunjungan ke node satu per satu (traversal), tree terbagi 3 jenis operasi yaitu
* Preorder : root/value => left => right 
* Inorder : left => value/root => right
* Postorder : left => right => value/root
## sumber materi
https://ieftimov.com/posts/golang-datastructures-trees/<br>
https://www.youtube.com/watch?v=-oYitelECuQ<br>
https://panjiologi.blogspot.com/2017/11/cara-mencari-preorder-inorder-postorder.html<br>
## sumber code
https://www.golangprograms.com/golang-program-to-implement-binary-tree.html<br>
https://santekno.com/cara-implementasi-binary-search-and-tree-golang/<br>
https://www.youtube.com/watch?v=nL7BHR5vJDc<br>
https://www.youtube.com/watch?v=-oYitelECuQ<br>
