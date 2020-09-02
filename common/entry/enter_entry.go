package entry

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type enterEntry struct {
	widget.Entry
	callback func(value string)
	keyName fyne.KeyName
}

func NewEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	if key.Name == e.keyName {
		e.callback(e.Entry.Text)
	}
	/*switch key.Name {
	case e.keyName:
		e.onEnter()
	default:
		e.Entry.KeyDown(key)
	}*/
}

func (e *enterEntry) SetCallback(callback func (value string)){
	e.callback = callback
}

func (e *enterEntry) SetKeyName(keyName fyne.KeyName){
	e.keyName = keyName
}