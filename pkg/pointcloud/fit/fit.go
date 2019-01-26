// Package fit will place a set of vectors correctly into a pointcloud.
package fit

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/pointcloud"
	"github.com/andrewbackes/autonoma/pkg/vector"
	"github.com/andrewbackes/autonoma/pkg/vector/transformation"
	"gonum.org/v1/gonum/mat"
	"math"
)

func ICP(source []vector.Vector, target *pointcloud.PointCloud, epsilon float64, iterations int) ([]vector.Vector, float64) {
	// handle starting condition
	if len(target.Points) == 0 {
		return source, 0
	}
	transformed := vector.Copy(source)
	dist := math.MaxFloat64
	for i := 0; (dist > epsilon) && (i < iterations); i++ {
		fmt.Println("Iteration", i)
		matched := closestPoints(transformed, target)
		dist = vector.AveDistance(transformed, matched)
		fmt.Println("Average Distance", dist)
		trans := nextTransformation(transformed, matched)
		apply(trans, transformed)
	}
	return transformed, dist
}

func apply(trans *transformation.Transformation, to []vector.Vector) {
	for i := range to {
		to[i] = trans.Apply(to[i])
	}
}

func closestPoints(source []vector.Vector, target *pointcloud.PointCloud) []vector.Vector {
	if len(target.Points) < len(source) {
		panic("not enough target vectors to match")
	}
	used := map[vector.Vector]struct{}{}
	matches := make([]vector.Vector, len(source))
	for i, v := range source {
		matches[i] = closestUnusedPoint(v, target, used)
		used[matches[i]] = struct{}{}
	}
	return matches
}

func closestUnusedPoint(v vector.Vector, to *pointcloud.PointCloud, used map[vector.Vector]struct{}) vector.Vector {
	if len(to.Points) <= len(used) {
		panic("no unused points left")
	}
	min := math.MaxFloat64
	closest := vector.Vector{}
	for w := range to.Points {
		d := vector.Distance(v, w)
		if _, taken := used[w]; !taken && d < min {
			min = d
			closest = w
		}
	}
	fmt.Println("--->", v, "closest to", closest)
	return closest
}

func nextTransformation(source, pairs []vector.Vector) *transformation.Transformation {
	sourceCentroid := vector.Centroid(source)
	pairsCentroid := vector.Centroid(pairs)
	crossCovarianceMatrix := crossCovarianceOf(source, pairs, sourceCentroid, pairsCentroid)
	rotation := rotationOf(crossCovarianceMatrix)
	return &transformation.Transformation{
		Rotation:    rotation,
		Translation: translationOf(sourceCentroid, pairsCentroid, rotation),
	}
}

func crossCovarianceOf(source, pairs []vector.Vector, sourceCentroid, pairedCentroid vector.Vector) mat.Matrix {
	// make new source and match sets with centroids subtracted
	shiftedSource := vector.Copy(source)
	vector.Shift(shiftedSource, sourceCentroid)

	shiftedPairs := vector.Copy(pairs)
	vector.Shift(shiftedPairs, pairedCentroid)

	// compute cross-covariance matrix
	sourceMatrix := vector.Matrix(shiftedSource)
	pairsMatrix := vector.Matrix(shiftedPairs)
	crosscovariance := &mat.Dense{}
	crosscovariance.Mul(sourceMatrix, pairsMatrix.T())
	return crosscovariance
}

// rotation matrix is is R=VU^T
func rotationOf(crossCovariance mat.Matrix) mat.Matrix {
	// perform SVD
	svd := &mat.SVD{}
	svd.Factorize(crossCovariance, mat.SVDFull)
	var u, v mat.Dense
	svd.UTo(&u)
	svd.VTo(&v)
	rotation := &mat.Dense{}
	rotation.Mul(&u, v.T())
	return rotation
}

// translation is C_s - RC_m
func translationOf(sourceCentroid, matchedCentroid vector.Vector, rotation mat.Matrix) vector.Vector {
	col := matchedCentroid.Matrix()
	var mult mat.Dense
	mult.Mul(rotation, col)
	v := vector.Subtract(sourceCentroid, vector.FromMatrix(&mult))
	return v
}
