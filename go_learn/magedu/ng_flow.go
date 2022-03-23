package magedu

func NgFlow(count int, dict map[int]float64) map[int]float64 {
	var new_dict map[int]float64
	new_value := float64(0)
	for k := range dict {
		value := float64(k/count) * dict[k]
		new_value += value
		new_dict[k] = value
	}
	for k := range dict {
		new_dict[k] = new_dict[k] / new_value
	}
	return new_dict
}
