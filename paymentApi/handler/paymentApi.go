package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/plutov/paypal/v3"
	payment "github.com/xing-you-ji/go-container-micro/payment/proto/payment"
	"github.com/xing-you-ji/go-container-micro/paymentApi/proto/paymentApi"
	"go.uber.org/zap"
)

type PaymentApi struct {
	PaymentService payment.PaymentService
}

var (
	ClientID string = "Af1ypmb8ljkjjV-1_OSnHrO2zJ3E7UnIHj8FAEMZ6kN33knndVpAVexJfHj2WzxWANbOH1MijcqqR_hy"
)

func (e *PaymentApi) PayPalRefund(ctx context.Context, req *paymentApi.Request,
	rsp *paymentApi.Response) error {
	// 验证payment
	if err := isOk("payment_id", req); err != nil {
		rsp.StatusCode = 500
		return err
	}

	// 退款号
	if err := isOk("refund_id", req); err != nil {
		rsp.StatusCode = 500
		return err
	}

	// 验证 退款金额
	if err := isOk("money", req); err != nil {
		rsp.StatusCode = 500
		return err
	}

	paymentID, err := strconv.ParseInt(req.Get["payment_id"].Value[0], 10, 64)
	if err != nil {
		zap.L().Error("获取payment_id error", zap.Error(err))
	}

	// 获取支付通道信息
	paymentInfo, err := e.PaymentService.FindPaymentByID(ctx, &payment.PaymentID{PaymentId: paymentID})
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return err
	}
	// SID 获取 paymentInfo.PaymentSid
	// 支付模式
	status := paypal.APIBaseSandBox
	if paymentInfo.PaymentStatus {
		status = paypal.APIBaseLive
	}
	// 退款例子
	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			EmailSubject: req.Get["refund_id"].Value[0] + " zhugeqing 提醒你收款！",
			EmailMessage: req.Get["refund_id"].Value[0] + " 您有一个收款信息！",
			// 每笔转账都要唯一
			SenderBatchID: req.Get["refund_id"].Value[0],
		},
		Items: []paypal.PayoutItem{
			{
				RecipientType: "EMAIL",
				// RecipientWallet: "",
				Receiver: "sb-47t3ft14558633@personal.example.com",
				Amount: &paypal.AmountPayout{
					// 币种
					Currency: "USD",
					Value:    req.Get["money"].Value[0],
				},
				Note:         req.Get["refund_id"].Value[0],
				SenderItemID: req.Get["refund_id"].Value[0],
			},
		},
	}

	// 创建支付客户端
	payPalClient, err := paypal.NewClient(ClientID, paymentInfo.PaymentSid, status)
	if err != nil {
		zap.L().Error("", zap.Error(err))
	}
	// 获取 token
	_, err = payPalClient.GetAccessToken()
	if err != nil {
		zap.L().Error("", zap.Error(err))
	}
	paymentResult, err := payPalClient.CreateSinglePayout(payout)
	if err != nil {
		zap.L().Error("", zap.Error(err))
	}
	zap.L().Info(",", zap.Any("paymentResult", paymentResult))
	rsp.Body = req.Get["refund_id"].Value[0] + "支付成功！"
	return nil
}

func isOk(key string, req *paymentApi.Request) error {
	if _, ok := req.Get[key]; !ok {
		err := errors.New(key + "参数异常")
		zap.L().Error("Get payment_id error", zap.Error(err))
		return err
	}
	return nil
}
