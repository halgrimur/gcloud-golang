// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/datastore/internal/proto/query.proto
// DO NOT EDIT!

package google_datastore_v1beta3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/cloud/datastore/internal/type_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specifies what data the 'entity' field contains.
// A `ResultType` is either implied (for example, in `LookupResponse.missing`
// from `datastore.proto`, it is always `KEY_ONLY`) or specified by context
// (for example, in message `QueryResultBatch`, field `entity_result_type`
// specifies a `ResultType` for all the values in field `entity_results`).
type EntityResult_ResultType int32

const (
	// Unspecified. This value is never used.
	EntityResult_RESULT_TYPE_UNSPECIFIED EntityResult_ResultType = 0
	// The key and properties.
	EntityResult_FULL EntityResult_ResultType = 1
	// A projected subset of properties. The entity may have no key.
	EntityResult_PROJECTION EntityResult_ResultType = 2
	// Only the key.
	EntityResult_KEY_ONLY EntityResult_ResultType = 3
)

var EntityResult_ResultType_name = map[int32]string{
	0: "RESULT_TYPE_UNSPECIFIED",
	1: "FULL",
	2: "PROJECTION",
	3: "KEY_ONLY",
}
var EntityResult_ResultType_value = map[string]int32{
	"RESULT_TYPE_UNSPECIFIED": 0,
	"FULL":       1,
	"PROJECTION": 2,
	"KEY_ONLY":   3,
}

func (x EntityResult_ResultType) String() string {
	return proto.EnumName(EntityResult_ResultType_name, int32(x))
}
func (EntityResult_ResultType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

// The sort direction.
type PropertyOrder_Direction int32

const (
	// Unspecified. This value must not be used.
	PropertyOrder_DIRECTION_UNSPECIFIED PropertyOrder_Direction = 0
	// Ascending.
	PropertyOrder_ASCENDING PropertyOrder_Direction = 1
	// Descending.
	PropertyOrder_DESCENDING PropertyOrder_Direction = 2
)

var PropertyOrder_Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "ASCENDING",
	2: "DESCENDING",
}
var PropertyOrder_Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"ASCENDING":             1,
	"DESCENDING":            2,
}

func (x PropertyOrder_Direction) String() string {
	return proto.EnumName(PropertyOrder_Direction_name, int32(x))
}
func (PropertyOrder_Direction) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{5, 0} }

// A composite filter operator.
type CompositeFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	CompositeFilter_OPERATOR_UNSPECIFIED CompositeFilter_Operator = 0
	// The results are required to satisfy each of the combined filters.
	CompositeFilter_AND CompositeFilter_Operator = 1
)

var CompositeFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	1: "AND",
}
var CompositeFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED": 0,
	"AND": 1,
}

func (x CompositeFilter_Operator) String() string {
	return proto.EnumName(CompositeFilter_Operator_name, int32(x))
}
func (CompositeFilter_Operator) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{7, 0} }

// A property filter operator.
type PropertyFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	PropertyFilter_OPERATOR_UNSPECIFIED PropertyFilter_Operator = 0
	// Less than.
	PropertyFilter_LESS_THAN PropertyFilter_Operator = 1
	// Less than or equal.
	PropertyFilter_LESS_THAN_OR_EQUAL PropertyFilter_Operator = 2
	// Greater than.
	PropertyFilter_GREATER_THAN PropertyFilter_Operator = 3
	// Greater than or equal.
	PropertyFilter_GREATER_THAN_OR_EQUAL PropertyFilter_Operator = 4
	// Equal.
	PropertyFilter_EQUAL PropertyFilter_Operator = 5
	// Has ancestor.
	PropertyFilter_HAS_ANCESTOR PropertyFilter_Operator = 11
)

var PropertyFilter_Operator_name = map[int32]string{
	0:  "OPERATOR_UNSPECIFIED",
	1:  "LESS_THAN",
	2:  "LESS_THAN_OR_EQUAL",
	3:  "GREATER_THAN",
	4:  "GREATER_THAN_OR_EQUAL",
	5:  "EQUAL",
	11: "HAS_ANCESTOR",
}
var PropertyFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED":  0,
	"LESS_THAN":             1,
	"LESS_THAN_OR_EQUAL":    2,
	"GREATER_THAN":          3,
	"GREATER_THAN_OR_EQUAL": 4,
	"EQUAL":                 5,
	"HAS_ANCESTOR":          11,
}

func (x PropertyFilter_Operator) String() string {
	return proto.EnumName(PropertyFilter_Operator_name, int32(x))
}
func (PropertyFilter_Operator) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{8, 0} }

// The possible values for the `more_results` field.
type QueryResultBatch_MoreResultsType int32

const (
	// Unspecified. This value is never used.
	QueryResultBatch_MORE_RESULTS_TYPE_UNSPECIFIED QueryResultBatch_MoreResultsType = 0
	// There may be additional batches to fetch from this query.
	QueryResultBatch_NOT_FINISHED QueryResultBatch_MoreResultsType = 1
	// The query is finished, but there may be more results after the limit.
	QueryResultBatch_MORE_RESULTS_AFTER_LIMIT QueryResultBatch_MoreResultsType = 2
	// The query is finished, but there may be more results after the end cursor.
	QueryResultBatch_MORE_RESULTS_AFTER_CURSOR QueryResultBatch_MoreResultsType = 4
	// The query has been exhausted.
	QueryResultBatch_NO_MORE_RESULTS QueryResultBatch_MoreResultsType = 3
)

var QueryResultBatch_MoreResultsType_name = map[int32]string{
	0: "MORE_RESULTS_TYPE_UNSPECIFIED",
	1: "NOT_FINISHED",
	2: "MORE_RESULTS_AFTER_LIMIT",
	4: "MORE_RESULTS_AFTER_CURSOR",
	3: "NO_MORE_RESULTS",
}
var QueryResultBatch_MoreResultsType_value = map[string]int32{
	"MORE_RESULTS_TYPE_UNSPECIFIED": 0,
	"NOT_FINISHED":                  1,
	"MORE_RESULTS_AFTER_LIMIT":      2,
	"MORE_RESULTS_AFTER_CURSOR":     4,
	"NO_MORE_RESULTS":               3,
}

func (x QueryResultBatch_MoreResultsType) String() string {
	return proto.EnumName(QueryResultBatch_MoreResultsType_name, int32(x))
}
func (QueryResultBatch_MoreResultsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{11, 0}
}

// The result of fetching an entity from Datastore.
type EntityResult struct {
	// The resulting entity.
	Entity *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	// A cursor that points to the position after the result entity.
	// Set only when the `EntityResult` is part of a `QueryResultBatch` message.
	Cursor []byte `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (m *EntityResult) Reset()                    { *m = EntityResult{} }
func (m *EntityResult) String() string            { return proto.CompactTextString(m) }
func (*EntityResult) ProtoMessage()               {}
func (*EntityResult) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EntityResult) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// A query for entities.
type Query struct {
	// The projection to return. Defaults to returning all properties.
	Projection []*Projection `protobuf:"bytes,2,rep,name=projection" json:"projection,omitempty"`
	// The kinds to query (if empty, returns entities of all kinds).
	// Currently at most 1 kind may be specified.
	Kind []*KindExpression `protobuf:"bytes,3,rep,name=kind" json:"kind,omitempty"`
	// The filter to apply.
	Filter *Filter `protobuf:"bytes,4,opt,name=filter" json:"filter,omitempty"`
	// The order to apply to the query results (if empty, order is unspecified).
	Order []*PropertyOrder `protobuf:"bytes,5,rep,name=order" json:"order,omitempty"`
	// The properties to make distinct. The query results will contain the first
	// result for each distinct combination of values for the given properties
	// (if empty, all results are returned).
	DistinctOn []*PropertyReference `protobuf:"bytes,6,rep,name=distinct_on" json:"distinct_on,omitempty"`
	// A starting point for the query results. Query cursors are
	// returned in query result batches.
	StartCursor []byte `protobuf:"bytes,7,opt,name=start_cursor,proto3" json:"start_cursor,omitempty"`
	// An ending point for the query results. Query cursors are
	// returned in query result batches.
	EndCursor []byte `protobuf:"bytes,8,opt,name=end_cursor,proto3" json:"end_cursor,omitempty"`
	// The number of results to skip. Applies before limit, but after all other
	// constraints. Optional. Must be >= 0 if specified.
	Offset int32 `protobuf:"varint,10,opt,name=offset" json:"offset,omitempty"`
	// The maximum number of results to return. Applies after all other
	// constraints. Optional.
	// Unspecified is interpreted as no limit.
	// Must be >= 0 if specified.
	Limit *google_protobuf2.Int32Value `protobuf:"bytes,12,opt,name=limit" json:"limit,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Query) GetProjection() []*Projection {
	if m != nil {
		return m.Projection
	}
	return nil
}

func (m *Query) GetKind() []*KindExpression {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Query) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *Query) GetOrder() []*PropertyOrder {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Query) GetDistinctOn() []*PropertyReference {
	if m != nil {
		return m.DistinctOn
	}
	return nil
}

func (m *Query) GetLimit() *google_protobuf2.Int32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

// A representation of a kind.
type KindExpression struct {
	// The name of the kind.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *KindExpression) Reset()                    { *m = KindExpression{} }
func (m *KindExpression) String() string            { return proto.CompactTextString(m) }
func (*KindExpression) ProtoMessage()               {}
func (*KindExpression) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// A reference to a property relative to the kind expressions.
type PropertyReference struct {
	// The name of the property.
	// If name includes "."s, it may be interpreted as a property name path.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PropertyReference) Reset()                    { *m = PropertyReference{} }
func (m *PropertyReference) String() string            { return proto.CompactTextString(m) }
func (*PropertyReference) ProtoMessage()               {}
func (*PropertyReference) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// A representation of a property in a projection.
type Projection struct {
	// The property to project.
	Property *PropertyReference `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
}

func (m *Projection) Reset()                    { *m = Projection{} }
func (m *Projection) String() string            { return proto.CompactTextString(m) }
func (*Projection) ProtoMessage()               {}
func (*Projection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Projection) GetProperty() *PropertyReference {
	if m != nil {
		return m.Property
	}
	return nil
}

// The desired order for a specific property.
type PropertyOrder struct {
	// The property to order by.
	Property *PropertyReference `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The direction to order by. Defaults to `ASCENDING`.
	Direction PropertyOrder_Direction `protobuf:"varint,2,opt,name=direction,enum=google.datastore.v1beta3.PropertyOrder_Direction" json:"direction,omitempty"`
}

func (m *PropertyOrder) Reset()                    { *m = PropertyOrder{} }
func (m *PropertyOrder) String() string            { return proto.CompactTextString(m) }
func (*PropertyOrder) ProtoMessage()               {}
func (*PropertyOrder) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *PropertyOrder) GetProperty() *PropertyReference {
	if m != nil {
		return m.Property
	}
	return nil
}

// A holder for any type of filter.
type Filter struct {
	// The type of filter.
	//
	// Types that are valid to be assigned to FilterType:
	//	*Filter_CompositeFilter
	//	*Filter_PropertyFilter
	FilterType isFilter_FilterType `protobuf_oneof:"filter_type"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type isFilter_FilterType interface {
	isFilter_FilterType()
}

type Filter_CompositeFilter struct {
	CompositeFilter *CompositeFilter `protobuf:"bytes,1,opt,name=composite_filter,oneof"`
}
type Filter_PropertyFilter struct {
	PropertyFilter *PropertyFilter `protobuf:"bytes,2,opt,name=property_filter,oneof"`
}

func (*Filter_CompositeFilter) isFilter_FilterType() {}
func (*Filter_PropertyFilter) isFilter_FilterType()  {}

func (m *Filter) GetFilterType() isFilter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return nil
}

func (m *Filter) GetCompositeFilter() *CompositeFilter {
	if x, ok := m.GetFilterType().(*Filter_CompositeFilter); ok {
		return x.CompositeFilter
	}
	return nil
}

func (m *Filter) GetPropertyFilter() *PropertyFilter {
	if x, ok := m.GetFilterType().(*Filter_PropertyFilter); ok {
		return x.PropertyFilter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Filter_OneofMarshaler, _Filter_OneofUnmarshaler, _Filter_OneofSizer, []interface{}{
		(*Filter_CompositeFilter)(nil),
		(*Filter_PropertyFilter)(nil),
	}
}

func _Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Filter)
	// filter_type
	switch x := m.FilterType.(type) {
	case *Filter_CompositeFilter:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CompositeFilter); err != nil {
			return err
		}
	case *Filter_PropertyFilter:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PropertyFilter); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Filter.FilterType has unexpected type %T", x)
	}
	return nil
}

func _Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Filter)
	switch tag {
	case 1: // filter_type.composite_filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CompositeFilter)
		err := b.DecodeMessage(msg)
		m.FilterType = &Filter_CompositeFilter{msg}
		return true, err
	case 2: // filter_type.property_filter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PropertyFilter)
		err := b.DecodeMessage(msg)
		m.FilterType = &Filter_PropertyFilter{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Filter)
	// filter_type
	switch x := m.FilterType.(type) {
	case *Filter_CompositeFilter:
		s := proto.Size(x.CompositeFilter)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Filter_PropertyFilter:
		s := proto.Size(x.PropertyFilter)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A filter that merges multiple other filters using the given operator.
type CompositeFilter struct {
	// The operator for combining multiple filters.
	Op CompositeFilter_Operator `protobuf:"varint,1,opt,name=op,enum=google.datastore.v1beta3.CompositeFilter_Operator" json:"op,omitempty"`
	// The list of filters to combine.
	// Must contain at least one filter.
	Filters []*Filter `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
}

func (m *CompositeFilter) Reset()                    { *m = CompositeFilter{} }
func (m *CompositeFilter) String() string            { return proto.CompactTextString(m) }
func (*CompositeFilter) ProtoMessage()               {}
func (*CompositeFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *CompositeFilter) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// A filter on a specific property.
type PropertyFilter struct {
	// The property to filter by.
	Property *PropertyReference `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	// The operator to filter by.
	Op PropertyFilter_Operator `protobuf:"varint,2,opt,name=op,enum=google.datastore.v1beta3.PropertyFilter_Operator" json:"op,omitempty"`
	// The value to compare the property to.
	Value *Value `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *PropertyFilter) Reset()                    { *m = PropertyFilter{} }
func (m *PropertyFilter) String() string            { return proto.CompactTextString(m) }
func (*PropertyFilter) ProtoMessage()               {}
func (*PropertyFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *PropertyFilter) GetProperty() *PropertyReference {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *PropertyFilter) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// A [GQL query](https://cloud.google.com/datastore/docs/apis/gql/gql_reference).
type GqlQuery struct {
	// A string of the format described
	// [here](https://cloud.google.com/datastore/docs/apis/gql/gql_reference).
	QueryString string `protobuf:"bytes,1,opt,name=query_string" json:"query_string,omitempty"`
	// When false, the query string must not contain any literals and instead
	// must bind all values. For example,
	// `SELECT * FROM Kind WHERE a = 'string literal'` is not allowed, while
	// `SELECT * FROM Kind WHERE a = @value` is.
	AllowLiterals bool `protobuf:"varint,2,opt,name=allow_literals" json:"allow_literals,omitempty"`
	// For each non-reserved named binding site in the query string,
	// there must be a named parameter with that name,
	// but not necessarily the inverse.
	// Key must match regex `[A-Za-z_$][A-Za-z_$0-9]*`, must not match regex
	// `__.*__`, and must not be `""`.
	NamedBindings map[string]*GqlQueryParameter `protobuf:"bytes,5,rep,name=named_bindings" json:"named_bindings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Numbered binding site @1 references the first numbered parameter,
	// effectively using 1-based indexing, rather than the usual 0.
	// For each binding site numbered i in `query_string`,
	// there must be an i-th numbered parameter.
	// The inverse must also be true.
	PositionalBindings []*GqlQueryParameter `protobuf:"bytes,4,rep,name=positional_bindings" json:"positional_bindings,omitempty"`
}

func (m *GqlQuery) Reset()                    { *m = GqlQuery{} }
func (m *GqlQuery) String() string            { return proto.CompactTextString(m) }
func (*GqlQuery) ProtoMessage()               {}
func (*GqlQuery) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *GqlQuery) GetNamedBindings() map[string]*GqlQueryParameter {
	if m != nil {
		return m.NamedBindings
	}
	return nil
}

func (m *GqlQuery) GetPositionalBindings() []*GqlQueryParameter {
	if m != nil {
		return m.PositionalBindings
	}
	return nil
}

// A binding parameter for a GQL query.
type GqlQueryParameter struct {
	// The type of parameter.
	//
	// Types that are valid to be assigned to ParameterType:
	//	*GqlQueryParameter_Value
	//	*GqlQueryParameter_Cursor
	ParameterType isGqlQueryParameter_ParameterType `protobuf_oneof:"parameter_type"`
}

func (m *GqlQueryParameter) Reset()                    { *m = GqlQueryParameter{} }
func (m *GqlQueryParameter) String() string            { return proto.CompactTextString(m) }
func (*GqlQueryParameter) ProtoMessage()               {}
func (*GqlQueryParameter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

type isGqlQueryParameter_ParameterType interface {
	isGqlQueryParameter_ParameterType()
}

type GqlQueryParameter_Value struct {
	Value *Value `protobuf:"bytes,2,opt,name=value,oneof"`
}
type GqlQueryParameter_Cursor struct {
	Cursor []byte `protobuf:"bytes,3,opt,name=cursor,proto3,oneof"`
}

func (*GqlQueryParameter_Value) isGqlQueryParameter_ParameterType()  {}
func (*GqlQueryParameter_Cursor) isGqlQueryParameter_ParameterType() {}

func (m *GqlQueryParameter) GetParameterType() isGqlQueryParameter_ParameterType {
	if m != nil {
		return m.ParameterType
	}
	return nil
}

func (m *GqlQueryParameter) GetValue() *Value {
	if x, ok := m.GetParameterType().(*GqlQueryParameter_Value); ok {
		return x.Value
	}
	return nil
}

func (m *GqlQueryParameter) GetCursor() []byte {
	if x, ok := m.GetParameterType().(*GqlQueryParameter_Cursor); ok {
		return x.Cursor
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GqlQueryParameter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GqlQueryParameter_OneofMarshaler, _GqlQueryParameter_OneofUnmarshaler, _GqlQueryParameter_OneofSizer, []interface{}{
		(*GqlQueryParameter_Value)(nil),
		(*GqlQueryParameter_Cursor)(nil),
	}
}

func _GqlQueryParameter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GqlQueryParameter)
	// parameter_type
	switch x := m.ParameterType.(type) {
	case *GqlQueryParameter_Value:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Value); err != nil {
			return err
		}
	case *GqlQueryParameter_Cursor:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Cursor)
	case nil:
	default:
		return fmt.Errorf("GqlQueryParameter.ParameterType has unexpected type %T", x)
	}
	return nil
}

func _GqlQueryParameter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GqlQueryParameter)
	switch tag {
	case 2: // parameter_type.value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Value)
		err := b.DecodeMessage(msg)
		m.ParameterType = &GqlQueryParameter_Value{msg}
		return true, err
	case 3: // parameter_type.cursor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ParameterType = &GqlQueryParameter_Cursor{x}
		return true, err
	default:
		return false, nil
	}
}

func _GqlQueryParameter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GqlQueryParameter)
	// parameter_type
	switch x := m.ParameterType.(type) {
	case *GqlQueryParameter_Value:
		s := proto.Size(x.Value)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GqlQueryParameter_Cursor:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Cursor)))
		n += len(x.Cursor)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A batch of results produced by a query.
type QueryResultBatch struct {
	// The number of results skipped, typically because of an offset.
	SkippedResults int32 `protobuf:"varint,6,opt,name=skipped_results" json:"skipped_results,omitempty"`
	// A cursor that points to the position after the last skipped result.
	// Will be set when `skipped_results` != 0.
	SkippedCursor []byte `protobuf:"bytes,3,opt,name=skipped_cursor,proto3" json:"skipped_cursor,omitempty"`
	// The result type for every entity in `entity_results`.
	EntityResultType EntityResult_ResultType `protobuf:"varint,1,opt,name=entity_result_type,enum=google.datastore.v1beta3.EntityResult_ResultType" json:"entity_result_type,omitempty"`
	// The results for this batch.
	EntityResults []*EntityResult `protobuf:"bytes,2,rep,name=entity_results" json:"entity_results,omitempty"`
	// A cursor that points to the position after the last result in the batch.
	EndCursor []byte `protobuf:"bytes,4,opt,name=end_cursor,proto3" json:"end_cursor,omitempty"`
	// The state of the query after the current batch.
	MoreResults QueryResultBatch_MoreResultsType `protobuf:"varint,5,opt,name=more_results,enum=google.datastore.v1beta3.QueryResultBatch_MoreResultsType" json:"more_results,omitempty"`
}

func (m *QueryResultBatch) Reset()                    { *m = QueryResultBatch{} }
func (m *QueryResultBatch) String() string            { return proto.CompactTextString(m) }
func (*QueryResultBatch) ProtoMessage()               {}
func (*QueryResultBatch) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *QueryResultBatch) GetEntityResults() []*EntityResult {
	if m != nil {
		return m.EntityResults
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityResult)(nil), "google.datastore.v1beta3.EntityResult")
	proto.RegisterType((*Query)(nil), "google.datastore.v1beta3.Query")
	proto.RegisterType((*KindExpression)(nil), "google.datastore.v1beta3.KindExpression")
	proto.RegisterType((*PropertyReference)(nil), "google.datastore.v1beta3.PropertyReference")
	proto.RegisterType((*Projection)(nil), "google.datastore.v1beta3.Projection")
	proto.RegisterType((*PropertyOrder)(nil), "google.datastore.v1beta3.PropertyOrder")
	proto.RegisterType((*Filter)(nil), "google.datastore.v1beta3.Filter")
	proto.RegisterType((*CompositeFilter)(nil), "google.datastore.v1beta3.CompositeFilter")
	proto.RegisterType((*PropertyFilter)(nil), "google.datastore.v1beta3.PropertyFilter")
	proto.RegisterType((*GqlQuery)(nil), "google.datastore.v1beta3.GqlQuery")
	proto.RegisterType((*GqlQueryParameter)(nil), "google.datastore.v1beta3.GqlQueryParameter")
	proto.RegisterType((*QueryResultBatch)(nil), "google.datastore.v1beta3.QueryResultBatch")
	proto.RegisterEnum("google.datastore.v1beta3.EntityResult_ResultType", EntityResult_ResultType_name, EntityResult_ResultType_value)
	proto.RegisterEnum("google.datastore.v1beta3.PropertyOrder_Direction", PropertyOrder_Direction_name, PropertyOrder_Direction_value)
	proto.RegisterEnum("google.datastore.v1beta3.CompositeFilter_Operator", CompositeFilter_Operator_name, CompositeFilter_Operator_value)
	proto.RegisterEnum("google.datastore.v1beta3.PropertyFilter_Operator", PropertyFilter_Operator_name, PropertyFilter_Operator_value)
	proto.RegisterEnum("google.datastore.v1beta3.QueryResultBatch_MoreResultsType", QueryResultBatch_MoreResultsType_name, QueryResultBatch_MoreResultsType_value)
}

var fileDescriptor2 = []byte{
	// 1085 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x5f, 0x53, 0xe3, 0x54,
	0x14, 0xdf, 0xf4, 0x0f, 0x94, 0xd3, 0x52, 0xb2, 0x77, 0x57, 0xb7, 0xcb, 0xba, 0x2b, 0x9b, 0x71,
	0x14, 0x75, 0x4c, 0xa5, 0x8c, 0xce, 0x0e, 0x23, 0x3b, 0x96, 0x36, 0x40, 0xa5, 0x24, 0x25, 0x29,
	0xce, 0xf0, 0x94, 0x09, 0xed, 0xa5, 0x46, 0x42, 0xd2, 0x4d, 0x6e, 0x77, 0xe5, 0x33, 0xf8, 0xec,
	0x8c, 0xef, 0x7e, 0x04, 0xfd, 0x26, 0xce, 0xf8, 0xe2, 0x97, 0xf1, 0xe4, 0xde, 0xa4, 0xd0, 0x42,
	0x2d, 0xcc, 0xbe, 0x25, 0xf7, 0x9e, 0xdf, 0xef, 0xfc, 0x3f, 0xe7, 0xc2, 0x77, 0x83, 0x20, 0x18,
	0x78, 0x54, 0x1d, 0x04, 0x9e, 0xe3, 0x0f, 0xd4, 0x20, 0x1c, 0x54, 0x7b, 0x5e, 0x30, 0xea, 0x57,
	0xfb, 0x0e, 0x73, 0x22, 0x16, 0x84, 0xb4, 0xea, 0xfa, 0x8c, 0x86, 0xbe, 0xe3, 0x55, 0x87, 0x61,
	0xc0, 0x82, 0xea, 0x9b, 0x11, 0x0d, 0x2f, 0x55, 0xfe, 0x4d, 0x2a, 0x09, 0x7a, 0x2c, 0xac, 0xbe,
	0xdd, 0x38, 0xa5, 0xcc, 0xd9, 0x5c, 0xdd, 0xbe, 0x37, 0x2f, 0xf5, 0x99, 0xcb, 0x12, 0xe2, 0xd5,
	0x17, 0x02, 0x2e, 0xae, 0x4e, 0x47, 0x67, 0xd5, 0x77, 0xa1, 0x33, 0x1c, 0xd2, 0x30, 0x4a, 0xee,
	0x77, 0xee, 0x41, 0xcf, 0x2e, 0x87, 0xd4, 0x16, 0x3a, 0x3c, 0x87, 0x79, 0x28, 0xce, 0x7f, 0x94,
	0x3f, 0x24, 0x28, 0x69, 0x5c, 0xa9, 0x49, 0xa3, 0x91, 0xc7, 0xc8, 0xd7, 0xb0, 0x20, 0x8c, 0xa8,
	0x48, 0x6b, 0xd2, 0x7a, 0xb1, 0xb6, 0xa6, 0xce, 0x72, 0x4f, 0x15, 0x38, 0x52, 0x86, 0x85, 0xde,
	0x28, 0x8c, 0x82, 0xb0, 0x92, 0x45, 0x44, 0x49, 0x39, 0x02, 0x10, 0x5c, 0x5d, 0xd4, 0x49, 0x9e,
	0xc1, 0x13, 0x53, 0xb3, 0x8e, 0xdb, 0x5d, 0xbb, 0x7b, 0xd2, 0xd1, 0xec, 0x63, 0xdd, 0xea, 0x68,
	0x8d, 0xd6, 0x6e, 0x4b, 0x6b, 0xca, 0x0f, 0x48, 0x01, 0x72, 0xbb, 0xc7, 0xed, 0xb6, 0x2c, 0x21,
	0x09, 0x74, 0x4c, 0xe3, 0x07, 0xad, 0xd1, 0x6d, 0x19, 0xba, 0x9c, 0x21, 0x25, 0x28, 0x1c, 0x68,
	0x27, 0xb6, 0xa1, 0xb7, 0x4f, 0xe4, 0xac, 0xf2, 0x6b, 0x16, 0xf2, 0x47, 0x71, 0xc8, 0xc9, 0x2b,
	0x00, 0x34, 0xfc, 0x67, 0xda, 0x63, 0x6e, 0xe0, 0x57, 0x32, 0x6b, 0x59, 0x34, 0xf1, 0x93, 0xd9,
	0x26, 0x76, 0xc6, 0xb2, 0xe4, 0x5b, 0xc8, 0x9d, 0xbb, 0x7e, 0x1f, 0x8d, 0x8c, 0x31, 0xeb, 0xb3,
	0x31, 0x07, 0x28, 0xa5, 0xfd, 0x32, 0x0c, 0x69, 0x14, 0xc5, 0x38, 0x0c, 0xc8, 0x99, 0xeb, 0x61,
	0x18, 0x2b, 0xb9, 0x79, 0x01, 0xd9, 0xe5, 0x72, 0xa8, 0x29, 0x1f, 0x84, 0x7d, 0x04, 0xe4, 0xb9,
	0xaa, 0xcf, 0xfe, 0xd7, 0x3c, 0xcc, 0x27, 0xbb, 0x34, 0x62, 0x71, 0xf2, 0x3d, 0x14, 0xfb, 0x6e,
	0xc4, 0x5c, 0xbf, 0xc7, 0x6c, 0x74, 0x6e, 0x81, 0xa3, 0xbf, 0x9c, 0x8f, 0x36, 0xe9, 0x19, 0x0d,
	0xa9, 0xdf, 0xa3, 0xe4, 0x31, 0x94, 0x22, 0xe6, 0x84, 0xcc, 0x4e, 0x12, 0xb2, 0x18, 0x27, 0x84,
	0x10, 0x00, 0xea, 0xf7, 0xd3, 0xb3, 0x02, 0x3f, 0xc3, 0xa4, 0x05, 0x67, 0x67, 0x11, 0x65, 0x15,
	0xc0, 0xff, 0x3c, 0xf9, 0x02, 0xf2, 0x9e, 0x7b, 0xe1, 0xb2, 0x4a, 0x89, 0x3b, 0xf9, 0x2c, 0xd5,
	0x9a, 0xd6, 0x9e, 0xda, 0xf2, 0xd9, 0x66, 0xed, 0x47, 0xc7, 0x1b, 0x51, 0xe5, 0x05, 0x94, 0xa7,
	0x62, 0x54, 0x82, 0x9c, 0xef, 0x5c, 0x50, 0x5e, 0x32, 0x4b, 0xca, 0x4b, 0x78, 0x78, 0xd3, 0xb4,
	0x54, 0x24, 0xc3, 0x45, 0x0e, 0x30, 0xdd, 0x57, 0xa9, 0xd9, 0x86, 0xc2, 0x30, 0x01, 0x24, 0x55,
	0x77, 0x1f, 0xaf, 0x95, 0x7f, 0x25, 0x58, 0x9e, 0x8c, 0xe4, 0xfb, 0x11, 0x92, 0x26, 0x2c, 0xf5,
	0xdd, 0x70, 0x5c, 0x63, 0xd2, 0x7a, 0xb9, 0xb6, 0x71, 0xc7, 0x24, 0xaa, 0xcd, 0x14, 0xa8, 0x68,
	0xb0, 0x34, 0xfe, 0x21, 0x4f, 0xe1, 0x83, 0x66, 0xcb, 0x14, 0xe5, 0x3d, 0xd5, 0x04, 0xcb, 0xb0,
	0x54, 0xb7, 0x1a, 0x9a, 0xde, 0x6c, 0xe9, 0x7b, 0xa2, 0x13, 0x9a, 0xda, 0xf8, 0x3f, 0x13, 0x77,
	0xe8, 0x42, 0x52, 0x58, 0x1a, 0xc8, 0xbd, 0xe0, 0x62, 0x18, 0x44, 0x2e, 0xa3, 0x76, 0x52, 0x94,
	0xc2, 0xbd, 0xcf, 0x67, 0x9b, 0xd7, 0x48, 0x11, 0x82, 0x64, 0xff, 0x01, 0x69, 0xc0, 0x4a, 0x1a,
	0x9d, 0x94, 0x25, 0xc3, 0x59, 0xd6, 0xe7, 0x3b, 0x99, 0x92, 0xec, 0x2c, 0x43, 0x51, 0x60, 0xed,
	0x78, 0xb4, 0x28, 0x7f, 0x4a, 0xb0, 0x32, 0xa5, 0x89, 0xbc, 0x86, 0x4c, 0x30, 0xe4, 0x06, 0x96,
	0x6b, 0xb5, 0x3b, 0x1b, 0xa8, 0x1a, 0xa8, 0xc8, 0x41, 0x09, 0xb2, 0x01, 0x8b, 0x42, 0x45, 0x94,
	0x34, 0xfa, 0xdc, 0xd6, 0x53, 0xbe, 0x82, 0xc2, 0x18, 0x5e, 0x81, 0xc7, 0x46, 0x47, 0x33, 0xeb,
	0x5d, 0xc3, 0x9c, 0x8a, 0xf8, 0x22, 0x64, 0xeb, 0x7a, 0x53, 0x96, 0x94, 0x7f, 0x32, 0x50, 0x9e,
	0xf4, 0xec, 0x7d, 0x4b, 0x67, 0x9b, 0xfb, 0x7c, 0xe7, 0x9a, 0x99, 0x76, 0x59, 0x85, 0xfc, 0xdb,
	0xb8, 0xc7, 0xf8, 0x28, 0x2d, 0xd6, 0x3e, 0x9e, 0xcd, 0x20, 0x5a, 0xf1, 0x37, 0xe9, 0x4e, 0x0e,
	0x63, 0x89, 0xb5, 0x35, 0xcb, 0xb2, 0xbb, 0xfb, 0x75, 0x1d, 0x4b, 0xec, 0x43, 0x20, 0xe3, 0x5f,
	0x1b, 0x85, 0xb5, 0xa3, 0xe3, 0x7a, 0x1b, 0x87, 0xae, 0x0c, 0xa5, 0x3d, 0x53, 0xab, 0x77, 0x35,
	0x53, 0x48, 0x66, 0xe3, 0xb2, 0xbd, 0x7e, 0x72, 0x25, 0x9c, 0x23, 0x4b, 0x90, 0x17, 0x9f, 0xf9,
	0x18, 0xb7, 0x5f, 0xb7, 0xec, 0xba, 0xde, 0xd0, 0x2c, 0x54, 0x2e, 0x17, 0x95, 0xbf, 0x32, 0x50,
	0xd8, 0x7b, 0xe3, 0x89, 0x99, 0x8d, 0x53, 0x89, 0xef, 0x4b, 0x3b, 0x62, 0xa1, 0xeb, 0x0f, 0xc4,
	0x94, 0x40, 0x23, 0xca, 0x8e, 0xe7, 0x05, 0xef, 0x6c, 0x0f, 0x53, 0x1f, 0x3a, 0x5e, 0xc4, 0xa3,
	0x56, 0x20, 0x87, 0x50, 0x8e, 0x07, 0x45, 0xdf, 0x3e, 0xc5, 0x19, 0x83, 0xe2, 0x51, 0x32, 0x46,
	0xbf, 0x99, 0x1d, 0x8b, 0x54, 0x93, 0xaa, 0xc7, 0xc0, 0x9d, 0x04, 0x87, 0xeb, 0x09, 0x95, 0xef,
	0xc3, 0x23, 0x5e, 0x5d, 0xd8, 0x84, 0x8e, 0x77, 0xc5, 0x99, 0x9b, 0x37, 0x5c, 0x53, 0xce, 0x8e,
	0x13, 0x22, 0x29, 0x5a, 0xb8, 0xea, 0x00, 0xb9, 0x85, 0xbf, 0x08, 0xd9, 0x73, 0x7a, 0x99, 0xf8,
	0xb4, 0x95, 0xa6, 0x2f, 0x33, 0xaf, 0x72, 0x6e, 0xd0, 0x6f, 0x65, 0x5e, 0x49, 0x8a, 0x0b, 0x0f,
	0x6f, 0x5c, 0xe0, 0x02, 0x9a, 0x20, 0x9d, 0x57, 0x13, 0xd8, 0xe0, 0xf2, 0xe4, 0x46, 0xc6, 0x6e,
	0x95, 0xa1, 0x3c, 0x4c, 0x09, 0x45, 0xc3, 0xfe, 0x9d, 0x05, 0x99, 0x2b, 0x12, 0xbb, 0x7a, 0xc7,
	0x61, 0xbd, 0x9f, 0xc8, 0x13, 0x58, 0x89, 0xce, 0x5d, 0x7c, 0x63, 0xf4, 0xed, 0x90, 0x1f, 0x47,
	0xb8, 0x85, 0xe2, 0xf5, 0x80, 0xc9, 0x4a, 0x2f, 0xae, 0x33, 0x63, 0xb2, 0x88, 0x78, 0x2d, 0x24,
	0xf2, 0x9c, 0x3b, 0x69, 0xf9, 0x8d, 0x79, 0x2f, 0x07, 0xa1, 0x59, 0xbd, 0xf6, 0x58, 0x78, 0x0d,
	0xe5, 0x09, 0xba, 0xb4, 0xf1, 0x3f, 0xbd, 0x1b, 0xd5, 0xd4, 0xa6, 0xcb, 0x71, 0x13, 0x3b, 0x50,
	0xba, 0x40, 0xc0, 0x98, 0x31, 0xcf, 0x8d, 0xdb, 0x9a, 0xcd, 0x38, 0x1d, 0x15, 0xf5, 0x10, 0x6f,
	0xc5, 0x7f, 0x14, 0x5b, 0xa9, 0xfc, 0x8e, 0xb3, 0x6e, 0xea, 0x8c, 0xbc, 0x84, 0xe7, 0x87, 0x86,
	0xa9, 0xd9, 0xe2, 0xad, 0x63, 0xdd, 0xf6, 0xd8, 0xc1, 0x2e, 0xd1, 0x8d, 0xae, 0xbd, 0xdb, 0xd2,
	0x5b, 0xd6, 0x3e, 0x9e, 0x48, 0xe4, 0x23, 0xa8, 0x4c, 0x80, 0xea, 0xbb, 0x71, 0xa3, 0xb5, 0x5b,
	0x87, 0xad, 0x2e, 0x76, 0xe3, 0x73, 0x78, 0x7a, 0xcb, 0x6d, 0xe3, 0xd8, 0xb4, 0xb0, 0xc5, 0x72,
	0xe4, 0x11, 0xac, 0xe8, 0x86, 0x7d, 0x5d, 0x42, 0xce, 0x9e, 0x2e, 0xf0, 0x7d, 0xbd, 0xf9, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x69, 0x9f, 0x5c, 0xd2, 0x0a, 0x00, 0x00,
}
