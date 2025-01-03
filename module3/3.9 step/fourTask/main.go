package main

// Вам необходимо написать функцию calculator следующего вида:
//
// func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
// В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.
//
// Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.
//
// Функция calculator должна быть неблокирующей, сразу возвращая управление.
//
// Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.
func main() {

}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {

	mk := make(chan int)
	res := 0

	go func() {
		defer close(mk)
		for {
			select {
			case v := <-arguments:
				res += v

			case <-done:
				mk <- res
				return
			}
		}

	}()
	return mk
}
