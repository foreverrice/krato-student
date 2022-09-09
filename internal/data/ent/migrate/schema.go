// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StoreBalanceAccountsColumns holds the columns for the "store_balance_accounts" table.
	StoreBalanceAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "account_no", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "store_code", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "upper_organ_no", Type: field.TypeString, Nullable: true},
		{Name: "pwd", Type: field.TypeString, Nullable: true},
		{Name: "pwd_salt", Type: field.TypeString, Nullable: true},
		{Name: "balance_fee", Type: field.TypeFloat64, Nullable: true},
		{Name: "total_charge_fee", Type: field.TypeFloat64, Nullable: true},
		{Name: "is_deleted", Type: field.TypeInt8, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
	}
	// StoreBalanceAccountsTable holds the schema information for the "store_balance_accounts" table.
	StoreBalanceAccountsTable = &schema.Table{
		Name:       "store_balance_accounts",
		Columns:    StoreBalanceAccountsColumns,
		PrimaryKey: []*schema.Column{StoreBalanceAccountsColumns[0]},
	}
	// StoreBalanceChangeDetailsColumns holds the columns for the "store_balance_change_details" table.
	StoreBalanceChangeDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "account_id", Type: field.TypeInt32},
		{Name: "store_code", Type: field.TypeString},
		{Name: "type", Type: field.TypeInt8},
		{Name: "in_batch_no", Type: field.TypeString},
		{Name: "third_pay_type", Type: field.TypeString},
		{Name: "third_pay_organ", Type: field.TypeString},
		{Name: "third_pay_no", Type: field.TypeString},
		{Name: "bc_bind_acc_id", Type: field.TypeInt32},
		{Name: "cash_account_no", Type: field.TypeString},
		{Name: "cash_principal", Type: field.TypeString},
		{Name: "cash_bank", Type: field.TypeString},
		{Name: "transaction_no", Type: field.TypeString},
		{Name: "transaction_at", Type: field.TypeTime},
		{Name: "before_fee", Type: field.TypeFloat64},
		{Name: "change_fee", Type: field.TypeFloat64},
		{Name: "after_fee", Type: field.TypeFloat64},
		{Name: "operator_no", Type: field.TypeString},
		{Name: "flow_no", Type: field.TypeString},
		{Name: "check_state", Type: field.TypeInt8},
		{Name: "check_at", Type: field.TypeTime},
		{Name: "is_deleted", Type: field.TypeInt8},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// StoreBalanceChangeDetailsTable holds the schema information for the "store_balance_change_details" table.
	StoreBalanceChangeDetailsTable = &schema.Table{
		Name:       "store_balance_change_details",
		Columns:    StoreBalanceChangeDetailsColumns,
		PrimaryKey: []*schema.Column{StoreBalanceChangeDetailsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StoreBalanceAccountsTable,
		StoreBalanceChangeDetailsTable,
	}
)

func init() {
}
