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
import 'package:flutter/services.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:playground/config/theme.dart';
import 'package:playground/constants/font_weight.dart';
import 'package:playground/constants/params.dart';
import 'package:playground/constants/sizes.dart';
import 'package:playground/modules/examples/repositories/models/shared_file_model.dart';
import 'package:playground/pages/playground/states/examples_state.dart';
import 'package:playground/pages/playground/states/playground_state.dart';

class ShareLinkField extends StatelessWidget {
  final bool isPressed;
  final ExampleState exampleState;
  final PlaygroundState playgroundState;
  final TextEditingController textEditingController = TextEditingController();

  ShareLinkField({
    Key? key,
    required this.isPressed,
    required this.exampleState,
    required this.playgroundState,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    AppLocalizations appLocale = AppLocalizations.of(context)!;

    if (isPressed && playgroundState.isExampleChanged) {
      return FutureBuilder(
        future: exampleState.getShareLink(
          [SharedFile(playgroundState.source, true, '')],
          playgroundState.sdk,
          playgroundState.pipelineOptions,
        ),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            setAndCopyLinkText(snapshot.data.toString());

            return TextFormField(
              controller: textEditingController,
              decoration: const InputDecoration(
                contentPadding: EdgeInsets.symmetric(vertical: kSmSpacing),
                border: InputBorder.none,
                isCollapsed: true,
              ),
              readOnly: true,
              style: TextStyle(
                fontSize: kLabelFontSize,
                fontWeight: kNormalWeight,
                color: ThemeColors.of(context).primary,
              ),
            );
          } else {
            return SizedBox(
              height: kContainerHeight / 2,
              width: kContainerHeight / 2,
              child: CircularProgressIndicator(
                color: ThemeColors.of(context).primary,
              ),
            );
          }
        },
      );
    } else if (isPressed && !playgroundState.isExampleChanged) {
      setAndCopyLinkText(
        '${Uri.base.toString()}?$kExampleParam=${playgroundState.selectedExample!.path}',
      );

      return TextFormField(
        controller: textEditingController,
        decoration: const InputDecoration(
          contentPadding: EdgeInsets.symmetric(vertical: kSmSpacing),
          border: InputBorder.none,
          isCollapsed: true,
        ),
        readOnly: true,
        style: TextStyle(
          fontSize: kLabelFontSize,
          fontWeight: kNormalWeight,
          color: ThemeColors.of(context).primary,
        ),
      );
    } else {
      return Text(
        appLocale.showAndCopyLink,
        style: TextStyle(
          fontSize: kLabelFontSize,
          fontWeight: kBoldWeight,
          color: ThemeColors.of(context).primary,
        ),
      );
    }
  }

  void setAndCopyLinkText(String link) async {
    textEditingController.text = link;
    await Clipboard.setData(ClipboardData(
      text: textEditingController.text,
    ));
  }
}
