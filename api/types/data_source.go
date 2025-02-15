package types

import (
	"fmt"

	"gorm.io/gorm"
)

type DataSourceType string

const (
	DataSourceTypeFact              DataSourceType = "fact"
	DataSourceTypeDimension         DataSourceType = "dimension"
	DataSourceTypeFactDimensionJoin DataSourceType = "fact_dimension_join"
	DataSourceTypeMergeJoin         DataSourceType = "merge_join"
)

type DataSource struct {
	Database  string         `json:"database"`
	Name      string         `json:"name"`
	AliasName string         `json:"alias"`
	Type      DataSourceType `json:"type"`
	JoinType  string         `json:"join_type"`
	Clause    Clause         `json:"clause"`

	expression string
}

func (d *DataSource) Expression() (string, error) {
	if d.expression != "" {
		switch d.Clause {
		case nil:
			return d.expression, nil
		default:
			return fmt.Sprintf("( %v )", d.expression), nil
		}
	}
	if d.Database == "" {
		return fmt.Sprintf("`%v`", d.Name), nil
	}
	return fmt.Sprintf("`%v`.`%v`", d.Database, d.Name), nil
}

func (d *DataSource) Alias() (string, error) {
	if d.AliasName == "" {
		return d.Name, nil
	}
	return d.AliasName, nil
}

func (d *DataSource) Statement() (string, error) {
	expression, _ := d.Expression()
	alias, _ := d.Alias()
	return fmt.Sprintf("%v AS %v", expression, alias), nil
}

func (d *DataSource) Init(tx *gorm.DB) (err error) {
	switch d.Clause {
	case nil:
		d.expression, err = d.Expression()
	default:
		d.expression, err = d.Clause.BuildSQL(tx)
	}
	return
}

func (d *DataSource) GetJoinType() string {
	if len(d.JoinType) == 0 {
		return "LEFT JOIN"
	}
	return d.JoinType
}
