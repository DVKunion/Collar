package log

// Chain
// log chain call
type Chain interface {
	Clone() Log
	Skip(skip int) Log
	Line(long ...bool) Log
	Prefix(prefix string) Log
	Cat(category interface{}) Log
}

func (l *defaultLog) Clone() Log {
	log := *l
	log.Logger = l.Logger.Clone()
	return &log
}

func (l *defaultLog) Skip(skip int) Log {
	log := *l
	log.Logger = l.Logger.Skip(skip)
	return &log
}

func (l *defaultLog) Line(long ...bool) Log {
	log := *l
	log.Logger = l.Logger.Line(long...)
	return &log
}

func (l *defaultLog) Prefix(prefix string) Log {
	log := *l
	log.Logger = l.Logger.Clone()
	log.Logger.SetPrefix(prefix)
	return &log
}

func (l *defaultLog) Cat(category interface{}) Log {
	var cat string
	switch t := category.(type) {
	case string:
		cat = t
	case interface{ String() string }:
		cat = t.String()
	default:
		cat = "default"
	}

	log := *l
	log.Logger = l.Logger.Cat(cat)
	return &log
}
