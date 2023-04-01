package main

import (
	"fmt"
	"os"
)

type Biodata struct {
    Absen   int
    Nama    string
    Alamat  string
    Pekerjaan string
    Alasan  string
}

var biodataList = []Biodata{
    {1, "Dani", "Bandung", "Freelancer", "Belajar Go-lang"},
    {2, "Setiawan", "Bandung Barat", "Network Administ", "Menambah Penghasilan"},
    {3, "Dani Setiawan", "Cililin", "Back-End", "Memahi Go-lang"},
}

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println("Usage: go run biodata.go [absen]")
        return
    }

    absen := args[0]
    for _, biodata := range biodataList {
        if fmt.Sprintf("%d", biodata.Absen) == absen {
            fmt.Println("Nama:", biodata.Nama)
            fmt.Println("Alamat:", biodata.Alamat)
            fmt.Println("Pekerjaan:", biodata.Pekerjaan)
            fmt.Println("Alasan memilih kelas Golang:", biodata.Alasan)
            return
        }
    }

    fmt.Println("Teman dengan absen", absen, "tidak ditemukan")
}