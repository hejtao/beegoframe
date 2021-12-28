package internal

import (
	"beegoframe/pkg/encoding/json"
	"bytes"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
)

var ctx = context.Background()

var wrongDocTypeErr = errors.New("wrong doc type, need struct or struct ptr")

type Index string

func (r Index) getName() string {
	return prefix + string(r)
}

func (r Index) Create(doc interface{}, id int64) (*elastic.IndexResponse, error) {
	idStr := strconv.Itoa(int(id))
	resp, err := client.Index().
		Index(r.getName()).
		Id(idStr).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r Index) ReadRow(container interface{}, id int64) error {
	v := reflect.ValueOf(container)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return wrongDocTypeErr
	}

	idStr := strconv.Itoa(int(id + 10000000))
	res, err := client.Get().
		Index(r.getName()).
		Id(idStr).
		Do(ctx)
	if err != nil {
		return err
	}

	if res.Found {
		data, err := res.Source.MarshalJSON()
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, container)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r Index) ReadRows(container interface{}, query Query, pageNum, pageSize int, order string, ascending bool) error {
	t := reflect.TypeOf(container)
	if t.Kind() != reflect.Ptr {
		return wrongDocTypeErr
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		return wrongDocTypeErr
	}

	if pageSize <= 0 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageSize = 1
	}
	res, err := client.Search().
		Index(r.getName()).
		Query(query).
		From((pageNum-1)*pageSize).
		Size(pageSize).
		Sort(order, ascending).
		Do(ctx)
	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}
	buffer.Write([]byte{'['})
	for k, hit := range res.Hits.Hits {
		if k == 0 {
			buffer.Write(hit.Source)
		} else {
			buffer.Write([]byte{','})
			buffer.Write(hit.Source)
		}
	}
	buffer.Write([]byte{']'})
	b := json.FormatTime(buffer.Bytes())
	err = json.Unmarshal(b, container)
	return err
}

func (r Index) Update(doc interface{}, id int64) (*elastic.UpdateResponse, error) {
	v := reflect.ValueOf(doc)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, wrongDocTypeErr
	}

	idStr := strconv.Itoa(int(id))

	resp, err := client.Update().
		Index(r.getName()).
		Id(idStr).
		Doc(doc).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//func (r Index) BulkUpdate(doc interface{}, query Query) (*elastic.BulkIndexByScrollResponse, error) {
//	v := reflect.ValueOf(doc)
//	if v.Kind() == reflect.Ptr {
//		v = v.Elem()
//	}
//	if v.Kind() != reflect.Struct {
//		return nil, wrongDocTypeErr
//	}
//
//	b, err := json.Marshal(doc)
//	_ = b
//	if err != nil {
//		return nil, err
//	}
//	resp, err := client.UpdateByQuery(r.getName()).
//		Query(query).
//		Script(elastic.NewScript( fmt.Sprintf("ctx._source['update_time']=%v", time.Now()))).
//		ProceedOnVersionConflict().
//		Do(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}

func (r Index) Delete(id int64) (*elastic.DeleteResponse, error) {
	idStr := strconv.Itoa(int(id))
	resp, err := client.Delete().
		Index(r.getName()).
		Id(idStr).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r Index) BulkDelete(query Query) (*elastic.BulkIndexByScrollResponse, error) {
	resp, err := client.DeleteByQuery(r.getName()).
		Query(query).
		ProceedOnVersionConflict().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
