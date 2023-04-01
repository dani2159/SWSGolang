package main

import "fmt"

func main() {
    
    i := 21
    j := true
    k := 123.456
    russia := 'Я'


    fmt.Printf("Nilai i: %v\n", i)
    fmt.Printf("Tipe data i: %T\n", i)
    fmt.Printf("Tanda persen: %%\n")
    fmt.Printf("Nilai j: %t\n", j)
    fmt.Printf("Unicode Russia: %c\n", russia)
    fmt.Printf("Nilai base 10 dari 21: %d\n", 21)
    fmt.Printf("Nilai base 8 dari 25: %o\n", 25)
    fmt.Printf("Nilai base 16 dari 15: %x\n", 15)
    fmt.Printf("Nilai base 16 dari 13: %X\n", 13)
    fmt.Printf("Unicode karakter Я: %U\n", 'Я')
    fmt.Printf("Nilai float: %f\n", k)
    fmt.Printf("Nilai float scientific: %e\n", k)
}