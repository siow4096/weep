/*
 * Copyright 2020 Netflix, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package errors

type Error string

func (e Error) Error() string { return string(e) }

const (
	NoCredentialsFoundInCache      = Error("no credentials found in cache")
	NoTokenFoundInCache            = Error("no token found in cache")
	NoDefaultRoleSet               = Error("no default role set")
	BrowserOpenError               = Error("could not launch browser, open link manually")
	CredentialRetrievalError       = Error("failed to retrieve credentials from broker")
	InvalidJWT                     = Error("JWT is invalid")
	InvalidArn                     = Error("requested ARN is invalid")
	MutualTLSCertNeedsRefreshError = Error("mTLS cert needs to be refreshed")
	MultipleMatchingRoles          = Error("more than one matching role for search string")
	NoMatchingRoles                = Error("no matching roles for search string")
	MalformedRequestError          = Error("malformed request sent to broker")
)
