package service

import (
	"encoding/json"
	"errors"
	"food_ordering_api/internal/utils"
	"log"
	"sync"
)

var (
	couponSets *utils.CouponSets
	loadOnce   sync.Once
	loading    = true
	mu         sync.RWMutex
)

func LoadCouponsInBackground() {
	go func() {
		log.Println("Loading coupons in the background...")
		cs, err := utils.LoadCoupons(
			"resources/couponbase1.gz",
			"resources/couponbase2.gz",
			"resources/couponbase3.gz",
		)
		if err != nil {
			log.Printf("Error loading coupons: %v", err)
			setLoading(false)
			return
		}
		SetCouponSets(cs)
		setLoading(false)
		log.Println("Coupons loaded successfully.")
	}()
}

func SetCouponSets(cs *utils.CouponSets) {
	loadOnce.Do(func() {
		couponSets = cs
	})
}

func setLoading(value bool) {
	mu.Lock()
	defer mu.Unlock()
	loading = value
}

func isLoading() bool {
	mu.RLock()
	defer mu.RUnlock()
	return loading
}

func ValidateCoupon(couponCode string) error {
	if isLoading() {
		return errors.New("coupon data not ready, try again later")
	}

	if couponSets == nil {
		return errors.New("coupon data not initialized")
	}

	valid, err := couponSets.IsValidPromoCode(couponCode)
	if !valid {
		return err
	}

	return nil
}

func CreateOrder(data map[string]interface{}) (map[string]interface{}, error) {
	// Validate coupon if provided
	if couponCodeRaw, ok := data["couponCode"]; ok {
		if couponCode, ok2 := couponCodeRaw.(string); ok2 && couponCode != "" {
			if err := ValidateCoupon(couponCode); err != nil {
				return nil, err
			}
		}
	}

	// Post order data to external API
	url := "https://orderfoodonline.deno.dev/api/order"
	respBytes, err := utils.PostData(url, data)
	if err != nil {
		return nil, err
	}

	// Parse JSON response bytes into a map
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(respBytes, &jsonResponse); err != nil {
		return nil, err
	}

	return jsonResponse, nil
}
