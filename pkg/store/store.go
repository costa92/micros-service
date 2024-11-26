package store

import (
	"context"
	"errors"

	"github.com/costa92/micros-service/pkg/store/logger/empty"
	"github.com/costa92/micros-service/pkg/store/where"
	"gorm.io/gorm"
)

// DBProvider is a marker interface for database providers.
type DBProvider interface {
	DB(ctx context.Context) *gorm.DB
}

// Option defines a function type for configuring the Store.
type Option[T any] func(*Store[T])

// Store is a generic store.
type Store[T any] struct {
	logger Logger
	store  DBProvider
}

// Logger is a generic logger.
func WithLogger[T any](logger Logger) Option[T] {
	return func(s *Store[T]) {
		s.logger = logger
	}
}

// NewStore creates a new store.
func NewStore[T any](storage DBProvider, logger Logger) *Store[T] {
	if logger == nil {
		logger = empty.NewLogger()
	}

	return &Store[T]{logger: logger, store: storage}
}

func (s *Store[T]) db(ctx context.Context, wheres ...where.Where) *gorm.DB {
	dbInstance := s.store.DB(ctx)
	for _, whr := range wheres {
		if whr != nil {
			dbInstance = whr.Where(dbInstance)
		}
	}
	return dbInstance
}

// Create inserts a new record into the database.
func (s *Store[T]) Create(ctx context.Context, obj *T) error {
	if err := s.db(ctx).Create(obj).Error; err != nil {
		s.logger.Error(err, "failed to create record")
		return err
	}
	return nil
}

func (s *Store[T]) Update(ctx context.Context, obj *T, wheres ...where.Where) error {
	if err := s.db(ctx).Save(obj).Error; err != nil {
		s.logger.Error(err, "Failed to update object in database", "object", obj)
		return err
	}
	return nil
}

// Get retrieves a record from the database.
func (s *Store[T]) Delete(ctx context.Context, opts *where.WhereOptions) error {
	err := s.db(ctx, opts).Delete(new(T)).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error(err, "Failed to delete object from database", "conditions", opts)
		return err
	}
	return nil
}

func (s *Store[T]) Get(ctx context.Context, opts *where.WhereOptions) (*T, error) {
	var obj T
	if err := s.db(ctx, opts).First(&obj).Error; err != nil {
		s.logger.Error(err, "Failed to retrieve object from database", "conditions", opts)
		return nil, err
	}
	return &obj, nil
}

// List retrieves a list of records from the database.

func (s *Store[T]) List(ctx context.Context, opts *where.WhereOptions) (count int64, ret []*T, err error) {
	err = s.db(ctx, opts).Order("id desc").Find(&ret).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		s.logger.Error(err, "Failed to list objects from database", "conditions", opts)
	}
	return
}
