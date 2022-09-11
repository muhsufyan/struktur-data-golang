# Graph
graph terdiri dari vertex (node), edge (konektor/jalur).<br>
jd sementara ini kita samakan saja antara node dg vertex<br>
selain itu terdpt istilah adjacency. Adjacency: Dua verteks x dan y yang berlainan disebut berhubungan langsung (adjacent) jika terdapat sisi {x, y} dalam E<br>
misal :<br>
V = {0, 1, 2, 3} artinya node nya ada 0,1,2,3<br>
E = {(0,1), (0,2), (0,3), (1,2)} artinya node 0 terhub ke node 1, node 0 terhub ke node 2, node 0 terhub ke node 3, node 1 terhub ke node 2<br> 
## pembuatan 
vertex (node) bisa dibuat dg struct, class as object<br>
edge bisa dibuat dg array dll. edge bisa merupakan matriks yang menyatakan hub antara vertex dimana 0 artinya tdk terhub & 1/angka apapun artinya terhub (if angka itu berarti weight nya)<br>
graph bisa dibuat menggunakan array, linked list, dll<br>
## penggambaran graph
## Jenis jenis
### berdasarkan arahnya
* direct : punya arah & sifat arahnya hanya 1 arah (dari vertex yyy ke vertex xxx)
* undirect : arahnya bolak balik/2 arah (dari vertex 1 ke vertex 2 & sebaliknya yaitu dr vertex 2 ke vertex 1)
### berdasarkan weightnya
* weight : punya nilai (misal jarak dari kota 1 ke kota 2 adlh 2 km)
* unweight : tidak punya nilai (ex jarak dari kota 1 ke kota 2 tdk ada jaraknya)
## operasi

## sumber materi
* https://www.trivusi.web.id/2022/07/struktur-data-graph.html
* http://aren.cs.ui.ac.id/sda/archive/1998/handout/handout19.html
* https://ocw.upj.ac.id/files/Handout-INF202-INF202-Struktur-Data-Wayan-Pertemuan-15.pdf
## sumber kode
* https://kalkicode.com/count-number-of-edges-in-an-undirected-graph-in-go