package api

type (
	Base struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}

	BaseWithMeta struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta"`
	}

	Error struct {
		Status string `json:"status"`
		Data   struct {
			Error string `json:"error"`
		} `json:"data"`
	}

	Message struct {
		Message string `json:"message"`
	}

	List struct {
		Data interface{} `json:"data"`
	}
)
