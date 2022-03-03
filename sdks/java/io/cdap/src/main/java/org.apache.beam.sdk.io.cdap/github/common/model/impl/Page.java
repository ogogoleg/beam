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
package org.apache.beam.sdk.io.cdap.github.common.model.impl;

import com.google.api.client.util.Key;
import org.apache.beam.sdk.io.cdap.github.common.model.GitHubModel;

/** Page model for github. */
@SuppressWarnings("UnusedVariable")
public class Page implements GitHubModel {

  @Key private String url;
  @Key private String status;
  @Key private String cname;

  @Key("custom_404")
  private Boolean custom404;

  @Key("html_url")
  private String htmlUrl;

  @Key private Source source;

  /** Page.Source model */
  public static class Source {
    @Key private String branch;
    @Key private String directory;
  }
}
