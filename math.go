package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Entrer deux parametre")
        return
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    var data []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        valeur := 0
        fmt.Sscanf(scanner.Text(), "%d", &valeur)
        data = append(data, valeur)
    }

    somme := 0
    for _, valeur := range data {
        somme += valeur
    }
    average := int(math.Round(float64(somme) / float64(len(data))))

    sort.Ints(data)
    milieu := len(data) / 2
    var median int
    if len(data)%2 == 0 {
        median = int(math.Round(float64(data[milieu-1]+data[milieu]) / 2))
    } else {
        median = data[milieu]
    }

    variance := 0.0
    for _, valeur := range data {
        variance += math.Pow(float64(valeur-average), 2)
    }
    variance = variance / float64(len(data))

    deviation := int(math.Round(math.Sqrt(variance)))

    fmt.Printf("Average: %d\n", average)
    fmt.Printf("Median: %d\n", median)
    fmt.Printf("Variance: %d\n", int(math.Round(variance)))
    fmt.Printf("Standard Deviation: %d\n", deviation)
}
