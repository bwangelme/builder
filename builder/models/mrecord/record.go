package mrecord

import (
	"builder/builder/libs/store"
	"builder/builder/pbs/recordpb"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Add(giturl string, revision string, image string, tag string) (*recordpb.Record, error) {
	sql := "insert into record(status, giturl, revision, img, tag) values (?,?,?,?)"
	ret, err := store.MySQL.Exec(sql, recordpb.StatusType_Init, giturl, revision, image, tag)
	if err != nil {
		return nil, errors.Wrapf(err, "insert failed")
	}

	theID, err := ret.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(err, "get LastInsertId failed")
	}

	return Get(uint64(theID))
}

func Get(id uint64) (*recordpb.Record, error) {
	records, err := GetMulti([]uint64{id})
	if err != nil {
		return nil, err
	}

	if len(records) != 1 {
		return nil, errors.New("get failed")
	}

	return records[0], nil
}

func GetMulti(ids []uint64) ([]*recordpb.Record, error) {
	if len(ids) > 500 {
		return nil, errors.New("get too much once")
	}

	var records []recordpb.Record
	query, args, err := sqlx.In("select * from record where id in (?)", ids)
	if err != nil {
		return nil, errors.Wrapf(err, "build sql err")
	}
	query = store.MySQL.Rebind(query)
	err = store.MySQL.Select(&records, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "select error |%s|", query)
	}

	var res []*recordpb.Record
	for _, r := range records {
		res = append(res, &r)
	}

	return res, nil
}
