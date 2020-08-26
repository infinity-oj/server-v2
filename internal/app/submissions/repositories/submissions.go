package repositories

import (
	"container/list"
	"sync"

	"github.com/google/uuid"

	"github.com/infinity-oj/server-v2/internal/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Repository interface {
	GetSubmission(submissionId string) (*models.Submission, error)
	GetSubmissionById(id uint64) (*models.Submission, error)
	GetSubmissions(offset, limit int, problemId string) ([]*models.Submission, error)
	Create(submitterID, problemId uint64, userSpace string) (s *models.Submission, err error)
	Update(s *models.Submission) error
}

type DefaultRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	queue  *list.List
	mutex  *sync.Mutex
}

func (m DefaultRepository) GetSubmissionById(id uint64) (*models.Submission, error) {
	submission := &models.Submission{}
	err := m.db.First(&submission, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (m DefaultRepository) GetSubmissions(offset, limit int, problemId string) (res []*models.Submission, err error) {
	if err = m.db.Table("submissions").Where("problem_id = ?", problemId).
		Offset(offset).Limit(limit).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (m DefaultRepository) GetSubmission(submissionId string) (*models.Submission, error) {
	submission := &models.Submission{}
	err := m.db.Where("submission_id = ?", submissionId).First(&submission).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (m DefaultRepository) Create(submitterId, problemId uint64, userSpace string) (s *models.Submission, err error) {
	s = &models.Submission{
		SubmissionId: uuid.New().String(),
		SubmitterId:  submitterId,
		ProblemId:    problemId,
		UserSpace:    userSpace,
	}

	if err = m.db.Create(s).Error; err != nil {
		return nil, errors.Wrapf(err,
			" create submission with username: %d, problemID: %s, userSpace: %s",
			submitterId, problemId, userSpace,
		)
	}

	return
}

func (m DefaultRepository) Update(s *models.Submission) (err error) {
	err = m.db.Save(s).Error
	return
}

func NewMysqlSubmissionsRepository(logger *zap.Logger, db *gorm.DB) Repository {
	return &DefaultRepository{
		logger: logger.With(zap.String("type", "Repository")),
		db:     db,
		queue:  list.New(),
		mutex:  &sync.Mutex{},
	}
}
