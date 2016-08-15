package service

import (
	"database/sql"
	"gosaic/database"
	"gosaic/model"

	"gopkg.in/gorp.v1"
)

var (
	gidx         model.Gidx
	gidxPartial  model.GidxPartial
	aspect       model.Aspect
	cover        model.Cover
	coverPartial model.CoverPartial
	macro        model.Macro
	macroPartial model.MacroPartial
)

func getTestDbMap() (*gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, err
	}

	_, err = database.Migrate(db)
	if err != nil {
		return nil, err
	}

	dbMap := gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	return &dbMap, nil
}

func getTestGidxService(dbMap *gorp.DbMap) (GidxService, error) {
	gidxService := NewGidxService(dbMap)
	err := gidxService.Register()
	if err != nil {
		return nil, err
	}

	return gidxService, nil
}

func getTestAspectService(dbMap *gorp.DbMap) (AspectService, error) {
	aspectService := NewAspectService(dbMap)
	err := aspectService.Register()
	if err != nil {
		return nil, err
	}

	return aspectService, nil
}

func getTestGidxPartialService(dbMap *gorp.DbMap) (GidxPartialService, error) {
	gidxPartialService := NewGidxPartialService(dbMap)
	err := gidxPartialService.Register()
	if err != nil {
		return nil, err
	}

	return gidxPartialService, nil
}

func getTestCoverService(dbMap *gorp.DbMap) (CoverService, error) {
	coverService := NewCoverService(dbMap)
	err := coverService.Register()
	if err != nil {
		return nil, err
	}

	return coverService, nil
}

func getTestCoverPartialService(dbMap *gorp.DbMap) (CoverPartialService, error) {
	coverPartialService := NewCoverPartialService(dbMap)
	err := coverPartialService.Register()
	if err != nil {
		return nil, err
	}

	return coverPartialService, nil
}

func getTestMacroService(dbMap *gorp.DbMap) (MacroService, error) {
	macroService := NewMacroService(dbMap)
	err := macroService.Register()
	if err != nil {
		return nil, err
	}

	return macroService, nil
}

func getTestMacroPartialService(dbMap *gorp.DbMap) (MacroPartialService, error) {
	macroPartialService := NewMacroPartialService(dbMap)
	err := macroPartialService.Register()
	if err != nil {
		return nil, err
	}

	return macroPartialService, nil
}

func getTestPartialComparisonService(dbMap *gorp.DbMap) (PartialComparisonService, error) {
	partialComparisonService := NewPartialComparisonService(dbMap)
	err := partialComparisonService.Register()
	if err != nil {
		return nil, err
	}

	return partialComparisonService, nil
}

func getTestMosaicService(dbMap *gorp.DbMap) (MosaicService, error) {
	mosaicService := NewMosaicService(dbMap)
	err := mosaicService.Register()
	if err != nil {
		return nil, err
	}

	return mosaicService, nil
}

func getTestMosaicPartialService(dbMap *gorp.DbMap) (MosaicPartialService, error) {
	mosaicPartialService := NewMosaicPartialService(dbMap)
	err := mosaicPartialService.Register()
	if err != nil {
		return nil, err
	}

	return mosaicPartialService, nil
}
