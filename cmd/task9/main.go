package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Количество входных каналов
const chanCount = 3

// Объединение каналов
func merge(channels ...chan int) chan int {
	mergeChanel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	go func() {
		wg.Wait()
		fmt.Println("merge chan close")
		close(mergeChanel)
	}()

	for _, chanel := range channels {
		go func() {
			for value := range chanel {
				mergeChanel <- value
			}
			wg.Done()
		}()

	}
	return mergeChanel
}

// Вспомогательная функция для демонстрации работы
func getRandomSlice(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(100)
	}
	return result
}
func main() {
	//Создание слайса каналов
	chanels := make([]chan int, chanCount)
	for index := range chanels {
		chanels[index] = make(chan int)
	}
	//Объединение каналов в один
	mergeChan := merge(chanels...)

	exit := make(chan struct{})
	once := sync.Once{}
	closeExit := func() {
		close(exit)
	}

	//Демонстрация работы
	//Создание горутины для каждого из каналов. В горутине создается слайс переменной длинны.
	// Горутина отправляет в канал значения из слайса до тех пор, пока они не закончатся или любая другая из горутин не завершит работу
	for index := range chanels {

		go func(index int, chanel chan int) {
			data := getRandomSlice(rand.Intn(10))
			fmt.Printf("Chanel %d data: %v\n", index, data)

			defer close(chanel)
			defer once.Do(closeExit)

			for {
				select {
				case <-time.After(time.Second):
					if len(data) > 0 {
						chanel <- data[len(data)-1]
						data = data[:len(data)-1]
					} else {
						fmt.Printf("Chanel %d is empty\n", index)
						return
					}
				case <-exit:
					fmt.Printf("Chanel %d stop by exit\n", index)
					return
				}
			}
		}(index, chanels[index])

	}
	//Вывод на экран значений из объединенного канала, пока он открыт
	for value := range mergeChan {
		fmt.Println(value)
	}

}
