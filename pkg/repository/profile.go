package repository

import (
	"AGZ/pkg/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

type ProfilePostgres struct {
	db *sqlx.DB
}

func (r *ProfilePostgres) AddPurchase(user structures.Params) error {
	query := fmt.Sprintf("call custom.add_basket('%s', '%s')", user.Purchase.Purchase, user.Purchase.Access)
	fmt.Println(query)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *ProfilePostgres) RemovePurchase(user structures.Params) error {
	query := fmt.Sprintf("call custom.remove_basket('%s', '%s')", user.Purchase.Purchase, user.Purchase.Access)
	fmt.Println(query)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *ProfilePostgres) GetBasket(user structures.Params) ([]structures.Purchases, error) {
	var purchase structures.Purchases
	var purchases []structures.Purchases
	var res []structures.Purchases
	var id string
	QueryId := fmt.Sprintf("select auth.t_auth_session.user_id from auth.t_auth_session where auth.t_auth_session.access_token::text = '%s';", user.Purchase.Access)
	err := r.db.Get(&id, QueryId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	QueryPurchases := fmt.Sprintf("select * from purchase.t_entity inner join custom.t_basket on custom.t_basket.entity_id::text=t_entity.entity_id::text where t_basket.profile_id::text='%s';", id)

	fmt.Println(QueryPurchases)
	rows, err := r.db.Queryx(QueryPurchases)
	for rows.Next() {
		err = rows.StructScan(&purchase)
		fmt.Println(err)
		purchases = append(purchases, purchase)
	}
	for _, purchase := range purchases {
		var doc structures.Documents
		var lot structures.LotItems
		var Drug structures.DrugLotItems

		var docs []structures.Documents
		var lots []structures.LotItems
		var Drugs []structures.DrugLotItems

		QueryDocs := fmt.Sprintf("select * from purchase.t_document inner join purchase.t_document_entity on purchase.t_document_entity.document_id::text = purchase.t_document.document_id::text where purchase.t_document_entity.entity_id::text = '%s';", purchase.Uid)
		fmt.Println(QueryDocs)
		rows, err := r.db.Queryx(QueryDocs)
		for rows.Next() {
			err = rows.StructScan(&doc)
			fmt.Println(err)
			docs = append(docs, doc)
		}
		QueryLots := fmt.Sprintf("select * from purchase.t_lot_item inner join purchase.t_lot_entity on t_lot_item.sid::text = purchase.t_lot_entity.lot_id::text where purchase.t_lot_entity.entity_id::text = '%s';", purchase.Uid)
		rows, err = r.db.Queryx(QueryLots)
		for rows.Next() {
			err = rows.StructScan(&lot)
			fmt.Println(err)
			lots = append(lots, lot)
		}
		QueryDrug := fmt.Sprintf("select * from purchase.t_drug_lot_item inner join purchase.t_drug_entity on purchase.t_drug_lot_item.sid::text = purchase.t_drug_entity.drug_id::text\nwhere purchase.t_drug_entity.entity_id::text = '%s';", purchase.Uid)
		fmt.Println(QueryDrug)
		rows, err = r.db.Queryx(QueryDrug)
		for rows.Next() {
			err = rows.StructScan(&Drug)
			fmt.Println(err)
			Drugs = append(Drugs, Drug)
		}
		purchase.LotItems = lots
		purchase.Documents = docs
		purchase.DrugLotItems = Drugs
		purchase.Customer.CustomerPhone = purchase.CustomerPhone
		purchase.Customer.CustomerEmail = purchase.CustomerEmail
		purchase.Customer.CustomerRegNum = purchase.CustomerRegNum
		purchase.Customer.CustomerLocation = purchase.CustomerLocation
		purchase.Customer.CustomerOrgShortName = purchase.CustomerOrgShortName
		purchase.Customer.CustomerOrgFullName = purchase.CustomerOrgFullName
		purchase.Customer.CustomerContact.FirstName = purchase.FirstName
		purchase.Customer.CustomerContact.MiddleName = purchase.MiddleName
		purchase.Customer.CustomerContact.Lastname = purchase.Lastname

		purchase.Delivery.DeliveryPlace = purchase.DeliveryPlace
		purchase.Delivery.KladrCode = purchase.KladrCode
		purchase.Delivery.Kladr = purchase.Kladr
		purchase.Delivery.FullName = purchase.FullName
		res = append(res, purchase)
	}
	return res, err
}

func (r *ProfilePostgres) AddLink(user structures.Params) error {
	query := fmt.Sprintf("call custom.add_link('%s', '%s', '%s')", user.Link.Link, user.Link.Access, user.Link.Name)
	fmt.Println(query)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *ProfilePostgres) RemoveLink(user structures.Params) error {
	query := fmt.Sprintf("call custom.remove_link('%s', '%s')", user.Link.Name, user.Link.Access)
	fmt.Println(query)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *ProfilePostgres) GetLinksBasket(user structures.Params) ([]structures.Links, error) {
	query := fmt.Sprintf("select * from custom.get_link_basket('%s')", user.Link.Access)
	fmt.Println(query)
	var links []structures.Links
	var link structures.Links
	rows, err := r.db.Queryx(query)
	if err != nil {
		return []structures.Links{}, err
	}
	for rows.Next() {
		fmt.Println(rows.Rows)
		err = rows.StructScan(&link)
		fmt.Println(err)
		links = append(links, link)
	}

	return links, err
}

func (r *ProfilePostgres) AddCustomNote(user structures.Params) error {
	query := fmt.Sprintf("select t.profile_id from custom.t_profile as t where t.user_id = (select t2.user_id from auth.t_auth_session as t2 where t2.access_token::text = '%s');", user.Purchase.Access)
	var row string
	err := r.db.Get(&row, query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("update custom.t_basket set custom_note = '%s' where entity_id = '%s' and profile_id = '%s';", user.Purchase.Note, user.Purchase.Purchase, row)
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return err
}

func (r *ProfilePostgres) AddCustomStatus(user structures.Params) error {
	query := fmt.Sprintf("select t.profile_id from custom.t_profile as t where t.user_id = (select t2.user_id from auth.t_auth_session as t2 where t2.access_token::text = '%s');", user.Purchase.Access)
	var row string
	err := r.db.Get(&row, query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("update custom.t_basket set custom_note = '%s' where entity_id = '%s' and profile_id = '%s';", user.Purchase.Note, user.Purchase.Purchase, row)
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return err
}

// GetProfileInfo
func (r *ProfilePostgres) GetProfileInfo(token structures.Tokens) (structures.UserPersonalInfo, error) {
	query := fmt.Sprintf("select t.profile_id from custom.t_profile as t where t.user_id = (select t2.user_id from auth.t_auth_session as t2 where t2.access_token::text = '%s');", token.Access)
	var row string
	err := r.db.Get(&row, query)
	if err != nil {
		return structures.UserPersonalInfo{}, err
	}
	perosnalQuery := fmt.Sprintf("select first_name, second_name, last_name, email, phone_number, organization, inn from custom.t_profile where profile_id = '%s'", row)
	rows, err := r.db.Queryx(perosnalQuery)
	if err != nil {
		return structures.UserPersonalInfo{}, err
	}
	var userData structures.UserPersonalInfo
	for rows.Next() {
		fmt.Println(rows.Rows)
		err = rows.StructScan(&userData)
	}
	return userData, err

}
func (r *ProfilePostgres) AddProfileInfo(profile structures.UserPersonalInfo) error {
	query := fmt.Sprintf("select t.profile_id from custom.t_profile as t where t.user_id = (select t2.user_id from auth.t_auth_session as t2 where t2.access_token::text = '%s');", profile.Access)
	var row string
	err := r.db.Get(&row, query)
	if err != nil {
		return err
	}
	fmt.Println(row)
	query = fmt.Sprintf("UPDATE custom.t_profile SET first_name = COALESCE('%s', first_name), last_name = COALESCE('%s', last_name),  second_name = COALESCE('%s', second_name),  organization = COALESCE('%s', organization),  email = COALESCE('%s', email),  phone_number = COALESCE('%s', phone_number),  inn = COALESCE('%s', inn)where profile_id::text = '%s'", profile.First, profile.Last, profile.Second, profile.Organization, profile.Email, profile.Phone, profile.INN, row)
	fmt.Println(query)
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return err
}
