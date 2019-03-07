package timeformatter

import (
	"syscall/js"
	"time"
)

// Counter is basic counter
type TimeFormatter struct {
	Time            time.Time
	Layout          string
	FormattedStr    string
	formattedTimeEl js.Value
	layoutFieldEl   js.Value
	timeFieldEl     js.Value
}

func NewTimeFormatter(elementID string) *TimeFormatter {
	tf := &TimeFormatter{}
	tf.initDOM(elementID)
	tf.Now()
	return tf
}

func (tf *TimeFormatter) initDOM(elementID string) {
	doc := js.Global().Get("document")
	el := doc.Call("getElementById", elementID)
	el.Call("insertAdjacentHTML", "beforeend", `
<div id="timeFormatter">
	<div>
		<label>Time: </label>
		<input id="timeField" placeholder="Time" value="2006-01-02T15:04:05Z07:00"></input>
	</div>
	<div>
		<label>Layout: </label>
		<input id="layoutField" placeholder="Layout" value="Mon Jan _2 15:04:05 MST 2006"></input>
	</div>
	<div id="formattedTime">
	</div>
	<button id="now">Now</button>
	<button id="format">Format</button>
</div>
`)
	tf.formattedTimeEl = doc.Call("getElementById", "formattedTime")
	tf.layoutFieldEl = doc.Call("getElementById", "layoutField")
	tf.timeFieldEl = doc.Call("getElementById", "timeField")
	now := doc.Call("getElementById", "now")
	format := doc.Call("getElementById", "format")
	now.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} { tf.Now(); return nil }))
	format.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} { tf.Format(); return nil }))
	tf.update()
}

func (tf *TimeFormatter) Now() {
	tf.Time = time.Now()
	layout := tf.layoutFieldEl.Get("value").String()
	tf.timeFieldEl.Set("value", tf.Time.Format(time.RFC3339))
	tf.FormattedStr = tf.Time.Format(layout)
	tf.update()
}

func (tf *TimeFormatter) Format() {
	layout := tf.layoutFieldEl.Get("value").String()
	timeStr := tf.timeFieldEl.Get("value").String()
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		tf.FormattedStr = `error parsing time str: Please input format of RFC3339("2006-01-02T15:04:05Z07:00")`
	}
	tf.Time = t
	tf.FormattedStr = t.Format(layout)
	tf.update()
}

func (tf *TimeFormatter) update() {
	tf.formattedTimeEl.Set("textContent", tf.FormattedStr)
}
