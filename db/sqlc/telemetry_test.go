package db

import (
	"context"
	"simpleiotapp/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTM(t *testing.T, gendevice Device) Telemetry {
	lan, lat := util.RandomLongLat()
	arg := CreateTMParams{
		Deviceid:  gendevice.ID,
		Longitude: lan,
		Latitude:  lat,
		Value:     util.RandomSQLNullFloat64(),
		Value2:    util.RandomSQLNullFloat64(),
		Value3:    util.RandomSQLNullFloat64(),
		Value4:    util.RandomSQLNullFloat64(),
	}

	tm, err := testQueries.CreateTM(context.Background(), arg)
	// check if no error
	require.NoError(t, err)
	// check if data created
	require.NotEmpty(t, tm)

	// check data content
	require.Equal(t, arg.Deviceid, tm.Deviceid)
	require.Equal(t, arg.Value, tm.Value)
	require.Equal(t, arg.Value2, tm.Value2)
	require.Equal(t, arg.Value3, tm.Value3)
	require.Equal(t, arg.Value4, tm.Value4)

	require.NotZero(t, tm.ID)
	require.NotZero(t, tm.Createdat)

	return tm
}

func TestGetTMByDeviceID(t *testing.T) {
	device := CreateRandomDevice(t)

	for i := 0; i < 10; i++ {
		createRandomTM(t, device)
	}

	arg := GetTMByDeviceIDParams{
		Deviceid: device.ID,
		Limit:    5,
		Offset:   5,
	}

	data, err := testQueries.GetTMByDeviceID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, data, 5)

	for _, datum := range data {
		require.NotEmpty(t, datum)
	}

}

func TestGetTMByID(t *testing.T) {
	// create new device
	device := CreateRandomDevice(t)
	// get new created TM data
	data1 := createRandomTM(t, device)

	// get data
	data2, err := testQueries.GetTMByID(context.Background(), data1.ID)

	// check get queries
	require.NoError(t, err)
	require.NotEmpty(t, data2)

	// check data content
	require.Equal(t, data1.ID, data2.ID)
	require.Equal(t, data1.Deviceid, data2.Deviceid)
	require.Equal(t, data1.Value, data2.Value)
	require.Equal(t, data1.Value2, data2.Value2)
	require.Equal(t, data1.Value3, data2.Value3)
	require.Equal(t, data1.Value4, data2.Value4)
	require.WithinDuration(t, data1.Createdat, data2.Createdat, time.Second)
}

func TestCreateTM(t *testing.T) {
	// create new device
	device := CreateRandomDevice(t)
	// get new created TM data
	data1 := createRandomTM(t, device)

	// get data
	data2, err := testQueries.GetTMByID(context.Background(), data1.ID)

	// check get queries
	require.NoError(t, err)
	require.NotEmpty(t, data2)

	// check data content
	require.Equal(t, data1.ID, data2.ID)
	require.Equal(t, data1.Deviceid, data2.Deviceid)
	require.Equal(t, data1.Value, data2.Value)
	require.Equal(t, data1.Value2, data2.Value2)
	require.Equal(t, data1.Value3, data2.Value3)
	require.Equal(t, data1.Value4, data2.Value4)
	require.WithinDuration(t, data1.Createdat, data2.Createdat, time.Second)
}
