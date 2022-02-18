package services

import (
	"context"
	"fmt"
	"strconv"

	"basic_golang/internal/adapter"
	zaplogger "basic_golang/internal/adapter/zap"
	authEntity "basic_golang/internal/domain/auth/entity"
	"basic_golang/internal/domain/fetch/entity"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type tempData struct {
	provinsi string
	amount   int
	tahun    string
	minggu   string
}

func (s *fetchDomain) GetResources(ctx context.Context, jwtToken string) (res []entity.Resource, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Fetch_GetResources")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	user, err := s.authDomain.CheckToken(ctx, jwtToken)
	if err != nil {
		logger.Error("error when GetResources|CheckToken", zap.Error(err))
		return res, fmt.Errorf("Not Authorized")
	}
	if (user == authEntity.User{}) {
		return res, fmt.Errorf("Not Authorized")
	}

	resources, err := s.fetchRepository.GetResources(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetResources", zap.Error(err))
		return res, err
	}

	usdIdrCurrency, err := s.fetchRepository.GetCurrencyIDRtoUSD(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetCurrencyUSDToIDR", zap.Error(err))
		return res, err
	}

	for _, resource := range resources {
		if (resource != entity.Resource{}) {
			var price float64
			if resource.Price != "" {
				price, err = strconv.ParseFloat(resource.Price, 32)
				if err != nil {
					logger.Error("error when GetResources|ParseFloat Price", zap.Error(err))
				}
			}
			priceUsd := price * usdIdrCurrency.IDRtoUSD

			res = append(res, entity.Resource{
				UUID:         resource.UUID,
				Komoditas:    resource.Komoditas,
				AreaProvinsi: resource.Komoditas,
				AreaKota:     resource.AreaKota,
				Size:         resource.Size,
				Price:        resource.Price,
				ParsedDate:   resource.ParsedDate,
				Timestamp:    resource.Timestamp,
				PriceUSD:     fmt.Sprintf("$%f", priceUsd),
			})
		}
	}

	return res, nil
}

func (s *fetchDomain) GetResourcesAdmin(ctx context.Context, jwtToken string) (res []entity.ResourceData, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "services_Fetch_GetResources")
	defer span.Finish()

	logger := zaplogger.For(ctx)

	// 1. Check token
	user, err := s.authDomain.CheckToken(ctx, jwtToken)
	if err != nil {
		logger.Error("error when GetResources|CheckToken", zap.Error(err))
		return res, fmt.Errorf("Not Authorized")
	}

	if (user == authEntity.User{}) {
		logger.Error("error when GetResources|CheckToken User not found", zap.Error(err))
		return res, fmt.Errorf("Not Authorized")
	}
	if user.Role != "admin" {
		logger.Error("error when GetResources|CheckToken Invalid User Role", zap.Error(err))
		return res, fmt.Errorf("Not Authorized - Invalid User Role")
	}

	// 2. fetch data resources
	resources, err := s.fetchRepository.GetResources(ctx)
	if err != nil {
		logger.Error("error when GetResources|GetResources", zap.Error(err))
		return res, err
	}

	// 3. Parse data
	var dataParsed []tempData
	for _, val := range resources {
		if val.ParsedDate == "" || val.Price == "" || val.Size == "" || val.AreaProvinsi == "" {
			continue
		}
		dateTime := adapter.ParseStringToTime(val.ParsedDate)
		year, week := dateTime.ISOWeek()

		price, _ := strconv.Atoi(val.Price)
		size, _ := strconv.Atoi(val.Size)
		amount := price * size

		temp := tempData{
			provinsi: val.AreaProvinsi,
			amount:   amount,
			tahun:    fmt.Sprintf("Tahun %s", strconv.Itoa(year)),
			minggu:   fmt.Sprintf("Minggu ke %s", strconv.Itoa(week)),
		}

		dataParsed = append(dataParsed, temp)
	}

	// 4. Agregasi data
	tempMap := make(map[string]map[string]map[string]int)
	for _, val := range dataParsed {
		// Create data set if data provinsi is not exist yet
		if prov, ok := tempMap[val.provinsi]; !ok {
			minggu := map[string]int{val.minggu: val.amount}
			tahun := map[string]map[string]int{val.tahun: minggu}
			tempMap[val.provinsi] = tahun
		} else {
			// Create data set if data provinsi is not exist yet
			if year, ok := prov[val.tahun]; !ok {
				minggu := map[string]int{val.minggu: val.amount}
				prov[val.tahun] = minggu
			} else {
				// Create data set if data profit is not exist yet
				// Add the amount if data profit is exist already
				if profit, ok := year[val.minggu]; !ok {
					year[val.minggu] = val.amount
				} else {
					year[val.minggu] += profit
				}
			}
		}
	}
	var result []entity.ResourceData

	for key, val := range tempMap {
		data := entity.ResourceData{
			AreaProvinsi: key,
			Profit:       val,
			Max:          s.FindMaxProfit(val),
			Min:          s.FindMinProfit(val),
			Avg:          s.FindAvgProfit(val),
			Median:       s.FindMedianProfit(val),
		}

		result = append(result, data)
	}

	return result, nil

}

func (s *fetchDomain) FindMaxProfit(data map[string]map[string]int) float64 {
	var max int
	for _, val := range data {
		for _, amount := range val {
			if amount >= max {
				max = amount
			}
		}
	}

	return float64(max)
}

func (s *fetchDomain) FindMinProfit(data map[string]map[string]int) float64 {
	min := int(^uint(0) >> 1)
	for _, val := range data {
		for _, amount := range val {
			if amount <= min {
				min = amount
			}
		}
	}

	return float64(min)
}

func (s *fetchDomain) FindAvgProfit(data map[string]map[string]int) float64 {
	var sum, counter int
	for _, val := range data {
		for _, amount := range val {
			sum += amount
			counter++
		}
	}

	return float64(sum / counter)
}

func (s *fetchDomain) FindMedianProfit(data map[string]map[string]int) float64 {
	var arr []int
	for _, val := range data {
		for _, amount := range val {
			arr = append(arr, amount)
		}
	}

	counter := len(arr)

	if counter+1%2 == 0 {
		a := arr[(counter / 2)]
		b := arr[(counter/2)+1]
		return float64((a + b) / 2)
	} else {
		return float64(arr[counter/2])
	}
}
