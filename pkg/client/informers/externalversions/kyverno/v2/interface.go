/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	internalinterfaces "github.com/kyverno/kyverno/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AdmissionReports returns a AdmissionReportInformer.
	AdmissionReports() AdmissionReportInformer
	// BackgroundScanReports returns a BackgroundScanReportInformer.
	BackgroundScanReports() BackgroundScanReportInformer
	// CleanupPolicies returns a CleanupPolicyInformer.
	CleanupPolicies() CleanupPolicyInformer
	// ClusterAdmissionReports returns a ClusterAdmissionReportInformer.
	ClusterAdmissionReports() ClusterAdmissionReportInformer
	// ClusterBackgroundScanReports returns a ClusterBackgroundScanReportInformer.
	ClusterBackgroundScanReports() ClusterBackgroundScanReportInformer
	// ClusterCleanupPolicies returns a ClusterCleanupPolicyInformer.
	ClusterCleanupPolicies() ClusterCleanupPolicyInformer
	// PolicyExceptions returns a PolicyExceptionInformer.
	PolicyExceptions() PolicyExceptionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AdmissionReports returns a AdmissionReportInformer.
func (v *version) AdmissionReports() AdmissionReportInformer {
	return &admissionReportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackgroundScanReports returns a BackgroundScanReportInformer.
func (v *version) BackgroundScanReports() BackgroundScanReportInformer {
	return &backgroundScanReportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CleanupPolicies returns a CleanupPolicyInformer.
func (v *version) CleanupPolicies() CleanupPolicyInformer {
	return &cleanupPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterAdmissionReports returns a ClusterAdmissionReportInformer.
func (v *version) ClusterAdmissionReports() ClusterAdmissionReportInformer {
	return &clusterAdmissionReportInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterBackgroundScanReports returns a ClusterBackgroundScanReportInformer.
func (v *version) ClusterBackgroundScanReports() ClusterBackgroundScanReportInformer {
	return &clusterBackgroundScanReportInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterCleanupPolicies returns a ClusterCleanupPolicyInformer.
func (v *version) ClusterCleanupPolicies() ClusterCleanupPolicyInformer {
	return &clusterCleanupPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PolicyExceptions returns a PolicyExceptionInformer.
func (v *version) PolicyExceptions() PolicyExceptionInformer {
	return &policyExceptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}