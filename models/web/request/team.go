package request

type TeamRequest struct {
	TeamCode     string `validate:"required,min=1,max=100" json:"team_code"`
	TeamName     string `validate:"required,min=1,max=100" json:"team_name"`
	Leader       string `validate:"required,min=1,max=100" json:"leader"`
	ProjectBased bool   `validate:"required" json:"project_based"`
	Level        int    `validate:"required,min=1,max=100" json:"level"`
	Status       string `validate:"required,min=1,max=100" json:"status"`
}
