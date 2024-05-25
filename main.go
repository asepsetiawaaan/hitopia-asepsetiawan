package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/asepsetiawaaan/hitopia-asepsetiawan/brackets"
	"github.com/asepsetiawaaan/hitopia-asepsetiawan/palindrome"
	"github.com/asepsetiawaaan/hitopia-asepsetiawan/weightedstrings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose a question to run:")
		fmt.Println("1. Weighted Strings")
		fmt.Println("2. Balanced Bracket")
		fmt.Println("3. Highest Palindrome")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			runWeightedStrings(reader)
		case 2:
			runBalancedBracket(reader)
		case 3:
			runHighestPalindrome(reader)
		case 4:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func runWeightedStrings(reader *bufio.Reader) {
	fmt.Print("Enter the string: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	fmt.Print("Enter the queries (space-separated): ")
	queryInput, _ := reader.ReadString('\n')
	queryInput = strings.TrimSpace(queryInput)
	queryStrings := strings.Split(queryInput, " ")
	queries := make([]int, len(queryStrings))
	for i, qs := range queryStrings {
		queries[i], _ = strconv.Atoi(qs)
	}

	weights := weightedstrings.ComputeWeights(s)
	results := weightedstrings.CheckQueries(weights, queries)
	fmt.Print("Output : ")
	fmt.Println(results)
}

func runBalancedBracket(reader *bufio.Reader) {
	fmt.Print("Enter the string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := brackets.IsBalanced(input)
	fmt.Println("Output : " + result)
}

func runHighestPalindrome(reader *bufio.Reader) {
	fmt.Print("Enter the string: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	fmt.Print("Enter the number of replacements (k): ")
	kStr, _ := reader.ReadString('\n')
	kStr = strings.TrimSpace(kStr)
	k, err := strconv.Atoi(kStr)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid number.")
		return
	}

	result := palindrome.HighestPalindrome(s, k)
	fmt.Println("Output : " + result)
}
