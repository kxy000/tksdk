package pddddkgoodspromotionrightauth

import (
	"encoding/json"
	"github.com/mimicode/tksdk/pddopensdk/response"
)

// Response pdd.ddk.goods.promotion.right.auth（多多进宝信息流渠道备案授权素材上传接口）
type Response struct {
	response.TopResponse
	GoodsPromotionRightAuthResponse GoodsPromotionRightAuthResponse `json:"goods_promotion_right_auth_response"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.ErrorCode = -1
		t.ErrorResponse.ErrorMsg = unmarshal.Error()
	}
}

type GoodsPromotionRightAuthResponse struct {
	RequestID string `json:"request_id"`
	Reason    string `json:"reason"`
	Result    bool   `json:"result"`
}
