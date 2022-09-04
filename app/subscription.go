package app

import (
	"net/http"
	"strconv"

	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/model"
)

// CreateSubscription creates a subscription and saves it to the store
func CreateSubscription(s db.Store, r *http.Request) (int, string) {
	_, err := s.Subscriptions().FindByEmail(r.FormValue("email"))
	if err == nil {
		message := "Subscription already exist!"
		return http.StatusBadRequest, message
	}

	_, err = s.Subscriptions().Create(model.RequestToSub(r))
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "Subscription created successfully."
}

func UpdateSubscription(s db.Store, r *http.Request) (int, string) {
	subscription, err := s.Subscriptions().FindByEmail(r.FormValue("email"))
	if err != nil {
		return http.StatusNotFound, err.Error()
	}

	subscriptionID, err := strconv.Atoi(r.FormValue("subscription_id"))
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	planID, err := strconv.Atoi(r.FormValue("subscription_plan_id"))
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	subscription.Type = "pro"
	subscription.SubscriptionID = subscriptionID
	subscription.PlanID = planID
	subscription.UserID = userID
	subscription.Status = r.FormValue("status")

	_, err = s.Subscriptions().Update(subscription)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "Subscription updated successfully."
}

func CancelSubscription(s db.Store, r *http.Request) (int, string) {
	subscription, err := s.Subscriptions().FindByEmail(r.FormValue("email"))
	if err != nil {
		return http.StatusNotFound, err.Error()
	}

	err = s.Subscriptions().Delete(subscription.ID)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "Subscription cancelled."
}

func PaymentSucceedSubscription(s db.Store, r *http.Request) (int, string) {
	subscription, err := s.Subscriptions().FindByEmail(r.FormValue("email"))
	if err != nil {
		return http.StatusNotFound, err.Error()
	}

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	_, err = s.Subscriptions().Update(subscription)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "Subscription payment succeeded."
}

func PaymentFailedSubscription(s db.Store, r *http.Request) (int, string) {
	subscription, err := s.Subscriptions().FindByEmail(r.FormValue("email"))
	if err != nil {
		return http.StatusNotFound, err.Error()
	}

	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	subscription.Status = r.FormValue("status")

	_, err = s.Subscriptions().Update(subscription)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return http.StatusOK, "Subscription payment failed."
}
