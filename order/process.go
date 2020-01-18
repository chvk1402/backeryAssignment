package order

import (
	"fmt"
	"github.com/backery/config"
	"github.com/backery/structs"
	"sort"
)

func ProcessOrder(code string, quantity int) (*structs.OrderResp, error) {
	var packs []structs.Price
	totalPrice := float32(0)
	//get priceMap for the code
	priceMap := config.PriceMatrix[code]
	// sort the price
	sortedQuantites := sortedKeys(priceMap)
	preVal := quantity
	for _, qty := range sortedQuantites {
		q := preVal % qty
		r := preVal / qty
		if preVal == quantity || q > qty || q == 0 {
			totalPrice += priceMap[qty] * float32(r)
			packs = append(packs, structs.Price{
				Pack:   qty,
				QtySet: r,
				Price:  priceMap[qty],
			})
			if q == 0 {
				preVal = 0
				break
			}
			//update preVal with current value
			preVal = q
		}
	}
	if preVal == 0 {
		return &structs.OrderResp{
			TotalPrice: totalPrice,
			Code:       code,
			Packs:      packs,
		}, nil
	}
	return nil, fmt.Errorf("the given order cannot be divided in to available packet sets")
}

func sortedKeys(m map[int]float32) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}
