# ルールエンジン - デザインパターン学習プロジェクト

Go言語でデザインパターンを実践的に学ぶためのプロジェクトです。

## 🎯 学習目標

このプロジェクトでは、以下の4つのデザインパターンを実装しながら学びます:

1. **Specification パターン** - 条件をオブジェクトとして表現
2. **Strategy パターン** - アルゴリズムの切り替え
3. **Chain of Responsibility パターン** - 責任の連鎖
4. **Factory パターン** - オブジェクト生成の抽象化

## 📁 プロジェクト構成

```
rule-engine-go/
├── main.go              # エントリーポイント
├── rule/
│   ├── order.go         # ドメインモデル(User, Order, Result)
│   ├── rule.go          # Specificationパターン
│   ├── engine.go        # Strategy & Chain of Responsibilityパターン
│   ├── factory.go       # Factoryパターン & RuleEngine統合
│   └── rule_test.go     # テストコード
└── README.md
```

## 🚀 学習の進め方

### ステップ1: Specificationパターン (rule.go)

条件をオブジェクトとして表現し、組み合わせ可能にします。

**実装するTODO:**
- [ ] `AndSpecification.IsSatisfiedBy()` - AND条件
- [ ] `OrSpecification.IsSatisfiedBy()` - OR条件
- [ ] `NotSpecification.IsSatisfiedBy()` - NOT条件
- [ ] `PremiumMemberSpec` - プレミアム会員判定
- [ ] `MinimumAgeSpec` - 最低年齢判定

**学習ポイント:**
- インターフェースを使った抽象化
- 条件の組み合わせ(Composite パターンの応用)
- if地獄を避ける設計

### ステップ2: Strategyパターン (engine.go)

割引計算のアルゴリズムを切り替え可能にします。

**実装するTODO:**
- [ ] `FixedDiscountStrategy` - 固定割引
- [ ] `MembershipDiscountStrategy` - 会員レベル別割引
- [ ] `VolumeDiscountStrategy` - 購入金額別割引

**学習ポイント:**
- アルゴリズムのカプセル化
- 実行時の戦略切り替え
- Open/Closed原則の実践

### ステップ3: Chain of Responsibilityパターン (engine.go)

ルールを順番に評価し、適切なハンドラーに処理を委譲します。

**実装するTODO:**
- [ ] `ApprovalRuleHandler.Handle()` - 承認ルール
- [ ] `DiscountRuleHandler.Handle()` - 割引ルール
- [ ] `RankRuleHandler` - ランク判定ルール

**学習ポイント:**
- 責任の連鎖
- 処理の委譲
- 柔軟なルール追加

### ステップ4: Factoryパターン (factory.go)

複雑なオブジェクト生成を抽象化します。

**実装するTODO:**
- [ ] `CreatePremiumApprovalChain()` - プレミアム承認チェーン
- [ ] `CreateDiscountChain()` - 割引チェーン
- [ ] `CreateRankChain()` - ランクチェーン
- [ ] `NewRuleEngine()` - エンジン初期化
- [ ] `RuleEngine.Evaluate()` - 統合評価

**学習ポイント:**
- 生成ロジックの分離
- 複雑な初期化の隠蔽
- 依存性の管理

### ステップ5: テスト実装 (rule_test.go)

TDD(テスト駆動開発)でパターンの理解を深めます。

**実装するTODO:**
- [ ] 各パターンの単体テスト
- [ ] 統合テスト
- [ ] エッジケースのテスト

## 🏃 実行方法

### 1. 依存関係の確認
```bash
go mod tidy
```

### 2. テストの実行
```bash
# すべてのテストを実行
go test ./rule/...

# 詳細表示
go test -v ./rule/...

# カバレッジ確認
go test -cover ./rule/...
```

### 3. プログラムの実行
```bash
go run main.go
```

## 📝 実装のヒント

### Specificationパターン
```go
// 複数の条件を組み合わせる例
spec := NewAndSpecification(
    NewVerifiedUserSpec(),
    NewMinimumAmountSpec(1000),
)

if spec.IsSatisfiedBy(user, order) {
    // 条件を満たす場合の処理
}
```

### Strategyパターン
```go
// 戦略を切り替える例
var strategy DiscountStrategy
if user.MembershipLevel == "gold" {
    strategy = NewMembershipDiscountStrategy()
} else {
    strategy = NewVolumeDiscountStrategy()
}
discount := strategy.CalculateDiscount(user, order)
```

### Chain of Responsibilityパターン
```go
// チェーンを構築する例
handler1 := NewApprovalRuleHandler(spec1)
handler2 := NewApprovalRuleHandler(spec2)
handler1.SetNext(handler2)

result := handler1.Handle(user, order)
```

## 🎓 学習後の発展課題

TODOをすべて実装したら、以下にチャレンジしてみましょう:

1. **新しいSpecificationを追加**
   - 特定カテゴリの商品判定
   - 時間帯による判定
   - 地域による判定

2. **新しいStrategyを追加**
   - 季節割引
   - クーポン割引
   - 複数割引の組み合わせ

3. **ロギング機能の追加**
   - Decorator パターンでログ追加
   - どのルールが適用されたか追跡

4. **設定ファイルからルール読み込み**
   - JSON/YAMLからルール定義を読み込む
   - Builder パターンで複雑なルールを構築

5. **パフォーマンス最適化**
   - キャッシュの追加
   - 並行処理の導入

## 📚 参考資料

- [Go言語でわかるデザインパターン](https://github.com/tmrts/go-patterns)
- [Specification Pattern](https://en.wikipedia.org/wiki/Specification_pattern)
- [Strategy Pattern](https://refactoring.guru/design-patterns/strategy)
- [Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility)

## ✅ チェックリスト

実装が完了したら、以下を確認しましょう:

- [ ] すべてのTODOコメントを実装した
- [ ] テストがすべてパスする
- [ ] `go run main.go`が正しく動作する
- [ ] コードレビュー: 各パターンの意図を理解している
- [ ] 発展課題に1つ以上チャレンジした

## 🤝 貢献

このプロジェクトは学習用です。改善案があれば自由に修正してください!

---

Happy Coding! 🚀
