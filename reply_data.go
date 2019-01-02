package echo

type (
	// 业务错误返回
	Err struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	// 集合分页返回
	Pagination struct {
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}
	Meta struct {
		Total     int `json:"total"`    //数据总数
		PageIndex int `json:"current"`  //当前页码
		PageSize  int `json:"pageSize"` //每页数据数量
		//PageTotal       int `json:"page_total"`        //页码总数 ? ng框架不需要
		//CurrentPageSize int `json:"current_page_size"` //当前页实际数据数量 ? ng框架不需要
	}
)

func (m *Err) Error() string {
	return m.Msg
}

const (
	ZH = "zh"
	EN = "en"

	DEFALT_LANG = ZH // 默认中文

	DEFAULT_CODE       = 1000 // 默认错误
	SYSTEM_ERROR       = 1001 // 系统错误
	PARAMS_BIND_ERROR  = 1002 // 参数绑定错误
	PARAMS_CHICK_ERROR = 1003 // 参数校验错误
)

type BilingualMsgGenerate struct {
	Zh string //中文
	En string //英文
}

type BilingualMsgConsume struct {
	ZhErr *Err // 中文错误
	EnErr *Err // 英文错误
}

func (b *BilingualMsgGenerate) Generate(code int) BusinessErr {
	consume := new(BilingualMsgConsume)
	consume.ZhErr = &Err{Code: code, Msg: b.Zh}
	consume.EnErr = &Err{Code: code, Msg: b.En}
	return consume
}

func (b *BilingualMsgConsume) ErrStruct(language string) *Err {
	if language == EN {
		return b.EnErr
	}
	return b.ZhErr
}

func (BilingualMsgConsume) Error() string {
	return "兼容Error interface"
}

var BASE_ERROR_CODE = map[int]BilingualMsgGenerate{
	DEFAULT_CODE:       {"默认错误", "Default error"},
	SYSTEM_ERROR:       {"系统错误", "system error"},
	PARAMS_BIND_ERROR:  {"参数绑定错误", "Parameter binding error"},
	PARAMS_CHICK_ERROR: {"参数校验错误", "Parameter check error"},
}

func getLang(c Context) string {
	lang, _ := c.Get(TRANSLATE_LANGUAGE_HEADER_KEY).(string)
	if lang == "" {
		lang = DEFALT_LANG
	}
	return lang
}
