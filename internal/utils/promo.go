package utils

import (
	"bufio"
	"compress/gzip"
	"errors"
	"os"
	"strings"
)

type CouponSets struct {
	Set1 map[string]struct{}
	Set2 map[string]struct{}
	Set3 map[string]struct{}
}

func LoadCoupons(file1, file2, file3 string) (*CouponSets, error) {
	set1, err := loadCouponFile(file1)
	if err != nil {
		return nil, err
	}
	set2, err := loadCouponFile(file2)
	if err != nil {
		return nil, err
	}
	set3, err := loadCouponFile(file3)
	if err != nil {
		return nil, err
	}
	return &CouponSets{Set1: set1, Set2: set2, Set3: set3}, nil
}

func loadCouponFile(filepath string) (map[string]struct{}, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	set := make(map[string]struct{})
	scanner := bufio.NewScanner(gz)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			set[line] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return set, nil
}

func (cs *CouponSets) IsValidPromoCode(code string) (bool, error) {
	if len(code) < 8 || len(code) > 10 {
		return false, errors.New("promo code length must be between 8 and 10 characters")
	}

	count := 0
	if _, ok := cs.Set1[code]; ok {
		count++
	}
	if _, ok := cs.Set2[code]; ok {
		count++
	}
	if _, ok := cs.Set3[code]; ok {
		count++
	}

	if count >= 2 {
		return true, nil
	}
	return false, errors.New("promo code not valid in at least two coupon files")
}
