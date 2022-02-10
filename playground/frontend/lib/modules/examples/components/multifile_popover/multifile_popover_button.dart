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

import 'package:aligned_dialog/aligned_dialog.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:playground/constants/assets.dart';
import 'package:playground/constants/sizes.dart';
import 'package:playground/modules/examples/components/multifile_popover/multifile_popover.dart';
import 'package:playground/modules/examples/models/example_model.dart';

class MultifilePopoverButton extends StatelessWidget {
  final BuildContext? parentContext;
  final ExampleModel example;
  final Alignment followerAnchor;
  final Alignment targetAnchor;

  const MultifilePopoverButton({
    Key? key,
    this.parentContext,
    required this.example,
    required this.followerAnchor,
    required this.targetAnchor,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return IconButton(
      iconSize: kIconSizeMd,
      splashRadius: kIconButtonSplashRadius,
      icon: SvgPicture.asset(kMultifileIconAsset),
      onPressed: () {
        _showMultifilePopover(
          parentContext ?? context,
          example,
          followerAnchor,
          targetAnchor,
        );
      },
    );
  }

  void _showMultifilePopover(
    BuildContext context,
    ExampleModel example,
    Alignment followerAnchor,
    Alignment targetAnchor,
  ) {
    // close previous dialogs
    Navigator.of(context, rootNavigator: true).popUntil((route) {
      return route.isFirst;
    });
    showAlignedDialog(
      context: context,
      builder: (dialogContext) => MultifilePopover(
        example: example,
      ),
      followerAnchor: followerAnchor,
      targetAnchor: targetAnchor,
      barrierColor: Colors.transparent,
    );
  }
}
