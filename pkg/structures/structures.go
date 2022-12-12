package structures

import (
	"context"
	"net/http"
	"time"
)

type UserPersonalInfo struct {
	Access       string `json:"access"`
	First        string `json:"first_name" db:"first_name"`
	Second       string `json:"second_name" db:"second_name"`
	Last         string `json:"last_name" db:"last_name"`
	Email        string `json:"email" db:"email"`
	Phone        string `json:"phone_number" db:"phone_number"`
	Organization string `json:"organization" db:"organization"`
	INN          string `json:"INN" db:"inn"`
}
type User struct {
	Login    string
	Email    string
	Password string
	DevInfo  string
	Hash     string
}
type Params struct {
	Basket   basket
	User     user
	EmailReg emailReg
	Purchase purchase
	Link     link
}

type user struct {
	Login   string `json:"login"`
	Pass    string `json:"pass"`
	DevInfo string `json:"devInfo"`
}

type basket struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	DevInfo string `json:"devInfo"`
}

type emailReg struct {
	Login string `json:"login"`
	Hash  string `json:"hash"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type purchase struct {
	Access   string `json:"access"`
	Purchase string `json:"purchase"`
	Note     string `json:"custom_note"`
	Status   string `json:"custom_status"`
}
type link struct {
	Access string `json:"access"`
	Link   string `json:"link" db:"link"`
	Name   string `json:"name"`
}
type Links struct {
	Links []string `json:"links"`
	Link  string   `json:"link" db:"link"`
	Name  string   `json:"name" db:"name"`
}

type Tokens struct {
	Access  string `json:"access" db:"c_access_token"`
	Refresh string `json:"refresh" db:"c_refresh_token"`
	DevInfo string `json:"devInfo"`
	AccessR string `json:"accessR" db:"a_access_token"`
}
type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

type Purchases struct {
	Status           string         `json:"status" db:"status"`
	Uid              string         `json:"Uid" db:"entity_id"`
	PurchaseNumber   string         `json:"PurchaseNumber" db:"purchase_number"`
	Description      string         `json:"Description" db:"description"`
	DateStart        string         `json:"DateStart" db:"date_start"`
	DateEnd          string         `json:"DateEnd" db:"date_end"`
	DateStartAuction string         `json:"DateStartAuction" db:"date_start_auction"`
	DateEndAuction   string         `json:"DateEndAuction" db:"date_end_auction"`
	OriginalUrl      string         `json:"OriginalUrl" db:"original_url"`
	NameEtp          string         `json:"NameEtp" db:"name_etp"`
	UrlNameEtp       string         `json:"UrlNameEtp" db:"url_name_etp"`
	Fz               string         `json:"Fz" db:"fz"`
	Type             string         `json:"Type" db:"type"`
	PurchaseStatus   string         `json:"PurchaseStatus" db:"purchase_status"`
	Customer         customer       `json:"Customer"`
	Delivery         deliveryPlace  `json:"Delivery"`
	DeliveryTime     string         `json:"DeliveryTime" db:"delivery_time"`
	Price            string         `json:"Price" db:"price"`
	Region           string         `json:"Region" db:"region"`
	Currency         string         `json:"Currency" db:"currency"`
	AmountCollateral string         `json:"AmountCollateral" db:"amount_collateral"`
	Documents        []Documents    `json:"Documents"`
	LotItems         []LotItems     `json:"LotItems"`
	DrugLotItems     []DrugLotItems `json:"DrugLotItems"`
	//db only fields
	KladrCode     string `db:"delivery_kladr_code"`
	FullName      string `db:"delivery_full_name"`
	Kladr         string `db:"delivery_kladr"`
	DeliveryPlace string `db:"delivery_place"`

	CustomerPhone        string `db:"customer_phone"`
	CustomerEmail        string `db:"customer_email"`
	CustomerLocation     string `db:"customer_location"`
	CustomerOrgFullName  string `db:"customer_org_full_name"`
	CustomerOrgShortName string `db:"customer_org_short_name"`
	CustomerRegNum       string `db:"customer_reg_num"`
	Lastname             string `db:"customer_contact_last_name"`
	FirstName            string `db:"customer_contact_first_name"`
	MiddleName           string `db:"customer_contact_middle_name"`
	ProfileId            string `db:"profile_id"`
	CustomNote           string `db:"custom_note"`
	CustomStatus         string `json:"custom_status" db:"custom_status"`
}
type deliveryPlace struct {
	KladrCode     string `json:"KladrCode"`
	FullName      string `json:"FullName"`
	Kladr         string `json:"Kladr"`
	DeliveryPlace string `json:"DeliveryPlace"`
}
type customer struct {
	CustomerPhone        string `json:"CustomerPhone" db:"customer_phone"`
	CustomerEmail        string `json:"CustomerEmail" db:"customer_email"`
	CustomerLocation     string `json:"CustomerLocation"`
	CustomerOrgFullName  string `json:"CustomerOrgFullName"`
	CustomerOrgShortName string `json:"CustomerOrgShortName"`
	CustomerRegNum       string `json:"customer_reg_num"`
	CustomerContact      customerContact
}
type customerContact struct {
	Lastname   string `json:"Lastname"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
}
type Documents struct {
	Did          string `db:"document_id"`
	DocumentName string `json:"DocumentName" db:"document_name"`
	DocumentUrl  string `json:"DocumentUrl" db:"document_url"`
	EntityId     string `json:"EntityId" db:"entity_id"`
}
type LotItems struct {
	Sid             string `json:"Sid" db:"sid"`
	Name            string `json:"Name" db:"name"`
	UnitPrice       string `json:"UnitPrice" db:"unit_price"`
	Sum             string `json:"Sum" db:"sum"`
	Quantity        string `json:"Quantity" db:"quantity"`
	Okpd2Code       string `json:"Okpd2Code" db:"okpd_2_code"`
	OkeiTitle       string `json:"OkeiTitle" db:"okei_title"`
	CodeMes         string `json:"CodeMes" db:"code_mes"`
	NationalCodeMes string `json:"NationalCodeMes" db:"national_code_mes"`
	NameMes         string `json:"NameMes" db:"name_mes"`
	LotId           string `json:"LotId" db:"lot_id"`
	EntityId        string `json:"EntityId" db:"entity_id"`
}
type DrugLotItems struct {
	Sid             string `json:"Sid" db:"sid"`
	MNNinfo         string `json:"MNNinfo" db:"MNNinfo"`
	MedFormName     string `json:"MedFormName" db:"med_form_name"`
	DosageGRLSValue string `json:"DosageGRLSValue" db:"dosage_grlsv_values"`
	DosageCode      string `json:"DosageCode" db:"dosage_code"`
	DosageOKEIName  string `json:"DosageOKEIName" db:"dosage_okei_name"`
	DosageUserName  string `json:"DosageUserName" db:"dosage_user_name"`
	DrugQuantity    string `json:"DrugQuantity" db:"drug_quantity"`
	EntityId        string `json:"EntityId" db:"entity_id"`
	LotId           string `json:"LotId" db:"lot_id"`
}
