package log

func Info(args... interface{}) {
	e := makeEntry()
	e.Info(args...)
}

func (this *Entry)Info(args... interface{}) {
	this.print(LevelInfo, LevelInfoColor, args...)
}