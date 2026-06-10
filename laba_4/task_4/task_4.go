package main

import (
	"fmt"
	"sort"
)

type VoterPreference struct {
	ranking []string
	count   int
}

func bordaCount(voters []VoterPreference, candidates []string) (string, map[string]int) {
	scores := make(map[string]int)
	for _, cand := range candidates {
		scores[cand] = 0
	}

	n := len(candidates)

	for _, voter := range voters {
		for pos, candidate := range voter.ranking {
			points := n - 1 - pos
			scores[candidate] += points * voter.count
		}
	}

	fmt.Println("\n=== РЕЗУЛЬТАТЫ ПО МЕТОДУ БОРДА ===")
	fmt.Println("----------------------------------------")
	fmt.Printf("%-15s %-10s\n", "Кандидат", "Баллы")
	fmt.Println("----------------------------------------")

	winner := ""
	maxScore := -1

	for _, cand := range candidates {
		fmt.Printf("%-15s %-10d\n", cand, scores[cand])
		if scores[cand] > maxScore {
			maxScore = scores[cand]
			winner = cand
		}
	}

	fmt.Println("----------------------------------------")
	return winner, scores
}

func condorcetMethod(voters []VoterPreference, candidates []string) string {
	n := len(candidates)
	pairwiseWins := make(map[[2]string]int)

	fmt.Println("\n=== ПАРНЫЕ СРАВНЕНИЯ ===")

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			candA := candidates[i]
			candB := candidates[j]
			votesA := 0
			votesB := 0

			for _, voter := range voters {
				posA := -1
				posB := -1
				for k, cand := range voter.ranking {
					if cand == candA {
						posA = k
					}
					if cand == candB {
						posB = k
					}
				}

				if posA < posB {
					votesA += voter.count
				} else {
					votesB += voter.count
				}
			}

			pairwiseWins[[2]string{candA, candB}] = votesA
			pairwiseWins[[2]string{candB, candA}] = votesB

			winner := ""
			if votesA > votesB {
				winner = candA
			} else if votesB > votesA {
				winner = candB
			} else {
				winner = "ничья"
			}

			fmt.Printf("%s vs %s: %d - %d (%s)\n", candA, candB, votesA, votesB, winner)
		}
	}

	fmt.Println("\n=== ПРОВЕРКА ПОБЕДИТЕЛЯ КОНДОРСЕ ===")

	condorcetWinner := ""
	for _, cand := range candidates {
		wins := 0
		for _, other := range candidates {
			if cand != other {
				if pairwiseWins[[2]string{cand, other}] > pairwiseWins[[2]string{other, cand}] {
					wins++
				}
			}
		}

		fmt.Printf("%s выигрывает в %d из %d парных сравнений\n", cand, wins, n-1)

		if wins == n-1 {
			condorcetWinner = cand
		}
	}

	return condorcetWinner
}

func printPreferences(voters []VoterPreference) {
	fmt.Println("\n=== ПРЕДПОЧТЕНИЯ ИЗБИРАТЕЛЕЙ ===")
	fmt.Println("--------------------------------------------------")

	for _, voter := range voters {
		fmt.Printf("%d избирателей: ", voter.count)
		for i, cand := range voter.ranking {
			fmt.Print(cand)
			if i < len(voter.ranking)-1 {
				fmt.Print(" > ")
			}
		}
		fmt.Println()
	}

	fmt.Println("--------------------------------------------------")
}

func main() {
	var n int
	fmt.Print("Введите количество кандидатов (n): ")
	fmt.Scan(&n)

	candidates := make([]string, n)
	fmt.Println("\nВведите имена кандидатов:")
	for i := 0; i < n; i++ {
		fmt.Printf("Кандидат %d: ", i+1)
		fmt.Scan(&candidates[i])
	}

	var k int
	fmt.Print("\nВведите количество групп избирателей (k): ")
	fmt.Scan(&k)

	voters := make([]VoterPreference, k)
	fmt.Println("\nВведите предпочтения для каждой группы:")

	for i := 0; i < k; i++ {
		fmt.Printf("\nГруппа %d:\n", i+1)
		fmt.Print("  Количество избирателей: ")
		fmt.Scan(&voters[i].count)

		fmt.Print("  Рейтинг (через пробел, от лучшего к худшему): ")
		voters[i].ranking = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&voters[i].ranking[j])
		}
	}

	printPreferences(voters)

	bordaWinner, _ := bordaCount(voters, candidates)
	fmt.Printf("\n   ПОБЕДИТЕЛЬ ПО БОРДА: %s\n", bordaWinner)

	fmt.Println()
	condorcetWinner := condorcetMethod(voters, candidates)
	fmt.Println()

	if condorcetWinner != "" {
		fmt.Printf("   ПОБЕДИТЕЛЬ ПО КОНДОРСЕ: %s\n", condorcetWinner)
	} else {
		fmt.Println("   ПОБЕДИТЕЛЬ ПО КОНДОРСЕ НЕ ОПРЕДЕЛЁН")
		fmt.Println("	(Парадокс Кондорсе: нет кандидата, выигрывающего")
		fmt.Println("		во всех парных сравнениях)")
	}

	if bordaWinner == condorcetWinner {
		fmt.Printf("   Оба метода дали одинаковый результат: %s\n", bordaWinner)
	} else if condorcetWinner == "" {
		fmt.Println("   Методы дали разные результаты:")
		fmt.Printf(" 	-  Борда: %s\n", bordaWinner)
		fmt.Println(" 	-  Кондорсе: победитель не определён")
	} else {
		fmt.Println("   Методы дали разные результаты:")
		fmt.Printf(" - Борда: %s\n", bordaWinner)
		fmt.Printf("   -  Кондорсе: %s\n", condorcetWinner)
	}

	_ = sort.StringsAreSorted
}
