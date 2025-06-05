package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"MergSortGoLanguage/internal/benchmark"
	"MergSortGoLanguage/internal/reader"
	"MergSortGoLanguage/internal/sort"
	"MergSortGoLanguage/internal/structures"
)

type result struct {
    Size           int
    Structure      string // lista | pilha | fila
    Representation string // linear | dinamica
    DurationS      float64
    MemBytes       uint64
}


func runBenchmark(label string, data []int) (durNS int64, mem uint64) {
	d, m := benchmark.MeasurePerformance(func() { _ = sort.MergeSort(data) })
	return d.Nanoseconds(), m
}

func main() {
    filePath := "data/ratings.csv"
    sizes := []int{100, 1000, 10000, 100000, 1000000}

    var results []result
    const nExec = 10

    for _, size := range sizes {
        src, err := reader.LoadRatings(filePath, size)
        if err != nil {
            log.Fatalf("erro ao ler csv: %v", err)
        }
        fmt.Printf("\n===== %d entradas =====\n", size)

        // -------- Lista (slice)
        var sumDur, sumMem int64
        for i := 0; i < nExec; i++ {
            dur, mem := runBenchmark("slice", src)
            sumDur += dur
            sumMem += int64(mem)
        }
        avgDur := float64(sumDur) / float64(nExec) / 1e9
        avgMem := sumMem / nExec
        results = append(results, result{size, "lista", "linear", avgDur, uint64(avgMem)})
        fmt.Printf("Lista linear        : %.6f s | %d B\n", avgDur, avgMem)

        // -------- Lista encadeada (ordenar vetor e inserir)
        sumDur, sumMem = 0, 0
        for i := 0; i < nExec; i++ {
            vetor := make([]int, len(src))
            copy(vetor, src)
            durSort, _ := runBenchmark("sort for linked list", vetor)
            durInsert, mem := benchmark.MeasurePerformance(func() {
                _ = structures.FromSlice(vetor)
            })
            sumDur += durSort + durInsert.Nanoseconds()
            sumMem += int64(mem)
        }
        avgDur = float64(sumDur) / float64(nExec) / 1e9
        avgMem = sumMem / nExec
        results = append(results, result{size, "lista", "dinamica", avgDur, uint64(avgMem)})
        fmt.Printf("Lista dinâmica      : %.6f s | %d B\n", avgDur, avgMem)

        // -------- Pilha linear (ordenar vetor e inserir)
        sumDur, sumMem = 0, 0
        for i := 0; i < nExec; i++ {
            vetor := make([]int, len(src))
            copy(vetor, src)
            durSort, _ := runBenchmark("sort for stack slice", vetor)
            sl := structures.NewStackLinear()
            durInsert, mem := benchmark.MeasurePerformance(func() {
                for _, v := range vetor { sl.Push(v) }
            })
            sumDur += durSort + durInsert.Nanoseconds()
            sumMem += int64(mem)
        }
        avgDur = float64(sumDur) / float64(nExec) / 1e9
        avgMem = sumMem / nExec
        results = append(results, result{size, "pilha", "linear", avgDur, uint64(avgMem)})
        fmt.Printf("Pilha linear        : %.6f s | %d B\n", avgDur, avgMem)

        // -------- Pilha dinâmica (ordenar vetor e inserir)
        sumDur, sumMem = 0, 0
        for i := 0; i < nExec; i++ {
            vetor := make([]int, len(src))
            copy(vetor, src)
            durSort, _ := runBenchmark("sort for stack dyn", vetor)
            sd := structures.NewStackDynamic()
            durInsert, mem := benchmark.MeasurePerformance(func() {
                for _, v := range vetor { sd.Push(v) }
            })
            sumDur += durSort + durInsert.Nanoseconds()
            sumMem += int64(mem)
        }
        avgDur = float64(sumDur) / float64(nExec) / 1e9
        avgMem = sumMem / nExec
        results = append(results, result{size, "pilha", "dinamica", avgDur, uint64(avgMem)})
        fmt.Printf("Pilha dinâmica      : %.6f s | %d B\n", avgDur, avgMem)

        // -------- Fila linear
        sumDur, sumMem = 0, 0
        for i := 0; i < nExec; i++ {
            ql := structures.NewQueueLinear()
            for _, v := range src { ql.Enqueue(v) }
            dur, mem := runBenchmark("queue slice", ql.ToSlice())
            sumDur += dur
            sumMem += int64(mem)
        }
        avgDur = float64(sumDur) / float64(nExec) / 1e9
        avgMem = sumMem / nExec
        results = append(results, result{size, "fila", "linear", avgDur, uint64(avgMem)})
        fmt.Printf("Fila linear         : %.6f s | %d B\n", avgDur, avgMem)

        // -------- Fila dinâmica
        sumDur, sumMem = 0, 0
        for i := 0; i < nExec; i++ {
            qd := structures.NewQueueDynamic()
            for _, v := range src { qd.Enqueue(v) }
            dur, mem := runBenchmark("queue dyn", qd.ToSlice())
            sumDur += dur
            sumMem += int64(mem)
        }
        avgDur = float64(sumDur) / float64(nExec) / 1e9
        avgMem = sumMem / nExec
        results = append(results, result{size, "fila", "dinamica", avgDur, uint64(avgMem)})
        fmt.Printf("Fila dinâmica       : %.6f s | %d B\n", avgDur, avgMem)
    }

    if err := saveCSV("results.csv", results); err != nil {
        log.Fatalf("erro ao salvar CSV: %v", err)
    }
    fmt.Println("\nResultados salvos em results.csv ✅")
}


func saveCSV(path string, rows []result) error {
    f, err := os.Create(path)
    if err != nil { return err }
    defer f.Close()

    w := csv.NewWriter(f)
    defer w.Flush()

    _ = w.Write([]string{"size", "structure", "representation", "duration_s", "mem_bytes"})
    for _, r := range rows {
        _ = w.Write([]string{
            strconv.Itoa(r.Size),
            r.Structure,
            r.Representation,
            fmt.Sprintf("%.6f", r.DurationS),
            strconv.FormatUint(r.MemBytes, 10),
        })
    }
    return w.Error()
}