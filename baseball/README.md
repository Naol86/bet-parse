Here's a revised version of your README adapted specifically for **baseball** bet evaluation:

---

# ⚾ Baseball Bet Evaluation Tool

This Go program evaluates **baseball betting outcomes** by comparing pre-match betting data with actual post-match results.

## 🔍 Features

- Parses **pre-match betting data** from `baseball_prematch.json`
- Parses **post-match results** from `baseball_result.json`
- Evaluates several common baseball bet types:

  - ✅ Moneyline (Home/Away)
  - ✅ Over/Under (Total Runs)
  - ✅ Correct Score
  - ✅ Run Line (Handicap)

- Provides **clear win/loss indicators** for each evaluated bet

---

## ⚙️ Prerequisites

- [Go 1.24+](https://golang.org/dl/) installed
- Two JSON data files:

  - `baseball_prematch.json` – contains pre-match betting odds
  - `baseball_result.json` – contains post-match scores and outcomes

---

## 🛠 Installation

1. Clone or download this repository:

   ```bash
   git clone https://github.com/yourusername/baseball-bet-evaluator.git
   cd baseball-bet-evaluator
   ```

2. Ensure both JSON files are placed in the root directory:

   - `baseball_prematch.json`
   - `baseball_result.json`

---

## 🚀 Usage

Run the tool using:

```bash
go run *.go
```

You’ll see output like:

- Match details (teams, league, date)
- Final score
- Evaluation results:

  - ✔ Win or ✖ Loss for each bet type

---

## 🎯 Customizing Bets

You can modify the sample selections in the `main()` function to test specific bets:

- **Moneyline Bet:**

  ```go
  selectionMoneyline := "home"  // or "away"
  ```

- **Over/Under Bet:**

  ```go
  selectionOUGoals := "Over"
  ouLine := 7.5
  ```

- **Correct Score Bet:**

  ```go
  correctScore := "8-3"
  ```

- **Run Line (Handicap) Bet:**

  ```go
  runLine := -1.5
  selectionRunLine := "home"  // or "away"
  ```
