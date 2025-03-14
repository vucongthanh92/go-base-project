package httpcommon

type listDataResponse struct {
	Data        interface{} `json:"data"`
	HasMoreData bool        `json:"hasMoreData"`
	Total       int64       `json:"total"`
}

func NewListResponse(data interface{}, hasMoreData bool, total int64) listDataResponse {
	return listDataResponse{
		Data:        data,
		HasMoreData: hasMoreData,
		Total:       total,
	}
}
