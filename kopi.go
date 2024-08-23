package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var kopiMap = map[string]bool{
    "Kopi Espresso":    true,
    "Kopi Latte":       true,
    "Kopi Americano":   true,
    "Kopi Cappuccino":  true,
    "Kopi Mocha":       true,
    "Kopi Campur Cuka": false,//kopi tidak enak
  }

  // Inisialisasi generator angka acak
  rand.Seed(time.Now().UnixNano())

  // Membuat slice dari key-key map kopi
  keys := make([]string, 0, len(kopiMap))
  for k := range kopiMap {
    keys = append(keys, k)
  }

  for len(keys) > 0 {
    // Pilih secara acak jenis kopi
    randomIndex := rand.Intn(len(keys))
    pilihan := keys[randomIndex]

    // Tampilkan hasilnya
    if kopiMap[pilihan] {
      fmt.Printf("%s : aman\n", pilihan)
    } else {
      fmt.Printf("%s : gagal\n", pilihan)
      break
    }

    // Hapus jenis kopi yang sudah dipilih
    keys = append(keys[:randomIndex], keys[randomIndex+1:]...)
  }
}