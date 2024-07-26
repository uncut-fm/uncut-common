package model

import "time"

type SubscriptionEventMetadata struct {
	SubscriptionName          string             `json:"subscription_name"`
	Action                    SubscriptionAction `json:"action"`
	SubscriptionType          SubscriptionType   `json:"subscription_type"`
	UserEmail                 string             `json:"user_email"`
	UserSubscriptionStartedAt time.Time          `json:"user_subscription_started_at"`
	UserSubscriptionExpiresAt time.Time          `json:"user_subscription_expires_at"`
	SubscriptionPeriod        SubscriptionPeriod `json:"subscription_period"`
}

// Action defines the type for the "action" enum field.
type SubscriptionAction string

// Action values.
const (
	ActionCancelAtExpiration SubscriptionAction = "CANCEL_AT_EXPIRATION"
	ActionRenew              SubscriptionAction = "RENEW"
	ActionFirstPurchase      SubscriptionAction = "FIRST_PURCHASE"
	ActionDelete             SubscriptionAction = "DELETE"
	ActionReactivate         SubscriptionAction = "REACTIVATE"
)

// SubscriptionType defines the type for the "type" enum field.
type SubscriptionType string

// SubscriptionType values.
const (
	TypeCollector SubscriptionType = "COLLECTOR"
	TypeCreator   SubscriptionType = "CREATOR"
)

type SubscriptionPeriod string

const (
	SubscriptionPeriodMonthly SubscriptionPeriod = "MONTHLY"
	SubscriptionPeriodYearly  SubscriptionPeriod = "YEARLY"
)
