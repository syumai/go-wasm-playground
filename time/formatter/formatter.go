package formatter

import (
	"errors"
	"syscall/js"
	"time"

	"github.com/syumai/go-wasm-playground/dom/button"
	"github.com/syumai/go-wasm-playground/dom/div"
	"github.com/syumai/go-wasm-playground/dom/input"
)

const defaultLayout = time.RFC3339Nano

// Formatter is time formatter
type Formatter struct {
	Time            time.Time
	Layout          string
	FormattedStr    string
	resultContainer *div.Div
	layoutField     *input.TextField
	timeField       *input.TextField
}

func New(elementID string) *Formatter {
	tf := &Formatter{}
	tf.initDOM(elementID)
	return tf
}

func (tf *Formatter) initDOM(elementID string) {
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
	<div>
		<button id="format">Format</button>
	</div>
	<div>
		<button id="now">Set current time</button>
	</div>
	<div>
		<label>Use layout constants</label>
		<div>
			<button id="ansic">ANSIC</button>
			<button id="unixDate">UnixDate</button>
			<button id="rubyDate">RubyDate</button>
			<button id="rfc822">RFC822</button>
			<button id="rfc822Z">RFC822Z</button>
			<button id="rfc850">RFC850</button>
			<button id="rfc1123">RFC1123</button>
			<button id="rfc1123Z">RFC1123Z</button>
			<button id="rfc3339">RFC3339</button>
			<button id="rfc3339Nano">RFC3339Nano</button>
			<button id="kitchen">Kitchen</button>
			<button id="stamp">Stamp</button>
			<button id="stampMilli">StampMilli</button>
			<button id="stampMicro">StampMicro</button>
			<button id="stampNano">StampNano</button>
		</div>
	</div>
</div>
`)
	tf.layoutField = input.NewTextField("layoutField")
	tf.timeField = input.NewTextField("timeField")
	tf.resultContainer = div.New("formattedTime")
	now := button.New("now")
	now.OnClick(tf.Now)
	format := button.New("format")
	format.OnClick(tf.Format)
	layoutButtonConfigs := []struct {
		id     string
		layout string
	}{
		{"ansic", time.ANSIC},
		{"unixDate", time.UnixDate},
		{"rubyDate", time.RubyDate},
		{"rfc822", time.RFC822},
		{"rfc822Z", time.RFC822Z},
		{"rfc850", time.RFC850},
		{"rfc1123", time.RFC1123},
		{"rfc1123Z", time.RFC1123Z},
		{"rfc3339", time.RFC3339},
		{"rfc3339Nano", time.RFC3339Nano},
		{"kitchen", time.Kitchen},
		{"stamp", time.Stamp},
		{"stampMilli", time.StampMilli},
		{"stampMicro", time.StampMicro},
		{"stampNano", time.StampNano},
	}
	for _, cnf := range layoutButtonConfigs {
		cnf := cnf
		b := button.New(cnf.id)
		b.OnClick(func() {
			tf.layoutField.SetValue(cnf.layout)
			tf.Format()
		})
	}
}

func (tf *Formatter) updateTime() error {
	timeStr := tf.timeField.Value()
	t, err := time.Parse(defaultLayout, timeStr)
	if err != nil {
		return errors.New("error parsing time str: please input date format in `" + defaultLayout + "`")
	}
	tf.Time = t
	return nil
}

func (tf *Formatter) updateLayout() {
	tf.Layout = tf.layoutField.Value()
}

func (tf *Formatter) Now() {
	tf.timeField.SetValue(time.Now().Format(defaultLayout))
	tf.Format()
}

func (tf *Formatter) Format() {
	tf.updateLayout()
	err := tf.updateTime()
	if err != nil {
		tf.resultContainer.SetText(err.Error())
		return
	}
	tf.FormattedStr = tf.Time.Format(tf.Layout)
	tf.resultContainer.SetText(tf.FormattedStr)
}
