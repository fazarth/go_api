package entity

type Greige struct {
	GreigeID         uint64 `gorm:"primary_key:auto_increment" json:"greige_id"`
	LocationID       uint64 `gorm:"type:int" json:"location_id"`
	GreigeNum        string `gorm:"type:varchar(255)" json:"greige_num"`
	GreigeBarcode    string `gorm:"type:varchar(255)" json:"greige_barcode"`
	CardCode         string `gorm:"type:varchar(255)" json:"card_code"`
	CardName         string `gorm:"type:text" json:"card_name"`
	ItemCode         string `gorm:"type:varchar(255)" json:"item_code"`
	ItemName         string `gorm:"type:varchar(255)" json:"item_name"`
	ItemColour       string `gorm:"type:text" json:"item_colour"`
	ItemProcess      string `gorm:"type:varchar(255)" json:"item_process"`
	IsMakloon        string `gorm:"type:varchar(255)" json:"is_makloon"`
	SoNum            uint64 `gorm:"type:int" json:"so_num"`
	SoLineNum        uint64 `gorm:"type:int" json:"so_line_num"`
	SoRefNum         string `gorm:"type:varchar(255)" json:"so_ref_num"`
	SoExpireTime     string `gorm:"type:varchar(255)" json:"so_expire_time"`
	PdoNum           uint64 `gorm:"type:int" json:"pdo_num"`
	PdoQty           uint64 `gorm:"type:int" json:"pdo_qty"`
	PdoWeight        uint64 `gorm:"type:int" json:"pdo_weight"`
	PdoTime          string `gorm:"type:varchar(255)" json:"pdo_time"`
	PdoShift         uint64 `gorm:"type:int" json:"pdo_shift"`
	PdoMachine       string `gorm:"type:varchar(255)" json:"pdo_machine"`
	PdoOperator      string `gorm:"type:varchar(255)" json:"pdo_operator"`
	QcTime           string `gorm:"type:varchar(255)" json:"qc_time"`
	QcOperator       string `gorm:"type:varchar(255)" json:"qc_operator"`
	ScaleTime        string `gorm:"type:varchar(255)" json:"scale_time"`
	ScaleOperator    string `gorm:"type:varchar(255)" json:"scale_operator"`
	WeightTypeCode   string `gorm:"type:varchar(255)" json:"weight_type_code"`
	WeightTypeName   string `gorm:"type:varchar(255)" json:"weight_type_name"`
	Bruto            uint64 `gorm:"type:int" json:"bruto"`
	BrutoRound       uint64 `gorm:"type:int" json:"bruto_round"`
	Netto            uint64 `gorm:"type:int" json:"netto"`
	NettoRound       uint64 `gorm:"type:int" json:"netto_round"`
	Weight           uint64 `gorm:"type:int" json:"weight"`
	WeightRemain     uint64 `gorm:"type:int" json:"weight_remain"`
	Length           uint64 `gorm:"type:int" json:"lenght"`
	LengthRemain     uint64 `gorm:"type:int" json:"lenght_remain"`
	GreigeStatusCode string `gorm:"type:varchar(255)" json:"greige_status_code"`
	GreigeStatusName string `gorm:"type:varchar(255)" json:"greige_status_name"`
	GreigeStatusTime string `gorm:"type:varchar(255)" json:"greige_status_time"`
	ArticleNum       string `gorm:"type:varchar(255)" json:"article_num"`
	NumAtCard        string `gorm:"type:varchar(255)" json:"num_at_card"`
	DocNumIso        string `gorm:"type:varchar(255)" json:"doc_num_iso"`
	Comments         string `gorm:"type:varchar(255)" json:"comments"`
	IsUsable         string `gorm:"type:varchar(255)" json:"is_usable"`
	Create_Id        uint64 `gorm:"type:int" json:"create_id"`
	Create_Time      string `gorm:"type:varchar(255)" json:"create_time"`
	Update_Id        uint64 `gorm:"type:int" json:"update_id"`
	Update_Time      string `gorm:"type:varchar(255)" json:"update_time"`
	LabdippCode      string `gorm:"type:varchar(255)" json:"labdipp_code"`
	CardAddress      string `gorm:"type:varchar(255)" json:"card_address"`
	SoNumBook        uint64 `gorm:"type:int" json:"so_num_book"`
	SoLineNumBook    uint64 `gorm:"type:int" json:"so_line_num_book"`
	SoRefNumBook     string `gorm:"type:varchar(255)" json:"so_ref_num_book"`
	SoExpireTimeBook string `gorm:"type:varchar(255)" json:"so_expire_time_book"`
	CardCodeBook     string `gorm:"type:varchar(255)" json:"card_code_book"`
	CardNameBook     string `gorm:"type:varchar(255)" json:"card_name_book"`
	NumAtCardBook    string `gorm:"type:varchar(255)" json:"num_at_book"`
	ItemColourBook   string `gorm:"type:varchar(255)" json:"item_colour_book"`
	ItemProcessBook  string `gorm:"type:varchar(255)" json:"time_process_book"`
	LabdippCodeBook  string `gorm:"type:varchar(255)" json:"labdip_code_book"`
	CardAddressBook  string `gorm:"type:varchar(255)" json:"card_address_book"`
}
