package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(inputs []string) FreqMap {
	comm := make(chan FreqMap)
	defer close(comm)

	for _, s := range inputs {
		go func(s string) {
			comm <- Frequency(s)
		}(s)
	}

	ret := <-comm
	for i := 1; i < len(inputs); i++ {
		for k, v := range <-comm {
			ret[k] += v
		}
	}

	return ret
}
