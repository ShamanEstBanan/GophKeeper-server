package service

import (
	"context"
	"fmt"

	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"ShamanEstBanan-GophKeeper-server/internal/errs"
)

func (s *service) GetAllRecords(ctx context.Context, userID entity.UserID) (*[]entity.RecordInfo, error) {
	return s.storage.GetAllRecords(ctx, userID)
}

func (s *service) GetRecordsByType(ctx context.Context, userID entity.UserID, datatype entity.DataType) (*[]entity.RecordInfo, error) {
	return s.storage.GetRecordsByType(ctx, userID, datatype)
}

func (s *service) CreateRecord(ctx context.Context, record entity.Record) (*entity.Record, error) {
	err := validateRecord(record)
	if err != nil {
		return nil, err
	}
	createdRecord, err := s.storage.CreateRecord(ctx, record)
	if err != nil {
		s.lg.Error("creation record error:")
		return nil, err
	}
	return createdRecord, nil
}

func (s *service) GetRecord(ctx context.Context, recID entity.RecordID, userID entity.UserID) (*entity.Record, error) {
	record, err := s.storage.GetRecord(ctx, recID, userID)
	if err != nil {
		return nil, err
	}
	if record.Name == "" {
		return nil, errs.ErrNotFound
	}
	return record, err
}

func (s *service) UpdateRecord(ctx context.Context, rec entity.Record) (*entity.Record, error) {
	if rec.Id == "" {
		return nil, fmt.Errorf("%w:%s", errs.ErrInvalidRecordInfo, errs.ErrInvalidRecordInfo)
	}
	err := validateRecord(rec)
	if err != nil {
		return nil, err
	}
	return s.storage.UpdateRecord(ctx, rec)
}

func (s *service) DeleteRecord(ctx context.Context, recID entity.RecordID, userID entity.UserID) error {
	return s.storage.DeleteRecord(ctx, recID, userID)
}

func validateRecord(record entity.Record) error {
	if record.Name == "" {
		return fmt.Errorf("%w:%s", errs.ErrInvalidRecordInfo, errs.ErrEmptyNameInRecord)
	}
	if record.Type == "" {
		return fmt.Errorf("%w:%s", errs.ErrInvalidRecordInfo, errs.ErrEmptyTypeInRecord)
	}
	if record.Data == nil {
		return fmt.Errorf("%w:%s", errs.ErrInvalidRecordInfo, errs.ErrEmptyDataInRecord)
	}
	return nil
}
