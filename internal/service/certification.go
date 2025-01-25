package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
)

type certificationService struct {
	certificationRepository domain.CertificationRepository
}

func Certification(certificationRepository domain.CertificationRepository) domain.CertificationService {
	return &certificationService{
		certificationRepository: certificationRepository,
	}
}

func (s certificationService) FindByUserId(ctx context.Context, userId string) (dto.BaseResp, error) {
	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		return util.ErrorResponse("400", "Invalid UUID format", "Invalid UUID format"), nil
	}

	data, err := s.certificationRepository.FindByUserId(ctx, parsedUserID)
	if err != nil || len(data) == 0 {
		return dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "200",
				ErrorMessage: dto.ErrorMessage{
					Id: "Sukses",
					En: "Success",
				},
			},
			OutputSchema: dto.CertificationResp{
				Certifications: []dto.Certification{},
			},
		}, nil
	}

	var response dto.CertificationResp
	var certificationList []dto.Certification
	// Convert data ke WorkExperience
	for _, certification := range data {
		var issueDateStr, expirationDateStr string

		// Mengkonversi IssueDate jika tidak nil
		if certification.IssueDate != nil {
			issueDateStr = certification.IssueDate.Format("2006-01-02")
		} else {
			issueDateStr = "" // Atau bisa menggunakan nilai default seperti "N/A"
		}

		// Mengkonversi ExpirationDate jika tidak nil
		if certification.ExpirationDate != nil {
			expirationDateStr = certification.ExpirationDate.Format("2006-01-02")
		} else {
			expirationDateStr = "" // Atau bisa menggunakan nilai default seperti "N/A"
		}

		// Tambahkan objek Education ke dalam educationList
		certificationList = append(certificationList, dto.Certification{
			CertificationID: certification.CertificationID,
			Name:            certification.Name,
			Body:            certification.Body,
			CredentialID:    certification.CredentialID,
			IssueDate:       issueDateStr,
			ExpirationDate:  expirationDateStr,
		})
	}

	response = dto.CertificationResp{
		Certifications: certificationList,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: response,
	}, nil
}

func (s certificationService) Insert(ctx context.Context, req dto.InsertCertificationReq) (dto.BaseResp, error) {
	_, err := s.certificationRepository.Insert(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), err
	}

	return util.ErrorResponse("200", "Sukses", "Success"), nil
}

func (s certificationService) Update(ctx context.Context, certificationId string, req dto.UpdateCertificationReq) (dto.BaseResp, error) {
	parsedSkillId, err := uuid.Parse(certificationId)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan tidak valid", "Bad request"), nil
	}

	_, err = s.certificationRepository.Update(ctx, parsedSkillId, req)
	if err != nil {
		return util.ErrorResponse("400", "Gagal", "Failed"), nil
	}

	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
