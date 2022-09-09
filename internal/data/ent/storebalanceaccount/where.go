// Code generated by ent, DO NOT EDIT.

package storebalanceaccount

import (
	"finance/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AccountNo applies equality check predicate on the "account_no" field. It's identical to AccountNoEQ.
func AccountNo(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNo), v))
	})
}

// StoreCode applies equality check predicate on the "store_code" field. It's identical to StoreCodeEQ.
func StoreCode(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStoreCode), v))
	})
}

// UpperOrganNo applies equality check predicate on the "upper_organ_no" field. It's identical to UpperOrganNoEQ.
func UpperOrganNo(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpperOrganNo), v))
	})
}

// Pwd applies equality check predicate on the "pwd" field. It's identical to PwdEQ.
func Pwd(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPwd), v))
	})
}

// PwdSalt applies equality check predicate on the "pwd_salt" field. It's identical to PwdSaltEQ.
func PwdSalt(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPwdSalt), v))
	})
}

// BalanceFee applies equality check predicate on the "balance_fee" field. It's identical to BalanceFeeEQ.
func BalanceFee(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalanceFee), v))
	})
}

// TotalChargeFee applies equality check predicate on the "total_charge_fee" field. It's identical to TotalChargeFeeEQ.
func TotalChargeFee(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalChargeFee), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// AccountNoEQ applies the EQ predicate on the "account_no" field.
func AccountNoEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNo), v))
	})
}

// AccountNoNEQ applies the NEQ predicate on the "account_no" field.
func AccountNoNEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountNo), v))
	})
}

// AccountNoIn applies the In predicate on the "account_no" field.
func AccountNoIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAccountNo), v...))
	})
}

// AccountNoNotIn applies the NotIn predicate on the "account_no" field.
func AccountNoNotIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAccountNo), v...))
	})
}

// AccountNoGT applies the GT predicate on the "account_no" field.
func AccountNoGT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountNo), v))
	})
}

// AccountNoGTE applies the GTE predicate on the "account_no" field.
func AccountNoGTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountNo), v))
	})
}

// AccountNoLT applies the LT predicate on the "account_no" field.
func AccountNoLT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountNo), v))
	})
}

// AccountNoLTE applies the LTE predicate on the "account_no" field.
func AccountNoLTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountNo), v))
	})
}

// AccountNoContains applies the Contains predicate on the "account_no" field.
func AccountNoContains(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountNo), v))
	})
}

// AccountNoHasPrefix applies the HasPrefix predicate on the "account_no" field.
func AccountNoHasPrefix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountNo), v))
	})
}

// AccountNoHasSuffix applies the HasSuffix predicate on the "account_no" field.
func AccountNoHasSuffix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountNo), v))
	})
}

// AccountNoIsNil applies the IsNil predicate on the "account_no" field.
func AccountNoIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAccountNo)))
	})
}

// AccountNoNotNil applies the NotNil predicate on the "account_no" field.
func AccountNoNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAccountNo)))
	})
}

// AccountNoEqualFold applies the EqualFold predicate on the "account_no" field.
func AccountNoEqualFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountNo), v))
	})
}

// AccountNoContainsFold applies the ContainsFold predicate on the "account_no" field.
func AccountNoContainsFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountNo), v))
	})
}

// StoreCodeEQ applies the EQ predicate on the "store_code" field.
func StoreCodeEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStoreCode), v))
	})
}

// StoreCodeNEQ applies the NEQ predicate on the "store_code" field.
func StoreCodeNEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStoreCode), v))
	})
}

// StoreCodeIn applies the In predicate on the "store_code" field.
func StoreCodeIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStoreCode), v...))
	})
}

// StoreCodeNotIn applies the NotIn predicate on the "store_code" field.
func StoreCodeNotIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStoreCode), v...))
	})
}

// StoreCodeGT applies the GT predicate on the "store_code" field.
func StoreCodeGT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStoreCode), v))
	})
}

// StoreCodeGTE applies the GTE predicate on the "store_code" field.
func StoreCodeGTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStoreCode), v))
	})
}

// StoreCodeLT applies the LT predicate on the "store_code" field.
func StoreCodeLT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStoreCode), v))
	})
}

// StoreCodeLTE applies the LTE predicate on the "store_code" field.
func StoreCodeLTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStoreCode), v))
	})
}

// StoreCodeContains applies the Contains predicate on the "store_code" field.
func StoreCodeContains(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStoreCode), v))
	})
}

// StoreCodeHasPrefix applies the HasPrefix predicate on the "store_code" field.
func StoreCodeHasPrefix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStoreCode), v))
	})
}

// StoreCodeHasSuffix applies the HasSuffix predicate on the "store_code" field.
func StoreCodeHasSuffix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStoreCode), v))
	})
}

// StoreCodeIsNil applies the IsNil predicate on the "store_code" field.
func StoreCodeIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStoreCode)))
	})
}

// StoreCodeNotNil applies the NotNil predicate on the "store_code" field.
func StoreCodeNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStoreCode)))
	})
}

// StoreCodeEqualFold applies the EqualFold predicate on the "store_code" field.
func StoreCodeEqualFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStoreCode), v))
	})
}

// StoreCodeContainsFold applies the ContainsFold predicate on the "store_code" field.
func StoreCodeContainsFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStoreCode), v))
	})
}

// UpperOrganNoEQ applies the EQ predicate on the "upper_organ_no" field.
func UpperOrganNoEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoNEQ applies the NEQ predicate on the "upper_organ_no" field.
func UpperOrganNoNEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoIn applies the In predicate on the "upper_organ_no" field.
func UpperOrganNoIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpperOrganNo), v...))
	})
}

// UpperOrganNoNotIn applies the NotIn predicate on the "upper_organ_no" field.
func UpperOrganNoNotIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpperOrganNo), v...))
	})
}

// UpperOrganNoGT applies the GT predicate on the "upper_organ_no" field.
func UpperOrganNoGT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoGTE applies the GTE predicate on the "upper_organ_no" field.
func UpperOrganNoGTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoLT applies the LT predicate on the "upper_organ_no" field.
func UpperOrganNoLT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoLTE applies the LTE predicate on the "upper_organ_no" field.
func UpperOrganNoLTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoContains applies the Contains predicate on the "upper_organ_no" field.
func UpperOrganNoContains(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoHasPrefix applies the HasPrefix predicate on the "upper_organ_no" field.
func UpperOrganNoHasPrefix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoHasSuffix applies the HasSuffix predicate on the "upper_organ_no" field.
func UpperOrganNoHasSuffix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoIsNil applies the IsNil predicate on the "upper_organ_no" field.
func UpperOrganNoIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpperOrganNo)))
	})
}

// UpperOrganNoNotNil applies the NotNil predicate on the "upper_organ_no" field.
func UpperOrganNoNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpperOrganNo)))
	})
}

// UpperOrganNoEqualFold applies the EqualFold predicate on the "upper_organ_no" field.
func UpperOrganNoEqualFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUpperOrganNo), v))
	})
}

// UpperOrganNoContainsFold applies the ContainsFold predicate on the "upper_organ_no" field.
func UpperOrganNoContainsFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUpperOrganNo), v))
	})
}

// PwdEQ applies the EQ predicate on the "pwd" field.
func PwdEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPwd), v))
	})
}

// PwdNEQ applies the NEQ predicate on the "pwd" field.
func PwdNEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPwd), v))
	})
}

// PwdIn applies the In predicate on the "pwd" field.
func PwdIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPwd), v...))
	})
}

// PwdNotIn applies the NotIn predicate on the "pwd" field.
func PwdNotIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPwd), v...))
	})
}

// PwdGT applies the GT predicate on the "pwd" field.
func PwdGT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPwd), v))
	})
}

// PwdGTE applies the GTE predicate on the "pwd" field.
func PwdGTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPwd), v))
	})
}

// PwdLT applies the LT predicate on the "pwd" field.
func PwdLT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPwd), v))
	})
}

// PwdLTE applies the LTE predicate on the "pwd" field.
func PwdLTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPwd), v))
	})
}

// PwdContains applies the Contains predicate on the "pwd" field.
func PwdContains(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPwd), v))
	})
}

// PwdHasPrefix applies the HasPrefix predicate on the "pwd" field.
func PwdHasPrefix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPwd), v))
	})
}

// PwdHasSuffix applies the HasSuffix predicate on the "pwd" field.
func PwdHasSuffix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPwd), v))
	})
}

// PwdIsNil applies the IsNil predicate on the "pwd" field.
func PwdIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPwd)))
	})
}

// PwdNotNil applies the NotNil predicate on the "pwd" field.
func PwdNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPwd)))
	})
}

// PwdEqualFold applies the EqualFold predicate on the "pwd" field.
func PwdEqualFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPwd), v))
	})
}

// PwdContainsFold applies the ContainsFold predicate on the "pwd" field.
func PwdContainsFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPwd), v))
	})
}

// PwdSaltEQ applies the EQ predicate on the "pwd_salt" field.
func PwdSaltEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltNEQ applies the NEQ predicate on the "pwd_salt" field.
func PwdSaltNEQ(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltIn applies the In predicate on the "pwd_salt" field.
func PwdSaltIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPwdSalt), v...))
	})
}

// PwdSaltNotIn applies the NotIn predicate on the "pwd_salt" field.
func PwdSaltNotIn(vs ...string) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPwdSalt), v...))
	})
}

// PwdSaltGT applies the GT predicate on the "pwd_salt" field.
func PwdSaltGT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltGTE applies the GTE predicate on the "pwd_salt" field.
func PwdSaltGTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltLT applies the LT predicate on the "pwd_salt" field.
func PwdSaltLT(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltLTE applies the LTE predicate on the "pwd_salt" field.
func PwdSaltLTE(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltContains applies the Contains predicate on the "pwd_salt" field.
func PwdSaltContains(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltHasPrefix applies the HasPrefix predicate on the "pwd_salt" field.
func PwdSaltHasPrefix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltHasSuffix applies the HasSuffix predicate on the "pwd_salt" field.
func PwdSaltHasSuffix(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltIsNil applies the IsNil predicate on the "pwd_salt" field.
func PwdSaltIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPwdSalt)))
	})
}

// PwdSaltNotNil applies the NotNil predicate on the "pwd_salt" field.
func PwdSaltNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPwdSalt)))
	})
}

// PwdSaltEqualFold applies the EqualFold predicate on the "pwd_salt" field.
func PwdSaltEqualFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPwdSalt), v))
	})
}

// PwdSaltContainsFold applies the ContainsFold predicate on the "pwd_salt" field.
func PwdSaltContainsFold(v string) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPwdSalt), v))
	})
}

// BalanceFeeEQ applies the EQ predicate on the "balance_fee" field.
func BalanceFeeEQ(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBalanceFee), v))
	})
}

// BalanceFeeNEQ applies the NEQ predicate on the "balance_fee" field.
func BalanceFeeNEQ(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBalanceFee), v))
	})
}

// BalanceFeeIn applies the In predicate on the "balance_fee" field.
func BalanceFeeIn(vs ...float64) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBalanceFee), v...))
	})
}

// BalanceFeeNotIn applies the NotIn predicate on the "balance_fee" field.
func BalanceFeeNotIn(vs ...float64) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBalanceFee), v...))
	})
}

// BalanceFeeGT applies the GT predicate on the "balance_fee" field.
func BalanceFeeGT(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBalanceFee), v))
	})
}

// BalanceFeeGTE applies the GTE predicate on the "balance_fee" field.
func BalanceFeeGTE(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBalanceFee), v))
	})
}

// BalanceFeeLT applies the LT predicate on the "balance_fee" field.
func BalanceFeeLT(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBalanceFee), v))
	})
}

// BalanceFeeLTE applies the LTE predicate on the "balance_fee" field.
func BalanceFeeLTE(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBalanceFee), v))
	})
}

// BalanceFeeIsNil applies the IsNil predicate on the "balance_fee" field.
func BalanceFeeIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBalanceFee)))
	})
}

// BalanceFeeNotNil applies the NotNil predicate on the "balance_fee" field.
func BalanceFeeNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBalanceFee)))
	})
}

// TotalChargeFeeEQ applies the EQ predicate on the "total_charge_fee" field.
func TotalChargeFeeEQ(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalChargeFee), v))
	})
}

// TotalChargeFeeNEQ applies the NEQ predicate on the "total_charge_fee" field.
func TotalChargeFeeNEQ(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalChargeFee), v))
	})
}

// TotalChargeFeeIn applies the In predicate on the "total_charge_fee" field.
func TotalChargeFeeIn(vs ...float64) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalChargeFee), v...))
	})
}

// TotalChargeFeeNotIn applies the NotIn predicate on the "total_charge_fee" field.
func TotalChargeFeeNotIn(vs ...float64) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalChargeFee), v...))
	})
}

// TotalChargeFeeGT applies the GT predicate on the "total_charge_fee" field.
func TotalChargeFeeGT(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalChargeFee), v))
	})
}

// TotalChargeFeeGTE applies the GTE predicate on the "total_charge_fee" field.
func TotalChargeFeeGTE(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalChargeFee), v))
	})
}

// TotalChargeFeeLT applies the LT predicate on the "total_charge_fee" field.
func TotalChargeFeeLT(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalChargeFee), v))
	})
}

// TotalChargeFeeLTE applies the LTE predicate on the "total_charge_fee" field.
func TotalChargeFeeLTE(v float64) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalChargeFee), v))
	})
}

// TotalChargeFeeIsNil applies the IsNil predicate on the "total_charge_fee" field.
func TotalChargeFeeIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTotalChargeFee)))
	})
}

// TotalChargeFeeNotNil applies the NotNil predicate on the "total_charge_fee" field.
func TotalChargeFeeNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTotalChargeFee)))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedIn applies the In predicate on the "is_deleted" field.
func IsDeletedIn(vs ...int8) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedNotIn applies the NotIn predicate on the "is_deleted" field.
func IsDeletedNotIn(vs ...int8) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIsDeleted), v...))
	})
}

// IsDeletedGT applies the GT predicate on the "is_deleted" field.
func IsDeletedGT(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedGTE applies the GTE predicate on the "is_deleted" field.
func IsDeletedGTE(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLT applies the LT predicate on the "is_deleted" field.
func IsDeletedLT(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedLTE applies the LTE predicate on the "is_deleted" field.
func IsDeletedLTE(v int8) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedIsNil applies the IsNil predicate on the "is_deleted" field.
func IsDeletedIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsDeleted)))
	})
}

// IsDeletedNotNil applies the NotNil predicate on the "is_deleted" field.
func IsDeletedNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsDeleted)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StoreBalanceAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StoreBalanceAccount) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StoreBalanceAccount) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StoreBalanceAccount) predicate.StoreBalanceAccount {
	return predicate.StoreBalanceAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}
