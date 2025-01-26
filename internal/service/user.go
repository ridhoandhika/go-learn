package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
)

type userService struct {
	userRepository domain.UserRepository
}

func User(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) GetUser(ctx context.Context, userID string) (dto.BaseResp, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.FindByID(ctx, parsedUserID)
	if err != nil {
		return util.ErrorResponse("404", "Tidak ditemukan", "Not found"), nil
	}

	userResp := dto.UserData{
		ID:       user.ID,
		Fullname: user.Fullname,
		Usename:  user.Username,
		Phone:    user.Phone,
	}
	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: userResp,
	}, nil
}

func (u userService) GetEducation(ctx context.Context, userID string) (dto.BaseResp, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.UserWithEducation(ctx, parsedUserID)
	if err != nil {
		return util.ErrorResponse("200", "Sukses", "Success"), nil
	}

	// Menyiapkan daftar Education
	var educationList []dto.Education
	for _, education := range user.Education {
		educationList = append(educationList, dto.Education{
			EducationID:  education.EducationID,
			Degree:       education.Degree,
			SchoolName:   education.SchoolName,
			FieldOfStudy: education.FieldOfStudy,
			Description:  education.Description,
			StartDate:    education.StartDate,
			EndDate:      education.EndDate,
		})
	}

	// Menyiapkan UserEducationResp untuk OutputSchema
	userEducationResp := dto.UserEducationResp{
		Education: educationList,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: userEducationResp,
	}, nil
}

func (u userService) GetPersonalInformation(ctx context.Context, userID string) (dto.BaseResp, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.UserWithPersonalInformation(ctx, parsedUserID)
	if err != nil {
		return dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "200",
				ErrorMessage: dto.ErrorMessage{
					Id: "Success",
					En: "Success",
				},
			},
			OutputSchema: dto.UserPersonalInformationResp{
				PersonalInformation: dto.PersonalInformationResp{},
			},
		}, nil
	}

	personalInformation := dto.PersonalInformationResp{
		PersonalInfoID: user.PersonalInfo.PersonalInfoID,
		FirstName:      user.PersonalInfo.FirstName,
		LastName:       user.PersonalInfo.LastName,
		PhoneNumber:    user.PersonalInfo.PhoneNumber,
		Email:          user.PersonalInfo.Email,
		Address:        user.PersonalInfo.Address,
		Summary:        user.PersonalInfo.Summary,
		DateOfBirth:    user.PersonalInfo.DateOfBirth,
	}

	personalInformationList := dto.UserPersonalInformationResp{
		PersonalInformation: personalInformation,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: personalInformationList,
	}, nil
}

func (u userService) GetWorkExperience(ctx context.Context, userID string) (dto.BaseResp, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.UserWithWorkExperience(ctx, parsedUserID)
	if err != nil {
		return dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "200",
				ErrorMessage: dto.ErrorMessage{
					Id: "Success",
					En: "Success",
				},
			},
			OutputSchema: dto.UserWorkExperienceResp{
				WorkExperience: []dto.WorkExperience{},
			},
		}, nil
	}

	var workExperienceList []dto.WorkExperience
	for _, data := range user.WorkExperience {
		workExperienceList = append(workExperienceList, dto.WorkExperience{
			WorkExperienceID: data.WorkExperienceID,
			JobTitle:         data.JobTitle,
			CompanyName:      data.CompanyName,
			StartDate:        data.StartDate,
			EndDate:          data.EndDate,
			JobDescription:   data.JobDescription,
		})
	}

	workExperienceResp := dto.UserWorkExperienceResp{
		WorkExperience: workExperienceList,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: workExperienceResp,
	}, nil
}

func (u userService) GetSkill(ctx context.Context, userID string) (dto.BaseResp, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.UserWithWorkExperience(ctx, parsedUserID)
	if err != nil {
		return dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "200",
				ErrorMessage: dto.ErrorMessage{
					Id: "Success",
					En: "Success",
				},
			},
			OutputSchema: dto.UserSkillResp{
				Skill: []dto.Skill{},
			},
		}, nil
	}

	var skillList []dto.Skill
	for _, data := range user.Skill {
		skillList = append(skillList, dto.Skill{
			SkillID: data.SkillID,
			Name:    data.Name,
			Level:   string(data.Level),
		})
	}

	skillResp := dto.UserSkillResp{
		Skill: skillList,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: skillResp,
	}, nil
}

func (u userService) GetCertification(ctx context.Context, userID string) (dto.BaseResp, error) {
	parsedUserID, _ := uuid.Parse(userID)
	user, err := u.userRepository.UserWithCertification(ctx, parsedUserID)
	if err != nil {
		return dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "200",
				ErrorMessage: dto.ErrorMessage{
					Id: "Success",
					En: "Success",
				},
			},
			OutputSchema: dto.UserCertificationResp{
				Certification: []dto.Certification{},
			},
		}, nil
	}

	var certificationList []dto.Certification
	for _, data := range user.Certification {
		var issueDateStr, expirationDateStr string

		// Mengkonversi IssueDate jika tidak nil
		if data.IssueDate != nil {
			issueDateStr = data.IssueDate.Format("2006-01-02")
		} else {
			issueDateStr = "" // Atau bisa menggunakan nilai default seperti "N/A"
		}

		// Mengkonversi ExpirationDate jika tidak nil
		if data.ExpirationDate != nil {
			expirationDateStr = data.ExpirationDate.Format("2006-01-02")
		} else {
			expirationDateStr = "" // Atau bisa menggunakan nilai default seperti "N/A"
		}
		certificationList = append(certificationList, dto.Certification{
			CertificationID: data.CertificationID,
			Name:            data.Name,
			Body:            data.Body,
			CredentialID:    data.Body,
			IssueDate:       issueDateStr,
			ExpirationDate:  expirationDateStr,
		})
	}

	certificationResp := dto.UserCertificationResp{
		Certification: certificationList,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: certificationResp,
	}, nil
}
