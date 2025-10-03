package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// added test grade C
func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam", 71, Exam)
	gradeCalculator.AddGrade("essay", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, actual_value)
	}
}

// added test grade D
func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam", 61, Exam)
	gradeCalculator.AddGrade("essay", 63, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay) // fix: changed grades to get F

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// added test for string type
func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected 'assignment'; got '%s'", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected 'exam'; got '%s'", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected 'essay'; got '%s'", Essay.String())
	}
}
