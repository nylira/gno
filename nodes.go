package gno

import (
	"fmt"
	"reflect"
	"strconv"
)

//----------------------------------------
// Primitives

type Word int

const (
	// Special words
	ILLEGAL Word = iota

	// Names and basic type literals
	// (these words stand for classes of literals)
	NAME   // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"

	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	BAND     // &
	BOR      // |
	XOR      // ^
	SHL      // <<
	SHR      // >>
	BAND_NOT // &^

	ADD_ASSIGN      // +=
	SUB_ASSIGN      // -=
	MUL_ASSIGN      // *=
	QUO_ASSIGN      // /=
	REM_ASSIGN      // %=
	BAND_ASSIGN     // &=
	BOR_ASSIGN      // |=
	XOR_ASSIGN      // ^=
	SHL_ASSIGN      // <<=
	SHR_ASSIGN      // >>=
	BAND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ    // !=
	LEQ    // <=
	GEQ    // >=
	DEFINE // :=

	// Keywords
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR
)

type Name string

// Lets Name be embedded and become "nameful".
func (n Name) GetName() Name {
	return n
}

//----------------------------------------
// Attributes
// All nodes have attributes for general analysis purposes.

type Attributes struct {
	Data map[interface{}]interface{}
}

func (a *Attributes) GetAttribute(key interface{}) interface{} {
	return a.Data[key]
}

func (a *Attributes) SetAttribute(key interface{}, value interface{}) {
	if a.Data == nil {
		a.Data = make(map[interface{}]interface{})
	}
	a.Data[key] = value
}

//----------------------------------------
// Node

type Node interface {
	assertNode()
	String() string
	Copy() Node
	GetAttribute(key interface{}) interface{}
	SetAttribute(key interface{}, value interface{})
}

// non-pointer receiver to help make immutable.
func (_ *NameExpr) assertNode()          {}
func (_ *BasicLitExpr) assertNode()      {}
func (_ *BinaryExpr) assertNode()        {}
func (_ *CallExpr) assertNode()          {}
func (_ *IndexExpr) assertNode()         {}
func (_ *SelectorExpr) assertNode()      {}
func (_ *SliceExpr) assertNode()         {}
func (_ *StarExpr) assertNode()          {}
func (_ *RefExpr) assertNode()           {}
func (_ *TypeAssertExpr) assertNode()    {}
func (_ *UnaryExpr) assertNode()         {}
func (_ *CompositeLitExpr) assertNode()  {}
func (_ *KeyValueExpr) assertNode()      {}
func (_ *FuncLitExpr) assertNode()       {}
func (_ *constExpr) assertNode()         {}
func (_ *FieldTypeExpr) assertNode()     {}
func (_ *ArrayTypeExpr) assertNode()     {}
func (_ *SliceTypeExpr) assertNode()     {}
func (_ *InterfaceTypeExpr) assertNode() {}
func (_ *ChanTypeExpr) assertNode()      {}
func (_ *FuncTypeExpr) assertNode()      {}
func (_ *MapTypeExpr) assertNode()       {}
func (_ *StructTypeExpr) assertNode()    {}
func (_ *constTypeExpr) assertNode()     {}
func (_ *AssignStmt) assertNode()        {}
func (_ *BlockStmt) assertNode()         {}
func (_ *BranchStmt) assertNode()        {}
func (_ *DeclStmt) assertNode()          {}
func (_ *DeferStmt) assertNode()         {}
func (_ *ExprStmt) assertNode()          {}
func (_ *ForStmt) assertNode()           {}
func (_ *GoStmt) assertNode()            {}
func (_ *IfStmt) assertNode()            {}
func (_ *IncDecStmt) assertNode()        {}
func (_ *LabeledStmt) assertNode()       {}
func (_ *RangeStmt) assertNode()         {}
func (_ *ReturnStmt) assertNode()        {}
func (_ *SelectStmt) assertNode()        {}
func (_ *SelectCaseStmt) assertNode()    {}
func (_ *SendStmt) assertNode()          {}
func (_ *SwitchStmt) assertNode()        {}
func (_ *SwitchCaseStmt) assertNode()    {}
func (_ *EmptyStmt) assertNode()         {}
func (_ *loopStmt) assertNode()          {}
func (_ *FuncDecl) assertNode()          {}
func (_ *ImportDecl) assertNode()        {}
func (_ *ValueDecl) assertNode()         {}
func (_ *TypeDecl) assertNode()          {}
func (_ *FileNode) assertNode()          {}
func (_ *PackageNode) assertNode()       {}

var _ = &NameExpr{}
var _ = &BasicLitExpr{}
var _ = &BinaryExpr{}
var _ = &CallExpr{}
var _ = &IndexExpr{}
var _ = &SelectorExpr{}
var _ = &SliceExpr{}
var _ = &StarExpr{}
var _ = &RefExpr{}
var _ = &TypeAssertExpr{}
var _ = &UnaryExpr{}
var _ = &CompositeLitExpr{}
var _ = &KeyValueExpr{}
var _ = &FuncLitExpr{}
var _ = &constExpr{}
var _ = &FieldTypeExpr{}
var _ = &ArrayTypeExpr{}
var _ = &SliceTypeExpr{}
var _ = &InterfaceTypeExpr{}
var _ = &ChanTypeExpr{}
var _ = &FuncTypeExpr{}
var _ = &MapTypeExpr{}
var _ = &StructTypeExpr{}
var _ = &constTypeExpr{}
var _ = &AssignStmt{}
var _ = &BlockStmt{}
var _ = &BranchStmt{}
var _ = &DeclStmt{}
var _ = &DeferStmt{}
var _ = &ExprStmt{}
var _ = &ForStmt{}
var _ = &GoStmt{}
var _ = &IfStmt{}
var _ = &IncDecStmt{}
var _ = &LabeledStmt{}
var _ = &RangeStmt{}
var _ = &ReturnStmt{}
var _ = &SelectStmt{}
var _ = &SelectCaseStmt{}
var _ = &SendStmt{}
var _ = &SwitchStmt{}
var _ = &SwitchCaseStmt{}
var _ = &EmptyStmt{}
var _ = &loopStmt{}
var _ = &FuncDecl{}
var _ = &ImportDecl{}
var _ = &ValueDecl{}
var _ = &TypeDecl{}
var _ = &FileNode{}
var _ = &PackageNode{}

//----------------------------------------
// Expr
//
// expressions generally have no side effects on the caller's context,
// except for channel blocks, type assertions, and panics.

type Expr interface {
	Node
	assertExpr()
}

type Exprs []Expr

// non-pointer receiver to help make immutable.
func (*NameExpr) assertExpr()         {}
func (*BasicLitExpr) assertExpr()     {}
func (*BinaryExpr) assertExpr()       {}
func (*CallExpr) assertExpr()         {}
func (*IndexExpr) assertExpr()        {}
func (*SelectorExpr) assertExpr()     {}
func (*SliceExpr) assertExpr()        {}
func (*StarExpr) assertExpr()         {}
func (*RefExpr) assertExpr()          {}
func (*TypeAssertExpr) assertExpr()   {}
func (*UnaryExpr) assertExpr()        {}
func (*CompositeLitExpr) assertExpr() {}
func (*KeyValueExpr) assertExpr()     {}
func (*FuncLitExpr) assertExpr()      {}
func (*constExpr) assertExpr()        {}

var _ Expr = &NameExpr{}
var _ Expr = &BasicLitExpr{}
var _ Expr = &BinaryExpr{}
var _ Expr = &CallExpr{}
var _ Expr = &IndexExpr{}
var _ Expr = &SelectorExpr{}
var _ Expr = &SliceExpr{}
var _ Expr = &StarExpr{}
var _ Expr = &RefExpr{}
var _ Expr = &TypeAssertExpr{}
var _ Expr = &UnaryExpr{}
var _ Expr = &CompositeLitExpr{}
var _ Expr = &KeyValueExpr{}
var _ Expr = &FuncLitExpr{}
var _ Expr = &constExpr{}

type NameExpr struct {
	Attributes
	// TODO rename .Path's to .ValuePaths.
	Path ValuePath // set by preprocessor.
	Name
}

type BasicLitExpr struct {
	Attributes
	// INT, FLOAT, IMAG, CHAR, or STRING
	Kind Word
	// literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo"
	// or `\m\n\o`
	Value string
}

type BinaryExpr struct { // (Left Op Right)
	Attributes
	Left  Expr // left operand
	Op    Word // operator
	Right Expr // right operand
}

type CallExpr struct { // Func(Args<Varg?...>)
	Attributes
	Func Expr  // function expression
	Args Exprs // function arguments, if any.
	Varg bool  // if true, final arg is variadic.
}

type IndexExpr struct { // X[Index]
	Attributes
	X     Expr // expression
	Index Expr // index expression
}

type SelectorExpr struct { // X.Sel
	Attributes
	X    Expr      // expression
	Path ValuePath // set by preprocessor.
	Sel  Name      // field selector
}

type SliceExpr struct { // X[Low:High:Max]
	Attributes
	X    Expr // expression
	Low  Expr // begin of slice range; or nil
	High Expr // end of slice range; or nil
	Max  Expr // maximum capacity of slice; or nil; added in Go 1.2
}

// A StarExpr node represents an expression of the form
// "*" Expression.  Semantically it could be a unary "*"
// expression, or a pointer type.
type StarExpr struct { // *X
	Attributes
	X Expr // operand
}

type RefExpr struct { // &X
	Attributes
	X Expr // operand
}

type TypeAssertExpr struct { // X.(Type)
	Attributes
	X     Expr // expression
	Type  Expr // asserted type; nil means type switch X.(type)
	HasOK bool // if true, is form: `_, ok := <X>.(<Type>)`
}

// A UnaryExpr node represents a unary expression. Unary
// "*" expressions (dereferencing and pointer-types) are
// represented with StarExpr nodes.  Unary & expressions
// (referencing) are represented with RefExpr nodes.
type UnaryExpr struct { // (Op X)
	Attributes
	X  Expr // operand
	Op Word // operator
}

// MyType{<key>:<value>} struct, array, slice, and map
// expressions.
type CompositeLitExpr struct {
	Attributes
	Type Expr          // literal type; or nil
	Elts KeyValueExprs // list of struct fields; if any
}

// A KeyValueExpr represents a single key-value pair in
// struct, array, slice, and map expressions.
type KeyValueExpr struct {
	Attributes
	Key   Expr // or nil
	Value Expr // never nil
}

type KeyValueExprs []KeyValueExpr

// A FuncLitExpr node represents a function literal.  Here one
// can reference statements from an expression, which
// completes the procedural circle.
type FuncLitExpr struct {
	Attributes
	StaticBlock
	Type FuncTypeExpr // function type
	Body Stmts        // function body
}

// The preprocessor replaces const expressions
// with *constExpr nodes.
type constExpr struct {
	Attributes
	Source Expr // (preprocessed) source of this value.
	TypedValue
}

//----------------------------------------
// Type(Expressions)
//
// In Go, Type expressions can be evaluated immediately
// without invoking the stack machine.  Exprs in type
// expressions are const (as in array len expr or map key type
// expr) or refer to an exposed symbol (with any pointer
// indirections).  this makes for more optimal performance.
//
// In Gno, type expressions are evaluated on the stack, with
// continuation opcodes, so the Gno VM could support types as
// first class objects.

type TypeExpr interface {
	Expr
	assertTypeExpr()
}

// non-pointer receiver to help make immutable.
func (_ *FieldTypeExpr) assertTypeExpr()     {}
func (_ *ArrayTypeExpr) assertTypeExpr()     {}
func (_ *SliceTypeExpr) assertTypeExpr()     {}
func (_ *InterfaceTypeExpr) assertTypeExpr() {}
func (_ *ChanTypeExpr) assertTypeExpr()      {}
func (_ *FuncTypeExpr) assertTypeExpr()      {}
func (_ *MapTypeExpr) assertTypeExpr()       {}
func (_ *StructTypeExpr) assertTypeExpr()    {}
func (_ *constTypeExpr) assertTypeExpr()     {}

func (_ *FieldTypeExpr) assertExpr()     {}
func (_ *ArrayTypeExpr) assertExpr()     {}
func (_ *SliceTypeExpr) assertExpr()     {}
func (_ *InterfaceTypeExpr) assertExpr() {}
func (_ *ChanTypeExpr) assertExpr()      {}
func (_ *FuncTypeExpr) assertExpr()      {}
func (_ *MapTypeExpr) assertExpr()       {}
func (_ *StructTypeExpr) assertExpr()    {}
func (_ *constTypeExpr) assertExpr()     {}

var _ TypeExpr = &FieldTypeExpr{}
var _ TypeExpr = &ArrayTypeExpr{}
var _ TypeExpr = &SliceTypeExpr{}
var _ TypeExpr = &InterfaceTypeExpr{}
var _ TypeExpr = &ChanTypeExpr{}
var _ TypeExpr = &FuncTypeExpr{}
var _ TypeExpr = &MapTypeExpr{}
var _ TypeExpr = &StructTypeExpr{}
var _ TypeExpr = &constTypeExpr{}

type FieldTypeExpr struct {
	Attributes
	Name
	Type Expr

	// Currently only BasicLitExpr allowed.
	// NOTE: In Go, only struct fields can have tags.
	Tag Expr
}

type FieldTypeExprs []FieldTypeExpr

type ArrayTypeExpr struct {
	Attributes
	Len Expr // if nil, variadic array lit
	Elt Expr // element type
}

type SliceTypeExpr struct {
	Attributes
	Elt Expr // element type
	Vrd bool // variadic arg expression
}

type InterfaceTypeExpr struct {
	Attributes
	Methods FieldTypeExprs // list of methods
}

type ChanDir int

const (
	SEND ChanDir = 1 << iota
	RECV
)
const (
	BOTH = SEND | RECV
)

type ChanTypeExpr struct {
	Attributes
	Dir   ChanDir // channel direction
	Value Expr    // value type
}

type FuncTypeExpr struct {
	Attributes
	Params  FieldTypeExprs // (incoming) parameters, if any.
	Results FieldTypeExprs // (outgoing) results, if any.
}

type MapTypeExpr struct {
	Attributes
	Key   Expr // const
	Value Expr // value type
}

type StructTypeExpr struct {
	Attributes
	Fields FieldTypeExprs // list of field declarations
}

// Like constExpr but for types.
type constTypeExpr struct {
	Attributes
	Source Expr
	Type
}

//----------------------------------------
// Stmt
//
// statements generally have side effects on the calling context.

type Stmt interface {
	Node
	assertStmt()
}

type Stmts []Stmt

// non-pointer receiver to help make immutable.
func (*AssignStmt) assertStmt()     {}
func (*BlockStmt) assertStmt()      {}
func (*BranchStmt) assertStmt()     {}
func (*DeclStmt) assertStmt()       {}
func (*DeferStmt) assertStmt()      {}
func (*EmptyStmt) assertStmt()      {} // useful for _ctif
func (*ExprStmt) assertStmt()       {}
func (*ForStmt) assertStmt()        {}
func (*GoStmt) assertStmt()         {}
func (*IfStmt) assertStmt()         {}
func (*IncDecStmt) assertStmt()     {}
func (*LabeledStmt) assertStmt()    {}
func (*RangeStmt) assertStmt()      {}
func (*ReturnStmt) assertStmt()     {}
func (*SelectStmt) assertStmt()     {}
func (*SelectCaseStmt) assertStmt() {}
func (*SendStmt) assertStmt()       {}
func (*SwitchStmt) assertStmt()     {}
func (*SwitchCaseStmt) assertStmt() {}
func (*loopStmt) assertStmt()       {}

var _ Stmt = &AssignStmt{}
var _ Stmt = &BlockStmt{}
var _ Stmt = &BranchStmt{}
var _ Stmt = &DeclStmt{}
var _ Stmt = &DeferStmt{}
var _ Stmt = &EmptyStmt{}
var _ Stmt = &ExprStmt{}
var _ Stmt = &ForStmt{}
var _ Stmt = &GoStmt{}
var _ Stmt = &IfStmt{}
var _ Stmt = &IncDecStmt{}
var _ Stmt = &LabeledStmt{}
var _ Stmt = &RangeStmt{}
var _ Stmt = &ReturnStmt{}
var _ Stmt = &SelectStmt{}
var _ Stmt = &SelectCaseStmt{}
var _ Stmt = &SendStmt{}
var _ Stmt = &SwitchStmt{}
var _ Stmt = &SwitchCaseStmt{}
var _ Stmt = &loopStmt{}

type AssignStmt struct {
	Attributes
	Lhs Exprs
	Op  Word // assignment word (DEFINE, ASSIGN)
	Rhs Exprs
}

type BlockStmt struct {
	Attributes
	StaticBlock
	Body Stmts
}

type BranchStmt struct {
	Attributes
	Op    Word // keyword word (BREAK, CONTINUE, GOTO, FALLTHROUGH)
	Label Name // label name; or empty
}

type DeclStmt struct {
	Attributes
	Decls Decls // (simple) ValueDecl or TypeDecl
}

type DeferStmt struct {
	Attributes
	Call CallExpr
}

// A compile artifact to use in place of nil.
// For example, _ctif() may return an empty statement.
type EmptyStmt struct {
	Attributes
}

type ExprStmt struct {
	Attributes
	X Expr
}

type ForStmt struct {
	Attributes
	StaticBlock
	Init Stmt // initialization (simple) statement; or nil
	Cond Expr // condition; or nil
	Post Stmt // post iteration (simple) statement; or nil
	Body Stmts
}

type GoStmt struct {
	Attributes
	Call CallExpr
}

// NOTE: syntactically, code may choose to chain if-else statements
// with `} else if ... {` constructions, but this is not represented
// in the logical AST.
type IfStmt struct {
	Attributes
	StaticBlock
	Init Stmt  // initialization (simple) statement; or nil
	Cond Expr  // condition; or nil
	Body Stmts // body statements
	Else Stmts // else statements
}

type IncDecStmt struct {
	Attributes
	X  Expr
	Op Word // INC or DEC
}

type LabeledStmt struct {
	Attributes
	Label Name
	Stmt  Stmt
}

type RangeStmt struct {
	Attributes
	StaticBlock
	X          Expr // value to range over
	Key, Value Expr // Key, Value may be nil
	Op         Word // ASSIGN or DEFINE
	Body       Stmts
}

type ReturnStmt struct {
	Attributes
	Results Exprs // result expressions; or nil
}

type SelectStmt struct {
	Attributes
	Cases []SelectCaseStmt
}

type SelectCaseStmt struct {
	Attributes
	StaticBlock
	Comm Stmt // send or receive statement; nil means default case
	Body Stmts
}

type SendStmt struct {
	Attributes
	Chan  Expr
	Value Expr
}

// type ReceiveStmt
// is just AssignStmt with a Receive unary expression.

type SwitchStmt struct {
	Attributes
	StaticBlock
	Init    Stmt             // initialization (simple) statement; or nil.
	X       Expr             // tag or _.(type) expression; or nil.
	Cases   []SwitchCaseStmt // cases
	VarName Name             // tag or type-switched value.
}

type SwitchCaseStmt struct {
	Attributes
	StaticBlock
	Cases Exprs // list of expressions or types; nil means default case
	Body  Stmts
}

//----------------------------------------
// Loop Statement (persistent)

type loopStmt struct {
	Attributes
	ForStmt   *ForStmt
	RangeStmt *RangeStmt
	ListLen   int  // for RangeStmt only
	ListIndex int  // for RangeStmt only
	BodyLen   int  // for continue
	BodyIndex int  // init:-2, cond/elem:-1, body:0..., post:n
	Active    Stmt // for PopStmt().
}

func (s *loopStmt) PopActiveStmt() (as Stmt) {
	as = s.Active
	s.Active = nil
	return
}

func (s *loopStmt) String() string {
	if s.ForStmt != nil {
		return fmt.Sprintf("loopStmt[%s %d]",
			s.ForStmt.String(),
			s.BodyIndex)
	} else {
		return fmt.Sprintf("loopStmt[%s %d/%d/%d]",
			s.RangeStmt.String(),
			s.ListLen,
			s.ListIndex,
			s.BodyIndex)
	}
}

//----------------------------------------
// Simple Statement
// NOTE: SimpleStmt is not used in nodes due to itable conversion costs.
//
// These are used in if, switch, and for statements for simple
// initialization.  The only allowed types are EmptyStmt, ExprStmt,
// SendStmt, IncDecStmt, and AssignStmt.

type SimpleStmt interface {
	Stmt
	assertSimpleStmt()
}

// non-pointer receiver to help make immutable.
func (*EmptyStmt) assertSimpleStmt()  {}
func (*ExprStmt) assertSimpleStmt()   {}
func (*SendStmt) assertSimpleStmt()   {}
func (*IncDecStmt) assertSimpleStmt() {}
func (*AssignStmt) assertSimpleStmt() {}

//----------------------------------------
// Decl

type Decl interface {
	Node
	GetName() Name
	assertDecl()
}

type Decls []Decl

// non-pointer receiver to help make immutable.
func (_ *FuncDecl) assertDecl()   {}
func (_ *ImportDecl) assertDecl() {}
func (_ *ValueDecl) assertDecl()  {}
func (_ *TypeDecl) assertDecl()   {}

var _ Decl = &FuncDecl{}
var _ Decl = &ImportDecl{}
var _ Decl = &ValueDecl{}
var _ Decl = &TypeDecl{}

type FuncDecl struct {
	Attributes
	StaticBlock
	NameExpr
	IsMethod bool
	Recv     FieldTypeExpr // receiver (if method); or empty (if function)
	Type     FuncTypeExpr  // function signature: parameters and results
	Body     Stmts         // function body; or empty for external (non-Go) function
}

type ImportDecl struct {
	Attributes
	NameExpr // local package name, or ".". required.
	PkgPath  string
}

type ValueDecl struct {
	Attributes
	NameExpr
	Type  Expr // value type; or nil
	Value Expr // initial value; or nil (unless const).
	Const bool
}

type TypeDecl struct {
	Attributes
	NameExpr
	Type    Expr // Name, SelectorExpr, StarExpr, or XxxTypes
	IsAlias bool // type alias since Go 1.9
}

//----------------------------------------
// SimpleDecl

type SimpleDecl interface {
	Decl
	assertSimpleDecl()
}

type SimpleDecls []SimpleDecl

func (_ *ValueDecl) assertSimpleDecl() {}
func (_ *TypeDecl) assertSimpleDecl()  {}

var _ SimpleDecl = &ValueDecl{}
var _ SimpleDecl = &TypeDecl{}

//----------------------------------------
// *FileSet

type FileSet struct {
	Files []*FileNode
}

func (fs *FileSet) AddFiles(fns ...*FileNode) {
	fs.Files = append(fs.Files, fns...)
}

func (fs *FileSet) GetFileByName(n Name) *FileNode {
	for _, fn := range fs.Files {
		if fn.Name == n {
			return fn
		}
	}
	return nil
}

// Returns a pointer to the file body decl (as well as
// the *FileNode which contains it) that declares n
// for the associated package with *FileSet.  Does not
// work for import decls which are for the file level.
// The file body decl can be replaced by reference
// assignment.
// TODO move to package?
func (fs *FileSet) GetDeclFor(n Name) (*FileNode, *Decl) {
	// XXX index to bound to linear time.
	for _, fn := range fs.Files {
		for i, dn := range fn.Body {
			if _, isImport := dn.(*ImportDecl); isImport {
				continue
			}
			if dn.GetName() == n {
				// found the decl that declares n.
				return fn, &fn.Body[i]
			}
		}
	}
	panic(fmt.Sprintf(
		"name %s not defined in fileset %v",
		n, fs))
}

//----------------------------------------
// FileNode, & PackageNode

type FileNode struct {
	Attributes
	StaticBlock
	Name
	PkgName Name
	Body    Decls
}

type PackageNode struct {
	Attributes
	StaticBlock
	PkgPath string
	PkgName Name
	*FileSet
}

func NewPackageNode(name Name, path string, fset *FileSet) *PackageNode {
	pn := &PackageNode{
		PkgPath: path,
		PkgName: Name(name),
		FileSet: fset,
	}
	pn.InitStaticBlock(pn, nil)
	return pn
}

func (pn *PackageNode) NewPackage(rlmr Realmer) *PackageValue {
	pv := &PackageValue{
		Block: Block{
			Source: pn,
		},
		PkgName: pn.PkgName,
		PkgPath: pn.PkgPath,
		FBlocks: make(map[Name]*Block),
	}
	if IsRealmPath(pn.PkgPath) {
		rlm := rlmr(pn.PkgPath)
		pv.SetRealm(rlm)
		rlm.pkg = pv // TODO
	}
	pn.UpdatePackage(pv)
	return pv
}

func (pn *PackageNode) UpdatePackage(pv *PackageValue) {
	if pv.Source != pn {
		panic("PackageNode.UpdatePackage() package mismatch")
	}
	pvl := len(pv.Values)
	pnl := len(pn.Values)
	if pvl < pnl {
		// XXX: deep copy heap values
		nvs := make([]TypedValue, pnl-pvl)
		copy(nvs, pn.Values[pvl:pnl])
		for _, nv := range nvs {
			if nv.IsUndefined() {
				continue
			}
			switch nv.T.Kind() {
			case FuncKind:
				// If package-level FuncLit function, value is nil,
				// and the closure will be set at run-time.
				if nv.V == nil {
					// nothing to do
				} else {
					// Set function closure for function declarations.
					switch fv := nv.V.(type) {
					case *FuncValue:
						// set fv.pkg.
						fv.pkg = pv
						// set fv.Closure.
						if fv.Closure != nil {
							panic("expected nil closure for static func")
						}
						if fv.FileName == "" {
							// Allow
							// m.RunDeclaration(FuncD(...))
							// without any file nodes, as long
							// as it uses no imports.
							continue
						}
						fb := pv.FBlocks[fv.FileName]
						fv.Closure = fb
					case *nativeValue:
						// do nothing for go native functions.
					default:
						panic("should not happen")
					}
				}
			case TypeKind:
				nt := nv.GetType()
				if dt, ok := nt.(*DeclaredType); ok {
					for i := 0; i < len(dt.Methods); i++ {
						mv := dt.Methods[i].V.(*FuncValue)
						// set mv.pkg.
						mv.pkg = pv
						// set mv.Closure.
						if mv.Closure != nil {
							panic("expected nil closure for static method")
						}
						fn, _ := pn.GetDeclFor(dt.Name)
						fb := pv.FBlocks[fn.Name]
						mv.Closure = fb
					}
				}
			default:
				// already shallowed copied.
			}
		}
		pv.Values = append(pv.Values, nvs...)
	} else if pvl > pnl {
		panic("package size error")
	} else {
		// nothing to do
	}
}

//----------------------------------------
// BlockNode

// Nodes that create their own scope satisfy this interface.
type BlockNode interface {
	Node
	InitStaticBlock(BlockNode, BlockNode)
	GetStaticBlock() *StaticBlock

	// StaticBlock promoted methods
	GetNames() []Name
	GetNumNames() uint16
	GetParent() BlockNode
	GetPathForName(Name) ValuePath
	GetLocalIndex(Name) (uint16, bool)
	GetValueRef(Name) *TypedValue
	GetStaticTypeOf(Name) Type
	GetStaticTypeOfAt(ValuePath) Type
	Define(Name, TypedValue)
}

//----------------------------------------
// StaticBlock

// Embed in node to make it a BlockNode.
// TODO rename to StaticBlock
type StaticBlock struct {
	Block
	NumNames uint16
	Names    map[Name]uint16
}

// Implements BlockNode
func (sb *StaticBlock) InitStaticBlock(source BlockNode, parent BlockNode) {
	if sb.Names != nil {
		panic("StaticBlock.Names already initalized")
	}
	if parent == nil {
		sb.Block = *NewBlock(source, nil)
	} else {
		sb.Block = *NewBlock(source, parent.GetStaticBlock().GetBlock())
	}
	sb.NumNames = 0
	sb.Names = make(map[Name]uint16)
	return
}

// Implements BlockNode.
func (sb *StaticBlock) GetStaticBlock() *StaticBlock {
	return sb
}

// Does not implement BlockNode to prevent confusion.
// To get the static *Block, call Blocknode.GetStaticBlock().GetBlock().
func (sb *StaticBlock) GetBlock() *Block {
	return &sb.Block
}

// Implements BlockNode.
func (sb *StaticBlock) GetNames() (ns []Name) {
	ns = make([]Name, sb.NumNames)
	for n, idx := range sb.Names {
		ns[int(idx)] = n
	}
	return ns
}

// Implements BlockNode.
func (sb *StaticBlock) GetNumNames() (nn uint16) {
	return sb.NumNames
}

// Implements BlockNode.
func (sb *StaticBlock) GetParent() BlockNode {
	if sb.Block.Parent == nil {
		return nil
	} else {
		return sb.Block.Parent.Source
	}
}

// Implements BlockNode.
func (sb *StaticBlock) GetPathForName(n Name) ValuePath {
	if debug {
		if n == "_" {
			panic("should not happen")
		}
	}
	gen := uint16(1)
	idx, ok := sb.GetLocalIndex(n)
	bp := sb.GetParent()
	for {
		if ok {
			return NewValuePath(n, gen, idx)
		} else if bp != nil {
			idx, ok = bp.GetLocalIndex(n)
			bp = bp.GetParent()
			gen++
		} else {
			panic(fmt.Sprintf("name %s not declared", n))
		}
	}
}

// Implements BlockNode.
func (sb *StaticBlock) GetStaticTypeOf(n Name) Type {
	idx, ok := sb.GetLocalIndex(n)
	vs := sb.Block.Values
	bp := sb.GetParent()
	for {
		if ok {
			return vs[idx].T
		} else if bp != nil {
			idx, ok = bp.GetLocalIndex(n)
			vs = bp.GetStaticBlock().GetBlock().Values
			bp = bp.GetParent()
		} else {
			panic(fmt.Sprintf("name %s not declared", n))
		}
	}
}

// Implements BlockNode.
func (sb *StaticBlock) GetStaticTypeOfAt(path ValuePath) Type {
	return sb.Block.GetValueRefAt(path).T
}

// Implements BlockNode.
func (sb *StaticBlock) GetLocalIndex(n Name) (uint16, bool) {
	idx, ok := sb.Names[n]
	if debug {
		nt := reflect.TypeOf(sb.Source).String()
		debug.Printf("StaticBlock(%p %v).GetLocalIndex(%s) = %v, %v\n",
			sb, nt, n, idx, ok)
	}
	return idx, ok
}

// Implemented BlockNode.
// This method is too slow for runtime, but it is used
// during preprocessing to compute types.
// Returns nil if not defined.
func (sb *StaticBlock) GetValueRef(n Name) *TypedValue {
	idx, ok := sb.GetLocalIndex(n)
	vs := sb.Block.Values
	bp := sb.GetParent()
	for {
		if ok {
			return &vs[idx]
		} else if bp != nil {
			idx, ok = bp.GetLocalIndex(n)
			vs = bp.GetStaticBlock().GetBlock().Values
			bp = bp.GetParent()
		} else {
			return nil
		}
	}
}

// Implements BlockNode
// Statically declares a name definition.
// At runtime, use *Block.GetValueRef() etc which take
// path values, which are pre-computeed in the
// preprocessor.  tv.Type is always set unless the type is
// unknown during transcription, in which case initially
// tv must be empty, then Define(n,tv) called again once
// more with tv.Type set.
// Currently tv.V is only set when the value represents a
// TypedValue.  The intent of tv is to describe the invariant
// of a named value, at the minimum its type, but also
// sometimes the typeval value; but we could go further and
// store preprocessed constant results here too.  See
// "anyValue()" and "asValue()" for usage.
func (sb *StaticBlock) Define(n Name, tv TypedValue) {
	if debug {
		debug.Printf(
			"StaticBlock.Define(%s, %v)\n",
			n, tv)
	}
	if len(n) == 0 {
		panic("name cannot be zero")
	}
	if int(sb.NumNames) != len(sb.Names) {
		panic("StaticBlock.NumNames and len(.Names) mismatch")
	}
	if (1<<16 - 1) < sb.NumNames {
		panic("too many variables in block")
	}
	if tv.T == nil && tv.V != nil {
		panic("StaticBlock.Define() requires .T if .V is set")
	}
	if idx, exists := sb.Names[n]; exists {
		// Re-definitions cases are limited.
		if tv.T == nil {
			panic("StaticBlock.Define() the second time requires .T")
		}
		old := sb.Block.Values[idx]
		if old.T != nil && tv.T != old.T {
			panic(fmt.Sprintf(
				"StaticBlock.Define() cannot change .T; was %v, new %v",
				old.T, tv.T))
		}
		if old.V != nil && tv.V != old.V {
			// NOTE This case exists to alert of a more
			// egregious error than the case below it where the
			// same thing is re-defined. After fixing this
			// immediate issue, the re-definition issue will
			// probably surface, as there are possibly two
			// bugs.
			panic("StaticBlock.Define() cannot change .V")
		}
		if old.T != nil && old.V != nil {
			panic(fmt.Sprintf(
				"StaticBlock.Define(`%s`) already defined as %s",
				n, old.String()))
		}
		sb.Block.Values[idx] = tv
	} else {
		// The general case without re-definition.
		sb.Names[n] = sb.NumNames
		sb.NumNames++
		sb.Block.Values = append(sb.Block.Values, tv)
	}
}

// Implements BlockNode
func (sb *StaticBlock) SetStaticBlock(osb StaticBlock) {
	*sb = osb
}

var _ BlockNode = &FuncLitExpr{}
var _ BlockNode = &BlockStmt{}
var _ BlockNode = &ForStmt{}
var _ BlockNode = &IfStmt{}
var _ BlockNode = &RangeStmt{}
var _ BlockNode = &SelectCaseStmt{}
var _ BlockNode = &SwitchCaseStmt{}
var _ BlockNode = &FuncDecl{}
var _ BlockNode = &FileNode{}
var _ BlockNode = &PackageNode{}

//----------------------------------------
// Value Path

// A relative pointer to a TypedValue value
//  (a) a Block scope var or const
//  (b) a StructValue field
//  (c) a DeclaredType method
//  (d) a PackageNode declaration
//
// It is not used for interface methods, they are looked up by
// name, and thus is slower.
//
// Depth 0 is reserved for the global uverse namespace, so
// local names have generation 1 and so on.  Depth 0 and index
// 0 refers to the untyped nil value (as first defined in
// uverse.go).
//
// In declared types, the methods have generation 1 and if the
// underlying type is a struct, the fields of that have
// generation 2.  For undeclared structs, the fields have
// generation 1.  There is currently no javascript prototype or
// inheritance, so generation 3 and above are illegal (though
// could may change in the future).  In all cases, generation 0
// is reserved for global fields, such as x.type. Embedded
// structs are always flattened for efficient access.
//
type ValuePath struct {
	// 0 for uverse, or num parents to traverse +1.
	Depth uint16
	// index of TV in aforementioned
	// block/struct/package/declared-type.
	Index uint16
	Name  Name
}

func NewValuePath(n Name, depth, index uint16) ValuePath {
	return ValuePath{
		Depth: depth,
		Index: index,
		Name:  n,
	}
}

func (vp ValuePath) IsZero() bool {
	return vp.Depth == 0 && vp.Index == 0
}

type ValuePather interface {
	GetPathForName(Name) ValuePath
}

//----------------------------------------
// Utility

func (blx *BasicLitExpr) GetString() string {
	str, err := strconv.Unquote(blx.Value)
	if err != nil {
		panic("error in parsing string literal: " + err.Error())
	}
	return str
}

func (blx *BasicLitExpr) GetInt() int {
	i, err := strconv.Atoi(blx.Value)
	if err != nil {
		panic(err)
	}
	return i
}

type GnoAttribute int

const (
	ATTR_PREPROCESSED GnoAttribute = iota
	ATTR_PREDEFINED
	ATTR_TYPE_VALUE
	ATTR_TYPEOF_VALUE
	ATTR_LABEL
	ATTR_IOTA
)
