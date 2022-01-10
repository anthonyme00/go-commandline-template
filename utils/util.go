package utils

type Utility struct {
	Logger *Logger
}

func NewUtility(filename string) *Utility {
	logger, err := newLogger(filename)
	if err != nil {
		panic(err)
	}

	return &Utility{
		Logger: logger,
	}
}

func (u *Utility) Clean() {
	u.Logger.clean()
}
