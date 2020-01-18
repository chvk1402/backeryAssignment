package order

import (
	"backery/config"
	"backery/structs"
	"fmt"
	"sort"
)

func ProcessOrder(code string, quantity int) (*structs.OrderResp, error) {
	var packs []structs.Price
	totalPrice := float32(0)
	//get priceMap for the code
	priceMap := config.PriceMatrix[code]
	// sort the price
	sortedQuantites := sortedKeys(priceMap)
	lenAr := len(sortedQuantites) - 1
	preVal := quantity
	for _, qty := range sortedQuantites {
		if qty > preVal {
			continue
		}
		rem := preVal % qty
		if rem == 0 {
			totalPrice += priceMap[qty] * float32(preVal/qty)
			packs = append(packs, structs.Price{
				Pack:   qty,
				QtySet: preVal/qty,
				Price:  priceMap[qty],
			})
			preVal = 0
			break
		}
		if rem >= sortedQuantites[lenAr] {
			totalPrice += priceMap[qty] * float32(preVal/qty)
			packs = append(packs, structs.Price{
				Pack:   qty,
				QtySet: preVal/qty,
				Price:  priceMap[qty],
			})
			preVal = rem
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