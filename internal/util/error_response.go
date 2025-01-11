package util

import "ridhoandhika/backend-api/dto"

// Helper function untuk membuat error response
func ErrorResponse(code, inMessage, enMessage string) dto.BaseResp {
	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: code,
			ErrorMessage: dto.ErrorMessage{
				Id: inMessage,
				En: enMessage,
			},
		},
	}
}
