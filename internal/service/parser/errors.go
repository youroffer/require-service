package parser

type pageFetchError struct {
	Page int
	Err  error
}

func(e *pageFetchError) Error() string {
	return e.Err.Error()
}

type vacancyFetchError struct {
	VacancyID string
	Err       error
}

func (e *vacancyFetchError) Error() string {
	return e.Err.Error()
}