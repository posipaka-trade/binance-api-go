package binance

func (manager *ExchangeManager) checkReqError(err error) {
	//if err == nil {
	//	return
	//}
	//
	//exchErr, isOkay := err.(*exchangeapi.ExchangeError)
	//if !isOkay {
	//	return
	//}
	//
	//if exchErr.Type == exchangeapi.HttpErr {
	//	if exchErr.Code == 429 || exchErr.Code == 418 {
	//		retryAfterStr, isOkay := exchErr.KeysDetails[bncresponse.RetryAfter]
	//		if isOkay {
	//			retryAfter, err := strconv.Atoi(retryAfterStr)
	//			if err == nil {
	//				manager.nextRequestTime = time.Now().Add(time.Duration(retryAfter) * time.Second)
	//			}
	//		}
	//	}
	//}
}
