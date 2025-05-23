// Code generated by goctl. DO NOT EDIT.
package types

type ChatInterviewReq struct {
	Answer    string `json:"answer"`
	SessionID int64  `json:"session_id"`
}

type ChatNewResp struct {
	SessionID int64 `json:"session_id"`
}

type CheckResumeReq struct {
	Condition string `json:"condition"`
	NeedNum   int    `json:"need_num"`
	FolderID  int64  `json:"folder_id"`
}

type CommonListReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type CommonListResp struct {
	Total int64 `json:"total"`
}

type CreateInterview struct {
	QuestionNum int    `json:"question_num,optional"`
	Hc          string `json:"hc,optional"`
	Level       int    `json:"level,optional"`
	PdfUrl      string `json:"pdf_url,optional"`
	Answer      string `json:"answer,optional"`
	SessionID   int64  `json:"session_id"`
}

type CreqteResumeFolderReq struct {
	ResumeName string `json:"name"`
}

type DeleteResumeFolderReq struct {
	FolderId int64 `path:"id"` // 要删除的文件夹Id
}

type DeleteResumeReq struct {
	ResumeId int64 `path:"id"` // 要删除的简历Id
}

type EvaluateResumeReq struct {
	ResumeUrl string `json:"resume_url"` // 简历地址
	Content   string `json:"content"`    // 简历内容
	Jd        string `json:"jd"`         // 职位描述
}

type FolderInfo struct {
	FolderId   int64  `json:"id"`
	FolderName string `json:"name"`
}

type GetResumeFolderListResp struct {
	CommonListResp
	List []FolderInfo `json:"list"` // 简历文件夹信息
}

type GetResumeListReq struct {
	CommonListReq
	FolderID int64 `form:"folder_id"` // 要查找的文件夹id，如果为空就是查找所有简历
}

type GetResumeListResp struct {
	CommonListResp
	List []ResumeInfo `json:"list"`
}

type ResumeInfo struct {
	ResumeId   int64  `json:"id"`
	ResumeName string `json:"name"`
	ResumeUrl  string `json:"url"`
	FolderId   int64  `json:"folder_id"`
	UploadTime string `json:"upload_time"`
	ResumeSize int64  `json:"size"`
	UserId     int64  `json:"user_id"`
}

type SSEReq struct {
	Data string `json:"data"`
}

type UpdateResumeFolderReq struct {
	FolderId   int64  `json:"folder_id"`
	FolderName string `json:"name"`
}

type UploadResumeResp struct {
	ResumeId int64  `json:"resume_id"`
	Address  string `json:"address"`
}

type UploadReusmeReq struct {
	FolderId int64 `form:"folder_id"` // 文件夹id
}
