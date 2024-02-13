// Пакет MinCoins содержит функцию для определения минимального набора монет, 
// которыми можно набрать некоторую сумму из полученного набора номиналов монет
package MinCoins

import "sort"

// MinCoins - принимает на вход сумму val и номиналы монет coins. Функция определяет минимальный по
// количеству набор монет, которыми в точности можно набрать указанную сумму val.
func MinCoins(val int, coins []int) []int {
	// Проверяется значение val на валидность; проверяется, что слайс coins не пустой
	if val <= 0 || len(coins) == 0 || coins == nil {
		return nil // Return empty slice
	}
	// Удаляются отрицательные и повторяющиеся значения из монет
	seen := make(map[int]bool) 
	i := 0                     
	for _, coin := range coins {
		if coin > 0 && !seen[coin] {
			seen[coin] = true
			coins[i] = coin
			i++              
		}
	}
	coins = coins[:i]
	// Сортировка монет в порядке возрастания
	sort.Ints(coins)
	// Использование динамического программирования
	dp := make([]int, val+1)
	for i := 1; i <= val; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	// Если dp[val] по-прежнему равен MaxInt32, нет никакого способа вычислить значение, используя данные монеты
	if dp[val] == math.MaxInt32 {
		return nil
	}
	// Использование слайса dp для построения списка используемых монет
	res := make([]int, 0)
	i = val
	for i > 0 {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 == dp[i] {
				res = append(res, coin)
				i -= coin
				break
			}
		}
	}
	// Переворачиваем список использованных монет и возвращаем его
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
