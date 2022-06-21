/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import 'package:flutter/material.dart';
import 'package:playground/components/dropdown_button/dropdown_button.dart';

const double kAdditionalDyAlignment = 10.0;

Offset findDropdownOffset({
  required GlobalKey selectorKey,
  DropdownAlignment alignment = DropdownAlignment.left,
  double widgetWidth = 0,
}) {
  RenderBox? rBox = selectorKey.currentContext?.findRenderObject() as RenderBox;

  switch (alignment) {
    case DropdownAlignment.left:
      return Offset(
        rBox.localToGlobal(Offset.zero).dx,
        rBox.localToGlobal(Offset.zero).dy +
            rBox.size.height +
            kAdditionalDyAlignment,
      );
    case DropdownAlignment.right:
      return Offset(
        rBox.localToGlobal(Offset.zero).dx - (widgetWidth - rBox.size.width),
        rBox.localToGlobal(Offset.zero).dy +
            rBox.size.height +
            kAdditionalDyAlignment,
      );
  }
}
