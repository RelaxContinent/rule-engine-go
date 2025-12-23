package rule

// ========================================
// 4. Factory パターン
// ========================================
// ルールやハンドラーの生成を抽象化

// RuleFactory はルールを生成するファクトリー
type RuleFactory struct{}

func NewRuleFactory() *RuleFactory {
	return &RuleFactory{}
}

// CreateStandardApprovalChain は標準的な承認ルールチェーンを生成
func (f *RuleFactory) CreateStandardApprovalChain() RuleHandler {
	// 例: 認証済み かつ 最低金額1000円以上
	verifiedAndMinAmount := NewAndSpecification(
		NewVerifiedUserSpec(),
		NewMinimumAmountSpec(1000),
	)

	handler := NewApprovalRuleHandler(verifiedAndMinAmount)
	return handler
}

// TODO: CreatePremiumApprovalChain を実装してください
// プレミアム会員向けの承認ルールチェーンを生成
// 条件: プレミアム会員 または (認証済み かつ 最低金額5000円以上)
func (f *RuleFactory) CreatePremiumApprovalChain() RuleHandler {
	// TODO: ここに実装
	// ヒント: OrSpecificationとAndSpecificationを組み合わせる
	return nil
}

// CreateDiscountChain は割引ルールチェーンを生成
func (f *RuleFactory) CreateDiscountChain() RuleHandler {
	// プレミアム会員向け: 会員レベル割引
	premiumHandler := NewDiscountRuleHandler(
		NewPremiumMemberSpec(),
		NewMembershipDiscountStrategy(),
	)

	// TODO: 大量購入向けの割引ハンドラーを追加してください
	// 条件: 金額10000円以上
	// 戦略: VolumeDiscountStrategy
	// ヒント: NewDiscountRuleHandlerを使い、premiumHandler.SetNextでチェーンに追加

	// TODO: ここに実装

	return premiumHandler
}

// TODO: CreateRankChain を実装してください
// ランク判定ルールチェーンを生成
// 例:
// - Platinum: プレミアム会員 かつ 総購入額100万円以上
// - Gold: プレミアム会員 かつ 総購入額50万円以上
// - Silver: 認証済み かつ 総購入額10万円以上
// - Bronze: 認証済み
func (f *RuleFactory) CreateRankChain() RuleHandler {
	// TODO: ここに実装
	// ヒント: 複数のRankRuleHandlerをSetNextでチェーンにつなぐ
	// 高いランクから順に評価する
	return nil
}

// ========================================
// RuleEngine - すべてを統合
// ========================================

// RuleEngine はルールエンジンの本体
type RuleEngine struct {
	approvalChain RuleHandler
	discountChain RuleHandler
	rankChain     RuleHandler
}

// TODO: NewRuleEngine を実装してください
// ファクトリーを使って各種チェーンを初期化
func NewRuleEngine(factory *RuleFactory) *RuleEngine {
	// TODO: ここに実装
	// ヒント: factory.CreateXXXChain()を呼び出して各チェーンを設定
	return nil
}

// Evaluate はユーザーと注文を評価して結果を返す
func (e *RuleEngine) Evaluate(user User, order Order) *Result {
	result := &Result{}

	// TODO: 承認チェックを実装
	// ヒント: e.approvalChain.Handle(user, order)を呼び出す
	// approvalChainがnilでないことを確認

	// TODO: 割引計算を実装
	// ヒント: e.discountChain.Handle(user, order)を呼び出す

	// TODO: ランク判定を実装
	// ヒント: e.rankChain.Handle(user, order)を呼び出す

	// TODO: 各チェーンの結果をマージしてresultに設定

	return result
}

// TODO: EvaluateApproval を実装してください
// 承認判定のみを行う
func (e *RuleEngine) EvaluateApproval(user User, order Order) bool {
	// TODO: ここに実装
	return false
}

// TODO: CalculateDiscount を実装してください
// 割引率の計算のみを行う
func (e *RuleEngine) CalculateDiscount(user User, order Order) int {
	// TODO: ここに実装
	return 0
}

// TODO: DetermineRank を実装してください
// ランク判定のみを行う
func (e *RuleEngine) DetermineRank(user User, order Order) string {
	// TODO: ここに実装
	return "Unknown"
}
