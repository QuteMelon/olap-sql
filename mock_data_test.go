package olapsql_test

import (
	"os"
	"time"

	"github.com/awatercolorpen/olap-sql"
	"github.com/awatercolorpen/olap-sql/api/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const mockWikiStatDataSet = "wikistat"
const mockWikiStatDataSetJoin = "wikistat_join"
const mockWikiStatDataSetMerge = "merge_uv"

type WikiStat struct {
	Date       time.Time `gorm:"column:date"       json:"date"`
	Time       time.Time `gorm:"column:time"       json:"time"`
	Project    string    `gorm:"column:project"    json:"project"`
	SubProject string    `gorm:"column:subproject" json:"subproject"`
	Path       string    `gorm:"column:path"       json:"path"`
	Hits       uint64    `gorm:"column:hits"       json:"hits"`
	Size       uint64    `gorm:"column:size"       json:"size"`
}

func (WikiStat) TableName() string {
	return mockWikiStatDataSet
}

type WikiStatRelate struct {
	Project string  `gorm:"column:project" json:"project"`
	Class   uint64  `gorm:"column:class"   json:"class"`
	Source  float64 `gorm:"column:source"  json:"source"`
}

func (WikiStatRelate) TableName() string {
	return mockWikiStatDataSet + "_relate"
}

type ClassRelate struct {
	ID   uint64 `gorm:"column:id"   json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (ClassRelate) TableName() string {
	return mockWikiStatDataSet + "_class"
}

type MergeUv struct {
	Date          time.Time `gorm:"column:date"       json:"date"`
	Time          time.Time `gorm:"column:time"       json:"time"`
	SubProject    string    `gorm:"column:sub_project" json:"sub_project"`
	UID           string    `gorm:"column:uid"   json:"uid"`
	ActivationCnt uint64    `gorm:"column:activation_cnt"  json:"activation_cnt"`
	Cost          float64   `gorm:"column:cost"  json:"cost"`
}

func (MergeUv) TableName() string {
	return mockWikiStatDataSet + "_uv"
}

func timeParseDate(in string) time.Time {
	t, _ := time.Parse("2006-01-02", in)
	return t
}

func timeParseTime(in string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z", in)
	return t
}

func DataWithClickhouse() bool {
	args := os.Args
	for _, arg := range args {
		if arg == "clickhouse" {
			return true
		}
	}
	return false
}

func DBType() types.DBType {
	if DataWithClickhouse() {
		return types.DBTypeClickHouse
	}
	return types.DBTypeSQLite
}

func MockWikiStatData(db *gorm.DB) error {
	if DataWithClickhouse() {
		return nil
	}

	if err := db.Debug().AutoMigrate(&WikiStat{}, &WikiStatRelate{}, &ClassRelate{}, &MergeUv{}); err != nil {
		return err
	}
	if err := db.Debug().Create([]*WikiStat{
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T11:45:26Z"), Project: "city", SubProject: "CHN", Path: "level1", Hits: 121, Size: 4098},
		{Date: timeParseDate("2021-05-06"), Time: timeParseTime("2021-05-06T11:45:26Z"), Project: "city", SubProject: "CHN", Path: "level1", Hits: 139, Size: 10086},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T12:43:56Z"), Project: "city", SubProject: "CHN", Path: "level2", Hits: 20, Size: 1024},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T07:00:12Z"), Project: "city", SubProject: "US", Path: "level1", Hits: 19, Size: 2048},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T21:23:48Z"), Project: "school", SubProject: "university", Path: "engineering", Hits: 2, Size: 156},
		{Date: timeParseDate("2021-05-06"), Time: timeParseTime("2021-05-06T21:16:39Z"), Project: "school", SubProject: "university", Path: "engineering", Hits: 3, Size: 158},
		{Date: timeParseDate("2021-05-06"), Time: timeParseTime("2021-05-06T20:32:41Z"), Project: "school", SubProject: "senior", Path: "*", Hits: 5, Size: 212},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T09:28:27Z"), Project: "music", SubProject: "pop", Path: "", Hits: 4783, Size: 37291},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T09:31:23Z"), Project: "music", SubProject: "pop", Path: "ancient", Hits: 391, Size: 2531},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T09:33:59Z"), Project: "music", SubProject: "rap", Path: "", Hits: 1842, Size: 12942},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T10:34:12Z"), Project: "music", SubProject: "rock", Path: "", Hits: 0, Size: 0},
	}).Error; err != nil {
		return err
	}
	if err := db.Debug().Create([]*WikiStatRelate{
		{Project: "city", Class: 1, Source: 4.872},
		{Project: "school", Class: 1, Source: 0.18742},
		{Project: "food", Class: 2, Source: 10.2484},
		{Project: "person", Class: 3, Source: 1.73},
		{Project: "music", Class: 4, Source: 93.20},
		{Project: "company", Class: 5, Source: 0.0281},
	}).Error; err != nil {
		return err
	}
	if err := db.Debug().Create([]*ClassRelate{
		{ID: 1, Name: "location"},
		{ID: 2, Name: "life"},
		{ID: 3, Name: "culture"},
		{ID: 4, Name: "entertainment"},
		{ID: 5, Name: "social"},
	}).Error; err != nil {
		return err
	}
	if err := db.Debug().Create([]*MergeUv{
		{Date: timeParseDate("2021-05-06"), Time: timeParseTime("2021-05-06T10:00:00Z"), SubProject: "CHN", UID: "aaa", ActivationCnt: 0, Cost: 0.0},
		{Date: timeParseDate("2021-05-06"), Time: timeParseTime("2021-05-06T21:00:00Z"), SubProject: "university", UID: "pl-okm", ActivationCnt: 3, Cost: 0.45},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T07:00:00Z"), SubProject: "US", UID: "aaa", ActivationCnt: 4, Cost: 0.51},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T09:00:00Z"), SubProject: "pop", UID: "12345678", ActivationCnt: 1, Cost: 0.08},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T09:00:00Z"), SubProject: "pop", UID: "qwerty", ActivationCnt: 2, Cost: 0.15},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T09:00:00Z"), SubProject: "rap", UID: "12345678", ActivationCnt: 3, Cost: 0.21},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T10:00:00Z"), SubProject: "rock", UID: "pl-okm", ActivationCnt: 10, Cost: 1.09},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T11:00:00Z"), SubProject: "CHN", UID: "12345678", ActivationCnt: 2, Cost: 0.34},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T11:00:00Z"), SubProject: "CHN", UID: "qwerty", ActivationCnt: 1, Cost: 0.22},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T12:00:00Z"), SubProject: "CHN", UID: "12345678", ActivationCnt: 1, Cost: 0.12},
		{Date: timeParseDate("2021-05-07"), Time: timeParseTime("2021-05-07T21:00:00Z"), SubProject: "university", UID: "12345678", ActivationCnt: 6, Cost: 1.12},
	}).Error; err != nil {
		return err
	}
	return nil
}

func MockLoad(manager *olapsql.Manager) error {
	client, _ := manager.GetClients()
	db, err := client.Get(DBType(), "")
	if err != nil {
		return err
	}
	return MockWikiStatData(db)
}

// MockQuery1 mock query for normal fact case
func MockQuery1() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSet,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"hits", "size_sum", "hits_avg", "hits_per_size"},
		Dimensions:   []string{"date"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeNotIn, Name: "path", Value: []interface{}{"*"}},
		},
		Orders: []*types.OrderBy{
			{Name: "size_sum", Direction: types.OrderDirectionTypeDescending},
		},
		Limit: &types.Limit{Limit: 2, Offset: 1},
	}
	return query
}

func MockQuery1ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 5)
	assert.Equal(t, "date", result.Dimensions[0])
	assert.Equal(t, "hits_per_size", result.Dimensions[4])
	assert.Len(t, result.Source, 1)
	assert.Len(t, result.Source[0], 5)
	assert.Equal(t, float64(10244), result.Source[0]["size_sum"])
	assert.Equal(t, 0.013861772745021476, getValue(result.Source[0]["hits_per_size"]))
}

// MockQuery2 mock query for normal fact dimension join case
func MockQuery2() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetJoin,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"hits", "size_sum", "hits_avg", "hits_per_size", "source_avg"},
		Dimensions:   []string{"date", "class_id"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeNotIn, Name: "path", Value: []interface{}{"*"}},
			{OperatorType: types.FilterOperatorTypeIn, Name: "class_id", Value: []interface{}{1, 2, 3, 4}},
		},
		Orders: []*types.OrderBy{
			{Name: "source_sum", Direction: types.OrderDirectionTypeDescending},
		},
		Limit: &types.Limit{Limit: 2, Offset: 1},
	}
	return query
}

func MockQuery2ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 7)
	assert.Equal(t, "date", result.Dimensions[0])
	assert.Equal(t, "source_avg", result.Dimensions[6])
	assert.Len(t, result.Source, 2)
	assert.Len(t, result.Source[0], 7)
	assert.Equal(t, float64(7326), result.Source[0]["size_sum"])

	assert.Equal(t, 0.022113022113022112, getValue(result.Source[0]["hits_per_size"]))
	assert.Equal(t, 3.700855, getValue(result.Source[0]["source_avg"]))
}

// MockQuery3 mock query for group by time case
func MockQuery3() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetJoin,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"hits", "size_sum", "hits_avg", "hits_per_size", "source_avg"},
		Dimensions:   []string{"time_by_hour", "class_id"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeNotIn, Name: "path", Value: []interface{}{"*"}},
			{OperatorType: types.FilterOperatorTypeIn, Name: "class_id", Value: []interface{}{1, 2, 3, 4}},
		},
		Orders: []*types.OrderBy{
			{Name: "time_by_hour", Direction: types.OrderDirectionTypeAscending},
		},
		Limit: &types.Limit{Limit: 10},
	}
	return query
}

func MockQuery3ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 7)
	assert.Equal(t, "time_by_hour", result.Dimensions[0])
	assert.Equal(t, "source_avg", result.Dimensions[6])
	assert.Len(t, result.Source, 8)
	assert.Len(t, result.Source[0], 7)
	assert.Equal(t, "2021-05-06 11:00:00", result.Source[0]["time_by_hour"])
	assert.Equal(t, float64(10086), result.Source[0]["size_sum"])
	assert.Equal(t, 0.013781479278207416, getValue(result.Source[0]["hits_per_size"]))
	assert.Equal(t, 4.872, getValue(result.Source[0]["source_avg"]))
}

// MockQuery4 mock query for nested join case
func MockQuery4() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetJoin,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"source_avg"},
		Dimensions:   []string{"project", "class_name"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeNotIn, Name: "path", Value: []interface{}{"*"}},
			{OperatorType: types.FilterOperatorTypeIn, Name: "project", Value: []interface{}{"city", "school", "music"}},
		},
		Orders: []*types.OrderBy{
			{Name: "project", Direction: types.OrderDirectionTypeDescending},
		},
	}
	return query
}

func MockQuery4ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 3)
	assert.Equal(t, "project", result.Dimensions[0])
	assert.Equal(t, "source_avg", result.Dimensions[2])
	assert.Len(t, result.Source, 3)
	assert.Len(t, result.Source[0], 3)
	assert.Equal(t, "school", result.Source[0]["project"])
	assert.Equal(t, "location", result.Source[0]["class_name"])
	assert.Equal(t, 0.18742, getValue(result.Source[0]["source_avg"]))
}

// MockQuery5 mock query for nan / inf case
func MockQuery5() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetJoin,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"hits_per_size"},
		Dimensions:   []string{"class_name"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeIn, Name: "time_by_hour", Value: []interface{}{"2021-05-07 10:00:00"}},
		},
	}
	return query
}

func MockQuery5ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 2)
	assert.Equal(t, "class_name", result.Dimensions[0])
	assert.Equal(t, "hits_per_size", result.Dimensions[1])
	assert.Len(t, result.Source, 1)
	assert.Len(t, result.Source[0], 2)
	assert.Equal(t, "entertainment", result.Source[0]["class_name"])
	assert.Equal(t, nil, result.Source[0]["hits_per_size"])
}

// MockQuery6 mock query for count distinct if
func MockQuery6() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetJoin,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"project_count"},
	}
	return query
}

func MockQuery6ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 1)
	assert.Len(t, result.Source, 1)
	assert.Len(t, result.Source[0], 1)
	assert.Equal(t, float64(3), result.Source[0]["project_count"])
}

// MockQuery7 mock query for sum if
func MockQuery7() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetJoin,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"hits_sum"},
	}
	return query
}

func MockQuery7ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 1)
	assert.Len(t, result.Source, 1)
	assert.Len(t, result.Source[0], 1)
	assert.Equal(t, float64(7325), getValue(result.Source[0]["hits_sum"]))
}

// MockQuery8 mock query for easy merge join
func MockQuery8() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetMerge,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"hits", "activation_cnt", "cost", "activation_rate"},
		Dimensions:   []string{"time_by_hour", "sub_project"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeGreaterEquals, Name: "time_by_hour", Value: []interface{}{"2021-05-07 06:00:00"}},
		},
		Orders: []*types.OrderBy{
			{Name: "time_by_hour", Direction: types.OrderDirectionTypeAscending},
			{Name: "sub_project", Direction: types.OrderDirectionTypeAscending},
		},
	}
	return query
}

func MockQuery8ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 6)
	assert.Len(t, result.Source, 7)
	assert.Len(t, result.Source[0], 6)
	assert.Equal(t, 0.21052631578947367, getValue(result.Source[0]["activation_rate"]))
}

// MockQuery9 mock query for mixed merge join
func MockQuery9() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetMerge,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"cost", "hits_per_user"},
		Dimensions:   []string{"sub_project", "class_name"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeIn, Name: "path", Value: []interface{}{"level1", "level2", "engineering"}},
		},
		Orders: []*types.OrderBy{
			{Name: "class_name", Direction: types.OrderDirectionTypeAscending},
		},
	}
	return query
}

func MockQuery9ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 4)
	assert.Len(t, result.Source, 3)
	assert.Len(t, result.Source[0], 4)
	assert.Equal(t, 93.33333333333333, getValue(result.Source[0]["hits_per_user"]))
}

// MockQuery10 mock query for merge join without dimension
func MockQuery10() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetMerge,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"cost", "hits_per_user"},
		Dimensions:   []string{"class_name"},
		Filters: []*types.Filter{
			{OperatorType: types.FilterOperatorTypeIn, Name: "sub_project", Value: []interface{}{"CHN", "pop", "university", "senior"}},
		},
		Orders: []*types.OrderBy{
			{Name: "class_name", Direction: types.OrderDirectionTypeAscending},
		},
	}
	return query
}

func MockQuery10ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 3)
	assert.Len(t, result.Source, 2)
	assert.Len(t, result.Source[0], 3)
	assert.Equal(t, 1293.5, getValue(result.Source[0]["hits_per_user"]))
}

// MockQuery11 mock query for merge join with dimension
func MockQuery11() *types.Query {
	query := &types.Query{
		DataSetName:  mockWikiStatDataSetMerge,
		TimeInterval: &types.TimeInterval{Name: "date", Start: "2021-05-06", End: "2021-05-08"},
		Metrics:      []string{"big_cost_per_hits"},
		Dimensions:   []string{"time_by_hour"},
		Orders: []*types.OrderBy{
			{Name: "time_by_hour", Direction: types.OrderDirectionTypeAscending},
		},
	}
	return query
}

func MockQuery11ResultAssert(t assert.TestingT, result *types.Result) {
	assert.Len(t, result.Dimensions, 2)
	assert.Len(t, result.Source, 9)
	assert.Len(t, result.Source[0], 2)
	assert.Equal(t, nil, result.Source[0]["big_cost_per_hits"])
	assert.Equal(t, 0.56, getValue(result.Source[8]["big_cost_per_hits"]))
}

// func MockClause() *types.NormalClause {
// 	clause1 := &types.NormalClause{
// 		Metrics: []*types.Metric{
// 			{
// 				Type:      types.MetricTypeValue,
// 				Table:     "wikistat",
// 				Name:      "hits",
// 				FieldName: "hits",
// 				DBType:    types.DBTypeSQLite,
// 			},
// 		},
// 		DataSource: []*types.DataSource{{Name: "wikistat"}},
// 	}
// 	source := []*types.DataSource{
// 		{
//
// 		},
// 	}
// 	clause := &types.NormalClause{
// 		DBType:  types.DBTypeSQLite,
// 		Dataset: mockWikiStatDataSet,
// 		Metrics: []*types.Metric{
// 			{
// 				Type:      types.MetricTypeValue,
// 				Table:     "t1",
// 				Name:      "hits",
// 				FieldName: "hits",
// 				DBType:    types.DBTypeSQLite,
// 			},
// 		},
// 		DataSource: &types.DataSource{
// 			Alias: "t1",
// 			Clause: clause1,
// 		},
// 		Joins: []*types.Join{
// 			{
// 				DataSource1: &types.DataSource{
// 					Alias: "t1",
// 					Clause: &types.NormalClause{
// 						DBType:  types.DBTypeSQLite,
// 						Dataset: mockWikiStatDataSet,
// 						Metrics: []*types.Metric{
// 							{
// 								Type:      types.MetricTypeValue,
// 								Table:     "wikistat",
// 								Name:      "hits",
// 								FieldName: "hits",
// 								DBType:    types.DBTypeSQLite,
// 							},
// 						},
// 						DataSource: &types.DataSource{
// 							Name: "wikistat",
// 						},
// 					},
// 				},
// 				DataSource2: &types.DataSource{
// 					Alias: "t2",
// 					Clause: &types.NormalClause{
// 						DBType:  types.DBTypeSQLite,
// 						Dataset: mockWikiStatDataSet,
// 						Metrics: []*types.Metric{
// 							{
// 								Type:      types.MetricTypeValue,
// 								Table:     "wikistat",
// 								Name:      "hits",
// 								FieldName: "hits",
// 								DBType:    types.DBTypeSQLite,
// 							},
// 						},
// 						DataSource: &types.DataSource{
// 							Name: "wikistat",
// 						},
// 					},
// 				},
// 				On: []*types.JoinOn{
// 					{
// 						Key1: "hits",
// 						Key2: "hits",
// 					},
// 				},
// 			},
// 		},
// 		Filters: []*types.Filter{
// 			{
// 				OperatorType: types.FilterOperatorTypeGreaterEquals,
// 				ValueType:    types.ValueTypeInteger,
// 				Name:         "hits",
// 				Value:        []interface{}{1},
// 			},
// 		},
// 	}
// 	return request
// }

func getValue(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case *float64:
		if v != nil {
			return *v
		}
		return -123456789
	default:
		return -123456789
	}
}
