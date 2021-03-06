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

// This file was copied from
// https://github.com/openshift/kubernetes-autoscaler/blob/95d5f0927b9347635f1720d6748503415e6921cb/cluster-autoscaler/cloudprovider/azure/azure_instance_types.go
// // https://github.com/kubernetes/kubernetes/issues/79384

package machineset

type instanceType struct {
	InstanceType string
	VCPU         int64
	MemoryMb     int64
	GPU          int64
}

// InstanceTypes is a map of azure resources
var InstanceTypes = map[string]*instanceType{
	"Standard_D4s_v3": {
		InstanceType: "Standard_D4s_v3",
		VCPU:         4,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_DS11_v2_Promo": {
		InstanceType: "Standard_DS11_v2_Promo",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_D8s_v3": {
		InstanceType: "Standard_D8s_v3",
		VCPU:         8,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_D32_v3": {
		InstanceType: "Standard_D32_v3",
		VCPU:         32,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_G1": {
		InstanceType: "Standard_G1",
		VCPU:         2,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_G2": {
		InstanceType: "Standard_G2",
		VCPU:         4,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_G3": {
		InstanceType: "Standard_G3",
		VCPU:         8,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_G4": {
		InstanceType: "Standard_G4",
		VCPU:         16,
		MemoryMb:     229376,
		GPU:          0,
	},
	"Standard_G5": {
		InstanceType: "Standard_G5",
		VCPU:         32,
		MemoryMb:     458752,
		GPU:          0,
	},
	"Standard_NV24": {
		InstanceType: "Standard_NV24",
		VCPU:         24,
		MemoryMb:     229376,
		GPU:          4,
	},
	"Standard_E4s_v3": {
		InstanceType: "Standard_E4s_v3",
		VCPU:         4,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_A2_v2": {
		InstanceType: "Standard_A2_v2",
		VCPU:         2,
		MemoryMb:     4096,
		GPU:          0,
	},
	"Standard_E2_v3": {
		InstanceType: "Standard_E2_v3",
		VCPU:         2,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_D2s_v3": {
		InstanceType: "Standard_D2s_v3",
		VCPU:         2,
		MemoryMb:     8192,
		GPU:          0,
	},
	"Standard_M64-16ms": {
		InstanceType: "Standard_M64-16ms",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_D5_v2": {
		InstanceType: "Standard_D5_v2",
		VCPU:         16,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_A1_v2": {
		InstanceType: "Standard_A1_v2",
		VCPU:         1,
		MemoryMb:     2048,
		GPU:          0,
	},
	"Standard_A8m_v2": {
		InstanceType: "Standard_A8m_v2",
		VCPU:         8,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_D11_v2_Promo": {
		InstanceType: "Standard_D11_v2_Promo",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_DS14-4_v2": {
		InstanceType: "Standard_DS14-4_v2",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_M64ms": {
		InstanceType: "Standard_M64ms",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_E8s_v3": {
		InstanceType: "Standard_E8s_v3",
		VCPU:         8,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_E64i_v3": {
		InstanceType: "Standard_E64i_v3",
		VCPU:         64,
		MemoryMb:     442368,
		GPU:          0,
	},
	"Standard_A4m_v2": {
		InstanceType: "Standard_A4m_v2",
		VCPU:         4,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_ND6s": {
		InstanceType: "Standard_ND6s",
		VCPU:         6,
		MemoryMb:     114688,
		GPU:          1,
	},
	"Standard_H16r": {
		InstanceType: "Standard_H16r",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_D4": {
		InstanceType: "Standard_D4",
		VCPU:         8,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_DS14-8_v2": {
		InstanceType: "Standard_DS14-8_v2",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_E8_v3": {
		InstanceType: "Standard_E8_v3",
		VCPU:         8,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_D1": {
		InstanceType: "Standard_D1",
		VCPU:         1,
		MemoryMb:     3584,
		GPU:          0,
	},
	"Standard_M64s": {
		InstanceType: "Standard_M64s",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_D3": {
		InstanceType: "Standard_D3",
		VCPU:         4,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_D2": {
		InstanceType: "Standard_D2",
		VCPU:         2,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_B1ms": {
		InstanceType: "Standard_B1ms",
		VCPU:         1,
		MemoryMb:     2048,
		GPU:          0,
	},
	"Standard_GS5-16": {
		InstanceType: "Standard_GS5-16",
		VCPU:         32,
		MemoryMb:     458752,
		GPU:          0,
	},
	"Standard_D2_v2_Promo": {
		InstanceType: "Standard_D2_v2_Promo",
		VCPU:         2,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_DS13_v2": {
		InstanceType: "Standard_DS13_v2",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_H8m": {
		InstanceType: "Standard_H8m",
		VCPU:         8,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_E32-8s_v3": {
		InstanceType: "Standard_E32-8s_v3",
		VCPU:         32,
		MemoryMb:     262144,
		GPU:          0,
	},
	"Standard_F72s_v2": {
		InstanceType: "Standard_F72s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_F32s_v2": {
		InstanceType: "Standard_F32s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_F2s": {
		InstanceType: "Standard_F2s",
		VCPU:         2,
		MemoryMb:     4096,
		GPU:          0,
	},
	"Standard_D14_v2": {
		InstanceType: "Standard_D14_v2",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_D12_v2_Promo": {
		InstanceType: "Standard_D12_v2_Promo",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_DS13-2_v2": {
		InstanceType: "Standard_DS13-2_v2",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_D4_v2_Promo": {
		InstanceType: "Standard_D4_v2_Promo",
		VCPU:         8,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_B2ms": {
		InstanceType: "Standard_B2ms",
		VCPU:         2,
		MemoryMb:     8192,
		GPU:          0,
	},
	"Standard_D3_v2_Promo": {
		InstanceType: "Standard_D3_v2_Promo",
		VCPU:         4,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_DS12_v2_Promo": {
		InstanceType: "Standard_DS12_v2_Promo",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_NC6s_v3": {
		InstanceType: "Standard_NC6s_v3",
		VCPU:         6,
		MemoryMb:     114688,
		GPU:          1,
	},
	"Standard_NC6s_v2": {
		InstanceType: "Standard_NC6s_v2",
		VCPU:         6,
		MemoryMb:     114688,
		GPU:          1,
	},
	"Standard_DS13-4_v2": {
		InstanceType: "Standard_DS13-4_v2",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_NC24": {
		InstanceType: "Standard_NC24",
		VCPU:         24,
		MemoryMb:     229376,
		GPU:          4,
	},
	"Standard_D64_v3": {
		InstanceType: "Standard_D64_v3",
		VCPU:         64,
		MemoryMb:     262144,
		GPU:          0,
	},
	"Standard_E32s_v3": {
		InstanceType: "Standard_E32s_v3",
		VCPU:         32,
		MemoryMb:     262144,
		GPU:          0,
	},
	"Standard_D32s_v3": {
		InstanceType: "Standard_D32s_v3",
		VCPU:         32,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_D4_v3": {
		InstanceType: "Standard_D4_v3",
		VCPU:         4,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_D4_v2": {
		InstanceType: "Standard_D4_v2",
		VCPU:         8,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_D13_v2_Promo": {
		InstanceType: "Standard_D13_v2_Promo",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_L8s": {
		InstanceType: "Standard_L8s",
		VCPU:         8,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_M64-32ms": {
		InstanceType: "Standard_M64-32ms",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_DS3_v2_Promo": {
		InstanceType: "Standard_DS3_v2_Promo",
		VCPU:         4,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_ND12s": {
		InstanceType: "Standard_ND12s",
		VCPU:         12,
		MemoryMb:     688128,
		GPU:          2,
	},
	"Standard_DS5_v2_Promo": {
		InstanceType: "Standard_DS5_v2_Promo",
		VCPU:         16,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_D5_v2_Promo": {
		InstanceType: "Standard_D5_v2_Promo",
		VCPU:         16,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_D16_v3": {
		InstanceType: "Standard_D16_v3",
		VCPU:         16,
		MemoryMb:     65638,
		GPU:          0,
	},
	"Standard_E64-32s_v3": {
		InstanceType: "Standard_E64-32s_v3",
		VCPU:         64,
		MemoryMb:     442368,
		GPU:          0,
	},
	"Standard_GS5-8": {
		InstanceType: "Standard_GS5-8",
		VCPU:         32,
		MemoryMb:     458752,
		GPU:          0,
	},
	"Standard_DS13_v2_Promo": {
		InstanceType: "Standard_DS13_v2_Promo",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_F16s_v2": {
		InstanceType: "Standard_F16s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_DS11_v2": {
		InstanceType: "Standard_DS11_v2",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_D11": {
		InstanceType: "Standard_D11",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_D13": {
		InstanceType: "Standard_D13",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_D12": {
		InstanceType: "Standard_D12",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_NC24r": {
		InstanceType: "Standard_NC24r",
		VCPU:         24,
		MemoryMb:     229376,
		GPU:          4,
	},
	"Standard_D14": {
		InstanceType: "Standard_D14",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_GS4-4": {
		InstanceType: "Standard_GS4-4",
		VCPU:         16,
		MemoryMb:     229376,
		GPU:          0,
	},
	"Standard_F4s": {
		InstanceType: "Standard_F4s",
		VCPU:         4,
		MemoryMb:     8192,
		GPU:          0,
	},
	"Standard_DS14": {
		InstanceType: "Standard_DS14",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_DS13": {
		InstanceType: "Standard_DS13",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_DS12": {
		InstanceType: "Standard_DS12",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_DS11": {
		InstanceType: "Standard_DS11",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_DS14_v2_Promo": {
		InstanceType: "Standard_DS14_v2_Promo",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_NV6": {
		InstanceType: "Standard_NV6",
		VCPU:         6,
		MemoryMb:     57344,
		GPU:          1,
	},
	"Standard_NV6_Promo": {
		InstanceType: "Standard_NV6_Promo",
		VCPU:         6,
		MemoryMb:     57344,
		GPU:          1,
	},
	"Standard_D15_v2": {
		InstanceType: "Standard_D15_v2",
		VCPU:         20,
		MemoryMb:     143360,
		GPU:          0,
	},
	"Standard_D3_v2": {
		InstanceType: "Standard_D3_v2",
		VCPU:         4,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_L32s": {
		InstanceType: "Standard_L32s",
		VCPU:         32,
		MemoryMb:     262144,
		GPU:          0,
	},
	"Standard_NC12s_v3": {
		InstanceType: "Standard_NC12s_v3",
		VCPU:         12,
		MemoryMb:     229376,
		GPU:          2,
	},
	"Standard_NC12s_v2": {
		InstanceType: "Standard_NC12s_v2",
		VCPU:         12,
		MemoryMb:     229376,
		GPU:          2,
	},
	"Standard_A2m_v2": {
		InstanceType: "Standard_A2m_v2",
		VCPU:         2,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_E32_v3": {
		InstanceType: "Standard_E32_v3",
		VCPU:         32,
		MemoryMb:     262144,
		GPU:          0,
	},
	"Standard_DS12_v2": {
		InstanceType: "Standard_DS12_v2",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_D14_v2_Promo": {
		InstanceType: "Standard_D14_v2_Promo",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_F4": {
		InstanceType: "Standard_F4",
		VCPU:         4,
		MemoryMb:     8192,
		GPU:          0,
	},
	"Standard_H16m": {
		InstanceType: "Standard_H16m",
		VCPU:         16,
		MemoryMb:     229376,
		GPU:          0,
	},
	"Standard_F8s": {
		InstanceType: "Standard_F8s",
		VCPU:         8,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_E64s_v3": {
		InstanceType: "Standard_E64s_v3",
		VCPU:         64,
		MemoryMb:     442368,
		GPU:          0,
	},
	"Standard_DS12-1_v2": {
		InstanceType: "Standard_DS12-1_v2",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_E8-4s_v3": {
		InstanceType: "Standard_E8-4s_v3",
		VCPU:         8,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_NV12": {
		InstanceType: "Standard_NV12",
		VCPU:         12,
		MemoryMb:     114688,
		GPU:          2,
	},
	"Standard_E16-4s_v3": {
		InstanceType: "Standard_E16-4s_v3",
		VCPU:         16,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_ND24rs": {
		InstanceType: "Standard_ND24rs",
		VCPU:         24,
		MemoryMb:     458752,
		GPU:          4,
	},
	"Standard_D8_v3": {
		InstanceType: "Standard_D8_v3",
		VCPU:         8,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_DS12-2_v2": {
		InstanceType: "Standard_DS12-2_v2",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_B8ms": {
		InstanceType: "Standard_B8ms",
		VCPU:         8,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_DS4_v2_Promo": {
		InstanceType: "Standard_DS4_v2_Promo",
		VCPU:         8,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_DS14_v2": {
		InstanceType: "Standard_DS14_v2",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_NC24s_v2": {
		InstanceType: "Standard_NC24s_v2",
		VCPU:         24,
		MemoryMb:     458752,
		GPU:          4,
	},
	"Standard_M128-64ms": {
		InstanceType: "Standard_M128-64ms",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_NC24rs_v3": {
		InstanceType: "Standard_NC24rs_v3",
		VCPU:         24,
		MemoryMb:     458752,
		GPU:          4,
	},
	"Standard_NC24rs_v2": {
		InstanceType: "Standard_NC24rs_v2",
		VCPU:         24,
		MemoryMb:     458752,
		GPU:          4,
	},
	"Standard_E8-2s_v3": {
		InstanceType: "Standard_E8-2s_v3",
		VCPU:         8,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_E16_v3": {
		InstanceType: "Standard_E16_v3",
		VCPU:         16,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_E4-2s_v3": {
		InstanceType: "Standard_E4-2s_v3",
		VCPU:         4,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_NC6": {
		InstanceType: "Standard_NC6",
		VCPU:         6,
		MemoryMb:     57344,
		GPU:          1,
	},
	"Standard_D11_v2": {
		InstanceType: "Standard_D11_v2",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_F16s": {
		InstanceType: "Standard_F16s",
		VCPU:         16,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_DS11-1_v2": {
		InstanceType: "Standard_DS11-1_v2",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_NC24s_v3": {
		InstanceType: "Standard_NC24s_v3",
		VCPU:         24,
		MemoryMb:     458752,
		GPU:          4,
	},
	"Standard_M128s": {
		InstanceType: "Standard_M128s",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_DS3_v2": {
		InstanceType: "Standard_DS3_v2",
		VCPU:         4,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_GS5": {
		InstanceType: "Standard_GS5",
		VCPU:         32,
		MemoryMb:     458752,
		GPU:          0,
	},
	"Standard_A4_v2": {
		InstanceType: "Standard_A4_v2",
		VCPU:         4,
		MemoryMb:     8192,
		GPU:          0,
	},
	"Standard_GS1": {
		InstanceType: "Standard_GS1",
		VCPU:         2,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_GS2": {
		InstanceType: "Standard_GS2",
		VCPU:         4,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_GS3": {
		InstanceType: "Standard_GS3",
		VCPU:         8,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_D64s_v3": {
		InstanceType: "Standard_D64s_v3",
		VCPU:         64,
		MemoryMb:     262144,
		GPU:          0,
	},
	"Standard_F64s_v2": {
		InstanceType: "Standard_F64s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_GS4": {
		InstanceType: "Standard_GS4",
		VCPU:         16,
		MemoryMb:     229376,
		GPU:          0,
	},
	"Standard_D13_v2": {
		InstanceType: "Standard_D13_v2",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_NC12": {
		InstanceType: "Standard_NC12",
		VCPU:         12,
		MemoryMb:     114688,
		GPU:          2,
	},
	"Standard_D1_v2": {
		InstanceType: "Standard_D1_v2",
		VCPU:         1,
		MemoryMb:     3584,
		GPU:          0,
	},
	"Standard_E16-8s_v3": {
		InstanceType: "Standard_E16-8s_v3",
		VCPU:         16,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_L16s": {
		InstanceType: "Standard_L16s",
		VCPU:         16,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_D2_v3": {
		InstanceType: "Standard_D2_v3",
		VCPU:         2,
		MemoryMb:     8192,
		GPU:          0,
	},
	"Standard_D2_v2": {
		InstanceType: "Standard_D2_v2",
		VCPU:         2,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_H16mr": {
		InstanceType: "Standard_H16mr",
		VCPU:         16,
		MemoryMb:     229376,
		GPU:          0,
	},
	"Standard_A6": {
		InstanceType: "Standard_A6",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_A7": {
		InstanceType: "Standard_A7",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_A4": {
		InstanceType: "Standard_A4",
		VCPU:         8,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_A5": {
		InstanceType: "Standard_A5",
		VCPU:         2,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_A2": {
		InstanceType: "Standard_A2",
		VCPU:         2,
		MemoryMb:     3584,
		GPU:          0,
	},
	"Standard_A3": {
		InstanceType: "Standard_A3",
		VCPU:         4,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_A0": {
		InstanceType: "Standard_A0",
		VCPU:         1,
		MemoryMb:     768,
		GPU:          0,
	},
	"Standard_A1": {
		InstanceType: "Standard_A1",
		VCPU:         1,
		MemoryMb:     1792,
		GPU:          0,
	},
	"Standard_DS5_v2": {
		InstanceType: "Standard_DS5_v2",
		VCPU:         16,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_A8": {
		InstanceType: "Standard_A8",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_A9": {
		InstanceType: "Standard_A9",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Basic_A4": {
		InstanceType: "Basic_A4",
		VCPU:         8,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_E64-16s_v3": {
		InstanceType: "Standard_E64-16s_v3",
		VCPU:         64,
		MemoryMb:     442368,
		GPU:          0,
	},
	"Basic_A2": {
		InstanceType: "Basic_A2",
		VCPU:         2,
		MemoryMb:     3584,
		GPU:          0,
	},
	"Basic_A3": {
		InstanceType: "Basic_A3",
		VCPU:         4,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Basic_A0": {
		InstanceType: "Basic_A0",
		VCPU:         1,
		MemoryMb:     768,
		GPU:          0,
	},
	"Basic_A1": {
		InstanceType: "Basic_A1",
		VCPU:         1,
		MemoryMb:     1792,
		GPU:          0,
	},
	"Standard_A10": {
		InstanceType: "Standard_A10",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_A11": {
		InstanceType: "Standard_A11",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_E4_v3": {
		InstanceType: "Standard_E4_v3",
		VCPU:         4,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_F8s_v2": {
		InstanceType: "Standard_F8s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_F4s_v2": {
		InstanceType: "Standard_F4s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_ND24s": {
		InstanceType: "Standard_ND24s",
		VCPU:         24,
		MemoryMb:     458752,
		GPU:          4,
	},
	"Standard_B2s": {
		InstanceType: "Standard_B2s",
		VCPU:         2,
		MemoryMb:     4096,
		GPU:          0,
	},
	"Standard_DS2_v2": {
		InstanceType: "Standard_DS2_v2",
		VCPU:         2,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_F1s": {
		InstanceType: "Standard_F1s",
		VCPU:         1,
		MemoryMb:     2048,
		GPU:          0,
	},
	"Standard_E16s_v3": {
		InstanceType: "Standard_E16s_v3",
		VCPU:         16,
		MemoryMb:     131072,
		GPU:          0,
	},
	"Standard_B4ms": {
		InstanceType: "Standard_B4ms",
		VCPU:         4,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_DS15_v2": {
		InstanceType: "Standard_DS15_v2",
		VCPU:         20,
		MemoryMb:     143360,
		GPU:          0,
	},
	"Standard_D12_v2": {
		InstanceType: "Standard_D12_v2",
		VCPU:         4,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_M128-32ms": {
		InstanceType: "Standard_M128-32ms",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_M128ms": {
		InstanceType: "Standard_M128ms",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_DS3": {
		InstanceType: "Standard_DS3",
		VCPU:         4,
		MemoryMb:     14336,
		GPU:          0,
	},
	"Standard_DS2": {
		InstanceType: "Standard_DS2",
		VCPU:         2,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_DS1": {
		InstanceType: "Standard_DS1",
		VCPU:         1,
		MemoryMb:     3584,
		GPU:          0,
	},
	"Standard_DS2_v2_Promo": {
		InstanceType: "Standard_DS2_v2_Promo",
		VCPU:         2,
		MemoryMb:     7168,
		GPU:          0,
	},
	"Standard_DS1_v2": {
		InstanceType: "Standard_DS1_v2",
		VCPU:         1,
		MemoryMb:     3584,
		GPU:          0,
	},
	"Standard_DS4": {
		InstanceType: "Standard_DS4",
		VCPU:         8,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_F2": {
		InstanceType: "Standard_F2",
		VCPU:         2,
		MemoryMb:     4096,
		GPU:          0,
	},
	"Standard_F1": {
		InstanceType: "Standard_F1",
		VCPU:         1,
		MemoryMb:     2048,
		GPU:          0,
	},
	"Standard_E2s_v3": {
		InstanceType: "Standard_E2s_v3",
		VCPU:         2,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_F2s_v2": {
		InstanceType: "Standard_F2s_v2",
		VCPU:         0,
		MemoryMb:     0,
		GPU:          0,
	},
	"Standard_E64is_v3": {
		InstanceType: "Standard_E64is_v3",
		VCPU:         64,
		MemoryMb:     442368,
		GPU:          0,
	},
	"Standard_L4s": {
		InstanceType: "Standard_L4s",
		VCPU:         4,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_H16": {
		InstanceType: "Standard_H16",
		VCPU:         16,
		MemoryMb:     114688,
		GPU:          0,
	},
	"Standard_F8": {
		InstanceType: "Standard_F8",
		VCPU:         8,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_GS4-8": {
		InstanceType: "Standard_GS4-8",
		VCPU:         16,
		MemoryMb:     229376,
		GPU:          0,
	},
	"Standard_DS4_v2": {
		InstanceType: "Standard_DS4_v2",
		VCPU:         8,
		MemoryMb:     28672,
		GPU:          0,
	},
	"Standard_D16s_v3": {
		InstanceType: "Standard_D16s_v3",
		VCPU:         16,
		MemoryMb:     65536,
		GPU:          0,
	},
	"Standard_H8": {
		InstanceType: "Standard_H8",
		VCPU:         8,
		MemoryMb:     57344,
		GPU:          0,
	},
	"Standard_A8_v2": {
		InstanceType: "Standard_A8_v2",
		VCPU:         8,
		MemoryMb:     16384,
		GPU:          0,
	},
	"Standard_B1s": {
		InstanceType: "Standard_B1s",
		VCPU:         1,
		MemoryMb:     1024,
		GPU:          0,
	},
	"Standard_F16": {
		InstanceType: "Standard_F16",
		VCPU:         16,
		MemoryMb:     32768,
		GPU:          0,
	},
	"Standard_E64_v3": {
		InstanceType: "Standard_E64_v3",
		VCPU:         64,
		MemoryMb:     442368,
		GPU:          0,
	},
	"Standard_E32-16s_v3": {
		InstanceType: "Standard_E32-16s_v3",
		VCPU:         32,
		MemoryMb:     262144,
		GPU:          0,
	},
}
