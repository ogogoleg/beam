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
import 'package:playground/modules/editor/components/share_dropdown/link_text_field.dart';
import 'package:playground/modules/editor/components/share_dropdown/share_link_field.dart';
import 'package:playground/pages/playground/states/examples_state.dart';
import 'package:playground/pages/playground/states/playground_state.dart';
import 'package:provider/provider.dart';

const kLoadingIndicatorSize = 20.0;

class ShareDropdownBody extends StatefulWidget {
  final VoidCallback close;

  const ShareDropdownBody({super.key, required this.close});

  @override
  State<ShareDropdownBody> createState() => _ShareDropdownBodyState();
}

class _ShareDropdownBodyState extends State<ShareDropdownBody> {
  final TextEditingController textEditingController = TextEditingController();
  bool isPressed = false;

  @override
  Widget build(BuildContext context) {
    AppLocalizations appLocale = AppLocalizations.of(context)!;

    return Consumer2<ExampleState, PlaygroundState>(
      builder: (context, exampleState, playgroundState, child) => Padding(
        padding: const EdgeInsets.symmetric(
          vertical: kXlSpacing,
          horizontal: kXlSpacing,
        ),
        child: Center(
          child: isPressed
              ? FutureBuilder(
                  builder: (context, snapshot) {
                    return _buildButton(playgroundState, snapshot, appLocale);
                  },
                )
              : Column(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    Text(appLocale.clickForLink),
                    GestureDetector(
                      onTap: () async {
                        setState(() {
                          isPressed = true;
                        });
                      },
                      child: ShareLinkField(
                        isPressed: isPressed,
                        child: Text(
                          appLocale.showAndCopyLink,
                          style: TextStyle(
                            fontSize: kLabelFontSize,
                            fontWeight: kBoldWeight,
                            color: ThemeColors.of(context).primary,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
        ),
      ),
    );
  }

  Widget _buildButton(
    PlaygroundState playgroundState,
    AsyncSnapshot snapshot,
    AppLocalizations appLocale,
  ) {
    if (playgroundState.isExampleChanged) {
      if (snapshot.hasData) {
        setAndCopyLinkText(snapshot.data.toString());

        return Column(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            Text(appLocale.linkCopied),
            ShareLinkField(
              isPressed: isPressed,
              child: LinkTextField(
                textEditingController: textEditingController,
              ),
            ),
          ],
        );
      } else {
        return Column(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            Text(appLocale.loadingLink),
            ShareLinkField(
              isPressed: isPressed,
              child: SizedBox(
                height: kLoadingIndicatorSize,
                width: kLoadingIndicatorSize,
                child: CircularProgressIndicator(
                  color: ThemeColors.of(context).primary,
                ),
              ),
            ),
          ],
        );
      }
    } else {
      setAndCopyLinkText(
        '${Uri.base.toString().split('?')[0]}?$kExampleParam=${playgroundState.selectedExample!.path}',
      );

      return Column(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: [
          Text(appLocale.linkCopied),
          ShareLinkField(
            isPressed: isPressed,
            child: LinkTextField(
              textEditingController: textEditingController,
            ),
          ),
        ],
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
