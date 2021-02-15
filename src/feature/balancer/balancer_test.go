package balancer

import "testing"

type TestDataItem struct {
	inputs   []*UsdAssetWrapper
	result   bool
	buy      float64
	sell     float64
	hasError bool
}

func TestShouldBuy(test *testing.T) {

	dataList := []TestDataItem{
		TestDataItem{
			inputs: []*UsdAssetWrapper{
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "DOT",
						Weight: 0.5,
					},
					assetAmount:     50,
					usdPricePerUnit: 1,
					totalUsdPrice:   50,
				},
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "USDCT",
						Weight: 0.5,
					},
					assetAmount:     50,
					usdPricePerUnit: 1,
					totalUsdPrice:   50,
				},
			},
			result:   false,
			buy:      0,
			sell:     0,
			hasError: false,
		},
		TestDataItem{
			inputs: []*UsdAssetWrapper{
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "DOT",
						Weight: 0.5,
					},
					assetAmount:     60,
					usdPricePerUnit: 1,
					totalUsdPrice:   60,
				},
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "USDCT",
						Weight: 0.5,
					},
					assetAmount:     50,
					usdPricePerUnit: 1,
					totalUsdPrice:   50,
				},
			},
			result:   false,
			buy:      0,
			sell:     -4.999999999999996,
			hasError: false,
		},
		TestDataItem{
			inputs: []*UsdAssetWrapper{
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "DOT",
						Weight: 0.5,
					},
					assetAmount:     80,
					usdPricePerUnit: 1,
					totalUsdPrice:   80,
				},
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "USDCT",
						Weight: 0.5,
					},
					assetAmount:     50,
					usdPricePerUnit: 1,
					totalUsdPrice:   50,
				},
			},
			result:   false,
			buy:      0,
			sell:     -15.000000000000004,
			hasError: false,
		},
		TestDataItem{
			inputs: []*UsdAssetWrapper{
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "DOT",
						Weight: 0.5,
					},
					assetAmount:     40,
					usdPricePerUnit: 1,
					totalUsdPrice:   40,
				},
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "USDCT",
						Weight: 0.5,
					},
					assetAmount:     50,
					usdPricePerUnit: 1,
					totalUsdPrice:   50,
				},
			},
			result:   true,
			buy:      5.000000000000002,
			sell:     0,
			hasError: false,
		},
		TestDataItem{
			inputs: []*UsdAssetWrapper{
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "DOT",
						Weight: 0.7,
					},
					assetAmount:     2,
					usdPricePerUnit: 4.20,
					totalUsdPrice:   8.40,
				},
				&UsdAssetWrapper{
					asset: AssetBalancer{
						Symbol: "USDCT",
						Weight: 0.3,
					},
					assetAmount:     50,
					usdPricePerUnit: 1,
					totalUsdPrice:   50,
				},
			},
			result:   true,
			buy:      32.48,
			sell:     0,
			hasError: false,
		},
	}

	for index, item := range dataList {
		result := ShouldBuy(item.inputs[0], item.inputs[1])

		if result != item.result {
			test.Errorf("[Index: %v] ShouldBuy with args %v : Failed, expected false but got %v", index, item.inputs, item.result)
		}

		buy := AmountToBuy(item.inputs[0], item.inputs[1])

		if buy != item.buy {
			test.Errorf("[Index: %v] AmountToBuy with args %v : Failed, expected %v but got %v", index, item.inputs, item.buy, buy)
		}

		sell := AmountToSell(item.inputs[0], item.inputs[1])

		if sell != item.sell {
			test.Errorf("[Index: %v] AmountToSell with args %v : Failed, expected %v but got %v", index, item.inputs, item.sell, sell)
		}
	}

}
