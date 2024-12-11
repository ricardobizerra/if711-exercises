package main

import (
	"fmt"
	"sync"
	"time"
)

type Carro struct {
	Placa string
}

func main() {
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
		carros := []Carro{Carro{Placa: "AAA1111"}, Carro{Placa: "AAA1112"}, Carro{Placa: "AAA1113"}, Carro{Placa: "AAA1114"}, Carro{Placa: "AAA1115"}}
		for index < len(carros) {
			carro := carros[index]
			select {
			case ponte_portao_esq_entrada <- carro:
				index += 1
			case carro := <-ponte_portao_esq_saida:
				fmt.Println("saiu no lado esquerdo", carro)
				lock_ponte.Unlock() // Libera a travessia para um novo carro
			}
		}
		for {
			carro := <-ponte_portao_esq_saida
			fmt.Println("saiu no lado esquerdo", carro)
			lock_ponte.Unlock() // Libera a travessia para um novo carro
		}
	}(&wg)

	// Portão direito
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		index := 0
		carros := []Carro{Carro{Placa: "BBB1111"}, Carro{Placa: "BBB1112"}, Carro{Placa: "BBB1113"}}
		for index < len(carros) {
			carro := carros[index]
			select {
			case ponte_portao_dir_entrada <- carro:
				index += 1
			case carro := <-ponte_portao_dir_saida:
				fmt.Println("saiu no lado direito", carro)
				lock_ponte.Unlock() // Libera a travessia para um novo carro
			}
		}
		for {
			carro := <-ponte_portao_dir_saida
			fmt.Println("saiu no lado direito", carro)
			lock_ponte.Unlock() // Libera a travessia para um novo carro
		}
	}(&wg)

	// Ponte
	go func() {
		for {
			lock_ponte.Lock()
			select {
			case carro := <-ponte_portao_esq_entrada:
				fmt.Println("entrou na ponte pela esquerda", carro)
				time.Sleep(2 * time.Second)     // Tempo de travessia
				ponte_portao_dir_saida <- carro // Envio ao portão de saída

			case carro := <-ponte_portao_dir_entrada:
				fmt.Println("entrou na ponte pela direita", carro)
				time.Sleep(2 * time.Second)     // Tempo de travessia
				ponte_portao_esq_saida <- carro // Envio ao portão de saída
			}
		}
	}()
	wg.Wait()

}
