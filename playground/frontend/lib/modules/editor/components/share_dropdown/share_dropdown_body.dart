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
import 'package:playground/constants/sizes.dart';

class ShareDropdownBody extends StatefulWidget {
  final Function close;

  const ShareDropdownBody({
    Key? key,
    required this.close,
  }) : super(key: key);

  @override
  State<ShareDropdownBody> createState() => _ShareDropdownBodyState();
}

class _ShareDropdownBodyState extends State<ShareDropdownBody> {
  bool isPressed = false;

  @override
  Widget build(BuildContext context) {
    AppLocalizations appLocale = AppLocalizations.of(context)!;

    return Padding(
      padding: const EdgeInsets.symmetric(
        vertical: kXlSpacing,
        horizontal: kXlSpacing,
      ),
      child: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            Text(isPressed ? appLocale.linkCopied : appLocale.clickForLink),
            GestureDetector(
              onTap: () async {
                setState(() {
                  isPressed = true;
                });
                await Clipboard.setData(
                  const ClipboardData(
                    text:
                        'https://play.beam.apache.org/fhudiu89lkd/mycodeexample',
                  ),
                );
              },
              child: Container(
                decoration: BoxDecoration(
                  color: isPressed
                      ? ThemeColors.of(context).greyColor
                      : ThemeColors.of(context).primaryBackground,
                  border: Border.all(
                    color: ThemeColors.of(context).primary,
                  ),
                  borderRadius: BorderRadius.circular(kSmBorderRadius),
                ),
                child: Center(
                  child: Padding(
                    padding: const EdgeInsets.symmetric(vertical: kLgSpacing),
                    child: isPressed
                        ? SelectableText(
                            'https://play.beam.apache.org/fhudiu89lkd/mycodeexample',
                            style: TextStyle(
                              fontSize: kLabelFontSize,
                              fontWeight: kNormalWeight,
                              color: ThemeColors.of(context).primary,
                            ),
                          )
                        : Text(
                            appLocale.showAndCopyLink,
                            style: TextStyle(
                              fontSize: kLabelFontSize,
                              fontWeight: kBoldWeight,
                              color: ThemeColors.of(context).primary,
                            ),
                          ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
