package uuid_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/invopop/gobl/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUUIDParsing(t *testing.T) {
	v1s := "03907310-8daa-11eb-8dcd-0242ac130003"
	v4s := "0def554c-54fd-4b3b-9ea0-4f2d288d4435"

	u1, err := uuid.Parse(v1s)
	assert.NoError(t, err)
	assert.Equal(t, uuid.Version(1), u1.Version())

	u4, err := uuid.Parse(v4s)
	assert.NoError(t, err)
	assert.Equal(t, uuid.Version(4), u4.Version())

	u1, err = uuid.Parse("")
	assert.NoError(t, err)
	assert.Equal(t, u1, uuid.Empty)

	u1, err = uuid.Parse("fooo")
	assert.ErrorContains(t, err, "invalid UUID length: 4")
	assert.Equal(t, u1, uuid.Empty)

	u1 = uuid.ShouldParse("")
	assert.True(t, u1.IsZero())
	u1 = uuid.ShouldParse("fooo")
	assert.True(t, u1.IsZero())
	u1 = uuid.ShouldParse(v1s)
	assert.Equal(t, v1s, u1.String())

	t.Run("must parse", func(t *testing.T) {
		assert.Panics(t, func() {
			uuid.MustParse("fooo")
		})
	})
}

func TestUUIDIsZero(t *testing.T) {
	var u1 uuid.UUID
	assert.True(t, u1.IsZero())
	var up1 *uuid.UUID
	assert.True(t, up1.IsZero())
	u1 = uuid.Zero
	assert.True(t, u1.IsZero())

	u1 = uuid.V1()
	assert.False(t, u1.IsZero())
	up1 = &u1
	assert.False(t, up1.IsZero())
}

func TestUUIDTimestasmp(t *testing.T) {
	u := uuid.V1()
	ts := u.Timestamp()
	assert.False(t, ts.IsZero())
	assert.WithinDuration(t, ts, time.Now(), 10*time.Second)

	u = uuid.V6()
	ts = u.Timestamp()
	assert.False(t, ts.IsZero())
	assert.WithinDuration(t, ts, time.Now(), 10*time.Second)

	u = uuid.V7()
	ts = u.Timestamp()
	assert.False(t, ts.IsZero())
	assert.WithinDuration(t, ts, time.Now(), 10*time.Second)

	u = uuid.V4()
	ts = u.Timestamp()
	assert.True(t, ts.IsZero())
}

func TestNodeID(t *testing.T) {
	a := uuid.NodeID()
	assert.Len(t, a, 12)

	uuid.SetRandomNodeID()
	assert.Len(t, uuid.NodeID(), 12)
	assert.NotEqual(t, a, uuid.NodeID())
}

func TestUUIDJSON(t *testing.T) {
	v1s := "03907310-8daa-11eb-8dcd-0242ac130003"
	type testJSON struct {
		ID      uuid.UUID  `json:"id"`
		EmptyID uuid.UUID  `json:"empty_id,omitempty"`
		OptID   *uuid.UUID `json:"opt_id,omitempty"`
	}

	v := testJSON{ID: uuid.V1()}
	assert.False(t, v.ID.IsZero())

	data, err := json.Marshal(v)
	require.NoError(t, err)
	assert.Equal(t, `{"id":"`+v.ID.String()+`"}`, string(data))

	v2 := testJSON{}
	assert.True(t, v2.ID.IsZero())

	b := []byte(`{"id":"` + v1s + `"}`)
	if err := json.Unmarshal(b, &v2); err != nil {
		t.Errorf("did not expect unmarshal to fail, err: %v", err.Error())
	}
	if v2.ID.String() != v1s {
		t.Errorf("did not get same string back, got: %v", v2.ID.String())
	}
}

func TestUUIDUnmarshalJSON(t *testing.T) {
	type m struct {
		ID  uuid.UUID  `json:"id"`
		PID *uuid.UUID `json:"pid"`
	}
	pid := uuid.UUID("03907310-8daa-11eb-8dcd-0242ac130003")
	tests := []struct {
		name string
		data string
		want m
		err  string
	}{
		{
			name: "valid UUID",
			data: `{"id":"03907310-8daa-11eb-8dcd-0242ac130003"}`,
			want: m{ID: "03907310-8daa-11eb-8dcd-0242ac130003"},
		},
		{
			name: "zero UUID",
			data: `{"id":"00000000-0000-0000-0000-000000000000"}`,
			want: m{ID: uuid.Zero},
		},
		{
			name: "invalid UUID",
			data: `{"id":"invalid-uuid"}`,
			want: m{},
			err:  "invalid UUID length: 12",
		},
		{
			name: "empty string",
			data: `{"id":""}`,
			want: m{},
			err:  "",
		},
		{
			name: "null",
			data: `{"id":null}`,
			want: m{},
			err:  "",
		},
		{
			name: "pointer valid UUID",
			data: `{"pid":"03907310-8daa-11eb-8dcd-0242ac130003"}`,
			want: m{PID: &pid},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out m
			err := json.Unmarshal([]byte(tt.data), &out)
			assert.Equal(t, tt.want, out)
			if tt.err != "" {
				assert.ErrorContains(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	type m struct {
		ID  uuid.UUID  `json:"id"`
		EID uuid.UUID  `json:"eid,omitempty"`
		PID *uuid.UUID `json:"pid,omitempty"`
	}
	zpid := uuid.Zero
	tests := []struct {
		name string
		data any
		want string
		err  string
	}{
		{
			name: "valid UUID",
			data: m{ID: "03907310-8daa-11eb-8dcd-0242ac130003"},
			want: `{"id":"03907310-8daa-11eb-8dcd-0242ac130003"}`,
		},
		{
			name: "empty UUID",
			data: m{ID: ""},
			want: `{"id":""}`,
		},
		{
			name: "zero UUID",
			data: m{ID: uuid.Zero},
			want: `{"id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name: "zero pointer UUID",
			data: m{PID: &zpid},
			want: `{"id":"","pid":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name: "zero empty UUID",
			data: m{EID: uuid.Zero},
			want: `{"id":"","eid":"00000000-0000-0000-0000-000000000000"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.data)
			assert.Equal(t, tt.want, string(data))
			if tt.err != "" {
				assert.ErrorContains(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	u := new(uuid.UUID)
	uuid.Normalize(u)
	assert.Empty(t, u)

	u2 := uuid.Zero
	uuid.Normalize(&u2)
	assert.Empty(t, u2)

	u3 := uuid.MustParse("03907310-8daa-11eb-8dcd-0242ac130003")
	uuid.Normalize(&u3)
	assert.Equal(t, "03907310-8daa-11eb-8dcd-0242ac130003", u3.String())

	assert.NotPanics(t, func() {
		uuid.Normalize(nil)
	})
}

func TestUUIDv3(t *testing.T) {
	ns := uuid.MustParse("0654a3f4-8ad5-44c8-828e-c25f7ccd6550")
	u := uuid.NewV3(ns, []byte("hello, world"))

	assert.Equal(t, 3, int(u.Version()))
	assert.Equal(t, "61cfb897-b1bb-382b-bab9-a7ba465a27fa", u.String())
}

func TestUUIDv5(t *testing.T) {
	ns := uuid.MustParse("0654a3f4-8ad5-44c8-828e-c25f7ccd6550")
	u := uuid.NewV5(ns, []byte("hello, world"))

	assert.Equal(t, 5, int(u.Version()))
	assert.Equal(t, "1f53a310-2a17-5acb-b76a-c39495e5356f", u.String())
}

func TestUUIDBase64(t *testing.T) {
	u := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")

	s := u.Base64()
	assert.Equal(t, "9HrBC1jMA3KFZw4CssPUeQ", s)

	u2, err := uuid.ParseBase64(s)
	require.NoError(t, err)
	assert.Equal(t, u.String(), u2.String())

	u2, err = uuid.ParseBase64(u.String())
	require.NoError(t, err)
	assert.Equal(t, u.String(), u2.String())
}

func TestUUIDBunary(t *testing.T) {
	u := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")

	b := u.Bytes()
	assert.Equal(t, 16, len(b))

	out, err := u.MarshalBinary()
	require.NoError(t, err)
	assert.Equal(t, b, out)

	u2 := new(uuid.UUID)
	err = u2.UnmarshalBinary(b)
	require.NoError(t, err)

	u2 = new(uuid.UUID)
	err = u2.UnmarshalBinary([]byte("invalid"))
	assert.ErrorContains(t, err, "invalid UUID (got 7 bytes)")
}
