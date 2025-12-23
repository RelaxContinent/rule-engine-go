package main

import (
	"fmt"
	"rule-engine-go/rule"
	"time"
)

func main() {
	fmt.Println("🚀 ルールエンジン - デザインパターン学習プロジェクト")
	fmt.Println("=" + "===========================================")

	// TODO: RuleFactoryを作成
	// ヒント: rule.NewRuleFactory()

	// TODO: RuleEngineを作成
	// ヒント: rule.NewRuleEngine(factory)

	// テストデータの準備
	user1 := rule.User{
		ID:              "user001",
		Name:            "山田太郎",
		Age:             25,
		MembershipLevel: "gold",
		TotalPurchases:  500000,
		IsVerified:      true,
	}

	order1 := rule.Order{
		ID:        "order001",
		UserID:    "user001",
		Amount:    15000,
		ItemCount: 3,
		CreatedAt: time.Now(),
		IsPremium: true,
		Category:  "electronics",
	}

	user2 := rule.User{
		ID:              "user002",
		Name:            "佐藤花子",
		Age:             30,
		MembershipLevel: "bronze",
		TotalPurchases:  50000,
		IsVerified:      true,
	}

	order2 := rule.Order{
		ID:        "order002",
		UserID:    "user002",
		Amount:    3000,
		ItemCount: 1,
		CreatedAt: time.Now(),
		IsPremium: false,
		Category:  "books",
	}

	// TODO実装までの一時的な使用(実装時に削除してください)
	_, _, _, _ = user1, order1, user2, order2

	// TODO: user1とorder1を評価
	// ヒント: engine.Evaluate(user1, order1)
	// 結果を表示

	fmt.Println("\n--- ユーザー1の評価結果 ---")
	// TODO: 結果を表示
	// fmt.Printf("承認: %v\n", result.Approved)
	// fmt.Printf("割引率: %d%%\n", result.DiscountRate)
	// fmt.Printf("ランク: %s\n", result.Rank)
	// fmt.Printf("理由: %s\n", result.Reason)

	// TODO: user2とorder2を評価して結果を表示
	fmt.Println("\n--- ユーザー2の評価結果 ---")
	// TODO: ここに実装

	// TODO: 個別の評価メソッドもテスト
	// EvaluateApproval, CalculateDiscount, DetermineRank
	fmt.Println("\n--- 個別評価のテスト ---")
	// TODO: ここに実装

	fmt.Println("\n✅ 実装が完了したら、このプログラムが正しく動作するはずです!")
}
