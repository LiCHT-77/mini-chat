// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/LiCHT-77/mini-chat/room/ent/room"
	"github.com/LiCHT-77/mini-chat/room/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	roomMixin := schema.Room{}.Mixin()
	roomMixinFields0 := roomMixin[0].Fields()
	_ = roomMixinFields0
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescCreateTime is the schema descriptor for create_time field.
	roomDescCreateTime := roomMixinFields0[0].Descriptor()
	// room.DefaultCreateTime holds the default value on creation for the create_time field.
	room.DefaultCreateTime = roomDescCreateTime.Default.(func() time.Time)
	// roomDescUpdateTime is the schema descriptor for update_time field.
	roomDescUpdateTime := roomMixinFields0[1].Descriptor()
	// room.DefaultUpdateTime holds the default value on creation for the update_time field.
	room.DefaultUpdateTime = roomDescUpdateTime.Default.(func() time.Time)
	// room.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	room.UpdateDefaultUpdateTime = roomDescUpdateTime.UpdateDefault.(func() time.Time)
}
