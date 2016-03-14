package utils

import (
	mesos "github.com/mesos/mesos-go/mesosproto"
	util "github.com/mesos/mesos-go/mesosutil"
)

func TaskMatches(data *TaskData, offer *mesos.Offer) string {
	if data.Cpu > GetScalarResources(offer, "cpus") {
		return "no cpus"
	}

	if data.Mem > GetScalarResources(offer, "mem") {
		return "no mem"
	}
	return ""
}

func GetScalarResources(offer *mesos.Offer, resourceName string) float64 {
	resources := 0.0
	filteredResources := util.FilterResources(offer.Resources, func(res *mesos.Resource) bool {
		return res.GetName() == resourceName
	})
	for _, res := range filteredResources {
		resources += res.GetScalar().GetValue()
	}
	return resources
}

func GetRangeResources(offer *mesos.Offer, resourceName string) []*mesos.Value_Range {
	resources := make([]*mesos.Value_Range, 0)
	filteredResources := util.FilterResources(offer.Resources, func(res *mesos.Resource) bool {
		return res.GetName() == resourceName
	})
	for _, res := range filteredResources {
		resources = append(resources, res.GetRanges().GetRange()...)
	}
	return resources
}
