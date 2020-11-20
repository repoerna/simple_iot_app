package db

import (
	"context"
	"database/sql"
	"simpleiotapp/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomDevice(t *testing.T) Device {
	arg := CreateDeviceParams{
		Name:      util.RandomName(),
		Shortname: util.RandomShortName(),
		Enabled:   true,
	}

	device, err := testQueries.CreateDevice(context.Background(), arg)

	// check if no error
	require.NoError(t, err)
	// check if data created
	require.NotEmpty(t, device)

	// check data content
	require.Equal(t, device.Name, arg.Name)
	require.Equal(t, device.Shortname, arg.Shortname)
	require.Equal(t, device.Enabled, arg.Enabled)

	require.NotZero(t, device.ID)
	require.NotZero(t, device.Createdat)
	require.NotZero(t, device.Updatedat)

	return device
}
func TestCreateDevice(t *testing.T) {
	CreateRandomDevice(t)
}

func TestGetDevice(t *testing.T) {
	// create new device
	device1 := CreateRandomDevice(t)
	// get new created device
	device2, err := testQueries.GetDevice(context.Background(), device1.ID)

	// check get queries
	require.NoError(t, err)
	require.NotEmpty(t, device2)

	// check data content
	require.Equal(t, device1.ID, device2.ID)
	require.Equal(t, device1.Name, device2.Name)
	require.Equal(t, device1.Shortname, device2.Shortname)
	require.Equal(t, device1.Enabled, device2.Enabled)
	require.WithinDuration(t, device1.Createdat, device2.Createdat, time.Second)
	require.WithinDuration(t, device1.Updatedat, device2.Updatedat, time.Second)
}

func TestUpdateDevice(t *testing.T) {
	// create new device
	device1 := CreateRandomDevice(t)

	arg := UpdateDeviceParams{
		ID:        device1.ID,
		Name:      util.RandomName(),
		Shortname: util.RandomShortName(),
		Enabled:   util.RandomBool(),
	}

	// get new created device
	device2, err := testQueries.UpdateDevice(context.Background(), arg)
	// check get queries
	require.NoError(t, err)
	require.NotEmpty(t, device2)

	// check data content
	require.Equal(t, device1.ID, device2.ID)
	require.Equal(t, arg.Name, device2.Name)
	require.Equal(t, arg.Shortname, device2.Shortname)
	require.Equal(t, arg.Enabled, device2.Enabled)
	require.WithinDuration(t, device1.Createdat, device2.Createdat, time.Second)
	require.WithinDuration(t, device1.Updatedat, device2.Updatedat, time.Second)
}

func TestDeleteDevice(t *testing.T) {
	// create new device
	device1 := CreateRandomDevice(t)

	// delete
	err := testQueries.DeleteDevice(context.Background(), device1.ID)
	// check process
	require.NoError(t, err)

	// try get created data
	device2, err := testQueries.GetDevice(context.Background(), device1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, device2)
}

func TestListDevices(t *testing.T) {
	// create devices
	for i := 0; i < 10; i++ {
		CreateRandomDevice(t)
	}

	// prepare argument
	arg := GetDevicesParams{
		Offset: 5,
		Limit:  5,
	}

	// get
	devices, err := testQueries.GetDevices(context.Background(), arg)
	// check process
	require.NoError(t, err)
	require.Len(t, devices, 5)

	for _, device := range devices {
		require.NotEmpty(t, device)
	}

}
