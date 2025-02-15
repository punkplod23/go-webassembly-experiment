package hashchecker

import (
	"strings"
	"sync"
)

var (
	// special characters
	specialChars    = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '"', '\'', '<', '>', ',', '.', '?', '/'}
	vowels          = []rune{'a', 'e', 'i', 'o', 'u'}
	constants       = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}
	numberChars     = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	frequency       = []rune{'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'c', 'u', 'm', 'w', 'f', 'g', 'y', 'p', 'b', 'v', 'k', 'j', 'x', 'q', 'z'}
	frequencyUpper  = []rune{'E', 'T', 'A', 'O', 'I', 'N', 'S', 'H', 'R', 'D', 'L', 'C', 'U', 'M', 'W', 'F', 'G', 'Y', 'P', 'B', 'V', 'K', 'J', 'X', 'Q', 'Z'}
	frequencyNo     = []rune{'4', '8', '3', '0', '1', '5', '7', '9', '6'}
	frequencyAll    = append(frequency, frequencyUpper...)
	specialCharsSub = map[rune]rune{
		'a': '@',
		'i': '!',
		'o': '0',
		'e': '3',
		's': '$',
		'l': '1',
	}
	hash = ""
)

func generateCandidates(length int, chars []rune, batchSize int, process func([]string)) {
	candidates := make([]string, 0, batchSize)
	var mu sync.Mutex
	var wg sync.WaitGroup

	batchChan := make(chan []string, 10)
	done := make(chan struct{})

	// Worker goroutines to process batches
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range batchChan {
				process(batch)
			}
		}()
	}

	var generate func(string, int)
	generate = func(current string, remaining int) {
		if remaining == 0 {
			mu.Lock()

			candidates = append(candidates, current)

			if len(candidates) >= batchSize {
				batch := make([]string, len(candidates))
				copy(batch, candidates)
				candidates = candidates[:0]
				mu.Unlock()
				batchChan <- batch
			} else {
				mu.Unlock()
			}
			return
		}
		for _, char := range chars {
			generate(current+string(char), remaining-1)
		}
	}

	generate("", length)

	mu.Lock()
	if len(candidates) > 0 {
		batch := make([]string, len(candidates))
		copy(batch, candidates)
		candidates = candidates[:0]
		mu.Unlock()
		batchChan <- batch
	} else {
		mu.Unlock()
	}

	close(batchChan)
	wg.Wait()
	close(done)
}

func compareCandidate(candidate string, check string) bool {

	candidateUppercase := strings.ToUpper(candidate)
	if checkLogic(check, candidate) {
		return true
	}
	if checkLogic(check, candidateUppercase) {
		return true
	}

	if checkUppercase(candidate, candidateUppercase, check) {
		return true
	}

	if checkSpecialChars(candidate, candidateUppercase, check) {
		return true
	}

	if checkNumberSubstitution(candidate, candidateUppercase, check) {
		return true
	}
	return false
}

func checkUppercase(candidate string, candidateUppercase string, check string) bool {
	for i := 1; i <= len(candidate); i++ {
		value := strings.ToUpper(candidate[:i]) + candidate[i:]
		valueUppercase := strings.ToUpper(candidateUppercase[:i]) + candidateUppercase[i:]
		if checkLogic(check, value) {
			return true
		}
		if checkLogic(check, valueUppercase) {
			return true
		}
		if checkSpecialChars(value, candidateUppercase, check) {
			return true
		}
	}
	return false
}

func checkSpecialChars(candidate string, candidateUppercase string, check string) bool {
	// Special character substitution
	for i := 0; i < len(candidate); i++ {
		if special, ok := specialCharsSub[rune(candidate[i])]; ok {
			value := candidate[:i] + string(special) + candidate[i+1:]
			valueUppercase := candidateUppercase[:i] + string(special) + candidateUppercase[i+1:]
			if checkLogic(check, value) {
				return true
			}
			if checkLogic(check, valueUppercase) {
				return true
			}
			if checkNumberSubstitution(value, candidateUppercase, check) {
				return true
			}
		}
	}
	return false
}

func checkNumberSubstitution(candidate string, candidateUppercase string, check string) bool {
	for i := 0; i < len(candidate); i++ {
		if candidate[i] >= '0' && candidate[i] <= '9' {
			value := candidate[:i] + string(candidate[i]) + candidate[i+1:]
			valueUppercase := candidateUppercase[:i] + string(candidateUppercase[i]) + candidateUppercase[i+1:]
			if checkLogic(check, value) {
				return true
			}
			if checkLogic(check, valueUppercase) {
				return true
			}
		}
	}
	return false
}

func checkFrequency(check string, chars []rune, option string, batchSize int, charLen int) string {
	found := ""
	var mu sync.Mutex

	processBatch := func(batch []string) {
		for _, candidate := range batch {
			if compareCandidate(candidate, check) {
				mu.Lock()
				found = candidate
				mu.Unlock()
				return
			}
		}
	}

	generateCandidates(charLen, chars, batchSize, processBatch)
	return found
}

func checkLogic(a string, b string) bool {

	//original := b
	hashStr := b

	if hashFunc, exists := hashFunctions[hash]; exists {
		hashStr = hashFunc(hashStr)
	}

	if a == hashStr {
		return true
	}

	return false
}

func BruteForce(hashType string, check string, batchSize int, charLen int) string {
	hash = hashType
	return checkFrequency(check, frequency, "", batchSize, charLen)
}
