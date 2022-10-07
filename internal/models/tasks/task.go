package tasks

import "math/rand"

const (
	ANS_COUNT = 4
	T_MAX     = 20
	T_MIN     = 10
	S_MAX     = 5
	S_MIN     = 2
)

// task A + B = C with possible answers
type Task struct {
	A, B, C int64
	Answers []int64
	Choosen int64
}

func getNum() int64 {
	return rand.Int63n(T_MAX-T_MIN) + T_MIN
}

func getShift() int64 {
	return rand.Int63n(S_MAX-S_MIN) + S_MIN
}

func New() *Task {

	A := getNum()
	B := getNum()
	C := A + B

	s_ans := [ANS_COUNT]int64{
		C,
		C + getShift(),
		C - getShift(),
		C + 2*getShift(),
	}

	p := rand.Perm(ANS_COUNT)

	us_ans := make([]int64, ANS_COUNT)

	for i, p := range p {
		us_ans[i] = s_ans[p]
	}

	t := Task{
		A:       A,
		B:       B,
		C:       C,
		Answers: us_ans,
		Choosen: -1,
	}

	return &t
}

func (t *Task) IsRight() bool {
	if t.Choosen < 0 || t.Choosen >= ANS_COUNT {
		return false
	}
	return t.Answers[t.Choosen] == t.C
}
