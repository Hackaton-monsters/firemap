package repository

//
//import (
//	"blogApi/internal/domain/contract"
//	"blogApi/internal/domain/entity"
//
//	"github.com/sirupsen/logrus"
//	"gorm.io/gorm"
//)
//
//type languageRepository struct {
//	db     *gorm.DB
//	logger *logrus.Logger
//}
//
//func NewLanguageRepository(
//	db *gorm.DB,
//	logger *logrus.Logger,
//) contract.LanguageRepository {
//	return &languageRepository{
//		db:     db,
//		logger: logger,
//	}
//}
//
//func (l languageRepository) GetByCode(code string) (*entity.Language, error) {
//	var language entity.Language
//	err := l.db.
//		Where("code = ?", code).
//		Find(&language).
//		Error
//	return &language, err
//}
