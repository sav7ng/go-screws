package geo

import (
	"github.com/shopspring/decimal"
)

func toDecimal(val string) decimal.Decimal {
	fromString, _ := decimal.NewFromString(val)
	return fromString
}

func BoxCenter(latitudeRange [2]decimal.Decimal, longitudeRange [2]decimal.Decimal) (centerLat decimal.Decimal, centerLng decimal.Decimal) {
	centerLat = latitudeRange[0].Add(latitudeRange[1]).Div(decimal.NewFromInt(2))
	centerLng = longitudeRange[0].Add(longitudeRange[1]).Div(decimal.NewFromInt(2))
	if centerLng.GreaterThan(decimal.NewFromInt(180)) {
		centerLng = centerLng.Sub(decimal.NewFromInt(360))
	}
	return
}

func NewGeoHash(lat string, lng string, desiredPrecision int) (int64, [2]decimal.Decimal, [2]decimal.Decimal) {
	latitudeRange := [2]decimal.Decimal{decimal.NewFromFloat(-90.0), decimal.NewFromFloat(90.0)}
	longitudeRange := [2]decimal.Decimal{decimal.NewFromFloat(-180), decimal.NewFromFloat(180)}
	significantBits := 0
	var bits int64 = 0
	isEvenBit := true
	for {
		if significantBits == desiredPrecision {
			break
		}
		if isEvenBit {
			bits = divideRangeEncode(toDecimal(lng), &longitudeRange, bits)
		} else {
			bits = divideRangeEncode(toDecimal(lat), &latitudeRange, bits)
		}
		significantBits++
		isEvenBit = !isEvenBit
	}
	bits <<= 64 - desiredPrecision
	return bits, latitudeRange, longitudeRange
}

func divideRangeEncode(val decimal.Decimal, r *[2]decimal.Decimal, bits int64) int64 {
	mid := r[0].Add(r[1]).Div(decimal.NewFromInt(2))
	if val.GreaterThanOrEqual(mid) {
		bits <<= 1
		bits |= 1
		r[0] = mid
	} else {
		bits <<= 1
		r[1] = mid
	}
	return bits
}
