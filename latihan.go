package main
import "fmt"

func primaUnik(n int) int {
	var hasil, i int
	hasil = 0
	i = 2
	for i <= n {
		if n%i == 0 {
			// Cek apakah i adalah bilangan prima
			isPrima := true
			for j := 2; j*j <= i; j++ {
				if i%j == 0 {
					isPrima = false
					break
				}
			}
			if isPrima {
				hasil += i
				// Hapus semua faktor i agar tidak terulang
				for n%i == 0 {
					n /= i
				}
			}
		}
		i++
	}
	return hasil
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Print(primaUnik(num))
}