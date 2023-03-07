// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package cache_policy

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig, b.ko.Spec.CachePolicyConfig) {
		delta.Add("Spec.CachePolicyConfig", a.ko.Spec.CachePolicyConfig, b.ko.Spec.CachePolicyConfig)
	} else if a.ko.Spec.CachePolicyConfig != nil && b.ko.Spec.CachePolicyConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.Comment, b.ko.Spec.CachePolicyConfig.Comment) {
			delta.Add("Spec.CachePolicyConfig.Comment", a.ko.Spec.CachePolicyConfig.Comment, b.ko.Spec.CachePolicyConfig.Comment)
		} else if a.ko.Spec.CachePolicyConfig.Comment != nil && b.ko.Spec.CachePolicyConfig.Comment != nil {
			if *a.ko.Spec.CachePolicyConfig.Comment != *b.ko.Spec.CachePolicyConfig.Comment {
				delta.Add("Spec.CachePolicyConfig.Comment", a.ko.Spec.CachePolicyConfig.Comment, b.ko.Spec.CachePolicyConfig.Comment)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.DefaultTTL, b.ko.Spec.CachePolicyConfig.DefaultTTL) {
			delta.Add("Spec.CachePolicyConfig.DefaultTTL", a.ko.Spec.CachePolicyConfig.DefaultTTL, b.ko.Spec.CachePolicyConfig.DefaultTTL)
		} else if a.ko.Spec.CachePolicyConfig.DefaultTTL != nil && b.ko.Spec.CachePolicyConfig.DefaultTTL != nil {
			if *a.ko.Spec.CachePolicyConfig.DefaultTTL != *b.ko.Spec.CachePolicyConfig.DefaultTTL {
				delta.Add("Spec.CachePolicyConfig.DefaultTTL", a.ko.Spec.CachePolicyConfig.DefaultTTL, b.ko.Spec.CachePolicyConfig.DefaultTTL)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.MaxTTL, b.ko.Spec.CachePolicyConfig.MaxTTL) {
			delta.Add("Spec.CachePolicyConfig.MaxTTL", a.ko.Spec.CachePolicyConfig.MaxTTL, b.ko.Spec.CachePolicyConfig.MaxTTL)
		} else if a.ko.Spec.CachePolicyConfig.MaxTTL != nil && b.ko.Spec.CachePolicyConfig.MaxTTL != nil {
			if *a.ko.Spec.CachePolicyConfig.MaxTTL != *b.ko.Spec.CachePolicyConfig.MaxTTL {
				delta.Add("Spec.CachePolicyConfig.MaxTTL", a.ko.Spec.CachePolicyConfig.MaxTTL, b.ko.Spec.CachePolicyConfig.MaxTTL)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.MinTTL, b.ko.Spec.CachePolicyConfig.MinTTL) {
			delta.Add("Spec.CachePolicyConfig.MinTTL", a.ko.Spec.CachePolicyConfig.MinTTL, b.ko.Spec.CachePolicyConfig.MinTTL)
		} else if a.ko.Spec.CachePolicyConfig.MinTTL != nil && b.ko.Spec.CachePolicyConfig.MinTTL != nil {
			if *a.ko.Spec.CachePolicyConfig.MinTTL != *b.ko.Spec.CachePolicyConfig.MinTTL {
				delta.Add("Spec.CachePolicyConfig.MinTTL", a.ko.Spec.CachePolicyConfig.MinTTL, b.ko.Spec.CachePolicyConfig.MinTTL)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.Name, b.ko.Spec.CachePolicyConfig.Name) {
			delta.Add("Spec.CachePolicyConfig.Name", a.ko.Spec.CachePolicyConfig.Name, b.ko.Spec.CachePolicyConfig.Name)
		} else if a.ko.Spec.CachePolicyConfig.Name != nil && b.ko.Spec.CachePolicyConfig.Name != nil {
			if *a.ko.Spec.CachePolicyConfig.Name != *b.ko.Spec.CachePolicyConfig.Name {
				delta.Add("Spec.CachePolicyConfig.Name", a.ko.Spec.CachePolicyConfig.Name, b.ko.Spec.CachePolicyConfig.Name)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin) {
			delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin)
		} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig) {
				delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig)
			} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior) {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior)
				} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior != nil {
					if *a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior != *b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies) {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies)
				} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies != nil {
					if !ackcompare.SliceStringPEqual(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items) {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli) {
				delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli)
			} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli != nil {
				if *a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli != *b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip) {
				delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip)
			} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip != nil {
				if *a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip != *b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig) {
				delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig)
			} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior) {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior)
				} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior != nil {
					if *a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior != *b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers) {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers)
				} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers != nil {
					if !ackcompare.SliceStringPEqual(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items) {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig) {
				delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig)
			} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior) {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior)
				} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior != nil {
					if *a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior != *b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings) {
					delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings)
				} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings != nil {
					if !ackcompare.SliceStringPEqual(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items) {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items)
					}
					if ackcompare.HasNilDifference(a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity) {
						delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity)
					} else if a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity != nil && b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity != nil {
						if *a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity != *b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity {
							delta.Add("Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity", a.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity, b.ko.Spec.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity)
						}
					}
				}
			}
		}
	}

	return delta
}
