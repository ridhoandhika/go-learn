// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Generate Token for Authentication",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AuthReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "output_schema": {
                                            "$ref": "#/definitions/dto.AuthResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "401": {
                        "description": "Authentication Failed",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/auth/refresh": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh JWT Token",
                "responses": {
                    "200": {
                        "description": "New JWT Token",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "output_schema": {
                                            "$ref": "#/definitions/dto.AuthResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Authentication Failed",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Registration Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Registration Success",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/education": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "education"
                ],
                "summary": "Insert Education",
                "parameters": [
                    {
                        "description": "Body Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.InsertEducationReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Education Details",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/education/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "education"
                ],
                "summary": "Update Education",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Education ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Body Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateEducationReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Work Experience resp",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/education/{userId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "education"
                ],
                "summary": "Get Education",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Education details successfully retrieved",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "output_schema": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/dto.EducationResp"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "education": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/dto.Education"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/personal-information": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "personal-information"
                ],
                "summary": "Insert Personal Information",
                "parameters": [
                    {
                        "description": "Body Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.InsertPersonalInformationReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Personal Information Details",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/personal-information/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "personal-information"
                ],
                "summary": "Get Personal Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Personal Information ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Personal Information Details",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "output_schema": {
                                            "$ref": "#/definitions/dto.PersonalInformationResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "401": {
                        "description": "Authentication Failed",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "personal-information"
                ],
                "summary": "Update Personal Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Personal Information ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Body Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePersonalInformationReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Personal Information Details",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/work-experience": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work-experience"
                ],
                "summary": "Insert Work Experience",
                "parameters": [
                    {
                        "description": "Body Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.InsertWorkExperienceReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Personal Information Details",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/work-experience/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work-experience"
                ],
                "summary": "Update Work Experience",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Work Experience ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Body Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateWorkExperienceReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Work Experience resp",
                        "schema": {
                            "$ref": "#/definitions/dto.BaseResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        },
        "/api/work-experience/{userId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "work-experience"
                ],
                "summary": "Get Work Experience",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Working Experiences",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.BaseResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "output_schema": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/dto.WorkExperiencesResp"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "work_experience": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/dto.WorkExperience"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorSchema"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AuthReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.AuthResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "dto.BaseResp": {
            "type": "object",
            "properties": {
                "error_schema": {
                    "$ref": "#/definitions/dto.ErrorSchema"
                },
                "output_schema": {
                    "description": "OutputSchema fleksibel (bisa berisi apa saja)"
                }
            }
        },
        "dto.Education": {
            "type": "object",
            "properties": {
                "degree": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "education_id": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "field_of_study": {
                    "type": "string"
                },
                "school_name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "dto.EducationResp": {
            "type": "object",
            "properties": {
                "education": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Education"
                    }
                }
            }
        },
        "dto.ErrorMessage": {
            "type": "object",
            "properties": {
                "en": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "dto.ErrorSchema": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "string"
                },
                "error_message": {
                    "$ref": "#/definitions/dto.ErrorMessage"
                }
            }
        },
        "dto.InsertEducationReq": {
            "type": "object",
            "properties": {
                "degree": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "education_id": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "field_of_study": {
                    "type": "string"
                },
                "school_name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.InsertPersonalInformationReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.InsertWorkExperienceReq": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "job_description": {
                    "type": "string"
                },
                "job_title": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.PersonalInformationResp": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "personal_info_id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateEducationReq": {
            "type": "object",
            "properties": {
                "degree": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "field_of_study": {
                    "type": "string"
                },
                "school_name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "dto.UpdatePersonalInformationReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateWorkExperienceReq": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "job_description": {
                    "type": "string"
                },
                "job_title": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "dto.UserRegisterReq": {
            "type": "object",
            "properties": {
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.WorkExperience": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "job_description": {
                    "type": "string"
                },
                "job_title": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "dto.WorkExperiencesResp": {
            "type": "object",
            "properties": {
                "work_experience": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.WorkExperience"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
