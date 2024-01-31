package request

import (
	"net/url"
)

// pdd.ddk.goods.promotion.right.auth（多多进宝信息流渠道备案授权素材上传接口）
// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.promotion.right.auth

type PddDdkGoodsPromotionRightAuthRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *PddDdkGoodsPromotionRightAuthRequest) CheckParameters() {

}

// 添加请求参数

func (tk *PddDdkGoodsPromotionRightAuthRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

// 返回接口名称

func (tk *PddDdkGoodsPromotionRightAuthRequest) GetApiName() string {
	return "pdd.ddk.goods.search"
}

// 返回请求参数

func (tk *PddDdkGoodsPromotionRightAuthRequest) GetParameters() url.Values {
	return *tk.Parameters
}
