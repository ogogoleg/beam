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
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:playground/components/dropdown_button/dropdown_button.dart';
import 'package:playground/config/theme.dart';
import 'package:playground/modules/editor/components/share_dropdown/share_dropdown_body.dart';

const kShareDropdownHeight = 120.0;
const kShareDropdownWidth = 460.0;

class ShareButton extends StatefulWidget {
  const ShareButton({Key? key}) : super(key: key);

  @override
  State<ShareButton> createState() => _ShareButtonState();
}

class _ShareButtonState extends State<ShareButton> {
  @override
  Widget build(BuildContext context) {
    AppLocalizations appLocale = AppLocalizations.of(context)!;

    return AppDropdownButton(
      buttonText: Text(appLocale.shareMyCode),
      buttonColor: ThemeColors.of(context).primary.withOpacity(0.2),
      dropdownBackgroundColor: ThemeColors.of(context).secondaryBackground,
      withArrowDown: false,
      leadingIcon: Icon(
        Icons.share_outlined,
        color: ThemeColors.of(context).primary,
      ),
      height: kShareDropdownHeight,
      width: kShareDropdownWidth,
      dropdownAlign: DropdownAlignment.right,
      createDropdown: (close) => ShareDropdownBody(close: close),
    );
  }
}
