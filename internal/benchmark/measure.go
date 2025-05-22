package benchmark

import (
    "runtime"
    "time"
)

// MeasurePerformance mede o tempo de execução e a memória usada por uma função
func MeasurePerformance(task func()) (time.Duration, uint64) {
    var m1, m2 runtime.MemStats
    runtime.GC() // Força GC para medições mais precisas
    runtime.ReadMemStats(&m1)

    start := time.Now()
    task()
    duration := time.Since(start)

    runtime.ReadMemStats(&m2)
    memoryUsed := m2.Alloc - m1.Alloc

    return duration, memoryUsed
}
