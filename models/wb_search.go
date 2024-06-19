package models

type WbSearch struct {
	Metadata struct {
		Name         string `json:"name"`
		CatalogType  string `json:"catalog_type"`
		CatalogValue string `json:"catalog_value"`
		Normquery    string `json:"normquery"`
		Rmi          string `json:"rmi"`
		Rs           int    `json:"rs"`
		Title        string `json:"title"`
	} `json:"metadata"`
	State          int `json:"state"`
	Version        int `json:"version"`
	PayloadVersion int `json:"payloadVersion"`
	Data           struct {
		Products []struct {
			Time1           int           `json:"time1"`
			Time2           int           `json:"time2"`
			Wh              int           `json:"wh"`
			Dtype           int           `json:"dtype"`
			Dist            int           `json:"dist"`
			Id              int           `json:"id"`
			Root            int           `json:"root"`
			KindId          int           `json:"kindId"`
			Brand           string        `json:"brand"`
			BrandId         int           `json:"brandId"`
			SiteBrandId     int           `json:"siteBrandId"`
			Colors          []interface{} `json:"colors"`
			SubjectId       int           `json:"subjectId"`
			SubjectParentId int           `json:"subjectParentId"`
			Name            string        `json:"name"`
			Supplier        string        `json:"supplier"`
			SupplierId      int           `json:"supplierId"`
			SupplierRating  float64       `json:"supplierRating"`
			SupplierFlags   int           `json:"supplierFlags"`
			Pics            int           `json:"pics"`
			Rating          int           `json:"rating"`
			ReviewRating    float64       `json:"reviewRating"`
			Feedbacks       int           `json:"feedbacks"`
			Volume          int           `json:"volume"`
			ViewFlags       int           `json:"viewFlags"`
			IsNew           bool          `json:"isNew,omitempty"`
			Sizes           []struct {
				Name     string `json:"name"`
				OrigName string `json:"origName"`
				Rank     int    `json:"rank"`
				OptionId int    `json:"optionId"`
				Wh       int    `json:"wh"`
				Dtype    int    `json:"dtype"`
				Price    struct {
					Basic     int `json:"basic"`
					Product   int `json:"product"`
					Total     int `json:"total"`
					Logistics int `json:"logistics"`
					Return    int `json:"return"`
				} `json:"price"`
				SaleConditions int    `json:"saleConditions"`
				Payload        string `json:"payload"`
			} `json:"sizes"`
			TotalQuantity int `json:"totalQuantity"`
			Log           struct {
				Cpm           int    `json:"cpm,omitempty"`
				Promotion     int    `json:"promotion,omitempty"`
				PromoPosition int    `json:"promoPosition,omitempty"`
				Position      int    `json:"position,omitempty"`
				AdvertId      int    `json:"advertId,omitempty"`
				Tp            string `json:"tp,omitempty"`
			} `json:"log"`
			PanelPromoId   int    `json:"panelPromoId,omitempty"`
			PromoTextCard  string `json:"promoTextCard,omitempty"`
			PromoTextCat   string `json:"promoTextCat,omitempty"`
			Logs           string `json:"logs,omitempty"`
			FeedbackPoints int    `json:"feedbackPoints,omitempty"`
		} `json:"products"`
		Total int `json:"total"`
	} `json:"data"`
}
