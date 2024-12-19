package main

import "fmt"

func main() {
    var num, jumlah, count float64
    for {
        fmt.Scan(&num)
        if num == 9999 {
            break
        }
        jumlah += num
        count++
    }
    if count > 0 {
        fmt.Printf("Rerata: %.2f\n", jumlah/count)
    } else {
        fmt.Println("Tidak ada bilangan yang dimasukkan.")
    }
}