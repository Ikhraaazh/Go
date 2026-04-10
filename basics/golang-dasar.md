# Go-Lang Dasar

> Materi by Eko Kurniawan Khannedy — [ProgrammerZamanNow](https://www.youtube.com/c/ProgrammerZamanNow)

---

## Daftar Isi

1. [Pengenalan Go-Lang](#pengenalan-go-lang)
2. [Menginstall Go-Lang](#menginstall-go-lang)
3. [Membuat Project](#membuat-project)
4. [Program Hello World](#program-hello-world)
5. [Tipe Data Number](#tipe-data-number)
6. [Tipe Data Boolean](#tipe-data-boolean)
7. [Tipe Data String](#tipe-data-string)
8. [Variable](#variable)
9. [Constant](#constant)
10. [Konversi Tipe Data](#konversi-tipe-data)
11. [Type Declarations](#type-declarations)
12. [Operasi Matematika](#operasi-matematika)
13. [Operasi Perbandingan](#operasi-perbandingan)
14. [Operasi Boolean](#operasi-boolean)
15. [Tipe Data Array](#tipe-data-array)
16. [Tipe Data Slice](#tipe-data-slice)
17. [Tipe Data Map](#tipe-data-map)
18. [If Expression](#if-expression)
19. [Switch Expression](#switch-expression)
20. [For Loops](#for-loops)
21. [Break & Continue](#break--continue)
22. [Function](#function)
23. [Function Parameter](#function-parameter)
24. [Function Return Value](#function-return-value)
25. [Returning Multiple Values](#returning-multiple-values)
26. [Named Return Values](#named-return-values)
27. [Variadic Function](#variadic-function)
28. [Function Value](#function-value)
29. [Function as Parameter](#function-as-parameter)
30. [Anonymous Function](#anonymous-function)
31. [Recursive Function](#recursive-function)
32. [Closure](#closure)
33. [Defer, Panic, dan Recover](#defer-panic-dan-recover)
34. [Struct](#struct)
35. [Struct Method](#struct-method)
36. [Interface](#interface)
37. [Interface Kosong](#interface-kosong)
38. [Nil](#nil)
39. [Type Assertions](#type-assertions)
40. [Pointer](#pointer)
41. [Package & Import](#package--import)
42. [Access Modifier](#access-modifier)
43. [Package Initialization](#package-initialization)
44. [Error](#error)

---

## Pengenalan Go-Lang

### Sejarah
- Dibuat di Google menggunakan bahasa C
- Di-rilis ke publik sebagai open source pada tahun **2009**
- Populer sejak digunakan untuk membuat **Docker** (2011)
- Dipakai juga di: Kubernetes, Prometheus, CockroachDB
- Saat ini populer untuk pembuatan **Backend API** di Microservices

### Kenapa Belajar Go-Lang?
- Sintaks sangat sederhana, mudah dipelajari
- Mendukung **concurrency programming** (cocok untuk multicore processor)
- Ada **garbage collector** — tidak perlu manage memory manual seperti C
- Bahasa yang sedang naik daun

### Proses Development

```
Source Code (.go) → Compiler (go build) → Binary → Dijalankan
```

Atau bisa langsung dijalankan tanpa kompilasi:
```
Source Code (.go) → go run → Output
```

---

## Menginstall Go-Lang

1. Download compiler di [https://golang.org/](https://golang.org/)
2. Install compiler
3. Verifikasi instalasi:

```bash
go version
```

**IDE yang direkomendasikan:**
- Visual Studio Code
- JetBrains GoLand

---

## Membuat Project

Project di Go-Lang disebut **module**. Untuk membuat module baru:

```bash
go mod init nama-module
```

Contoh:
```bash
go mod init belajar-golang
```

Perintah ini akan menghasilkan file `go.mod`.

---

## Program Hello World

### Main Function

- Go-Lang mirip C/C++, butuh `main function` sebagai entry point
- Keyword untuk membuat function: `func`
- `main` function harus ada di dalam `package main`
- Titik koma di akhir baris tidak wajib

```go
package main

func main() {

}
```

### Println

Untuk mencetak output, perlu import module `fmt` terlebih dahulu:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

### Kompilasi dan Menjalankan

**Kompilasi:**
```bash
go build helloworld.go
./helloworld
```

**Menjalankan langsung tanpa kompilasi:**
```bash
go run helloworld.go
```

### Multiple Main Function

Di Go-Lang, nama function di satu module harus **unik**. Jika ada dua file dengan `func main()`, build akan error:

```
# ./...
./sample.go:5:6: main redeclared in this block
```

> **Solusi saat belajar:** jalankan file satu per satu dengan `go run namafile.go`.

---

## Tipe Data Number

### Integer

| Tipe Data | Nilai Minimum | Nilai Maksimum |
|-----------|--------------|----------------|
| `int8` | -128 | 127 |
| `int16` | -32768 | 32767 |
| `int32` | -2147483648 | 2147483647 |
| `int64` | -9223372036854775808 | 9223372036854775807 |
| `uint8` | 0 | 255 |
| `uint16` | 0 | 65535 |
| `uint32` | 0 | 4294967295 |
| `uint64` | 0 | 18446744073709551615 |

### Floating Point

| Tipe Data | Keterangan |
|-----------|-----------|
| `float32` | Presisi rendah |
| `float64` | Presisi tinggi |
| `complex64` | Bilangan kompleks dengan float32 |
| `complex128` | Bilangan kompleks dengan float64 |

### Alias

| Tipe Data | Alias untuk |
|-----------|------------|
| `byte` | `uint8` |
| `rune` | `int32` |
| `int` | Minimal `int32` |
| `uint` | Minimal `uint32` |

```go
package main

import "fmt"

func main() {
    var a int8 = 100
    var b int16 = 30000
    var c int32 = 1000000
    var d int64 = 9000000000

    var e float32 = 3.14
    var f float64 = 3.14159265358979

    fmt.Println(a, b, c, d, e, f)
}
```

---

## Tipe Data Boolean

Tipe data dengan dua nilai: **benar** atau **salah**.

| Nilai | Keterangan |
|-------|-----------|
| `true` | Benar |
| `false` | Salah |

```go
package main

import "fmt"

func main() {
    var benar bool = true
    var salah bool = false

    fmt.Println(benar)  // true
    fmt.Println(salah)  // false
}
```

---

## Tipe Data String

- Kumpulan karakter
- Diapit dengan tanda `"` (petik dua)
- Jumlah karakter bisa 0 sampai tak terhingga

### Function untuk String

| Function | Keterangan |
|----------|-----------|
| `len("string")` | Menghitung jumlah karakter |
| `"string"[index]` | Mengambil karakter pada posisi tertentu |

```go
package main

import "fmt"

func main() {
    var nama string = "Eko Kurniawan Khannedy"
    fmt.Println(nama)
    fmt.Println(len(nama))     // panjang string
    fmt.Println(nama[0])       // karakter pertama (dalam bentuk byte)
    fmt.Println(string(nama[0])) // "E"
}
```

---

## Variable

- Tempat untuk menyimpan data
- Di Go-Lang, variable hanya bisa menyimpan **satu tipe data**
- Keyword: `var`

### Cara Deklarasi

```go
package main

import "fmt"

func main() {
    // cara 1: deklarasi dengan tipe data
    var nama string
    nama = "Eko"

    // cara 2: deklarasi + inisialisasi (tipe bisa dihilangkan)
    var umur = 30

    // cara 3: shorthand (tanpa var, langsung :=)
    alamat := "Jakarta"

    fmt.Println(nama, umur, alamat)
}
```

### Multiple Variable

```go
package main

import "fmt"

func main() {
    var (
        firstName = "Eko"
        lastName  = "Khannedy"
        umur      = 30
    )

    fmt.Println(firstName, lastName, umur)
}
```

---

## Constant

- Nilai **tidak bisa diubah** setelah pertama kali diberi nilai
- Keyword: `const`
- Wajib langsung diinisialisasi saat deklarasi

```go
package main

import "fmt"

func main() {
    const PI = 3.14
    const MAX_VALUE = 100

    fmt.Println(PI)
    fmt.Println(MAX_VALUE)

    // PI = 3.15 // ERROR: cannot assign to PI
}
```

### Multiple Constant

```go
package main

import "fmt"

func main() {
    const (
        PI        = 3.14
        MAX_VALUE = 100
        APP_NAME  = "Belajar Go"
    )

    fmt.Println(PI, MAX_VALUE, APP_NAME)
}
```

---

## Konversi Tipe Data

Untuk mengkonversi tipe data, gunakan syntax `TipeData(variable)`:

```go
package main

import "fmt"

func main() {
    var nilai32 int32 = 32768
    var nilai64 int64 = int64(nilai32)

    var nilai float64 = float64(nilai32)

    fmt.Println(nilai32)  // 32768
    fmt.Println(nilai64)  // 32768
    fmt.Println(nilai)    // 32768
}
```

```go
package main

import "fmt"

func main() {
    var a int64 = 100000
    var b int32 = int32(a)  // konversi int64 → int32
    var c int16 = int16(b)  // konversi int32 → int16
    var d int8  = int8(c)   // konversi int16 → int8 (bisa overflow!)

    fmt.Println(a, b, c, d)
}
```

---

## Type Declarations

Membuat **tipe data baru** dari tipe data yang sudah ada. Berguna untuk membuat alias yang lebih mudah dimengerti.

```go
package main

import "fmt"

type NoKTP string
type Umur int

func main() {
    var ktp NoKTP = "1234567890123456"
    var umur Umur = 30

    fmt.Println(ktp)
    fmt.Println(umur)
}
```

---

## Operasi Matematika

| Operator | Keterangan |
|----------|-----------|
| `+` | Penjumlahan |
| `-` | Pengurangan |
| `*` | Perkalian |
| `/` | Pembagian |
| `%` | Sisa Pembagian (modulo) |

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b int = 3

    fmt.Println(a + b)  // 13
    fmt.Println(a - b)  // 7
    fmt.Println(a * b)  // 30
    fmt.Println(a / b)  // 3
    fmt.Println(a % b)  // 1
}
```

### Augmented Assignments

| Operasi Matematika | Augmented |
|--------------------|-----------|
| `a = a + 10` | `a += 10` |
| `a = a - 10` | `a -= 10` |
| `a = a * 10` | `a *= 10` |
| `a = a / 10` | `a /= 10` |
| `a = a % 10` | `a %= 10` |

```go
package main

import "fmt"

func main() {
    var a = 10
    a += 5
    fmt.Println(a)  // 15
    a *= 2
    fmt.Println(a)  // 30
}
```

### Unary Operator

| Operator | Keterangan |
|----------|-----------|
| `++` | `a = a + 1` |
| `--` | `a = a - 1` |
| `-` | Negative |
| `+` | Positive |
| `!` | Boolean kebalikan |

```go
package main

import "fmt"

func main() {
    var a = 10
    a++
    fmt.Println(a)  // 11
    a--
    fmt.Println(a)  // 10

    var b bool = true
    fmt.Println(!b) // false
}
```

---

## Operasi Perbandingan

Hasil operasi perbandingan selalu berupa **boolean** (`true` / `false`).

| Operator | Keterangan |
|----------|-----------|
| `>` | Lebih Dari |
| `<` | Kurang Dari |
| `>=` | Lebih Dari Sama Dengan |
| `<=` | Kurang Dari Sama Dengan |
| `==` | Sama Dengan |
| `!=` | Tidak Sama Dengan |

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b int = 20

    fmt.Println(a > b)   // false
    fmt.Println(a < b)   // true
    fmt.Println(a >= b)  // false
    fmt.Println(a <= b)  // true
    fmt.Println(a == b)  // false
    fmt.Println(a != b)  // true
}
```

---

## Operasi Boolean

| Operator | Keterangan |
|----------|-----------|
| `&&` | Dan (AND) |
| `\|\|` | Atau (OR) |
| `!` | Kebalikan (NOT) |

### Tabel Kebenaran &&

| Nilai 1 | Operator | Nilai 2 | Hasil |
|---------|----------|---------|-------|
| `true` | `&&` | `true` | `true` |
| `true` | `&&` | `false` | `false` |
| `false` | `&&` | `true` | `false` |
| `false` | `&&` | `false` | `false` |

### Tabel Kebenaran \|\|

| Nilai 1 | Operator | Nilai 2 | Hasil |
|---------|----------|---------|-------|
| `true` | `\|\|` | `true` | `true` |
| `true` | `\|\|` | `false` | `true` |
| `false` | `\|\|` | `true` | `true` |
| `false` | `\|\|` | `false` | `false` |

```go
package main

import "fmt"

func main() {
    var nilaiUjian = 70
    var nilaiAbsensi = 80

    var lulus = nilaiUjian >= 75 && nilaiAbsensi >= 75
    fmt.Println(lulus) // false

    var bolehIkutUjian = nilaiUjian >= 75 || nilaiAbsensi >= 75
    fmt.Println(bolehIkutUjian) // true
}
```

---

## Tipe Data Array

- Kumpulan data dengan **tipe yang sama**
- Ukuran harus ditentukan saat dibuat dan **tidak bisa berubah**
- Index dimulai dari **0**

```go
package main

import "fmt"

func main() {
    // deklarasi array dengan ukuran 3
    var names [3]string
    names[0] = "Eko"
    names[1] = "Kurniawan"
    names[2] = "Khannedy"

    fmt.Println(names)       // [Eko Kurniawan Khannedy]
    fmt.Println(names[0])    // Eko
    fmt.Println(len(names))  // 3
}
```

### Membuat Array Langsung

```go
package main

import "fmt"

func main() {
    // inisialisasi langsung
    names := [3]string{"Eko", "Kurniawan", "Khannedy"}

    // atau pakai ... untuk auto-count
    hobbies := [...]string{"Coding", "Gaming", "Reading"}

    fmt.Println(names)
    fmt.Println(hobbies)
    fmt.Println(len(hobbies)) // 3
}
```

### Function Array

| Operasi | Keterangan |
|---------|-----------|
| `len(array)` | Mendapatkan panjang array |
| `array[index]` | Mendapat data di posisi index |
| `array[index] = value` | Mengubah data di posisi index |

---

## Tipe Data Slice

- Mirip Array, tapi **ukurannya bisa berubah**
- Slice selalu **terhubung** ke Array
- Memiliki 3 properti: **pointer**, **length**, **capacity**

### Membuat Slice dari Array

| Cara | Keterangan |
|------|-----------|
| `array[low:high]` | Dari index `low` sampai sebelum `high` |
| `array[low:]` | Dari index `low` sampai akhir |
| `array[:high]` | Dari index 0 sampai sebelum `high` |
| `array[:]` | Seluruh array |

```go
package main

import "fmt"

func main() {
    names := [...]string{"Eko", "Kurniawan", "Khannedy", "Budi", "Joko"}

    slice1 := names[1:3]  // ["Kurniawan", "Khannedy"]
    slice2 := names[1:]   // ["Kurniawan", "Khannedy", "Budi", "Joko"]
    slice3 := names[:3]   // ["Eko", "Kurniawan", "Khannedy"]

    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(slice3)
}
```

### Function Slice

| Operasi | Keterangan |
|---------|-----------|
| `len(slice)` | Mendapatkan panjang |
| `cap(slice)` | Mendapat kapasitas |
| `append(slice, data)` | Tambah data ke akhir slice |
| `make([]TypeData, length, capacity)` | Membuat slice baru |
| `copy(destination, source)` | Menyalin slice |

```go
// Append
package main

import "fmt"

func main() {
    var hobbies []string  // slice kosong

    hobbies = append(hobbies, "Coding")
    hobbies = append(hobbies, "Gaming", "Reading")

    fmt.Println(hobbies)       // [Coding Gaming Reading]
    fmt.Println(len(hobbies))  // 3
    fmt.Println(cap(hobbies))  // tergantung alokasi
}
```

```go
// Make Slice
package main

import "fmt"

func main() {
    slice := make([]string, 3, 5)  // length 3, capacity 5
    slice[0] = "Eko"
    slice[1] = "Kurniawan"
    slice[2] = "Khannedy"

    fmt.Println(slice)
    fmt.Println(len(slice))  // 3
    fmt.Println(cap(slice))  // 5
}
```

```go
// Copy Slice
package main

import "fmt"

func main() {
    source := []string{"Eko", "Kurniawan", "Khannedy"}
    destination := make([]string, len(source))

    copy(destination, source)

    fmt.Println(source)      // [Eko Kurniawan Khannedy]
    fmt.Println(destination) // [Eko Kurniawan Khannedy]
}
```

### Hati-Hati: Array vs Slice

```go
package main

import "fmt"

func main() {
    // INI ARRAY (pakai angka atau ...)
    array := [3]string{"Eko", "Kurniawan", "Khannedy"}

    // INI SLICE (tanpa angka)
    slice := []string{"Eko", "Kurniawan", "Khannedy"}

    fmt.Println(array)
    fmt.Println(slice)
}
```

---

## Tipe Data Map

- Kumpulan **key-value** (kamus)
- Key bersifat **unik** — jika key sama, data lama akan tertimpa
- Berbeda dengan Array/Slice, jumlah data tidak dibatasi

```go
package main

import "fmt"

func main() {
    // Membuat map
    var person map[string]string = map[string]string{}
    person["name"] = "Eko"
    person["address"] = "Jakarta"

    fmt.Println(person)
    fmt.Println(person["name"])    // Eko
    fmt.Println(person["address"]) // Jakarta
}
```

### Function Map

| Operasi | Keterangan |
|---------|-----------|
| `len(map)` | Jumlah data |
| `map[key]` | Mengambil data |
| `map[key] = value` | Mengubah/menambah data |
| `make(map[TypeKey]TypeValue)` | Membuat map baru |
| `delete(map, key)` | Menghapus data |

```go
package main

import "fmt"

func main() {
    person := map[string]string{
        "name":    "Eko",
        "address": "Jakarta",
    }

    fmt.Println(len(person))  // 2

    delete(person, "address")
    fmt.Println(person)       // map[name:Eko]
}
```

---

## If Expression

```go
package main

import "fmt"

func main() {
    nilai := 70

    if nilai >= 75 {
        fmt.Println("Lulus")
    }
}
```

### Else Expression

```go
package main

import "fmt"

func main() {
    nilai := 70

    if nilai >= 75 {
        fmt.Println("Lulus")
    } else {
        fmt.Println("Tidak Lulus")
    }
}
```

### Else If Expression

```go
package main

import "fmt"

func main() {
    nilai := 85

    if nilai >= 90 {
        fmt.Println("A")
    } else if nilai >= 80 {
        fmt.Println("B")
    } else if nilai >= 70 {
        fmt.Println("C")
    } else {
        fmt.Println("D")
    }
}
```

### If dengan Short Statement

```go
package main

import "fmt"

func main() {
    if nilai := 80; nilai >= 75 {
        fmt.Println("Lulus dengan nilai", nilai)
    } else {
        fmt.Println("Tidak lulus dengan nilai", nilai)
    }
    // variabel 'nilai' tidak bisa diakses di luar if
}
```

---

## Switch Expression

```go
package main

import "fmt"

func main() {
    nilai := "B"

    switch nilai {
    case "A":
        fmt.Println("Sangat Baik")
    case "B":
        fmt.Println("Baik")
    case "C":
        fmt.Println("Cukup")
    default:
        fmt.Println("Tidak diketahui")
    }
}
```

### Switch dengan Short Statement

```go
package main

import "fmt"

func main() {
    switch nilai := "B"; nilai {
    case "A":
        fmt.Println("Sangat Baik")
    case "B":
        fmt.Println("Baik")
    default:
        fmt.Println("Tidak diketahui")
    }
}
```

### Switch Tanpa Kondisi

```go
package main

import "fmt"

func main() {
    nilai := 75

    switch {
    case nilai >= 90:
        fmt.Println("A")
    case nilai >= 80:
        fmt.Println("B")
    case nilai >= 70:
        fmt.Println("C")
    default:
        fmt.Println("D")
    }
}
```

---

## For Loops

Go-Lang hanya punya satu keyword loop: `for`.

```go
package main

import "fmt"

func main() {
    // for seperti while
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }
}
```

### For dengan Statement

```go
package main

import "fmt"

func main() {
    // for klasik: init; kondisi; post
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

### For Range

```go
package main

import "fmt"

func main() {
    names := []string{"Eko", "Kurniawan", "Khannedy"}

    // iterasi slice
    for index, value := range names {
        fmt.Println(index, value)
    }

    // iterasi map
    person := map[string]string{
        "name":    "Eko",
        "address": "Jakarta",
    }
    for key, value := range person {
        fmt.Println(key, ":", value)
    }
}
```

---

## Break & Continue

- `break` — hentikan seluruh perulangan
- `continue` — skip iterasi saat ini, lanjut ke iterasi berikutnya

```go
// Break
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break  // berhenti di 5
        }
        fmt.Println(i)
    }
}
```

```go
// Continue
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue  // skip bilangan genap
        }
        fmt.Println(i)  // hanya bilangan ganjil
    }
}
```

---

## Function

- Blok kode yang bisa digunakan **berulang-ulang**
- Keyword: `func`

```go
package main

import "fmt"

func sayHello() {
    fmt.Println("Hello!")
}

func main() {
    sayHello()
    sayHello()
    sayHello()
}
```

---

## Function Parameter

```go
package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
    fmt.Println("Hello", firstName, lastName)
}

func main() {
    sayHelloTo("Eko", "Kurniawan")
    sayHelloTo("Budi", "Santoso")
}
```

---

## Function Return Value

```go
package main

import "fmt"

func getFullName(firstName string, lastName string) string {
    return firstName + " " + lastName
}

func main() {
    fullName := getFullName("Eko", "Kurniawan")
    fmt.Println(fullName)
}
```

---

## Returning Multiple Values

```go
package main

import "fmt"

func getMinMax(array []int) (int, int) {
    min := array[0]
    max := array[0]

    for _, value := range array {
        if value < min {
            min = value
        }
        if value > max {
            max = value
        }
    }

    return min, max
}

func main() {
    numbers := []int{1, 5, 3, 8, 2}
    min, max := getMinMax(numbers)
    fmt.Println("Min:", min)  // 1
    fmt.Println("Max:", max)  // 8
}
```

### Menghiraukan Return Value

Gunakan `_` (underscore) untuk mengabaikan salah satu return value:

```go
package main

import "fmt"

func getMinMax(array []int) (int, int) {
    min := array[0]
    max := array[0]
    for _, value := range array {
        if value < min { min = value }
        if value > max { max = value }
    }
    return min, max
}

func main() {
    numbers := []int{1, 5, 3, 8, 2}

    min, _ := getMinMax(numbers)  // abaikan max
    fmt.Println("Min:", min)

    _, max := getMinMax(numbers)  // abaikan min
    fmt.Println("Max:", max)
}
```

---

## Named Return Values

```go
package main

import "fmt"

func getMinMax(array []int) (min int, max int) {
    min = array[0]
    max = array[0]

    for _, value := range array {
        if value < min { min = value }
        if value > max { max = value }
    }

    return  // return tanpa nilai, otomatis return min dan max
}

func main() {
    numbers := []int{1, 5, 3, 8, 2}
    min, max := getMinMax(numbers)
    fmt.Println(min, max)
}
```

---

## Variadic Function

- Parameter terakhir bisa menerima **lebih dari satu nilai** (seperti array)
- Notasi: `...TipeData`

```go
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))           // 6
    fmt.Println(sum(1, 2, 3, 4, 5))     // 15
    fmt.Println(sum(10, 20))            // 30
}
```

### Slice sebagai Variadic Parameter

```go
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}

    // kirim slice ke variadic pakai ...
    total := sum(numbers...)
    fmt.Println(total)  // 15
}
```

---

## Function Value

Function adalah **first-class citizen** di Go-Lang — bisa disimpan dalam variable:

```go
package main

import "fmt"

func main() {
    // simpan function di variable
    add := func(a, b int) int {
        return a + b
    }

    fmt.Println(add(3, 5))   // 8
    fmt.Println(add(10, 20)) // 30
}
```

---

## Function as Parameter

```go
package main

import "fmt"

func printResult(value int, transformer func(int) int) {
    result := transformer(value)
    fmt.Println("Result:", result)
}

func double(value int) int {
    return value * 2
}

func square(value int) int {
    return value * value
}

func main() {
    printResult(10, double)  // Result: 20
    printResult(10, square)  // Result: 100
}
```

### Function Type Declarations

```go
package main

import "fmt"

// buat alias untuk tipe function
type Transformer func(int) int

func printResult(value int, transformer Transformer) {
    fmt.Println(transformer(value))
}

func double(value int) int {
    return value * 2
}

func main() {
    printResult(10, double)      // 20

    // atau pakai anonymous function
    printResult(10, func(v int) int {
        return v * v
    }) // 100
}
```

---

## Anonymous Function

Function tanpa nama, langsung dibuat di variable atau saat dipanggil:

```go
package main

import "fmt"

func main() {
    // anonymous function di variable
    hello := func(name string) {
        fmt.Println("Hello", name)
    }
    hello("Eko")

    // langsung dipanggil (IIFE)
    func(name string) {
        fmt.Println("Hello", name)
    }("Budi")
}
```

---

## Recursive Function

Function yang **memanggil dirinya sendiri**. Contoh kasus: Faktorial.

```go
// Cara 1: Menggunakan for loop
package main

import "fmt"

func factorialLoop(value int) int {
    result := 1
    for i := 2; i <= value; i++ {
        result *= i
    }
    return result
}

func main() {
    fmt.Println(factorialLoop(5))   // 120
    fmt.Println(factorialLoop(10))  // 3628800
}
```

```go
// Cara 2: Menggunakan recursive
package main

import "fmt"

func factorialRecursive(value int) int {
    if value == 1 {
        return 1  // base case
    }
    return value * factorialRecursive(value-1)
}

func main() {
    fmt.Println(factorialRecursive(5))   // 120
    fmt.Println(factorialRecursive(10))  // 3628800
}
```

---

## Closure

Function yang bisa **mengakses variable di scope sekitarnya**:

```go
package main

import "fmt"

func main() {
    counter := 0

    increment := func() {
        counter++  // akses variable di scope luar
    }

    increment()
    increment()
    increment()

    fmt.Println(counter)  // 3
}
```

```go
package main

import "fmt"

func makeCounter() func() int {
    counter := 0
    return func() int {
        counter++
        return counter
    }
}

func main() {
    counter1 := makeCounter()
    counter2 := makeCounter()

    fmt.Println(counter1())  // 1
    fmt.Println(counter1())  // 2
    fmt.Println(counter2())  // 1 (independent)
}
```

---

## Defer, Panic, dan Recover

### Defer

- Menjadwalkan eksekusi function **di akhir** function saat ini selesai
- Tetap dieksekusi walau terjadi error/panic

```go
package main

import "fmt"

func endProgram() {
    fmt.Println("Program selesai")
}

func main() {
    defer endProgram()  // dijadwalkan di akhir

    fmt.Println("Hello")
    fmt.Println("World")
    // Output: Hello → World → Program selesai
}
```

### Panic

- Menghentikan program secara paksa
- `defer` tetap akan dijalankan sebelum program berhenti

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Defer tetap jalan")

    fmt.Println("Mulai program")
    panic("terjadi error parah!")  // program berhenti di sini
    fmt.Println("Ini tidak akan dieksekusi")
}
```

### Recover

- Menangkap panic agar program **tidak mati**
- Harus digunakan di dalam `defer`

```go
// SALAH - recover di luar defer tidak akan berfungsi
package main

import "fmt"

func main() {
    // ini TIDAK akan menangkap panic
    message := recover()
    fmt.Println(message)

    panic("ups!")
}
```

```go
// BENAR - recover harus di dalam defer
package main

import "fmt"

func safeRun() {
    defer func() {
        if message := recover(); message != nil {
            fmt.Println("Recovered:", message)
        }
    }()

    fmt.Println("Mulai")
    panic("terjadi panic!")
    fmt.Println("Ini tidak dieksekusi")
}

func main() {
    safeRun()
    fmt.Println("Program tetap lanjut")
}
// Output:
// Mulai
// Recovered: terjadi panic!
// Program tetap lanjut
```

---

## Struct

- Template/prototype data
- Kumpulan dari **field** (atribut)
- Keyword: `type NamaStruct struct`

```go
package main

import "fmt"

type Customer struct {
    Name    string
    Email   string
    Age     int
}

func main() {
    // cara 1: buat object lalu assign
    var customer1 Customer
    customer1.Name = "Eko"
    customer1.Email = "eko@example.com"
    customer1.Age = 30

    fmt.Println(customer1)
    fmt.Println(customer1.Name)
}
```

### Struct Literals

```go
package main

import "fmt"

type Customer struct {
    Name  string
    Email string
    Age   int
}

func main() {
    // cara 2: struct literal (urutan sesuai field)
    customer1 := Customer{"Eko", "eko@example.com", 30}

    // cara 3: struct literal dengan nama field (lebih aman)
    customer2 := Customer{
        Name:  "Budi",
        Email: "budi@example.com",
        Age:   25,
    }

    fmt.Println(customer1)
    fmt.Println(customer2)
}
```

---

## Struct Method

Method adalah function yang menempel pada struct:

```go
package main

import "fmt"

type Customer struct {
    Name  string
    Email string
}

// method SayHello pada Customer
func (c Customer) SayHello() {
    fmt.Println("Hello, nama saya", c.Name)
}

func main() {
    customer := Customer{Name: "Eko", Email: "eko@example.com"}
    customer.SayHello()  // Hello, nama saya Eko
}
```

---

## Interface

- Tipe data **abstract** — hanya berisi definisi method, tidak ada implementasi
- Digunakan sebagai **kontrak**
- Implementasi dilakukan secara **otomatis** (duck typing)

```go
package main

import "fmt"

type HasName interface {
    GetName() string
}

type Person struct {
    Name string
}

// Person mengimplementasi HasName secara otomatis
// karena punya method GetName()
func (p Person) GetName() string {
    return p.Name
}

func sayHello(hasName HasName) {
    fmt.Println("Hello", hasName.GetName())
}

func main() {
    person := Person{Name: "Eko"}
    sayHello(person)  // Hello Eko
}
```

### Implementasi Interface

```go
package main

import "fmt"

type MakeSoundable interface {
    MakeSound()
}

type Cat struct {
    Name string
}

type Dog struct {
    Name string
}

func (c Cat) MakeSound() {
    fmt.Println(c.Name, "bersuara: Meow")
}

func (d Dog) MakeSound() {
    fmt.Println(d.Name, "bersuara: Woof")
}

func makeSound(animal MakeSoundable) {
    animal.MakeSound()
}

func main() {
    makeSound(Cat{Name: "Kitty"})  // Kitty bersuara: Meow
    makeSound(Dog{Name: "Doggy"})  // Doggy bersuara: Woof
}
```

---

## Interface Kosong

- Interface tanpa method sama sekali
- Semua tipe data otomatis mengimplementasinya
- Alias: `any`

```go
package main

import "fmt"

func printAnything(value interface{}) {
    fmt.Println(value)
}

// atau dengan 'any' (Go 1.18+)
func printAny(value any) {
    fmt.Println(value)
}

func main() {
    printAnything("Hello")
    printAnything(123)
    printAnything(true)
    printAnything([]string{"a", "b"})
}
```

**Contoh penggunaan di standard library:**
- `fmt.Println(a ...interface{})`
- `panic(v interface{})`
- `recover() interface{}`

---

## Nil

- Data **kosong** di Go-Lang
- Hanya bisa dipakai di: interface, function, map, slice, pointer, channel
- Variable dengan tipe data biasa (int, string, bool) punya **default value**, bukan nil

```go
package main

import "fmt"

func main() {
    // nil di interface
    var inter interface{} = nil
    fmt.Println(inter)  // <nil>

    // nil di slice
    var slice []string = nil
    fmt.Println(slice)  // []

    // nil di map
    var m map[string]string = nil
    fmt.Println(m)  // map[]
}
```

```go
package main

import "fmt"

type SayHello interface {
    Hello()
}

func callHello(say SayHello) {
    if say != nil {
        say.Hello()
    } else {
        fmt.Println("nil!")
    }
}

func main() {
    callHello(nil)  // nil!
}
```

---

## Type Assertions

Mengubah/mengecek tipe data dari interface ke tipe konkret:

```go
package main

import "fmt"

func main() {
    var value interface{} = "Hello World"

    // type assertion
    str := value.(string)
    fmt.Println(str)  // Hello World

    // type assertion aman (dengan cek ok)
    str2, ok := value.(string)
    if ok {
        fmt.Println("String:", str2)
    }

    num, ok2 := value.(int)
    if !ok2 {
        fmt.Println("Bukan int, num default:", num)  // 0
    }
}
```

### Type Assertions Menggunakan Switch

Lebih aman dari panic jika tipe salah:

```go
package main

import "fmt"

func checkType(value interface{}) {
    switch v := value.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Int:", v)
    case bool:
        fmt.Println("Bool:", v)
    default:
        fmt.Println("Tipe tidak diketahui:", v)
    }
}

func main() {
    checkType("Hello")  // String: Hello
    checkType(123)      // Int: 123
    checkType(true)     // Bool: true
    checkType(3.14)     // Tipe tidak diketahui: 3.14
}
```

---

## Pointer

### Pass by Value (Default Go-Lang)

Secara default, semua variable di Go-Lang di-passing sebagai **copy** (by value):

```go
package main

import "fmt"

func change(value int) {
    value = 200  // hanya mengubah copy
}

func main() {
    number := 100
    change(number)
    fmt.Println(number)  // masih 100 (tidak berubah)
}
```

### Pass by Reference dengan Pointer

- `&variable` → ambil alamat memory variable
- `*variable` → akses/ubah nilai di alamat memory tersebut

```go
package main

import "fmt"

func changeWithPointer(value *int) {
    *value = 200  // ubah nilai di alamat aslinya
}

func main() {
    number := 100
    changeWithPointer(&number)  // kirim alamat memory
    fmt.Println(number)  // 200 (berubah!)
}
```

### Operator &

```go
package main

import "fmt"

func main() {
    number := 100
    pointer := &number  // pointer ke number

    fmt.Println(number)   // 100
    fmt.Println(pointer)  // alamat memory, contoh: 0xc0000b4010
    fmt.Println(*pointer) // 100 (nilai di alamat tersebut)

    // ubah pointer, tidak mengubah number
    pointer2 := &number
    fmt.Println(pointer == pointer2) // true, sama-sama acu ke number
}
```

### Operator *

```go
package main

import "fmt"

func main() {
    number := 100
    pointer1 := &number
    pointer2 := &number

    // tanpa *, hanya mengubah pointer1 saja
    // pointer1 = ... // ini mengubah pointer, bukan nilai

    // dengan *, mengubah nilai di memory aslinya
    *pointer1 = 200

    fmt.Println(number)   // 200
    fmt.Println(*pointer1) // 200
    fmt.Println(*pointer2) // 200 (ikut berubah karena sama-sama acu ke number)
}
```

### Operator new

```go
package main

import "fmt"

func main() {
    // new() membuat pointer ke data kosong (zero value)
    pointer := new(int)
    fmt.Println(*pointer)  // 0

    *pointer = 100
    fmt.Println(*pointer)  // 100
}
```

### Pointer di Method

Sangat direkomendasikan pakai pointer di method agar tidak boros memory:

```go
package main

import "fmt"

type Customer struct {
    Name string
}

// TANPA pointer — data di-copy, perubahan tidak berpengaruh ke aslinya
func (c Customer) ChangeName(name string) {
    c.Name = name  // hanya ubah copy
}

// DENGAN pointer — ubah data aslinya
func (c *Customer) ChangeNamePointer(name string) {
    c.Name = name  // ubah data asli
}

func main() {
    customer := Customer{Name: "Eko"}

    customer.ChangeName("Budi")
    fmt.Println(customer.Name)  // Eko (tidak berubah)

    customer.ChangeNamePointer("Budi")
    fmt.Println(customer.Name)  // Budi (berubah!)
}
```

---

## Package & Import

- **Package** = folder untuk mengorganisir kode
- Satu folder = satu package
- Nama file bebas, yang penting nama package-nya sama

```
project/
├── main.go         → package main
└── helper/
    └── helper.go   → package helper
```

```go
// helper/helper.go
package helper

func SayHello(name string) string {
    return "Hello " + name
}
```

```go
// main.go
package main

import (
    "fmt"
    "belajar-golang/helper"
)

func main() {
    result := helper.SayHello("Eko")
    fmt.Println(result)  // Hello Eko
}
```

---

## Access Modifier

Di Go-Lang tidak ada keyword `public`/`private`. Cukup dari **huruf pertama nama**:

| Awalan | Akses |
|--------|-------|
| Huruf **BESAR** | Public (bisa diakses dari package lain) |
| Huruf **kecil** | Private (hanya di dalam package sendiri) |

```go
// helper/helper.go
package helper

var PublicVariable = "bisa diakses dari luar"  // public
var privateVariable = "tidak bisa dari luar"    // private

func PublicFunction() string {                  // public
    return "public function"
}

func privateFunction() string {                 // private
    return "private function"
}
```

```go
// main.go
package main

import (
    "fmt"
    "belajar-golang/helper"
)

func main() {
    fmt.Println(helper.PublicVariable)    // OK
    fmt.Println(helper.PublicFunction())  // OK
    // fmt.Println(helper.privateVariable)  // ERROR!
}
```

---

## Package Initialization

Fungsi `init()` dipanggil **otomatis** saat package diakses, sebelum `main()`:

```go
// database/database.go
package database

import "fmt"

var connection string

func init() {
    // dijalankan otomatis saat package di-import
    connection = "MySQL Connected"
    fmt.Println("Database initialized")
}

func GetConnection() string {
    return connection
}
```

```go
// main.go
package main

import (
    "fmt"
    "belajar-golang/database"
)

func main() {
    fmt.Println(database.GetConnection())
    // Output:
    // Database initialized
    // MySQL Connected
}
```

### Blank Identifier

Import package hanya untuk menjalankan `init()`, tanpa menggunakan function-nya:

```go
package main

import (
    "fmt"
    _ "belajar-golang/database"  // import hanya untuk init()
)

func main() {
    fmt.Println("Program jalan")
    // init() dari package database tetap dijalankan
}
```

---

## Error

### error Interface

```go
type error interface {
    Error() string
}
```

### Membuat Error

Gunakan package `errors`:

```go
package main

import (
    "errors"
    "fmt"
)

func pembagian(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("tidak bisa dibagi dengan nol")
    }
    return a / b, nil
}

func main() {
    result, err := pembagian(10, 2)
    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Result:", result)  // 5
    }

    result2, err2 := pembagian(10, 0)
    if err2 != nil {
        fmt.Println("Error:", err2.Error())  // tidak bisa dibagi dengan nol
    } else {
        fmt.Println("Result:", result2)
    }
}
```

### Custom Error

Karena `error` adalah interface, kita bisa membuat struct yang mengimplementasinya:

```go
package main

import "fmt"

// custom error
type ValidationError struct {
    Message string
}

// implementasi interface error
func (v *ValidationError) Error() string {
    return v.Message
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Message: "umur tidak boleh minus"}
    }
    if age > 150 {
        return &ValidationError{Message: "umur terlalu besar"}
    }
    return nil
}

func main() {
    err := validateAge(-1)
    if err != nil {
        fmt.Println("Error:", err.Error())  // umur tidak boleh minus
    }
}
```

### Mengecek Jenis Error

```go
package main

import (
    "errors"
    "fmt"
)

type ValidationError struct {
    Message string
}

func (v *ValidationError) Error() string {
    return v.Message
}

type NotFoundError struct {
    Message string
}

func (n *NotFoundError) Error() string {
    return n.Message
}

func process(value int) error {
    if value < 0 {
        return &ValidationError{Message: "nilai tidak boleh minus"}
    }
    if value > 100 {
        return &NotFoundError{Message: "nilai tidak ditemukan"}
    }
    return nil
}

func main() {
    err := process(-1)
    if err != nil {
        var validErr *ValidationError
        var notFoundErr *NotFoundError

        if errors.As(err, &validErr) {
            fmt.Println("Validation Error:", validErr.Message)
        } else if errors.As(err, &notFoundErr) {
            fmt.Println("Not Found Error:", notFoundErr.Message)
        }
    }
}
```

---

## Materi Selanjutnya

Setelah menguasai dasar Go-Lang, lanjutkan ke:

1. **Go-Lang Standard Library** — package bawaan Go (strings, math, os, dll)
2. **Go-Lang Modules** — manajemen dependency
3. **Go-Lang Unit Test** — testing dengan package `testing`
4. **Go-Lang Goroutine** — concurrency & parallelism
5. **Go-Lang Database** — koneksi ke database SQL
6. **Go-Lang Web** — membuat HTTP server & REST API
