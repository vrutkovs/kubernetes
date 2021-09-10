// Copyright 2015, 2018 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dbus

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strconv"

	"github.com/godbus/dbus/v5"
)

// Who can be used to specify which process to kill in the unit via the KillUnitWithTarget API
type Who string

const (
	// All sends the signal to all processes in the unit
	All Who = "all"
	// Main sends the signal to the main process of the unit
	Main Who = "main"
	// Control sends the signal to the control process of the unit
	Control Who = "control"
)

func (c *Conn) jobComplete(signal *dbus.Signal) {
	var id uint32
	var job dbus.ObjectPath
	var unit string
	var result string
	dbus.Store(signal.Body, &id, &job, &unit, &result)
	c.jobListener.Lock()
	out, ok := c.jobListener.jobs[job]
	if ok {
		out <- result
		delete(c.jobListener.jobs, job)
	}
	c.jobListener.Unlock()
}

func (c *Conn) startJob(ctx context.Context, ch chan<- string, job string, args ...interface{}) (int, error) {
	if ch != nil {
		c.jobListener.Lock()
		defer c.jobListener.Unlock()
	}

	var p dbus.ObjectPath
	err := c.sysobj.CallWithContext(ctx, job, 0, args...).Store(&p)
	if err != nil {
		return 0, err
	}

	if ch != nil {
		c.jobListener.jobs[p] = ch
	}

	// ignore error since 0 is fine if conversion fails
	jobID, _ := strconv.Atoi(path.Base(string(p)))

	return jobID, nil
}

// Deprecated: use StartUnitContext instead.
func (c *Conn) StartUnit(name string, mode string, ch chan<- string) (int, error) {
	return c.StartUnitContext(context.Background(), name, mode, ch)
}

// StartUnitContext enqueues a start job and depending jobs, if any (unless otherwise
// specified by the mode string).
//
// Takes the unit to activate, plus a mode string. The mode needs to be one of
// replace, fail, isolate, ignore-dependencies, ignore-requirements. If
// "replace" the call will start the unit and its dependencies, possibly
// replacing already queued jobs that conflict with this. If "fail" the call
// will start the unit and its dependencies, but will fail if this would change
// an already queued job. If "isolate" the call will start the unit in question
// and terminate all units that aren't dependencies of it. If
// "ignore-dependencies" it will start a unit but ignore all its dependencies.
// If "ignore-requirements" it will start a unit but only ignore the
// requirement dependencies. It is not recommended to make use of the latter
// two options.
//
// If the provided channel is non-nil, a result string will be sent to it upon
// job completion: one of done, canceled, timeout, failed, dependency, skipped.
// done indicates successful execution of a job. canceled indicates that a job
// has been canceled  before it finished execution. timeout indicates that the
// job timeout was reached. failed indicates that the job failed. dependency
// indicates that a job this job has been depending on failed and the job hence
// has been removed too. skipped indicates that a job was skipped because it
// didn't apply to the units current state.
//
// If no error occurs, the ID of the underlying systemd job will be returned. There
// does exist the possibility for no error to be returned, but for the returned job
// ID to be 0. In this case, the actual underlying ID is not 0 and this datapoint
// should not be considered authoritative.
//
// If an error does occur, it will be returned to the user alongside a job ID of 0.
<<<<<<< HEAD
func (c *Conn) StartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.StartUnit", name, mode)
||||||| 5e58841cce7
func (c *Conn) StartUnit(name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.StartUnit", name, mode)
=======
// Deprecated: use StartUnitContext instead
func (c *Conn) StartUnit(name string, mode string, ch chan<- string) (int, error) {
	return c.StartUnitContext(context.Background(), name, mode, ch)
}

// StartUnitContext same as StartUnit with context
func (c *Conn) StartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.StartUnit", name, mode)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use StopUnitContext instead.
||||||| 5e58841cce7
// StopUnit is similar to StartUnit but stops the specified unit rather
// than starting it.
=======
// StopUnit is similar to StartUnit but stops the specified unit rather
// than starting it.
// Deprecated: use StopUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) StopUnit(name string, mode string, ch chan<- string) (int, error) {
<<<<<<< HEAD
	return c.StopUnitContext(context.Background(), name, mode, ch)
}

// StopUnitContext is similar to StartUnitContext, but stops the specified unit
// rather than starting it.
func (c *Conn) StopUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.StopUnit", name, mode)
||||||| 5e58841cce7
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.StopUnit", name, mode)
=======
	return c.StopUnitContext(context.Background(), name, mode, ch)
}

// StopUnitContext same as StopUnit with context
func (c *Conn) StopUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.StopUnit", name, mode)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use ReloadUnitContext instead.
||||||| 5e58841cce7
// ReloadUnit reloads a unit.  Reloading is done only if the unit is already running and fails otherwise.
=======
// ReloadUnit reloads a unit.  Reloading is done only if the unit is already running and fails otherwise.
// Deprecated: use ReloadUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) ReloadUnit(name string, mode string, ch chan<- string) (int, error) {
<<<<<<< HEAD
	return c.ReloadUnitContext(context.Background(), name, mode, ch)
||||||| 5e58841cce7
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.ReloadUnit", name, mode)
=======
	return c.ReloadUnitContext(context.Background(), name, mode, ch)
}

// ReloadUnitContext same as ReloadUnit with context
func (c *Conn) ReloadUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.ReloadUnit", name, mode)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// ReloadUnitContext reloads a unit. Reloading is done only if the unit
// is already running, and fails otherwise.
func (c *Conn) ReloadUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.ReloadUnit", name, mode)
}

// Deprecated: use RestartUnitContext instead.
||||||| 5e58841cce7
// RestartUnit restarts a service.  If a service is restarted that isn't
// running it will be started.
=======
// RestartUnit restarts a service.  If a service is restarted that isn't
// running it will be started.
// Deprecated: use RestartUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) RestartUnit(name string, mode string, ch chan<- string) (int, error) {
<<<<<<< HEAD
	return c.RestartUnitContext(context.Background(), name, mode, ch)
||||||| 5e58841cce7
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.RestartUnit", name, mode)
=======
	return c.RestartUnitContext(context.Background(), name, mode, ch)
}

// RestartUnitContext same as RestartUnit with context
func (c *Conn) RestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.RestartUnit", name, mode)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// RestartUnitContext restarts a service. If a service is restarted that isn't
// running it will be started.
func (c *Conn) RestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.RestartUnit", name, mode)
}

// Deprecated: use TryRestartUnitContext instead.
||||||| 5e58841cce7
// TryRestartUnit is like RestartUnit, except that a service that isn't running
// is not affected by the restart.
=======
// TryRestartUnit is like RestartUnit, except that a service that isn't running
// is not affected by the restart.
// Deprecated: use TryRestartUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) TryRestartUnit(name string, mode string, ch chan<- string) (int, error) {
<<<<<<< HEAD
	return c.TryRestartUnitContext(context.Background(), name, mode, ch)
}

// TryRestartUnitContext is like RestartUnitContext, except that a service that
// isn't running is not affected by the restart.
func (c *Conn) TryRestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.TryRestartUnit", name, mode)
||||||| 5e58841cce7
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.TryRestartUnit", name, mode)
=======
	return c.TryRestartUnitContext(context.Background(), name, mode, ch)
}

// TryRestartUnitContext same as TryRestartUnit with context
func (c *Conn) TryRestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.TryRestartUnit", name, mode)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use ReloadOrRestartUnitContext instead.
||||||| 5e58841cce7
// ReloadOrRestartUnit attempts a reload if the unit supports it and use a restart
// otherwise.
=======
// ReloadOrRestartUnit attempts a reload if the unit supports it and use a restart
// otherwise.
// Deprecated: use ReloadOrRestartUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) ReloadOrRestartUnit(name string, mode string, ch chan<- string) (int, error) {
<<<<<<< HEAD
	return c.ReloadOrRestartUnitContext(context.Background(), name, mode, ch)
}

// ReloadOrRestartUnitContext attempts a reload if the unit supports it and use
// a restart otherwise.
func (c *Conn) ReloadOrRestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.ReloadOrRestartUnit", name, mode)
||||||| 5e58841cce7
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.ReloadOrRestartUnit", name, mode)
=======
	return c.ReloadOrRestartUnitContext(context.Background(), name, mode, ch)
}

// ReloadOrRestartUnitContext same as ReloadOrRestartUnit with context
func (c *Conn) ReloadOrRestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.ReloadOrRestartUnit", name, mode)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use ReloadOrTryRestartUnitContext instead.
||||||| 5e58841cce7
// ReloadOrTryRestartUnit attempts a reload if the unit supports it and use a "Try"
// flavored restart otherwise.
=======
// ReloadOrTryRestartUnit attempts a reload if the unit supports it and use a "Try"
// flavored restart otherwise.
// Deprecated: use ReloadOrTryRestartUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) ReloadOrTryRestartUnit(name string, mode string, ch chan<- string) (int, error) {
<<<<<<< HEAD
	return c.ReloadOrTryRestartUnitContext(context.Background(), name, mode, ch)
}

// ReloadOrTryRestartUnitContext attempts a reload if the unit supports it,
// and use a "Try" flavored restart otherwise.
func (c *Conn) ReloadOrTryRestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.ReloadOrTryRestartUnit", name, mode)
}

// Deprecated: use StartTransientUnitContext instead.
func (c *Conn) StartTransientUnit(name string, mode string, properties []Property, ch chan<- string) (int, error) {
	return c.StartTransientUnitContext(context.Background(), name, mode, properties, ch)
||||||| 5e58841cce7
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.ReloadOrTryRestartUnit", name, mode)
=======
	return c.ReloadOrTryRestartUnitContext(context.Background(), name, mode, ch)
}

// ReloadOrTryRestartUnitContext same as ReloadOrTryRestartUnit with context
func (c *Conn) ReloadOrTryRestartUnitContext(ctx context.Context, name string, mode string, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.ReloadOrTryRestartUnit", name, mode)
>>>>>>> v1.21.4
}

// StartTransientUnitContext may be used to create and start a transient unit, which
// will be released as soon as it is not running or referenced anymore or the
// system is rebooted. name is the unit name including suffix, and must be
// unique. mode is the same as in StartUnitContext, properties contains properties
// of the unit.
<<<<<<< HEAD
func (c *Conn) StartTransientUnitContext(ctx context.Context, name string, mode string, properties []Property, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.StartTransientUnit", name, mode, properties, make([]PropertyCollection, 0))
||||||| 5e58841cce7
func (c *Conn) StartTransientUnit(name string, mode string, properties []Property, ch chan<- string) (int, error) {
	return c.startJob(ch, "org.freedesktop.systemd1.Manager.StartTransientUnit", name, mode, properties, make([]PropertyCollection, 0))
=======
// Deprecated: use StartTransientUnitContext instead
func (c *Conn) StartTransientUnit(name string, mode string, properties []Property, ch chan<- string) (int, error) {
	return c.StartTransientUnitContext(context.Background(), name, mode, properties, ch)
}

// StartTransientUnitContext same as StartTransientUnit with context
func (c *Conn) StartTransientUnitContext(ctx context.Context, name string, mode string, properties []Property, ch chan<- string) (int, error) {
	return c.startJob(ctx, ch, "org.freedesktop.systemd1.Manager.StartTransientUnit", name, mode, properties, make([]PropertyCollection, 0))
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use KillUnitContext instead.
||||||| 5e58841cce7
// KillUnit takes the unit name and a UNIX signal number to send.  All of the unit's
// processes are killed.
=======
// KillUnit takes the unit name and a UNIX signal number to send.  All of the unit's
// processes are killed.
// Deprecated: use KillUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) KillUnit(name string, signal int32) {
<<<<<<< HEAD
	c.KillUnitContext(context.Background(), name, signal)
}

// KillUnitContext takes the unit name and a UNIX signal number to send.
// All of the unit's processes are killed.
func (c *Conn) KillUnitContext(ctx context.Context, name string, signal int32) {
	c.KillUnitWithTarget(ctx, name, All, signal)
}

// KillUnitWithTarget is like KillUnitContext, but allows you to specify which
// process in the unit to send the signal to.
func (c *Conn) KillUnitWithTarget(ctx context.Context, name string, target Who, signal int32) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.KillUnit", 0, name, string(target), signal).Store()
||||||| 5e58841cce7
	c.sysobj.Call("org.freedesktop.systemd1.Manager.KillUnit", 0, name, "all", signal).Store()
=======
	c.KillUnitContext(context.Background(), name, signal)
}

// KillUnitContext same as KillUnit with context
func (c *Conn) KillUnitContext(ctx context.Context, name string, signal int32) {
	c.KillUnitWithTarget(ctx, name, All, signal)
}

// KillUnitWithTarget is like KillUnitContext, but allows you to specify which process in the unit to send the signal to
func (c *Conn) KillUnitWithTarget(ctx context.Context, name string, target Who, signal int32) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.KillUnit", 0, name, string(target), signal).Store()
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use ResetFailedUnitContext instead.
||||||| 5e58841cce7
// ResetFailedUnit resets the "failed" state of a specific unit.
=======
// ResetFailedUnit resets the "failed" state of a specific unit.
// Deprecated: use ResetFailedUnitContext instead
>>>>>>> v1.21.4
func (c *Conn) ResetFailedUnit(name string) error {
<<<<<<< HEAD
	return c.ResetFailedUnitContext(context.Background(), name)
}

// ResetFailedUnitContext resets the "failed" state of a specific unit.
func (c *Conn) ResetFailedUnitContext(ctx context.Context, name string) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ResetFailedUnit", 0, name).Store()
||||||| 5e58841cce7
	return c.sysobj.Call("org.freedesktop.systemd1.Manager.ResetFailedUnit", 0, name).Store()
=======
	return c.ResetFailedUnitContext(context.Background(), name)
}

// ResetFailedUnitContext same as ResetFailedUnit with context
func (c *Conn) ResetFailedUnitContext(ctx context.Context, name string) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ResetFailedUnit", 0, name).Store()
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use SystemStateContext instead.
||||||| 5e58841cce7
// SystemState returns the systemd state. Equivalent to `systemctl is-system-running`.
=======
// SystemState returns the systemd state. Equivalent to `systemctl is-system-running`.
// Deprecated: use SystemStateContext instead
>>>>>>> v1.21.4
func (c *Conn) SystemState() (*Property, error) {
<<<<<<< HEAD
	return c.SystemStateContext(context.Background())
}

// SystemStateContext returns the systemd state. Equivalent to
// systemctl is-system-running.
func (c *Conn) SystemStateContext(ctx context.Context) (*Property, error) {
||||||| 5e58841cce7
=======
	return c.SystemStateContext(context.Background())
}

// SystemStateContext same as SystemState with context
func (c *Conn) SystemStateContext(ctx context.Context) (*Property, error) {
>>>>>>> v1.21.4
	var err error
	var prop dbus.Variant

	obj := c.sysconn.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")
	err = obj.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.systemd1.Manager", "SystemState").Store(&prop)
	if err != nil {
		return nil, err
	}

	return &Property{Name: "SystemState", Value: prop}, nil
}

<<<<<<< HEAD
// getProperties takes the unit path and returns all of its dbus object properties, for the given dbus interface.
func (c *Conn) getProperties(ctx context.Context, path dbus.ObjectPath, dbusInterface string) (map[string]interface{}, error) {
||||||| 5e58841cce7
// getProperties takes the unit path and returns all of its dbus object properties, for the given dbus interface
func (c *Conn) getProperties(path dbus.ObjectPath, dbusInterface string) (map[string]interface{}, error) {
=======
// getProperties takes the unit path and returns all of its dbus object properties, for the given dbus interface
func (c *Conn) getProperties(ctx context.Context, path dbus.ObjectPath, dbusInterface string) (map[string]interface{}, error) {
>>>>>>> v1.21.4
	var err error
	var props map[string]dbus.Variant

	if !path.IsValid() {
		return nil, fmt.Errorf("invalid unit name: %v", path)
	}

	obj := c.sysconn.Object("org.freedesktop.systemd1", path)
	err = obj.CallWithContext(ctx, "org.freedesktop.DBus.Properties.GetAll", 0, dbusInterface).Store(&props)
	if err != nil {
		return nil, err
	}

	out := make(map[string]interface{}, len(props))
	for k, v := range props {
		out[k] = v.Value()
	}

	return out, nil
}

<<<<<<< HEAD
// Deprecated: use GetUnitPropertiesContext instead.
||||||| 5e58841cce7
// GetUnitProperties takes the (unescaped) unit name and returns all of its dbus object properties.
=======
// GetUnitProperties takes the (unescaped) unit name and returns all of its dbus object properties.
// Deprecated: use GetUnitPropertiesContext instead
>>>>>>> v1.21.4
func (c *Conn) GetUnitProperties(unit string) (map[string]interface{}, error) {
<<<<<<< HEAD
	return c.GetUnitPropertiesContext(context.Background(), unit)
}

// GetUnitPropertiesContext takes the (unescaped) unit name and returns all of
// its dbus object properties.
func (c *Conn) GetUnitPropertiesContext(ctx context.Context, unit string) (map[string]interface{}, error) {
||||||| 5e58841cce7
=======
	return c.GetUnitPropertiesContext(context.Background(), unit)
}

// GetUnitPropertiesContext same as GetUnitPropertiesContext with context
func (c *Conn) GetUnitPropertiesContext(ctx context.Context, unit string) (map[string]interface{}, error) {
>>>>>>> v1.21.4
	path := unitPath(unit)
	return c.getProperties(ctx, path, "org.freedesktop.systemd1.Unit")
}

<<<<<<< HEAD
// Deprecated: use GetUnitPathPropertiesContext instead.
||||||| 5e58841cce7
// GetUnitPathProperties takes the (escaped) unit path and returns all of its dbus object properties.
=======
// GetUnitPathProperties takes the (escaped) unit path and returns all of its dbus object properties.
// Deprecated: use GetUnitPathPropertiesContext instead
>>>>>>> v1.21.4
func (c *Conn) GetUnitPathProperties(path dbus.ObjectPath) (map[string]interface{}, error) {
<<<<<<< HEAD
	return c.GetUnitPathPropertiesContext(context.Background(), path)
}

// GetUnitPathPropertiesContext takes the (escaped) unit path and returns all
// of its dbus object properties.
func (c *Conn) GetUnitPathPropertiesContext(ctx context.Context, path dbus.ObjectPath) (map[string]interface{}, error) {
	return c.getProperties(ctx, path, "org.freedesktop.systemd1.Unit")
||||||| 5e58841cce7
	return c.getProperties(path, "org.freedesktop.systemd1.Unit")
=======
	return c.GetUnitPathPropertiesContext(context.Background(), path)
}

// GetUnitPathPropertiesContext same as GetUnitPathProperties with context
func (c *Conn) GetUnitPathPropertiesContext(ctx context.Context, path dbus.ObjectPath) (map[string]interface{}, error) {
	return c.getProperties(ctx, path, "org.freedesktop.systemd1.Unit")
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use GetAllPropertiesContext instead.
||||||| 5e58841cce7
// GetAllProperties takes the (unescaped) unit name and returns all of its dbus object properties.
=======
// GetAllProperties takes the (unescaped) unit name and returns all of its dbus object properties.
// Deprecated: use GetAllPropertiesContext instead
>>>>>>> v1.21.4
func (c *Conn) GetAllProperties(unit string) (map[string]interface{}, error) {
<<<<<<< HEAD
	return c.GetAllPropertiesContext(context.Background(), unit)
}

// GetAllPropertiesContext takes the (unescaped) unit name and returns all of
// its dbus object properties.
func (c *Conn) GetAllPropertiesContext(ctx context.Context, unit string) (map[string]interface{}, error) {
||||||| 5e58841cce7
=======
	return c.GetAllPropertiesContext(context.Background(), unit)
}

// GetAllPropertiesContext same as GetAllProperties with context
func (c *Conn) GetAllPropertiesContext(ctx context.Context, unit string) (map[string]interface{}, error) {
>>>>>>> v1.21.4
	path := unitPath(unit)
	return c.getProperties(ctx, path, "")
}

func (c *Conn) getProperty(ctx context.Context, unit string, dbusInterface string, propertyName string) (*Property, error) {
	var err error
	var prop dbus.Variant

	path := unitPath(unit)
	if !path.IsValid() {
		return nil, errors.New("invalid unit name: " + unit)
	}

	obj := c.sysconn.Object("org.freedesktop.systemd1", path)
	err = obj.CallWithContext(ctx, "org.freedesktop.DBus.Properties.Get", 0, dbusInterface, propertyName).Store(&prop)
	if err != nil {
		return nil, err
	}

	return &Property{Name: propertyName, Value: prop}, nil
}

<<<<<<< HEAD
// Deprecated: use GetUnitPropertyContext instead.
||||||| 5e58841cce7
=======
// Deprecated: use GetUnitPropertyContext instead
>>>>>>> v1.21.4
func (c *Conn) GetUnitProperty(unit string, propertyName string) (*Property, error) {
<<<<<<< HEAD
	return c.GetUnitPropertyContext(context.Background(), unit, propertyName)
}

// GetUnitPropertyContext takes an (unescaped) unit name, and a property name,
// and returns the property value.
func (c *Conn) GetUnitPropertyContext(ctx context.Context, unit string, propertyName string) (*Property, error) {
	return c.getProperty(ctx, unit, "org.freedesktop.systemd1.Unit", propertyName)
||||||| 5e58841cce7
	return c.getProperty(unit, "org.freedesktop.systemd1.Unit", propertyName)
=======
	return c.GetUnitPropertyContext(context.Background(), unit, propertyName)
}

// GetUnitPropertyContext same as GetUnitProperty with context
func (c *Conn) GetUnitPropertyContext(ctx context.Context, unit string, propertyName string) (*Property, error) {
	return c.getProperty(ctx, unit, "org.freedesktop.systemd1.Unit", propertyName)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use GetServicePropertyContext instead.
||||||| 5e58841cce7
// GetServiceProperty returns property for given service name and property name
=======
// GetServiceProperty returns property for given service name and property name
// Deprecated: use GetServicePropertyContext instead
>>>>>>> v1.21.4
func (c *Conn) GetServiceProperty(service string, propertyName string) (*Property, error) {
<<<<<<< HEAD
	return c.GetServicePropertyContext(context.Background(), service, propertyName)
}

// GetServiceProperty returns property for given service name and property name.
func (c *Conn) GetServicePropertyContext(ctx context.Context, service string, propertyName string) (*Property, error) {
	return c.getProperty(ctx, service, "org.freedesktop.systemd1.Service", propertyName)
||||||| 5e58841cce7
	return c.getProperty(service, "org.freedesktop.systemd1.Service", propertyName)
=======
	return c.GetServicePropertyContext(context.Background(), service, propertyName)
}

// GetServicePropertyContext same as GetServiceProperty with context
func (c *Conn) GetServicePropertyContext(ctx context.Context, service string, propertyName string) (*Property, error) {
	return c.getProperty(ctx, service, "org.freedesktop.systemd1.Service", propertyName)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use GetUnitTypePropertiesContext instead.
||||||| 5e58841cce7
// GetUnitTypeProperties returns the extra properties for a unit, specific to the unit type.
// Valid values for unitType: Service, Socket, Target, Device, Mount, Automount, Snapshot, Timer, Swap, Path, Slice, Scope
// return "dbus.Error: Unknown interface" if the unitType is not the correct type of the unit
=======
// GetUnitTypeProperties returns the extra properties for a unit, specific to the unit type.
// Valid values for unitType: Service, Socket, Target, Device, Mount, Automount, Snapshot, Timer, Swap, Path, Slice, Scope
// return "dbus.Error: Unknown interface" if the unitType is not the correct type of the unit
// Deprecated: use GetUnitTypePropertiesContext instead
>>>>>>> v1.21.4
func (c *Conn) GetUnitTypeProperties(unit string, unitType string) (map[string]interface{}, error) {
<<<<<<< HEAD
	return c.GetUnitTypePropertiesContext(context.Background(), unit, unitType)
}

// GetUnitTypePropertiesContext returns the extra properties for a unit, specific to the unit type.
// Valid values for unitType: Service, Socket, Target, Device, Mount, Automount, Snapshot, Timer, Swap, Path, Slice, Scope.
// Returns "dbus.Error: Unknown interface" error if the unitType is not the correct type of the unit.
func (c *Conn) GetUnitTypePropertiesContext(ctx context.Context, unit string, unitType string) (map[string]interface{}, error) {
||||||| 5e58841cce7
=======
	return c.GetUnitTypePropertiesContext(context.Background(), unit, unitType)
}

// GetUnitTypePropertiesContext same as GetUnitTypeProperties with context
func (c *Conn) GetUnitTypePropertiesContext(ctx context.Context, unit string, unitType string) (map[string]interface{}, error) {
>>>>>>> v1.21.4
	path := unitPath(unit)
<<<<<<< HEAD
	return c.getProperties(ctx, path, "org.freedesktop.systemd1."+unitType)
}

// Deprecated: use SetUnitPropertiesContext instead.
func (c *Conn) SetUnitProperties(name string, runtime bool, properties ...Property) error {
	return c.SetUnitPropertiesContext(context.Background(), name, runtime, properties...)
||||||| 5e58841cce7
	return c.getProperties(path, "org.freedesktop.systemd1."+unitType)
=======
	return c.getProperties(ctx, path, "org.freedesktop.systemd1."+unitType)
>>>>>>> v1.21.4
}

// SetUnitPropertiesContext may be used to modify certain unit properties at runtime.
// Not all properties may be changed at runtime, but many resource management
// settings (primarily those in systemd.cgroup(5)) may. The changes are applied
// instantly, and stored on disk for future boots, unless runtime is true, in which
// case the settings only apply until the next reboot. name is the name of the unit
// to modify. properties are the settings to set, encoded as an array of property
// name and value pairs.
<<<<<<< HEAD
func (c *Conn) SetUnitPropertiesContext(ctx context.Context, name string, runtime bool, properties ...Property) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.SetUnitProperties", 0, name, runtime, properties).Store()
||||||| 5e58841cce7
func (c *Conn) SetUnitProperties(name string, runtime bool, properties ...Property) error {
	return c.sysobj.Call("org.freedesktop.systemd1.Manager.SetUnitProperties", 0, name, runtime, properties).Store()
=======
// Deprecated: use SetUnitPropertiesContext instead
func (c *Conn) SetUnitProperties(name string, runtime bool, properties ...Property) error {
	return c.SetUnitPropertiesContext(context.Background(), name, runtime, properties...)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use GetUnitTypePropertyContext instead.
||||||| 5e58841cce7
=======
// SetUnitPropertiesContext same as SetUnitProperties with context
func (c *Conn) SetUnitPropertiesContext(ctx context.Context, name string, runtime bool, properties ...Property) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.SetUnitProperties", 0, name, runtime, properties).Store()
}

// Deprecated: use GetUnitTypePropertyContext instead
>>>>>>> v1.21.4
func (c *Conn) GetUnitTypeProperty(unit string, unitType string, propertyName string) (*Property, error) {
<<<<<<< HEAD
	return c.GetUnitTypePropertyContext(context.Background(), unit, unitType, propertyName)
}

// GetUnitTypePropertyContext takes a property name, a unit name, and a unit type,
// and returns a property value. For valid values of unitType, see GetUnitTypePropertiesContext.
func (c *Conn) GetUnitTypePropertyContext(ctx context.Context, unit string, unitType string, propertyName string) (*Property, error) {
	return c.getProperty(ctx, unit, "org.freedesktop.systemd1."+unitType, propertyName)
||||||| 5e58841cce7
	return c.getProperty(unit, "org.freedesktop.systemd1."+unitType, propertyName)
=======
	return c.GetUnitTypePropertyContext(context.Background(), unit, unitType, propertyName)
}

// GetUnitTypePropertyContext same as GetUnitTypeProperty with context
func (c *Conn) GetUnitTypePropertyContext(ctx context.Context, unit string, unitType string, propertyName string) (*Property, error) {
	return c.getProperty(ctx, unit, "org.freedesktop.systemd1."+unitType, propertyName)
>>>>>>> v1.21.4
}

type UnitStatus struct {
	Name        string          // The primary unit name as string
	Description string          // The human readable description string
	LoadState   string          // The load state (i.e. whether the unit file has been loaded successfully)
	ActiveState string          // The active state (i.e. whether the unit is currently started or not)
	SubState    string          // The sub state (a more fine-grained version of the active state that is specific to the unit type, which the active state is not)
	Followed    string          // A unit that is being followed in its state by this unit, if there is any, otherwise the empty string.
	Path        dbus.ObjectPath // The unit object path
	JobId       uint32          // If there is a job queued for the job unit the numeric job id, 0 otherwise
	JobType     string          // The job type as string
	JobPath     dbus.ObjectPath // The job object path
}

type storeFunc func(retvalues ...interface{}) error

func (c *Conn) listUnitsInternal(f storeFunc) ([]UnitStatus, error) {
	result := make([][]interface{}, 0)
	err := f(&result)
	if err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	status := make([]UnitStatus, len(result))
	statusInterface := make([]interface{}, len(status))
	for i := range status {
		statusInterface[i] = &status[i]
	}

	err = dbus.Store(resultInterface, statusInterface...)
	if err != nil {
		return nil, err
	}

	return status, nil
}

// Deprecated: use ListUnitsContext instead.
func (c *Conn) ListUnits() ([]UnitStatus, error) {
	return c.ListUnitsContext(context.Background())
}

// ListUnitsContext returns an array with all currently loaded units. Note that
// units may be known by multiple names at the same time, and hence there might
// be more unit names loaded than actual units behind them.
// Also note that a unit is only loaded if it is active and/or enabled.
// Units that are both disabled and inactive will thus not be returned.
<<<<<<< HEAD
func (c *Conn) ListUnitsContext(ctx context.Context) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnits", 0).Store)
||||||| 5e58841cce7
func (c *Conn) ListUnits() ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.Call("org.freedesktop.systemd1.Manager.ListUnits", 0).Store)
=======
// Deprecated: use ListUnitsContext instead
func (c *Conn) ListUnits() ([]UnitStatus, error) {
	return c.ListUnitsContext(context.Background())
}

// ListUnitsContext same as ListUnits with context
func (c *Conn) ListUnitsContext(ctx context.Context) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnits", 0).Store)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use ListUnitsFilteredContext instead.
||||||| 5e58841cce7
// ListUnitsFiltered returns an array with units filtered by state.
// It takes a list of units' statuses to filter.
=======
// ListUnitsFiltered returns an array with units filtered by state.
// It takes a list of units' statuses to filter.
// Deprecated: use ListUnitsFilteredContext instead
>>>>>>> v1.21.4
func (c *Conn) ListUnitsFiltered(states []string) ([]UnitStatus, error) {
<<<<<<< HEAD
	return c.ListUnitsFilteredContext(context.Background(), states)
||||||| 5e58841cce7
	return c.listUnitsInternal(c.sysobj.Call("org.freedesktop.systemd1.Manager.ListUnitsFiltered", 0, states).Store)
=======
	return c.ListUnitsFilteredContext(context.Background(), states)
}

// ListUnitsFilteredContext same as ListUnitsFiltered with context
func (c *Conn) ListUnitsFilteredContext(ctx context.Context, states []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitsFiltered", 0, states).Store)
>>>>>>> v1.21.4
}

// ListUnitsFilteredContext returns an array with units filtered by state.
// It takes a list of units' statuses to filter.
func (c *Conn) ListUnitsFilteredContext(ctx context.Context, states []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitsFiltered", 0, states).Store)
}

// Deprecated: use ListUnitsByPatternsContext instead.
func (c *Conn) ListUnitsByPatterns(states []string, patterns []string) ([]UnitStatus, error) {
	return c.ListUnitsByPatternsContext(context.Background(), states, patterns)
}

// ListUnitsByPatternsContext returns an array with units.
// It takes a list of units' statuses and names to filter.
// Note that units may be known by multiple names at the same time,
// and hence there might be more unit names loaded than actual units behind them.
<<<<<<< HEAD
func (c *Conn) ListUnitsByPatternsContext(ctx context.Context, states []string, patterns []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitsByPatterns", 0, states, patterns).Store)
||||||| 5e58841cce7
func (c *Conn) ListUnitsByPatterns(states []string, patterns []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.Call("org.freedesktop.systemd1.Manager.ListUnitsByPatterns", 0, states, patterns).Store)
=======
// Deprecated: use ListUnitsByPatternsContext instead
func (c *Conn) ListUnitsByPatterns(states []string, patterns []string) ([]UnitStatus, error) {
	return c.ListUnitsByPatternsContext(context.Background(), states, patterns)
}

// ListUnitsByPatternsContext same as ListUnitsByPatterns with context
func (c *Conn) ListUnitsByPatternsContext(ctx context.Context, states []string, patterns []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitsByPatterns", 0, states, patterns).Store)
>>>>>>> v1.21.4
}

// Deprecated: use ListUnitsByNamesContext instead.
func (c *Conn) ListUnitsByNames(units []string) ([]UnitStatus, error) {
	return c.ListUnitsByNamesContext(context.Background(), units)
}

// ListUnitsByNamesContext returns an array with units. It takes a list of units'
// names and returns an UnitStatus array. Comparing to ListUnitsByPatternsContext
// method, this method returns statuses even for inactive or non-existing
// units. Input array should contain exact unit names, but not patterns.
<<<<<<< HEAD
//
// Requires systemd v230 or higher.
func (c *Conn) ListUnitsByNamesContext(ctx context.Context, units []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitsByNames", 0, units).Store)
||||||| 5e58841cce7
// Note: Requires systemd v230 or higher
func (c *Conn) ListUnitsByNames(units []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.Call("org.freedesktop.systemd1.Manager.ListUnitsByNames", 0, units).Store)
=======
// Note: Requires systemd v230 or higher
// Deprecated: use ListUnitsByNamesContext instead
func (c *Conn) ListUnitsByNames(units []string) ([]UnitStatus, error) {
	return c.ListUnitsByNamesContext(context.Background(), units)
}

// ListUnitsByNamesContext same as ListUnitsByNames with context
func (c *Conn) ListUnitsByNamesContext(ctx context.Context, units []string) ([]UnitStatus, error) {
	return c.listUnitsInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitsByNames", 0, units).Store)
>>>>>>> v1.21.4
}

type UnitFile struct {
	Path string
	Type string
}

func (c *Conn) listUnitFilesInternal(f storeFunc) ([]UnitFile, error) {
	result := make([][]interface{}, 0)
	err := f(&result)
	if err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	files := make([]UnitFile, len(result))
	fileInterface := make([]interface{}, len(files))
	for i := range files {
		fileInterface[i] = &files[i]
	}

	err = dbus.Store(resultInterface, fileInterface...)
	if err != nil {
		return nil, err
	}

	return files, nil
}

<<<<<<< HEAD
// Deprecated: use ListUnitFilesContext instead.
||||||| 5e58841cce7
// ListUnitFiles returns an array of all available units on disk.
=======
// ListUnitFiles returns an array of all available units on disk.
// Deprecated: use ListUnitFilesContext instead
>>>>>>> v1.21.4
func (c *Conn) ListUnitFiles() ([]UnitFile, error) {
<<<<<<< HEAD
	return c.ListUnitFilesContext(context.Background())
}

// ListUnitFiles returns an array of all available units on disk.
func (c *Conn) ListUnitFilesContext(ctx context.Context) ([]UnitFile, error) {
	return c.listUnitFilesInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitFiles", 0).Store)
||||||| 5e58841cce7
	return c.listUnitFilesInternal(c.sysobj.Call("org.freedesktop.systemd1.Manager.ListUnitFiles", 0).Store)
=======
	return c.ListUnitFilesContext(context.Background())
}

// ListUnitFilesContext same as ListUnitFiles with context
func (c *Conn) ListUnitFilesContext(ctx context.Context) ([]UnitFile, error) {
	return c.listUnitFilesInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitFiles", 0).Store)
>>>>>>> v1.21.4
}

<<<<<<< HEAD
// Deprecated: use ListUnitFilesByPatternsContext instead.
||||||| 5e58841cce7
// ListUnitFilesByPatterns returns an array of all available units on disk matched the patterns.
=======
// ListUnitFilesByPatterns returns an array of all available units on disk matched the patterns.
// Deprecated: use ListUnitFilesByPatternsContext instead
>>>>>>> v1.21.4
func (c *Conn) ListUnitFilesByPatterns(states []string, patterns []string) ([]UnitFile, error) {
<<<<<<< HEAD
	return c.ListUnitFilesByPatternsContext(context.Background(), states, patterns)
}

// ListUnitFilesByPatternsContext returns an array of all available units on disk matched the patterns.
func (c *Conn) ListUnitFilesByPatternsContext(ctx context.Context, states []string, patterns []string) ([]UnitFile, error) {
	return c.listUnitFilesInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitFilesByPatterns", 0, states, patterns).Store)
||||||| 5e58841cce7
	return c.listUnitFilesInternal(c.sysobj.Call("org.freedesktop.systemd1.Manager.ListUnitFilesByPatterns", 0, states, patterns).Store)
=======
	return c.ListUnitFilesByPatternsContext(context.Background(), states, patterns)
}

// ListUnitFilesByPatternsContext same as ListUnitFilesByPatterns with context
func (c *Conn) ListUnitFilesByPatternsContext(ctx context.Context, states []string, patterns []string) ([]UnitFile, error) {
	return c.listUnitFilesInternal(c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListUnitFilesByPatterns", 0, states, patterns).Store)
>>>>>>> v1.21.4
}

type LinkUnitFileChange EnableUnitFileChange

// Deprecated: use LinkUnitFilesContext instead.
func (c *Conn) LinkUnitFiles(files []string, runtime bool, force bool) ([]LinkUnitFileChange, error) {
	return c.LinkUnitFilesContext(context.Background(), files, runtime, force)
}

// LinkUnitFilesContext links unit files (that are located outside of the
// usual unit search paths) into the unit search path.
//
// It takes a list of absolute paths to unit files to link and two
// booleans.
//
// The first boolean controls whether the unit shall be
// enabled for runtime only (true, /run), or persistently (false,
// /etc).
//
// The second controls whether symlinks pointing to other units shall
// be replaced if necessary.
//
// This call returns a list of the changes made. The list consists of
// structures with three strings: the type of the change (one of symlink
// or unlink), the file name of the symlink and the destination of the
// symlink.
<<<<<<< HEAD
func (c *Conn) LinkUnitFilesContext(ctx context.Context, files []string, runtime bool, force bool) ([]LinkUnitFileChange, error) {
||||||| 5e58841cce7
func (c *Conn) LinkUnitFiles(files []string, runtime bool, force bool) ([]LinkUnitFileChange, error) {
=======
// Deprecated: use LinkUnitFilesContext instead
func (c *Conn) LinkUnitFiles(files []string, runtime bool, force bool) ([]LinkUnitFileChange, error) {
	return c.LinkUnitFilesContext(context.Background(), files, runtime, force)
}

// LinkUnitFilesContext same as LinkUnitFiles with context
func (c *Conn) LinkUnitFilesContext(ctx context.Context, files []string, runtime bool, force bool) ([]LinkUnitFileChange, error) {
>>>>>>> v1.21.4
	result := make([][]interface{}, 0)
	err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.LinkUnitFiles", 0, files, runtime, force).Store(&result)
	if err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	changes := make([]LinkUnitFileChange, len(result))
	changesInterface := make([]interface{}, len(changes))
	for i := range changes {
		changesInterface[i] = &changes[i]
	}

	err = dbus.Store(resultInterface, changesInterface...)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

// Deprecated: use EnableUnitFilesContext instead.
func (c *Conn) EnableUnitFiles(files []string, runtime bool, force bool) (bool, []EnableUnitFileChange, error) {
	return c.EnableUnitFilesContext(context.Background(), files, runtime, force)
}

// EnableUnitFilesContext may be used to enable one or more units in the system
// (by creating symlinks to them in /etc or /run).
//
// It takes a list of unit files to enable (either just file names or full
// absolute paths if the unit files are residing outside the usual unit
// search paths), and two booleans: the first controls whether the unit shall
// be enabled for runtime only (true, /run), or persistently (false, /etc).
// The second one controls whether symlinks pointing to other units shall
// be replaced if necessary.
//
// This call returns one boolean and an array with the changes made. The
// boolean signals whether the unit files contained any enablement
// information (i.e. an [Install]) section. The changes list consists of
// structures with three strings: the type of the change (one of symlink
// or unlink), the file name of the symlink and the destination of the
// symlink.
<<<<<<< HEAD
func (c *Conn) EnableUnitFilesContext(ctx context.Context, files []string, runtime bool, force bool) (bool, []EnableUnitFileChange, error) {
||||||| 5e58841cce7
func (c *Conn) EnableUnitFiles(files []string, runtime bool, force bool) (bool, []EnableUnitFileChange, error) {
=======
// Deprecated: use EnableUnitFilesContext instead
func (c *Conn) EnableUnitFiles(files []string, runtime bool, force bool) (bool, []EnableUnitFileChange, error) {
	return c.EnableUnitFilesContext(context.Background(), files, runtime, force)
}

// EnableUnitFilesContext same as EnableUnitFiles with context
func (c *Conn) EnableUnitFilesContext(ctx context.Context, files []string, runtime bool, force bool) (bool, []EnableUnitFileChange, error) {
>>>>>>> v1.21.4
	var carries_install_info bool

	result := make([][]interface{}, 0)
	err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.EnableUnitFiles", 0, files, runtime, force).Store(&carries_install_info, &result)
	if err != nil {
		return false, nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	changes := make([]EnableUnitFileChange, len(result))
	changesInterface := make([]interface{}, len(changes))
	for i := range changes {
		changesInterface[i] = &changes[i]
	}

	err = dbus.Store(resultInterface, changesInterface...)
	if err != nil {
		return false, nil, err
	}

	return carries_install_info, changes, nil
}

type EnableUnitFileChange struct {
	Type        string // Type of the change (one of symlink or unlink)
	Filename    string // File name of the symlink
	Destination string // Destination of the symlink
}

// Deprecated: use DisableUnitFilesContext instead.
func (c *Conn) DisableUnitFiles(files []string, runtime bool) ([]DisableUnitFileChange, error) {
	return c.DisableUnitFilesContext(context.Background(), files, runtime)
}

// DisableUnitFilesContext may be used to disable one or more units in the
// system (by removing symlinks to them from /etc or /run).
//
// It takes a list of unit files to disable (either just file names or full
// absolute paths if the unit files are residing outside the usual unit
// search paths), and one boolean: whether the unit was enabled for runtime
// only (true, /run), or persistently (false, /etc).
//
// This call returns an array with the changes made. The changes list
// consists of structures with three strings: the type of the change (one of
// symlink or unlink), the file name of the symlink and the destination of the
// symlink.
<<<<<<< HEAD
func (c *Conn) DisableUnitFilesContext(ctx context.Context, files []string, runtime bool) ([]DisableUnitFileChange, error) {
||||||| 5e58841cce7
func (c *Conn) DisableUnitFiles(files []string, runtime bool) ([]DisableUnitFileChange, error) {
=======
// Deprecated: use DisableUnitFilesContext instead
func (c *Conn) DisableUnitFiles(files []string, runtime bool) ([]DisableUnitFileChange, error) {
	return c.DisableUnitFilesContext(context.Background(), files, runtime)
}

// DisableUnitFilesContext same as DisableUnitFiles with context
func (c *Conn) DisableUnitFilesContext(ctx context.Context, files []string, runtime bool) ([]DisableUnitFileChange, error) {
>>>>>>> v1.21.4
	result := make([][]interface{}, 0)
	err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.DisableUnitFiles", 0, files, runtime).Store(&result)
	if err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	changes := make([]DisableUnitFileChange, len(result))
	changesInterface := make([]interface{}, len(changes))
	for i := range changes {
		changesInterface[i] = &changes[i]
	}

	err = dbus.Store(resultInterface, changesInterface...)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

type DisableUnitFileChange struct {
	Type        string // Type of the change (one of symlink or unlink)
	Filename    string // File name of the symlink
	Destination string // Destination of the symlink
}

<<<<<<< HEAD
// Deprecated: use MaskUnitFilesContext instead.
||||||| 5e58841cce7
// MaskUnitFiles masks one or more units in the system
//
// It takes three arguments:
//   * list of units to mask (either just file names or full
//     absolute paths if the unit files are residing outside
//     the usual unit search paths)
//   * runtime to specify whether the unit was enabled for runtime
//     only (true, /run/systemd/..), or persistently (false, /etc/systemd/..)
//   * force flag
=======
// MaskUnitFiles masks one or more units in the system
//
// It takes three arguments:
//   * list of units to mask (either just file names or full
//     absolute paths if the unit files are residing outside
//     the usual unit search paths)
//   * runtime to specify whether the unit was enabled for runtime
//     only (true, /run/systemd/..), or persistently (false, /etc/systemd/..)
//   * force flag
// Deprecated: use MaskUnitFilesContext instead
>>>>>>> v1.21.4
func (c *Conn) MaskUnitFiles(files []string, runtime bool, force bool) ([]MaskUnitFileChange, error) {
<<<<<<< HEAD
	return c.MaskUnitFilesContext(context.Background(), files, runtime, force)
}

// MaskUnitFilesContext masks one or more units in the system.
//
// The files argument contains a  list of units to mask (either just file names
// or full absolute paths if the unit files are residing outside the usual unit
// search paths).
//
// The runtime argument is used to specify whether the unit was enabled for
// runtime only (true, /run/systemd/..), or persistently (false,
// /etc/systemd/..).
func (c *Conn) MaskUnitFilesContext(ctx context.Context, files []string, runtime bool, force bool) ([]MaskUnitFileChange, error) {
||||||| 5e58841cce7
=======
	return c.MaskUnitFilesContext(context.Background(), files, runtime, force)
}

// MaskUnitFilesContext same as MaskUnitFiles with context
func (c *Conn) MaskUnitFilesContext(ctx context.Context, files []string, runtime bool, force bool) ([]MaskUnitFileChange, error) {
>>>>>>> v1.21.4
	result := make([][]interface{}, 0)
	err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.MaskUnitFiles", 0, files, runtime, force).Store(&result)
	if err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	changes := make([]MaskUnitFileChange, len(result))
	changesInterface := make([]interface{}, len(changes))
	for i := range changes {
		changesInterface[i] = &changes[i]
	}

	err = dbus.Store(resultInterface, changesInterface...)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

type MaskUnitFileChange struct {
	Type        string // Type of the change (one of symlink or unlink)
	Filename    string // File name of the symlink
	Destination string // Destination of the symlink
}

<<<<<<< HEAD
// Deprecated: use UnmaskUnitFilesContext instead.
||||||| 5e58841cce7
// UnmaskUnitFiles unmasks one or more units in the system
//
// It takes two arguments:
//   * list of unit files to mask (either just file names or full
//     absolute paths if the unit files are residing outside
//     the usual unit search paths)
//   * runtime to specify whether the unit was enabled for runtime
//     only (true, /run/systemd/..), or persistently (false, /etc/systemd/..)
=======
// UnmaskUnitFiles unmasks one or more units in the system
//
// It takes two arguments:
//   * list of unit files to mask (either just file names or full
//     absolute paths if the unit files are residing outside
//     the usual unit search paths)
//   * runtime to specify whether the unit was enabled for runtime
//     only (true, /run/systemd/..), or persistently (false, /etc/systemd/..)
// Deprecated: use UnmaskUnitFilesContext instead
>>>>>>> v1.21.4
func (c *Conn) UnmaskUnitFiles(files []string, runtime bool) ([]UnmaskUnitFileChange, error) {
<<<<<<< HEAD
	return c.UnmaskUnitFilesContext(context.Background(), files, runtime)
}

// UnmaskUnitFilesContext unmasks one or more units in the system.
//
// It takes the list of unit files to mask (either just file names or full
// absolute paths if the unit files are residing outside the usual unit search
// paths), and a boolean runtime flag to specify whether the unit was enabled
// for runtime only (true, /run/systemd/..), or persistently (false,
// /etc/systemd/..).
func (c *Conn) UnmaskUnitFilesContext(ctx context.Context, files []string, runtime bool) ([]UnmaskUnitFileChange, error) {
||||||| 5e58841cce7
=======
	return c.UnmaskUnitFilesContext(context.Background(), files, runtime)
}

// UnmaskUnitFilesContext same as UnmaskUnitFiles with context
func (c *Conn) UnmaskUnitFilesContext(ctx context.Context, files []string, runtime bool) ([]UnmaskUnitFileChange, error) {
>>>>>>> v1.21.4
	result := make([][]interface{}, 0)
	err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.UnmaskUnitFiles", 0, files, runtime).Store(&result)
	if err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	changes := make([]UnmaskUnitFileChange, len(result))
	changesInterface := make([]interface{}, len(changes))
	for i := range changes {
		changesInterface[i] = &changes[i]
	}

	err = dbus.Store(resultInterface, changesInterface...)
	if err != nil {
		return nil, err
	}

	return changes, nil
}

type UnmaskUnitFileChange struct {
	Type        string // Type of the change (one of symlink or unlink)
	Filename    string // File name of the symlink
	Destination string // Destination of the symlink
}

<<<<<<< HEAD
// Deprecated: use ReloadContext instead.
||||||| 5e58841cce7
// Reload instructs systemd to scan for and reload unit files. This is
// equivalent to a 'systemctl daemon-reload'.
=======
// Reload instructs systemd to scan for and reload unit files. This is
// equivalent to a 'systemctl daemon-reload'.
// Deprecated: use ReloadContext instead
>>>>>>> v1.21.4
func (c *Conn) Reload() error {
<<<<<<< HEAD
	return c.ReloadContext(context.Background())
}

// ReloadContext instructs systemd to scan for and reload unit files. This is
// an equivalent to systemctl daemon-reload.
func (c *Conn) ReloadContext(ctx context.Context) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.Reload", 0).Store()
||||||| 5e58841cce7
	return c.sysobj.Call("org.freedesktop.systemd1.Manager.Reload", 0).Store()
=======
	return c.ReloadContext(context.Background())
}

// ReloadContext same as Reload with context
func (c *Conn) ReloadContext(ctx context.Context) error {
	return c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.Reload", 0).Store()
>>>>>>> v1.21.4
}

func unitPath(name string) dbus.ObjectPath {
	return dbus.ObjectPath("/org/freedesktop/systemd1/unit/" + PathBusEscape(name))
}

// unitName returns the unescaped base element of the supplied escaped path.
func unitName(dpath dbus.ObjectPath) string {
	return pathBusUnescape(path.Base(string(dpath)))
}
<<<<<<< HEAD

// JobStatus holds a currently queued job definition.
type JobStatus struct {
	Id       uint32          // The numeric job id
	Unit     string          // The primary unit name for this job
	JobType  string          // The job type as string
	Status   string          // The job state as string
	JobPath  dbus.ObjectPath // The job object path
	UnitPath dbus.ObjectPath // The unit object path
}

// Deprecated: use ListJobsContext instead.
func (c *Conn) ListJobs() ([]JobStatus, error) {
	return c.ListJobsContext(context.Background())
}

// ListJobsContext returns an array with all currently queued jobs.
func (c *Conn) ListJobsContext(ctx context.Context) ([]JobStatus, error) {
	return c.listJobsInternal(ctx)
}

func (c *Conn) listJobsInternal(ctx context.Context) ([]JobStatus, error) {
	result := make([][]interface{}, 0)
	if err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListJobs", 0).Store(&result); err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	status := make([]JobStatus, len(result))
	statusInterface := make([]interface{}, len(status))
	for i := range status {
		statusInterface[i] = &status[i]
	}

	if err := dbus.Store(resultInterface, statusInterface...); err != nil {
		return nil, err
	}

	return status, nil
}
||||||| 5e58841cce7
=======

// Currently queued job definition
type JobStatus struct {
	Id       uint32          // The numeric job id
	Unit     string          // The primary unit name for this job
	JobType  string          // The job type as string
	Status   string          // The job state as string
	JobPath  dbus.ObjectPath // The job object path
	UnitPath dbus.ObjectPath // The unit object path
}

// ListJobs returns an array with all currently queued jobs
// Deprecated: use ListJobsContext instead
func (c *Conn) ListJobs() ([]JobStatus, error) {
	return c.ListJobsContext(context.Background())
}

// ListJobsContext same as ListJobs with context
func (c *Conn) ListJobsContext(ctx context.Context) ([]JobStatus, error) {
	return c.listJobsInternal(ctx)
}

func (c *Conn) listJobsInternal(ctx context.Context) ([]JobStatus, error) {
	result := make([][]interface{}, 0)
	if err := c.sysobj.CallWithContext(ctx, "org.freedesktop.systemd1.Manager.ListJobs", 0).Store(&result); err != nil {
		return nil, err
	}

	resultInterface := make([]interface{}, len(result))
	for i := range result {
		resultInterface[i] = result[i]
	}

	status := make([]JobStatus, len(result))
	statusInterface := make([]interface{}, len(status))
	for i := range status {
		statusInterface[i] = &status[i]
	}

	if err := dbus.Store(resultInterface, statusInterface...); err != nil {
		return nil, err
	}

	return status, nil
}
>>>>>>> v1.21.4
