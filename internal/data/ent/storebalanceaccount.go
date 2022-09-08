// Code generated by ent, DO NOT EDIT.

package ent

import (
	"finance/internal/data/ent/storebalanceaccount"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// StoreBalanceAccount is the model entity for the StoreBalanceAccount schema.
type StoreBalanceAccount struct {
	config `json:"-"`
	// ID of the ent.
	// 主键自增
	ID uint32 `json:"id,omitempty"`
	// 门店编号
	StoreCode string `json:"store_code,omitempty"`
	// 上级机构号 加盟商franchisee_id
	UpperOrganNo string `json:"upper_organ_no,omitempty"`
	// 密码
	Pwd string `json:"pwd,omitempty"`
	// 密码加盐
	PwdSalt string `json:"pwd_salt,omitempty"`
	// 余额
	BalanceFee float64 `json:"balance_fee,omitempty"`
	// 总充值金额
	TotalChargeFee float64 `json:"total_charge_fee,omitempty"`
	// 0未删除 1已删除
	IsDeleted int8 `json:"is_deleted,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StoreBalanceAccount) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case storebalanceaccount.FieldBalanceFee, storebalanceaccount.FieldTotalChargeFee:
			values[i] = new(sql.NullFloat64)
		case storebalanceaccount.FieldID, storebalanceaccount.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case storebalanceaccount.FieldStoreCode, storebalanceaccount.FieldUpperOrganNo, storebalanceaccount.FieldPwd, storebalanceaccount.FieldPwdSalt:
			values[i] = new(sql.NullString)
		case storebalanceaccount.FieldUpdatedAt, storebalanceaccount.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type StoreBalanceAccount", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StoreBalanceAccount fields.
func (sba *StoreBalanceAccount) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case storebalanceaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sba.ID = uint32(value.Int64)
		case storebalanceaccount.FieldStoreCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_code", values[i])
			} else if value.Valid {
				sba.StoreCode = value.String
			}
		case storebalanceaccount.FieldUpperOrganNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field upper_organ_no", values[i])
			} else if value.Valid {
				sba.UpperOrganNo = value.String
			}
		case storebalanceaccount.FieldPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pwd", values[i])
			} else if value.Valid {
				sba.Pwd = value.String
			}
		case storebalanceaccount.FieldPwdSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pwd_salt", values[i])
			} else if value.Valid {
				sba.PwdSalt = value.String
			}
		case storebalanceaccount.FieldBalanceFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance_fee", values[i])
			} else if value.Valid {
				sba.BalanceFee = value.Float64
			}
		case storebalanceaccount.FieldTotalChargeFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_charge_fee", values[i])
			} else if value.Valid {
				sba.TotalChargeFee = value.Float64
			}
		case storebalanceaccount.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				sba.IsDeleted = int8(value.Int64)
			}
		case storebalanceaccount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sba.UpdatedAt = value.Time
			}
		case storebalanceaccount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sba.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this StoreBalanceAccount.
// Note that you need to call StoreBalanceAccount.Unwrap() before calling this method if this StoreBalanceAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (sba *StoreBalanceAccount) Update() *StoreBalanceAccountUpdateOne {
	return (&StoreBalanceAccountClient{config: sba.config}).UpdateOne(sba)
}

// Unwrap unwraps the StoreBalanceAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sba *StoreBalanceAccount) Unwrap() *StoreBalanceAccount {
	_tx, ok := sba.config.driver.(*txDriver)
	if !ok {
		panic("ent: StoreBalanceAccount is not a transactional entity")
	}
	sba.config.driver = _tx.drv
	return sba
}

// String implements the fmt.Stringer.
func (sba *StoreBalanceAccount) String() string {
	var builder strings.Builder
	builder.WriteString("StoreBalanceAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sba.ID))
	builder.WriteString("store_code=")
	builder.WriteString(sba.StoreCode)
	builder.WriteString(", ")
	builder.WriteString("upper_organ_no=")
	builder.WriteString(sba.UpperOrganNo)
	builder.WriteString(", ")
	builder.WriteString("pwd=")
	builder.WriteString(sba.Pwd)
	builder.WriteString(", ")
	builder.WriteString("pwd_salt=")
	builder.WriteString(sba.PwdSalt)
	builder.WriteString(", ")
	builder.WriteString("balance_fee=")
	builder.WriteString(fmt.Sprintf("%v", sba.BalanceFee))
	builder.WriteString(", ")
	builder.WriteString("total_charge_fee=")
	builder.WriteString(fmt.Sprintf("%v", sba.TotalChargeFee))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", sba.IsDeleted))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sba.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sba.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// StoreBalanceAccounts is a parsable slice of StoreBalanceAccount.
type StoreBalanceAccounts []*StoreBalanceAccount

func (sba StoreBalanceAccounts) config(cfg config) {
	for _i := range sba {
		sba[_i].config = cfg
	}
}
