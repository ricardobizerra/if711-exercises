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
	var lock_ponte sync.Mutex
	var buffer_ponte *Carro
	var buffer_ponte_lock sync.Mutex
	var buffer_esq *Carro
	var buffer_esq_lock sync.Mutex
	var buffer_dir *Carro
	var buffer_dir_lock sync.Mutex
	direcao_atual := ""
	var wg sync.WaitGroup

	// Portão esquerdo
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		index := 0
		carros := []Carro{Carro{Placa: "AAA1111"}, Carro{Placa: "AAA1112"}, Carro{Placa: "AAA1113"}, Carro{Placa: "AAA1114"}, Carro{Placa: "AAA1115"}}
		for index < len(carros) {
			carro := carros[index]

			if lock_ponte.TryLock() {
				buffer_ponte_lock.Lock()
				buffer_ponte = &carro
				buffer_ponte_lock.Unlock()
				index += 1
				direcao_atual = "esquerda"
			} else if buffer_esq_lock.TryLock() {
				carro_vindo := buffer_esq
				buffer_esq_lock.Unlock()
				buffer_esq = nil
				if carro_vindo == nil {
					continue
				}
				fmt.Println("saiu no lado esquerdo", carro_vindo)
				lock_ponte.Unlock()
			}
		}
		for {
			if buffer_esq_lock.TryLock() {
				carro_vindo := buffer_esq
				buffer_esq_lock.Unlock()
				buffer_esq = nil
				if carro_vindo == nil {
					continue
				}
				fmt.Println("saiu no lado esquerdo", carro_vindo)
				lock_ponte.Unlock()
			}
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

			if lock_ponte.TryLock() {
				buffer_ponte_lock.Lock()
				buffer_ponte = &carro
				buffer_ponte_lock.Unlock()
				index += 1
				direcao_atual = "direita"
			} else if buffer_dir_lock.TryLock() {
				carro_vindo := buffer_dir
				buffer_dir_lock.Unlock()
				buffer_dir = nil
				if carro_vindo == nil {
					continue
				}
				fmt.Println("saiu no lado direito", carro_vindo)
				lock_ponte.Unlock()
			}
		}
		for {
			if buffer_dir_lock.TryLock() {
				carro_vindo := buffer_dir
				buffer_dir_lock.Unlock()
				buffer_dir = nil
				if carro_vindo == nil {
					continue
				}
				fmt.Println("saiu no lado direito", carro_vindo)
				lock_ponte.Unlock()
			}
		}
	}(&wg)

	// Ponte
	go func() {
		for {
			carro_na_ponte := buffer_ponte
			if carro_na_ponte == nil {
				continue
			} else {
				buffer_ponte = nil

				fmt.Println("entrou na ponte pela", direcao_atual, carro_na_ponte)
				time.Sleep(time.Second)
				if direcao_atual == "esquerda" {
					fmt.Println("saindo da ponte pela direita", carro_na_ponte)
					buffer_dir_lock.Lock()
					buffer_dir = carro_na_ponte
					buffer_dir_lock.Unlock()
				} else if direcao_atual == "direita" {
					fmt.Println("saindo da ponte pela esquerda", carro_na_ponte)
					buffer_esq_lock.Lock()
					buffer_esq = carro_na_ponte
					buffer_esq_lock.Unlock()
				}
			}
		}
	}()
	wg.Wait()

}
