package pub

import "fmt"

type Err struct {
	Code uint
	Msg  string
}

func (e *Err) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Msg)
}

func (e *Err) Format(args ...interface{}) *Err {
	return &Err{
		Code: e.Code,
		Msg:  fmt.Sprintf(e.Msg, args...),
	}
}

func errorf(format string, args ...interface{}) {
	CheckErr(fmt.Errorf("gob: "+format, args...))
}

func CheckErr(err error) {
	if err != nil {
		panic(&Err{Msg: err.Error()})
	}
}

func CatchError(err *error) {
	if e := recover(); e != nil {
		//fmt.Println("test")
		ge, ok := e.(*Err)
		if !ok {
			panic(e)
		}
		*err = ge
	}
	return
}
