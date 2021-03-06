package gno

import (
	"encoding/binary"
)

type ObjectID struct {
	RealmID        // base
	Ordinal uint64 // counter
}

func (oid ObjectID) Bytes() []byte {
	bz := make([]byte, HashSize+8)
	copy(bz[:HashSize], oid.RealmID.Bytes())
	binary.BigEndian.PutUint64(
		bz[HashSize:], uint64(oid.Ordinal))
	return bz
}

func (oid ObjectID) IsZero() bool {
	if debug {
		if oid.RealmID.IsZero() && oid.Ordinal != 0 {
			panic("should not happen")
		}
	}
	return oid.RealmID.IsZero()
}

type Object interface {
	GetObjectInfo() *ObjectInfo
	GetObjectID() ObjectID
	GetOwner() Object
	SetOwner(Object)
	GetIsOwned() bool
	GetIsReal() bool
	IncRefCount() int
	DecRefCount() int
	GetRefCount() int
	GetIsNewReal() bool
	SetIsNewReal(bool)
	GetIsDirty() bool
	SetIsDirty(bool)
	GetIsDeleted() bool
	SetIsDeleted(bool)

	ValuePreimage(rlm *Realm, owned bool) ValuePreimage
}

var _ Object = &ArrayValue{}
var _ Object = &StructValue{}
var _ Object = &MapValue{}
var _ Object = &Block{}

type ObjectInfo struct {
	ID        ObjectID  // set if real.
	Hash      ValueHash // if dirty, outdated.
	Owner     Object    // parent in the ownership tree.
	RefCount  int       // deleted/gc'd if 0.
	IsNewReal bool      // if new and owner is real.
	IsDirty   bool      // if real but modified; hash is outdated if true.
	IsDeleted bool      // if real but no longer referenced.
}

func (oi *ObjectInfo) GetObjectInfo() *ObjectInfo {
	return oi
}

func (oi *ObjectInfo) GetObjectID() ObjectID {
	return oi.ID
}

func (oi *ObjectInfo) GetOwner() Object {
	return oi.Owner
}

func (oi *ObjectInfo) SetOwner(po Object) {
	oi.Owner = po
}

func (oi *ObjectInfo) GetIsOwned() bool {
	return oi.Owner != nil
}

// NOTE: does not return true for new reals.
func (oi *ObjectInfo) GetIsReal() bool {
	return !oi.ID.IsZero()
}

func (oi *ObjectInfo) IncRefCount() int {
	oi.RefCount++
	return oi.RefCount
}

func (oi *ObjectInfo) DecRefCount() int {
	oi.RefCount--
	if debug {
		if oi.RefCount < 0 {
			panic("should not happen")
		}
	}
	return oi.RefCount
}

func (oi *ObjectInfo) GetRefCount() int {
	return oi.RefCount
}

func (oi *ObjectInfo) GetIsNewReal() bool {
	return oi.IsNewReal
}

func (oi *ObjectInfo) SetIsNewReal(x bool) {
	oi.IsNewReal = x
}

func (oi *ObjectInfo) GetIsDirty() bool {
	return oi.IsDirty
}

func (oi *ObjectInfo) SetIsDirty(x bool) {
	oi.IsDirty = x
}

func (oi *ObjectInfo) GetIsDeleted() bool {
	return oi.IsDeleted
}

func (oi *ObjectInfo) SetIsDeleted(x bool) {
	oi.IsDeleted = x
}
