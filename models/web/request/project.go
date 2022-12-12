package request

type ProjectRequest struct {
	ProjectCode string `validate:"required,min=1,max=100" json:"project_code"`
	ProjectName string `validate:"required,min=1,max=100" json:"project_name"`
	Deadline    string `json:"deadline"`
	Status      string `validate:"required,min=1,max=100" json:"status"`
}

type ProjectRequestUpdate struct {
	ProjectCode string `json:"project_code"`
	ProjectName string `validate:"required,min=1,max=100" json:"project_name"`
	Deadline    string `json:"deadline"`
	Status      string `validate:"required,min=1,max=100" json:"status"`
}