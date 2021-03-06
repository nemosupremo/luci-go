# Copyright 2018 The LUCI Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

load("go.chromium.org/luci/starlark/starlarkproto/testprotos/test.proto", "testprotos")

m = testprotos.Complex()

# Oneof alternatives have no defaults.
assert.eq(m.simple, None)
assert.eq(m.inner_msg, None)

# Setting one alternative resets the other (and doesn't touch other fields).
m.i64 = 123
m.simple = testprotos.Simple()
assert.true(m.simple != None)
assert.true(m.inner_msg == None)
assert.eq(m.i64, 123)
m.inner_msg = testprotos.Complex.InnerMessage()
assert.true(m.simple == None)
assert.true(m.inner_msg != None)
assert.eq(m.i64, 123)

# Setting a "picked" alternative to None resets it.
assert.true(m.inner_msg != None)
m.inner_msg = None
assert.true(m.inner_msg == None)

# Setting some other alternative to None does nothing.
m.simple = testprotos.Simple()
m.inner_msg = None
assert.true(m.simple != None)

# In constructors the last kwarg wins (starlark dicts preserve order).
m2 = testprotos.Complex(
    simple=testprotos.Simple(),
    inner_msg=testprotos.Complex.InnerMessage())
assert.true(m2.simple == None)
assert.true(m2.inner_msg != None)
m3 = testprotos.Complex(
    inner_msg=testprotos.Complex.InnerMessage(),
    simple=testprotos.Simple())
assert.true(m3.simple != None)
assert.true(m3.inner_msg == None)

# Serialization works.
assert.eq(
    proto.to_pbtext(testprotos.Complex(simple=testprotos.Simple(i=1))),
    "simple: <\n  i: 1\n>\n")
