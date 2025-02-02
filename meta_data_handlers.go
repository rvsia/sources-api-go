package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RedHatInsights/sources-api-go/dao"
	m "github.com/RedHatInsights/sources-api-go/model"
	"github.com/RedHatInsights/sources-api-go/util"
	"github.com/labstack/echo/v4"
)

// function that defines how we get the dao - default implementation below.
var getMetaDataDao func(c echo.Context) (dao.MetaDataDao, error)

func getMetaDataDaoWithTenant(c echo.Context) (dao.MetaDataDao, error) {
	var tenantID int64
	var ok bool
	tenantVal := c.Get("tenantID")

	// if we set the tenant on this request - include it. otherwise do not.
	if tenantVal != nil {
		if tenantID, ok = tenantVal.(int64); !ok {
			return nil, fmt.Errorf("failed to pull tenant from request")
		}

		return &dao.MetaDataDaoImpl{TenantID: &tenantID}, nil
	} else {
		return &dao.MetaDataDaoImpl{}, nil
	}
}

func MetaDataList(c echo.Context) error {
	metaDataDB, err := getMetaDataDao(c)
	if err != nil {
		return err
	}

	filters, err := getFilters(c)
	if err != nil {
		return err
	}

	limit, offset, err := getLimitAndOffset(c)
	if err != nil {
		return err
	}

	var (
		metaDatas []m.MetaData
		count     int64
	)

	metaDatas, count, err = metaDataDB.List(limit, offset, filters)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorDoc(err.Error(), "400"))
	}

	out := make([]interface{}, len(metaDatas))
	for i := 0; i < len(metaDatas); i++ {
		out[i] = metaDatas[i].ToResponse()
	}

	return c.JSON(http.StatusOK, util.CollectionResponse(out, c.Request(), int(count), limit, offset))
}

func ApplicationTypeListMetaData(c echo.Context) error {
	metaDataDB, err := getMetaDataDao(c)
	if err != nil {
		return err
	}

	filters, err := getFilters(c)
	if err != nil {
		return err
	}

	limit, offset, err := getLimitAndOffset(c)
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(c.Param("application_type_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorDoc(err.Error(), "400"))
	}

	var (
		metaDatas []m.MetaData
		count     int64
	)

	metaDatas, count, err = metaDataDB.SubCollectionList(m.ApplicationType{Id: id}, limit, offset, filters)

	if err != nil {
		if errors.Is(err, util.ErrNotFoundEmpty) {
			return err
		}
		return c.JSON(http.StatusBadRequest, util.ErrorDoc(err.Error(), "400"))
	}

	out := make([]interface{}, len(metaDatas))
	for i := 0; i < len(metaDatas); i++ {
		out[i] = metaDatas[i].ToResponse()
	}

	return c.JSON(http.StatusOK, util.CollectionResponse(out, c.Request(), int(count), limit, offset))
}

func MetaDataGet(c echo.Context) error {
	metaDataDB, err := getMetaDataDao(c)
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorDoc(err.Error(), "400"))
	}

	c.Logger().Infof("Getting MetaData ID %v", id)

	app, err := metaDataDB.GetById(&id)

	if err != nil {
		if errors.Is(err, util.ErrNotFoundEmpty) {
			return err
		}
		return c.JSON(http.StatusBadRequest, util.ErrorDoc(err.Error(), "400"))
	}

	return c.JSON(http.StatusOK, app.ToResponse())
}
