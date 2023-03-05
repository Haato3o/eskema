package syntax

type Primitive int

const (
	String = iota
	Char
	UInt8
	UInt16
	UInt32
	UInt64
	Int8
	Int16
	Int32
	Int64
	Float
	Double
	TimeStamp
	Date
	DateTime
	Array
	Map
	Bool
)
