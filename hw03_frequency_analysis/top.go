package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[\p{L}^-]+`)

func additionalTask(inputWord string) []string {
	inputWord = strings.ToLower(inputWord)
	/*
		rezult := re.FindAllString(strings.ToLower(inputWord), -1)

		if len(rezult) != 1 && rezult[1] != "-" {

			fmt.Println("ERRRRROOOORRRRR", "len = ", len(rezult), "  rez = ", rezult, "   inputWord = ", inputWord)
			fmt.Printf("%q", rezult)
			return "", errors.New("additionalTask: problem with input data")
		}
	*/
	return re.FindAllString(strings.ToLower(inputWord), -1)
}

const rezultLen int = 10

type wordCountPair struct {
	word  string
	count uint64
}

func Top10(inputData string) []string {
	// test()
	// return nil

	// inputData = "dog,two"
	if len(inputData) == 0 {
		return nil
	}

	// inputData = "a  c b   a  g bc"
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
