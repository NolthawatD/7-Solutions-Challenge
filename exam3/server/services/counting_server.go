package services

import (
	context "context"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type countingServer struct{}

func NewCountingServer() CountingServer {
	return countingServer{}
}

func (countingServer) mustEmbedUnimplementedCountingServer() {}

func (countingServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	if req.Name == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Name is required",
		)
	}

	result := fmt.Sprintf("Hello %v", req.Name)
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

func (countingServer) CountBeef(ctx context.Context, req *CountBeefRequest) (*CountBeefResponse, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	data, err := fetchData(url)
	if err != nil {
		return nil, err
	}

	wordsToCount := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}
	wordCounts := countWords(data, wordsToCount)
	fmt.Printf("beef: %v\n", wordCounts)

	res := CountBeefResponse{
		Beef: wordCounts,
	}

	return &res, nil
}

func fetchData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

func countWords(text string, wordsToCount []string) map[string]int32 {
	wordCounts := make(map[string]int32)
	for _, word := range wordsToCount {
		wordCounts[word] = 0
	}

	// remove punctuation (except hyphens) and convert to lower case
	re := regexp.MustCompile(`[^\w\s-]`)
	normalizedText := re.ReplaceAllString(text, "")
	normalizedText = strings.ToLower(normalizedText)

	words := strings.Fields(normalizedText)
	// fmt.Printf("\nwords: %v", words)

	// count each word
	for _, word := range words {
		// fmt.Printf("\n word : %s", word)
		if _, exists := wordCounts[word]; exists {
			wordCounts[word]++
		}
	}

	return wordCounts
}
