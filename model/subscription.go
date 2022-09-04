package model

import (
	"net/http"
	"strconv"
	"time"
)

// AlertName ...
type AlertName struct {
	AlertName string `json:"alert_name"`
}

// Subscription ...
type Subscription struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	CancelledAt    time.Time  `json:"cancelled_at"`
	Type           string     `json:"type"`
	SubscriptionID int        `json:"subscription_id"`
	PlanID         int        `json:"plan_id"`
	UserID         int        `json:"user_id"`
	Email          string     `json:"email"`
	Status         string     `json:"status"`
	UpdateURL      string     `json:"update_url"`
	CancelURL      string     `json:"cancel_url"`
}

type SubscriptionHook struct {
	AlertID             string `json:"alert_id"`
	AlertName           string `json:"alert_name"`
	CancelURL           string `json:"cancel_url"`
	CheckoutID          string `json:"checkout_id"`
	Currency            string `json:"currency"`
	Email               string `json:"email"`
	EventTime           string `json:"event_time"`
	LinkedSubscriptions string `json:"linked_subscriptions"`
	MarketingConsent    string `json:"marketing_consent"`
	Passthrough         string `json:"passthrough"`
	Quantity            string `json:"quantity"`
	Source              string `json:"source"`
	Status              string `json:"status"`
	SubscriptionID      string `json:"subscription_id"`
	SubscriptionPlanID  string `json:"subscription_plan_id"`
	UnitPrice           string `json:"unit_price"`
	UpdateURL           string `json:"update_url"`
	UserID              string `json:"user_id"`
	PSignature          string `json:"p_signature"`
}

// ToSubscription ...
func RequestToSub(r *http.Request) *Subscription {
	subID, _ := strconv.Atoi(r.FormValue("subscription_id"))
	planID, _ := strconv.Atoi(r.FormValue("subscription_plan_id"))
	userID, _ := strconv.Atoi(r.FormValue("user_id"))

	status := r.FormValue("status")

	return &Subscription{
		Type:           "pro",
		SubscriptionID: subID,
		PlanID:         planID,
		UserID:         userID,
		Email:          r.FormValue("email"),
		Status:         status,
		UpdateURL:      r.FormValue("update_url"),
		CancelURL:      r.FormValue("cancel_url"),
	}
}

// SubscriptionDTO DTO object for Subscription type
type SubscriptionDTO struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CancelledAt    time.Time `json:"cancelled_at"`
	SubscriptionID int       `json:"subscription_id"`
	PlanID         int       `json:"plan_id"`
	UserID         int       `json:"user_id"`
	Email          string    `json:"email"`
	Status         string    `json:"status"`
	UpdateURL      string    `json:"update_url"`
	CancelURL      string    `json:"cancel_url"`
}

type SubscriptionAuthDTO struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	UpdateURL string `json:"update_url"`
	CancelURL string `json:"cancel_url"`
}

// ToSubscription ...
func ToSubscription(subscriptionDTO *SubscriptionDTO) *Subscription {
	return &Subscription{
		ID:             subscriptionDTO.ID,
		CancelledAt:    subscriptionDTO.CancelledAt,
		SubscriptionID: subscriptionDTO.SubscriptionID,
		PlanID:         subscriptionDTO.PlanID,
		UserID:         subscriptionDTO.UserID,
		Email:          subscriptionDTO.Email,
		Status:         subscriptionDTO.Status,
		UpdateURL:      subscriptionDTO.UpdateURL,
		CancelURL:      subscriptionDTO.CancelURL,
	}
}

// ToSubscriptionDTO ...
func ToSubscriptionDTO(subscription *Subscription) *SubscriptionDTO {
	return &SubscriptionDTO{
		ID:             subscription.ID,
		CancelledAt:    subscription.CancelledAt,
		SubscriptionID: subscription.SubscriptionID,
		PlanID:         subscription.PlanID,
		UserID:         subscription.UserID,
		Email:          subscription.Email,
		Status:         subscription.Status,
		UpdateURL:      subscription.UpdateURL,
		CancelURL:      subscription.CancelURL,
	}
}

// ToSubscriptionAuthDTO ...
func ToSubscriptionAuthDTO(subscription *Subscription) *SubscriptionAuthDTO {
	return &SubscriptionAuthDTO{
		Type:      subscription.Type,
		Status:    subscription.Status,
		UpdateURL: subscription.UpdateURL,
		CancelURL: subscription.CancelURL,
	}
}
