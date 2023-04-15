package storage

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"context"
	"fmt"
	"time"
)

func (s *storage) GetAllRecords(ctx context.Context, userID entity.UserID) (*[]entity.RecordInfo, error) {

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `SELECT id, name, datatype FROM records WHERE user_id = $1`
	rows, err := s.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	var records []entity.RecordInfo
	rec := entity.RecordInfo{}
	for rows.Next() {
		err = rows.Scan(
			&rec.Id,
			&rec.Name,
			&rec.Type)
		records = append(records, rec)
	}
	return &records, nil
}

func (s *storage) GetRecordsByType(ctx context.Context, userID entity.UserID, datatype entity.DataType) (*[]entity.RecordInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `SELECT id, name, datatype FROM records WHERE user_id = $1 AND datatype = $2`
	rows, err := s.db.Query(ctx, query, userID, datatype)
	if err != nil {
		return nil, err
	}
	var records []entity.RecordInfo
	rec := entity.RecordInfo{}
	for rows.Next() {
		err = rows.Scan(
			&rec.Id,
			&rec.Name,
			&rec.Type)
		records = append(records, rec)
	}

	return &records, nil
}

func (s *storage) GetRecord(ctx context.Context, recordID entity.RecordID, userID entity.UserID) (*entity.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `SELECT name, datatype, data FROM records WHERE user_id = $1 AND id = $2`
	rows := s.db.QueryRow(ctx, query, userID, recordID)
	rec := &entity.Record{
		Id:     recordID,
		UserID: userID,
	}
	err := rows.Scan(&rec.Name, &rec.Type, &rec.Data)
	if err != nil {
		return nil, err
	}

	return rec, nil
}

func (s *storage) CreateRecord(ctx context.Context, record entity.Record) (*entity.Record, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `INSERT INTO records (user_id, datatype, data, created_at, updated_at) VALUES ($1, $2, $3, $4,$5) RETURNING records.id`

	var recordID string
	err := s.db.QueryRow(ctx, query, record.UserID, record.Type, record.Data, time.Now(), time.Now()).Scan(&recordID)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &entity.Record{
		Id:     recordID,
		Name:   record.Name,
		Type:   record.Type,
		Data:   record.Data,
		UserID: record.UserID,
	}, nil
}

func (s *storage) UpdateRecord(ctx context.Context, record entity.Record) (*entity.Record, error) {

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `UPDATE records SET name = $1, datatype = $2, data = $3 WHERE id = $4`
	_, err := s.db.Exec(ctx, query, record.Name, record.Type, record.Data, record.Id)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (s *storage) DeleteRecord(ctx context.Context, recordID string, userID entity.UserID) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `DELETE FROM records WHERE id = $1 AND user_id = $2`
	_, err := s.db.Exec(ctx, query, recordID, userID)
	if err != nil {
		return err
	}
	return nil
}
