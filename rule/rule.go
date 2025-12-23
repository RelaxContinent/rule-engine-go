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

func (a *AndSpecification) IsSatisfiedBy(user User, order Order) bool {
	for _, spec := range a.specs {
		if !spec.IsSatisfiedBy(user, order) {
			return false
		}
	}
	return true
}

// OrSpecification は複数の条件をORで結合
type OrSpecification struct {
	specs []Specification
}

func NewOrSpecification(specs ...Specification) *OrSpecification {
	return &OrSpecification{specs: specs}
}

func (o *OrSpecification) IsSatisfiedBy(user User, order Order) bool {
	for _, spec := range o.specs {
		if spec.IsSatisfiedBy(user, order) {
			return true
		}
	}
	return false
}

// NotSpecification は条件を反転
type NotSpecification struct {
	spec Specification
}

func NewNotSpecification(spec Specification) *NotSpecification {
	return &NotSpecification{spec: spec}
}

func (n *NotSpecification) IsSatisfiedBy(user User, order Order) bool {
	return !n.spec.IsSatisfiedBy(user, order)
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

type PremiumMemberSpec struct {
	MembershipLevels string
}

func NewPremiumMemberSpec() *PremiumMemberSpec {
	return &PremiumMemberSpec{}
}

func (p *PremiumMemberSpec) IsSatisfiedBy(user User, order Order) bool {
	return user.MembershipLevel == "gold" || user.MembershipLevel == "platinum"
}

type MinimumAgeSpec struct {
	minAge int
}

func NewMinimumAgeSpec(minAge int) *MinimumAgeSpec {
	return &MinimumAgeSpec{minAge: minAge}
}

func (m *MinimumAgeSpec) IsSatisfiedBy(user User, order Order) bool {
	return user.Age >= m.minAge
}
