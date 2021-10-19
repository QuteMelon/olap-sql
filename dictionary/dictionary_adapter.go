package dictionary

import (
	"fmt"
	"github.com/awatercolorpen/olap-sql/api/models"
)

type AdapterType string

const (
	DBAdapter   AdapterType = "DB"
	FileAdapter AdapterType = "FILE"
)

// Adapter Adapter适配器
type Adapter interface {
	// TODO 有待考虑(这里应该设计的不好)
	NewAdapter(interface{}) (interface{}, error)
}

// AdapterOption Adapter配置
type AdapterOption struct {
	Type AdapterType
	Dsn  string
}

func NewAdapter(option *AdapterOption) (*DictionaryAdapter, error) {
	// 根据不同的Type去实例化不同的Adapter
	switch option.Type {
	case DBAdapter:
		return NewDictionaryAdapterByDB(option)
	case FileAdapter:
		return NewDictionaryAdapterByYaml(option)
	}
	return nil, nil
}

// DataSaveCenter 用于保存指标的逻辑数据信息
type DictionaryAdapter struct {
	// TODO
	set        []*models.DataSet
	sources    []*models.DataSource
	metrics    []*models.Metric
	dimensions []*models.Dimension
}

func (d *DictionaryAdapter) Create(item interface{}) error {
	switch v := item.(type) {
	case *models.DataSet:
		if err := d.isValidDataSetSchema(v.Schema); err != nil {
			return err
		}
		d.set = append(d.set, v)
	case []*models.DataSet:
		for _, i := range item.([]*models.DataSet) {
			if err := d.isValidDataSetSchema(i.Schema); err != nil {
				return err
			}
			d.set = append(d.set, i)
		}
	case *models.DataSource:
		d.sources = append(d.sources, v)
	case []*models.DataSource:
		d.sources = append(d.sources, v...)
	case *models.Metric:
		d.metrics = append(d.metrics, v)
	case []*models.Metric:
		d.metrics = append(d.metrics, v...)
	case *models.Dimension:
		d.dimensions = append(d.dimensions, v)
	case []*models.Dimension:
		d.dimensions = append(d.dimensions, v...)
	}
	return nil
}


func NewDictionaryAdapterByDB(option *AdapterOption) (*DictionaryAdapter, error) {
	return nil, nil
}

func NewDictionaryAdapterByYaml(option *AdapterOption) (*DictionaryAdapter, error) {
	return nil, nil
}

func (d *DictionaryAdapter) GetDataSetByName(name string) (*models.DataSet, error) {
	for _, data := range d.set {
		if data.Name == name {
			return checkDataSetActive(data)
		}
	}
	return nil, fmt.Errorf("can not find '%v' data set", name)
}

func (d *DictionaryAdapter) GetSourcesByIds(ids []uint64) ([]*models.DataSource, error) {
	idsMap := getIdsMap(ids)
	metricsSourcesIdsMap := make(map[uint64]bool)
	for _, metric := range d.metrics {
		metricsSourcesIdsMap[metric.DataSourceID] = true
	}

	dimensionsSourcesIdsMap := make(map[uint64]bool)
	for _, dimension := range d.dimensions {
		dimensionsSourcesIdsMap[dimension.DataSourceID] = true
	}

	result := make([]*models.DataSource, 0)
	for _, source := range d.sources {
		_, ok := idsMap[source.ID]
		_, ok2 := metricsSourcesIdsMap[source.ID]
		_, ok3 := dimensionsSourcesIdsMap[source.ID]
		if ok && (ok2 || ok3) {
			result = append(result, source)
		}
	}
	return result, nil
}

func (d *DictionaryAdapter) GetMetricsByIds(ids []uint64) ([]*models.Metric, error) {
	idsMap := getIdsMap(ids)
	metrics := make([]*models.Metric, 0)
	for _, metric := range d.metrics {
		if _, ok := idsMap[metric.DataSourceID]; ok {
			metrics = append(metrics, metric)
		}
	}
	return metrics, nil
}

// GetDimensionsByIds 通过ids筛选Dimensions信息
func (d *DictionaryAdapter) GetDimensionsByIds(ids []uint64) ([]*models.Dimension, error) {
	idsMap := getIdsMap(ids)
	dimensions := make([]*models.Dimension, 0)
	for _, dimension := range d.dimensions {
		if _, ok := idsMap[dimension.DataSourceID]; ok {
			dimensions = append(dimensions, dimension)
		}
	}
	return dimensions, nil
}

func checkDataSetActive(set *models.DataSet) (*models.DataSet, error) {
	if set.Schema == nil {
		return nil, fmt.Errorf("schema is nil for data_set %v", set.Name)
	}
	return set, nil
}

func (d *DictionaryAdapter) isValidJoinOns(joinOns models.JoinOns) (id1, id2 uint64, err error) {
	in1, in2 := joinOns.ID()

	in1Map := getIdsMap(in1)
	in2Map := getIdsMap(in2)

	out1 := make(map[uint64]bool, 0)
	out2 := make(map[uint64]bool, 0)

	for _, dimension := range d.dimensions {
		id := dimension.ID
		if _, ok := in1Map[id]; ok {
			out1[dimension.DataSourceID] = true
		}
		if _, ok := in2Map[id]; ok {
			out2[dimension.DataSourceID] = true
		}
	}

	if len(out1) != 1 {
		return 0, 0, fmt.Errorf("invalid data_source_id=%v", out1)
	}
	if len(out2) != 1 {
		return 0, 0, fmt.Errorf("invalid data_source_id=%v", out2)
	}

	for id := range out1 {
		id1 = id
	}

	for id := range out2 {
		id2 = id
	}
	return
}

func (d *DictionaryAdapter) isValidSecondary(secondary *models.Secondary) error {
	id1, id2, err := d.isValidJoinOns(models.JoinOns(secondary.JoinOn))
	if err != nil {
		return err
	}
	if id1 != secondary.DataSourceID1 {
		return fmt.Errorf("unmatched data_source_ids, %v != %v", id1, secondary.DataSourceID1)
	}
	if id2 != secondary.DataSourceID2 {
		return fmt.Errorf("unmatched data_source_ids, %v != %v", id2, secondary.DataSourceID2)
	}
	return nil
}

func (d *DictionaryAdapter) isValidDataSetSchema(schema *models.DataSetSchema) error {
	if _, err := schema.Tree(); err != nil {
		return err
	}

	for _, v := range schema.Secondary {
		if err := d.isValidSecondary(v); err != nil {
			return err
		}
	}
	return nil
}

// isValidDataSet 检查DataSet的合法性
func (d *DictionaryAdapter) isValidDataSet(set *models.DataSet) error {
	return d.isValidDataSetSchema(set.Schema)
}

// isValidAdapterCheck 检查Adapter的合法性
func (d *DictionaryAdapter) isValidAdapterCheck() error {
	for _, set := range d.set {
		if err := d.isValidDataSet(set); err != nil {
			return err
		}
	}
	return nil
}

// fillSourceMetricsAndDimensions 填充 Sources 的外键信息 Metrics 和 Dimensions
func (d *DictionaryAdapter) fillSourceMetricsAndDimensions(){
	for _, source := range d.sources {
		for _, metric := range d.metrics {
			if metric.DataSourceID == source.ID {
				source.Metrics = append(source.Metrics, metric)
			}
		}
		for _, dimension := range d.dimensions {
			if dimension.DataSourceID == source.ID {
				source.Dimensions = append(source.Dimensions, dimension)
			}
		}
	}
}

func getIdsMap(ids []uint64) map[interface{}]interface{} {
	idsMap := make(map[interface{}]interface{})
	for _, id := range ids {
		idsMap[id] = true
	}
	return idsMap
}
