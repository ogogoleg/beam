// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by starcgen. DO NOT EDIT.
// File: filter.shims.go

package filter

import (
	"context"
	"reflect"

	// Library imports
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/exec"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/graphx/schema"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/sdf"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/reflectx"
)

func init() {
	runtime.RegisterFunction(mapFn)
	runtime.RegisterFunction(mergeFn)
	runtime.RegisterType(reflect.TypeOf((*filterFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*filterFn)(nil)).Elem())
	reflectx.RegisterStructWrapper(reflect.TypeOf((*filterFn)(nil)).Elem(), wrapMakerFilterFn)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int, int) int)(nil)).Elem(), funcMakerIntIntГInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func(typex.T, func(typex.T)))(nil)).Elem(), funcMakerTypex۰TEmitTypex۰TГ)
	reflectx.RegisterFunc(reflect.TypeOf((*func(typex.T) (typex.T, int))(nil)).Elem(), funcMakerTypex۰TГTypex۰TInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func())(nil)).Elem(), funcMakerГ)
	exec.RegisterEmitter(reflect.TypeOf((*func(typex.T))(nil)).Elem(), emitMakerTypex۰T)
}

func wrapMakerFilterFn(fn interface{}) map[string]reflectx.Func {
	dfn := fn.(*filterFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 typex.T, a1 func(typex.T)) { dfn.ProcessElement(a0, a1) }),
		"Setup":          reflectx.MakeFunc(func() { dfn.Setup() }),
	}
}

type callerIntIntГInt struct {
	fn func(int, int) int
}

func funcMakerIntIntГInt(fn interface{}) reflectx.Func {
	f := fn.(func(int, int) int)
	return &callerIntIntГInt{fn: f}
}

func (c *callerIntIntГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntIntГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntIntГInt) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(int), args[1].(int))
	return []interface{}{out0}
}

func (c *callerIntIntГInt) Call2x1(arg0, arg1 interface{}) interface{} {
	return c.fn(arg0.(int), arg1.(int))
}

type callerTypex۰TEmitTypex۰TГ struct {
	fn func(typex.T, func(typex.T))
}

func funcMakerTypex۰TEmitTypex۰TГ(fn interface{}) reflectx.Func {
	f := fn.(func(typex.T, func(typex.T)))
	return &callerTypex۰TEmitTypex۰TГ{fn: f}
}

func (c *callerTypex۰TEmitTypex۰TГ) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerTypex۰TEmitTypex۰TГ) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerTypex۰TEmitTypex۰TГ) Call(args []interface{}) []interface{} {
	c.fn(args[0].(typex.T), args[1].(func(typex.T)))
	return []interface{}{}
}

func (c *callerTypex۰TEmitTypex۰TГ) Call2x0(arg0, arg1 interface{}) {
	c.fn(arg0.(typex.T), arg1.(func(typex.T)))
}

type callerTypex۰TГTypex۰TInt struct {
	fn func(typex.T) (typex.T, int)
}

func funcMakerTypex۰TГTypex۰TInt(fn interface{}) reflectx.Func {
	f := fn.(func(typex.T) (typex.T, int))
	return &callerTypex۰TГTypex۰TInt{fn: f}
}

func (c *callerTypex۰TГTypex۰TInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerTypex۰TГTypex۰TInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerTypex۰TГTypex۰TInt) Call(args []interface{}) []interface{} {
	out0, out1 := c.fn(args[0].(typex.T))
	return []interface{}{out0, out1}
}

func (c *callerTypex۰TГTypex۰TInt) Call1x2(arg0 interface{}) (interface{}, interface{}) {
	return c.fn(arg0.(typex.T))
}

type callerГ struct {
	fn func()
}

func funcMakerГ(fn interface{}) reflectx.Func {
	f := fn.(func())
	return &callerГ{fn: f}
}

func (c *callerГ) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerГ) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerГ) Call(args []interface{}) []interface{} {
	c.fn()
	return []interface{}{}
}

func (c *callerГ) Call0x0() {
	c.fn()
}

type emitNative struct {
	n   exec.ElementProcessor
	fn  interface{}
	est *sdf.WatermarkEstimator

	ctx   context.Context
	ws    []typex.Window
	et    typex.EventTime
	value exec.FullValue
}

func (e *emitNative) Init(ctx context.Context, ws []typex.Window, et typex.EventTime) error {
	e.ctx = ctx
	e.ws = ws
	e.et = et
	return nil
}

func (e *emitNative) Value() interface{} {
	return e.fn
}

func (e *emitNative) AttachEstimator(est *sdf.WatermarkEstimator) {
	e.est = est
}

func emitMakerTypex۰T(n exec.ElementProcessor) exec.ReusableEmitter {
	ret := &emitNative{n: n}
	ret.fn = ret.invokeTypex۰T
	return ret
}

func (e *emitNative) invokeTypex۰T(val typex.T) {
	e.value = exec.FullValue{Windows: e.ws, Timestamp: e.et, Elm: val}
	if e.est != nil {
		(*e.est).(sdf.TimestampObservingEstimator).ObserveTimestamp(e.et.ToTime())
	}
	if err := e.n.ProcessElement(e.ctx, &e.value); err != nil {
		panic(err)
	}
}

// DO NOT MODIFY: GENERATED CODE
