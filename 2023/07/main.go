package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return getScore(parseInput(1, input))
}

func part2(input []string) any {
	return getScore(parseInput(2, input))
}

func getScore(players []hand) int {
	total := 0
	sort.SliceStable(players, func(i, j int) bool {
		p1 := players[i]
		p2 := players[j]
		if p1.handStrength != p2.handStrength {
			return p1.handStrength < p2.handStrength
		}
		for i := 0; i < 5; i++ {
			if p1.cardsStrength[i] == p2.cardsStrength[i] {
				continue
			}
			return p1.cardsStrength[i] < p2.cardsStrength[i]
		}
		return false
	})
	for i, player := range players {
		total += player.bid * (i + 1)
	}
	return total
}

type hand struct {
	handStrength  HandStrength
	cardsStrength [5]int
	bid           int
}

type HandStrength int

const (
	HighCard HandStrength = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func computeHandStrength(part int, cards string) HandStrength {
	occ := map[uint8]int{}
	for _, card := range cards {
		occ[uint8(card)]++
	}
	type occurrence struct {
		card  uint8
		count int
	}
	var sortedOccurrences []occurrence
	for k, v := range occ {
		sortedOccurrences = append(sortedOccurrences, occurrence{card: k, count: v})
	}
	sort.SliceStable(sortedOccurrences, func(i, j int) bool {
		return sortedOccurrences[i].count > sortedOccurrences[j].count
	})
	if sortedOccurrences[0].count == 5 {
		return FiveOfAKind // quick exit
	}
	if part == 2 {
		var jokerReplacementIdx int
		for i := 0; i < len(sortedOccurrences); i++ {
			if sortedOccurrences[i].card != 'J' {
				jokerReplacementIdx = i
				break
			}
		}
		sortedOccurrences[jokerReplacementIdx].count += occ['J']
	}

	handStrength := HighCard
	for _, v := range sortedOccurrences {
		if part == 2 && v.card == 'J' {
			continue
		}
		switch v.count {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			if handStrength == OnePair {
				return FullHouse
			}
			handStrength = ThreeOfAKind
		case 2:
			if handStrength == ThreeOfAKind {
				return FullHouse
			}
			if handStrength == OnePair {
				return TwoPair
			}
			handStrength = OnePair
		}
	}
	return handStrength
}

func parseInput(part int, input []string) []hand {
	headCardsStrengthMap := map[rune]int{'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	if part == 2 {
		headCardsStrengthMap['J'] = 1
	}
	var players []hand
	for _, line := range input {
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		cardStrengths := [5]int{}
		for i, card := range parts[0] {
			s := int(card - 48)
			if card >= 65 {
				s = headCardsStrengthMap[card]
			}
			cardStrengths[i] = s
		}
		players = append(players, hand{
			handStrength:  computeHandStrength(part, parts[0]),
			cardsStrength: cardStrengths,
			bid:           bid,
		})
	}
	return players
}
