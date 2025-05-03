# Table Tennis Bet Evaluation Tool

This Go program evaluates **table tennis betting outcomes** by comparing pre-match betting data with actual post-match results.

## 🔍 Features

- Parses **pre-match betting data** from `table_tennis_prematch.json`
- Parses **post-match results** from `table_tennis_result.json`
- Evaluates several common bet types:
  - ✅ 1X2 (Home/Draw/Away)
  - ✅ Over/Under
  - ✅ Correct Score
  - ✅ Double Chance
- Provides **clear win/loss indicators** for each evaluated bet

---

## ⚙️ Prerequisites

- [Go 1.24+](https://golang.org/dl/) installed
- Two JSON data files:
  - `table_tennis_prematch.json` – contains pre-match betting odds
  - `table_tennis_result.json` – contains post-match scores and outcomes

---

## 🛠 Installation

1. Clone or download this repository:

   ```bash
   git clone https://github.com/yourusername/table-tennis-bet-evaluator.git
   cd table-tennis-bet-evaluator
   ```

2. Ensure both JSON files are placed in the root directory:

   - `table_tennis_prematch.json`
   - `table_tennis_result.json`

---

## 🚀 Usage

Run the tool using:

```bash
go run *.go
```

You’ll see output like:

- Match details (teams, league)
- Final score
- Evaluation results:

  - ✔ Win or ✖ Loss for each bet type

---

## 🎯 Customizing Bets

You can modify the sample selections in the `main()` function to test specific bets:

- **1X2 Bet:**

  ```go
  selection1X2 := "1"  // or "X" or "2"
  ```

- **Over/Under Bet:**

  ```go
  selectionOUGoals := "Over"
  ouLine := 75.5
  ```

- **Correct Score Bet:**

  ```go
  correctScore := "11-9"
  ```

- **Double Chance Bet:**

  ```go
  doubleChance := "1X"  // options: "1X", "X2", "12"
  ```

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

## 🤝 Contributions

Feel free to open issues or pull requests to improve the tool or add new features like:

- Handicap betting
- Live odds comparison
- GUI interface
