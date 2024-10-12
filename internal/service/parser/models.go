package parser

type VacancyID struct {
	ID string `json:"id"`
}

type PageResp struct {
	Pages int         `json:"pages"`
	Found int         `json:"found"`
	Items []VacancyID `json:"items"`
}

type VacancySkill struct {
	Name string `json:"name"`
}

type Vacancy struct {
	Description string         `json:"description"`
	Skills      []VacancySkill `json:"key_skills"`
}
