package rule

// ========================================
// 2. Strategy パターン
// ========================================
// 判定ロジックを切り替え可能にする

// DiscountStrategy は割引率を計算する戦略のインターフェース
type DiscountStrategy interface {
	CalculateDiscount(user User, order Order) int
}

// NoDiscountStrategy は割引なし
type NoDiscountStrategy struct{}

func (n *NoDiscountStrategy) CalculateDiscount(user User, order Order) int {
	return 0
}

// TODO: FixedDiscountStrategy を実装してください
// 固定割引率を返す戦略
type FixedDiscountStrategy struct {
	// TODO: フィールドを定義(固定の割引率を保持)
}

// TODO: コンストラクタを実装
func NewFixedDiscountStrategy(rate int) *FixedDiscountStrategy {
	// TODO: ここに実装
	return nil
}

// TODO: CalculateDiscountメソッドを実装
func (f *FixedDiscountStrategy) CalculateDiscount(user User, order Order) int {
	// TODO: ここに実装
	return 0
}

// TODO: MembershipDiscountStrategy を実装してください
// 会員レベルに応じた割引率を返す戦略
// bronze: 5%, silver: 10%, gold: 15%, platinum: 20%
type MembershipDiscountStrategy struct {
	// TODO: 必要なフィールドがあれば追加
}

func NewMembershipDiscountStrategy() *MembershipDiscountStrategy {
	return &MembershipDiscountStrategy{}
}

// TODO: CalculateDiscountメソッドを実装
func (m *MembershipDiscountStrategy) CalculateDiscount(user User, order Order) int {
	// TODO: ここに実装
	// ヒント: switch文やmapを使って会員レベルごとの割引率を返す
	return 0
}

// TODO: VolumeDiscountStrategy を実装してください
// 購入金額に応じた割引率を返す戦略
// 10000円以上: 5%, 50000円以上: 10%, 100000円以上: 15%
type VolumeDiscountStrategy struct{}

func NewVolumeDiscountStrategy() *VolumeDiscountStrategy {
	return &VolumeDiscountStrategy{}
}

// TODO: CalculateDiscountメソッドを実装
func (v *VolumeDiscountStrategy) CalculateDiscount(user User, order Order) int {
	// TODO: ここに実装
	return 0
}

// ========================================
// 3. Chain of Responsibility パターン
// ========================================
// ルールを順番に評価し、最初にマッチしたものを適用

// RuleHandler はルールチェーンのハンドラー
type RuleHandler interface {
	SetNext(handler RuleHandler) RuleHandler
	Handle(user User, order Order) *Result
}

// BaseRuleHandler は基本的なハンドラー実装
type BaseRuleHandler struct {
	next RuleHandler
}

func (b *BaseRuleHandler) SetNext(handler RuleHandler) RuleHandler {
	b.next = handler
	return handler
}

// ApprovalRuleHandler は承認ルールのハンドラー
type ApprovalRuleHandler struct {
	BaseRuleHandler
	spec Specification
}

func NewApprovalRuleHandler(spec Specification) *ApprovalRuleHandler {
	return &ApprovalRuleHandler{spec: spec}
}

// TODO: Handleメソッドを実装してください
// ヒント:
// 1. specの条件を満たす場合、Approved=trueのResultを返す
// 2. 満たさない場合、次のハンドラー(b.next)があればそれを呼び出す
// 3. 次のハンドラーがない場合、Approved=falseのResultを返す
func (a *ApprovalRuleHandler) Handle(user User, order Order) *Result {
	// TODO: ここに実装
	return &Result{
		Approved: false,
		Reason:   "No matching rule",
	}
}

// DiscountRuleHandler は割引ルールのハンドラー
type DiscountRuleHandler struct {
	BaseRuleHandler
	spec     Specification
	strategy DiscountStrategy
}

func NewDiscountRuleHandler(spec Specification, strategy DiscountStrategy) *DiscountRuleHandler {
	return &DiscountRuleHandler{
		spec:     spec,
		strategy: strategy,
	}
}

// TODO: Handleメソッドを実装してください
// ヒント:
// 1. specの条件を満たす場合、strategyで割引率を計算してResultを返す
// 2. 満たさない場合、次のハンドラーを呼び出す
func (d *DiscountRuleHandler) Handle(user User, order Order) *Result {
	// TODO: ここに実装
	return &Result{
		DiscountRate: 0,
		Reason:       "No matching discount rule",
	}
}

// TODO: RankRuleHandler を実装してください
// ユーザーのランクを判定するハンドラー
// 条件を満たす場合、Rankフィールドにランク名を設定
type RankRuleHandler struct {
	BaseRuleHandler
	// TODO: 必要なフィールドを定義
}

// TODO: コンストラクタを実装
func NewRankRuleHandler(spec Specification, rank string) *RankRuleHandler {
	// TODO: ここに実装
	return nil
}

// TODO: Handleメソッドを実装
func (r *RankRuleHandler) Handle(user User, order Order) *Result {
	// TODO: ここに実装
	return &Result{
		Rank:   "Unknown",
		Reason: "No matching rank rule",
	}
}
