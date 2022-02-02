package models

import (
  "github.com/google/uuid"
  "time"
  "encoding/json"

  // "khaerus-mini-wallet/db/database"
)

type Res1 struct {
	Status string `json:"status"`
	Requirement *json.RawMessage `json:"data,omitempty"`
}

type Res2 struct {
	Token string `json:"token"`
}

type Res3 struct {
	Wallet ShowWallet `json:"wallet"`
}

type Res4 struct {
	Deposit ShowDeposit `json:"deposit"`
}

type Res5 struct {
	Withdraw ShowWithdraw `json:"withdrawn"`
}

type Account struct {
  ID       uuid.UUID   `gorm:"id,primary_key;type:uuid;default:uuid_generate_v4();"`
  Name     string      `gorm:"size:255"`
  Email    string      `gorm:"size:255", unique`
  Phone    string      `gorm:"size:255", unique`
  Password   string    `gorm:"type:text"`
  Created_at time.Time
  Updated_at time.Time
}

type Wallet struct {
  ID            uuid.UUID   `gorm:"id,primary_key;type:uuid;default:uuid_generate_v4();"`
  OwnedBy       string      `gorm:"size:255", unique`
  Status     string      `gorm:"size:255"`
  Balance     int64      `gorm:""`
  EnableAt    time.Time    
  Created_at time.Time
  Updated_at time.Time
}

type Deposit struct {
  ID            uuid.UUID   `gorm:"id,primary_key;type:uuid;default:uuid_generate_v4();"`
  DepositedBy     string     `gorm:"size:255", unique`
  Status     string      `gorm:"size:255"`
  DepositeAt    time.Time
  Amount     int64      `gorm:""`
  ReferenceId     string     `gorm:"size:255", unique`
  Created_at time.Time
  Updated_at time.Time
}

type ShowWallet struct {
  ID          uuid.UUID   `json:"id" bson:"id"`
  OwnedBy     string    `json:"owned_by" bson:"owned_by"`
  Status     string   `json:"status" bson:"status"`
  Balance     int64   `json:"balance" bson:"balance"`
  EnableAt    time.Time    `json:"enable_at" bson:"enable_at"`
}

type ShowDeposit struct {
  ID            uuid.UUID   `json:"id" bson:"id"`
  DepositedBy     string   `json:"deposited_by" bson:"deposited_by"`
  Status     string   `json:"status" bson:"status"`
  DepositeAt    time.Time   `json:"deposited_at" bson:"deposited_at"`
  Amount     int64   `json:"amount" bson:"amount"`
  ReferenceId     string     `json:"reference_id" bson:"reference_id"`
}

type Withdraw struct {
  ID            uuid.UUID   `gorm:"id,primary_key;type:uuid;default:uuid_generate_v4();"`
  WithdrawnBy     string     `gorm:"size:255", unique`
  Status     string      `gorm:"size:255"`
  WithdrawnAt    time.Time
  Amount     int64      `gorm:""`
  ReferenceId     string     `gorm:"size:255", unique`
  Created_at time.Time
  Updated_at time.Time
}

type ShowWithdraw struct {
  ID          uuid.UUID   `json:"id" bson:"id"`
  WithdrawnBy     string    `json:"withdrawn_by" bson:"withdrawn_by"`
  Status     string   `json:"status" bson:"status"`
  WithdrawnAt    time.Time    `json:"withdrawn_at" bson:"withdrawn_at"`
  Amount     int64   `json:"amount" bson:"amount"`
  ReferenceId     string     `json:"reference_id" bson:"reference_id"`
}

func (b *Account) TableName() string {
	return "account"
}

func (b *Wallet) TableName() string {
	return "wallets"
}