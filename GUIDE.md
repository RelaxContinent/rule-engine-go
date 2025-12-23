# 実装ガイド

このドキュメントでは、各TODOの実装方法を段階的に解説します。

## 📖 目次

1. [Specificationパターン](#1-specificationパターン)
2. [Strategyパターン](#2-strategyパターン)
3. [Chain of Responsibilityパターン](#3-chain-of-responsibilityパターン)
4. [Factoryパターン](#4-factoryパターン)
5. [RuleEngine統合](#5-ruleengine統合)
6. [テスト実装](#6-テスト実装)

---

## 1. Specificationパターン

### 1.1 AndSpecification.IsSatisfiedBy()

**目的:** 複数の条件がすべて満たされているか判定

```go
func (a *AndSpecification) IsSatisfiedBy(user User, order Order) bool {
    for _, spec := range a.specs {
        if !spec.IsSatisfiedBy(user, order) {
            return false
        }
    }
    return true
}
```

**ポイント:**
- すべてのspecsをループで確認
- 1つでもfalseがあればfalseを返す
- すべてtrueならtrueを返す

### 1.2 OrSpecification.IsSatisfiedBy()

**目的:** 複数の条件のいずれかが満たされているか判定

```go
func (o *OrSpecification) IsSatisfiedBy(user User, order Order) bool {
    for _, spec := range o.specs {
        if spec.IsSatisfiedBy(user, order) {
            return true
        }
    }
    return false
}
```

**ポイント:**
- 1つでもtrueがあればtrueを返す
- すべてfalseならfalseを返す

### 1.3 NotSpecification.IsSatisfiedBy()

**目的:** 条件を反転

```go
func (n *NotSpecification) IsSatisfiedBy(user User, order Order) bool {
    return !n.spec.IsSatisfiedBy(user, order)
}
```

### 1.4 PremiumMemberSpec

**目的:** ゴールドまたはプラチナ会員かどうか判定

```go
type PremiumMemberSpec struct{}

func NewPremiumMemberSpec() *PremiumMemberSpec {
    return &PremiumMemberSpec{}
}

func (p *PremiumMemberSpec) IsSatisfiedBy(user User, order Order) bool {
    return user.MembershipLevel == "gold" || user.MembershipLevel == "platinum"
}
```

### 1.5 MinimumAgeSpec

**目的:** 最低年齢を満たしているか判定

```go
type MinimumAgeSpec struct {
    minAge int
}

func NewMinimumAgeSpec(minAge int) *MinimumAgeSpec {
    return &MinimumAgeSpec{minAge: minAge}
}

func (m *MinimumAgeSpec) IsSatisfiedBy(user User, order Order) bool {
    return user.Age >= m.minAge
}
```

---

## 2. Strategyパターン

### 2.1 FixedDiscountStrategy

**目的:** 固定の割引率を返す

```go
type FixedDiscountStrategy struct {
    rate int
}

func NewFixedDiscountStrategy(rate int) *FixedDiscountStrategy {
    return &FixedDiscountStrategy{rate: rate}
}

func (f *FixedDiscountStrategy) CalculateDiscount(user User, order Order) int {
    return f.rate
}
```

### 2.2 MembershipDiscountStrategy

**目的:** 会員レベルに応じた割引率を返す

```go
func (m *MembershipDiscountStrategy) CalculateDiscount(user User, order Order) int {
    switch user.MembershipLevel {
    case "bronze":
        return 5
    case "silver":
        return 10
    case "gold":
        return 15
    case "platinum":
        return 20
    default:
        return 0
    }
}
```

**別解: mapを使う方法**
```go
type MembershipDiscountStrategy struct {
    rates map[string]int
}

func NewMembershipDiscountStrategy() *MembershipDiscountStrategy {
    return &MembershipDiscountStrategy{
        rates: map[string]int{
            "bronze":   5,
            "silver":   10,
            "gold":     15,
            "platinum": 20,
        },
    }
}

func (m *MembershipDiscountStrategy) CalculateDiscount(user User, order Order) int {
    if rate, ok := m.rates[user.MembershipLevel]; ok {
        return rate
    }
    return 0
}
```

### 2.3 VolumeDiscountStrategy

**目的:** 購入金額に応じた割引率を返す

```go
func (v *VolumeDiscountStrategy) CalculateDiscount(user User, order Order) int {
    if order.Amount >= 100000 {
        return 15
    } else if order.Amount >= 50000 {
        return 10
    } else if order.Amount >= 10000 {
        return 5
    }
    return 0
}
```

---

## 3. Chain of Responsibilityパターン

### 3.1 ApprovalRuleHandler.Handle()

**目的:** 承認条件を満たすか判定し、満たさなければ次へ

```go
func (a *ApprovalRuleHandler) Handle(user User, order Order) *Result {
    if a.spec.IsSatisfiedBy(user, order) {
        return &Result{
            Approved: true,
            Reason:   "Approved by rule",
        }
    }
    
    if a.next != nil {
        return a.next.Handle(user, order)
    }
    
    return &Result{
        Approved: false,
        Reason:   "No matching approval rule",
    }
}
```

**ポイント:**
1. 条件を満たせば承認結果を返す
2. 満たさず次のハンドラーがあれば委譲
3. 次もなければ却下

### 3.2 DiscountRuleHandler.Handle()

**目的:** 割引条件を満たすか判定し、割引率を計算

```go
func (d *DiscountRuleHandler) Handle(user User, order Order) *Result {
    if d.spec.IsSatisfiedBy(user, order) {
        discount := d.strategy.CalculateDiscount(user, order)
        return &Result{
            DiscountRate: discount,
            Reason:       "Discount applied",
        }
    }
    
    if d.next != nil {
        return d.next.Handle(user, order)
    }
    
    return &Result{
        DiscountRate: 0,
        Reason:       "No matching discount rule",
    }
}
```

### 3.3 RankRuleHandler

**目的:** ランク条件を満たすか判定

```go
type RankRuleHandler struct {
    BaseRuleHandler
    spec Specification
    rank string
}

func NewRankRuleHandler(spec Specification, rank string) *RankRuleHandler {
    return &RankRuleHandler{
        spec: spec,
        rank: rank,
    }
}

func (r *RankRuleHandler) Handle(user User, order Order) *Result {
    if r.spec.IsSatisfiedBy(user, order) {
        return &Result{
            Rank:   r.rank,
            Reason: "Rank determined",
        }
    }
    
    if r.next != nil {
        return r.next.Handle(user, order)
    }
    
    return &Result{
        Rank:   "Unknown",
        Reason: "No matching rank rule",
    }
}
```

---

## 4. Factoryパターン

### 4.1 CreatePremiumApprovalChain()

**目的:** プレミアム会員向けの承認ルールを生成

```go
func (f *RuleFactory) CreatePremiumApprovalChain() RuleHandler {
    // プレミアム会員 OR (認証済み AND 最低金額5000円以上)
    premiumOrVerifiedWithAmount := NewOrSpecification(
        NewPremiumMemberSpec(),
        NewAndSpecification(
            NewVerifiedUserSpec(),
            NewMinimumAmountSpec(5000),
        ),
    )
    
    return NewApprovalRuleHandler(premiumOrVerifiedWithAmount)
}
```

### 4.2 CreateDiscountChain()

**目的:** 割引ルールチェーンを生成

```go
func (f *RuleFactory) CreateDiscountChain() RuleHandler {
    // プレミアム会員向け: 会員レベル割引
    premiumHandler := NewDiscountRuleHandler(
        NewPremiumMemberSpec(),
        NewMembershipDiscountStrategy(),
    )
    
    // 大量購入向け: ボリューム割引
    volumeHandler := NewDiscountRuleHandler(
        NewMinimumAmountSpec(10000),
        NewVolumeDiscountStrategy(),
    )
    
    // チェーンに接続
    premiumHandler.SetNext(volumeHandler)
    
    return premiumHandler
}
```

### 4.3 CreateRankChain()

**目的:** ランク判定ルールチェーンを生成

```go
func (f *RuleFactory) CreateRankChain() RuleHandler {
    // Platinum: プレミアム会員 AND 総購入額100万円以上
    platinumSpec := NewAndSpecification(
        NewPremiumMemberSpec(),
        // 総購入額の判定用に新しいSpecを作成する必要があります
        // 簡易版として、ここでは直接実装
    )
    
    // より実用的な実装例:
    // 各ランクのハンドラーを作成(高いランクから順に)
    platinumHandler := NewRankRuleHandler(
        NewAndSpecification(
            NewPremiumMemberSpec(),
            // TotalPurchasesSpecを実装する必要があります
        ),
        "Platinum",
    )
    
    goldHandler := NewRankRuleHandler(
        NewPremiumMemberSpec(),
        "Gold",
    )
    
    silverHandler := NewRankRuleHandler(
        NewVerifiedUserSpec(),
        "Silver",
    )
    
    bronzeHandler := NewRankRuleHandler(
        NewVerifiedUserSpec(),
        "Bronze",
    )
    
    // チェーンに接続
    platinumHandler.SetNext(goldHandler)
    goldHandler.SetNext(silverHandler)
    silverHandler.SetNext(bronzeHandler)
    
    return platinumHandler
}
```

**注意:** TotalPurchasesSpecなど、追加のSpecificationが必要になります。

---

## 5. RuleEngine統合

### 5.1 NewRuleEngine()

**目的:** ファクトリーを使ってエンジンを初期化

```go
func NewRuleEngine(factory *RuleFactory) *RuleEngine {
    return &RuleEngine{
        approvalChain: factory.CreateStandardApprovalChain(),
        discountChain: factory.CreateDiscountChain(),
        rankChain:     factory.CreateRankChain(),
    }
}
```

### 5.2 RuleEngine.Evaluate()

**目的:** すべてのルールを評価して統合結果を返す

```go
func (e *RuleEngine) Evaluate(user User, order Order) *Result {
    result := &Result{}
    
    // 承認チェック
    if e.approvalChain != nil {
        approvalResult := e.approvalChain.Handle(user, order)
        result.Approved = approvalResult.Approved
    }
    
    // 割引計算
    if e.discountChain != nil {
        discountResult := e.discountChain.Handle(user, order)
        result.DiscountRate = discountResult.DiscountRate
    }
    
    // ランク判定
    if e.rankChain != nil {
        rankResult := e.rankChain.Handle(user, order)
        result.Rank = rankResult.Rank
    }
    
    result.Reason = "Evaluation completed"
    return result
}
```

### 5.3 個別評価メソッド

```go
func (e *RuleEngine) EvaluateApproval(user User, order Order) bool {
    if e.approvalChain == nil {
        return false
    }
    result := e.approvalChain.Handle(user, order)
    return result.Approved
}

func (e *RuleEngine) CalculateDiscount(user User, order Order) int {
    if e.discountChain == nil {
        return 0
    }
    result := e.discountChain.Handle(user, order)
    return result.DiscountRate
}

func (e *RuleEngine) DetermineRank(user User, order Order) string {
    if e.rankChain == nil {
        return "Unknown"
    }
    result := e.rankChain.Handle(user, order)
    return result.Rank
}
```

---

## 6. テスト実装

### 6.1 基本的なテストパターン

```go
func TestAndSpecification(t *testing.T) {
    spec1 := NewVerifiedUserSpec()
    spec2 := NewMinimumAmountSpec(1000)
    andSpec := NewAndSpecification(spec1, spec2)
    
    // 両方満たす場合
    user := User{IsVerified: true}
    order := Order{Amount: 1500}
    if !andSpec.IsSatisfiedBy(user, order) {
        t.Error("Expected true when both conditions are met")
    }
    
    // 片方満たさない場合
    order.Amount = 500
    if andSpec.IsSatisfiedBy(user, order) {
        t.Error("Expected false when one condition is not met")
    }
}
```

### 6.2 テーブル駆動テスト

```go
func TestMembershipDiscountStrategy(t *testing.T) {
    strategy := NewMembershipDiscountStrategy()
    
    tests := []struct {
        name     string
        level    string
        expected int
    }{
        {"Bronze member", "bronze", 5},
        {"Silver member", "silver", 10},
        {"Gold member", "gold", 15},
        {"Platinum member", "platinum", 20},
        {"Unknown level", "unknown", 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            user := User{MembershipLevel: tt.level}
            order := Order{}
            
            got := strategy.CalculateDiscount(user, order)
            if got != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, got)
            }
        })
    }
}
```

---

## 🎯 次のステップ

1. 各セクションのコードを実装
2. `go test -v ./rule/...` でテスト実行
3. `go run main.go` で動作確認
4. 発展課題にチャレンジ!

Happy Coding! 🚀
