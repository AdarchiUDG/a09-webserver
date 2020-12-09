package school

type Grade struct {
	Student, Class string
	Value float64
}


type School struct {
	Students map[string]map[string]float64
}

func getAverage(student map[string]float64) float64 {
	var total float64

	for _, grade := range student {
		total += grade
	}

	return total / float64(len(student))
}

func (t *School) AddGrade(grade Grade) {
	if t.Students[grade.Student] == nil {
		t.Students[grade.Student] = make(map[string]float64)
	}
	t.Students[grade.Student][grade.Class] = grade.Value
} 

func (t *School) GetStudentAverage(name string) float64 {
	student, ok := t.Students[name]
	if !ok {
		return 0
	} else {
		return getAverage(student)
	}
}

func (t *School) GetGeneralAverage() float64 {
	var total float64
	for _, student := range t.Students {
		total += getAverage(student)
	}

	if len(t.Students) > 0 {
		return total / float64(len(t.Students))
	}

	return float64(0)
}

func (t *School) GetClassAverage(name string) float64 {
	var total float64
	var found bool
	var totalClasses int
	for _, student := range t.Students {
		grade, ok := student[name]
		if ok {
			total += grade
			totalClasses += 1
			found = true
		}
	}

	if found {
		return total / float64(totalClasses)
	} else {
		return 0
	}
	
}