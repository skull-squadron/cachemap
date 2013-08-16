package cachemap

import (
    "testing"
)

var testKey = uint64(1)
var testKeyBytes = uint64(8)
var testKey2 = "sex"
var testKeyBytes2 = uint64(3)
var testValue = uint64(0xdeadbeef)
var testValueBytes = uint64(8)
var testValue2 = uint64(0xc0ffee)
var testValueBytes2 = uint64(8)

type CreateDouble int

func (x *CreateDouble) Create(key Key, item *Item) error {
    *x = 1
    return nil
}

func (x *CreateDouble) Read(key Key) (item *Item, err error) {
    return nil, nil
}

func (x *CreateDouble) Update(key Key, item *Item) error {
    return nil
}

func (x *CreateDouble) UpdateMeta(key Key, item *Item) error {
    return nil
}

func (x *CreateDouble) Delete(key Key) error {
    return nil
}

func (x *CreateDouble) GetBytes() (uint64, error) {
    return 0, nil
}

func TestCreate(t *testing.T) {
    var p CreateDouble
    cm := NewCacheMap(&p)
    err := cm.Add(testKey, testValue, testKeyBytes, testValueBytes)
    if err != nil {
        t.Error(err)
    }
    v, err := cm.Get(testKey)
    expected := testValue
    if v.(uint64) != expected {
        t.Error("Failed to Add actual=", v, " expected=", expected)
    }
    if p != CreateDouble(testKey) {
        t.Error("Failed to call persistent Create")
    }
}



type ReadDouble int

func (x *ReadDouble) Create(key Key, item *Item) error {
    return nil
}

func (x *ReadDouble) Read(key Key) (item *Item, err error) {
    return
}

func (x *ReadDouble) Update(key Key, item *Item) error {
    return nil
}

func (x *ReadDouble) UpdateMeta(key Key, item *Item) error {
    return nil
}

func (x *ReadDouble) Delete(key Key) error {
    return nil
}

func (x *ReadDouble) GetBytes() (uint64, error) {
    return 0, nil
}

func TestRead(t *testing.T) {
    var p ReadDouble
    cm := NewCacheMap(&p)
    err := cm.Add(testKey, testValue, testKeyBytes, testValueBytes)
    if err != nil {
        t.Error(err)
    }
    v, err := cm.Get(testKey)
    if err != nil {
        t.Error(err)
    }
    expected := ReadDouble(testValue)
    if v.(uint64) != uint64(expected) {
        t.Error("Failed to call Read actual=", v, " expected=", expected)
    }
}

type UpdateDouble int

func (x *UpdateDouble) Create(key Key, item *Item) error {
    return nil
}

func (x *UpdateDouble) Read(key Key) (item *Item, err error) {
    return
}

func (x *UpdateDouble) Update(key Key, item *Item) error {
    *x = 1
    return nil
}

func (x *UpdateDouble) UpdateMeta(key Key, item *Item) error {
    return nil
}

func (x *UpdateDouble) Delete(key Key) error {
    return nil
}

func (x *UpdateDouble) GetBytes() (uint64, error) {
    return 0, nil
}


func TestUpdate(t *testing.T) {
    var p UpdateDouble
    cm := NewCacheMap(&p)
    err := cm.Add(testKey, testValue, testKeyBytes, testValueBytes)
    if err != nil {
        t.Error(err)
    }
    err = cm.Add(testKey, testValue2, testKeyBytes, testValueBytes2)
    if err != nil {
        t.Error(err)
    }
    v, err := cm.Get(testKey)
    if err != nil {
        t.Error(err)
    }
    expected := UpdateDouble(testValue2)
    if v.(uint64) != uint64(expected) {
        t.Error("Failed to call Update actual=", v, " expected=", expected)
    }
    if p != 1 {
        t.Error("Failed to call persistent Update")
    }
}

type DeleteDouble string

func (x *DeleteDouble) Create(key Key, item *Item) error {
    return nil
}

func (x *DeleteDouble) Read(key Key) (item *Item, err error) {
    return
}

func (x *DeleteDouble) Update(key Key, item *Item) error {
    return nil
}

func (x *DeleteDouble) UpdateMeta(key Key, item *Item) error {
    return nil
}

func (x *DeleteDouble) Delete(key Key) error {
    *x = "works"
    return nil
}

func (x *DeleteDouble) GetBytes() (uint64, error) {
    return 0, nil
}

func TestDelete(t *testing.T) {
    var p DeleteDouble
    cm := NewCacheMap(&p)
    err := cm.Add(testKey2, testValue, testKeyBytes, testValueBytes)
    if err != nil {
        t.Error(err)
    }
    err = cm.Delete(testKey2)
    if err != nil {
        t.Error(err)
    }
    v, err := cm.Get(testKey2)
    if err != nil {
        t.Error(err)
    }
    if v != nil {
        t.Error("Failed to call Delete actual=", v, " expected=", nil)
    }
    if p != "works" {
        t.Error("Failed to call persistent Delete")
    }
}
