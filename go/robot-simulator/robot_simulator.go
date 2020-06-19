package robot

func Left() {
	Step1Robot.Lock()
	defer Step1Robot.Unlock()
	Step1Robot.X = Step1Robot.X - 1
	Step1Robot.Dir = W
}

func Right() {
	Step1Robot.Lock()
	defer Step1Robot.Unlock()
	Step1Robot.X = Step1Robot.X + 1
	Step1Robot.Dir = E
}

func Advance() {
	Step1Robot.Lock()
	defer Step1Robot.Unlock()

	if Step1Robot.Dir == E {
		Step1Robot.X = Step1Robot.X + 1
		return
	}

	if Step1Robot.Dir == W {
		Step1Robot.X = Step1Robot.X - 1
		return
	}

	if Step1Robot.Dir == S {
		Step1Robot.X = Step1Robot.Y - 1
		return
	}

	if Step1Robot.Dir == N {
		Step1Robot.X = Step1Robot.Y + 1
		return
	}
}
