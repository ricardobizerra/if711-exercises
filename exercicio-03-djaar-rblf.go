package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Carro struct {
	Placa string
}

func Exercise2() {
	var lock_ponte sync.Mutex

	var buffer_ponte *Carro
	var buffer_ponte_lock sync.Mutex
	var buffer_esq *Carro
	var buffer_esq_lock sync.Mutex
	var buffer_dir *Carro
	var buffer_dir_lock sync.Mutex
	direcao_atual := ""
	var direcao_atual_lock sync.Mutex
	var wg sync.WaitGroup

	done_esq := false
	var done_esq_lock sync.Mutex
	done_dir := false
	var done_dir_lock sync.Mutex

	// Portão esquerdo
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		index := 0
		carros_chegando := []Carro{}
		carros := []Carro{}
		for i := range 1000 {
			carros = append(carros, Carro{Placa: fmt.Sprintf("AAA%d", i)})
		}
		for index < len(carros) {
			carro := carros[index]

			if lock_ponte.TryLock() {
				direcao_atual_lock.Lock()
				direcao_atual = "esquerda"
				direcao_atual_lock.Unlock()
				buffer_ponte_lock.Lock()
				buffer_ponte = &carro
				buffer_ponte_lock.Unlock()
				index += 1
			} else if buffer_esq_lock.TryLock() {
				carro_vindo := buffer_esq
				buffer_esq = nil
				buffer_esq_lock.Unlock()
				if carro_vindo == nil {
					continue
				}
				carros_chegando = append(carros_chegando, *carro_vindo)
				// fmt.Println("saiu no lado esquerdo", carro_vindo)
				lock_ponte.Unlock()
			}
		}
		for {
			// fmt.Println("esq", len(carros_chegando))
			if len(carros_chegando) == 1000 {
				done_esq_lock.Lock()
				done_esq = true
				done_esq_lock.Unlock()
				break
			}

			if buffer_esq_lock.TryLock() {
				carro_vindo := buffer_esq
				buffer_esq = nil
				buffer_esq_lock.Unlock()
				if carro_vindo == nil {
					continue
				}

				carros_chegando = append(carros_chegando, *carro_vindo)
				// fmt.Println("saiu no lado esquerdo", carro_vindo)
				lock_ponte.Unlock()
			}
		}
	}(&wg)

	// Portão direito
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		index := 0
		carros_chegando := []Carro{}
		carros := []Carro{}
		for i := range 1000 {
			carros = append(carros, Carro{Placa: fmt.Sprintf("BBB%d", i)})
		}
		for index < len(carros) {
			carro := carros[index]

			if lock_ponte.TryLock() {
				direcao_atual_lock.Lock()
				direcao_atual = "direita"
				direcao_atual_lock.Unlock()
				buffer_ponte_lock.Lock()
				buffer_ponte = &carro
				buffer_ponte_lock.Unlock()
				index += 1
			} else if buffer_dir_lock.TryLock() {
				carro_vindo := buffer_dir
				buffer_dir = nil
				buffer_dir_lock.Unlock()
				if carro_vindo == nil {
					continue
				}
				carros_chegando = append(carros_chegando, *carro_vindo)
				// fmt.Println("saiu no lado direito", carro_vindo)
				lock_ponte.Unlock()
			}
		}
		for {
			// fmt.Println("dir", len(carros_chegando))
			if len(carros_chegando) == 1000 {
				done_dir_lock.Lock()
				done_dir = true
				done_dir_lock.Unlock()
				break
			}
			if buffer_dir_lock.TryLock() {

				carro_vindo := buffer_dir
				buffer_dir = nil
				buffer_dir_lock.Unlock()
				if carro_vindo == nil {
					continue
				}
				carros_chegando = append(carros_chegando, *carro_vindo)
				// fmt.Println("saiu no lado direito", carro_vindo)
				lock_ponte.Unlock()
			}
		}
	}(&wg)

	// Ponte
	go func() {
		done_ponte := false
		for {
			if done_ponte {
				return
			}
			if done_esq_lock.TryLock() {
				if done_dir_lock.TryLock() {
					if done_esq && done_dir {
						done_ponte = true
					}
					done_dir_lock.Unlock()
				}
				done_esq_lock.Unlock()
			}
			buffer_ponte_lock.Lock()
			carro_na_ponte := buffer_ponte
			buffer_ponte_lock.Unlock()
			if carro_na_ponte == nil {
				continue
			} else {
				buffer_ponte_lock.Lock()
				buffer_ponte = nil
				buffer_ponte_lock.Unlock()

				// fmt.Println("entrou na ponte pela", direcao_atual, carro_na_ponte)
				// time.Sleep(time.Second)
				direcao_atual_lock.Lock()
				if direcao_atual == "esquerda" {
					buffer_dir_lock.Lock()
					// fmt.Println("saindo da ponte pela direita", carro_na_ponte)
					buffer_dir = carro_na_ponte
					buffer_dir_lock.Unlock()
				} else if direcao_atual == "direita" {
					// fmt.Println("saindo da ponte pela esquerda", carro_na_ponte)
					buffer_esq_lock.Lock()
					buffer_esq = carro_na_ponte
					buffer_esq_lock.Unlock()
				}
				direcao_atual_lock.Unlock()
			}

		}
	}()

	wg.Wait()
}

func Exercise1() {
	ponte_portao_esq_entrada := make(chan Carro, 1)
	ponte_portao_esq_saida := make(chan Carro, 1)
	ponte_portao_dir_entrada := make(chan Carro, 1)
	ponte_portao_dir_saida := make(chan Carro, 1)

	var lock_ponte sync.Mutex

	var wg sync.WaitGroup

	// Portão esquerdo
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		index := 0

		carros_chegando := []Carro{}
		carros := []Carro{}
		for i := range 1000 {
			carros = append(carros, Carro{Placa: fmt.Sprintf("AAA%d", i)})
		}

		for index < len(carros) {
			carro := carros[index]
			select {
			case ponte_portao_esq_entrada <- carro:
				index += 1
			case carro := <-ponte_portao_esq_saida:
				carros_chegando = append(carros_chegando, carro)
				// fmt.Println("saiu no lado esquerdo", carro)
				lock_ponte.Unlock() // Libera a travessia para um novo carro
			}
		}
		for {
			if len(carros_chegando) == 1000 {
				break
			}
			carros_chegando = append(carros_chegando, <-ponte_portao_esq_saida)
			// fmt.Println("saiu no lado esquerdo", carro)
			lock_ponte.Unlock() // Libera a travessia para um novo carro
		}
	}(&wg)

	// Portão direito
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		index := 0

		carros_chegando := []Carro{}
		carros := []Carro{}
		for i := range 1000 {
			carros = append(carros, Carro{Placa: fmt.Sprintf("BBB%d", i)})
		}
		for index < len(carros) {
			carro := carros[index]
			select {
			case ponte_portao_dir_entrada <- carro:
				index += 1
			case carro := <-ponte_portao_dir_saida:
				carros_chegando = append(carros_chegando, carro)
				// fmt.Println("saiu no lado direito", carro)
				lock_ponte.Unlock() // Libera a travessia para um novo carro
			}
		}
		for {
			if len(carros_chegando) == 1000 {
				break
			}
			carros_chegando = append(carros_chegando, <-ponte_portao_dir_saida)
			// fmt.Println("saiu no lado direito", carro)
			lock_ponte.Unlock() // Libera a travessia para um novo carro
		}
	}(&wg)

	// Ponte
	go func() {
		for {
			lock_ponte.Lock()
			select {
			case carro := <-ponte_portao_esq_entrada:
				// fmt.Println("entrou na ponte pela esquerda", carro)
				ponte_portao_dir_saida <- carro // Envio ao portão de saída

			case carro := <-ponte_portao_dir_entrada:
				//fmt.Println("entrou na ponte pela direita", carro)
				ponte_portao_esq_saida <- carro // Envio ao portão de saída
			}
		}
	}()
	wg.Wait()
}

func Media(numeros []float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	soma := 0.0
	for i := range len(numeros) {
		soma += numeros[i]
	}

	return soma / float64(len(numeros))
}

func calcularVariancia(arr []float64, media float64) float64 {
	var somaQuadrados float64
	for _, v := range arr {
		somaQuadrados += (v - media) * (v - media)
	}
	return somaQuadrados / float64(len(arr))
}

func calcularMediana(arr []float64) float64 {
	sort.Float64s(arr)

	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2.0
	}
}

func main() {
	runs_1 := [](float64){}
	for i := range 1000 {
		_ = i
		start_time := time.Now()
		Exercise1()
		run_time := time.Since(start_time)
		runs_1 = append(runs_1, float64(run_time.Milliseconds()))
	}

	fmt.Println("Exercício 1")
	fmt.Println("====================================")
	fmt.Println("Média:    ", Media(runs_1), "ms")
	fmt.Println("Variancia:", calcularVariancia(runs_1, Media(runs_1)), "ms")
	fmt.Println("Mediana:  ", calcularMediana(runs_1), "ms")

	runs_2 := [](float64){}
	for i := range 1000 {
		_ = i
		start_time := time.Now()
		Exercise2()
		run_time := time.Since(start_time)
		runs_2 = append(runs_2, float64(run_time.Milliseconds()))
	}

	fmt.Println()
	fmt.Println("Exercício 2")
	fmt.Println("====================================")
	fmt.Println("Média:    ", Media(runs_2), "ms")
	fmt.Println("Variancia:", calcularVariancia(runs_2, Media(runs_2)), "ms")
	fmt.Println("Mediana:  ", calcularMediana(runs_2), "ms")
}
