package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[\p{L}^-]+`)

func additionalTask(inputWord string) []string {
	return re.FindAllString(strings.ToLower(inputWord), -1)
}

const rezultLen int = 10

type wordCountPair struct {
	word  string
	count uint64
}

func Top10(inputData string) []string {
	if len(inputData) == 0 {
		return nil
	}

	wordsSlice := strings.Fields(inputData)
	wordCountMap := make(map[string]uint64)

	// fmt.Printf("%q\n", wordsSlice)
	for _, word := range wordsSlice {
		for _, word := range additionalTask(word) {
			if word == "-" {
				continue
			}
			wordCountMap[word]++
		}
	}

	pairSlice := make([]wordCountPair, 0, len(wordCountMap))

	for key, val := range wordCountMap {
		// fmt.Println("key = ", key, "  val = ", val)
		pairSlice = append(pairSlice, wordCountPair{key, val})
	}

	// Можно сделать без слайса структур. нужен только слайс слов и мапа слово - счётчик.
	// В функции сортировки слайса используем замыкание на мапу, и в итоге получим отсортированный
	// слайс слов, где первые N - это и есть TopN.
	// Но сложность алгоритма не должна измениться, мне все-равно надо создавать слайс слов так как
	// они изменились
	// Вообще можно создать функцию возвращающую топN элементов и не сортирующую весь массив. Но что-то пока-что сложно((
	sort.SliceStable(pairSlice, func(i, j int) bool {
		if pairSlice[i].count != pairSlice[j].count {
			return pairSlice[i].count > pairSlice[j].count
		}
		return pairSlice[i].word < pairSlice[j].word
	})

	rezultSlice := make([]string, 0, rezultLen)
	for i := 0; i < rezultLen && i < len(pairSlice); i++ {
		rezultSlice = append(rezultSlice, pairSlice[i].word)
	}
	return rezultSlice
}
