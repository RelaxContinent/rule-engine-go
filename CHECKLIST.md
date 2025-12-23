# デザインパターン学習チェックリスト

## 📋 実装進捗

### ステップ1: Specificationパターン (rule/rule.go)

- [ ] `AndSpecification.IsSatisfiedBy()` 
  - すべての条件がtrueの場合のみtrueを返す
  - テスト: `TestAndSpecification`

- [ ] `OrSpecification.IsSatisfiedBy()`
  - いずれかの条件がtrueならtrueを返す
  - テスト: `TestOrSpecification`

- [ ] `NotSpecification.IsSatisfiedBy()`
  - 条件を反転させる
  - テスト: `TestNotSpecification`

- [ ] `PremiumMemberSpec`
  - MembershipLevelが"gold"または"platinum"ならtrue
  - テスト: `TestPremiumMemberSpec`

- [ ] `MinimumAgeSpec`
  - 最低年齢を満たすか判定
  - テスト: 自分で追加

---

### ステップ2: Strategyパターン (rule/engine.go)

- [ ] `FixedDiscountStrategy`
  - 固定の割引率を返す
  - テスト: `TestFixedDiscountStrategy`

- [ ] `MembershipDiscountStrategy.CalculateDiscount()`
  - bronze: 5%, silver: 10%, gold: 15%, platinum: 20%
  - テスト: `TestMembershipDiscountStrategy`

- [ ] `VolumeDiscountStrategy.CalculateDiscount()`
  - 10000円以上: 5%, 50000円以上: 10%, 100000円以上: 15%
  - テスト: `TestVolumeDiscountStrategy`

---

### ステップ3: Chain of Responsibilityパターン (rule/engine.go)

- [ ] `ApprovalRuleHandler.Handle()`
  - 条件を満たせば承認、満たさなければ次へ
  - テスト: `TestApprovalRuleHandler`

- [ ] `DiscountRuleHandler.Handle()`
  - 条件を満たせば割引計算、満たさなければ次へ
  - テスト: `TestDiscountRuleHandler`

- [ ] `RankRuleHandler` (全体)
  - 構造体、コンストラクタ、Handleメソッド
  - テスト: `TestRankRuleHandler`

---

### ステップ4: Factoryパターン (rule/factory.go)

- [ ] `CreatePremiumApprovalChain()`
  - プレミアム会員 OR (認証済み AND 5000円以上)
  - テスト: `TestRuleFactory`

- [ ] `CreateDiscountChain()`
  - 大量購入向けの割引ハンドラーを追加
  - プレミアム会員 → 大量購入のチェーン

- [ ] `CreateRankChain()`
  - Platinum, Gold, Silver, Bronzeのランク判定チェーン
  - 高いランクから順に評価

---

### ステップ5: RuleEngine統合 (rule/factory.go)

- [ ] `NewRuleEngine()`
  - ファクトリーを使って各チェーンを初期化

- [ ] `RuleEngine.Evaluate()`
  - 承認、割引、ランクをすべて評価
  - 結果をマージして返す

- [ ] `RuleEngine.EvaluateApproval()`
  - 承認判定のみ

- [ ] `RuleEngine.CalculateDiscount()`
  - 割引計算のみ

- [ ] `RuleEngine.DetermineRank()`
  - ランク判定のみ

---

### ステップ6: メイン実装 (main.go)

- [ ] RuleFactoryの作成

- [ ] RuleEngineの作成

- [ ] user1とorder1の評価

- [ ] user2とorder2の評価

- [ ] 個別評価メソッドのテスト

---

### ステップ7: テスト実装 (rule/rule_test.go)

- [ ] `TestAndSpecification`
- [ ] `TestOrSpecification`
- [ ] `TestNotSpecification`
- [ ] `TestPremiumMemberSpec`
- [ ] `TestFixedDiscountStrategy`
- [ ] `TestMembershipDiscountStrategy`
- [ ] `TestVolumeDiscountStrategy`
- [ ] `TestApprovalRuleHandler`
- [ ] `TestDiscountRuleHandler`
- [ ] `TestRankRuleHandler`
- [ ] `TestRuleFactory`
- [ ] `TestRuleEngine`

---

## 🎯 動作確認

### コンパイル確認
```bash
go build .
```

### テスト実行
```bash
# すべてのテスト
go test ./rule/...

# 詳細表示
go test -v ./rule/...

# カバレッジ
go test -cover ./rule/...
```

### プログラム実行
```bash
go run main.go
```

---

## 🚀 発展課題

基本実装が完了したら、以下にチャレンジ!

### レベル1: 新しいSpecificationを追加

- [ ] `CategorySpec` - 特定カテゴリの商品判定
- [ ] `TimeRangeSpec` - 時間帯による判定
- [ ] `TotalPurchasesSpec` - 総購入額による判定

### レベル2: 新しいStrategyを追加

- [ ] `SeasonalDiscountStrategy` - 季節割引
- [ ] `CouponDiscountStrategy` - クーポン割引
- [ ] `CombinedDiscountStrategy` - 複数割引の組み合わせ

### レベル3: 機能拡張

- [ ] ログ機能の追加(Decoratorパターン)
- [ ] JSON/YAMLからルール定義を読み込む
- [ ] Builderパターンで複雑なルールを構築
- [ ] キャッシュ機能の追加

### レベル4: パフォーマンス最適化

- [ ] ベンチマークテストの追加
- [ ] 並行処理の導入
- [ ] メモリ使用量の最適化

---

## 📝 学習メモ

### 学んだこと

- Specificationパターン:
  - 
  - 

- Strategyパターン:
  - 
  - 

- Chain of Responsibilityパターン:
  - 
  - 

- Factoryパターン:
  - 
  - 

### つまずいたポイント

- 
- 

### 次に学びたいパターン

- 
- 

---

## ✅ 完了確認

すべてのTODOを実装したら、以下を確認:

- [ ] すべてのテストがパス
- [ ] `go run main.go`が正しく動作
- [ ] コードレビュー: 各パターンの意図を理解している
- [ ] README.mdを読み直して理解を深めた
- [ ] GUIDE.mdの解答例と比較した
- [ ] 発展課題に1つ以上チャレンジした

---

**おめでとうございます!🎉**

すべて完了したら、次のデザインパターンにチャレンジしましょう!
