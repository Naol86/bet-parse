package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ========== PRE-MATCH TYPES ==========


type SP_TableTennis struct {
	MatchLines                     Market `json:"match_lines"`
	FirstGame                      Market `json:"1st_game"`
	FirstGameWinnerAndTotalDouble  Market `json:"1st_game_winner_and_total_double"`
	FirstGameCorrectScore          Market `json:"1st_game_correct_score"`
	FirstGameWinningMargin         Market `json:"1st_game_winning_margin"`
	FirstGameRaceTo                Market `json:"1st_game_race_to"`
	FirstGameLeadAfter             Market `json:"1st_game_lead_after"`
	FirstGameToGoToExtraPoints     Market `json:"1st_game_to_go_to_extra_points"`
	FirstGameTotalOddEven          Market `json:"1st_game_total_odd_even"`
	TotalPointsOddEven             Market `json:"total_points_odd_even"`
}

type PreMatchGame struct {
	UpdatedAt string `json:"updated_at"`
	Key       string `json:"key"`
	SP        SP_TableTennis `json:"sp"`
}

type PreMatchResult struct {
	FI       string        `json:"FI"`
	EventID  string        `json:"event_id"`
	Game     PreMatchGame  `json:"game"`
	Main     PreMatchGame  `json:"main"`
	Match    PreMatchGame  `json:"match"`
	Others   []interface{} `json:"others"`
	Schedule PreMatchGame  `json:"schedule"`
}

type PreMatchResponse struct {
	Success int             `json:"success"`
	Results []PreMatchResult `json:"results"`
}

// ========== POST-MATCH TYPES ==========

type League struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	CC   string `json:"cc"`
}

type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ImageID string `json:"image_id"` // ✅ correct
	CC      string `json:"cc"`
}

type GameScore struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type PostMatchResult struct {
	ID               string         `json:"id"`
	SportID          string         `json:"sport_id"`
	Time             string         `json:"time"`
	TimeStatus       string         `json:"time_status"`
	League           League         `json:"league"`
	Home             Team           `json:"home"`
	Away             Team           `json:"away"`
	SS               string         `json:"ss"`
	InplayCreatedAt  string         `json:"inplay_created_at"`
	InplayUpdatedAt  string         `json:"inplay_updated_at"`
	ConfirmedAt      string         `json:"confirmed_at"`
	Bet365ID        string         `json:"bet365_id"`
}

type PostMatchResponse struct {
	Success int              `json:"success"`
	Results []PostMatchResult `json:"results"`
}

// ========== COMBINED MATCH DATA ==========

type CompleteMatchData struct {
	PreMatch  *PreMatchResponse
	PostMatch *PostMatchResponse
}

// ========== READER FUNCTIONS ==========
func ReadPreMatchData(filename string) (*PreMatchResponse, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var response PreMatchResponse
	err = json.Unmarshal(byteValue, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &response, nil
}

func ReadPostMatchData(filename string) (*PostMatchResponse, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var response PostMatchResponse
	err = json.Unmarshal(byteValue, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &response, nil
}

// ----------- Evaluation ------------

func evaluate1X2(home, away int, selection string) string {
    result := "Draw"
    if home > away {
        result = "1"
    } else if away > home {
        result = "2"
    }

    if selection == result {
        return "✅ Win"
    }
    return "❌ Loss"
}

func evaluateOverUnder(totalGoals int, selection string, line float64) string {
    if selection == "Over" && float64(totalGoals) > line {
        return "✅ Win"
    } else if selection == "Under" && float64(totalGoals) < line {
        return "✅ Win"
    }
    return "❌ Loss"
}

func evaluateCorrectScore(home, away int, selection string) string {
    actual := fmt.Sprintf("%d-%d", home, away)
    if actual == selection {
        return "✅ Win"
    }
    return "❌ Loss"
}

func evaluateDoubleChance(home, away int, selection string) string {
    switch selection {
    case "1X":
			if home >= away {
					return "✅ Win"
			}
    case "X2":
        if away >= home {
            return "✅ Win"
        }
    case "12":
        if home != away {
            return "✅ Win"
        }
    }
    return "❌ Loss"
}

func main() {
	// Load pre-match data
	preMatch, err := ReadPreMatchData("baseball_prematch.json")
	if err != nil {
		fmt.Println("Error reading pre-match data:", err)
		return
	}

	// Load post-match data
	postMatch, err := ReadPostMatchData("baseball_result.json")
	if err != nil {
		fmt.Println("Error reading post-match data:", err)
		return
	}

	// Combine datasets
	matchData := CompleteMatchData{
		PreMatch:  preMatch,
		PostMatch: postMatch,
	}

	if len(matchData.PreMatch.Results) > 0 && len(matchData.PostMatch.Results) > 0 {
		// pre := matchData.PreMatch.Results[0]
		post := matchData.PostMatch.Results[0]

		fmt.Printf("Match Analysis: %s vs %s\n", post.Home.Name, post.Away.Name)
		fmt.Printf("League: %s\n", post.League.Name)
		fmt.Printf("Final Score: %s\n", post.SS)

		// Evaluate 1X2 using final score
		fmt.Println("\n--- 1X2 Evaluation ---")
		selection1X2 := "2" // Sample selection: "1" = Home win, "2" = Away win, "Draw"
		selectionOUGoals := "Under"// Over goals
		ouLine := 2.5// line
		correctScore := "2-3"// exact score
		doubleChance := "12"// double chance

		parts := strings.Split(post.SS, "-")
		if len(parts) == 2 {
			homeScore, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			awayScore, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			totalGoals := homeScore + awayScore

			fmt.Println("Evaluating Sample Selections:")
			fmt.Println("1X2 →", selection1X2, ":", evaluate1X2(homeScore, awayScore, selection1X2))
			fmt.Println("Over/Under", ouLine, "→", selectionOUGoals, ":", evaluateOverUnder(totalGoals, selectionOUGoals, ouLine))
			fmt.Println("Correct Score →", correctScore, ":", evaluateCorrectScore(homeScore, awayScore, correctScore))
			fmt.Println("Double Chance →", doubleChance, ":", evaluateDoubleChance(homeScore, awayScore, doubleChance))
		} else {
			fmt.Println("Invalid SS format:", post.SS)
		}
	}
}