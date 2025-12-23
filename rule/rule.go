package rule

// ========================================
// 1. Specification パターン
// ========================================
// 条件をオブジェクトとして表現し、組み合わせ可能にする

// Specification は条件を表すインターフェース
type Specification interface {
	IsSatisfiedBy(user User, order Order) bool
}

// AndSpecification は複数の条件をANDで結合
type AndSpecification struct {
	specs []Specification
}

func NewAndSpecification(specs ...Specification) *AndSpecification {
	return &AndSpecification{specs: specs}
}

// TODO: IsSatisfiedByメソッドを実装してください
// ヒント: すべてのspecsがtrueを返す場合のみtrueを返す
func (a *AndSpecification) IsSatisfiedBy(user User, order Order) bool {
	// TODO: ここに実装
	return false
}

// OrSpecification は複数の条件をORで結合
type OrSpecification struct {
	specs []Specification
}

func NewOrSpecification(specs ...Specification) *OrSpecification {
	return &OrSpecification{specs: specs}
}

// TODO: IsSatisfiedByメソッドを実装してください
// ヒント: いずれかのspecsがtrueを返せばtrueを返す
func (o *OrSpecification) IsSatisfiedBy(user User, order Order) bool {
	// TODO: ここに実装
	return false
}

// NotSpecification は条件を反転
type NotSpecification struct {
	spec Specification
}

func NewNotSpecification(spec Specification) *NotSpecification {
	return &NotSpecification{spec: spec}
}

// TODO: IsSatisfiedByメソッドを実装してください
// ヒント: specの結果を反転させる
func (n *NotSpecification) IsSatisfiedBy(user User, order Order) bool {
	// TODO: ここに実装
	return false
}

// ========================================
// 具体的な条件の実装例
// ========================================

// MinimumAmountSpec は最低金額の条件
type MinimumAmountSpec struct {
	minAmount int
}

func NewMinimumAmountSpec(minAmount int) *MinimumAmountSpec {
	return &MinimumAmountSpec{minAmount: minAmount}
}

func (m *MinimumAmountSpec) IsSatisfiedBy(user User, order Order) bool {
	return order.Amount >= m.minAmount
}

// VerifiedUserSpec は認証済みユーザーの条件
type VerifiedUserSpec struct{}

func NewVerifiedUserSpec() *VerifiedUserSpec {
	return &VerifiedUserSpec{}
}

func (v *VerifiedUserSpec) IsSatisfiedBy(user User, order Order) bool {
	return user.IsVerified
}

// TODO: PremiumMemberSpec を実装してください
// プレミアム会員(MembershipLevel が "gold" または "platinum")かどうかを判定
type PremiumMemberSpec struct {
	// TODO: 必要なフィールドがあれば追加
}

func NewPremiumMemberSpec() *PremiumMemberSpec {
	// TODO: ここに実装
	return &PremiumMemberSpec{}
}

// TODO: IsSatisfiedByメソッドを実装してください
func (p *PremiumMemberSpec) IsSatisfiedBy(user User, order Order) bool {
	// TODO: ここに実装
	return false
}

// TODO: MinimumAgeSpec を実装してください
// 最低年齢の条件を判定
type MinimumAgeSpec struct {
	// TODO: フィールドを定義
}

// TODO: コンストラクタを実装
func NewMinimumAgeSpec(minAge int) *MinimumAgeSpec {
	// TODO: ここに実装
	return nil
}

// TODO: IsSatisfiedByメソッドを実装
func (m *MinimumAgeSpec) IsSatisfiedBy(user User, order Order) bool {
	// TODO: ここに実装
	return false
}
