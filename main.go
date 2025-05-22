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
	Size          int
	Structure     string // lista | pilha | fila
	Representation string // linear | dinamica
	DurationNS    int64
	MemBytes      uint64
}

func runBenchmark(label string, data []int) (durNS int64, mem uint64) {
	d, m := benchmark.MeasurePerformance(func() { _ = sort.MergeSort(data) })
	return d.Nanoseconds(), m
}

func main() {
	filePath := "data/ratings.csv"               // caminho relativo recomendado
	sizes := []int{100, 1000, 10000, 100000, 1000000}

	var results []result

	for _, size := range sizes {
		src, err := reader.LoadRatings(filePath, size)
		if err != nil {
			log.Fatalf("erro ao ler csv: %v", err)
		}
		fmt.Printf("\n===== %d entradas =====\n", size)

		// -------- Lista (slice)
		dur, mem := runBenchmark("slice", src)
		results = append(results, result{size, "lista", "linear", dur, mem})
		fmt.Printf("Lista linear        : %v ns | %d B\n", dur, mem)

		// -------- Lista encadeada
		ll := structures.FromSlice(src)
		dur, mem = runBenchmark("linked list", ll.ToSlice())
		results = append(results, result{size, "lista", "dinamica", dur, mem})
		fmt.Printf("Lista dinâmica      : %v ns | %d B\n", dur, mem)

		// -------- Pilha linear
		sl := structures.NewStackLinear()
		for _, v := range src { sl.Push(v) }
		dur, mem = runBenchmark("stack slice", sl.ToSlice())
		results = append(results, result{size, "pilha", "linear", dur, mem})
		fmt.Printf("Pilha linear        : %v ns | %d B\n", dur, mem)

		// -------- Pilha dinâmica
		sd := structures.NewStackDynamic()
		for _, v := range src { sd.Push(v) }
		dur, mem = runBenchmark("stack dyn", sd.ToSlice())
		results = append(results, result{size, "pilha", "dinamica", dur, mem})
		fmt.Printf("Pilha dinâmica      : %v ns | %d B\n", dur, mem)

		// -------- Fila linear
		ql := structures.NewQueueLinear()
		for _, v := range src { ql.Enqueue(v) }
		dur, mem = runBenchmark("queue slice", ql.ToSlice())
		results = append(results, result{size, "fila", "linear", dur, mem})
		fmt.Printf("Fila linear         : %v ns | %d B\n", dur, mem)

		// -------- Fila dinâmica
		qd := structures.NewQueueDynamic()
		for _, v := range src { qd.Enqueue(v) }
		dur, mem = runBenchmark("queue dyn", qd.ToSlice())
		results = append(results, result{size, "fila", "dinamica", dur, mem})
		fmt.Printf("Fila dinâmica       : %v ns | %d B\n", dur, mem)
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

	_ = w.Write([]string{"size", "structure", "representation", "duration_ns", "mem_bytes"})
	for _, r := range rows {
		_ = w.Write([]string{
			strconv.Itoa(r.Size),
			r.Structure,
			r.Representation,
			strconv.FormatInt(r.DurationNS, 10),
			strconv.FormatUint(r.MemBytes, 10),
		})
	}
	return w.Error()
}
