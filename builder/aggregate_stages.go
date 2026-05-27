package builder

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Bucket function returns a mongo $bucket operator used in aggregations.
func Bucket(groupBy, boundaries, def, output interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// BucketAuto function returns a mongo $bucketAuto operator used in aggregations.
func BucketAuto(groupBy, buckets, output, granularity interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// CollStats function returns a mongo $collStats operator used in aggregations.
func CollStats(latencyStats, storageStats, count interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// CurrentOp function returns a mongo $currentOp operator used in aggregations.
func CurrentOp(allUsers, idleConnections, idleCursors, idleSessions, localOps interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// $geoNear,$graphLookup has many params, those functions
// will have too many params and do not make readable code.

// Group function returns a mongo $group operator used in aggregations.
func Group(ID interface{}, params bson.M) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// Lookup function returns a mongo $lookup operator used in aggregations.
func Lookup(from, localField, foreignField, as interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// UncorrelatedLookup function returns a mongo $lookup operator used in aggregations.
func UncorrelatedLookup(from, let, pipeline, as interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// Merge function returns a mongo $merge operator used in aggregations.
func Merge(into, on, let, whenMatched, whenNotMatched interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}

// ReplaceRoot function returns a mongo $replaceRoot operator used in aggregations.
func ReplaceRoot(newRoot interface{}) Operator { _ = "STUB: not implemented"; return *new(Operator) }

// Sample function returns a mongo sample operator used in aggregations.
func Sample(size interface{}) Operator { _ = "STUB: not implemented"; return *new(Operator) }

// Unwind function returns a mongo $unwind operator used in aggregations.
func Unwind(path, includeArrayIndex, preserveNullAndEmptyArrays interface{}) Operator {
	_ = "STUB: not implemented"
	return *new(Operator)
}
