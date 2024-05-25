# Penjelasan Soal Nomor 2



# Penjelasan kompleksitas 

## Kompleksitas Waktu (Time Complexity)

## Inisialisasi Struktur Data:
- stack := []rune{}: Menginisialisasi stack sebagai slice kosong, O(1).
- brackets := map[rune]rune{...}: Menginisialisasi map untuk pasangan tanda kurung, O(1).

## Iterasi Melalui String:
- for _, char := range s: Melakukan iterasi melalui setiap karakter dalam string s. Karena kita mengiterasi seluruh string sekali, kompleksitasnya adalah O(n), di mana n adalah panjang string.

## Operation Pada Stack:
- stack = append(stack, char): Menambahkan elemen ke stack (push operation), O(1).
- len(stack) == 0 || stack[len(stack)-1] != brackets[char]: Mengecek kondisi stack dan mencocokkan pasangan tanda kurung, O(1).
- stack = stack[:len(stack)-1]: Menghapus elemen dari stack (pop operation), O(1).

## Pengecekan Akhir:

- if len(stack) == 0: Mengecek apakah stack kosong setelah semua iterasi, O(1).
Secara keseluruhan, kompleksitas waktu dari fungsi ini adalah O(n) karena setiap operasi dalam loop memiliki kompleksitas O(1), dan kita hanya mengiterasi string sekali.

## Kompleksitas Ruang (Space Complexity)

## Penggunaan Stack:

Dalam skenario terburuk, stack akan menyimpan semua karakter pembuka dalam string sebelum menemukan pasangan penutupnya. Oleh karena itu, dalam kasus terburuk, stack dapat menyimpan hingga n elemen, di mana n adalah panjang string.

## Map untuk Pasangan Tanda Kurung:

Map brackets memiliki ukuran tetap dengan 3 pasangan, sehingga kompleksitas ruangnya adalah O(1).
Secara keseluruhan, kompleksitas ruang dari fungsi ini adalah O(n) karena dalam skenario terburuk, stack akan menyimpan hingga n elemen.

## Ringkasan Kompleksitas
- Kompleksitas Waktu: O(n), karena kita mengiterasi setiap karakter dalam string satu kali.
- Kompleksitas Ruang: O(n), karena dalam skenario terburuk, kita mungkin perlu menyimpan semua karakter pembuka dalam stack.

Function IsBalanced ini cukup efisien baik dalam hal waktu maupun ruang, menggunakan struktur data stack untuk memastikan bahwa tanda kurung dalam string seimbang.