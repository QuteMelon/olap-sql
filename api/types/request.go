package types

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Request struct {
	Metrics    []*Metric    `json:"metrics"`
	Dimensions []*Dimension `json:"dimensions"`
	Filters    []*Filter    `json:"filters"`
	Joins      []*Join      `json:"jsons"`
	DataSource *DataSource  `json:"data_source"`
}

func (r *Request) Clause(tx *gorm.DB) (*gorm.DB, error) {
	select1, err := r.metricStatement()
	if err != nil {
		return nil, err
	}
	tx = tx.Select(select1)

	select2, err := r.dimensionStatement()
	if err != nil {
		return nil, err
	}
	tx = tx.Select(select2)

	where1, err := r.filterStatement()
	if err != nil {
		return nil, err
	}
	for _, v := range where1 {
		tx = tx.Where(v)
	}

	join1, err := r.joinStatement()
	if err != nil {
		return nil, err
	}
	for _, v := range join1 {
		tx = tx.Joins(v)
	}

	group1, err := r.groupStatement()
	if err != nil {
		return nil, err
	}
	for _, v := range group1 {
		tx = tx.Group(v)
	}

	table1, err := r.tableStatement()
	if err != nil {
		return nil, err
	}
	tx = tx.Table(table1)

	return tx, nil
}


func (r *Request) Statement() (string, error) {
	statement1, err := r.metricStatement()
	if err != nil {
		return "", err
	}

	statement2, err := r.dimensionStatement()
	if err != nil {
		return "", err
	}

	statement3, err := r.filterStatement()
	if err != nil {
		return "", err
	}

	statement4, err := r.joinStatement()
	if err != nil {
		return "", err
	}

	statement5, err := r.groupStatement()
	if err != nil {
		return "", err
	}

	statement6, err := r.tableStatement()
	if err != nil {
		return "", err
	}

	return r.buildSql(statement1, statement2, statement3, statement4, statement5, statement6), nil
}

func (r *Request) metricStatement() ([]string, error) {
	var statement []string
	for _, v := range r.Metrics {
		s, err := v.Statement()
		if err != nil {
			return nil, err
		}
		statement = append(statement, s)
	}

	return statement, nil
}

func (r *Request) dimensionStatement() ([]string, error) {
	var statement []string
	for _, v := range r.Dimensions {
		s, err := v.Statement()
		if err != nil {
			return nil, err
		}
		statement = append(statement, s)
	}

	return statement, nil
}

func (r *Request) filterStatement() ([]string, error) {
	var statement []string
	for _, v := range r.Filters {
		s, err := v.Statement()
		if err != nil {
			return nil, err
		}
		statement = append(statement, s)
	}
	return statement, nil
}

func (r *Request) joinStatement() ([]string, error) {
	if r.DataSource == nil {
		return nil, fmt.Errorf("nil data source")
	}

	var statement []string
	for _, v := range r.Joins {
		switch r.DataSource.Type {
		case DataSourceTypeKylin, DataSourceTypePresto, DataSourceTypeClickHouse:
			var on []string
			for _, u := range v.On {
				on = append(on, fmt.Sprintf("%v.%v = %v.%v", v.Table1, u.Key1, v.Table2, u.Key2))
			}
			statement = append(statement, fmt.Sprintf("LEFT JOIN %v ON %v", v.Table2, strings.Join(on, " AND ")))
		// case DataSourceTypeClickHouse:
		// 	statement = append(statement, fmt.Sprintf("t1 LEFT JOIN %v ON t1.%v = %v.%v", v.Table2, v.Key1, v.Table2, v.Key2))
		default:
			return nil, fmt.Errorf("not supported data source type %v", r.DataSource.Type)
		}
	}
	return statement, nil
}

func (r *Request) groupStatement() ([]string, error) {
	var statement []string
	for _, v := range r.Dimensions {
		statement = append(statement, v.Name)
	}

	return statement, nil
}

func (r *Request) tableStatement() (string, error) {
	return r.DataSource.Statement()
}

func (r *Request) buildSql(metrics, dimensions, filters, joins, groups []string, table string) string {
	selectCol := append([]string{}, metrics...)
	selectCol = append(selectCol, dimensions...)

	selectStatement := strings.Join(selectCol, " , ")
	groupStatement := strings.Join(groups, " , ")
	whereStatement := strings.Join(filters, " AND ")
	joinStatement := strings.Join(joins, " ")

	sql := fmt.Sprintf("SELECT %v FROM %v", selectStatement, table)
	if joinStatement != "" {
		sql = fmt.Sprintf("%v %v", sql, joinStatement)
	}
	if whereStatement != "" {
		sql = fmt.Sprintf("%v WHERE %v", sql, whereStatement)
	}
	if groupStatement != "" {
		sql = fmt.Sprintf("%v GROUP BY %v", sql, groupStatement)
	}

	return sql
}
