package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ { // Repetimos varias veces para mayor probabilidad de GC
		go triggerGCWithPointerLeak(&wg)
	}
	wg.Wait()

}

func triggerGCWithPointerLeak(wg *sync.WaitGroup) {
	// 1. Crear una gran cantidad de datos para forzar presión de memoria
	data := make([]int, 1_000_000)
	for index, _ := range data {
		data[index] = index
	}

	// 2. Obtenemos un puntero a un elemento de `data`
	p := unsafe.Pointer(&data[0])
	x := uintptr(p) // Convertimos el puntero a uintptr, el cual no será seguido por el GC

	// 3. Sugerimos al garbage collector que libere memoria
	runtime.GC()

	time.Sleep(5 * time.Second) // Damos un poco de tiempo al GC para que actúe

	// Intentamos volver a convertir `x` a unsafe.Pointer y leer el valor
	p2 := unsafe.Pointer(x)
	fmt.Println(*(*int)(p2))

	//si comentamos este print, el GC se da cuenta de que data ya no se usa y por tanto el GC lo libera por lo que la linea anterior
	//da panic debido a que cuando se ejecuta, el GC libero la memoria por no usarse
	fmt.Println(data[0])

	wg.Done()
}
